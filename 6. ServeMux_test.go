package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

/*
ServeMux
-Saat membuat web biasanya kita ingin membuat banyak sekali endpoint URL
misalnya kita ingin membuat endpoint login, product, report sesuai dengan aplikasi web yang kita buat
-HandlerFunc sayangnya tidak mendukung itu, jadi denganb menggunakan handlerfunc kita hanya bisa menggunakan
satu endpoint saja yaitu nama domain saja
-Alternative implementasi dari handler adalah ServeMux
-ServeMux adalah implementasi dari handler yang mendukung multiple endpoint
karena handler dibuat dalam bentuk interface makan bisa menggunakan banyak implementasi
contohnya adalah servemux dan handlerfunc yang dibahas sebelumnya
-Servemux sebenarnya adalah handler yang digunakan untuk menggabungkan banyak handler lainnya
jadi kita bisa membuat banyak handler dan digabungkan kedalam servemux
-saat digabungkan kedalam servemux kita bisa tentukan di handler A untuk endpoint mana dan handler B untuk endpoint mana

*/

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprint(writer, "Hello World")
	})

	mux.HandleFunc("/hi", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprint(writer, "Hi")
	})

	mux.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Images")
	})

	mux.HandleFunc("/images/thumbnails", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "thumbnails")
	})

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
Penjelasan
-membuat variable mux dengan tipe http.serverMux
-dari var mux kita membuuat 2 buah handler/handlerfunc yang satu mempunyai endpoint default
dan yang satu mempunyai endpoint "/hi"
-lalu membuat server dan mengisi handler nya dengan mux yang sudah didefinisikan diawal
-jadi ketika mengetikkan URL default, kaan muncul halaman dengan handler default yang sudah disetting
pada handler pertama, yaitu mencetak tulisan "Hello World"
-ketika mengetikkan URL kedua yang mempunya endpoint "/hi", maka itu secara otomatis akan ditangani
oleh handler dengan endpoint yang sudah disetting "/hi" dan akan muncul tulisan "Hi"
-jadi servemux disini sama dengan router pada bahasa pemograman lain
-jadi bisa membuat banyak sekali handler dan di merge kedalam satu ServeMux
-hal yang perlu diingat ketikan menuliskan endpoint yang sama maka salah satu akan tertimpa, jadi endpoint itu unique
*/

/*
URL Pattern
-URL Pattern dalam serveMux itu sederhana, tinggal menambahkan string yang ingin digunakan sebagai endpoint
tanpa harus memasukkan domain web
-jika URL Pattern dalam serveMux ditambahkan dengan garis miring, maka url tersebut akan menerimma path dengan awalan tersebut
misa /images/ artinya akan dieksekusi jika endpointnya /images/, /images/contoh dan /images/contoh/lagi
-Namun jika terdapat URL Patterm yang lebih panjang maka akan diprioritaskan yang lebih panjang
misal jika terdapat URL /images dan /images/thumbnails, maka jika mengakses /images/thumbnails akan mengakses
/images/thumbnails, bukan /images
*/
