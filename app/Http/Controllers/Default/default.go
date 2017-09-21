package DefaultController

import (
	"net/http"

	//Framework
	e "github.com/etf-one/zhongda/app/http/Errors"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	e.ReturnSuccess(w, "hello world")
}

func SayError(w http.ResponseWriter, r *http.Request) {
	e.ReturnError(w, e.GenericError, "something went wrong")
}
