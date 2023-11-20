package main

import (
	"carservice/infra/config"
	"carservice/infra/database"
	"carservice/infra/logger"
	"carservice/service"
	"carservice/service/chain"
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

	// // init logger
	l := logger.Initialize(*isDebug)

	// init conifg
	config.Initialize(*isDebug)

	// init database
	database.Initialize(*isDebug)

	// 初始化链的client
	chainClient, err := chain.NewChainClient()
	if err != nil {
		logger.Panic("new chain client failed", err)
	}
	// 初始化contracts
	chainClient.InitContracts()
	user, err := chain.GetUser("cmtestuser1")
	if err != nil {
		logger.Panic("get cmtestuser1 failed", err)
	}
	client1AddrInt, client1EthAddr, client1AddrSki, err := chain.MakeAddrAndSkiFromCrtFilePath(user.SignCrtPath)
	if err != nil {
		logger.Panic("make addr and ski from crt file failed", err)
	}
	logger.Debugf("ethaddr => %s\n", client1EthAddr)
	logger.Debugf("skiaddr => %s\n", client1AddrSki)
	logger.Debugf("addrint => %d\n", client1AddrInt)

	// init service
	eg := router.Initialize(*isDebug)
	srv := service.NewService(eg, chainClient)
	srv.Run()

	// 优雅退出
	// kill -9 是捕捉不到的
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	srv.Shutdown()
	chainClient.Shutdown()
	l.Shutdown()

}
