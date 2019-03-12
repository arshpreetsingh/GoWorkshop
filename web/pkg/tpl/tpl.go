package tpl

import (
	"encoding/json"
	"net/http"
)

// Data struct
type Data struct {
	Data interface{}
}

func (d Data) RenderJson(w http.ResponseWriter, r *http.Request) {
	data := Data{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		panic(err)
	}
	dataJson, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	//Set Content-Type header so that clients will know how to read response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//Write json response back to response
	w.Write(dataJson)
	return
}
