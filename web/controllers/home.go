package controllers

import (
	"net/http"
)

func (app *Application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	value, err := app.Fabric.Query("hello")
	if err != nil {
		http.Error(w, "Unable to query the blockchain", 500)
	}

	data := &struct {
		Hello string
	}{
		Hello: value,
	}
	app.renderTemplate(w, r, "home.html", data)
}
