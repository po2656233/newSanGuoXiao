package brtuitongzi

import (
	protoMsg "superman/internal/protocol/gofile"
	. "superman/nodes/leaf/jettengame/game/internal/category" // 注意这里不能这样导入 "../../category" 因为本地导入是根据gopath路径设定的
)

const (
	Area_Shun   = iota //顺
	Area_Tian          //天
	Area_Di            //地
	Area_Zhuang        //庄

	AREA_MAX           = 4  //最大区域
	CardCount          = 40 //牌数目
	CardAmount         = 2
	DiceCount          = 2
	BrTuitongziHostMax = 11
)
const (
	ERROR_CARD  int32 = 1 * iota //错误的牌型
	SINGLE_CARD                  //单牌
	ERBA_CARD                    //二八杠
	DOUBLE_CARD                  //对子
)

// 赔率
var Odds = map[int]float32{
	Area_Shun:   1.00,
	Area_Tian:   1.00,
	Area_Di:     1.00,
	Area_Zhuang: 1.00,
}

// 大数定义统一置文本后
var CardListData = [CardCount]byte{
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, //筒子和白板
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, //筒子和白板
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, //筒子和白板
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, //筒子和白板
}

// 后者是否比前者大
func CompareCard(former, latter *protoMsg.CardInfo) bool {
	if nil == former && nil != latter {
		return true
	}

	if nil == latter {
		return false
	}

	if ERROR_CARD == former.CardType || ERROR_CARD == latter.CardType {
		return false
	}
	if latter.CardType == former.CardType {
		return former.CardValue < latter.CardValue
	}

	return former.CardType < latter.CardType
}

func Equal(former, latter *protoMsg.CardInfo) bool {
	return latter.CardType == former.CardType && former.CardValue == latter.CardValue
}

// 判断牌型
func JudgeCarType(cards []byte) *protoMsg.CardInfo {
	info := &protoMsg.CardInfo{}
	length := len(cards)
	if CardAmount != length {
		return info
	}
	// ->获取牌值
	info.Cards = SortCards(cards)
	value1 := GetCardValue(info.Cards[0]) * 10
	value2 := GetCardValue(info.Cards[1]) * 10
	sum := value1 + value2
	//是否是豹子
	if info.Cards[0] == info.Cards[1] {
		info.CardType = DOUBLE_CARD
	} else {
		info.CardType = SINGLE_CARD
		sum = sum % 100
		if 0 == sum {
			if (value1 == 20 || value1 == 80) && (value2 == 20 || value2 == 80) {
				info.CardType = ERBA_CARD
			} else {
				sum = 0
			}
		}
		if value1 == 100 || value2 == 100 {
			//白板算半点
			sum += 5
		}
	}
	info.CardValue = int32(sum)
	return info
}
