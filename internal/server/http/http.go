package http

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"gow/internal/pkg/proxy"
)

func InitRouters() *server.Hertz {
	h := server.New(server.WithHostPorts("0.0.0.0:8877"))

	testG := h.Group("/api/v1/test")
	{
		testG.GET("/ping", proxy.PingTest)
		testG.GET("/proxy", proxy.PTest)
	}

	return h
}

func StartHttpServer(h *server.Hertz) {
	h.Spin()
}
