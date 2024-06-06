// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package sqlmodel

const TableNameAchievement = "achievement"

// Achievement mapped from table <achievement>
type Achievement struct {
	ID        int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:成就" json:"id"` // 成就
	Type      int32  `gorm:"column:type;comment:成就类型" json:"type"`                         // 成就类型
	UID       int64  `gorm:"column:uid;comment:用户id" json:"uid"`                           // 用户id
	Count     int32  `gorm:"column:count;comment:达成次数" json:"count"`                       // 达成次数
	Timestamp int64  `gorm:"column:timestamp;comment:达成的时间点" json:"timestamp"`             // 达成的时间点
	Remard    string `gorm:"column:remard;comment:备注" json:"remard"`                       // 备注
}

// TableName Achievement's table name
func (*Achievement) TableName() string {
	return TableNameAchievement
}