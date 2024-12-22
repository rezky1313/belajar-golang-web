package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
XSS Cross Site Scripting
-XSS adalah security issue pada saat membuat web
-XSS adalah celah keamanan dimana seseorang bisa memasukkan parameter javascript agar dirender oleh halaman web
-biasanya tujuan dari XSS adalah untuk mencuri cookie broowser pengguna yang sedang mengakses web
-XSS bisa menyebabkan pengguna kehilangan akun

Auto Escape
-pada goalng template masalah XSS sudah teratasi secara otomatis
-Golang template memiliki fitur auto escape, dimana bisa mendeteksi mana data yang perlu ditampilkan di template
jika mengandung tag2 HTML atau script, secara otomatis akan diescape
-https://github/golang/go/blob/master/src/html/escape.go
-htts://golang.org/pkg/html/template/#hdr-Contexts
-jadi script akan otomatis diubah menjadi teks dan tida dirender sebagaimana mestinya

Mematikan auto escape
-bagaimana jika kita ingin benar2 memasukkan script ,html atau gambar?
-jika mau auto escape bisa juga dimatikan
-namun harus beritahu secara eksplisit ketikan menambahkan template data, bahwa data ini mengantung HTML
dan hars dirender sebagaimana mestinya
-bisa menggunakan data
-template.HTML , jika ini adalah HTML
-template.CSS , jika ini adalah CSS
-template.javascript , jika ini adalah javascript

-jika mematikan auto escape pastikan data bernar2 aman unutk dimatikan auto escape nya
*/

// auto escape
func TemplateAutoEscape(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Go-Lang Auto Escape",
		//	"Body":  template.HTMLEscapeString("<h2>Selamat Belajar Golang Web</h2>"),
		"Body": "<p>ini adalh body</p>",
	})
}

func TestTemplateAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscape),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
