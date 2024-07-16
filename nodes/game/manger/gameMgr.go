package manger

// CalculateInfo 结算信息
type CalculateInfo struct {
	UserID   int64  // 结算对象(接受充值的ID)
	ByUID    int64  // 该用户发起
	PreMoney int64  // 结算前的钱
	Payment  int64  // 付款金额
	Code     int32  // 操作码(即为什么付款)
	Order    string // 订单号(或牌局号)
	Remark   string // 备注
}

// ProductCallback 实例回调(根据桌子号创建游戏)
type ProductCallback func(table *Table) IGameOperate

type CalculateSQL func(info CalculateInfo) (nowMoney, factDeduct int64, isOK bool)

// ClearCallback 清场回调
type ClearCallback func() (gameID int64)

// IGameOperate 子游戏接口
type IGameOperate interface {
	Scene(args []interface{})        //场 景
	Start(args []interface{})        //开 始
	Playing(args []interface{})      //游 戏(下分|下注)
	Over(args []interface{})         //结 算
	UpdateInfo(args []interface{})   //更新信息
	SuperControl(args []interface{}) //超级控制 在检测到没真实玩家时,且处于空闲状态时,自动关闭
}

// IAgainst 对战类
type IAgainst interface {
	Ready(args []interface{})     //准备
	DispatchCard()                //发牌
	CallScore(args []interface{}) //叫分
	OutCard(args []interface{})   //出牌
	Discard(args []interface{})   //操作
}

// IMultiPlayer 百人类
type IMultiPlayer interface {
	DispatchCard()                              //发牌
	DeduceWin() []byte                          //开奖区域
	BonusArea(area int32, betScore int64) int64 //区域赔额
}

type IDevice interface {
	Ready() bool              //准备
	Start(time int64) bool    //开启(游戏)
	Maintain(time int64) bool //维护(暂停)
	Clear(time int64) bool    //清场
	Close() bool              //关闭
}
