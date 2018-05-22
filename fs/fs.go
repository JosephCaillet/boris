package fs

import (
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strings"
)

type FilePathInfos struct {
	FullPath string
	os.FileInfo
}

func Tree(rootPath string) ([]FilePathInfos, error) {
	var filePathInfo []FilePathInfos
	rootPath = path.Clean(rootPath)

	rootStat, err := os.Stat(rootPath)
	if err != nil {
		return nil, err
	}
	filePathInfo = append(filePathInfo, FilePathInfos{FullPath: rootPath, FileInfo: rootStat})

	subFilePathInfo, err := tree(rootPath)
	if err != nil {
		return nil, err
	}
	filePathInfo = append(filePathInfo, subFilePathInfo...)
	return filePathInfo, nil
}

func tree(rootPath string) ([]FilePathInfos, error) {
	var filePathInfo []FilePathInfos
	files, err := ioutil.ReadDir(rootPath)
	if err != nil {
		return nil, err
	}

	sort.Slice(files, func(i int, j int) bool {
		if files[i].IsDir() != files[j].IsDir() {
			return !files[i].IsDir() && files[j].IsDir()
		}
		return strings.ToLower(files[i].Name()) < strings.ToLower(files[j].Name())
	})

	for _, file := range files {
		filePathInfo = append(filePathInfo, FilePathInfos{FullPath: rootPath + "/" + file.Name(), FileInfo: file})
		if !file.IsDir() {
			continue
		}
		subFilePathInfo, err := tree(rootPath + "/" + file.Name())
		if err != nil {
			return nil, err
		}
		filePathInfo = append(filePathInfo, subFilePathInfo...)
	}
	return filePathInfo, nil
}
