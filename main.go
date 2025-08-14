package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"
)

var (
	startTime = time.Now()
	imageTag  = os.Getenv("IMAGE_TAG")
)

type AppInfo struct {
	Hostname  string `json:"hostname"`
	StartTime string `json:"start_time"`
	ImageTag  string `json:"image_tag"`
}

func getAppInfo() AppInfo {
	hostname, _ := os.Hostname()
	return AppInfo{
		Hostname:  hostname,
		StartTime: startTime.Format(time.RFC3339),
		ImageTag:  imageTag,
	}
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := `
		<h1>App Info</h1>
		<ul>
			<li><strong>Hostname:</strong> {{.Hostname}}</li>
			<li><strong>Start Time:</strong> {{.StartTime}}</li>
			<li><strong>Image Tag:</strong> {{.ImageTag}}</li>
		</ul>
	`
	t := template.Must(template.New("page").Parse(tmpl))
	t.Execute(w, getAppInfo())
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getAppInfo())
}

func main() {
	http.HandleFunc("/", htmlHandler)
	http.HandleFunc("/api", jsonHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	fmt.Printf("Server starting on port %s...\n", port)
	http.ListenAndServe(":"+port, nil)
}
