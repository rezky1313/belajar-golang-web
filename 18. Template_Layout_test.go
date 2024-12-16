package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

/*
Template Layout
-Saat Membuat halaman website, ada beberapa bagian yang selalu sama seperti footer dan headder
-Best practicenya pada bagian yang selalu sama, seharusnya disimpan pada template terpisah
-di golang mendukung untuk import file template lain

Import Template
-{{template "nama template"}} artinya kita import template tanpa memberikan data apapun
-{[templatew "Nama tempalte" .Value]} import template dengan data, jika ingin kirim semua data cukup dengan "." saja

*membuat template header dan footer
*/

func TemplateLayout(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/layout.gohtml", "./templates/header.gohtml", "./templates/footer.gohtml"))
	/*
		harus memasukkan semua file yang akan diimport kedalam template layout
		disarankan menggnakan parseGlobe atau golang embed saja
	*/
	t.ExecuteTemplate(writer, "layout", map[string]interface{}{
		"Name":  "Muhammad Sumbul",
		"Title": "Template Layout",
	})
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/*
Template NAme
-kita biasanuya kletika menulisakan nama file template meggunakan .gohtml
-cara supaya tidak perlu menggunakan .gohtml adalah dengan menggunakan {{define "nama template"}}
lalu diakhiri dengan {{end}}

-dengan template menggunakan define kita bisa membuat 2 tempalte dalam 1 file
dengan wrap 2 template dalam 2 define
*/
