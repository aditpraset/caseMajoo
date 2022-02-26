package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"caseMajoo/api/auth"
	"caseMajoo/api/models"
	"caseMajoo/api/responses"
	"crypto/md5"
	"encoding/hex"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	var response models.ResponseJson

	hashPassword := GetMD5Hash(user.Password)

	if err != nil {
		response.Success = "false"
		response.Message = "Login Failed! Please enter your Username and Password"
		responses.JSON(w, http.StatusUnprocessableEntity, response)
		return
	}

	err = user.Validate("login")
	if err != nil {
		response.Success = "false"
		response.Message = "Login Failed! Please enter your Username and Password"
		responses.JSON(w, http.StatusUnprocessableEntity, response)
		return
	}

	err = server.DB.Where("user_name = ?  and password = ?", user.UserName, hashPassword).First(&user).Error

	if err != nil {
		response.Success = "false"
		response.Message = "Login Failed!  Invalid Username or Password"
		responses.JSON(w, http.StatusUnprocessableEntity, response)
		return
	}

	token, _ := auth.CreateToken(user.ID)

	response.Token = token
	response.Message = "Login Success"
	response.Success = "true"
	response.Data = user

	responses.JSON(w, http.StatusOK, response)
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
