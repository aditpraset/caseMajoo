package models

import (
	"caseMajoo/api/helpers"
	"time"

	"github.com/jinzhu/gorm"
)

type Merchants struct {
	ID           int64     `gorm:"primary_key;auto_increment" json:"id"`
	UserID       int64     `json:"user_id"`
	MerchantName string    `json:"merchant_name"`
	User         User      `gorm:"ForeignKey:user_id" json:"user"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	CreatedBy    *int64    `json:"created_by"`
	Created      User      `gorm:"Foreignkey:id;association_foreignkey:created_by;" json:"created"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	UpdatedBy    *int      `json:"updated_by"`
}

type MerchantReport struct {
	MerchantName string    `json:"merchant_name"`
	Date         time.Time `json:"date"`
	Omzet        float64   `json:"omzet"`
}

func (merchant *Merchants) MerchantReport(start, page, limit int, db *gorm.DB, userID int64, dates []Date) (int, []MerchantReport) {
	var merchants []Merchants
	var merchantReport MerchantReport
	arrMerchantReport := []MerchantReport{}
	MerchantReports := []MerchantReport{}
	arrReport := make([]MerchantReport, len(arrMerchantReport))

	var startRange string
	var endRange string

	if dates != nil {
		for _, date := range dates {
			startRange = date.Start
			endRange = date.End
		}
	} else {
		startRange = "2021-11-01"
		endRange = "2021-11-30"
	}

	layoutFormat := "2006-01-02"

	startDate, _ := time.Parse(layoutFormat, startRange)
	endDate, _ := time.Parse(layoutFormat, endRange)

	for rd := helpers.RangeDate(startDate, endDate); ; {
		date := rd()
		if date.IsZero() {
			break
		}
		tanggal := date.Format("2006-01-02")

		db.Select("m.merchant_name, date(t.created_at) as date, sum(t.bill_total) as omzet").
			Table("Transactions as t").
			Joins("join Merchants as m on t.merchant_id = m.id").
			Group("m.merchant_name, date").
			Where("m.user_id = ?", userID).
			Where("date(t.created_at) = ?", tanggal).
			Find(&arrMerchantReport)

		if len(arrMerchantReport) > 0 {
			merchantReport = MerchantReport{MerchantName: arrMerchantReport[0].MerchantName, Date: arrMerchantReport[0].Date, Omzet: arrMerchantReport[0].Omzet}
			arrMerchantReport = append(arrMerchantReport, merchantReport)
		} else {
			date, _ := time.Parse(layoutFormat, tanggal)
			db.Where("user_id = ?", userID).Find(&merchants)
			for i := 0; i < len(merchants); i++ {

				merchantReport = MerchantReport{MerchantName: merchants[i].MerchantName, Date: date, Omzet: 0}
				arrMerchantReport = append(arrMerchantReport, merchantReport)
			}
		}
		arrReport = append(arrReport, arrMerchantReport[0])
	}

	max := limit * page

	for i := start; i < max; i++ {
		MerchantReports = append(MerchantReports, arrReport[i])
	}
	count := len(arrReport)
	return count, MerchantReports
}
