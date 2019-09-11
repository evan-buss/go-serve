package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// Show the file and directory views.
func browseFilesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// The new path is the base path + the requested path
		// The server accepts paths relative to the base
		currentPath := ensureDirPath(basePath + r.URL.Path)

		// Ensure that the path exists on the system
		dir, err := ioutil.ReadDir(currentPath)
		if err != nil {
			http.Error(w, "Invalid location", http.StatusInternalServerError)
			return
		}

		// Convert os.FileInfo to fileInfo
		contents := make([]FileInfo, 0)

		for _, val := range dir {
			if !*showDots && val.Name()[0] == '.' {
				continue
			}

			temp := FileInfo{
				Info:         val,
				Link:         ensureDirPath(r.URL.Path) + val.Name(),
				DownloadLink: "/download" + ensureDirPath(r.URL.Path) + val.Name(),
				FDate:        val.ModTime().Format(time.UnixDate),
				DirSize:      -1,
			}

			if temp.Info.IsDir() {
				temp.DirSize = getDirSize(currentPath + temp.Info.Name())
			}

			contents = append(contents, temp)
		}

		headerName := strings.Split(r.URL.Path, "/")

		// Execute template with newly created Directory object
		err = tmpl.Execute(w, Directory{
			Name:         headerName[len(headerName)-1],
			Path:         basePath,
			RelativePath: r.URL.Path,
			Nav:          generateNav(r.URL.Path),
			Contents:     contents,
		})

		if err != nil {
			log.Fatal("Template Error: " + err.Error())
		}
	}
}

func downloadHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		remove := false
		path := basePath + r.URL.Path[len("/download"):]

		// Make sure the file exists
		info, err := os.Stat(path)
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		// Zip directories and set path to zip file created in /tmp
		if info.IsDir() {
			remove = true
			path = ZipWriter(path)
		}

		// Send the file
		http.ServeFile(w, r, path)

		// Only remove the temp zip file
		if remove {
			os.Remove(path)
		}
	}
}
