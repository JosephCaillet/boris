package musicorganizer

import (
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"
	"time"

	"github.com/JosephCaillet/boris/fs"
	"github.com/cleversoap/go-cp"
	"github.com/dhowden/tag"
)

type readTagsError struct {
	error
	filepath string
}

func (e readTagsError) Error() string {
	return fmt.Sprintf("opening tags of file %q: %v", e.filepath, e.error)
}

var tmpl *template.Template

func Reorganize() error {
	var err error
	tmpl, err = template.New("TreeTemplate").Parse(config.TreeTemplate)
	if err != nil {
		return fmt.Errorf("parsing template: %v", err)
	}

	if config.Preview {
		fmt.Println("⏺  Preview mode ⏺")
	} else {
		if config.Move {
			fmt.Println("⏺  Move mode ⏺")
		} else {
			fmt.Println("⏺  Copy mode ⏺")
		}
	}

	fmt.Print("\n🔎  Listing files...\n\n")
	exploredPathes, err := fs.Tree(config.MusicIn)
	if err != nil {
		return err
	}

	err = reorganizeFiles(&exploredPathes)
	if err != nil {
		return fmt.Errorf("error reorganizing file: %v", err)
	}

	if !config.Preview && config.DeleteMusicIn {
		if err = os.RemoveAll(config.MusicIn); err != nil {
			return fmt.Errorf("delete input music directory: %v", err)
		}

	}

	return nil
}

func computeNewFilePath(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("opening file %q: %v", filePath, err)
	}
	defer file.Close()

	tags, err := tag.ReadFrom(file)
	if err != nil {
		return "", readTagsError{filepath: filePath, error: err}
	}

	var pathBuilder strings.Builder
	err = tmpl.Execute(&pathBuilder, newMetadata(tags, filePath))
	if err != nil {
		return "", fmt.Errorf("executing template: %v", err)
	}

	newPath := strings.Replace(pathBuilder.String(), "\n", "", -1)
	newPath = strings.Replace(newPath, "\t", "", -1)
	return config.MusicOut + "/" + newPath, nil
}

func reorganizeFiles(exploredPathes *[]fs.FilePathInfos) error {
	start := time.Now()
	pathesNb := len(*exploredPathes)
	nonMusicFile := make([]string, 0)
	lastDestinationDir := config.MusicOut
	musicFoundInDir := false

	for i, filePathInfo := range *exploredPathes {
		progressPrefix := fmt.Sprintf("[ %d%% ][ %s ]\t",
			int(float32(i+1)/float32(pathesNb)*100.0),
			time.Since(start).Round(time.Second),
		)

		if filePathInfo.FileInfo.IsDir() {
			musicFoundInDir = false
			fmt.Printf("%s\n%s↳ %s/\n", progressPrefix, progressPrefix, filePathInfo.FullPath)
			continue
		}

		newPath, err := computeNewFilePath(filePathInfo.FullPath)
		if _, ok := err.(readTagsError); ok {
			fmt.Printf("%s\t❌ error: %v\n", progressPrefix, err)
			newPath = lastDestinationDir + "/" + path.Base(filePathInfo.FullPath)
			nonMusicFile = append(nonMusicFile, filePathInfo.FullPath)
		} else if err != nil {
			return fmt.Errorf("computing new path: %v", err)
		} else {
			musicFoundInDir = true
			lastDestinationDir = path.Dir(newPath)

			fmt.Printf("%s\t♫ %s\t➜\t%s\n", progressPrefix, path.Base(filePathInfo.FullPath), newPath)

			if !config.Preview {
				if err = os.MkdirAll(lastDestinationDir, 0777); err != nil {
					return fmt.Errorf("creating folder: %v", err)
				}
				if err = reorganizeFile(filePathInfo.FullPath, newPath); err != nil {
					return err
				}
			}
		}

		if len(nonMusicFile) != 0 &&
			(i+1 < len(*exploredPathes) && (*exploredPathes)[i+1].FileInfo.IsDir()) ||
			i+1 == len(*exploredPathes) {
			if !musicFoundInDir {
				fmt.Printf("%s\t⚠ No tagged music file found, moving file(s) below to last computed location.\n", progressPrefix)
			}
			for _, srcPath := range nonMusicFile {
				newPath = lastDestinationDir + "/" + path.Base(srcPath)
				fmt.Printf("%s\t🖺 %s\t➜\t%s\n", progressPrefix, path.Base(srcPath), newPath)
				if !config.Preview {
					if err = reorganizeFile(srcPath, newPath); err != nil {
						return err
					}
				}
			}
			nonMusicFile = nil
		}
	}

	return nil
}

func reorganizeFile(oldPath, newPath string) error {
	if config.Move {
		if err := os.Rename(oldPath, newPath); err != nil {
			return fmt.Errorf("moving file: %v", err)
		}
	} else {
		if err := cp.Copy(oldPath, newPath); err != nil {
			return fmt.Errorf("copy file: %v", err)
		}
	}
	return nil
}
