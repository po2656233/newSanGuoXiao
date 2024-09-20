package controller

import superGin "github.com/po2656233/superplace/components/gin"

type Controller struct {
	superGin.BaseController
}

func (p *Controller) Init() {
	group := p.Group("/")
	group.GET("/", p.index)

	group.POST("/register", p.register) //用户注册
	group.POST("/login", p.login)       //用户登录
	group.POST("/recharge", p.recharge) //用户登录

	list := &superGin.Group{
		RouterGroup: group.Group("/list"),
	}
	list.GET("/pid", p.pidList)            //包列表
	list.GET("/server", p.serverList0)     //服务器网关地址
	list.GET("/server/:pid", p.serverList) //服务器网关地址
	list.GET("/class", p.classList)        //分类列表
	list.GET("/room", p.roomList)          //房间列表
	list.GET("/table", p.tableList)        //牌桌列表
	list.GET("/game", p.gameList)          //游戏列表

	create := &superGin.Group{
		RouterGroup: group.Group("/create"),
	}
	create.POST("/room", p.roomCreate)   //房间创建
	create.POST("/table", p.tableCreate) //房间创建

	del := &superGin.Group{
		RouterGroup: group.Group("/delete"),
	}
	del.POST("/table", p.tableDelete) //删除牌桌

}
