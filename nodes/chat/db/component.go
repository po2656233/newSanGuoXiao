package db

import (
	superGorm "github.com/po2656233/superplace/components/gorm"
	"github.com/po2656233/superplace/facade"
	log "github.com/po2656233/superplace/logger"
	"gorm.io/gorm"
	. "superman/internal/constant"
	protoMsg "superman/internal/protocol/go_file/common"
	sqlmodel "superman/internal/sql_model/social"
	"superman/internal/utils"
	"sync"
)

type Component struct {
	facade.Component
	db    *gorm.DB
	curDB string
	sync.Mutex
}

func New() *Component {
	return &Component{}
}
func (c *Component) Name() string {
	return ChatDb
}

func (c *Component) Init() {

}

func (c *Component) OnAfterInit() {
	// 初始化数据库连接后的操作
	component := superGorm.NewComponent()
	name := component.Name()

	// 获取gorm组件
	gormCpt := c.App().Find(name).(*superGorm.Component)
	if gormCpt == nil {
		log.Panic("[component = %s] not found.", name)
		return
	}
	// 获取 db_id = "center_db_1" 的配置
	dbID := c.App().Settings().GetConfig(DbList).GetString(c.Name())
	if c.db != nil {
		dbObj, err := c.db.DB()
		utils.CheckError(err)
		err = dbObj.Close()
		utils.CheckError(err)
		c.db = nil
	}
	c.db = gormCpt.GetDb(dbID)
	if c.db == nil {
		log.Panic(dbID, " not found")
	}
}

func (c *Component) OnStop() {
	if c.db != nil {
		sqlDB, _ := c.db.DB()
		_ = sqlDB.Close()
	}
}

// Chat表操作
func (c *Component) AddChat(chat *sqlmodel.Chat) error {
	c.Lock()
	defer c.Unlock()
	return c.db.Create(chat).Error
}

func (c *Component) GetChatByID(id int64) (*sqlmodel.Chat, error) {
	var chat sqlmodel.Chat
	err := c.db.First(&chat, id).Error
	return &chat, err
}

func (c *Component) UpdateChat(chat *sqlmodel.Chat) error {
	c.Lock()
	defer c.Unlock()
	err := c.db.Save(chat).Error
	return err
}

func (c *Component) DeleteChat(id int64) error {
	return c.db.Delete(&sqlmodel.Chat{}, id).Error
}

// Club表操作
func (c *Component) AddClub(club *sqlmodel.Club) error {
	c.Lock()
	defer c.Unlock()
	err := c.db.Create(club).Error
	return err
}

func (c *Component) GetClubByID(id int64) (*sqlmodel.Club, error) {
	var club sqlmodel.Club
	err := c.db.First(&club, id).Error
	return &club, err
}

func (c *Component) GetClubsByUID(id int64, page, pageSize int) ([]*sqlmodel.Club, error) {
	var clubs []*sqlmodel.Club
	if pageSize <= 0 {
		pageSize = 20
	}
	if page <= 0 {
		page = 1
	}
	err := c.db.Limit(pageSize).Offset((page-1)*pageSize).Find(&clubs, "master = ?", id).Error
	return clubs, err
}

func (c *Component) UpdateClub(club *sqlmodel.Club) error {
	c.Lock()
	defer c.Unlock()
	err := c.db.Save(club).Error
	return err
}

func (c *Component) DeleteClub(id int64) error {
	return c.db.Delete(&sqlmodel.Club{}, id).Error
}

// AddClubMember 表操作
func (c *Component) AddClubMember(member *sqlmodel.Clubmember) error {
	c.Lock()
	defer c.Unlock()
	err := c.db.Create(member).Error
	return err
}

func (c *Component) GetClubMemberByID(id int64) (*sqlmodel.Clubmember, error) {
	var member sqlmodel.Clubmember
	err := c.db.First(&member, id).Error
	return &member, err
}

