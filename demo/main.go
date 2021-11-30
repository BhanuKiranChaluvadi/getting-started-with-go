package main

import (
	"net/http"

	"github.com/BhanuKiranChaluvadi/getting-started-with-go/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
