package main

import (
	"embed"
	"net/http"
)

//go:embed index.html
var content embed.FS

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Read the embedded file
        file, err := content.ReadFile("index.html")
        if err != nil {
            http.Error(w, "File not found", http.StatusNotFound)
            return
        }

        // Serve the HTML content
        w.Header().Set("Content-Type", "text/html")
        w.Write(file)
    })

    // Start the HTTP server
    http.ListenAndServe(":8080", nil)
}
