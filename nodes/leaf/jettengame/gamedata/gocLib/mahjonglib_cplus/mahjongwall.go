package mahjonglib

import (
	"superman/nodes/leaf/jettengame/gamedata/gocLib/util"
)

// Wall 牌墙
type Wall struct {
	tiles    []int // 所有牌
	forward  int   // 前游标
	backward int   // 后游标
}

// NewWall 新建一个牌墙
func NewWall() *Wall {
	wall := &Wall{}
	return wall
}

// GetForward 获取前游标
func (wall *Wall) GetForward() int {
	return wall.forward
}

// GetBackwoad 获取后游标
func (wall *Wall) GetBackwoad() int {
	return wall.backward
}

// SetTiles 设置牌墙的牌
func (wall *Wall) SetTiles(tiles []int) {
	wall.tiles = tiles
	wall.forward = 0
	wall.backward = 0
}

// SetStart 从第几张开始
func (wall *Wall) SetStart(pos int) {
	tiles := make([]int, 0)
	for i := pos; i < len(wall.tiles); i++ {
		tiles = append(tiles, wall.tiles[i])
	}
	for j := 0; j < pos; j++ {
		tiles = append(tiles, wall.tiles[j])
	}
	wall.tiles = tiles
}

// GetTile 取某个索引的牌面值
func (wall *Wall) GetTile(index int) int {
	return wall.tiles[index]
}

// GetTiles 获取牌墙所有的牌
func (wall *Wall) GetTiles() []int {
	return wall.tiles
}

// Shuffle 洗牌
func (wall *Wall) Shuffle() {
	wall.tiles = util.ShuffleSliceInt(wall.tiles)
}

// Length 获取牌墙长度
func (wall *Wall) Length() int {
	return len(wall.tiles)
}

// RemainLength 牌墙剩余张数
func (wall *Wall) RemainLength() int {
	return wall.Length() - wall.forward - wall.backward
}

func (wall *Wall) Remains() []int {
	size := len(wall.tiles)
	remainLen := size - wall.forward - wall.backward
	return wall.tiles[size-remainLen:]
}

// IsLast 是否绝张
func (wall *Wall) IsLast(tile int) bool {
	size := len(wall.tiles)

	cards := wall.tiles[wall.forward : size-wall.backward]
	count := 0
	for _, card := range cards {
		if card == tile {
			count++
		}
	}
	return count == 0
}

// GetFrowrdNextTile 获取下一张被抓的牌的牌面值
func (wall *Wall) GetFrowrdNextTile() int {
	var index int
	// 当牌池只剩下最后一张牌时，需要根据前抓的张数和后抓的张数，拿到最后一张牌的索引
	// 如果后面抓了奇数张牌，则下张要抓的牌的索引，要向后移动一位
	if wall.forward+wall.backward == wall.Length()-1 && wall.backward%2 == 1 {
		index = wall.forward + 1
	} else {
		index = wall.forward
	}
	return wall.GetTile(index)
}

// ForwardDraw 从前面抓一张牌
func (wall *Wall) ForwardDraw() int {
	defer func() {
		wall.forward++
	}()
	return wall.GetFrowrdNextTile()
}

// ForwardDrawMulti 从前面抓N张牌
func (wall *Wall) ForwardDrawMulti(n int) []int {
	defer func() {
		wall.forward += n
	}()
	tiles := make([]int, 0, n)
	return append(tiles, wall.tiles[wall.forward:wall.forward+n]...)
}

// BackwardDraw 从后面抓牌
// 根据wall.backword的当前值计算该拿的牌的
// eg: 如果牌的总数是108张，牌是按照0 ~ 107的顺序摆的，上面的都是偶数张，下面的是奇数张
// 如果只剩一张牌，若后面拿的是偶数张牌，那么直接拿就可以了；如果后面拿的是基数张，则拿forward+1(eg:后面未抓过时，107=this.forward+1，后面抓过一张的话，107=this.forward)
// 如果wall.backword == 0, 此时从后面拿，应该是拿第106张
// 如果wall.backword == 1, 此时从后面拿，应该是拿第107张
// 如果wall.backword == 2, 此时从后面拿，应该是拿第104张
// 如果wall.backword == 3, 此时从后面拿，应该是拿第105张
func (wall *Wall) BackwardDraw() int {
	defer func() {
		wall.backward++
	}()

	var index int
	if wall.forward+wall.backward == wall.Length()-1 {
		// 如果只剩一张牌，那么就直接拿了

		if wall.backward%2 == 1 {
			index = wall.forward + 1
		} else {
			index = wall.forward
		}
	} else if wall.backward%2 == 0 {
		// 如果从后面拿了偶数张，公式为：总张数-2-已抓张数
		index = (wall.Length() - 2) - wall.backward
	} else {
		// 如果从后面拿了奇数张，公式为：总张数-已抓牌数
		index = wall.Length() - wall.backward
	}

	return wall.GetTile(index)

}

// IsAllDrawn 是否已经抓完了
func (wall *Wall) IsAllDrawn() bool {
	return wall.forward+wall.backward >= wall.Length()
}

// IsDrawn 某张牌是否被抓过
// 需要考虑后面被抓奇数张的情况，如果牌总数是108，后面第一张抓的应该是106，107还在
func (wall *Wall) IsDrawn(index int) bool {
	if index < wall.forward {
		return true
	}
	if wall.backward%2 == 0 {
		return index >= wall.Length()-wall.backward
	}
	return index >= wall.Length()-wall.backward-1 && index != wall.Length()-wall.backward
}
