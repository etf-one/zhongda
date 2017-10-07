package UserController

import (
	"encoding/json"
	"net/http"

	e "github.com/etf-one/zhongda/app/http/Errors"
)

type SignUpInfo struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	Lastname  string `json:"lastname"`
}

func Signup(w http.ResponseWriter, req *http.Request) {
	var info SignUpInfo
	decoder := json.NewDecoder(req.Body)
	defer req.Body.Close()

	err := decoder.Decode(&info)

	if err != nil {
		e.ReturnError(w, e.RequestEmptyOrInvalid, err.Error())
		return
	}
	w.WriteHeader(200)
	w.Write([]byte("Successful"))
}
