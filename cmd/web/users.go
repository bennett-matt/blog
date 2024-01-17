package main

import "net/http"

type userSignupForm struct {
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password string `form:"password"`
	// validator.Validator `form:"-"`
}

func (app *application) userSignupHandler(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, Template{View: "sign_up.html", Layout: "base"})
}

func (app *application) userSignupPost(w http.ResponseWriter, r *http.Request) {
}

func (app *application) activateUserHandler(w http.ResponseWriter, r *http.Request) {
}

func (app *application) updateUserPasswordHandler(w http.ResponseWriter, r *http.Request) {
}
