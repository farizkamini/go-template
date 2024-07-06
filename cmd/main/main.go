package main

import (
	"context"
	"fmt"
	"gomodule/config/db"
	"gomodule/config/server"
	"gomodule/config/serverstatic"
	"gomodule/pkg/vip"
	"gomodule/pkg/zlog"

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
		zlog.Info(fmt.Sprintf("port asset: %d", vipp.AppPortAsset), "server asset run")
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
