package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

/*
Midleware
-Dalam pembuatan web ada konsep yang bernama midle ware atau filter atau interceptor
-midldeware adalah sebuah fitur dimana kita bisa melakukan sesuatu atau menambahkan kode sebelum atau sesudah handler
dieksekusi
-jadi kadang kita butuh melakukan sesuatu sebelum dan sesudah handler dieksekusi, biasanya setiap midleware melakukan
hal yang sama, contoh ketika daalam saebuah web ada fitur login, maka di awal harus ada validasi apaka user sudah
login atau belum, jika dilakuakn pada setiap handler akan membuat kode tidak efisien, maka dari itu digunakan midleware
agar pengecekan logon dilakukan satukali saja, semua handler harus melewati midleware dulu ketika akan dieksekusi

Diagram midleware

[request/response]<>--------<>[Midleware]<>-------<>[handler]

-user melakukan request
-request masuk ke  midleware terlebih dahulu
-dari midleware baru diteruskan ke handler
-response handler akan melewati midleware lagi hingga sampai ke user

Midleware di Golang Web
-golang tidak menyediakan midleware, di goang cuma ada handler, tidak ada konsep midleware
-namun karena struktur handler di golang menggunakan interface, kita bisa membuat midleware sendiri menggunakan handler
kita bisa membuat midleware sendiri asalkan sesuai dengan kontrak interface dari handler http.handler
*/

// membuat log midleware/ kontrak middleware
type logMiddleware struct {
	Handler http.Handler
}

// method dari logMiddleware dengan nama ServeHTTP
func (middleware *logMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Before Execute Handler") //bisa menambahkan logic apapun disini
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("After Execute Handler")
}

// membuat error Handler
type ErrorHandler struct {
	Handler http.Handler
}

func (handler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		err := recover()
		fmt.Println("Recover: ", err)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error: %s", err)
		}
	}()
	handler.Handler.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler Executed")
		fmt.Fprintf(writer, "Ini handler pertama")
	})

	mux.HandleFunc("/foo", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler Executed")
		fmt.Fprintf(writer, "Ini Foo handler")
	})
	//menambahkan error handler
	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler Executed")
		panic("Ups its not you, its from Us")
	})
	//diatas belum menggunakan error handler sehingga program berhenti

	// middleware := new(logMiddleware)
	// middleware.Handler = mux
	logMiddleware := &logMiddleware{
		Handler: mux,
	}

	errorHandler := &ErrorHandler{
		Handler: logMiddleware,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

/*
-Server runing, menerima request dan kirim ke logMiddleware yang berperan sebagai handler pada server
-pada saat logMiddleware berjalan maka kode/logic pada function ServeHTTP akan berjalan
-lalu logMiddleware menerima request dan mengirimkan ke mux, mux sebagai handler pada logMiddleware
-lalu mux akan menjalankan handler yang sebenarnya setelah melewati logMiddleware dan mengerjakan logika/kodenya

-server menerima request lalu mengarahkan ke logMidleware karna berperan sebahai handler pada server
-lalu logMiddleware menjalankan function method ServeHTTP
-pada method serveHTTP kode filter berjalan, mengerjakan logic diawal, lalu dilanjutkan dengan kode handler yang sebenarnya
-dilanjutkan dengan menjalankan handler mux yang ada pada middleware sesuai path url yang diminta
-lalu SErveHTTP menjalankan kode pada baris terakhir
*/

/*
error Handler
-middleware bisa juga digunakan untuk melakukan error handler
-dikarenakan middleware kita buat sendiri menggunakan interface maka error handler harus dibuat sendiri juga
-sehingga jika terjadi panic kita bisa recover di middlweware dan mengubah panic tersebut menjadi error response
-jadi program tidak berhenti mendadak tanpa pesan error
*/
