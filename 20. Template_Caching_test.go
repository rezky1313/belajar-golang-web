package belajar_golang_web

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

/*
Temppalte caching
-kode sebelumnya sebenarnya tidak efisien
-dikarenakan setiap handler dipalnggil kita melakukan parsing ulang temaplate terus menerus
-idealnya tempalte hanya diparsing 1 kali di awal, selanjutnya data tempalte disimpan di memory
*/

//go:embed templates/*.gohtml
var templatesCaching embed.FS

// memparsing semua templates menggunakan go embed dan memasukkan kedalam variable global
// var myTemplates = template.Must(template.ParseFS(templatesCaching, "templates/*.gohtml"))
var myTemplates = template.Must(template.ParseFS(templatesCaching, "templates/*.gohtml"))

// handler tidak memerlukan parsing lagi
func templateCaching(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "simple.gohtml", MyPage{
		Name: "aaaaa",
	})
}

func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	templateCaching(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
