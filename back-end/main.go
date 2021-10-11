package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	//Configuracion de CORS
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Access-Control-Allow-Origin", "Accept", "Accept-Language", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		Debug:            true,
	})
	r.Use(cors.Handler)

	// Ruta de Index
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		//Respuesta
		w.Write([]byte("INDEX"))
	})

	//Obtiene la ruta raiz
	workDir, _ := os.Getwd()
	//Hace referencia la ruta raiz concatenando la carpeta projects
	filesDir := http.Dir(filepath.Join(workDir, "projects"))

	// Funcion que define el path de las rutas
	FileServer(r, "/files", filesDir)

	//Puerto del servidor
	http.ListenAndServe(":3333", r)
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	//Verifica que no permita ingresar parametros en la ruta principal
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	//Revisa que el path de la ruta no sea "/" y que tampoco termine en "/"
	if path != "/" && path[len(path)-1] != '/' {

		//Ruta para subir archivos
		r.Post(path, func(w http.ResponseWriter, r *http.Request) { uploader(w, r) })

		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}

func uploader(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(5000)

	file, fileInfo, err := r.FormFile("file")

	f, err := os.OpenFile("./projects/"+fileInfo.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		log.Fatal(err)
		return
	}

	defer f.Close()

	io.Copy(f, file)

	fmt.Fprintf(w, fileInfo.Filename)
}
