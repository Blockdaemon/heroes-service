package controllers

import (
	"net/http"
)

func (app *Application) RequestHandler(w http.ResponseWriter, r *http.Request) {
	key := "hello"
	value, err := app.Fabric.Query(key)
	if err != nil {
		http.Error(w, "Unable to query the blockchain", 500)
	}
	data := &struct {
		Key           string
		Value         string
		TransactionId string
		Success       bool
		Response      bool
	}{
		Key:           key,
		Value:         value,
		TransactionId: "",
		Success:       false,
		Response:      false,
	}
	if r.FormValue("submitted") == "true" {
		value := r.FormValue(key)
		txid, err := app.Fabric.Invoke(key, value)
		if err != nil {
			http.Error(w, "Unable to invoke in the blockchain", 500)
		}
		data.TransactionId = txid
		data.Success = true
		data.Response = true
	}
	app.renderTemplate(w, r, "request.html", data)
}
