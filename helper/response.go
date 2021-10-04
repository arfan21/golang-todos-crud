package helper

import (
	"encoding/json"
	"net/http"
)

func Response(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	dataJson, _ := json.Marshal(data)
	w.Write(dataJson)
}
