package musicorganizer

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
	"time"

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

type filePathInfos struct {
	fullPath string
	os.FileInfo
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

	exploredPathes := make([]filePathInfos, 0)
	fmt.Print("\n🔎  Listing files...\n\n")

	err = filepath.Walk(config.MusicIn, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("accessing a path %q: %v", filePath, err)
		}

		exploredPathes = append(exploredPathes, filePathInfos{fullPath: filePath, FileInfo: info})
		return nil
	})
	if err != nil {
		fmt.Printf("❌ error walking the path %q: %v\n", config.MusicIn, err)
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

func reorganizeFiles(exploredPathes *[]filePathInfos) error {
	start := time.Now()
	pathesNb := len(*exploredPathes)
	nonMusicFile := make([]string, 0)
	var lastDestinationDir string

	for i, filePathInfo := range *exploredPathes {
		progressPrefix := fmt.Sprintf("[ %d%% ][ %s ]\t",
			int(float32(i+1)/float32(pathesNb)*100.0),
			time.Since(start).Round(time.Second),
		)

		if filePathInfo.FileInfo.IsDir() {
			fmt.Printf("%s\n%s↳ %s/\n", progressPrefix, progressPrefix, filePathInfo.fullPath)
			continue
		}

		newPath, err := computeNewFilePath(filePathInfo.fullPath)
		if _, ok := err.(readTagsError); ok {
			fmt.Printf("%s\t❌ error: %v\n", progressPrefix, err)
			newPath = lastDestinationDir + "/" + path.Base(filePathInfo.fullPath)
			nonMusicFile = append(nonMusicFile, filePathInfo.fullPath)
		} else if err != nil {
			return fmt.Errorf("computing new path: %v", err)
		} else {
			lastDestinationDir = path.Dir(newPath)

			fmt.Printf("%s\t♫ %s\t➜\t%s\n", progressPrefix, path.Base(filePathInfo.fullPath), newPath)

			if !config.Preview {
				if err = os.MkdirAll(lastDestinationDir, 0777); err != nil {
					return fmt.Errorf("creating folder: %v", err)
				}
				if config.Move {
					if err = os.Rename(filePathInfo.fullPath, newPath); err != nil {
						return fmt.Errorf("moving file: %v", err)
					}
				} else {
					if err = cp.Copy(filePathInfo.fullPath, newPath); err != nil {
						return fmt.Errorf("copy file: %v", err)
					}
				}
			}
		}

		if len(nonMusicFile) != 0 &&
			(i+1 < len(*exploredPathes) && (*exploredPathes)[i+1].FileInfo.IsDir()) ||
			i+1 == len(*exploredPathes) {
			for _, srcPath := range nonMusicFile {
				newPath = lastDestinationDir + "/" + path.Base(srcPath)
				fmt.Printf("%s\t🖺 %s\t➜\t%s\n", progressPrefix, path.Base(srcPath), newPath)
			}
			nonMusicFile = nil
		}
	}

	return nil
}
