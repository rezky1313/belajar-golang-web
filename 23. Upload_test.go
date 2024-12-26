package belajar_golang_web

/*
Upload file
-Saat membuat web, selain menerima input form dan query parameter, kada juga menerima input berupa file
dari client
-golang menyediakan fitur upload file

Multipart
-saat ingin menerima file maka harus melakukan parsing dulu menggnakan request.ParseMultipartForm(size)
atau juga bisa langsung mengambil data file nya menggunakan request.FormFile(name), didalam nya sebenarnya ada
proses parsing multipart form
-hasilnya merupakan data data yang terdapat pada package multipart, seperti multipart.file file sebagai
representasi file dari client dan multipart.FileHeader sebagai informasi file dari client
*/

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// handler tempalte form upload
func TemplateUploadForm(writer http.ResponseWriter, request *http.Request) {
	err := myTemplates.ExecuteTemplate(writer, "upload.form.gohtml", nil)
	if err != nil {
		panic(err)
	}
}

// handler upload success
func UploadSuccess(writer http.ResponseWriter, request *http.Request) {
	//mengambil file
	file, fileHeader, err := request.FormFile("file")
	if err != nil {
		panic(err)
	}

	//menyimpan file upload dari client ke directory
	//menentukan target penyimpanan
	//nama file nantinya akan disamakan dengan fileHeader dengan menggunakan fileHeader.FileName
	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}

	//setelah menentukan target penyimpanan maka sekarang pindahkan data upload ke directory target
	//data darii file akan dicopy ke fileDestination
	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}

	//mengambil inputan yang bukan file
	name := request.PostFormValue("name")
	myTemplates.ExecuteTemplate(writer, "upload.success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

func TestFormUploadServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", TemplateUploadForm)
	//handler untuk upload file
	//path harus disesuaikan dengan action template form
	mux.HandleFunc("/success", UploadSuccess)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

/*
penjabaran alur kode upload file diatas
-ada 2 handler, handler pertama adalah handler untuk menampilkan form upload yang kedua adalah handler untuk
menerima file upload , menyimpan dan menampilkan hasil upload
-handler pertama adalah TemplateUploadForm, yang akan menampilkan form upload
-handler kedua adalah UploadSuccess, yang akan menerima file upload, menyimpan
-pada handle uploadsucceess kita menggunakan request.FormFile("file") unutk mengambil file dari client
-fileHeader.Filename untuk mendapatkan informasi seperti nama, size dan mime typew file
-fileDestination, err := os.Create("./resources/" + fileHeader.Filename) untuk menentukan target penyimpanan
-menggunakan os create untuk membuat file baru lalu menentukan lokasi penyimpanan dan nama file
-untuk memindahkan data upload ke directory target, kita menggunakan io.Copy(fileDestination, file)
-parameter pertama adalah target lokasi penyimpanan, parameter kedua adalah data ayang akan disimpan
-untuk mengambil inputan yang bukan file kita bisa menggunakan request.postFromValue("name")
-untuk menampilkan hasil upload kita menggunakan myTemplates.executeTemplate(writer, "filetemplate", map)
*/

/*
Implementasi Unit Test
-
*/

// menggunakan Embed untuk mengambil contoh file yang akan diupload dan memasukkannya sebagai bytes
// lalu dimasukkan ke dalam veriable body
//
//go:embed resources/invasion_gig1.png
var uploadFileTest []byte

func TestUpload1(t *testing.T) {
	//untuk menyimpan body dalam bentuk binary
	body := new(bytes.Buffer)

	//merupakan simulasi input data
	//untuk membuat format multipart
	//memberi isian ke body
	writer := multipart.NewWriter(body)
	//simulasi input teks
	//memasukkan data yang bukan file dengan menggunakan field
	writer.WriteField("name", "Rezky Wahyudi")

	//simulasi uplaod file menggunakan Create FormFile
	//menggunakan parameter field name(nama field, sesuaikan dengan form upload) dan
	//file name kita custom, jadi setiap file yang terupload akan menggunakan nama tersebut
	file, _ := writer.CreateFormFile("file", "NewLogo.png")
	//file yang kana diupload
	file.Write(uploadFileTest)
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", body)
	//selalu set content type, jika tidak maka parsing akan gagal
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	UploadSuccess(recorder, request)

	body1, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body1))

}
