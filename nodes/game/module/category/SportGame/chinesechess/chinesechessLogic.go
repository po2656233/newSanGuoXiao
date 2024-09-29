package chinesechess

import (
	. "superman/internal/constant"
	gameMsg "superman/internal/protocol/go_file/game"
)

const (
	maxTimeout = 30
	maxRow     = 10
	maxCol     = 9
)
const (
	moveOk = iota + 1
	moveFail
	moveJueSha
	moveBeJiangJu
	moveJiangJu
	moveKunBi
)

func CanMove(board *gameMsg.XQBoardInfo, origin, target *gameMsg.XQGrid) int {
	if isValidMove(board, origin, target) {
		// 绝杀
		if target.Core == gameMsg.XQPiece_RedShuai || target.Core == gameMsg.XQPiece_BlackJu {
			return moveJueSha
		}
		// 被将军
		return checkJiang(board, origin, target)
	}
	return moveFail
}

// 校验棋子移动规则
func isValidMove(board *gameMsg.XQBoardInfo, origin, target *gameMsg.XQGrid) bool {
	// 根据棋子类型校验移动规则
	// 同一方不能走
	if origin.Core == gameMsg.XQPiece_NoXQPiece ||
		(gameMsg.XQPiece_RedShuai < origin.Core && gameMsg.XQPiece_RedShuai < target.Core) ||
		(origin.Core <= gameMsg.XQPiece_RedShuai && target.Core <= gameMsg.XQPiece_RedShuai) {
		return false
	}

	ok := false
	switch origin.Core {
	case gameMsg.XQPiece_RedBing:
		ok = isRedBingMove(origin, target)
	case gameMsg.XQPiece_RedPao:
		ok = isPaoMove(board, origin, target)
	case gameMsg.XQPiece_RedJu:
		ok = isJuMove(board, origin, target)
	case gameMsg.XQPiece_RedMa:
		ok = isMaMove(board, origin, target)
	case gameMsg.XQPiece_RedXiang:
		ok = isXiangMove(board, origin, target)
	case gameMsg.XQPiece_RedShi:
		ok = isShiMove(origin, target)
	case gameMsg.XQPiece_RedShuai:
		ok = isJiangMove(origin, target)
	case gameMsg.XQPiece_BlackZu:
		ok = isBlackZuMove(origin, target)
	case gameMsg.XQPiece_BlackPao:
		ok = isPaoMove(board, origin, target)
	case gameMsg.XQPiece_BlackJu:
		ok = isJuMove(board, origin, target)
	case gameMsg.XQPiece_BlackMa:
		ok = isMaMove(board, origin, target)
	case gameMsg.XQPiece_BlackXiang:
		ok = isXiangMove(board, origin, target)
	case gameMsg.XQPiece_BlackShi:
		ok = isShiMove(origin, target)
	case gameMsg.XQPiece_BlackJiang:
		ok = isJiangMove(origin, target)
	default:
	}
	return ok
}

func checkJiang(board *gameMsg.XQBoardInfo, origin, target *gameMsg.XQGrid) int {
	tempBoard := gameMsg.XQBoardInfo{
		Cells: make([]*gameMsg.XQGrid, len(board.Cells)),
	}
	copy(tempBoard.Cells, board.Cells)
	defer func() {
		tempBoard.Cells = nil
	}()
	// 红将位置
	red := gameMsg.XQGrid{}
	black := gameMsg.XQGrid{}
	for _, cell := range tempBoard.Cells {
		if cell.Row == origin.Row && cell.Col == origin.Col {
			cell.Core = gameMsg.XQPiece_NoXQPiece
		} else if cell.Row == target.Row && cell.Col == target.Col {
			cell.Core = origin.Core
		}

		if cell.Core == gameMsg.XQPiece_BlackJiang {
			black.Row = cell.Row
			black.Col = cell.Col
			black.Core = cell.Core
		} else if cell.Core == gameMsg.XQPiece_RedShuai {
			red.Row = cell.Row
			red.Col = cell.Col
			red.Core = cell.Core
		}
	}

	// 走完之后，两将军是否见面了
	if red.Col == black.Col && 0 == checkColPassBy(&tempBoard, &red, &black) {
		return moveBeJiangJu
	}
	newPos := gameMsg.XQGrid{
		Row:  target.Row,
		Col:  target.Col,
		Core: origin.Core,
	}

	var enemy *gameMsg.XQGrid
	// 对方的棋子(除仕、象外)是否可以吃到我方将军
	if gameMsg.XQPiece_RedShuai < origin.Core {
		// 黑方走完，红方是否可以吃到黑方的将军
		for _, cell := range tempBoard.Cells {
			if cell.Core <= gameMsg.XQPiece_RedShuai && cell.Core != gameMsg.XQPiece_RedShi && cell.Core != gameMsg.XQPiece_RedXiang {
				if isValidMove(&tempBoard, cell, &black) {
					return moveBeJiangJu
				}
			}
		}
		enemy = &red
	} else {
		// 红方走完，黑方是否可以吃到红方的帅
		for _, cell := range tempBoard.Cells {
			if gameMsg.XQPiece_RedShuai < cell.Core && cell.Core != gameMsg.XQPiece_BlackShi && cell.Core != gameMsg.XQPiece_BlackXiang {
				if isValidMove(&tempBoard, cell, &red) {
					return moveBeJiangJu
				}
			}
		}
		enemy = &black
	}
	if isValidMove(&tempBoard, &newPos, enemy) {
		return moveJiangJu
	}

	return moveOk
}

