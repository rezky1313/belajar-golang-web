package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

/*
Download file
-Selain upload file kita juga ingin halaman web bisa mendonwnload sesuatu dari web
-sebenarnya digolang sudah disediakan fileServer dan serveFile, akan tetapi dengan menggunakan metode tersebut
kita akan langusng merender file download di browser
-jika kita ingin memaksa download file tanpa dirender oleh browser menggnakan header Content-Disposition
-fitur untuk memaksa file didownload sebenarnya bukan merulpakan fitur golang, akan tetapi fitur statndart HTTP
-http://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/content-Disposition
*/

// handler Donwnload
func Downloadfile(writer http.ResponseWriter, request *http.Request) {
	//masukkan quesry parameter file yang akan didwonload
	fileName := request.URL.Query().Get("file")
	//jika file name kosong maka akan ewrror bad request
	if fileName == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "BAD REQUEST")
		return
	}

	//jika file tidak kosong maka download file menggunakan file disposition
	writer.Header().Add("Content-Disposition", "attachment; filename=\""+fileName+"\"")
	http.ServeFile(writer, request, "./resources/"+fileName)
	//serve file artinya tampilkan file
	//sedangkan content disaposition artinya memindahkan file yang ditampilkan
}

func TestDownloadFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(Downloadfile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

/*
-pada browser kita menginputkan query parameter sesuai nama file, maka file akan langsung terdownload
-jika tidak menggunakan content disposition  maka file akan langsung dirender dibrowser

*/
