package musicorganizer

import (
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/JosephCaillet/boris/fs"
	"github.com/JosephCaillet/boris/log"
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
		log.Mode("Preview mode")
	} else {
		if config.Move {
			log.Mode("Move mode")
		} else {
			log.Mode("Copy mode")
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
	return path.Clean(config.MusicOut + "/" + newPath), nil
}

func reorganizeFiles(exploredPathes *[]fs.FilePathInfos) error {
	nonMusicFile := make([]string, 0)
	lastDestinationDir := config.MusicOut
	musicFoundInDir := false
	log.StartOperation(len(*exploredPathes))

	for i, filePathInfo := range *exploredPathes {
		log.ProgressOperation()
		if filePathInfo.FileInfo.IsDir() {
			musicFoundInDir = false
			log.EnteringFolder(filePathInfo.FullPath)
			continue
		}

		newPath, err := computeNewFilePath(filePathInfo.FullPath)
		if _, ok := err.(readTagsError); ok {
			if IsFileOpenable(filePathInfo.FullPath) {
				log.ErrorTag(err)
			}
			nonMusicFile = append(nonMusicFile, filePathInfo.FullPath)
		} else if err != nil {
			return fmt.Errorf("computing new path: %v", err)
		} else {
			musicFoundInDir = true
			lastDestinationDir = path.Dir(newPath)

			log.MoveFile(path.Base(filePathInfo.FullPath), newPath, true)

			if !config.Preview {
				if err = reorganizeFile(filePathInfo.FullPath, newPath); err != nil {
					return err
				}
			}
		}

		if len(nonMusicFile) != 0 &&
			(i+1 < len(*exploredPathes) && (*exploredPathes)[i+1].FileInfo.IsDir()) ||
			i+1 == len(*exploredPathes) {
			if !musicFoundInDir {
				log.WarnWrongMove()
			}
			for _, srcPath := range nonMusicFile {
				if musicFoundInDir {
					newPath = lastDestinationDir + "/" + path.Base(srcPath)
				} else {
					newPath = config.UnorganizedFiles + "/" + strings.Join(strings.Split(srcPath, "/")[1:], "/")
				}
				log.MoveFile(path.Base(srcPath), newPath, false)
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
	if err := os.MkdirAll(path.Dir(newPath), os.ModePerm); err != nil {
		return fmt.Errorf("creating folder: %v", err)
	}
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
