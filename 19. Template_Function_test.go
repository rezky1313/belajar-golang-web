package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"text/template"
)

/*
Template Function
-selain mengakses Field, didalma template juga bisa mengakses function
-cara emmanggil funciton sama dengan memmanggil field, hanya saja jika funciton memiliki parameter
maka tambahkan parameter diakhir setelah nama function
-{{.FunctionName}}, memnaggil field fucntionName atau fucntion functionNAme() tanpa parameter
-golang tempalte secara otomatis bisa mendeteksi yang mana field bisasa dan mana yang function
-{{.FcuntionName"param1","param2"}}, memanggil function dengan parameter

*/

type MyPage struct {
	Name string
}

// function method dari struct MyPage
func (myPage MyPage) SayHello(name string) string {
	return "Hello" + name + "Myname is" + myPage.Name
}

// disini kita tidak buat template dengan file, tapi langsung dengan "template.New"
func Templatefunction(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "Budi"}}`))

	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Rezky",
	})
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	Templatefunction(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/*
Global Fucntion
-Global function adalah function yang bisa kita definisikan sendiri
dan digunakan secara langsung pada golang template tanpa menggunakan template data
-jadi sebelumnya kita menggunakan template data jika ingin menggunakna function
-sebenarnya golang sudah memiliki beberapa function bawaan seperti yang digunakan pada template action
-https://github.com/goalng/go/blob/master/src/text/template/funcs.go

*/

// menggunakan global function len yang sudah built in di goalng
func TemplateGlobalFucntion(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{len .Name}}`))
	//pemanggilan  func  global tidak perlu menggunakan titik di awal nama function
	// t.ExecuteTemplate(writer, "FUCNTION", map[string]interface{}{
	// 	"Name": "Tutor Golang",
	// })
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Rezzky",
	})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateGlobalFucntion(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/*
Menambahkan Global Fcuntion
-untuk menambahkan global function bisa menggunakan method "funcs" pada template
-Perlu diingat, bahwa menambahkan global function harus dilakukan sebelum parsing template
jika menambahkan function global setelah parsing makan akan terjadi error karna goalng tidak mengenali function
-
*/

func TemplateFunctionGlobalCreate(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FUNCTION")
	//membuat global function global sendiri
	//dengan menggunakan map, lalu menggunakan anonymous function
	t = t.Funcs(map[string]interface{}{
		"upperyow": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	//yang akan jadi value untuk function upper yang didefine diatas adalah value .Name pada struct
	t = template.Must(t.Parse(`{{upperyow .Name}}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "mantap kan rizki",
	})
}

func TestTemplateFunctionGlobalCreate(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobalCreate(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/*
fcuntion pipelines
-function pipelines artinya hasil dari function dikirim ke function berikutnya
-untuk menggunakan funcitonpipelines bisa menggunakan tanda |
-{{SayHello .Name | upper }}, artinya hasil dari fucntion say hello akan dikirin ke function upper sebagai parameter
-bisa menambahkan function pipelines lebih dari satu
*/

func TemplateFunctionPipelines(writer http.ResponseWriter, request *http.Request) {
	t := template.New("INIFUNCTION")
	t = t.Funcs(map[string]interface{}{
		"SayHello": func(value string) string {
			return "Hello " + value
		},
		"upperyo": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse(`{{SayHello .Name | upperyo}}`))
	t.ExecuteTemplate(writer, "INIFUNCTION", MyPage{
		Name: "sherlocxk holmes",
	})
}

func TestTemplateFunctionPipelines(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionPipelines(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
