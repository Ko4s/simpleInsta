package views

import (
	"html/template"
	"path/filepath"
)

// Template paths and extensions
const (
	TemplateDir       string = "./views/"
	LayoutDir         string = "./views/layouts/"
	TemplateExtension string = ".html"
)

// NewView is a View constructor
func NewView(layout string, files ...string) *View {
	appendTemplateAndExtensionToFilePath(files)
	files = append(files, layoutFiles()...)
	t := template.Must(template.ParseFiles(files...))

	return &View{
		Template: t,
		Layout: layout,
	}
}

// View is a struct for menaging and rendering pages(views)
type View struct {
	Template *template.Template
	Layout   string
}

func layoutFiles() []string {
	layoutFileNames, err := filepath.Glob(LayoutDir + "*" + ".html") // ./views/layouts/*.html"

	if err != nil {
		panic(err)
	}

	return layoutFileNames
}

// appendTemplateAndExtensionToFilePath take a slice of string and
// appends view folder file path and extension to given file name
func appendTemplateAndExtensionToFilePath(files []string) {
	for i, file := range files {
		files[i] = TemplateDir + file + TemplateExtension
	}
}
