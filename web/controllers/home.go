package controllers

import (
	"net/http"
)

func (app *Application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	key := "hello"
	value, err := app.Fabric.Query(key)
	if err != nil {
		http.Error(w, "Unable to query the blockchain", 500)
	}

	data := &struct {
		Key   string
		Value string
	}{
		Key:   key,
		Value: value,
	}
	app.renderTemplate(w, r, "home.html", data)
}
