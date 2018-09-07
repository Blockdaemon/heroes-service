package web

import (
	"fmt"
	"github.com/Blockdaemon/hlf-webapp/web/controllers"
	"log"
	"net/http"
)

func Serve(app *controllers.Application) {
	fs := http.FileServer(http.Dir(app.WebRoot + "/web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/home.html", app.HomeHandler)
	http.HandleFunc("/request.html", app.RequestHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/home.html", http.StatusTemporaryRedirect)
	})

	fmt.Printf("Listening (http://localhost:%d/) ...\n", app.WebPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", app.WebPort), nil))
}