func (c *Component) UpdateClubMember(member *sqlmodel.Clubmember) error {
	c.Lock()
	defer c.Unlock()
	err := c.db.Save(member).Error
	return err
}

func (c *Component) DeleteClubMember(id int64) error {
	return c.db.Delete(&sqlmodel.Clubmember{}, id).Error
}

func (c *Component) DeleteClubOneMember(club, uid int64) error {
	return c.db.Delete(&sqlmodel.Clubmember{}, "club_id = ? AND uid = ?", club, uid).Error
}

// AddClubApply 表操作
func (c *Component) AddClubApply(apply *sqlmodel.Clubapply) error {
	c.Lock()
	defer c.Unlock()
	err := c.db.Create(apply).Error
	return err
}

func (c *Component) GetClubApplyByID(id int64) (*sqlmodel.Clubapply, error) {
	var apply sqlmodel.Clubapply
	err := c.db.First(&apply, id).Error
	return &apply, err
}

func (c *Component) GetClubApply(club, uid int64) (*sqlmodel.Clubapply, error) {
	var apply sqlmodel.Clubapply
	apply.ClubID = club
	apply.UID = uid
	err := c.db.First(&apply).Error
	return &apply, err
}

func (c *Component) GetClubMaster(club *sqlmodel.Club) (uid int64, err error) {
	err = c.db.Model(club).Select("master").First(&uid, club).Error
	return
}

func (c *Component) UpdateClubApply(apply *sqlmodel.Clubapply) error {
	c.Lock()
	defer c.Unlock()
	err := c.db.Save(apply).Error
	return err
}

func (c *Component) DeleteClubApply(id int64) error {
	return c.db.Delete(&sqlmodel.Clubapply{}, id).Error
}

// ClubInvite表操作
func (c *Component) AddClubInvite(invite *sqlmodel.Clubinvite) error {
	c.Lock()
	defer c.Unlock()
	err := c.db.Create(invite).Error
	return err
}

func (c *Component) GetClubInviteByID(id int64) (*sqlmodel.Clubinvite, error) {
	var invite sqlmodel.Clubinvite
	err := c.db.First(&invite, id).Error
	return &invite, err
}

func (c *Component) UpdateClubInvite(invite *sqlmodel.Clubinvite) error {
	c.Lock()
	defer c.Unlock()
	err := c.db.Save(invite).Error
	return err
}

func (c *Component) DeleteClubInvite(id int64) error {
	return c.db.Delete(&sqlmodel.Clubinvite{}, id).Error
}

// Friend表操作
func (c *Component) AddFriend(friend *sqlmodel.Friend) error {
	c.Lock()
	defer c.Unlock()
	err := c.db.Create(friend).Error
	return err
}

func (c *Component) GetFriendByID(id int64) (*sqlmodel.Friend, error) {
	var friend sqlmodel.Friend
	err := c.db.First(&friend, id).Error
	return &friend, err
}

func (c *Component) UpdateFriend(friend *sqlmodel.Friend) error {
	c.Lock()
	defer c.Unlock()
	err := c.db.Save(friend).Error
	return err
}

// FriendApply表操作
func (c *Component) AddFriendApply(apply *sqlmodel.Friendapply) error {
	c.Lock()
	defer c.Unlock()
	err := c.db.Create(apply).Error
	return err
}

func (c *Component) GetFriendApplyByID(id int64) (*sqlmodel.Friendapply, error) {
	var apply sqlmodel.Friendapply
	err := c.db.First(&apply, id).Error
	return &apply, err
}

func (c *Component) GetFriendApplyByUID(senderUID, targetUID int64) (*sqlmodel.Friendapply, error) {
	var apply sqlmodel.Friendapply
	apply.SenderUID = senderUID
	apply.TargetUID = targetUID
	err := c.db.Model(apply).First(&apply).Error
	return &apply, err
}

func (c *Component) UpdateFriendApply(apply *sqlmodel.Friendapply) error {
	c.Lock()
	defer c.Unlock()
	err := c.db.Save(apply).Error
	return err
}

