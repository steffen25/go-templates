package templates

import (
	"html/template"
	"log"
	"path/filepath"
	"net/http"
	"os"
	"fmt"
	"strings"
)

var templates map[string]*template.Template
const BASE_TEMPLATE_PATH = "/views/base.html"

func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	layouts, err := filepath.Glob(workDir + "/views/layouts/*.html")
	if err != nil {
		log.Fatal(err)
	}

	frontends, err := filepath.Glob(workDir + "/views/frontend/*.html")
	if err != nil {
		log.Fatal(err)
	}

	layoutsPlusBase := append(layouts, workDir + BASE_TEMPLATE_PATH)
	var files []string
	for _, frontend := range frontends {
		if len(files) == 0 || files == nil {
			// We only need to append the layoutsPlusBase once
			files = append(layoutsPlusBase, frontend)
		} else {
			files = append(files, frontend)
		}

		templates[filepath.Base(frontend)] = template.Must(template.ParseFiles(files...))
	}
}

func Render(w http.ResponseWriter, name string, data map[string]interface{}) error {
	tmpl, ok := templates[name]
	if !ok {
		return fmt.Errorf("The template %s does not exist.", name)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	baseTemplate := strings.Split(BASE_TEMPLATE_PATH, "/")
	baseTemplateName := baseTemplate[len(baseTemplate)-1]
	err := tmpl.ExecuteTemplate(w, baseTemplateName, data)
	if err != nil {
		return fmt.Errorf("Could not execute template err %v", err)
	}

	return nil
}