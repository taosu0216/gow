package main

import "gow/internal/server/http"

func main() {
	h := http.InitRouters()
	http.StartHttpServer(h)
}
