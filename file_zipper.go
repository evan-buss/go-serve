package main

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// ZipWriter takes a directory, creates a zip file, and returns the path
func ZipWriter(directory string) string {
	strings.LastIndex(directory, "/")
	dirName := directory[strings.LastIndex(directory, "/")+1:]
	zipPath := os.TempDir() + "/" + dirName + ".zip"

	// Get a Buffer to Write To
	outFile, err := os.Create(zipPath)
	if err != nil {
		fmt.Println(err)
	}
	defer outFile.Close()

	// Create a new zip archive.
	w := zip.NewWriter(outFile)

	// Add some files to the archive.
	addFiles(w, directory+"/", "")

	if err != nil {
		fmt.Println(err)
	}

	// Make sure to check the error on Close.
	err = w.Close()
	if err != nil {
		fmt.Println(err)
	}
	return zipPath
}

// Recursively walk through files to zip entire directories and subdirectories
func addFiles(w *zip.Writer, basePath, baseInZip string) {
	// Open the Directory
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		// fmt.Println(basePath + file.Name())
		if !file.IsDir() {
			dat, err := ioutil.ReadFile(basePath + file.Name())
			if err != nil {
				fmt.Println(err)
			}

			// Add some files to the archive.
			f, err := w.Create(baseInZip + file.Name())
			if err != nil {
				fmt.Println(err)
			}
			_, err = f.Write(dat)
			if err != nil {
				fmt.Println(err)
			}
		} else if file.IsDir() {

			// Recurse
			newBase := basePath + file.Name() + "/"
			// fmt.Println("Recursing and Adding SubDir: " + file.Name())
			// fmt.Println("Recursing and Adding SubDir: " + newBase)

			addFiles(w, newBase, baseInZip+file.Name()+"/")
		}
	}
}
