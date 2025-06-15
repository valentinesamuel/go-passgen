package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func HandleGeneratePassword(w http.ResponseWriter, req *http.Request) {

	passwordLength, err := strconv.Atoi(req.URL.Query().Get("length"))
	if passwordLength == 0 {
		sendHttpResponse(w, http.StatusBadRequest, nil, "Invalid length")
		return
	}

	if err != nil {
		sendHttpResponse(w, http.StatusBadRequest, nil, err.Error())
		return
	}

	upperCase, err := strconv.ParseBool(req.URL.Query().Get("upper"))
	if err != nil {
		upperCase = true
	}

	lowerCase, err := strconv.ParseBool(req.URL.Query().Get("lower"))
	if err != nil {
		lowerCase = true
	}

	digits, err := strconv.ParseBool(req.URL.Query().Get("digits"))
	if err != nil {
		digits = true
	}

	symbols, err := strconv.ParseBool(req.URL.Query().Get("symbols"))
	if err != nil {
		symbols = true
	}

	pool, err := generatePassword(passwordLength, upperCase, lowerCase, digits, symbols)

	if err != nil {
		sendHttpResponse(w, http.StatusBadRequest, nil, err.Error())
		return
	}

	finalPassword, err := buildPasswordPool(passwordLength, pool)
	if err != nil {
		sendHttpResponse(w, http.StatusBadRequest, nil, err.Error())
		return
	}

	warningMsg := ""
	if passwordLength > 10 {
		warningMsg = "This password length might be too long to remember"
	}
	sendHttpResponse(w, http.StatusOK, finalPassword, warningMsg)

}

func sendHttpResponse(w http.ResponseWriter, statusCode int, data interface{}, extraMessage string) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)

	var message string
	var successStatus bool

	if statusCode > 199 && statusCode < 399 {
		message = "✅  Request Successful"
		successStatus = true
	} else {
		message = "❌  Request Failed!!"
		successStatus = false
	}

	response := ResponseData{
		Data:     data,
		Success:  successStatus,
		Response: message,
		Message:  extraMessage,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
