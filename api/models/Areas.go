package models

import "github.com/jinzhu/gorm" // tambahkan import untuk menggunakan method yg ada pada gorm

type (
	Area struct {
		ID        int64
		AreaValue float64 //tipe data diubah menjadi float64 dikarenakan terdapat proses yg menghasilkan bilangan desimal
		AreaType  string
	}
)

func (_r *Area) InsertArea(param1 float64, param2 float64, AreaType string, DB *gorm.DB, ar *Area) (err error) { //*AreaRepository diubah menjadi *Area, tipe data pada param1 diubah menjadi float64 karena harus disesuaikan dengan apa yg ada di struct,  parameter type diubah menjadi AreaType, tambahkan DB *gorm.D untuk menggunakan method DB, *ModelArea diubah menjadi *Area
	inst := DB.Model(&Area{}) //_r dihapus dan ar diubah menjadi &Area{}
	var area float64          //penulisan Var seharusnya menggunakan var atau area := 0 dan tipe data diubah menjadi float dikarenakan terdapat proses yg menghasilkan angka desimal
	// area = 0  tidak perlu di deklarasikan karena pada line 15 secara default nilai area = 0
	switch AreaType {
	case "persegi panjang": //penulisan string menggunakan "" bukan ``
		area = param1 * param2 //karena area sudah di declare pada line 15, jadi tidak perlu menggunakan var atau :=
		ar.AreaValue = area
		ar.AreaType = "persegi panjang" //penulisan string menggunakan "" bukan ``
		err = inst.Create(&ar).Error    //diawal penulisan create menggunakan huruf besar menjadi Create
		if err != nil {
			return err
		}
	case "persegi": //penulisan string menggunakan "" bukan ``
		var area = param1 * param2 //penulisan variabel area tidak perlu menggunakan var, karena sudah di declare sebelumnya pada line 15
		ar.AreaValue = area
		ar.AreaType = "persegi"      //penulisan string menggunakan "" bukan ``
		err = inst.Create(&ar).Error //diawal penulisan create menggunakan huruf besar menjadi Create

		if err != nil {
			return err
		}
	case "segitiga": //karena type data nya string maka penulisan segitiga menggunakan ""
		area = 0.5 * (param1 * param2)
		ar.AreaValue = area
		ar.AreaType = "segitiga"     //penulisan string menggunakan "" bukan ``
		err = inst.Create(&ar).Error //diawal penulisan create menggunakan huruf besar menjadi Create
		if err != nil {
			return err
		}

	default:
		ar.AreaValue = 0
		ar.AreaType = "undefined data" //penulisan string menggunakan "" bukan ``
		err = inst.Create(&ar).Error   //diawal penulisan create menggunakan huruf besar menjadi Create
		if err != nil {
			return err
		}
	}

	return err //tambahkan return
}
