package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

type handler struct {
	story Story
	tpl   *template.Template
}

func NewHandler(s Story) http.Handler {
	return handler{
		story: s,
		tpl:   template.Must(template.ParseFiles("cyoa/templates/layout.html")),
	}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	arc := "intro"
	arcParam := r.URL.Query().Get("arc")
	if arcParam != "" {
		arc = arcParam
	}
	chapter, ok := h.story[arc]
	if !ok {
		http.Error(w, "chapter not found", http.StatusNotFound)
		return
	}
	err := h.tpl.Execute(w, chapter)
	if err != nil {
		http.Error(w, "Sothing went wrong...", http.StatusInternalServerError)
	}
}

func main() {
	fileName := flag.String("json", "cyoa/stories.json", "json file name")
	flag.Parse()

	f, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	story, err := JsonStory(f)
	if err != nil {
		log.Fatal(err)
	}
	handler := NewHandler(story)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.ServeHTTP)

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mux)
}
