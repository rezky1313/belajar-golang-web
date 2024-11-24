package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

/*
Handler
-Server hanya bertugas sebagai web server, sedangkan untuk menerima HTTP Request yang masuk ke server, kita butuh handler
-Handler di golang direpresentasikan dalam interface, dimana didalam kontraknya terdapat sebuah function bernama ServeHTTP()
yang digunakan sebagai function yang akan dieksekusi ketika menerima HTTP Request

type Handler interfece {
	ServeHTTP(ResponseWriter, *Request)
}

-Pada interface Handler Terdapat methosd bernama SarveHTTP
-pada method  terdapat 2 parameter ResponseWriter dan *Request
-jadi jika menggunakan handler wajib mengikuti kontrak yang ada pada interface handler

HandleFunc
-untungnya ada banyak implementasi dari handler, contohnya handler function/handlerfunc
jadi tidak perlu buat handler manual
-kita bisa menggunakan handlerfunc untuk membuat funtion handler HTTP

type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w,r)
}

-jadi handler function adalah sebuah tipe sebagai alias untuk function method pada interface handller
-
*/

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		//logic Web
		fmt.Fprint(writer, "Hello World")
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

/*
-membuat variable dengan tipe handlerfunc yang beriisikan anonymous function dengan parameter writer dan request
-Writer merupakan response yang akna diberikan kepada client
-sedangkan request adalah requst yang diterima dari client
-jika ingin mengembalikan response dengan menggunakan writer
-dengan cara writer.Write hanya saja itu membutuhkan binary atau byte arary
-agar mudah kita bisa menggunbakan fmt format, jika sebelumnya kita sering menggunakan println atau printf
sekarang bisa menggnakan Fprint, didalam fprint ada parameter dan data yang akan dikirim ke writernya
jadi kita tidak perlu lagi konversi data menjadi byte secara manual

*/
