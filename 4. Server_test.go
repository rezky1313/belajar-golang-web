package belajar_golang_web

import (
	"net/http"
	"testing"
)

/*
-Server adalah struct yang terdapat pada package net/HTTP yang digunakan sebagai representasi web server digolang
-untuk membuat web kita wajib menggunakan server
-saat membuat server ada beberapa yang harus ditentukan seperti host dan port tempat dimana web berjalan
-jika kita mengakkses google sebernarnya web google berjalan pada port 80
-tapi disarankan untuk keperluan belajar tidak perlu menggunakan port 80
-karena dibeberapa sistem operasi diwajibkan sebagai administrator, jadi ketika menggunakan windows
pada saat running harus sebagai administrator direkomendasikan port 8080
-untuk host bisa menggunakan localhost saja, localhoist adalah penanda kita menjalankan server pada komputer sendiri
-Setelah membuat sever kita bisa menjalankan server tersebut dengan funciton ListenAndServe()
itu adalah method yang ada didalam server, tinggal panggil maka web server akan berjalan
*/

func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

/*
Handler
-Server hanya bertugas sebagai web server, sedangkan untuk menerima HTTP Request yang masuk ke server, kita butuh handler
-Handler di golang direpresentasikan dalam interface, dimana didalam kontraknya terdapat sebuah function bernama ServeHTTP()
yang digunakan sebagai function yang akan dieksekusi ketika menerima HTTP Request

type Handler interfece {
	ServeHTTP(ResponseWriter, *Request)
}

-Pada interface Handler Terdapat methosd bernama SarveHTTP
-pada method  terdapat 2 parameter ResponseWriter dan *Request
-jadi jika menggunakan handler wajib mengikuti kontrak yang ada pada interface handler

HandleFunc
-untungnya ada banyak implementasi dari handler, contohnya handler function/handlerfunc
jadi tidak perlu buat handler manual
-kita bisa menggunakan handlerfunc untuk membuat funtion handler HTTP

type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w,r)
}

-jadi handler function adalah sebuah tipe sebagai alias untuk function method pada interface handller
-
*/
