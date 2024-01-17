package main

import "net/http"

func (app *application) signInHandler(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, Template{Layout: "base", View: "sign_in.html"})
}

func (app *application) signInPost(w http.ResponseWriter, r *http.Request) {

}

func (app *application) signOutHandler(w http.ResponseWriter, r *http.Request) {

}

func (app *application) signOutPost(w http.ResponseWriter, r *http.Request) {

}

func (app *application) resetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, Template{Layout: "base", View: "reset_password.html"})
}

func (app *application) createPasswordResetTokenHandler(w http.ResponseWriter, r *http.Request) {

}

func (app *application) createActivationTokenHandler(w http.ResponseWriter, r *http.Request) {

}
