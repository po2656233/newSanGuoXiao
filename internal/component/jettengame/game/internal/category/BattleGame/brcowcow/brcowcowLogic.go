package brcowcow

import (
	"fmt"
	"sync"

	"github.com/po2656233/goleaf/log"
)

const (
	Area_Tian   = iota //天
	Area_Di            //地
	Area_Xuan          //玄
	Area_Huang         //黄
	Area_Banker        //庄(不可下注)

	MULTIPLE_ONE   = 1.0
	MULTIPLE_TWO   = 1.0
	MULTIPLE_THREE = 1.0
	MULTIPLE_FOUR  = 1.0
	MULTIPLE_FIVE  = 1.0
	MULTIPLE_SIX   = 1.0
	MULTIPLE_SEVEN = 2.0
	MULTIPLE_EIGHT = 2.0
	MULTIPLE_NINE  = 2.0
	MULTIPLE_TEN   = 3.0

	AREA_MAX        = 4  //最大区域
	CardCount       = 52 //牌数目
	CardAmount      = 5
	BrcowcowHostMax = 5
)

// 大数定义统一置文本后
var CardListData = [CardCount]byte{
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, //方块 A - K
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, //梅花 A - K
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2A, 0x2B, 0x2C, 0x2D, //红桃 A - K
	0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3A, 0x3B, 0x3C, 0x3D, //黑桃 A - K ------
}

// ----------------------------------
type WinType struct {
	Name string
	Type int
	Odds float64
}

type StuCowCow struct {
	BetType       []*BetType
	SameAsOdds    float64
	WinType       []*WinType
	BetWinInfoMap *BetWinInfoMap
}

func (g *StuCowCow) Init() {
	g.SameAsOdds = 1.0

	//下注类型
	g.AddBetType("天", Area_Tian, MULTIPLE_TEN)
	g.AddBetType("地", Area_Di, MULTIPLE_TEN)
	g.AddBetType("玄", Area_Xuan, MULTIPLE_TEN)
	g.AddBetType("黄", Area_Huang, MULTIPLE_TEN)

	//赔率
	g.AddWinType("五小牛", 13, MULTIPLE_TEN)
	g.AddWinType("炸弹牛", 12, MULTIPLE_TEN)
	g.AddWinType("五花牛", 11, MULTIPLE_TEN)
	g.AddWinType("牛牛", 10, MULTIPLE_TEN)
	g.AddWinType("牛九", 9, MULTIPLE_NINE)
	g.AddWinType("牛八", 8, MULTIPLE_EIGHT)
	g.AddWinType("牛七", 7, MULTIPLE_SEVEN)
	g.AddWinType("牛六", 6, MULTIPLE_SIX)
	g.AddWinType("牛五", 5, MULTIPLE_FIVE)
	g.AddWinType("牛四", 4, MULTIPLE_FOUR)
	g.AddWinType("牛三", 3, MULTIPLE_THREE)
	g.AddWinType("牛二", 2, MULTIPLE_TWO)
	g.AddWinType("牛一", 1, MULTIPLE_ONE)
	g.AddWinType("无牛", 0, MULTIPLE_ONE)

}

// 获取最大的赔付金额
//func (g *StuCowCow) HasLockMoneyByOdds(betInfo *BetInfo) float64 {
//	var moneyCount float64 = 0
//	//var info = betInfo.GetAll()
//	//for k1 := range info {
//	//	for k2 := range info[k1] {
//	//		moneyCount += info[k1][k2].Odds * info[k1][k2].BetMoney
//	//	}
//	//}
//	return moneyCount
//}

// 添加投注类型
func (g *StuCowCow) AddBetType(name string, betType int, odds float64) {
	g.BetType = append(g.BetType, &BetType{Name: name, Type: betType, Odds: odds})
}

// 添加赢钱类型
func (g *StuCowCow) AddWinType(name string, winType int, odds float64) {
	g.WinType = append(g.WinType, &WinType{Name: name, Type: winType, Odds: odds})
}

// 根据代号获取类型
func (g *StuCowCow) GetBetInfoByType(betType int) (*BetType, error) {
	for _, v := range g.BetType {
		if v.Type == betType {
			return v, nil
		}
	}
	return nil, fmt.Errorf("code is not exists")
}

func (g *StuCowCow) GetWinInfoByWinType(winType int) (*WinType, error) {
	for _, v := range g.WinType {
		if v.Type == winType {
			return v, nil
		}
	}
	return nil, fmt.Errorf("win is not exists:%v", winType)
}

// 获取赔率最大的类型
func (g *StuCowCow) GetMaxOddsType() *BetType {
	var max = 0.0
	var r *BetType
	for _, v := range g.BetType {
		if v.Odds > max {
			max = v.Odds
			r = v
		}
	}
	return r
}

