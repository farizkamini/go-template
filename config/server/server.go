package server

import (
	"fmt"
	"mereetmoi/internal/routes"
	"mereetmoi/pkg/vip"
	"mereetmoi/pkg/zlog"
	"net/http"

	"strconv"
)

type SrvConfig struct {
	Port    string
	Handler http.Handler
}

func New() *SrvConfig {
	return &SrvConfig{}
}

func (s *SrvConfig) Serve() {
	vipp, errVip := vip.New().App()
	if errVip != nil {
		zlog.Fatal(errVip)
		return
	}

	srv := &SrvConfig{
		Port:    ":" + strconv.Itoa(vipp.AppPort),
		Handler: routes.Master(),
	}

	zlog.Info(nil, fmt.Sprintf("lets rawk at %s", srv.Port))
	err := http.ListenAndServe(srv.Port, srv.Handler)
	if err != nil {
		zlog.Fatal(err)
	}
}
