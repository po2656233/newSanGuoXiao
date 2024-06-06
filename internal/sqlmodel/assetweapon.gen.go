// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package sqlmodel

const TableNameAssetweapon = "assetweapon"

// Assetweapon mapped from table <assetweapon>
type Assetweapon struct {
	ID       int64 `gorm:"column:id;primaryKey;autoIncrement:true;comment:assetweapon武器" json:"id"` // assetweapon武器
	UID      int64 `gorm:"column:uid;primaryKey;comment:所得的人" json:"uid"`                           // 所得的人
	Weaponid int64 `gorm:"column:weaponid;primaryKey;comment:武器id" json:"weaponid"`                 // 武器id
	Amount   int32 `gorm:"column:amount;comment:当前拥有数量" json:"amount"`                              // 当前拥有数量
	Spending int32 `gorm:"column:spending;comment:已花费数量" json:"spending"`                           // 已花费数量
}

// TableName Assetweapon's table name
func (*Assetweapon) TableName() string {
	return TableNameAssetweapon
}