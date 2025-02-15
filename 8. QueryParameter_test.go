package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

/*
Query Parameter
-Query parameter adalah salah satu fitur yang digunakan dalam pembuatan web
-biasanya digunakan untuk mengirim data dari client ke server
-ditempatkan pada URL
-Untuk manambahkan query parameter, kita bisa gunakan ?namaQuery=valueQuery pada URLnya
-jika sebelumnya Pada HTML ketika menggunakan action method "GET" itu semua datanya menjadi query parameter
jadi kita bisa menangkap datanya dengan Query parameter pada Golang

url.URL
-query parameter sebenarnya datanya menempel di URL dan representasi URL pada golang adalah  url.URL
-Dalam parameter REquest terdaoat attribute URL yang berisi data url.URL
-Dari data URL ini kita bisa mengambil query parameter dari client dengan menggunakan method Query()
yang akan mengembalikan Map dimana key=QueryParameterName dan value=valueParameter


*/

/*
-ada sebuah function bernama SayHello yang mempunya writer dan request sebagai parameter
-membuat satu buah variable yang berisikan:
-untuk menggunakan query parameter cukup gunakan request.URL.Query.Get("name")
-.URL untuk mendapatkan data struct URLnya
-lalu gunakan .Query untuk mendapatkan data Querynya saja karena didalam URL terdapat banyak data seperti URL,port,domain
-ini akakn menggembalikan data query berupa Map
-lalu implementasi GET untuk mendapatkan query berdasarkan nama querynya
-lalu cek jika kita tidak mendapatkan data nama, maka print hello saja
-jika mendapatkan data nama print hello+nama
*/

func SayHello(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(writer, "Hello")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/sayhello?name=rezky", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
	response.Body.Close()

	status := recorder.Result().StatusCode
	if status != http.StatusOK {
		t.Errorf("Unexpected Status code: Want %v, Got %v", http.StatusOK, status)
	}
	fmt.Println(status)

}

/*
-jadi nantinya query parameter ini digunakan untuk firut seperti pencarian, pagination, sort data
-misalnya client mengetikkan data pencarian, data tersebut akan masuk ke url
-lalu client mengirimkan url ke server, lalu server akan memproses data apa yang ditampilkan sesuai query parameternya
-query parameter bisa multiple
-pada request diatas name sebagai key dan rezky sebagai query parameter valuenya
*/

/*
Multiple Query parameter
-untuk parameter lebih dari 1 contoh pada menggunaakn fitur filter pada parameter cukup dengan menggunakan "&"
-misalnya localhost:8080/katalog?search=nike&page=2
-lalu pisahkan method get nya antara parameter pertama dan kedua
-jika pada cara pertama digabung kedalam 1 line dan 1 variable sekarang dimasukkan kedalam 2 variable
-validasi pada key bisa dilakukan bisa tidak
-validasi dilakukan misalnya untuk input data penting nantinya akan menghasilkan "data tidak boleh kosong"
*/

func HandlerCatalogue(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query() //tidak menggunakan get karena akan dibagi menjadi 2

	querySearch := query.Get("search")
	queryPage := query.Get("page")

	//tanpa pengecekan/validasi
	fmt.Fprintf(writer, "Search: %v", querySearch)
	fmt.Fprintf(writer, "Page: %v", queryPage)

}

func TestMultiParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080/catalogue", nil)
	recorder := httptest.NewRecorder()

	HandlerCatalogue(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
	defer response.Body.Close()

	status := recorder.Result().StatusCode
	if status != http.StatusOK {
		t.Errorf("unexpected status code: Want %v, Got %v", http.StatusOK, status)
	}
	fmt.Println(status)
}

/*
Validasi Parameter
-digunakan untuk cek apakah parameter sesuai dengan yang diinginkan atau cek apakah parameter key ada atau tidak
-misalnya cek apakah query kosong atau memang tidak ada, query yang tidak boleh kosong/input penting,
sampai menentukan nilai default nilai default pada query parameter
*/

func HandlerValidasiParameter1(writer http.ResponseWriter, request *http.Request) {
	//validasi key paremeter ada dan tidak ada
	query := request.URL.Query()

	//pada pengecekan key "search" menggunakan "has"
	if query.Has("search") {
		fmt.Println("Key 'search' ditemukan")
	} else {
		fmt.Println("Key 'search' tidak ditemukan")
	}

	searchQuery := query.Get("search")
	//pageQuery := query.Get("page")

	fmt.Fprintf(writer, "Search: %v", searchQuery)
	//fmt.Fprintf(writer, "Page: %v", pageQuery)

}
func HandlerValidasiParameter2(writer http.ResponseWriter, request *http.Request) {
	//validasi apakah key ada tapi kosong dan key memang tidak ada serta jika query ditemukan
	query := request.URL.Query()

	if !query.Has("search") {
		fmt.Println("Query key search tidak ditemukan")
		// search := query.Get("search")
		// fmt.Fprintf(writer, "No query:%v", search)
	} else if query.Get("search") == "" {
		fmt.Println("query ditemukan tapi kosong")
		search := query.Get("search")
		fmt.Fprintf(writer, "Search:%v", search)
	} else {
		fmt.Println("Query search ditemukan")
		search := query.Get("search")
		fmt.Fprintf(writer, "search:%v", search)
	}
}
func HandlerValidasiParameter3(writer http.ResponseWriter, request *http.Request) {
	//validasi input data penting
	query := request.URL.Query()

	search := query.Get("search")
	if search == "" {
		//pesan error bisa menggnakan http.error atau t.errorf atau menggunakan print teks biasa
		//jika menggunakan error handling maka program akan terhenti karena menggunakan return dan test jadi gagal
		http.Error(writer, "Parameter tidak boleh kosong", http.StatusBadRequest)
		return
	} else {
		fmt.Fprintf(writer, "search:%v", search)
	}

}
func HandlerValidasiParameter4(writer http.ResponseWriter, request *http.Request) {
	//validasi menggunakan nilai default jika kosong
	query := request.URL.Query()

	search := query.Get("search")
	if search == "" {
		search = "adidas"
	}
	fmt.Fprintf(writer, "search: %v", search)
}

func HandlerMultipleParameterValues(writer http.ResponseWriter, request *http.Request) {
	//menambahkan banyak value ke query parameter
	//query sebenarnya merupakan map, jadi ketika menggunakan metode GET itu yang diambil data pertamanya
	query := request.URL.Query()

	names := query["names"]
	fmt.Fprint(writer, strings.Join(names, " ")) //menggabungkan string dengan menambahkan separator dibelakang
}

func TestParameterValidation(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080/validation?names=rezky&names=Wahyudi&names=kurniawan", nil)
	recorder := httptest.NewRecorder()

	//HandlerValidasiParameter4(recorder, request)
	HandlerMultipleParameterValues(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
	defer response.Body.Close()

	status := recorder.Result().StatusCode
	if status != http.StatusOK {
		t.Errorf("unexpected status code: want %v, got %v", http.StatusOK, status)
	}
	fmt.Println(status)
}
