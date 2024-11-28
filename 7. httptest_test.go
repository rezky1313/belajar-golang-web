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
	fmt.Fprintln(writer, "Hello World")
}

func TestHelloHandler(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	response.Body.Close()

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

//cara kedua

func TestHttptest(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Worl")
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	response, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("Failed to Send Request: %v", err)
	}

	body, _ := io.ReadAll(response.Body)
	response.Body.Close()

	if string(body) != "Hello World" {
		t.Errorf("Expected 'Hello World' but got this: %s", string(body))
	}
}

//Cara Ketiga

func HelloHandler2(w http.ResponseWriter, r *http.Request) {
	//menetapkann header Content-Type
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, "Hello World")
}

func TestHttptest2(t *testing.T) {
	//Membuat Permintaan HTTP Get Ke Endpoint, pada method bisa diketik manual bisa juga dengan menggnakan http.methodGEt
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/Hello2", nil) //menambahkan http:// pada URL
	recorder := httptest.NewRecorder()                                                  //recorder untuk menangkap response

	//memanggil handler dengan recorder dan request sebagai parameter yang akan dikirim
	HelloHandler2(recorder, request)

	//mangambil response dari recorder
	response := recorder.Result()

	//membaca body response
	body, _ := io.ReadAll(response.Body)
	//disini bisa tambahkan error handling
	response.Body.Close()

	//memverifikasi isi body response
	if string(body) != "Hello World" {
		t.Errorf("Expected 'Hello World' but got this: %s", string(body))
	}
	//cetak response recorder
	cetakBody := string(body) //response.body mengembalikan []byte dan error, ognore error dan dikonvert []byte ke string
	fmt.Println(cetakBody)

	//verifikasi status kode
	if response.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code, got %d, want %d", response.StatusCode, http.StatusOK)
	}
	//cetak statusOK
	cetakStatus := response.StatusCode //menembalikan int jadi tidak perlu dikonvert
	fmt.Println(cetakStatus)

	//verifikasi header (opsional jika ada header khusus)
	contentType := response.Header.Get("Content-Type")
	ExpectedContentType := "text/plain; charset=utf-8"
	if contentType != ExpectedContentType {
		t.Errorf("Unexpected Content Type, got: %q, Want %q", contentType, ExpectedContentType)
	}
	//Cetak Header
	cetakHeader := response.Header.Get("Content-Type")
	fmt.Println(cetakHeader)
}

/*
pada kode diatas jika memeriksa content-type pastikan menyeting header pada handler
jika tidak maka golang akan otomatis menganggap bahwa content type text-plain dengan fotmat huruf kecil semua
*/

//https://chatgpt.com/share/67481f06-be74-8005-a3c7-ee13b4e6e90e
