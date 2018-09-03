package views

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var mapTemplate = make(map[string]*template.Template)

func PopulateTemplate() {
	var allFiles []string
	templatesDir := "./files/www/html/"

	files, err := ioutil.ReadDir(templatesDir)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		fileName := file.Name()
		if strings.HasSuffix(fileName, ".html") {
			allFiles = append(allFiles, templatesDir+fileName)
		}
	}

	templates, err := template.ParseFiles(allFiles...)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	mapTemplate = map[string]*template.Template{
		"login": templates.Lookup("login.html"),
		"home":  templates.Lookup("home.html"),
	}
}

func GetMapTemplate() map[string]*template.Template {
	return mapTemplate
}