func (c *Component) DeleteFriendApply(id int64) error {
	return c.db.Delete(&sqlmodel.Friendapply{}, id).Error
}

// GetClubInvitesByTargetUid 根据目标用户ID获取俱乐部邀请列表
func (c *Component) GetClubInvitesByTargetUid(targetUid int64) ([]*sqlmodel.Clubinvite, error) {
	var invites []*sqlmodel.Clubinvite
	err := c.db.Where("target_uid = ?", targetUid).Find(&invites).Error
	return invites, err
}

// GetClubMembersByClubId 根据俱乐部ID获���成员列表
func (c *Component) GetClubMembersByClubId(clubId int64) ([]*sqlmodel.Clubmember, error) {
	var members []*sqlmodel.Clubmember
	err := c.db.Model(members).Where("club_id = ?", clubId).Find(&members).Error
	return members, err
}

// GetFriendsByUid 根据用户ID获取好友列表
func (c *Component) GetFriendsByUid(uid int64) ([]*sqlmodel.Friend, error) {
	var friends []*sqlmodel.Friend
	err := c.db.Where("uid = ?", uid).Find(&friends).Error
	return friends, err
}

// GetFriendsIds 根据用户ID获取好友列表
func (c *Component) GetFriendsIds(uid int64) ([]int64, error) {
	var friends []int64
	err := c.db.Select("friend_uid").Where("uid = ?", uid).Find(&friends).Error
	return friends, err
}

// DeleteFriend 删除好友关系
func (c *Component) DeleteFriend(uid, friendUid int64) error {
	return c.db.Where("uid = ? AND friend_uid = ?", uid, friendUid).Delete(&sqlmodel.Friend{}).Error
}

// GetFriendAppliesByTargetUid 获取针对特定用户的好友申请列表
func (c *Component) GetFriendAppliesByTargetUid(targetUid int64) ([]*sqlmodel.Friendapply, error) {
	var applies []*sqlmodel.Friendapply
	err := c.db.Where("target_uid = ?", targetUid).Find(&applies).Error
	return applies, err
}

// GetClubMember 获取特定俱乐部的特定成员
func (c *Component) GetClubMember(clubId, uid int64) (*sqlmodel.Clubmember, error) {
	var member sqlmodel.Clubmember
	err := c.db.Where("club_id = ? AND uid = ?", clubId, uid).First(&member).Error
	return &member, err
}

// GetClubMembers 获取特定俱乐部的成员
func (c *Component) GetClubMembers(clubId int64) (member []*sqlmodel.Clubmember, err error) {
	err = c.db.Where("club_id = ? ", clubId).Find(&member).Error
	return
}

// GetClubMemberIds 获取特定俱乐部成员的IDs
func (c *Component) GetClubMemberIds(clubId int64) (members []int64, err error) {
	err = c.db.Select("uid").Where("club_id = ? ", clubId).Find(&members).Error
	return
}

// GetChatHistory 获取聊天历史记录
func (c *Component) GetChatHistory(channel int32, senderUid, targetUid, clubId, startTime, endTime int64) ([]*sqlmodel.Chat, error) {
	var chats []*sqlmodel.Chat
	query := c.db.Where("channel = ?", channel)
	if senderUid != 0 {
		query = query.Where("sender_uid = ?", senderUid)
	}
	if targetUid != 0 {
		query = query.Where("target_uid = ?", targetUid)
	}
	if clubId != 0 {
		query = query.Where("club_id = ?", clubId)
	}
	if startTime != 0 {
		query = query.Where("timestamp >= ?", startTime)
	}
	if endTime != 0 {
		query = query.Where("timestamp <= ?", endTime)
	}
	err := query.Find(&chats).Error
	return chats, err
}

