package mahjonglib

/*
#include <stdio.h>
#include <stdlib.h>
#include "mahjong.h"
*/
import "C"
import (
	"fmt"
	"github.com/po2656233/goleaf/log"
	"strconv"
	"unsafe"
)

// OtherSit 返回其他家的座位值
func OtherSit(iSit, uSit int, isAddGang bool) string {
	faceSit := iSit + 2
	if 4 < faceSit {
		faceSit = faceSit - 4
	}
	preSit := iSit - 1
	if preSit == 0 {
		preSit = 4
	}
	nextSit := iSit + 1
	if 4 < nextSit {
		nextSit = nextSit - 4
	}
	sit := 0
	switch uSit {
	case faceSit:
		sit = 1
	case preSit:
		sit = 2
	case nextSit:
		sit = 3
	}
	if isAddGang {
		sit += 4
	}
	if sit != 0 && sit != 4 {
		return strconv.Itoa(sit)
	}
	return ""
}

func AdjustCard(card byte) int32 {
	switch card {
	case 10, 11, 12, 13, 14, 15, 16, 17, 18:
		card += 1
	case 19, 20, 21, 22, 23, 24, 25, 26, 27:
		card += 2
	case 28, 29, 30, 31:
		card += 3
	case 32:
		card = 42
	case 33:
		card = 41
	case 34:
		card = 43
	case 39, 40, 41, 42:
		card += 22
	case 35, 36, 37, 38:
		card += 16
	}
	return int32(card)
}
func isSuit(card int32) bool {
	//return card < MAHJONG_DOT_PLACEHOLDER
	return card < 30
}

// IsSequence 是否顺子
// 传入的牌必须是已排序的
// 非万、筒、条肯定不是顺
func IsSequence(tileA, tileB, tileC int32) bool {
	if !isSuit(tileA) || !isSuit(tileB) || !isSuit(tileC) {
		return false
	}
	if tileB == tileA+1 && tileC == tileB+1 {
		return true
	}
	return false
}
func IsMenCard(men, card int32) bool {
	//if men == 1 && MAHJONG_PLACEHOLDER < card && card < MAHJONG_CRAK_PLACEHOLDER {
	//	return true
	//} else if men == 2 && MAHJONG_BAM_PLACE_HOLDER < card && card < MAHJONG_DOT_PLACEHOLDER {
	//	return true
	//} else if men == 3 && MAHJONG_CRAK_PLACEHOLDER < card && card < MAHJONG_BAM_PLACE_HOLDER {
	//	return true
	//}
	if men == 1 && 0 < card && card < 10 {
		return true
	} else if men == 2 && 20 < card && card < 30 {
		return true
	} else if men == 3 && 10 < card && card < 20 {
		return true
	}
	return false
}
func CheckTing(tiles string) ([]byte, bool) {
	code := C.int(0)
	cards := C.CString(tiles)
	defer C.free(unsafe.Pointer(cards))
	ting := C.CheckTing(cards, &code)
	//fmt.Printf("可以听的牌值:%v---%v\n", C.GoString(ting), code)
	defer C.free(unsafe.Pointer(ting))
	return []byte(C.GoString(ting)), 0 < code
}

func CheckHu(tiles, hua string, quanfeng, menfeng int, isMy, isJueZhang, isHaiDi, isGang, isWordReturn bool) (fanTypes []byte, fanShu int32, ok bool) {
	mjText := map[int]string{
		1: "E",
		2: "S",
		3: "W",
		4: "N",
	}
	if _, ok1 := mjText[quanfeng]; !ok1 {
		return nil, 0, false
	}
	if _, ok1 := mjText[menfeng]; !ok1 {
		return nil, 0, false
	}
	//场况 如|EE0000 分别表示圈风、门风、是否为自摸、是否为绝张、是否为海底牌、是否为抢杠
	strAnnotation := "|" + mjText[quanfeng] + mjText[menfeng]
	if isMy {
		strAnnotation += "1"
	} else {
		strAnnotation += "0"
	}
	if isJueZhang {
		strAnnotation += "1"
	} else {
		strAnnotation += "0"
	}
	if isHaiDi {
		strAnnotation += "1"
	} else {
		strAnnotation += "0"
	}
	if isGang {
		strAnnotation += "1"
	} else {
		strAnnotation += "0"
	}
	strAnnotation += "|"
	tiles += strAnnotation + hua

	code := C.int(0)
	fanCout := C.int(0)
	needWord := C.int(0)
	if isWordReturn {
		needWord = C.int(1)
	}
	cards := C.CString(tiles)
	defer C.free(unsafe.Pointer(cards))
	hu := C.CanHu(cards, &code, &fanCout, needWord)
	defer C.free(unsafe.Pointer(hu))
	if code < 0 {
		log.Debug("当前不能胡牌的原因:%v 牌值:%v", code, tiles)
	}
	return []byte(C.GoString(hu)), int32(fanCout), code == 1

}

func RunMJ() {
	//[678m,3][789m,1][5555m,]9m2s6s6sN
	//[456p,3][456s,1]145ss55m[1111s]
	var cards *C.char = C.CString("[WWWW,7]4m4m4m5m5m7m7m7m8mCC") //C.CString("[456s,1][456s,1][456s,3]45s55m |EE0000|fah")
	defer C.free(unsafe.Pointer(cards))

	huCard := C.int(15)

	code := C.int(0)
	fanCout := C.int(0)
	isWordReturn := C.int(1)

	name := C.CheckTing(cards, &code)
	fmt.Printf("可以听的牌值:%v---%v\n", C.GoString(name), code)
	defer C.free(unsafe.Pointer(name))

	hu := C.CanHu(cards, &code, &fanCout, isWordReturn)
	defer C.free(unsafe.Pointer(hu))
	if code != 0 {
		if isWordReturn == 0 {
			str := []byte(C.GoString(hu))
			for _, v := range str {
				fmt.Println(v)
			}
			fmt.Printf("总番数:%v \n---", fanCout)
		} else {
			fmt.Printf("总番数:%v\n%v\n---", fanCout, C.GoString(hu))
		}

		//fmt.Printf("总番数:%v %v\n---", fan, C.GoString(hu))

	} else {
		fmt.Printf("总番不能胡牌!牌值:%v 当前牌值%v", cards, huCard)
	}

}
