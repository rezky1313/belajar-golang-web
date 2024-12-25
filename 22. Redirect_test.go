package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

/*
Redirect
-Saat membuat website kita pasti selalu membuat redirect
-misalnya setelah login, akan selalu redirect ke halaman dashboard
-redirect sudah memiliki standard response status code
-redirect sebenarnya adalah response status code 301 atau 302
-301 adalah permanent redirect
-302 adalah temporary redirect
-kita bisa gunakan status code diatas dan menambah header location
-pada golang dipermudah dengan menggunakan http.Redirect
*/

func RedirectTo(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Redirect Ke Halaman Dashboard")
}

func RedirectFromYow(writer http.ResponseWriter, request *http.Request) {
	http.Redirect(writer, request, "/redirect-to", http.StatusTemporaryRedirect)
}
func RedirectOut(writer http.ResponseWriter, request *http.Request) {
	http.Redirect(writer, request, "https://www.programmerzamannow.com/", http.StatusTemporaryRedirect)
}

/*
-menggunakan http redirect lalu mengisi writer dengan pesan redirect ke halaman dashboard
-pada function redirect from pada http redirect kita masuukan writer, request, lokasi redirect dan status code
-jika halama redirect tetap pada website yang sama, masukkan path saja
-jika redirect ke website lain, masukkan full url
*/

func TestRedirect(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect-from", RedirectFromYow)
	mux.HandleFunc("/redirect-to", RedirectTo)
	mux.HandleFunc("/redirect-out", RedirectOut)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
