package controllers

import (
	"net/http"
)

func (app *Application) RequestHandler(w http.ResponseWriter, r *http.Request) {
	data := &struct {
		TransactionId string
		Success       bool
		Response      bool
	}{
		TransactionId: "",
		Success:       false,
		Response:      false,
	}
	if r.FormValue("submitted") == "true" {
		value := r.FormValue("hello")
		txid, err := app.Fabric.Invoke("hello", value)
		if err != nil {
			http.Error(w, "Unable to invoke in the blockchain", 500)
		}
		data.TransactionId = txid
		data.Success = true
		data.Response = true
	}
	app.renderTemplate(w, r, "request.html", data)
}
