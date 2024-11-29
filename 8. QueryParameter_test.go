package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
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
	request := httptest.NewRequest(http.MethodGet, "localhost:8080/catalogue?search=nike&page=2", nil)
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
