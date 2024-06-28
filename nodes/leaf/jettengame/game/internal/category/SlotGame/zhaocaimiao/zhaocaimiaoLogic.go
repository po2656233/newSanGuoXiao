package zhaocaimiao

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	rand2 "math/rand"
	"superman/nodes/leaf/jettengame/base"
	"time"
)

const (
	RowCount        = 6
	ColCount        = 6
	AwardCol        = 3
	ElementMaxCount = 13

	ZhaoCaiMiao = 11 // 招财喵
	DuoBao      = 12 // 夺宝
	BaiDa       = 13 // 百搭
	BaseBet     = 20

	MaxLoop    = 20 // 检测次数
	ChangeRate = 30 //变框概率
)
const (
	Placeholder  = -1 //变框占位
	MarkerYin    = -3
	MarkerJin    = -4
	MarkerBaiDa  = -5
	MarkerNormal = -6 //消除占位
)

const (
	OccYin   = 0x10
	OccJin   = 0x20
	OccTwo   = 0x40 //占位
	OccThree = 0x80
	OccFour  = 0x100
)

type AwardInfo struct {
	Index   int        // 对应第一列的索引
	Element int32      // 中奖元素
	Pos     []Position // 中奖位置
}

type Position struct {
	Rx row
	Cy col
}

// CreateLines 生成一轮结果线
func CreateLines(colCount, rowCount int) map[col][]row {
	lines := make(map[col][]row, 0)
	for i := 0; i < colCount; i++ {
		for j := 0; j < rowCount; j++ {
			lines[col(i)] = append(lines[col(i)], getNewElement(ElementMaxCount))
		}
	}
	lines[0][0] = 0
	lines[ColCount-1][0] = 0
	return lines
}

// getNewElement 获取新元素
func getNewElement(max int) row {
	// [0,max) 真随机
	result, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return row(rand2.Int63n(int64(max)) + 1)
	}
	return row(result.Int64() + 1)
}

func getCtrlElement(isAward bool, firCol []row) row {
	// [0,max) 真随机
	rand2.Seed(time.Now().UnixNano())
	if isAward {
		if 1 < len(firCol) {
			index := 1 + rand2.Int()%(RowCount-1)
			return firCol[index]
		}
		return 0
	}

	noAwards := make([]row, 0)
	for i := 1; i < ElementMaxCount; i++ {
		isHave := false
		for _, r := range firCol {
			if int32(i) == originalElement(int32(r)) {
				isHave = true
				break
			}
		}

		if !isHave {
			noAwards = append(noAwards, row(i))
		}
	}
	if 0 < len(noAwards) {
		index := rand2.Int() % len(noAwards)
		return noAwards[index]
	}

	return 0
}

// sortRows 按行排列
func sortRows(lines map[col][]row) map[col][]row {
	if lines == nil || 0 == len(lines[0]) {
		return nil
	}
	rowLines := make(map[col][]row, 0)
	rowLines[0] = make([]row, 0)
	sizeCol := len(lines)
	sizeRow := len(lines[0])

	for i := 0; i < sizeCol; i++ {
		for index, r := range lines[col(i)] {
			for j := 0; j < sizeRow; j++ {
				if j == index {
					if 0 == len(rowLines[col(j)]) {
						rowLines[col(j)] = make([]row, 0)
					}
					rowLines[col(j)] = append(rowLines[col(j)], row(r))
				}
			}
		}
	}
	return rowLines
}

func originalElement(element int32) int32 {
	if element < 0 || element == BaiDa {
		return element
	}
	element = getOriginal(element, OccYin)
	element = getOriginal(element, OccJin)
	element = getOriginal(element, OccTwo)
	element = getOriginal(element, OccThree)
	element = getOriginal(element, OccFour)
	return element
}

func IsYin(element int32) bool {
	if element < 0 {
		return false
	}
	return element&OccYin == OccYin
}

func IsJin(element int32) bool {
	if element < 0 {
		return false
	}
	return element&OccJin == OccJin
}

func IsBaiDa(element int32) bool {
	if element < 0 {
		return false
	}
	return element&BaiDa == BaiDa
}

func GetKuangSize(element int32) int {
	if element < 0 {
		return 0
	}
	if element&OccTwo == OccTwo {
		return 2
	}
	if element&OccThree == OccThree {
		return 3
	}
	if element&OccFour == OccFour {
		return 4
	}
	return 0
}

