package controllers

import (
	"fmt"
	"net/http"

	"caseMajoo/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Selamat Datang API Test Case Majoo")

}

func (server *Server) Deret(w http.ResponseWriter, r *http.Request) {
	var x, next int
	var pertama, kedua int = 0, 1
	fmt.Printf("\nMasukkan jumlah Value: ")
	fmt.Scanf("%d", &x)

	fmt.Printf("\nInputan deret Pertama: ")
	fmt.Scanf("%d", &pertama)

	fmt.Printf("\nInputan deret Kedua: ")
	fmt.Scanf("%d", &kedua)

	selisih := kedua - pertama

	for i := 0; i < (x); i++ {
		fmt.Printf("%d ", pertama)
		next = pertama + selisih
		pertama = next
	}
	fmt.Printf("\n")
}
