package belajar_golang_web

/*
Golang Web
-Golang pada saat ini dijadikan dijadikan salah satu bahasa pemograman untuk membuat web terutama web API (backend)
-di Go sudah disediakan package untuk pembuatan web dan juga package unit testing untuk implementasi unit test untuk web
-hal ini memjadikan pembuatan web di golang menjadi lebih mudah karena tidak butuh library atau framework

diagram cara kerja golang web


											1. Request				2. eksekusi
		5. render[Web Browser]<=======================================>[golang] 3. render
											4. Response

Alur kerja
-Web browser akan mengirimkan request ke palikasi golang
-di golang sudah tersedia yang namanya web server jadi tidak perlu menginstal web server tambahan seperti PHP
-lalu aplikasi golang akan mengeksekusi request
-setelah eksekusi goang akan melakukan render contoh render HTML, CSS dan Javascript dll
-lalu akan dikembalikan sebagai response kepada web browser
-web browser akan melakukan render kembali  untuk menterjemahkan bahasa mesin yang di render oleh golang

Alur kerja Detail
-Web Browser melakukan HTTP Request ke web server
-Golang menerima request, lalu mengeksekusi request tersebut sesuai dengan yang diminta
-Setelah melakukan eksekusi request golang akan mengembalikan data dan dirender sesuai dengan kebutuhannya
misal HTML, CSS dll
-Golang akan mengembalikan hasil render tersebut sebagai HTTP Response ke Web Browser
-Web browser akan menerima content dari web server, lalu me render content tersebut sesuai tipe contentnya

Package Net/HTTP
Pada Golang untuk pembuata web golang menyediakan package net/HTTP
-pada beberapa bahasa pemograman lain, seperti java misalnya membuttuh kan library atau framework
-sedangkan digolang sudah disediakan secara built-in ole dev nya
-sehingga ketika pembuatan web tidak membutukan library, bisa nenggunakan package yang sudah tersedia
-Walaupun ketika membuat web dengan skala besar kita direkomendasikan menggunakan framework karena beberapa hal sudah
dipermudah oleh framework
-pada saat awal ini disarankan untu fokus kepada package net/HTTP untuk membnuat webnya, karena semua
framework golang menggunakan net/HTTP sebagai dasar pembuatan framework nya
*/
