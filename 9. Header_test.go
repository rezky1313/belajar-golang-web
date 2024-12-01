package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
Header
-Selain query dalam HTTP ada yang namanya header
-header  adalah informasi tambahan yang bisa dikirim dari client ke server atau sebaliknya
-jadi header tidak hanya ada pada HTTP request, pada HTTP response pun kita juga bisa menambahkan informasi header
-saat menggunakan browser biasanya header ditambahkan secara otomatis oleh browser, seperti infomasi browser, content type
yang dikirm dann diterima browser dll
-untuk menangkap request header yang dikirim oleh client kita bisa mengambilnya dengan request.header
-header mirip dengan query parameter isinya map[string][]string
-query parameter case sensitive sedangkan header tidak
contoh nya jika browser mengirimkan data header huruf kecil semua sedangkan kita menangkapnya
dengan huruf besar semua itu tidak masalah
*/

func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("content-type")
	fmt.Fprint(writer, contentType)
}

func TestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080/header", nil)
	//1. simulasi mengirim custom content type dari browser/client ke server
	request.Header.Add("content-type", "aplication/json")
	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	//2. mengirim header otomatis dari client
	//header := response.Header.Get("content-type")
	//fmt.Println(header)
	fmt.Println(string(body))
}
