package main

import (
	"github.com/macropusgiganteus/covid-summarizer-api/internal/server"
)

func main() {
	r := server.SetupRouter()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
