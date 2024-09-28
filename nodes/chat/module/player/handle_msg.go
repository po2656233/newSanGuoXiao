package player

import (
	cproto "github.com/po2656233/superplace/net/proto"
	protoMsg "superman/internal/protocol/gofile"
)

func (p *ActorPlayer) registerLocalMsg() {
	p.Local().Register(p.FriendApply)
	p.Local().Register(p.ClubInviteList)
	p.Local().Register(p.ChatSgxInvite)
	p.Local().Register(p.ClubMembers)
	p.Local().Register(p.FriendApplyBatchDeal)
	p.Local().Register(p.FriendList)
	p.Local().Register(p.ClubNew)
	p.Local().Register(p.ClubInviteDeal)
	p.Local().Register(p.FriendDel)
	p.Local().Register(p.FriendOnline)
	p.Local().Register(p.FriendApplyList)
	p.Local().Register(p.ClubApply)
	p.Local().Register(p.ClubDissolve)
	p.Local().Register(p.ClubGive)
	p.Local().Register(p.FriendApplyBatch)
	p.Local().Register(p.FriendApplyDeal)
	p.Local().Register(p.ClubInvite)
	p.Local().Register(p.ClubJob)
	p.Local().Register(p.ChatHis)
	p.Local().Register(p.ClubApplyDeal)
	p.Local().Register(p.ClubMine)
	p.Local().Register(p.ClubApplyList)
	p.Local().Register(p.ClubList)
	p.Local().Register(p.ChatText)
}

// 申请加好友 SenderUid申请加TargetUid为好友
func (p *ActorPlayer) FriendApply(session *cproto.Session, m *protoMsg.FriendApplyReq) {
	// TODO: 实现 FriendApply 处理逻辑
}

// 邀请信息查询
func (p *ActorPlayer) ClubInviteList(session *cproto.Session, m *protoMsg.ClubInviteListReq) {
	// TODO: 实现 ClubInviteList 处理逻辑
}

// 邀请对战
func (p *ActorPlayer) ChatSgxInvite(session *cproto.Session, m *protoMsg.ChatSgxInviteReq) {
	// TODO: 实现 ChatSgxInvite 处理逻辑
}

// 群成员列表
func (p *ActorPlayer) ClubMembers(session *cproto.Session, m *protoMsg.ClubMembersReq) {
	// TODO: 实现 ClubMembers 处理逻辑
}

// 批量处理好友申请
func (p *ActorPlayer) FriendApplyBatchDeal(session *cproto.Session, m *protoMsg.FriendApplyBatchDealReq) {
	// TODO: 实现 FriendApplyBatchDeal 处理逻辑
}

// 好友列表
func (p *ActorPlayer) FriendList(session *cproto.Session, m *protoMsg.FriendListReq) {
	// TODO: 实现 FriendList 处理逻辑
}

// 建群
func (p *ActorPlayer) ClubNew(session *cproto.Session, m *protoMsg.ClubNewReq) {
	// TODO: 实现 ClubNew 处理逻辑
}

// 同意/拒绝邀请
func (p *ActorPlayer) ClubInviteDeal(session *cproto.Session, m *protoMsg.ClubInviteDealReq) {
	// TODO: 实现 ClubInviteDeal 处理逻辑
}

// 删除好友
func (p *ActorPlayer) FriendDel(session *cproto.Session, m *protoMsg.FriendDelReq) {
	// TODO: 实现 FriendDel 处理逻辑
}

// 好友在线情况
func (p *ActorPlayer) FriendOnline(session *cproto.Session, m *protoMsg.FriendOnlineReq) {
	// TODO: 实现 FriendOnline 处理逻辑
}

// 好友申请列表
func (p *ActorPlayer) FriendApplyList(session *cproto.Session, m *protoMsg.FriendApplyListReq) {
	// TODO: 实现 FriendApplyList 处理逻辑
}

// 申请入群
func (p *ActorPlayer) ClubApply(session *cproto.Session, m *protoMsg.ClubApplyReq) {
	// TODO: 实现 ClubApply 处理逻辑
}

// 解散群
func (p *ActorPlayer) ClubDissolve(session *cproto.Session, m *protoMsg.ClubDissolveReq) {
	// TODO: 实现 ClubDissolve 处理逻辑
}

// 为成员分配积分
func (p *ActorPlayer) ClubGive(session *cproto.Session, m *protoMsg.ClubGiveReq) {
	// TODO: 实现 ClubGive 处理逻辑
}

// 批量加好友
func (p *ActorPlayer) FriendApplyBatch(session *cproto.Session, m *protoMsg.FriendApplyBatchReq) {
	// TODO: 实现 FriendApplyBatch 处理逻辑
}

// 处理好友申请
func (p *ActorPlayer) FriendApplyDeal(session *cproto.Session, m *protoMsg.FriendApplyDealReq) {
	// TODO: 实现 FriendApplyDeal 处理逻辑
}

// 邀请入群
func (p *ActorPlayer) ClubInvite(session *cproto.Session, m *protoMsg.ClubInviteReq) {
	// TODO: 实现 ClubInvite 处理逻辑
}

// 切换群成员职务
func (p *ActorPlayer) ClubJob(session *cproto.Session, m *protoMsg.ClubJobReq) {
	// TODO: 实现 ClubJob 处理逻辑
}

// 聊天记录
func (p *ActorPlayer) ChatHis(session *cproto.Session, m *protoMsg.ChatHisReq) {
	// TODO: 实现 ChatHis 处理逻辑
}

// 处理入群申请
func (p *ActorPlayer) ClubApplyDeal(session *cproto.Session, m *protoMsg.ClubApplyDealReq) {
	// TODO: 实现 ClubApplyDeal 处理逻辑
}

// 我的群
func (p *ActorPlayer) ClubMine(session *cproto.Session, m *protoMsg.ClubMineReq) {
	// TODO: 实现 ClubMine 处理逻辑
}

// 申请列表
func (p *ActorPlayer) ClubApplyList(session *cproto.Session, m *protoMsg.ClubApplyListReq) {
	// TODO: 实现 ClubApplyList 处理逻辑
}

// 群列表
func (p *ActorPlayer) ClubList(session *cproto.Session, m *protoMsg.ClubListReq) {
	// TODO: 实现 ClubList 处理逻辑
}

// 纯文本聊天消息
func (p *ActorPlayer) ChatText(session *cproto.Session, m *protoMsg.ChatTextReq) {
	// TODO: 实现 ChatText 处理逻辑
}
