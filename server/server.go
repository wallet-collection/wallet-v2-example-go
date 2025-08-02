package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"wallet/api"
	"wallet/client"
	"wallet/config"
	"wallet/job"
	"wallet/model"
	"wallet/redis"
)

// Start 启动服务
func Start(isSwag bool, configPath string) {

	conf, err := config.NewConfig(configPath)

	if err != nil {
		panic("Failed to load configuration")
	}

	db, err := model.NewDB(conf.Mysql)
	if err != nil {
		panic(err)
	}

	_ = redis.NewRedisDB(conf.Redis)

	// 创建定时任务
	job.CreateJob(db)

	// 创建钱包客户端
	client.NewWalletClient(conf.Wallet.Appid, conf.Wallet.SecretKey, conf.Wallet.Url)

	if isSwag {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	server := gin.Default()

	// 设置受信任代理,如果不设置默认信任所有代理，不安全
	err = server.SetTrustedProxies(nil)
	if err != nil {
		panic(err)
	}

	// 中间件
	server.Use(gin.Logger())
	server.Use(gin.Recovery())
	server.Use(api.Cors())

	interfaceGroup := server.Group("/")
	{
		// 币种相关
		// 账单的筛选信息
		interfaceGroup.GET("/memberBill/businessType", api.MemberBillBusinessType)
		interfaceGroup.GET("/coinConf/list", api.CoinConfList)
		// 回调
		interfaceGroup.POST("/recharge/call", api.RechargeCall) // 需要和 api 里面的 recharge.go 文件里面的路由一致
		// 这是 recharge.go 文件的路由代码地址 callUrl := config.CONF.Wallet.CallUrl + "/recharge/call"

		interfaceGroup.POST("/withdraw/call", api.WithdrawCall) // 需要和 api 里面的 withdraw.go 文件里面的路由一致
		// 这是 withdraw.go 文件的路由代码地址 callUrl := config.CONF.Wallet.CallUrl + "/withdraw/call"

		// 登录
		interfaceGroup.POST("/login", api.AuthLoginByPwd)
		interfaceGroup.POST("/register", api.AuthRegister)

		//authId := interfaceGroup.Group("/", api.AuthId())
		//{
		//
		//}

		auth := interfaceGroup.Group("/", api.AuthRequired())
		{
			auth.GET("/memberBill/list", api.MemberBillList)
			auth.GET("/memberCoin/info", api.MemberCoinInfo)
			auth.GET("/memberCoin/list", api.MemberCoinList)

			// 获取登录用户信息
			auth.GET("/member/loginInfo", api.MemberLoginInfo)

			// 币种
			auth.GET("/recharge/address", api.RechargeAddress)
			auth.GET("/recharge/list", api.RechargeList)
			auth.GET("/recharge/info", api.RechargeInfo)
			auth.POST("/withdraw", api.WithdrawCreate)
			auth.GET("/withdraw/list", api.WithdrawList)
			auth.GET("/withdraw/info", api.WithdrawInfo)
		}

	}

	err = server.Run(fmt.Sprintf(":%v", conf.App.Port))
	if err != nil {
		panic("start error")
	}

	fmt.Println("start success")

}
