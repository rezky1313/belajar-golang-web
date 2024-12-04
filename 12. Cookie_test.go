package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
Stateless
-HTTP merupakan stateless antara server dan client, artinya server tidak akan menyimpan data apapapun requst dari client
-Hal ini dilakukan untuk scalability dari sisi server
ketika membuautu/running web pasti akan membuat/menggunakan banyak server, jadi belum tentu satu request masuk ke server yang sama
terus menerus, dan juga server tidak mungkin akan mengingat satu request ini sebelumnya datang dari server yang mana
jika server harus mengingat request yang sama maka request harus masuk ke server yang sama terus menerus, itu tidak efisien
tidak scalable
-bagaimana cara server mengingat sebuah client? misalnya ketika satu client sudah login maka client tidak perlu login lagi
-untuk melakukan ini bisa memanfaatkan cookie

Cookie
-Cookie adalah fitur HTTP dimana server bisa memberi response cookie berisi (key - value) dan client akan menyimpan cookie tersebut di browser
-pada request selanjut nya client akan selalu membawa cookie tersebut secara otomatis
jadi contohnya ketika client sukses login maka server akan mengiri data cookie berupa "login success" dan browser akan menyimpan
pada request kedua maka server tidak perlu meminta login pada client
-dan server secara otomatis akan selalu menerima data cookie yang dibawa client setiap melakukan request

Cara Membuat cookie
-Cookie adalah data yang dibuat oleh server dan sengaja agar disimpan oleh browser
-untuk membuat cookie pada server, bisa menggunakan function http.SetCookie()
*/

func SetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie := new(http.Cookie)
	//menentukan nama cookie
	cookie.Name = "X-Rezky-Name"
	//menentukan value cookie, contoh disini menggunakan value nama
	cookie.Value = request.URL.Query().Get("name")
	// menentukan cookie menempel pada url apa, jadi cookie bisa tentukan cookie aktif pada url yang mana saja
	//defaultnya cookie di letakkan pada url pada handler, conmtoh pada handler login
	//jadi cookie aktif pada saat login saja
	cookie.Path = "/"

	//setelah memnuat cookie kita bisa set cookie dengan menggunakan function http.SetCookie
	http.SetCookie(writer, cookie)
	fmt.Fprint(writer, "Sukses create cookie")
	//cookie coleh lebih dari satu, jadi bisa membuat cookie sebanyak2 nya tapi dianjurkan secukupnya saja
	//karena cookie akan disimpan dalam browser dan dikirimkan lagi ke server, kakn membuat operasi menjadi lambat

}

// membaca/mengambil cookie
func GetCookie(writer http.ResponseWriter, request *http.Request) {
	//cookies untuk mengambil data cookie berupa slices
	//cookie untuk mengambil berdasarkan nama
	cookie, err := request.Cookie("X-Rezky-Name")
	if err != nil {
		fmt.Fprint(writer, "No Cookie")
	} else {
		fmt.Fprintf(writer, "Hello %s", cookie.Value)
	}

	//jika terjadi error maka cookie tidak ada
}

func TestCookie(t *testing.T) {
	//membuat dummy server

	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)
	//mux merupakan handler multiplexer untuk handler yang sudah dibuat sebelumnya
	//dengan handler mux bisa memuat beberapa handler
	//berbeda dengan server test yang langsung dibuat kedalam satu url penuh

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

// membnuuat unit test sset cookie, simulasi membuat/mendapatkan cookie dari server
func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?name=Rezky", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	//membaca cookies
	cookies := recorder.Result().Cookies()
	for _, cookie := range cookies {
		fmt.Printf("%s : %s\n", cookie.Name, cookie.Value)
	}
}

// simulasi mendapatkan cokoie dari server
func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080/", nil)
	cookie := new(http.Cookie)
	cookie.Name = "X-Rezky-Name"
	cookie.Value = "Rezky Wahyudi"
	request.AddCookie(cookie)
	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
