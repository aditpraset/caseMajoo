package controllers

import (
	"caseMajoo/api/auth"
	"caseMajoo/api/models"
	"caseMajoo/api/responses"
	"net/http"
)

func (server *Server) GetMerchantReport(w http.ResponseWriter, r *http.Request) {
	var res models.Merchants
	var arrMerchantReport []models.MerchantReport
	var response models.ResponseJson
	paging := models.Paging(r)
	userID, _ := auth.ExtractTokenID(r)

	// var date []models.Date
	date := paging.Dates

	total, arrMerchantReport := res.MerchantReport(paging.FromRow, paging.Page, paging.Limit, server.DB, userID, date)

	pages := (total / paging.Limit)
	if (total % paging.Limit) != 0 {
		pages++
	}
	paging.TotalRows = total
	paging.LastPage = pages
	response.Message = "Success"
	response.Success = "true"
	response.Paging = paging
	response.Data = arrMerchantReport
	token, err := auth.CreateToken(userID)

	if err != nil {
		response.Success = "false"
		response.Message = "Something Wrong"
		responses.JSON(w, http.StatusUnprocessableEntity, response)
		return
	}
	response.Token = token
	responses.JSON(w, http.StatusOK, response)
}
