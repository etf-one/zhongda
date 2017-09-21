package route

import (
	"github.com/gorilla/mux"

	//Framework
	//c "github.com/etf-one/zhongda/config"

	//Controllers
	cD "github.com/etf-one/zhongda/app/http/Controllers/Default"
)

func New() *mux.Router {

	//Create main router
	mainRouter := mux.NewRouter().StrictSlash(true)
	mainRouter.KeepContext = true

	//App Routes
	mainRouter.Methods("GET").Path("/").HandlerFunc(cD.SayHello)
	mainRouter.Methods("GET").Path("/error").HandlerFunc(cD.SayError)

	return mainRouter

}
