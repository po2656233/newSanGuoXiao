package brtoubao

const (
	//下注区域
	AREA_Big    = iota //大
	AREA_Small         //小
	AREA_Single        //单
	AREA_Double        //双
	AREA_Baozi
	AREA_Baozi_1
	AREA_Baozi_2
	AREA_Baozi_3
	AREA_Baozi_4
	AREA_Baozi_5
	AREA_Baozi_6
	AREA_4
	AREA_5
	AREA_6
	AREA_7
	AREA_8
	AREA_9
	AREA_10
	AREA_11
	AREA_12
	AREA_13
	AREA_14
	AREA_15
	AREA_16
	AREA_17

	AREA_MAX = 25 //最大区域

	Small   = 3
	Middle  = 10
	Largest = 18

	CardCount       = 3 //牌数目
	BrtoubaoHostMax = 11
)

//赔率
var Odds = map[int]int{
	AREA_Big:     1,
	AREA_Small:   1,
	AREA_Single:  1,
	AREA_Double:  1,
	AREA_Baozi:   30,
	AREA_Baozi_1: 180,
	AREA_Baozi_2: 180,
	AREA_Baozi_3: 180,
	AREA_Baozi_4: 180,
	AREA_Baozi_5: 180,
	AREA_Baozi_6: 180,
	AREA_4:       60,
	AREA_5:       30,
	AREA_6:       18,
	AREA_7:       12,
	AREA_8:       8,
	AREA_9:       6,
	AREA_10:      6,
	AREA_11:      6,
	AREA_12:      6,
	AREA_13:      8,
	AREA_14:      12,
	AREA_15:      18,
	AREA_16:      30,
	AREA_17:      60,
}
