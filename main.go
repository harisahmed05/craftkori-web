package main

import (
    "html/template"
    "log"
    "net/http"
)

func main() {
    // Serve static files
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    tmpl := template.Must(template.ParseFiles("templates/index.html"))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        err := tmpl.Execute(w, nil)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    })

    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
