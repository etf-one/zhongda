package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/pjebs/restgate"
	"github.com/urfave/negroni"

	//Framework
	route "github.com/etf-one/zhongda/app/Http"
	s "github.com/etf-one/zhongda/app/Providers"
	c "github.com/etf-one/zhongda/config"
)

var addr *string

func init() {
	addr = flag.String("addr", ":3001", "Host Address. Default - :3001")
	flag.Parse()
}

func main() {

	app := negroni.New()

	app.Use(negroni.NewRecovery())
	app.Use(negroni.NewLogger())
	//Add Security Credential Requirements to API routes.
	//For production, keep HTTPSProtection = true
	HTTPSProtection := false
	if HTTPSProtection {
		app.Use(restgate.New("X-Auth-Key", "X-Auth-Secret", restgate.Static, restgate.Config{HTTPSProtectionOff: false, Key: []string{c.API_ENDPOINT_KEY}, Secret: []string{c.API_ENDPOINT_SECRET}}))
	}
	app.Use(negroni.HandlerFunc(s.ServiceProviders)) //Initializes Service Providers
	app.UseHandler(route.New())
	http.Handle("/", app)

	log.Fatal(http.ListenAndServe(*addr, nil))
}
