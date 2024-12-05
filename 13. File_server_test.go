package belajar_golang_web

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

/*
-Golang memiliki fitur yang bernaman file server
-dengan ini bisa membuat handler yang digunakan sebagai static file server
-dengan file server kita tidak perlu me-load file secara manual lagi
-fileServer adalh handler, jadi bisa di tambahkan kedalam http.server atau http.serveMux

cara memenggunakan fileServer
-tetukan dulu direktori/ lokasi dari lokasi file meenggunakan http.Dir masukan lokasi conoth pada folder (./resources)
-lalu buat variable file server dengan isi http.fileServer(direktori)
-buat server mux
-pada mux gunakan funciton Handle lalu masukkan pattern ("/static", fileServer)
-ketika membuka static nanti akan dihandle oleh file server
*/

func TestFileServer(t *testing.T) {
	directory := http.Dir("./resources")
	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()
	//*menghapus prefix pada url dengan stripe prefix
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

/*
-setelah dijalankan dengan url /static/index.html maka akn tampil error404 not found
-hal ini dikarenakan fileServer membaca url, lalu mencari file berdasarkan urlnya
jadi jika kita membuat static/index.html maka file server akan mencari ke file ./resources/static/index.html
-itu menyebabkan 404 error karenak tidak ada folder static dalam path
-oleh karena itu kita bisa menggunakan function http.StripPrefix() untuk menghapus prefix di URL
-jadi sebelum memnaggil urlnya dia akan menghapus prefix "static"
*/

/*
file server dengan golang embed
-Padan golang embed kita bisa embed file kedalam binary distribution file, hal ini  emmpermudah sehingga tidak
perlu meng-copy file lagi
-golang embed juga memiliki fitur yang bernama embed.Fs fitur ini diintegrasikan dengan fileServer
-jika tidak menggunakan embed, pada saat compile program, kita harus upload juga file resource jadi 2 kali kerja
-jika menggunakna embed semua file resource akan dimasukkan kedalam program
*/

//go:embed resources
var resources embed.FS

func TestFileServerGoEmbed(t *testing.T) {
	//tidak perlu menggunakan directory lagi (perlu jika menggunakan embed)
	//jika menggunakan embed harus menggunakan fs.sub()
	directory, _ := fs.Sub(resources, "resources")
	fileServer := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

/*
-ketika dijalankan terjadi lagi error 404
-karena di golang embed, nama folder ikut menjadi nama resource nya, misal resources/index.js
jadi untuk menagksesnya kita butuh gunakan url /static/resources/index.js
-jadi jika ingin langsung mengakses file index.js tanpa menggunakan resources, kita bisa menggunakan function
fs.Sub() untuk mendapatkan sub directory
*/
