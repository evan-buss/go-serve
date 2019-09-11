package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

// Basepath is set to the current working directory
var basePath string
var tmpl *template.Template
var port = flag.String("port", "8080", "Change the server's port")
var unsafe = flag.Bool("unsafe", false, "Sets the base file server directory to the root '/' directory. Use with caution.")
var home = flag.Bool("home", false, "Sets the base file server directory to your home directory")
var showDots = flag.Bool("showDots", false, "Enable to make dotfiles visible")

func init() {
	flag.Parse()

	var err error
	if *unsafe {
		basePath = "/"
	}

	if *home {
		basePath, err = os.UserHomeDir()
		if err != nil {
			log.Fatal("Error setting base directory to home directory")
		}
	}

	if !*home && !*unsafe {
		basePath, err = os.Getwd()
		if err != nil {
			log.Fatal("Could not find current working directory.")
		}
	}

	tplFuncMap := make(template.FuncMap)
	tplFuncMap["Split"] = strings.Split
	tplFuncMap["getFormattedSize"] = getFormattedSize

	data, err := Asset("data/index.html")
	if err != nil {
		log.Fatal("Could not load html template from binary data")
	}

	// tmpl = template.Must(template.New("index.html").Funcs(tplFuncMap).ParseFiles("index.html"))
	tmpl = template.Must(template.New("index.html").Funcs(tplFuncMap).Parse(string(data)))

	// Find the file server's local IP adddress
	fmt.Println("Serving Files on: " + findIP())
}

func main() {
	http.HandleFunc("/download/", downloadHandler())
	http.HandleFunc("/", browseFilesHandler())
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
