package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"

	"github.com/gorilla/mux"
	"github.com/mholt/archiver"
)

type serverConfig struct {
	dirPath    string
	portNumber int
}

type fileViewModel struct {
	DirPath string
	DirInfo []DirInfo
}

func main() {
	config := serverConfig{
		dirPath:    GetEnv("DIR_PATH", "").(string),
		portNumber: GetEnv("PORT", 3000).(int),
	}

	if config.dirPath == "" {
		log.Panic("Please set DIR_PATH environment variable")
	}

	r := mux.NewRouter()
	r.HandleFunc("/", fileHandler(&config))
	r.HandleFunc("/files/{id}", fileIDHandler(&config))
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	log.Printf("Server started at port %v", config.portNumber)
	log.Printf("Serving tar files from folder = %v", config.dirPath)
	log.Panic(http.ListenAndServe(fmt.Sprintf(":%v", config.portNumber), r))
}

func fileHandler(config *serverConfig) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ch, errCh := GetDir(config.dirPath)

		select {
		case dir := <-ch:
			t, _ := template.ParseFiles("views/index.html")

			t.Execute(w, fileViewModel{
				DirPath: config.dirPath,
				DirInfo: dir,
			})

		case <-errCh:
			w.WriteHeader(http.StatusExpectationFailed)
		}
	}
}

func fileIDHandler(config *serverConfig) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		fileName := params["id"]
		p := path.Join(config.dirPath, fileName)

		w.Header().Set(
			"Content-Disposition",
			fmt.Sprintf("attachment; filename=\"%v.tar\"", fileName))

		archiver.Tar.Write(w, []string{p})

		log.Println(fmt.Sprintf("File served at [%v]", p))
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}