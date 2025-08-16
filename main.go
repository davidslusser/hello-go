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
	<!DOCTYPE html>
	<html>
	<head>
		<meta charset="UTF-8">
		<title>Hello Go!</title>
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.4.2/css/all.min.css">
		<style>
			body {
				background: linear-gradient(135deg, #00769c 0%, #007d9c 40%, #50b7e0 70%, #5dc9e2 100%);
				color: #fff;
				font-family: Arial, sans-serif;
				min-height: 100vh;
				margin: 0;
				display: flex;
				justify-content: center;
				align-items: center;
			}
			.center-container {
				display: flex;
				flex-direction: column;
				align-items: center;
				justify-content: center;
			}
			ul { list-style-type: none; padding: 0; }
			li { 
				margin-bottom: 24px;
				text-align: center; 
			}
			h1 { color: #fff; text-align: center; }
			strong {
				color: #fddd00;
			}
			.go-icon {
				font-size: 64px;
				margin-bottom: 20px;
				color: #fddd00;
			}
		</style>
	</head>
	<body>
		<div class="center-container">
			<div class="go-icon">
				<i class="fa-brands fa-golang"></i>
			</div>
			<h1><strong>Hello Go!</strong></h1>
			<ul>
				<li><strong>Hostname:</strong><br/> {{.Hostname}}</li>
				<li><strong>Start Time:</strong><br/> {{.StartTime}}</li>
				<li><strong>Image Tag:</strong><br/> {{.ImageTag}}</li>
			</ul>
		</div>
	</body>
	</html>
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
