package musicorganizer

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/dhowden/tag"
)

var (
	lastExploredDir string
	tmpl            *template.Template
	exploredPathes  []string
)

func Reorganize() error {
	var err error
	tmpl, err = template.New("TreeTemplate").Parse(config.TreeTemplate)
	if err != nil {
		return fmt.Errorf("failure parsing templagte: %v", err)
	}

	if config.Preview {
		fmt.Println("⏺ Preview mode ⏺")
		err = filepath.Walk(config.MusicIn, reorganizeFilePreview)
		if err != nil {
			fmt.Printf("❌ error walking the path %q: %v\n", config.MusicIn, err)
		}
	} else {
		fmt.Println("⏺ No preview ⏺")
		exploredPathes = make([]string, 1)
		err = filepath.Walk(config.MusicIn, explorePathes)
		if err != nil {
			fmt.Printf("❌ error walking the path %q: %v\n", config.MusicIn, err)
		}
		err = reorganizeFile()
		if err != nil {
			fmt.Printf("❌ error moving file: %v\n", err)
		}
	}

	return nil
}

func reorganizeFilePreview(filePath string, info os.FileInfo, err error) error {
	if err != nil {
		return fmt.Errorf("failure accessing a path %q: %v", filePath, err)
	}

	if info.IsDir() {
		return nil
	}

	currentDir := path.Dir(filePath)
	if currentDir != lastExploredDir {
		fmt.Printf("\n↳ %s/\n", currentDir)
	}
	lastExploredDir = currentDir

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failure opening file %q: %v", filePath, err)
	}
	defer file.Close()

	tags, err := tag.ReadFrom(file)
	if err != nil {
		fmt.Printf("\t❌ failure opening tags of file %q: %v\n", filePath, err)
		return nil
	}

	var pathBuilder strings.Builder
	err = tmpl.Execute(&pathBuilder, newMetadata(tags, filePath))
	if err != nil {
		return fmt.Errorf("failure executing template: %v", err)
	}
	newPath := strings.Replace(pathBuilder.String(), "\n", "", -1)
	newPath = strings.Replace(newPath, "\t", "", -1)
	fmt.Printf("\t%s\t→\t%s%s\n", path.Base(filePath), config.MusicOut, newPath)

	return nil
}

func explorePathes(filePath string, info os.FileInfo, err error) error {
	if err != nil {
		return fmt.Errorf("failure accessing a path %q: %v", filePath, err)
	}

	if info.IsDir() {
		return nil
	}

	exploredPathes = append(exploredPathes, filePath)
	return nil
}

func reorganizeFile() error {
	for _, filePath := range exploredPathes {
		fmt.Println(filePath)
	}

	return nil
}
