package main

import (
	"carservice/infra/config"
	"carservice/infra/database"
	"carservice/infra/logger"
	"carservice/service"
	"carservice/service/router"
	"flag"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// is debug
	var isDebug = flag.Bool("debug", false, "当前是否在debug下开启服务")
	flag.Parse()

	// init logger
	l := logger.Initialize(*isDebug)

	// init conifg
	config.Initialize(*isDebug)

	// init database
	database.Initialize(*isDebug)

	// init service
	eg := router.Initialize(*isDebug)
	srv := service.NewService(eg)
	srv.Run()

	// chainClient, err := chain.NewChainClient()

	// _ = chainClient
	// _ = err

	// user, err := chain.GetUser("cmtestuser1")
	// if err != nil {
	// 	logger.Panic("get cmtestuser1 failed", err)
	// }
	// client1AddrInt, client1EthAddr, client1AddrSki, err := chain.MakeAddrAndSkiFromCrtFilePath(user.SignCrtPath)
	// chainClient.TokenContractBalanceOf(client1AddrInt)
	// if err != nil {
	// 	logger.Panic("get token balance failed", err)
	// }
	// logger.Debugf("ethaddr => %s\n", client1EthAddr)
	// logger.Debugf("skiaddr => %s\n", client1AddrSki)

	// 优雅退出
	// kill -9 是捕捉不到的
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	srv.Shutdown()
	l.Shutdown()

}
