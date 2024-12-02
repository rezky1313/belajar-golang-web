package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
Response code
-Dalam HTTP, Response code merupakan representasi kode response
-dari response code ini kita bisa mengetahui apakah request yang kita kirim ke server sukses diproses atau gagal
-ada banyak sekali response code dalam pembuatan web

response code dikelompokkan menjadi 5 bagian
1. informational response: 100 - 199
2. Successful request: 200 - 299
3. Redirect: 300 - 399
4. client Error: 400 - 499
5. Server Error: 500 - 599
Semuanya mengacu pada RFC 7231

Mengubah Response Code
-
-Secara default jika kita tidak menyebutkan response code maka response code nya adalah 200 OK
-jika ingin mengubahnya, kita bisa menggunakan function REsponseWriter.WriteHeader(int)
-Semua data status code sudah disediakan oleh golang, kita bisa menggunakan variable yang sidah disediakan jika ingin
 https://github.com/golang/go/blob/master/src/net/http/status.go

*/

func ResponseCodeHandler(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		//writer.WriteHeader(400) // bad request
		//jika takut salah menggunakan kode, bisa menggunakan variable yang sudak disediakan
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(writer, "Name is empty")
	} else {
		//writer.WriteHeader(200) //ini tidak perlu lagi di cek karena jika benar maka kode akan otomatis 200
		fmt.Fprintf(writer, "Hi %s", name)
	}
}

func TestResponseCode(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080/responsecode?name=", nil)
	//request.Header.Add("Content-Type/application")
	recorder := httptest.NewRecorder()

	ResponseCodeHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	defer response.Body.Close()
	fmt.Println(string(body))

	//statusCode := recorder.Result().StatusCode
	//fmt.Println(statusCode)
	fmt.Println(response.StatusCode) //print kode
	fmt.Println(response.Status)     // print reasponse code dengan stringnya

}
