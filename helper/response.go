package helper

import (
	"encoding/json"
	"net/http"
)

type BaseResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(w http.ResponseWriter, data BaseResponse) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(data.Status)
	dataJson, _ := json.Marshal(data)
	w.Write(dataJson)
}