// 同排途中经过多少个棋子
func checkColPassBy(board *gameMsg.XQBoardInfo, origin, target *gameMsg.XQGrid) int {
	count := 0
	if origin.Col == target.Col {
		maxR := origin.Row
		minR := target.Row
		if maxR < target.Row {
			maxR = target.Row
			minR = origin.Row
		}

		for _, cell := range board.Cells {
			if origin.Col == cell.Col && minR < cell.Row && cell.Row < maxR {
				count++
			}
		}
	}
	return count
}

// 同行途中经过多少个棋子
func checkRowPassBy(board *gameMsg.XQBoardInfo, origin, target *gameMsg.XQGrid) int {
	count := 0
	if origin.Row == target.Row {
		maxC := origin.Col
		minC := target.Col
		if maxC < target.Col {
			maxC = target.Col
			minC = origin.Col
		}

		for _, cell := range board.Cells {
			if origin.Row == cell.Row && minC < cell.Col && cell.Col < maxC {
				count++
			}
		}
	}
	return count
}

//////////////////////////////////////////////////////////////////////////////////////////

func isRedBingMove(origin, target *gameMsg.XQGrid) bool {
	// 红方兵（卒）的初始位置
	if origin.Row <= 4 {
		// 未过河
		if target.Col == origin.Col && target.Row == origin.Row+1 {
			return true
		}
	} else {
		// 已过河 只能向前,且不能移出底部
		if target.Col == origin.Col && target.Row == origin.Row+1 {
			return true
		}
		if target.Row == origin.Row && (target.Col == origin.Col-1 || target.Col == origin.Col+1) {
			return true
		}
	}
	return false
}

func isBlackZuMove(origin, target *gameMsg.XQGrid) bool {
	// 红方兵（卒）的初始位置
	if origin.Row > 4 {
		// 未过河
		if target.Col == origin.Col && target.Row == origin.Row-1 {
			return true
		}
	} else {
		// 已过河 只能向前,且不能移出底部
		if target.Col == origin.Col && target.Row == origin.Row-1 {
			return true
		}
		if target.Row == origin.Row && (target.Col == origin.Col-1 || target.Col == origin.Col+1) {
			return true
		}
	}
	return false
}
func isPaoMove(board *gameMsg.XQBoardInfo, origin, target *gameMsg.XQGrid) bool {
	if origin.Row == target.Row || origin.Col == target.Col {
		if target.Core == gameMsg.XQPiece_NoXQPiece {
			// 不需要炮架
			if checkColPassBy(board, origin, target) == 0 && checkRowPassBy(board, origin, target) == 0 {
				return true
			} else if origin.Row == target.Row && checkColPassBy(board, origin, target) == 1 {
				return true
			} else if origin.Col == target.Col && checkRowPassBy(board, origin, target) == 1 {
				return true
			}
		}
	}
	return false
}
func isJuMove(board *gameMsg.XQBoardInfo, origin, target *gameMsg.XQGrid) bool {
	if origin.Row == target.Row || origin.Col == target.Col {
		if checkColPassBy(board, origin, target) == 0 && checkRowPassBy(board, origin, target) == 0 {
			return true
		}
	}
	return false
}