// GetChatHistoryUid 获取聊天历史记录
func (c *Component) GetChatHistoryUid(channel int32, uid, clubId, startTime, endTime int64) ([]*sqlmodel.Chat, error) {
	var chats []*sqlmodel.Chat
	query := c.db.Where("channel = ?", channel)
	if uid != 0 {
		query = query.Where("sender_uid = ? or target_uid = ?", uid, uid)
	}
	if clubId != 0 {
		query = query.Where("club_id = ?", clubId)
	}
	if startTime != 0 {
		query = query.Where("timestamp >= ?", startTime)
	}
	if endTime != 0 {
		query = query.Where("timestamp <= ?", endTime)
	}
	err := query.Find(&chats).Error
	return chats, err
}

// GetClubsByUid 获取用户加入的俱乐部列表
func (c *Component) GetClubsByUid(uid int64) ([]*sqlmodel.Club, error) {
	var clubIds []int64
	err := c.db.Model(&sqlmodel.Clubmember{}).Where("uid = ?", uid).Pluck("club_id", &clubIds).Error
	if err != nil {
		return nil, err
	}
	var clubs []*sqlmodel.Club
	err = c.db.Where("id IN ?", clubIds).Find(&clubs).Error
	return clubs, err
}

// GetClubAppliesByClubId 获取针对特定俱乐部���申请列表
func (c *Component) GetClubAppliesByClubId(clubId int64) ([]*sqlmodel.Clubapply, error) {
	var applies []*sqlmodel.Clubapply
	err := c.db.Where("club_id = ?", clubId).Find(&applies).Error
	return applies, err
}

// GetAllClubs 获取所有俱乐部列表
func (c *Component) GetAllClubs() ([]*sqlmodel.Club, error) {
	var clubs []*sqlmodel.Club
	err := c.db.Find(&clubs).Error
	return clubs, err
}

// FriendOnlineStatus 结构体用于表示好友在线状态
type FriendOnlineStatus struct {
	Uid    int64
	Status int
}

// GetClubInvitesByInviteType 根据邀请类型获取俱乐部邀请列表
func (c *Component) GetClubInvitesByInviteType(uid int64, inviteType int32) ([]*sqlmodel.Clubinvite, error) {
	var invites []*sqlmodel.Clubinvite
	err := c.db.Where("target_uid = ? AND invite_type = ?", uid, inviteType).Find(&invites).Error
	return invites, err
}

// GetFriendApplyBySenderAndTarget 根据发送者和接收者获取好友申请
func (c *Component) GetFriendApplyBySenderAndTarget(senderUid, targetUid int64) (*sqlmodel.Friendapply, error) {
	var apply sqlmodel.Friendapply
	err := c.db.Where("sender_uid = ? AND target_uid = ?", senderUid, targetUid).First(&apply).Error
	return &apply, err
}

// GetClubInviteByClubIdAndTargetUid 根据俱乐部ID和目标用户ID获取俱乐部邀请
func (c *Component) GetClubInviteByClubIdAndTargetUid(clubId, targetUid int64) (*sqlmodel.Clubinvite, error) {
	var invite sqlmodel.Clubinvite
	err := c.db.Where("club_id = ? AND target_uid = ?", clubId, targetUid).First(&invite).Error
	return &invite, err
}

// GetClubInviteById 根据俱乐部ID和目标用户ID获取俱乐部邀请
func (c *Component) GetClubInviteById(id int64) (*sqlmodel.Clubinvite, error) {
	var invite sqlmodel.Clubinvite
	err := c.db.Model(invite).First(&invite, id).Error
	return &invite, err
}

// GetUserBaseByID 根据用户ID获取用户信息
func (c *Component) GetUserBaseByID(uid int64) (*protoMsg.UserFullInfo, error) {
	var user protoMsg.UserFullInfo
	err := c.db.Model(user).Where("id = ?", uid).First(&user).Error
	return &user, err
}

// GetUserByID 根据用户ID获取用户信息
func (c *Component) GetUserByID(uid int64) (*protoMsg.UserInfo, error) {
	var user protoMsg.UserInfo
	err := c.db.Where("id = ?", uid).First(&user).Error
	return &user, err
}
