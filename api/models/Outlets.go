package models

import (
	"caseMajoo/api/helpers"
	"time"

	"github.com/jinzhu/gorm"
)

type Outlets struct {
	ID         int64     `gorm:"primary_key;auto_increment" json:"id"`
	MerchantID int64     `json:"merchant_id"`
	OutletName string    `json:"outlet_name"`
	Merchant   Merchants `gorm:"ForeignKey:merchant_id" json:"merchant"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	CreatedBy  *int64    `json:"created_by"`
	Created    User      `gorm:"Foreignkey:id;association_foreignkey:created_by;" json:"created"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	UpdatedBy  *int      `json:"updated_by"`
}

type OutletReport struct {
	MerchantID   int64     `json:"merchant_id"`
	MerchantName string    `json:"merchant_name"`
	OutletName   string    `json:"outlet_name"`
	Date         time.Time `json:"date"`
	Omzet        float64   `json:"omzet"`
}

func (outlet *Outlets) PaginateOutletReport(begin, page, limit int, db *gorm.DB, userID int64, dates []Date) (int, []OutletReport) {
	var outlets []Outlets
	var outletReport OutletReport
	OutletReports := []OutletReport{}
	arrOutletReport := []OutletReport{}

	arrReport := make([]OutletReport, len(arrOutletReport))

	var start string
	var end string

	if dates != nil {
		for _, date := range dates {
			start = date.Start
			end = date.End
		}
	} else {
		start = "2021-11-01"
		end = "2021-11-30"
	}

	layoutFormat := "2006-01-02"

	startDate, _ := time.Parse(layoutFormat, start)
	endDate, _ := time.Parse(layoutFormat, end)

	for rd := helpers.RangeDate(startDate, endDate); ; {
		date := rd()
		if date.IsZero() {
			break
		}
		tanggal := date.Format("2006-01-02")

		db.Select("date(t.created_at)  as date, m.merchant_name, o.outlet_name, SUM(t.bill_total) as omzet").
			Table("Transactions as t ").
			Joins("join Merchants as m on t.merchant_id = m.id").
			Joins("join Outlets as o on t.outlet_id = o.id").
			Group("m.merchant_name, date, o.outlet_name").
			Where("m.user_id = ?", userID).
			Where("date(t.created_at) = ?", tanggal).
			Find(&arrOutletReport)

		if len(arrOutletReport) > 0 {
			outletReport = OutletReport{MerchantName: arrOutletReport[0].MerchantName, OutletName: arrOutletReport[0].OutletName, Date: arrOutletReport[0].Date, Omzet: arrOutletReport[0].Omzet}
			arrOutletReport = append(arrOutletReport, outletReport)
		} else {
			date, _ = time.Parse(layoutFormat, tanggal)

			db.
				Joins("join merchants on merchants.id = outlets.merchant_id").
				Where("merchants.user_id = ?", userID).
				Where("outlets.created_by = ?", userID).
				Preload("Merchant").
				Find(&outlets)

			outletReport = OutletReport{MerchantName: outlets[0].Merchant.MerchantName, OutletName: outlets[0].OutletName, Date: date, Omzet: 0}
			arrOutletReport = append(arrOutletReport, outletReport)
		}

		arrReport = append(arrReport, arrOutletReport[0])
	}
	max := limit * page
	for i := begin; i < max; i++ {
		OutletReports = append(OutletReports, arrReport[i])
	}
	count := len(arrReport)
	return count, OutletReports
}
