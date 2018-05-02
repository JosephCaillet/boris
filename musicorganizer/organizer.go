package musicorganizer

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/dhowden/tag"
)

func Reorganize() error {
	err := filepath.Walk(config.MusicIn, reorganizeFile)
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", config.MusicIn, err)
	}

	return nil
}

func reorganizeFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		return fmt.Errorf("failure accessing a path %q: %v", path, err)
	}

	if info.IsDir() {
		return nil
	}

	fmt.Printf("visited file: %s\n", path)
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failure opening file %q: %v", path, err)
	}
	defer file.Close()

	tags, err := tag.ReadFrom(file)
	if err != nil {
		fmt.Printf("failure opening tags of file %q: %v", path, err)
		return nil
	}

	tmpl, err := template.New("TreeTemplate").Parse(config.TreeTemplate)
	if err != nil {
		return fmt.Errorf("failure parsing templagte: %v", err)
	}

	var pathBuilder strings.Builder
	err = tmpl.Execute(&pathBuilder, newMetadata(tags, path))
	if err != nil {
		return fmt.Errorf("failure executing template: %v", err)
	}
	newPath := strings.Replace(pathBuilder.String(), "\n", "", -1)
	fmt.Printf("%s\n\n", newPath)

	return nil
}
