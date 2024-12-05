package belajar_golang_web

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

/*
-jika sebelumnya dengann menggunakan embed semua file terbawa semua ketika compile
-kadang ada kasus misal kita hanya ingin menggunakan static file sesuai yang kita inginkan
-hal ini bisa dilakukan menggunakan function http.servefile()
-dengan menggunakan function ini, kita bisa menentukan file mana yang ingin kita tulis ke http response
*/

func ServeFile(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		http.ServeFile(writer, request, "./resources/ok.html")
	} else {
		http.ServeFile(writer, request, "./resources/notFound.html")
	}
}

func TestServeFileServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

/*
-ketika parameter tidak ada maka yang ditampilkan halaman not found
-jika ada parameter maka tampilkan fil ok html
*/

/*
menggunakan golang Embed
-paramereter function http.ServeFile hanya berisi string file name, sehingga tidak bisa menggunakan golang embed
karenam embed menggunakan file system, bukan file name
-namun bukan berarti kita tidak bisa menggunakan golang embed, karena jika untuk melakukan load file,
kita hanya butuh menggunakan package fmt dan responsewriter saja tidak perlu menggunakan ServeFile
*/

//go:embed resources/ok.html
var resourcesOk string

//go:embed resources/notfound.html
var resourcesNotFound string

func ServefileEmbed(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		fmt.Fprint(writer, resourcesOk)
	} else {
		fmt.Fprint(writer, resourcesNotFound)
	}
}

func TestServeFileServerEmbed(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServefileEmbed),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
