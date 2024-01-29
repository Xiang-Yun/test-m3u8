package main

import "net/http"

func (app *application) Video(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "video", nil); err != nil {
		app.errorLog.Println(err)
	}
}
