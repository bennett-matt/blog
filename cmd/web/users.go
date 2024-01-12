package main

import "net/http"

type userSignupForm struct {
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password string `form:"password"`
	// validator.Validator `form:"-"`
}

func (app *application) userSignup(w http.ResponseWriter, r *http.Request) {
}

func (app *application) userSignupPost(w http.ResponseWriter, r *http.Request) {
}

func (app *application) activateUserHandler(w http.ResponseWriter, r *http.Request) {
}

func (app *application) updateUserPasswordHandler(w http.ResponseWriter, r *http.Request) {
}
