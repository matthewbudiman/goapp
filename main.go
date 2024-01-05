package main

import (
	"log"
	"net/http"
	"os"

	"github.com/thedevsaddam/renderer"
)

var rnd *renderer.Render

func init() {
	opts := renderer.Options{
		ParseGlobPattern: "./templates/*.html",
	}
	rnd = renderer.New(opts)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	port := ":9000"
	log.Println("Listening on port ", port)
	http.ListenAndServe(port, mux)
}

func home(w http.ResponseWriter, r *http.Request) {
	message := getEnv("MOTD", "Today Rocks!")
	motd := struct {
		MOTD string
	}{message}

	rnd.HTML(w, http.StatusOK, "home", motd)
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
