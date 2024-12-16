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
Tempalte Action
-golang Template mendukung perintah action, seperti percabangan, perulangan dan lain2


Pada template goalng memasukkan logika agak sedikit tricky, karena pada html memang tidak bisa
memasukkakn kode golang

If Else
-{{if .Value}} T1 {{end}}, Jika Value tidak kosong, maka T1 akan dieksekusi, jika kosong makan tidak ada yang dieksekusi
-{{if .Value}} T1 {{else}} T2 {{end}}, jika Value tidak kosong, maka T1 akan dieksekusi, jika kosong
maka T2 yang akan dieksekusi
-{{if .Value}} T1 {{else if .Value2}} T2 {{else}} T3 {{end}}, jika Value1 tidak kosong, T1 dieksekusi
jika T2 tidak kosong, T2 dieksekusi, jika tidak semuanya, T3 dieksekusi

-jadi menggunakan kata kunci if dilanjutkan dengan value yang akan dicek
-setelah value dicek maka eksekusi apapun yang ada setelah pengecekan dalam kondisi ini T1
-setelah itu gunakan kataq kunci end untuk mengakiri logika

*membuat template html
*/

type Address2 struct {
	Street  string
	City    string
	ZipCode string
}

type File2 struct {
	Name    string
	Title   string
	Address Address
}

func TemplateAction(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))

	t.ExecuteTemplate(writer, "if.gohtml", File2{
		//Name:  "Johnny Winter",
		Title: "Template Action Mf",
		Address: Address{
			Street: "22 Baker Street",
		},
	})
}

func TestTemplateAction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAction(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/*
Operator Perbandingan
-ketika butuh melakukan perbandingan pada if statement

-eq: Equal Arg1 == Arg2
-ne: Not Equal/tidak sama dengan Arg1 != Arg2
-lt: less than Arg1 < Arg2
-le: less than Equals Arg1 <= Arg2
-gt: greater than Arg1 > Arg2
-ge: greater than equal Arg1 >= Arg2

contoh
{{if ge .FinalValue 80}}
-artinya jika final value grater than equal 80, mkaa eksekusi program  dibawahnya
-operator memang diletakkan didepan value, Kenapa didepan?
-hal ini dikarenakan operator perbandingan tersebut sebenarnya adalah function
-jadi saat memanggil 	{{eq first second}} sebenarnya dia memanggil funtion eq dengan parameter
first dan second: eq(first/arg1, second/arg2)
*/

func TemplateActionOperator(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/operators.gohtml"))

	t.ExecuteTemplate(writer, "operators.gohtml", map[string]interface{}{
		"FinalValue": 50,
		"Title":      "Template Action Operator",
	})
}

func TestTemplateActionOperator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionOperator(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}

/*
Range
-Range digunakan untuk melakukan iterasi data template
-jika ingin mnampilkan data berupa table bisa menggunakan range ini
-tidak ada pruralngan for pada golang template
-yang bisa digunakan adalah iterasi dengan menggunakan range untuk tiap data pada array, slice, map, atau chanel
-{{range $index, $element := .value}} T1 {{end}}, jika value memmiliki data maka T1 akan dieksekusi sebanyak
-{{range $index, $element := .value}} T1 {{else}} T2 {{end}}, jika value memmiliki data maka T1 akan dieksekusi sebanyak
element value dan jika koosng maka T2 akan dieksekusi kita bisa gunakan $index untuk menagkses index
dan $element untuk mengakses element
-
*/
var Hobbies = [3]string{
	"memasak", "Band", "Golang",
}

func TemplateActionRange(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))

	t.ExecuteTemplate(writer, "range.gohtml", map[string]interface{}{
		"Hobbies": []string{
			// "Read", "Cook", "Code",
		},
	})
}

func TestTemplateActionRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionRange(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/*
template dengan With
-Kadang kita sering menggunakan nested struct
-jika menggunakan template biasanya harus menyebutkan value awal dan nested valuenya
-pada golang template terdapat action with, digunakan untuk mengubah scope dot menjadi object yang kita mau
-{{with .Value}} T1 {{end}}, value merupakan nested struct, jadi tidakl perlu menuliskan parentnya
*/

func TemplateActionwith(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/with.gohtml"))

	t.ExecuteTemplate(writer, "with.gohtml", map[string]interface{}{
		"Title": "template Golang With",
		"Name":  "Budi wicaksono",
		"Address": map[string]interface{}{
			"Street": "Mampang",
		},
	})
}

func TestTemplateActionWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionwith(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
