搭建长安链环境
感觉提示来搭建即可，如果是服务器环境，要注意暴露对应的端口。
https://docs.chainmaker.org.cn/v2.3.2/html/quickstart/%E9%80%9A%E8%BF%87%E7%AE%A1%E7%90%86%E5%8F%B0%E4%BD%93%E9%AA%8C%E9%93%BE.html

```
git clone -b v2.3.2 --depth=1 https://git.chainmaker.org.cn/chainmaker/management-backend.git
cd management-backend
docker compose up -d
```
访问对应ip的80端口即可

区块链管理->配置链之后，
下载连文件并部署，如果是服务器环境，就上传到服务器上。
![部署链](images/chain.png)
```bash
scp ~/Downloads/test.zip  MyTencent:/root/test.zip
```



```bash
docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.20 bash -c 'go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct && CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -v'

nohup ./carservice &
```
运行环境

git	/	源码管理	是
golang	1.16+	编译环境	是
docker	18+	独立运行容器	否
docker-compose	/	容器管理组件	否
gcc	7.3+	编译环境依赖	是
glibc	2.18	智能合约执行环境依赖	是

https://docs.chainmaker.org.cn/recovery/%E9%80%9A%E8%BF%87%E5%91%BD%E4%BB%A4%E8%A1%8C%E5%B7%A5%E5%85%B7%E5%90%AF%E5%8A%A8%E9%93%BE.html#normalError


```shell
gcc --version
# 如果gcc<7.3，则运行一下命令升级
sudo yum install centos-release-scl
sudo yum install devtoolset-7-gcc*
scl enable devtoolset-7 bash
```


下载配置文件：
![download-config](images/download-config.png)

将配置文件解压到根路径下，并根据项目目录修改配置文件的路径
![config-path](images/config-path.png)

然后直接运行：

```golang
go run main.go
```

如果合约没有部署，则会部署合约

```bash
2023/11/11 15:35:23 chain client initialize successfully
合约[token]不存在
====================== 创建合约 ======================
2023/11/11 15:35:24 contract_result:<result:"\n\005token\022\0051.0.0\030\005*\226\001\n\nTestCMorg1\020\001\032 \357\377\010\250\271\243\2570\017\201\350\367\221\311\261\177\200\211\017;g\337M,\226\0243\222!:\364\305\"\033cmtestuser1.sign.TestCMorg1*\005ADMIN2@3eddab9da000facf9b3fc840660748cdb49b2ef303c995853bf6f3a2d18945112(6795f86e288bc16c31d22f9ea23955224121d74d" gas_used:1061794 contract_event:<topic:"ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef" tx_id:"17968157d5691844ca5656f4803126714e7174eb62be44c3b2cc0b211c7bc54e" contract_name:"token" contract_version:"1.0.0" event_data:"0000000000000000000000000000000000000000000000000000000000000000" event_data:"0000000000000000000000000000000000000000000000000000000000000000" event_data:"000000000000000000000000000000000000000000000000016345785d8a0000" > > tx_id:"17968157d5691844ca5656f4803126714e7174eb62be44c3b2cc0b211c7bc54e" tx_timestamp:1699688123 tx_block_height:2
addr [470532945834136871529055253531705118185281158937] => [%!d(string=0)]
2023/11/11 15:35:24 ethaddr => 0x526b6d51426e2603b97880cf44d9b34ae5e5ff19
2023/11/11 15:35:24 skiaddr => 3eddab9da000facf9b3fc840660748cdb49b2ef303c995853bf6f3a2d1894511

```

合约部署成功则可以在管理平台中查看合约
![contract-manage](images/contract-manage.png)

// 初始化数据库数据

这里可以直接用`carserie.sql`来初始化数据

```
go test -v infra/database/database_test.go -run TestPrePareHighwayData
go test -v infra/database/database_test.go -run TestPrePareChargeData
go test -v infra/database/database_test.go -run TestPrePareParkData
```

生成 QRCode 的 text，然后去草料二维码生成二维码
https://cli.im/text

```
go test -v util/rsa_test.go -run TestGenerateQRCodeText
```

## highway Qrcode

高速入口:
![高速入口](images/car_highway_start.png)

高速出口:
![高速出口](images/car_highway_end.png)

充电桩，开始充电
![开始充电](images/car_charge_start.png)

充电桩，结束充电
![结束充电](images/car_charge_end.png)

停车，开始停车
![开始停车](images/car_park_start.png)

停车，结束停车
![结束停车](images/car_park_end.png)
