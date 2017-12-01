package controller

import (
	"net/http"
	"app/shared/view"
)


func IndexGET(w http.ResponseWriter, r *http.Request) {
	v := view.New(r)
	v.Name = "base"
	v.Render(w)
}