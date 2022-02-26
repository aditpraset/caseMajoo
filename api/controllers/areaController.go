package controllers

import (
	"caseMajoo/api/models"
	"errors"
	"log"
	"net/http"
)

func (server *Server) Area(w http.ResponseWriter, r *http.Request) {

	var area models.Area                                        //declare variable area terhadap struct Area pada model
	err := area.InsertArea(10, 10, "persegi", server.DB, &area) //Penulisan persegi menggunakan "", _u.Repository diubah menjadi area karena saya mendeklarasikan Struct pada variable area, deklarasi variable err menggunakan := dikarenakan sebelumnya tidak ada, setelah persegi sya menambahkan parameter server.DB, &area karena harus disesuakan dengan model yg sudah dibuat
	if err != nil {
		// log.Error().Msg(err.Error()) diubah pada line 16
		log.Fatal(err.Error())
		err = errors.New("ERROR_DATABASE") //en dihapus lalu "ERROR_DATABASE" diubah menjadi string
		return                             //err dihapus karena return yg di inginkan kosong
	}

}
