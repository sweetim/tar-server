package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/mholt/archiver"
	"github.com/sweetim/tar-server/util"
)

type serverConfig struct {
	dirPath    string
	portNumber int
}

func main() {
	config := serverConfig{
		dirPath:    util.GetEnv("DIR_PATH", "/data").(string),
		portNumber: util.GetEnv("PORT", 3000).(int),
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
		ch, errCh := util.GetDir(config.dirPath)

		select {
		case dir := <-ch:
			t, err := template.New("index.gohtml").
				Funcs(
					template.FuncMap{
						"UnitSuffix":    util.UnitSuffix,
						"BoolMapString": util.BoolMapString,
						"IndexColor":    util.IndexColor,
					}).
				ParseFiles("views/index.gohtml")

			if err != nil {
				panic(err)
			}

			err = t.Execute(w, struct {
				DirPath string
				DirInfo []util.DirInfo
			}{
				DirPath: config.dirPath,
				DirInfo: dir,
			})

			if err != nil {
				panic(err)
			}

		case e := <-errCh:
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte("Please set the DIR_PATH correctly\n"))
			w.Write([]byte(e.Error()))

			log.Println(e)
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

		err := archiver.DefaultTar.Create(w)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}

		defer archiver.DefaultTar.Close()

		err = filepath.Walk(p, func(path string, info os.FileInfo, err error) error {
			file, err := os.Open(path)
			if err != nil {
				return err
			}

			defer file.Close()

			fileInfo, err := file.Stat()
			if err != nil {
				return err
			}

			archiver.DefaultTar.Write(archiver.File{
				FileInfo: archiver.FileInfo{
					FileInfo:   fileInfo,
					CustomName: path[len(filepath.Dir(p)):],
				},
				ReadCloser: file,
			})

			return nil
		})

		if err != nil {
			http.Error(w, err.Error(), 500)
		}

		log.Println(fmt.Sprintf("File served at [%v]", p))
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}
