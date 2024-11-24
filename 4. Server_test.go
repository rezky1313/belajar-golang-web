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