func (g *StuCowCow) Compare(betType int, bankerType PokerType, personType PokerType) *BetWinInfo {

	var betWinInfo = &BetWinInfo{}

	betTypeInfo, _ := g.GetBetInfoByType(betType)

	if IsBankerWin(bankerType, personType) {
		// 如果庄家赢了 根据相应赔率算钱
		// 获取专家牌型的赔率
		// log.Printf("%+v",bankerType)
		isWin, err := g.GetWinInfoByWinType(bankerType.Type)
		if err != nil {
			log.Debug("错误：%v", err.Error())
			return nil
		}
		betWinInfo = &BetWinInfo{
			LoseOdds: isWin.Odds,
			WinOdds:  0,
			BetType:  betTypeInfo,
			IsWin:    false,
		}
	} else {
		// log.Printf("%+v",personType)
		isWin, err := g.GetWinInfoByWinType(personType.Type)
		if err != nil {
			log.Debug(err.Error())
			return nil
		}
		betWinInfo = &BetWinInfo{
			LoseOdds: 0,
			WinOdds:  isWin.Odds,
			BetType:  betTypeInfo,
			IsWin:    true,
		}
	}

	return betWinInfo
}

// CalcPoker CalcPoker
func CalcPoker(pokers Pokers) PokerType {

	var pokerType = PokerType{}

	var p1 = pokers.List[0]
	var p2 = pokers.List[1]
	var p3 = pokers.List[2]
	var p4 = pokers.List[3]
	var p5 = pokers.List[4]

	// 成员
	pokerType.Member = pokers.List

	// 最大牌
	pokerType.MaxPoint = p1.Number

	// 最大牌的花色
	pokerType.MaxColor = p1.Color

	// 最大牌的原始点数
	pokerType.MaxRaw = p1.Raw

	// 值
	pokerType.Value = p1.Number

	pokerType.ValueColor = p1.Color

	// 统计每个数字出现的个数
	var mapSlice = make(map[int]int)

	// 余数
	var leave = 0

	for i := 0; i < 5; i++ {
		leave += pokers.List[i].Mark
		if _, ok := mapSlice[pokers.List[i].Number]; ok {
			mapSlice[pokers.List[i].Number]++
		} else {
			mapSlice[pokers.List[i].Number] = 1
		}
	}

	leave = leave % 10

	// 五小牛
	if p1.Mark+p2.Mark+p3.Mark+p4.Mark+p5.Mark < 10 {
		if p1.Mark < 5 {
			pokerType.Type = 13
			return pokerType
		}
	}

	// 炸弹牛
	for k, v := range mapSlice {
		if v == 4 {
			pokerType.Value = k
			pokerType.Type = 12
			return pokerType
		}
	}

	// 五花牛
	if p1.Number > 10 && p2.Number > 10 && p3.Number > 10 && p4.Number > 10 && p5.Number > 10 {
		pokerType.Type = 11
		return pokerType
	}

	// 牛-?
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 5; j++ {
			// log.Println(pokers.List[i].Mark + pokers.List[j].Mark)
			if (pokers.List[i].Mark+pokers.List[j].Mark)%10 == leave {
				// 有牛
				if leave == 0 {
					// 牛牛
					pokerType.Type = 10
					return pokerType
				} else {
					// 牛X
					pokerType.Type = leave
					return pokerType
				}
			}
		}
	}

	pokerType.Type = 0
	return pokerType
}

// IsBankerWin IsBankerWin
func IsBankerWin(bankerType PokerType, personType PokerType) bool {

	// 直接对比类型 如果类型相等 比值的大小 值的大小相等 比花色 花色相等 庄家赢
	if bankerType.Type > personType.Type {
		return true
	}

	if bankerType.Type < personType.Type {
		return false
	}

	if bankerType.Type == personType.Type {

		// 炸弹牛直接比值
		if bankerType.Type == 12 {
			if bankerType.Value > personType.Value {
				return true
			}
			if bankerType.Value < personType.Value {
				return false
			}
			if bankerType.Value == personType.Value {
				return true
			}
		}

		// 不是炸弹牛 比最大的牌的点和花色
		if bankerType.MaxPoint > personType.MaxPoint {
			return true
		}
		if bankerType.MaxPoint < personType.MaxPoint {
			return false
		}
		if bankerType.MaxPoint == personType.MaxPoint {
			return true
		}
	}

	return true
}

func removePoker(pokers []Poker, value int) []Poker {
	if len(pokers) == 0 {
		return pokers
	}
	for i, v := range pokers {
		if v.Points == value {
			pokers = append(pokers[:i], pokers[i+1:]...)
			return removePoker(pokers, value)
		}
	}
	return pokers
}

