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

const DirPath = `c:\Go\src\`

type serverConfig struct {
	dirPath    string
	portNumber int
}

func main() {
	config := serverConfig{
		dirPath:    `c:\Go\src\`,
		portNumber: 3000,
	}

	r := mux.NewRouter()
	r.HandleFunc("/", fileHandler)
	r.HandleFunc("/files/{id}", fileIDHandler)
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	log.Printf("Server started at port %v", config.portNumber)
	log.Panic(http.ListenAndServe(fmt.Sprintf(":%v", config.portNumber), r))
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	ch, errCh := GetDir(DirPath)

	select {
	case dir := <-ch:
		t, _ := template.ParseFiles("views/index.html")
		t.Execute(w, dir)

	case <-errCh:
		w.WriteHeader(http.StatusExpectationFailed)
	}
}

func fileIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fileName := params["id"]
	p := path.Join(DirPath, fileName)

	w.Header().Set(
		"Content-Disposition",
		fmt.Sprintf("attachment; filename=\"%v.tar\"", fileName))

	archiver.Tar.Write(w, []string{p})

	log.Println(fmt.Sprintf("File served at [%v]", p))
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}
