package handlers

import (
	"net/http"
	"shohag/github.com/util"
)

func Test(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, "test", 200)
}
