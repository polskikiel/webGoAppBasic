package controller

import (
	"net/http"
	"app/shared/view"
)

func LoginGET(w http.ResponseWriter, r *http.Request) {
	v := view.New(r)
	v.Name = "login/login"
	v.Render(w)
}
