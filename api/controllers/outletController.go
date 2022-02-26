package controllers

import (
	"caseMajoo/api/auth"
	"caseMajoo/api/models"
	"caseMajoo/api/responses"
	"net/http"
)

func (server *Server) GetOutletReport(w http.ResponseWriter, r *http.Request) {
	var res models.Outlets
	var arrOutletReport []models.OutletReport
	var response models.ResponseJson
	paging := models.Paging(r)
	userID, _ := auth.ExtractTokenID(r)

	date := paging.Dates

	total, arrOutletReport := res.PaginateOutletReport(paging.FromRow, paging.Page, paging.Limit, server.DB, userID, date)
	pages := (total / paging.Limit)
	if (total % paging.Limit) != 0 {
		pages++
	}
	paging.TotalRows = total
	paging.LastPage = pages
	response.Success = "true"
	response.Message = "Success"
	response.Paging = paging
	response.Data = arrOutletReport

	signedToken, err := auth.CreateToken(userID)

	if err != nil {
		response.Success = "false"
		response.Message = "Something Wrong"
		responses.JSON(w, http.StatusUnprocessableEntity, response)
		return
	}
	response.Token = signedToken
	responses.JSON(w, http.StatusOK, response)
}
