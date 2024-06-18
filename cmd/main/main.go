package main

import (
	"context"
	"mereetmoi/config/db"
	"mereetmoi/config/server"
	serverstatic "mereetmoi/config/serverstatic"
	"mereetmoi/pkg/vip"
	"mereetmoi/pkg/zlog"
	"net/http"
	"strconv"
)

func main() {
	server.CreateLogFile()
	vipp, errVip := vip.New().App()
	if errVip != nil {
		zlog.Fatal(errVip)
		return
	}
	go func() {
		zlog.Info("asset run", "asset run")
		err := http.ListenAndServe(":"+strconv.Itoa(vipp.AppPortAsset), serverstatic.Master())
		if err != nil {
			zlog.Fatal(err)
			return
		}
	}()

	server.CreateDirAssets()
	ctx := context.Background()
	_, err := db.New(ctx).Conn()
	if err != nil {
		zlog.Fatal(err)
		return
	}
	server.New().Serve()
}
