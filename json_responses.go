package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type BaseResponseStructure struct {
	StatusCode 		uint							`json:"status_code"'`
	Body 			map[string]interface{}			`json:"body"`
	ServerTime 		string   						`json:"serverTime"`
}

func ReturnBaseResponse(w *http.ResponseWriter, code uint, obj map[string]interface{}) {
	response := &BaseResponseStructure{
		StatusCode: 	code,
		Body:       	obj,
		ServerTime: 	time.Now().String(),
	}
	if err := json.NewEncoder(*w).Encode(response); err != nil {
		log.Println("Error encoding json:", err)
	}
}

type ErrorResponseStructure struct {
	StatusCode 		uint		`json:"status_code"'`
	Error 			string		`json:"error"`
	ServerTime 		string   	`json:"serverTime"`
}

func ReturnErrorResponse(w *http.ResponseWriter, code uint, err error) {
	response := &ErrorResponseStructure{
		StatusCode: 	code,
		Error:       	err.Error(),
		ServerTime: 	time.Now().String(),
	}

	if err := json.NewEncoder(*w).Encode(response); err != nil {
		log.Println("Error encoding json:", err)
	}
}