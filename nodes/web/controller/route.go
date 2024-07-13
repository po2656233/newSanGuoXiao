package controller

import superGin "github.com/po2656233/superplace/components/gin"

type Controller struct {
	superGin.BaseController
}

func (p *Controller) Init() {
	group := p.Group("/")
	group.GET("/", p.index)
	group.GET("/pid/list", p.pidList)            //包列表
	group.GET("/server/list/:pid", p.serverList) //服务器网关地址
	group.POST("/register", p.register)          //用户注册
	group.POST("/login", p.login)                //用户登录
	group.GET("/class/list", p.classList)        //分类列表
	group.GET("/room/list", p.roomList)          //房间列表
	group.GET("/table/list", p.tableList)        //牌桌列表
	group.GET("/game/list", p.gameList)          //游戏列表
	group.POST("/room/create", p.roomCreate)     //房间创建
	group.POST("/table/create", p.tableCreate)   //房间创建
}
