package category

//面向子游戏,服务于子游戏

import (
	"math/rand"
	"sort"
	"superman/nodes/leaf/jettengame/manger"
	"superman/nodes/leaf/jettengame/sql/mysql"
	"superman/nodes/leaf/jettengame/sql/redis"
	"time"
)

const (
	LOGIC_MASK_COLOR = 0xF0 //花色掩码
	LOGIC_MASK_VALUE = 0x0F //数值掩码
	COLOR_Diamond    = 0x00
	COLOR_Club       = 0x10
	COLOR_Heart      = 0x20
	COLOR_Spade      = 0x30
	WAITTIME         = 5
)

// 管理类(数据库、客户端、平台、用户)
var GlobalSqlHandle = mysql.SqlHandle()
var redisHandle = redis.RedisHandle()
var GlobalSender = manger.GetClientManger()
var GlobalPlatformManger = manger.GetPlatformManger()
var GlobalPlayerManger = manger.GetPlayerManger()

// GetCardValue 获取数值
func GetCardValue(cbCardData byte) byte {
	return cbCardData & LOGIC_MASK_VALUE
}

// GetCardValues 获取所有数值
func GetCardValues(cbCardData []byte) []byte {
	var cards []byte
	for _, v := range cbCardData {
		cards = append(cards, GetCardValue(v))
	}
	return cards
}

// GetCardColor 获取花色
func GetCardColor(cbCardData byte) byte {
	return cbCardData & LOGIC_MASK_COLOR
}

// GetCardPip 获取牌点
func GetCardPip(cbCardData byte) byte {
	//计算牌点
	cbCardValue := GetCardValue(cbCardData)
	var cbPipCount byte = 0
	if cbCardValue < 10 {
		cbPipCount = cbCardValue
	}
	return cbPipCount
}

// GetCardListPip 获取所有牌的最终点数
func GetCardListPip(cbCardData []byte) byte {
	//变量定义
	var cbPipCount byte = 0

	//获取牌点
	cbCardCount := len(cbCardData)
	for i := 0; i < cbCardCount; i++ {
		cbPipCount = (GetCardPip(cbCardData[i]) + cbPipCount) % 10
	}
	return cbPipCount
}

// GetCardText 文字转换
func GetCardText(cbCardData byte) string {
	color := GetCardColor(cbCardData)
	value := GetCardValue(cbCardData)
	strTxt := string("")
	switch color {
	case 0x00:
		strTxt = "♦"
	case 0x10:
		strTxt = "♣"
	case 0x20:
		strTxt = "♥"
	case 0x30:
		strTxt = "♠"
	}

	switch value {
	case 0x00:
		return ""
	case 0x01:
		strTxt += "1"
	case 0x02:
		strTxt += "2"
	case 0x03:
		strTxt += "3"
	case 0x04:
		strTxt += "4"
	case 0x05:
		strTxt += "5"
	case 0x06:
		strTxt += "6"
	case 0x07:
		strTxt += "7"
	case 0x08:
		strTxt += "8"
	case 0x09:
		strTxt += "9"
	case 0x0A:
		strTxt += "10"
	case 0x0B:
		strTxt += "J"
	case 0x0C:
		strTxt += "Q"
	case 0x0D:
		strTxt += "K"
	case 0x0E:
		strTxt += "MinJoke"
	case 0x0F:
		strTxt += "MaxJoke"
	}
	return strTxt
}
func GetCardsText(cbCardData []byte) string {
	strText := ""
	for _, v := range cbCardData {
		if card := GetCardText(v); card != "" {
			strText += GetCardText(v) + ","
		}

	}
	return strText
}

// --------------------牌值处理-------------------------------------
// 排序
func SortCards(cards []byte) []byte {
	var localCards PlayerCards = cards
	sort.Sort(localCards)
	return localCards
}

// SortDDZCards 主要用于斗地主的牌值排序
func SortDDZCards(cards []byte) []byte {
	tempCard := make([]byte, 0)
	curValue := byte(0)
	for i := byte(0x03); i < byte(0x14); i++ {
		if i == 0x0E || i == 0x0F {
			continue
		} else if i == 0x10 {
			curValue = 0x01
		} else if i == 0x11 {
			curValue = 0x02
		} else if i == 0x12 {
			curValue = 0x0E
		} else if i == 0x13 {
			curValue = 0x0F
		} else {
			curValue = i
		}

		for _, v := range cards {
			if GetCardValue(v) == curValue {
				tempCard = append(tempCard, v)

				//花色排序
				colorCard := make([]byte, 0)
				for j := 0; j < 4; j++ {
					if 0 <= len(tempCard)-1-j && curValue == GetCardValue(tempCard[len(tempCard)-1-j]) {
						colorCard = append(colorCard, tempCard[len(tempCard)-1-j])
					}
				}

				if 1 < len(colorCard) {
					var localCards PlayerCards = colorCard
					sort.Sort(localCards)
					tempCard = append(tempCard[:len(tempCard)-len(colorCard)], colorCard...)
				}

			}
		}
	}
	return tempCard
}

// 实现排序用
type PlayerCards []byte

// Len()
func (s PlayerCards) Len() int {
	return len(s)
}

// Less():成绩将有低到高排序
func (s PlayerCards) Less(i, j int) bool {
	return s[i] < s[j]
}

// Swap()
func (s PlayerCards) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

////////////////////////

// //////////////////////牌局操作///////////////////////////////////////
// 洗牌
func Shuffle(cards []byte) []byte {
	count := len(cards)
	tempCards := make([]byte, count)
	copy(tempCards, cards)

	var index int
	var temp byte
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		index = rand.Int() % count
		temp = tempCards[i]
		tempCards[i] = tempCards[index]
		tempCards[index] = temp
	}
	return tempCards
}

// 发牌 参数:牌\座号\座位总数   座位号==索引号 idx不从0开始
func Deal(cards []byte, oneCount int, idx int) []byte {
	if idx <= 0 || oneCount <= 0 {
		return nil
	}
	size := len(cards)
	if size < oneCount*idx {
		return nil
	}

	data := make([]byte, oneCount)
	copy(data, cards[(idx-1)*oneCount:idx*oneCount])

	return data
}

// RollDice 骰子数目
func RollDice(count int) []byte {
	rand.Seed(time.Now().Unix())
	nums := make([]byte, 0)
	for i := 0; i < count; i++ {
		nums = append(nums, byte(rand.Int()%6+1))
	}
	return nums
}

// Settle 结账
func Settle(payments []int64, odds ...int32) int64 {
	minSize := len(payments)
	if len(odds) < minSize {
		minSize = len(odds)
	}

	total := int64(0)
	for i := 0; i < minSize; i++ {
		total += payments[i] * int64(odds[i])
	}

	return total / 100
}
