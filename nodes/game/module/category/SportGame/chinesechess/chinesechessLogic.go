package chinesechess

import (
	. "superman/internal/constant"
	protoMsg "superman/internal/protocol/gofile"
)

const (
	maxTimeout = 30
	maxRow     = 10
	maxCol     = 9
)

// 校验棋子移动规则
func isValidMove(board *protoMsg.XQBoardInfo, origin, target *protoMsg.XQGrid) bool {
	// 根据棋子类型校验移动规则
	// 同一方不能走
	if protoMsg.XQPiece_RedShuai < origin.Core && protoMsg.XQPiece_RedShuai < target.Core {
		return false
	}
	if origin.Core <= protoMsg.XQPiece_RedShuai && target.Core <= protoMsg.XQPiece_RedShuai {
		return false
	}
	switch origin.Core {
	case protoMsg.XQPiece_RedBing:
		return isRedBingMove(origin, target)
	case protoMsg.XQPiece_RedPao:
		return isPaoMove(board, origin, target)
	case protoMsg.XQPiece_RedJu:
		return isJuMove(board, origin, target)
	case protoMsg.XQPiece_RedMa:
		return isMaMove(board, origin, target)
	case protoMsg.XQPiece_RedXiang:
		return isXiangMove(board, origin, target)
	case protoMsg.XQPiece_RedShi:
		return isShiMove(origin, target)
	case protoMsg.XQPiece_RedShuai:
		return isJiangMove(origin, target)
	case protoMsg.XQPiece_BlackZu:
		return isBlackZuMove(origin, target)
	case protoMsg.XQPiece_BlackPao:
		return isPaoMove(board, origin, target)
	case protoMsg.XQPiece_BlackJu:
		return isJuMove(board, origin, target)
	case protoMsg.XQPiece_BlackMa:
		return isMaMove(board, origin, target)
	case protoMsg.XQPiece_BlackXiang:
		return isXiangMove(board, origin, target)
	case protoMsg.XQPiece_BlackShi:
		return isShiMove(origin, target)
	case protoMsg.XQPiece_BlackJiang:
		return isJiangMove(origin, target)
	default:
	}
	return false
}

// 同排途中经过多少个棋子
func checkColPassBy(board *protoMsg.XQBoardInfo, origin, target *protoMsg.XQGrid) int {
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
func checkRowPassBy(board *protoMsg.XQBoardInfo, origin, target *protoMsg.XQGrid) int {
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

func isRedBingMove(origin, target *protoMsg.XQGrid) bool {
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

func isBlackZuMove(origin, target *protoMsg.XQGrid) bool {
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
func isPaoMove(board *protoMsg.XQBoardInfo, origin, target *protoMsg.XQGrid) bool {
	if origin.Row == target.Row || origin.Col == target.Col {
		if target.Core == protoMsg.XQPiece_NoXQPiece {
			// 不需要炮架
			if checkColPassBy(board, origin, target) == 0 && checkRowPassBy(board, origin, target) == 0 {
				return true
			}
		}
	}
	return false
}
func isJuMove(board *protoMsg.XQBoardInfo, origin, target *protoMsg.XQGrid) bool {
	if origin.Row == target.Row || origin.Col == target.Col {
		if checkColPassBy(board, origin, target) == 0 && checkRowPassBy(board, origin, target) == 0 {
			return true
		}
	}
	return false
}

func isMaMove(board *protoMsg.XQBoardInfo, origin, target *protoMsg.XQGrid) bool {
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
				if target.Core != protoMsg.XQPiece_NoXQPiece {
					return false
				}
				break
			}
		}
	}

	return isOk
}

func isXiangMove(board *protoMsg.XQBoardInfo, origin, target *protoMsg.XQGrid) bool {
	if origin.Core == protoMsg.XQPiece_RedXiang && 4 < target.Row {
		return false
	}
	if origin.Core == protoMsg.XQPiece_BlackXiang && target.Row <= 4 {
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

	// 检测马脚
	if isOk {
		for _, cell := range board.Cells {
			if cell.Col == col && cell.Row == row {
				if target.Core != protoMsg.XQPiece_NoXQPiece {
					return false
				}
				break
			}
		}
	}

	return isOk
}
func isShiMove(origin, target *protoMsg.XQGrid) bool {
	isOk := false
	if origin.Core == protoMsg.XQPiece_RedShi && target.Row <= 2 && 2 < target.Col && target.Col < 6 {
		isOk = true
	} else if origin.Core == protoMsg.XQPiece_BlackXiang && 6 < target.Row && target.Row < maxRow && 2 < target.Col && target.Col < 6 {
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

func isJiangMove(origin, target *protoMsg.XQGrid) bool {
	isOk := false
	if origin.Core == protoMsg.XQPiece_RedShuai && target.Row <= 2 && 2 < target.Col && target.Col < 6 {
		isOk = true
	} else if origin.Core == protoMsg.XQPiece_BlackJiang && 6 < target.Row && target.Row < maxRow && 2 < target.Col && target.Col < 6 {
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
