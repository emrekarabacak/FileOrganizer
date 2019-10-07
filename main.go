package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

var path = flag.String("src", "", "Source folder path")
var dest = flag.String("dst", "", "Destination folder path")
var copy = flag.Bool("copy", true, "Copies files if true, otherwise files will be moved")

func main() {
	flag.Parse()

	fmt.Println(*copy)

	if len(*path) == 0 {
		fmt.Println("source path is mandatory")
		return
	}

	if len(*dest) == 0 {
		fmt.Println("Destination path is mandatory")
		return
	}

	files := make([]fileInformation, 0, 15)
	retrieveFileInformation(*path, &files)
	groupedFiles := groupFilesByDate(&files)

	if *copy {
		copyToPath(groupedFiles, dest)
	} else {
		moveToPath(groupedFiles, dest)
	}
}

func moveToPath(result map[string][]fileInformation, dest *string) {
	for key, val := range result {
		path := fmt.Sprintf("%s/%s", *dest, key)
		os.Mkdir(path, os.ModePerm)

		for _, f := range val {
			currPath := fmt.Sprintf("%s/%s", f.dir, f.name)
			destPath := fmt.Sprintf("%s/%s", path, f.name)
			err := os.Rename(currPath, destPath)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
}

func copyToPath(result map[string][]fileInformation, dest *string) {
	for key, val := range result {
		path := fmt.Sprintf("%s/%s", *dest, key)
		os.Mkdir(path, os.ModePerm)

		for _, f := range val {
			inDir := fmt.Sprintf("%s/%s", f.dir, f.name)
			in, _ := os.Open(inDir)
			defer in.Close()

			outDir := fmt.Sprintf("%s/%s", path, f.name)
			out, _ := os.Create(outDir)
			defer out.Close()

			_, err := io.Copy(out, in)
			if err != nil {
				fmt.Println("Copying file failed ", f.name, err)
			}
		}
	}
}

func groupFilesByDate(fileInfoArray *[]fileInformation) map[string][]fileInformation {
	result := make(map[string][]fileInformation)
	for _, fileInfo := range *fileInfoArray {
		cd := fileInfo.createDate
		val := fmt.Sprintf("%d-%d", cd.Year(), int(cd.Month()))
		result[val] = append(result[val], fileInfo)
	}
	return result
}

func retrieveFileInformation(dir string, fileInfo *[]fileInformation) {
	files, _ := ioutil.ReadDir(dir)
	for _, f := range files {
		if f.IsDir() {
			dirPath := fmt.Sprintf("%s/%s", dir, f.Name())
			retrieveFileInformation(dirPath, fileInfo)
		} else {
			*fileInfo = append(*fileInfo, fileInformation{dir: dir, name: f.Name(), createDate: f.ModTime()})
		}
	}
}

type fileInformation struct {
	dir        string
	name       string
	createDate time.Time
}
