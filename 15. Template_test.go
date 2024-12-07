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
Web dinamis
-sampai pada saat sekarang ini hanya mmebahas tentang membuat response menggunakan string dan static file
-pada kenyataannya, saat membuat web pasti akan membuat halaman yang dinamis, bisa berubah2 sesuai data yang diakses oleh user
-di golang terdapat fitur HTML Template, yaitu fitur template yang bisa digunakan untuk membuat HTML yang dinamis

HTML Template
-fitur HTML terdapat pada package HTML Template tidak perlu library/framework
-sebelum menggunakan HTML Template, perlu terlebih dahulu membuat template nya
konsepnya berbeda dengan PHP, jika dengan php bisa langsung menggnakan html didalma phpnya
kalau golang konsepnya templating, jadi memisahkan logic kode golang dengan templatenya
ini bagus unutk program dengan skala besar, karena terpisah jadi tidak membingungkan, apa iya?
karena dipisah akan lebih mudah memaintain kode programnya
-Template bisa berupa file atau string
-bagian dinamis pada HTML template adalah bagian yang menggunakan  tanda {{}}


Membuat Template
-saat membuat template dengan string, kita perlu memberi tahu nama template nya
jika membuat template dengan file, maka nanti nama templatenya akan otomatis sama dengan nama filenya
-untuk membuat text template, cukup dengan text html, dna untuk konten yang dinamis cukup tambahkan {{.}}
contohnya:
<html><body>{{.}}</body></html>
*/

func SimpleHtml(writer http.ResponseWriter, request *http.Request) {
	templateText := `<html><body>{{.}}</html></body>`
	// t, err := template.New("SIMPLE").Parse(templateText)
	// if err != nil {
	// 	panic(err)
	// }

	//menggunakan cara otomatis untuk pengecekan error
	t := template.Must(template.New("SIMPLE").Parse(templateText))

	t.ExecuteTemplate(writer, "SIMPLE", "Hello HTML Template")
}

/*
-jadi setelah template text dibuat dalam bentuk file/text harus melakukan pasring template
-caranya dengan menggunakan package html/template.New dilanjutkan dengan nama "SIMPLE" (jika template menggunakna string) template
lalu gunakan function parse lalu masukkan stirng templatenya yang sudah dibuat sebelumnya
-untuk eksekusi template bisa menggunakan executeTemplate lalu masukkan writernya, lalu masukkan nama templatenya
nama template harus dimasukkan lagi karena nantinya bisa ada banyak nama template yang digunakan
lalu diikuti data dinamisnya, otomatis titik didalam kurung kurawal tadi akandisubtitusi kedalam kurung kurawal tadi

*/

func TestTemplateHtml(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "https:localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHtml(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/*
Template dari File
-Selain menggunakan string seperti diatas, template bisa juga dibuat menggunakan file
-hal ini mempermudah karena bisa langsung bisa membuat file html
-nama template = nama file
-biasanya untuk template di golang menggnakna gohtml, aga text editor tau kalau itu adalah template golang.

*/

func SimpleHTMLFile(writer http.ResponseWriter, request *http.Request) {
	// t, err := template.ParseFiles("./templates/simple.gohtml")
	// if err != nil {
	// 	panic(err)
	// }

	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))

	t.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Template")
}

func TestSimpleHTMLFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "https://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLFile(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/*
Template Directory
-kadang biasanya kita jarang sekali menyebutkan nama template satu per satu
jika dalam folder template ada ratusa file maka akan di sebutkan satu persatu maka itu akan sangat menyulitkan
-alangkah baiknya jika template disimpan kedalam satu file directory, lalu golang akan load semua data
yang ada pada folder template
-Golang template mendukung proses load template dari directory
-hal ini mempermudah karena tidak perlu load template sat per satu
-menggnakan parseGlob("./templates/*.gohtml")
*/

func TemplateDirectory(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))

	t.ExecuteTemplate(writer, "simple.gohtml", "Template HTML")
}

func TestTemplateDirectory(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDirectory(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

/*
Menggunakna golang Embed
-Golang embed direkomendasikan untuk menyimpan data template
-dengan goalng embed tidak perlu lagi mengCopy template file lagi, karena sudah otomatis diembed kedalam distribution file
pada saat complie program
*/

//go:embed templates/*gohtml
var templates embed.FS

func TemplateEmbed(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFS(templates, "templates/*gohtml") //tidak menggunkan titik/slash
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(writer, "simple.gohtml", "HTML Template") //pada nama bisa dipilih sesuai isi file dirctory
}

func TestTemplateEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