//------------Poker------------------------引入

type BetType struct {
	Name string  `json:"name" bson:"name"`
	Type int     `json:"type" bson:"type"`
	Odds float64 `json:"odds" bson:"odds"`
}

type BetWinInfoMap struct {
	WinInfoMap map[string]map[int]*BetWinInfo
	mux        sync.RWMutex
}

type BetWinInfo struct {
	LoseOdds float64
	WinOdds  float64
	IsWin    bool
	BetType  *BetType
}

func (m *BetWinInfoMap) Init() {
	m.mux.Lock()
	defer m.mux.Unlock()
	m.WinInfoMap = make(map[string]map[int]*BetWinInfo)
}

func (m *BetWinInfoMap) Set(sourceName string, betType int, betWinInfo *BetWinInfo) {
	m.mux.Lock()
	defer m.mux.Unlock()
	m.WinInfoMap[sourceName][betType] = betWinInfo
}

func (m *BetWinInfoMap) Get(sourceName string, betType int) (*BetWinInfo, error) {
	m.mux.RLock()
	defer m.mux.RUnlock()
	if _, ok := m.WinInfoMap[sourceName][betType]; ok {
		return m.WinInfoMap[sourceName][betType], nil
	}
	return nil, fmt.Errorf("%d is not found", betType)
}

func (m *BetWinInfoMap) InitSourceMap(sourceName string) {
	m.mux.Lock()
	defer m.mux.Unlock()
	m.WinInfoMap[sourceName] = make(map[int]*BetWinInfo)
}

func (m *BetWinInfoMap) GetMap(sourceName string) (map[int]*BetWinInfo, error) {
	m.mux.RLock()
	defer m.mux.RUnlock()
	if _, ok := m.WinInfoMap[sourceName]; ok {
		return m.WinInfoMap[sourceName], nil
	}
	return nil, fmt.Errorf("%s is not found", sourceName)
}

// Poker Poker
type Poker struct {
	// 1-13 牌面
	Number int
	// 1-方块 2-梅花 3-红桃 4-黑桃
	Color int
	// 2-14 点数
	Points int
	// Mark JQK 为10点
	Mark int
	// Raw 原始
	Raw int
}

// Pokers Pokers
type Pokers struct {
	List []Poker
}

// PokerType PokerType
type PokerType struct {
	// 值
	Value      int
	ValueColor int
	// 第一位:9 豹子, 8 顺子, 7 对子, 6单张
	Type int
	// 第二位:最大点数
	MaxPoint int
	// 第三位:最大点数的花色
	MaxColor int
	// 最大牌值
	MaxRaw int
	// 第四位:是否是同花
	SameColor bool
	// 成员
	Member []Poker
}

// AddPokers AddPokers
func (p *Pokers) AddPokers(pokers ...Poker) *Pokers {
	p.List = append(p.List, pokers...)
	return p
}

// ArrangeByPoints ArrangeByPoints
func (p *Pokers) ArrangeByPoints() []Poker {
	for i := 0; i < len(p.List); i++ {
		for j := i + 1; j < len(p.List); j++ {
			if p.List[i].Points < p.List[j].Points {
				p.List[j], p.List[i] = p.List[i], p.List[j]
			}
			if p.List[i].Points == p.List[j].Points {
				if p.List[i].Color < p.List[j].Color {
					p.List[j], p.List[i] = p.List[i], p.List[j]
				}
			}
		}
	}
	return p.List
}

// ArrangeByNumber ArrangeByNumber
func (p *Pokers) ArrangeByNumber() []Poker {
	for i := 0; i < len(p.List); i++ {
		for j := i + 1; j < len(p.List); j++ {
			if p.List[i].Number < p.List[j].Number {
				p.List[j], p.List[i] = p.List[i], p.List[j]
			}
			if p.List[i].Number == p.List[j].Number {
				if p.List[i].Color < p.List[j].Color {
					p.List[j], p.List[i] = p.List[i], p.List[j]
				}
			}
		}
	}
	return p.List
}

// CreatePoker
func CreatePoker(list []int) []Poker {

	var pokerList []Poker

	for _, v := range list {

		// 牌面
		var number = int(v % 16)
		if number == 0 {
			number = 13
		}

		// 点数
		var point = int(v % 16)
		if point == 0 {
			point = 13
		}
		//if point == 1 {//A作为10时可解开
		//	point = 14
		//}

		var mark = point
		if mark >= 11 && mark <= 16 {
			mark = 10
		}
		if mark == 14 {
			mark = 1
		}

		var color = int((v-1)/16 + 1)

		var poker = Poker{Number: number, Color: color, Points: point, Mark: mark, Raw: int(v)}

		pokerList = append(pokerList, poker)
	}

	return pokerList
}
