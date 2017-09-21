package main

func (a *App) initializeRoutes() {

	a.router.HandleFunc("/signup", signup).Methods("POST")
}