func getOriginal(element, occ int32) int32 {
	if element&occ == occ {
		element -= occ
	}
	return element
}

//// GetNewLines 消同类,并生成新的结果
//func GetNewLines(awardPos map[col]row, lines map[col][]row) map[col][]row {
//	for c, r := range awardPos {
//		rows, ok := lines[c]
//		if !ok || row(len(rows)) <= r || lines[c][r] == DuoBao { // 夺宝不删除
//			continue
//		}
//
//		//删除元素
//		lines[c] = append(lines[c][:r], lines[c][r+1:]...)
//
//		//补充删除的元素
//		lines[c] = append(lines[c], getNewElement(ElementMaxCount))
//	}
//	return lines
//}

// checkAward 检测中奖
func checkAward(lines map[col][]row) []AwardInfo {
	if len(lines) < ColCount {
		return nil
	}
	// 第一列首位和最后一列首位 不纳入中奖
	lines[0][0] = 0
	lines[ColCount-1][0] = 0

	//  根据第一列的数值 从第二列开始比对、统计
	awardLines := make([]AwardInfo, 0)
	for i := 1; i < RowCount; i++ {
		element := int32(lines[0][i])
		if element < 0 {
			continue
		}
		awardPos := getAwardInfo(element, lines)
		if AwardCol <= len(awardPos) {
			//中奖
			award := AwardInfo{
				Index:   i,
				Element: element,
				Pos:     awardPos,
			}
			awardLines = append(awardLines, award)
		}
	}
	return awardLines
}

// 获取中奖位置
func getAwardInfo(element int32, lines map[col][]row) []Position {
	//从第二列开始查
	pos := make([]Position, 0)

	colNum := 0
	isHave := false
	for i := col(0); i < ColCount; i++ {
		isHave = false
		for j := row(0); j < RowCount; j++ {
			if element == originalElement(int32(lines[i][j])) {
				pos = append(pos, Position{
					Rx: j,
					Cy: i,
				})
				isHave = true
			}
		}
		if !isHave {
			break
		}
		colNum++
	}
	if colNum < AwardCol {
		return nil
	}
	return pos
}

// 占框
func changeKuang(rate int, lines map[col][]row) map[col][]row {
	size := len(lines)
	if size < ColCount {
		return nil
	}

	//if rate == 0 {
	//	rate = 100
	//}
	//rate = 100 / rate
	for i := 1; i < ColCount; i++ {
		for j := 1; j < RowCount; j++ {
			// 是否命中占格
			if lines[col(i)][j] == Placeholder || lines[col(i)][j] == DuoBao || !base.IsSatisfy(rate) {
				continue
			}

			// 命中四
			if base.IsSatisfy(rate) {
				if len(lines[col(i)]) <= j+3 { // 不满足格子数
					continue
				}
				// 不再删除 采取占位
				lines[col(i)][row(j)] += row(OccYin) + row(OccFour)
				//lines[col(i)] = append(lines[col(i)][:j+1], lines[col(i)][j+4:]...)
				lines[col(i)][j+1] = Placeholder
				lines[col(i)][j+2] = Placeholder
				lines[col(i)][j+3] = Placeholder
				break
			}
			// 命中三
			if base.IsSatisfy(rate) {
				if len(lines[col(i)]) <= j+2 { // 不满足格子数
					continue
				}
				// 不再删除 采取占位
				lines[col(i)][row(j)] += row(OccYin) + row(OccThree)
				//lines[col(i)] = append(lines[col(i)][:j+1], lines[col(i)][j+3:]...)
				lines[col(i)][j+1] = Placeholder
				lines[col(i)][j+2] = Placeholder
				break
			}
			// 命中两个
			if base.IsSatisfy(rate) {
				// 格子数是否满足
				if len(lines[col(i)]) <= j+1 { // 不满足格子数
					continue
				}
				// 不再删除 采取占位
				lines[col(i)][row(j)] += row(OccYin) + row(OccTwo)
				//lines[col(i)] = append(lines[col(i)][:j+1], lines[col(i)][j+2:]...)
				lines[col(i)][j+1] = Placeholder
			}
		}
	}

	return lines
}

