package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

/*
Form Post
-Saat belajar HTML biasanya saat membuat form kita submit datanya menggunakan method GET atau POST
-perbedaannya jika menggunakan method GET Maka semua data di form akan menjadi/dikirim via query parameter ke server
-Sedangkan jika menggunakan method POST, maka semua ada form akan dikirm via body HTTP request ke server
-Di Golang unutuk mengambil data form POST sangat mudah

Request.PostForm
-Semua data yang dikirim dari client secara otomatis akan disimpan kedalam attribute request.PostForm
-berbeda dengan url kita tidak perlu melakukan parsing, pada url sudah langsung terdapat query parameter
-Namun sebelum mengambil data di attribute PostForm kita wajib memanggil method Request.ParseForm() terlebih dahulu
method ini digunakan untuk parsing data body apakah bisa diparsing menjadi form atau tidak, jika tidak bisa parsing
maka akan menyebabkan error
*/

func FormPost(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm() //parsing form menjadi body
	if err != nil {            //jika tidak bisa parsing maka error
		panic(err)
	}

	//otomatis parsing bisa dilakukan dengan
	//request.PostFormValue("firstName")

	firstName := request.PostForm.Get("firstName")
	lastName := request.PostForm.Get("lastName")

	fmt.Fprintf(writer, "%s %s", firstName, lastName) //print ke writer
}

func TestFormPost(t *testing.T) {
	//membuat simulasi request untuk body unutk testing terlebih dahulu
	requestBody := strings.NewReader("firstName=Rezky&lastName=Wahyudi")
	//request body isinya mirip dengan Query parameter hanya saja tidak dikirimkan melalui url, dikirim lewat body
	//membuatt simulasi http request
	request := httptest.NewRequest(http.MethodPost, "localhost:8080", requestBody)
	//membuat custom header untuk test
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//header sudah merupakan standard untuk form post jika berbeda maka data form POST tidak akan tampil
	//membuat recorder untuk server membaca data yang dikirimkan client
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

}
