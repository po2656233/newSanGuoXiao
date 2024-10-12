package db

import (
	"errors"
	sqlmodel "superman/internal/sql_model/minigame"
	. "superman/internal/utils"
	"time"
)

// //////////////////////////////以下有关写操作 强烈要求加锁////////////////////////////////////////////////////////////////////

// AddUser 完成注册 【新增一个玩家】
func (self *Component) AddUser(user sqlmodel.User) (int64, error) {
	uid := self.checkUser(user.Account)
	if 0 < uid {
		return 0, errors.New("用户已经存在")
	}
	// SignInTime
	now := time.Now()
	user.Signintime = now.Unix()
	user.CreatedAt = now
	self.Lock()
	defer self.Unlock()
	err := self.db.Table(user.TableName()).Create(&user).Error
	if !CheckError(err) {
		return 0, err
	}
	return user.ID, nil
}

func (self *Component) UpdateNickName(uid int64, nickname string) error {
	self.Lock()
	defer self.Unlock()
	user := sqlmodel.User{}
	return self.db.Table(user.TableName()).Where("id=?", uid).UpdateColumn("name", nickname).Error
}
func (self *Component) UpdateLoginTime(uid int64) error {
	self.Lock()
	defer self.Unlock()
	user := sqlmodel.User{}

	return self.db.Table(user.TableName()).Where("id=?", uid).UpdateColumn("logintime", time.Now().Unix()).Error
}
func (self *Component) UpdateLeaveTime(uid, now int64) error {
	self.Lock()
	defer self.Unlock()
	user := sqlmodel.User{}
	return self.db.Table(user.TableName()).Where("id=?", uid).UpdateColumn("leavetime", now).Error
}

///////////////////////Check//////////////////////////////////////////////

// CheckUser 获取玩家ID
func (self *Component) checkUser(account string) (uid int64) {
	user := sqlmodel.User{}
	err := self.db.Table(user.TableName()).Select("id").Where("account = ?", account).Find(&uid).Error
	if !CheckError(err) {
		return 0
	}
	return uid
}

// CheckUserInfo 获取玩家信息
func (self *Component) CheckUserInfo(account, password string) (*sqlmodel.User, error) {
	user := &sqlmodel.User{}
	query := "`account`= ? AND `password` = ?"
	err := self.db.Table(user.TableName()).Select("*").Where(query, account, password).Find(user).Error
	CheckError(err)
	return user, err
}

// CheckUserSimpInfo 获取玩家信息
func (self *Component) CheckUserSimpInfo(uid int64) (*sqlmodel.User, error) {
	user := &sqlmodel.User{}
	query := "`id`= ? "
	selectField := "id,name,account,head,face,gender,age,empiric,vip,yuanbao,coin,money"
	err := self.db.Table(user.TableName()).Select(selectField).Where(query, uid).Find(user).Error
	CheckError(err)
	return user, err
}

// CheckUserID 获取玩家ID
func (self *Component) CheckUserID(account, password string) (uid int64, err error) {
	user := &sqlmodel.User{}
	query := "`account`= ? AND `password` = ?"
	err = self.db.Table(user.TableName()).Select("id").Where(query, account, password).Find(&uid).Error
	CheckError(err)
	return
}

// CheckUserMoney 获取玩家ID
func (self *Component) CheckUserMoney(uid int64) (money int64, err error) {
	user := &sqlmodel.User{}
	err = self.db.Table(user.TableName()).Select("money").Where("id=?", uid).Find(&money).Error
	CheckError(err)
	return
}
func (self *Component) CheckUserYuanBao(uid int64) (yuanbao int64, err error) {
	user := &sqlmodel.User{}
	err = self.db.Table(user.TableName()).Select("yuanbao").Where("id=?", uid).Find(&yuanbao).Error
	CheckError(err)
	return
}
func (self *Component) CheckUserEmpiric(uid int64) (empiric int64, err error) {
	user := &sqlmodel.User{}
	err = self.db.Table(user.TableName()).Select("empiric").Where("id=?", uid).Find(&empiric).Error
	CheckError(err)
	return
}
func (self *Component) CheckUserCoin(uid int64) (coin int64, err error) {
	user := &sqlmodel.User{}
	err = self.db.Table(user.TableName()).Select("coin").Where("id=?", uid).Find(&coin).Error
	CheckError(err)
	return
}