// 消除并补充元素
func eraseRule(rate int, lines map[col][]row, erasePos []AwardInfo) (newLines map[col][]row, freeCount int, err error) {
	freeCount = 0
	newLines = make(map[col][]row, 0)
	for i := col(0); i < ColCount; i++ {
		if 0 == len(newLines[i]) {
			newLines[i] = make([]row, 0)
		}
		newLines[i] = lines[i]
	}

	// 标记需要删除的元素 和 免费局数
	for _, award := range erasePos {
		for _, p := range award.Pos {
			if rows, ok := newLines[p.Cy]; ok {
				if len(rows) <= int(p.Rx) {
					err = errors.New(fmt.Sprintf("第%v列的第%v行 不存在", p.Cy+1, p.Rx+1))
					return
				}

				// 标记元素
				if IsYin(int32(rows[p.Rx])) {
					rows[p.Rx] = MarkerYin
				} else if IsJin(int32(rows[p.Rx])) {
					rows[p.Rx] = MarkerJin
				} else if IsBaiDa(int32(rows[p.Rx])) {
					// 消除百搭,框内的元素也要消除
					size := GetKuangSize(int32(rows[p.Rx]))
					for i := 0; i < size; i++ {
						if int(p.Rx)+i < len(rows) {
							rows[p.Rx+row(i)] = MarkerNormal
						}
					}
					rows[p.Rx] = MarkerBaiDa
				} else if rows[p.Rx] != DuoBao { // 夺宝不消除
					rows[p.Rx] = MarkerNormal
				} else {
					freeCount++
				}
			}
		}
	}
	// 统计免费局数
	if 3 == freeCount {
		freeCount = 10
	} else if 3 < freeCount {
		freeCount = 10 + 2*(freeCount-3)
	}

	// 消除中奖元素 并补充
	// 第一行元素暂不删除,删除其他行元素 即(第一行按行消除，其他的按列消除。)
	isAward := base.IsSatisfy(rate)

	// 先消除第一列
	newLines[0] = eraseCol(isAward, newLines[0], newLines[0])

	// 根据第一列进行消除其他列,控制是否开奖
	for c, rows := range newLines {
		// 消除银框元素
		if c != 0 {
			newLines[c] = eraseCol(isAward, newLines[0], rows)
		}
	}

	// 第一行的消除
	return eraseFirstRow(newLines), freeCount, err
}

// 消除第一行元素  只删除中间的元素
func eraseFirstRow(lines map[col][]row) map[col][]row {
	newElements := make([]row, 0)
	curRows := make([]row, 0)
	for i := col(0); i < ColCount-1; i++ { // 由于最后一个为固定值0 则不纳入计算
		if lines[i][0] == MarkerNormal {
			// 第一行消除后
			newElements = append(newElements, getCtrlElement(false, lines[0]))
			//newElements = append(newElements, getNewElement(ElementMaxCount))
		} else {
			curRows = append(curRows, lines[i][0])
		}
	}
	if 0 < len(newElements) {
		curRows = append(curRows, newElements...)
		curRows = append(curRows, 0)
		for c, rows := range lines {
			rows[0] = curRows[c]
		}
	}
	return lines
}

// 消除列元素
func eraseCol(isAward bool, firstCol, rows []row) []row {
	curRows := make([]row, 0)
	marks := make([]int, 0)
	for i := 1; i < len(rows); i++ { // 保留非标记的元素
		if rows[i] < Placeholder {
			marks = append(marks, int(rows[i]))
		}
	}

	// 不需要删除
	if len(marks) == 0 {
		return rows
	}

	//
	for _, mark := range marks {
		curRows = make([]row, 0)
		for i := 1; i < len(rows); i++ { // 保留非标记的元素
			if rows[i] != row(mark) {
				curRows = append(curRows, rows[i])
			}
		}

		// 补充元素
		lossLen := RowCount - len(curRows) - 1
		if 0 < lossLen {
			newElements := make([]row, 0) // 元素补充
			newElements = append(newElements, rows[0])
			for i := 0; i < lossLen; i++ {
				element := getCtrlElement(isAward, firstCol)
				if mark == MarkerYin {
					element += OccJin
				} else if mark == MarkerJin {
					element = BaiDa
				}
				newElements = append(newElements, element)
			}
			rows = append(newElements, curRows...)
		}
	}

	return rows
}
