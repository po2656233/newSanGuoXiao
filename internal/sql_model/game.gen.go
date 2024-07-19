// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package sql_model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameGame = "game"

// Game mapped from table <game>
type Game struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Kid       int32          `gorm:"column:kid;not null;comment:种类" json:"kid"`                           // 种类
	EnName    string         `gorm:"column:en_name;comment:英文名，用来匹配生成游戏" json:"en_name"`                  // 英文名，用来匹配生成游戏
	Name      string         `gorm:"column:name;comment:游戏名称" json:"name"`                                // 游戏名称
	Lessscore int64          `gorm:"column:lessscore;comment:底分" json:"lessscore"`                        // 底分
	State     int32          `gorm:"column:state;comment:状态(0未开放 1正常 2维护 3关闭)" json:"state"`              // 状态(0未开放 1正常 2维护 3关闭)
	MaxPlayer int32          `gorm:"column:max_player;default:-1;comment:最大人数(-1:无限制)" json:"max_player"` // 最大人数(-1:无限制)
	Remark    string         `gorm:"column:remark;comment:备注" json:"remark"`                              // 备注
	HowToPlay string         `gorm:"column:how_to_play;comment:玩法介绍" json:"how_to_play"`                  // 玩法介绍
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	UpdateBy  int64          `gorm:"column:update_by" json:"update_by"`
	CreateBy  int64          `gorm:"column:create_by" json:"create_by"`
}

// TableName Game's table name
func (*Game) TableName() string {
	return TableNameGame
}
