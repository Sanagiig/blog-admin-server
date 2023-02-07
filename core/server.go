package core

import (
	"fmt"
	"go-blog/global/settings"
	"go-blog/initialize"
	"go-blog/service"
	"net/http"
	"time"
)

func RunServer() {
	r := initialize.InitRouter()
	s := http.Server{
		Addr:           settings.HttpPort,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	service.SysStore.Routes = r.Routes()
	fmt.Println("server is running at ", s.Addr)
	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
