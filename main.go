package main

import (
	"net/http"

	"github.com/puhelan/webserver/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
