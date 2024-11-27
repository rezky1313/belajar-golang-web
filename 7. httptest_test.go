package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
HTTP Test
-Golang sudah menyiapkan package khusus untuk melakukan test pada fitur web
-Semuanya ada didalam package net/http/httptest/
-dengan menggunakan package ini kita bisa melakukan test terhadap aplikasi web tanpa harus menjalankan servernya
-kita bisa langsung fokus terhadap handler function yang ingin kita test

httptest.NewRequest()
-NewRequest(method,url,body) merupakan function yang digunakan untuk membuat http.Request
yang sebelumnya kita membuat httprequest melalui browser
-kita bisa menentukan method, url, body yang  akan kita kirim seagai simulasi unit test
-kita juga bisa menambahkan informasi tambahan lainnya pada request yang ingin kita kirim
seperti header, cookie dll

httptest.NewRecorder
-merupakan function yang digunakan untuk membuat response recorder atau writer
-RespoonseRecorder merupakan struct bantuan untuk merekam HTTP Response dari hasil testing yang dilakukan
sama artinya dengan response writer, conotnya kita bisa menangkan response bodynya apa, reaponse statusnya apa

*/

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Hello world")
}

func TestHelloHandler(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}

/*
-membuat function handler yang berisikan tulisan "Hello World" dengan parameter/kontrak writer dan request
-lalu print ke writer tulisan "Hello world"
-membuat function test dengan naman TestHello Handler, yang bertujuan untuk melakukan test terhadap simulasi response dan request
-jika sebelumnya kita membuat server untuk menguji write/response dan request dengan httptest itu tidak diperlukan
-cukup dengan menggunakan httptest.newRequest dan htttptest.NewRecorder
-Membuat variable Request dengan isi httptest.NewRequest
-lalu tentukan method, target URL, dan body nya
-menentukan method bisa dengan mengetikan secara manual atau bisa dengan http.methodGet dll
-lalu menentukan target nya dengan mengetiikan "localhost:8080/hello" ini sama dengan setting address pada server
-body disetting nil sana karna pada saat ini belum dibutuhkan
-membuat variable recorder dengan isi httptest.NewRecorder
-setelah itu panggil handler y ang sudah  dibuat tadi dengan memasukkan requst dan recorder sebagai parameter untuk dikirim
-lalu lakukan pengecekan apakah hasil test berjalan dengan benar
-pada recorder kita bisa menangkap appaun yang dibutuhkan
-untuk melihat hasil test bisa menggunakan recorder.result
-membuat variable response yaang berisikan recorder.Result
-pada ressponse kita akan mengambil body
-body merupakan byte array jadi haus dikonvert dulu
-untuk membaca response bisa dengan helper io.ReadAll
-test ini bisa menggunakan assertion jika diperlukan untuk cek apakah data yang keluar sudah benar atau belum

-jadi tidak perlu running server untuk test web server sudah berjalan atau belum
*/