func isMaMove(board *gameMsg.XQBoardInfo, origin, target *gameMsg.XQGrid) bool {
	isOk := false
	row := int32(INVALID)
	col := int32(INVALID)
	if origin.Row == target.Row-1 && origin.Col == target.Col-2 {
		row = target.Row - 1
		col = target.Col - 1
		isOk = true
	}
	if origin.Row == target.Row-1 && origin.Col == target.Col+2 {
		row = target.Row - 1
		col = target.Col + 1
		isOk = true
	}
	if origin.Row == target.Row+1 && origin.Col == target.Col-2 {
		row = target.Row + 1
		col = target.Col - 1
		isOk = true
	}
	if origin.Row == target.Row+1 && origin.Col == target.Col+2 {
		row = target.Row + 1
		col = target.Col + 1
		isOk = true
	}
	//
	if origin.Row == target.Row+2 && origin.Col == target.Col+1 {
		row = target.Row + 1
		col = target.Col + 1
		isOk = true
	}
	if origin.Row == target.Row+2 && origin.Col == target.Col-1 {
		row = target.Row + 1
		col = target.Col - 1
		isOk = true
	}
	if origin.Row == target.Row-2 && origin.Col == target.Col-1 {
		row = target.Row - 1
		col = target.Col - 1
		isOk = true
	}
	if origin.Row == target.Row-2 && origin.Col == target.Col+1 {
		row = target.Row - 1
		col = target.Col + 1
		isOk = true
	}

	// 检测马脚
	if isOk {
		for _, cell := range board.Cells {
			if cell.Col == col && cell.Row == row {
				if target.Core != gameMsg.XQPiece_NoXQPiece {
					return false
				}
				break
			}
		}
	}

	return isOk
}

func isXiangMove(board *gameMsg.XQBoardInfo, origin, target *gameMsg.XQGrid) bool {
	if origin.Core == gameMsg.XQPiece_RedXiang && 4 < target.Row {
		return false
	}
	if origin.Core == gameMsg.XQPiece_BlackXiang && target.Row <= 4 {
		return false
	}
	isOk := false
	row := int32(INVALID)
	col := int32(INVALID)
	if origin.Row == target.Row-2 && origin.Col == target.Col-2 {
		row = target.Row - 1
		col = target.Col - 1
		isOk = true
	}
	if origin.Row == target.Row-2 && origin.Col == target.Col+2 {
		row = target.Row - 1
		col = target.Col + 1
		isOk = true
	}
	if origin.Row == target.Row+2 && origin.Col == target.Col-2 {
		row = target.Row + 1
		col = target.Col - 1
		isOk = true
	}
	if origin.Row == target.Row+2 && origin.Col == target.Col+2 {
		row = target.Row + 1
		col = target.Col + 1
		isOk = true
	}

	// 检测象脚
	if isOk {
		for _, cell := range board.Cells {
			if cell.Col == col && cell.Row == row {
				if target.Core != gameMsg.XQPiece_NoXQPiece {
					return false
				}
				break
			}
		}
	}

	return isOk
}
func isShiMove(origin, target *gameMsg.XQGrid) bool {
	isOk := false
	if origin.Core == gameMsg.XQPiece_RedShi && target.Row <= 2 && 2 < target.Col && target.Col < 6 {
		isOk = true
	} else if origin.Core == gameMsg.XQPiece_BlackXiang && 6 < target.Row && target.Row < maxRow && 2 < target.Col && target.Col < 6 {
		isOk = true
	}
	if isOk {
		if origin.Row == target.Row+1 && origin.Col == target.Col+1 {
			return true
		}
		if origin.Row == target.Row+1 && origin.Col == target.Col-1 {
			return true
		}
		if origin.Row == target.Row-1 && origin.Col == target.Col+1 {
			return true
		}
		if origin.Row == target.Row-1 && origin.Col == target.Col-1 {
			return true
		}
	}
	return false
}

func isJiangMove(origin, target *gameMsg.XQGrid) bool {
	isOk := false
	if origin.Core == gameMsg.XQPiece_RedShuai && target.Row <= 2 && 2 < target.Col && target.Col < 6 {
		isOk = true
	} else if origin.Core == gameMsg.XQPiece_BlackJiang && 6 < target.Row && target.Row < maxRow && 2 < target.Col && target.Col < 6 {
		isOk = true
	}
	if isOk {
		if origin.Row == target.Row && origin.Col == target.Col+1 {
			return true
		}
		if origin.Row == target.Row && origin.Col == target.Col-1 {
			return true
		}
		if origin.Row == target.Row+1 && origin.Col == target.Col {
			return true
		}
		if origin.Row == target.Row-1 && origin.Col == target.Col {
			return true
		}
	}
	return false
}
