package views

import (
  "html/template"
  "path/filepath"
  "net/http"
)

var LayoutDir string = "views/layouts"

func NewView(layout string, files...string) *View {
  files = append(files, layoutFiles()...)
  t, err := template.ParseFiles(files...)
  if err != nil {
    panic(err)
  }

  return &View{
    Template: t,
    Layout:   layout,
  }
}

type View struct {
  Template *template.Template
  Layout   string
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
  w.Header().Set("Content-Type", "text/html")
  return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  if err := v.Render(w, nil); err != nil {
    panic(err)
  }
}

func layoutFiles() []string {
  files, err := filepath.Glob(LayoutDir + "/*.gohtml")
  if err != nil {
    panic(err)
  }
  return files
}
