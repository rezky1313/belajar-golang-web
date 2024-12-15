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
-saat ingin membuat web menggunakan data dinamis banyak tempat tapi datanya berberda2
-utnuk membuat data dinamis bisa dengan menggunakan data struct atau map
-Namun perlu dilakukan perubahan pada template nya, kita perlu memberi tahu field atau key mana
yang akan digunakan untuk mengisi data dinamis pada template
-untuk mennyebut kann field bisa seperti ini {{.FieldName}}
*/

func TemplateDataMap(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))

	t.ExecuteTemplate(writer, "name.gohtml", map[string]interface{}{
		"Title": "Template Data struct",
		"Name":  "Rezky",
		"Address": map[string]interface{}{
			"Street": "Padank",
		},
	})
	//menggunakan map string sbg key dan interface sbg value/data
	//menggunakan interface dikarenakan tipe data bisa berupa string, int dll
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

type Address struct {
	Street   string
	City     string
	Province string
	PostCode int
}

type File struct {
	Title   string
	Name    string
	Address Address
}

func TemplateDataStruct(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))

	t.ExecuteTemplate(writer, "name.gohtml", File{
		Title: "Template Data Menggunakan Struct",
		Name:  "Rezky Wahyudi",
		Address: Address{
			Street: "Jalan Bandar Buaat",
		},
	})
}
func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
