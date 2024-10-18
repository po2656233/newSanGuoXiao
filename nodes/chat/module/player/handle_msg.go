package player

import (
	"fmt"
	log "github.com/po2656233/superplace/logger"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/proto"
	. "superman/internal/constant"
	commMsg "superman/internal/protocol/go_file/common"
	. "superman/internal/redis_cluster"
	"superman/internal/rpc"
	sqlmodel "superman/internal/sql_model/social"
	//pb "superman/internal/protocol/go_file/common"
	chatMsg "superman/internal/protocol/go_file/chat"
	gateMsg "superman/internal/protocol/go_file/gate"

	//"superman/nodes/chat/db"
	"time"

	cproto "github.com/po2656233/superplace/net/proto"
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
	p.Local().Register(p.ClubDelMember)
}

// 获取用户基本信息
func (p *ActorPlayer) getUserInfo(uid int64) *commMsg.UserInfo {
	ret := SingleRedis().DB.Get(context.Background(), GetUserKey(uid))
	if ret.Err() == nil {
		data, _ := ret.Bytes()
		if INVALID < len(data) {
			info := commMsg.UserInfo{}
			proto.Unmarshal(data, &info)
			log.Infof("[ActorAccount] getUserInfo SingleRedis uid:%v (redis-uid:%v) info:%#v", uid, info.UserID, info)
			return &info
		}
	}

	resp, code := rpc.SendDataToAcc(p.App(), &gateMsg.GetUserInfoReq{
		Uid: uid,
	})
	if code != SUCCESS {
		log.Errorf("[getUserInfo] uid:%v fail code:%v", uid, code)
		return nil
	}
	user, ok := resp.(*gateMsg.GetUserInfoResp)
	if !ok || user == nil || user.Info == nil {
		log.Errorf("[getUserInfo] uid:%v nodata ", uid)
		return nil
	}
	if user.Info.UserID == INVALID {
		log.Errorf("[getUserInfo] uid:%v nodata ", user)
		return nil
	}
	return user.Info
}

// 获取用户基本信息
func (p *ActorPlayer) getUserNameInfo(username string) *commMsg.UserInfo {
	resp, code := rpc.SendDataToAcc(p.App(), &gateMsg.GetUserInfoReq{
		Username: username,
	})
	if code != SUCCESS {
		log.Errorf("[getUserNameInfo] username:%v fail code:%v", username, code)
		return nil
	}
	user, ok := resp.(*gateMsg.GetUserInfoResp)
	if !ok || user == nil || user.Info == nil {
		log.Errorf("[getUserNameInfo] uid:%v nodata ", username)
		return nil
	}
	if user.Info.UserID == INVALID {
		log.Errorf("[getUserInfo] uid:%v nodata ", user)
		return nil
	}
	return user.Info
}

// 获取用户基本信息
func (p *ActorPlayer) getUserBaseInfo(uid int64) *commMsg.UserFullInfo {
	user, err := p.dbComponent.GetUserBaseByID(uid) // 假设有一个方法可以根据用户ID获取用户信息
	if err != nil {
		log.Errorf("[getUserBaseInfo] uid:%v  err:%v", uid, err)
		return nil // 如果获取失败，返回nil
	}

	return user
}

// ///////////////////////////////////////////////////////////////////////////////

// FriendApply 申请加好友
func (p *ActorPlayer) FriendApply(session *cproto.Session, m *chatMsg.FriendApplyReq) {
	apply := &sqlmodel.Friendapply{
		SenderUID: session.Uid,
		TargetUID: m.TargetUid,
		Cont:      m.Cont,
		ApplyTime: time.Now().Unix(),
		Status:    Pending, // 0: 待处理
	}

	if session.Uid == m.TargetUid || m.SenderUid != session.Uid {
		rpc.SendResult(&p.ActorBase, Chat039)
		return
	}

	if err := p.dbComponent.AddFriendApply(apply); err != nil {
		log.Warnf("[FriendApply] uid:%v m:%v  err:%v", session.Uid, m, err)
		rpc.SendResult(&p.ActorBase, Chat001)
		return
	}
	resp := &chatMsg.FriendApplyResp{
		ApplyData: &chatMsg.FriendApplyReq{
			SenderUid: apply.SenderUID,
			TargetUid: apply.TargetUID,
			Cont:      apply.Cont,
		},
	}
	p.NotifyTo([]int64{apply.TargetUID, session.Uid}, resp)
}

// ClubInviteList 邀请信息查询
func (p *ActorPlayer) ClubInviteList(session *cproto.Session, m *chatMsg.ClubInviteListReq) {
	invites, err := p.dbComponent.GetClubInvitesByInviteType(session.Uid, m.InviteType)
	resp := &chatMsg.ClubInviteListResp{
		InviteType: m.InviteType,
		DataArr:    make([]*chatMsg.ClubInviteInfo, 0),
	}
	if err != nil {
		log.Warnf("[ClubInviteList] uid:%v m:%v  err:%v", session.Uid, m, err)
		rpc.SendResult(&p.ActorBase, Chat002)
		return
	}

	for _, invite := range invites {
		targetData := p.getUserInfo(invite.TargetUID)
		if targetData == nil {
			log.Warnf("[ClubInviteList] getUserInfo uid:%v m:%v  TargetUid:%v", session.Uid, m, invite.TargetUID)
		}
		senderData := p.getUserInfo(invite.SenderUID)
		if senderData == nil {
			log.Warnf("[ClubInviteList] getUserInfo uid:%v m:%v  SenderUID:%v", session.Uid, m, invite.SenderUID)
		}
		resp.DataArr = append(resp.DataArr, &chatMsg.ClubInviteInfo{
			ClubId:     invite.ClubID,
			TargetData: targetData,
			SenderData: senderData,
		})
	}
	log.Infof("[ClubInviteList] req:%v res:%v", m, resp)
	p.SendMsg(resp)
}

// 邀请对战
func (p *ActorPlayer) ChatSgxInvite(session *cproto.Session, m *chatMsg.ChatSgxInviteReq) {
	chat := &sqlmodel.Chat{
		//Channel:   m.Channel,
		SenderUID: session.Uid,
		TargetUID: m.TargetUid,
		GameEid:   m.GameEid,
		Cont:      m.Cont,
		Timestamp: time.Now().Unix(),
		MsgType:   1, // 假设1表示邀请对战
	}
	err := p.dbComponent.AddChat(chat)

	resp := &chatMsg.ChatSgxInviteResp{
		SenderUid:  chat.SenderUID,
		TargetUid:  chat.TargetUID,
		GameEid:    chat.GameEid,
		Cont:       chat.Cont,
		SenderData: p.getUserInfo(chat.SenderUID), // 获取发送者基本信息
	}
	if err != nil {
		log.Warnf("[ChatSgxInvite] uid:%v m:%v  err:%v", session.Uid, m, err)
		rpc.SendResult(&p.ActorBase, Chat003)
		return
	}
	p.SendMsg(resp)
}

// 群成员列表
func (p *ActorPlayer) ClubMembers(session *cproto.Session, m *chatMsg.ClubMembersReq) {
	members, err := p.dbComponent.GetClubMembersByClubId(m.ClubId)

	resp := &chatMsg.ClubMembersResp{
		ClubId:  m.ClubId,
		DataArr: make([]*chatMsg.ClubMemberInfo, 0),
	}
	if err != nil {
		log.Warnf("[ClubMembers] uid:%v m:%v  err:%v", session.Uid, m, err)
		rpc.SendResult(&p.ActorBase, Chat004)
		return
	}

	for _, member := range members {
		resp.DataArr = append(resp.DataArr, &chatMsg.ClubMemberInfo{
			Id:            member.ID,
			Uid:           member.UID,
			Job:           member.Job,
			Liveness:      member.Liveness,
			TotalLiveness: member.TotalLiveness,
			Score:         member.Score,
			RefereeUid:    member.RefereeUID,
			ClubId:        member.ClubID,
		})
	}
	p.SendMsg(resp)
}

// 批量处理好友申请
func (p *ActorPlayer) FriendApplyBatchDeal(session *cproto.Session, m *chatMsg.FriendApplyBatchDealReq) {
	resp := &chatMsg.FriendApplyBatchDealResp{
		TargetUid:    session.Uid,
		SenderUidArr: m.SenderUidArr,
		IsAgree:      m.IsAgree,
		FriendArr:    make([]*commMsg.UserInfo, 0),
	}

	now := time.Now().Unix()
	for _, senderUid := range m.SenderUidArr {
		apply, err := p.dbComponent.GetFriendApplyBySenderAndTarget(senderUid, session.Uid)
		if err != nil {
			log.Warnf("[FriendApplyBatchDeal]-->GetFriendApplyBySenderAndTarget uid:%v m:%v apply:%v err:%v", session.Uid, m, apply, err)
			continue
		}

		apply.Status = m.IsAgree
		err = p.dbComponent.UpdateFriendApply(apply)
		if err != nil {
			log.Warnf("[FriendApplyBatchDeal]-->UpdateFriendApply uid:%v m:%v apply:%v err:%v", session.Uid, m, apply, err)
			continue
		}
		if m.IsAgree == False {
			continue
		}
		err = p.dbComponent.AddFriend(&sqlmodel.Friend{
			UID:       session.Uid,
			FriendUID: senderUid,
			IsBlack:   False,
			AddTime:   now,
		})
		if err != nil {
			log.Warnf("[FriendApplyBatchDeal]-->AddFriend uid:%v m:%v apply:%v err:%v", session.Uid, m, apply, err)
			continue
		}
		friendInfo := p.getUserInfo(senderUid)
		if friendInfo != nil {
			resp.FriendArr = append(resp.FriendArr, friendInfo)
		}

	}

	p.SendMsg(resp)
}

// 好友列表
func (p *ActorPlayer) FriendList(session *cproto.Session, m *chatMsg.FriendListReq) {
	friends, err := p.dbComponent.GetFriendsByUid(session.Uid)

	resp := &chatMsg.FriendListResp{
		DataArr: make([]*commMsg.UserInfo, 0),
	}
	if err != nil {
		log.Warnf("[FriendList] uid:%v m:%v err:%v", session.Uid, m, err)
		rpc.SendResult(&p.ActorBase, Chat005)
		return
	}

	for _, friend := range friends {
		userInfo := p.getUserInfo(friend.FriendUID)
		if userInfo != nil {
			resp.DataArr = append(resp.DataArr, userInfo)
		}
	}
	p.SendMsg(resp)
}

// 建群
func (p *ActorPlayer) ClubNew(session *cproto.Session, m *chatMsg.ClubNewReq) {
	club := &sqlmodel.Club{
		Master:    session.Uid,
		Builder:   session.Uid,
		CreatedAt: time.Now().Unix(),
		Icon:      m.Icon,
		Name:      m.Name,
		Notice:    m.Notice,
	}

	resp := &chatMsg.ClubNewResp{
		Data: &chatMsg.ClubInfo{
			Id:        club.ID,
			Master:    club.Master,
			Builder:   club.Builder,
			CreatedAt: club.CreatedAt,
			Icon:      int32(club.Icon),
			Name:      club.Name,
			Notice:    club.Notice,
		},
	}
	err := p.dbComponent.AddClub(club)
	if err != nil {
		log.Warnf("[ClubNew] uid:%v m:%v err:%v", session.Uid, m, err)
		rpc.SendResult(&p.ActorBase, Chat006)
		return
	}
	err = p.dbComponent.AddClubMember(&sqlmodel.Clubmember{
		ClubID:   club.ID,
		UID:      club.Builder,
		Job:      1,
		Liveness: 0,
	})
	if err != nil {
		log.Warnf("[ClubNew] uid:%v m:%v err:%v", session.Uid, m, err)
		rpc.SendResult(&p.ActorBase, Chat006)
		return
	}
	p.SendMsg(resp)
}

// 处理邀请
func (p *ActorPlayer) ClubInviteDeal(session *cproto.Session, m *chatMsg.ClubInviteDealReq) {
	invite, err := p.dbComponent.GetClubInviteById(m.Id)
	if err != nil {
		log.Warnf("[ClubInviteDeal] uid:%v m:%v err:%v", session.Uid, m, err)
		rpc.SendResult(&p.ActorBase, Chat007)
		return
	}

	if invite.TargetUID != session.Uid {
		log.Warnf("[ClubInviteDeal] uid:%v m:%v err:%v", session.Uid, m, err)
		rpc.SendResult(&p.ActorBase, Chat008)
		return
	}
	//mem, err := p.dbComponent.GetClubMember(invite.ClubID, session.Uid)
	//if err != nil {
	//	log.Warnf("[ClubInviteDeal] uid:%v m:%v err:%v", session.Uid, m, err)
	//	rpc.SendResult(&p.ActorBase, Chat007)
	//	return
	//}
	//if mem.Job == 0 {
	//	rpc.SendResult(&p.ActorBase, Chat036)
	//	return
	//}

	invite.Status = m.IsAgree
	err = p.dbComponent.UpdateClubInvite(invite)
	if err != nil {
		log.Warnf("[ClubInviteDeal] uid:%v m:%v err:%v", session.Uid, m, err)
		rpc.SendResult(&p.ActorBase, Chat008)
		return
	}

	username := ""
	if info := p.getUserInfo(invite.TargetUID); info != nil {
		username = info.Account
	}

	hint := fmt.Sprintf("%v 已经拒绝了您的请求!", username)
	if m.IsAgree == True {
		err = p.dbComponent.AddClubMember(&sqlmodel.Clubmember{
			ClubID:   invite.ClubID,
			UID:      invite.TargetUID,
			Job:      0,
			Liveness: 0,
		})
		if err != nil {
			log.Warnf("[ClubNew] uid:%v m:%v err:%v", session.Uid, m, err)
			rpc.SendResult(&p.ActorBase, Chat006)
			return
		}
		hint = fmt.Sprintf("%v 已经同意了您的请求!", username)
	}

	resp := &chatMsg.ClubInviteDealResp{
		Id:      m.Id,
		IsAgree: m.IsAgree,
	}

	p.NotifyTo([]int64{invite.SenderUID}, &gateMsg.ResultResp{
		State: SUCCESS,
		Hints: hint,
	})
	p.SendMsg(resp)
}

// 删除好友
func (p *ActorPlayer) FriendDel(session *cproto.Session, m *chatMsg.FriendDelReq) {
	err := p.dbComponent.DeleteFriend(session.Uid, m.FriendUid)

	resp := &chatMsg.FriendDelResp{
		FriendUid: m.FriendUid,
	}
	if err != nil {
		log.Warnf("[FriendDel] uid:%v m:%v err:%v", session.Uid, m, err)
		rpc.SendResult(&p.ActorBase, Chat009)
		return
	}
	p.SendMsg(resp)
}

// FriendOnline 好友在线情况
func (p *ActorPlayer) FriendOnline(session *cproto.Session, m *chatMsg.FriendOnlineReq) {
	friends, err := p.dbComponent.GetFriendsIds(session.Uid)
	if err != nil {
		log.Warnf("[FriendOnline] uid:%v m:%v err:%v", session.Uid, m, err)
		rpc.SendResult(&p.ActorBase, Chat010)
		return
	}
	resp := &chatMsg.FriendOnlineResp{
		UidList: make([]int64, 0),
	}

	uids := p.GetOnline()
	size := len(friends)
	count := 0
	for _, uid := range uids {
		for _, friend := range friends {
			if uid == friend {
				resp.UidList = append(resp.UidList, friend)
				count++
				break
			}
		}
		if size <= count {
			break
		}
	}
	p.SendMsg(resp)
}

// FriendApplyList 好友申请列表
func (p *ActorPlayer) FriendApplyList(session *cproto.Session, m *chatMsg.FriendApplyListReq) {
	applies, err := p.dbComponent.GetFriendAppliesByTargetUid(session.Uid)
	if err != nil {
		log.Warnf("[FriendApplyList] uid:%v m:%v err:%v", session.Uid, m, err)
		rpc.SendResult(&p.ActorBase, Chat011)
		return
	}

	resp := &chatMsg.FriendApplyListResp{
		DataArr: make([]*commMsg.UserInfo, 0),
	}
	for _, apply := range applies {
		resp.DataArr = append(resp.DataArr, p.getUserInfo(apply.SenderUID))
	}
	p.SendMsg(resp)
}

// ClubApply 申请入群
func (p *ActorPlayer) ClubApply(session *cproto.Session, m *chatMsg.ClubApplyReq) {
	master, err := p.dbComponent.GetClubMaster(&sqlmodel.Club{
		ID: m.ClubId,
	})
	if err != nil {
		log.Warnf("[ClubApply] uid:%v m:%v err:%v", session.Uid, m, err)
		rpc.SendResult(&p.ActorBase, Chat033)
		return
	}
	if session.Uid == INVALID {
		session.Uid = p.uid
	}
	apply := &sqlmodel.Clubapply{
		ClubID:    m.ClubId,
		UID:       session.Uid,
		ApplyTime: time.Now().Unix(),
		Status:    Pending, // 0: 待处理
	}
	err = p.dbComponent.AddClubApply(apply)
	if err != nil {
		log.Warnf("[ClubApply] uid:%v m:%v err:%v", session.Uid, m, err)
		rpc.SendResult(&p.ActorBase, Chat012)
		return
	}

	resp := &chatMsg.ClubApplyResp{
		ClubId:   m.ClubId,
		ClubName: m.ClubName,
		ApplyMan: p.getUserInfo(session.Uid), // 获取申请人信息
	}
	p.NotifyTo([]int64{master}, resp)
	rpc.SendResult(&p.ActorBase, Chat031)
}

// ClubDissolve 解散群
func (p *ActorPlayer) ClubDissolve(session *cproto.Session, m *chatMsg.ClubDissolveReq) {
	// 获取所有成员
	members, err := p.dbComponent.GetClubMemberIds(m.ClubId)
	if err != nil {
		log.Warnf("[ClubDissolve] uid:%v m:%v err:%v", session.Uid, m, err)
		rpc.SendResult(&p.ActorBase, Chat013)
		return
	}

	err = p.dbComponent.DeleteClub(m.ClubId)
	if err != nil {
		log.Warnf("[ClubDissolve] uid:%v m:%v err:%v", session.Uid, m, err)
		rpc.SendResult(&p.ActorBase, Chat013)
		return
	}
	resp := &chatMsg.ClubDissolveResp{
		ClubId: m.ClubId,
	}
	p.NotifyTo(members, resp)
}

// 为成员分配积分
func (p *ActorPlayer) ClubGive(session *cproto.Session, m *chatMsg.ClubGiveReq) {
	member, err := p.dbComponent.GetClubMember(m.ClubId, m.TargetUid)
	if err != nil {
		log.Warnf("[ClubGive] uid:%v m:%v err:%v", session.Uid, m, err)
		rpc.SendResult(&p.ActorBase, Chat014)
		return
	}

	member.Score += m.Count // 假设Count是要分配的积分
	err = p.dbComponent.UpdateClubMember(member)
	if err != nil {
		log.Warnf("[ClubGive] uid:%v m:%v err:%v", session.Uid, m, err)
		rpc.SendResult(&p.ActorBase, Chat015)
		return
	}

	resp := &chatMsg.ClubGiveResp{
		ReqData:   m,
		SenderUid: session.Uid,
	}
	p.NotifyTo([]int64{session.Uid, m.TargetUid}, resp)
}

// FriendApplyBatch 批量加好友
func (p *ActorPlayer) FriendApplyBatch(session *cproto.Session, m *chatMsg.FriendApplyBatchReq) {
	resp := &chatMsg.FriendApplyBatchResp{
		ApplyData: &chatMsg.FriendApplyBatchReq{},
	}

	for _, applyUID := range m.TargetUidArr {
		friendApply := &sqlmodel.Friendapply{
			SenderUID: m.SenderUid,
			TargetUID: applyUID,
			Cont:      m.Cont,
			ApplyTime: time.Now().Unix(),
			Status:    0, // 0: 待处理
		}
		err := p.dbComponent.AddFriendApply(friendApply)
		if err != nil {
			log.Warnf("[ActorPlayer] FriendApplyBatch err:%v", err)
			continue
		}
		resp.ApplyData.SenderUid = m.SenderUid
		resp.ApplyData.TargetUidArr = append(resp.ApplyData.TargetUidArr, applyUID)
		resp.ApplyData.Cont = m.Cont
	}

	p.SendMsg(resp)
}

// FriendApplyDeal 处理好友申请
func (p *ActorPlayer) FriendApplyDeal(session *cproto.Session, m *chatMsg.FriendApplyDealReq) {
	apply, err := p.dbComponent.GetFriendApplyByUID(m.SenderUid, session.Uid)
	if err != nil {
		log.Warnf("[FriendApplyDeal] uid:%v m:%v err:%v", session.Uid, m, err)
		rpc.SendResult(&p.ActorBase, Chat016)
		return
	}

	if m.IsAgree == True {
		apply.Status = Agree
	} else if m.IsAgree == False {
		apply.Status = Reject
	}
	err = p.dbComponent.UpdateFriendApply(apply)

	if err != nil {
		log.Warnf("[FriendApplyDeal] uid:%v m:%v err:%v", session.Uid, m, err)
		rpc.SendResult(&p.ActorBase, Chat017)
		return
	}

	// 获取玩家信息
	resp := &chatMsg.FriendApplyDealResp{
		SenderUid: m.SenderUid,
		TargetUid: apply.TargetUID,
		IsAgree:   m.IsAgree,
	}
	if resp.IsAgree == True {
		err = p.dbComponent.AddFriend(&sqlmodel.Friend{
			UID:       session.Uid,
			FriendUID: apply.TargetUID,
			IsBlack:   False,
			AddTime:   time.Now().Unix(),
		})
		if err != nil {
			log.Warnf("[FriendApplyBatchDeal]-->AddFriend uid:%v m:%v apply:%v err:%v", session.Uid, m, apply, err)
		}
		resp.FriendData = p.getUserInfo(apply.TargetUID)
	}
	p.SendMsg(resp)

	// 通知对方
	if resp.IsAgree == True {
		resp.FriendData = p.getUserInfo(apply.SenderUID)
	}
	p.NotifyTo([]int64{apply.TargetUID}, resp)
}

// ClubInvite 邀请入群
func (p *ActorPlayer) ClubInvite(session *cproto.Session, m *chatMsg.ClubInviteReq) {

	if m.TargetUid == INVALID {
		userInfo := p.getUserNameInfo(m.TargetName)
		if userInfo == nil {
			rpc.SendResult(&p.ActorBase, Chat042)
			return
		}
		m.TargetUid = userInfo.UserID
	} else {
		userInfo := p.getUserInfo(m.TargetUid)
		if userInfo == nil {
			rpc.SendResult(&p.ActorBase, Chat042)
			return
		}
	}

	if session.Uid == m.TargetUid {
		rpc.SendResult(&p.ActorBase, Chat039)
		return
	}

	mem, err := p.dbComponent.GetClubMember(m.ClubId, m.TargetUid)
	if INVALID < mem.ID {
		rpc.SendResult(&p.ActorBase, Chat040)
		return
	}

	invite := &sqlmodel.Clubinvite{
		ClubID:     m.ClubId,
		SenderUID:  session.Uid,
		TargetUID:  m.TargetUid,
		InviteTime: time.Now().Unix(),
		Status:     Pending, // 0: 待处理
	}
	err = p.dbComponent.AddClubInvite(invite)
	if err != nil {
		log.Warnf("[ClubInvite] uid:%v m:%v err:%v", session.Uid, m, err)
		rpc.SendResult(&p.ActorBase, Chat018)
		return
	}

	resp := &chatMsg.ClubInviteResp{
		Data: &chatMsg.ClubInviteInfo{
			ClubId:     m.ClubId,
			SenderData: p.getUserInfo(session.Uid),
			TargetData: p.getUserInfo(m.TargetUid),
		},
	}
	rpc.SendResult(&p.ActorBase, Chat018)
	p.NotifyTo([]int64{m.TargetUid}, resp)
}

// ClubList 群列表
func (p *ActorPlayer) ClubList(session *cproto.Session, m *chatMsg.ClubListReq) {
	clubs, err := p.dbComponent.GetAllClubs()

	resp := &chatMsg.ClubListResp{
		DataArr: make([]*chatMsg.ClubInfo, 0),
	}
	if err != nil {
		log.Warnf("[ClubList] uid:%v m:%v err:%v", session.Uid, m, err)
		rpc.SendResult(&p.ActorBase, Chat019)
		return
	}

	for _, club := range clubs {
		resp.DataArr = append(resp.DataArr, &chatMsg.ClubInfo{
			Id:        club.ID,
			Master:    club.Master,
			Builder:   club.Builder,
			CreatedAt: club.CreatedAt,
			Icon:      int32(club.Icon),
			Mode:      int32(club.Mode),
			Name:      club.Name,
			Notice:    club.Notice,
		})
	}
	p.SendMsg(resp)
}

// ChatText 纯文本聊天消息
func (p *ActorPlayer) ChatText(session *cproto.Session, m *chatMsg.ChatTextReq) {
	chat := &sqlmodel.Chat{
		Channel:   m.Channel,
		SenderUID: session.Uid,
		TargetUID: m.TargetUid,
		ClubID:    m.ClubId,
		Timestamp: time.Now().Unix(),
		Cont:      m.Cont,
		MsgType:   0, // 假设0表示普通文本消息
	}
	err := p.dbComponent.AddChat(chat)

	resp := &chatMsg.ChatTextResp{
		Channel:   m.Channel,
		SenderUid: chat.SenderUID,
		TargetUid: chat.TargetUID,
		ClubId:    chat.ClubID,
		TimeStamp: chat.Timestamp,
		Cont:      chat.Cont,
	}
	if err != nil {
		log.Warnf("[ChatText] uid:%v m:%v err:%v", session.Uid, m, err)
		rpc.SendResult(&p.ActorBase, Chat020)
		return
	}

	log.Infof("uid:%v \t sid:%v", session.Uid, session.Sid)

	// 发送消息
	p.NotifyTo([]int64{session.Uid, m.TargetUid}, resp)
}

// ClubJob 切换群成员职务
func (p *ActorPlayer) ClubJob(session *cproto.Session, m *chatMsg.ClubJobReq) {

	// 检验当前用户的权限
	master, err := p.dbComponent.GetClubMember(m.ClubId, session.Uid)
	if master.Job == 0 {
		rpc.SendResult(&p.ActorBase, Chat035)
		return
	}

	if master.Job == m.Job {
		rpc.SendResult(&p.ActorBase, Chat037)
		return
	}
	if m.Job == 2 && master.Job != 1 {
		rpc.SendResult(&p.ActorBase, Chat038)
		return
	}

	if m.Job == 1 && master.Job != 1 {
		rpc.SendResult(&p.ActorBase, Chat038)
		return
	}

	member, err := p.dbComponent.GetClubMember(m.ClubId, m.Uid)
	if err != nil {
		log.Warnf("[ClubJob] uid:%v m:%v err:%v", session.Uid, m, err)
		rpc.SendResult(&p.ActorBase, Chat021)
		return
	}

	member.Job = m.Job
	err = p.dbComponent.UpdateClubMember(member)
	if err != nil {
		log.Warnf("[ClubJob] uid:%v m:%v err:%v", session.Uid, m, err)
		rpc.SendResult(&p.ActorBase, Chat022)
		return
	}

	resp := &chatMsg.ClubJobResp{
		ClubId: m.ClubId,
		Uid:    m.Uid,
		Job:    m.Job,
	}
	p.SendMsg(resp)
}

// ChatHis 聊天记录
func (p *ActorPlayer) ChatHis(session *cproto.Session, m *chatMsg.ChatHisReq) {
	resp := &chatMsg.ChatHisResp{
		Channel: m.Channel,
		Datas:   make([]*chatMsg.ChatTextResp, 0),
	}

	// 他发给别人的
	chats, err := p.dbComponent.GetChatHistoryUid(m.Channel, session.Uid, 0, 0, 0)
	if err != nil {
		rpc.SendResult(&p.ActorBase, Chat030)
		return
	}

	for _, chat := range chats {
		resp.Datas = append(resp.Datas, &chatMsg.ChatTextResp{
			Channel:   chat.Channel,
			SenderUid: chat.SenderUID,
			TargetUid: chat.TargetUID,
			ClubId:    chat.ClubID,
			TimeStamp: chat.Timestamp,
			Cont:      chat.Cont,
		})
	}
	p.SendMsg(resp)
}

// ClubApplyDeal 处理入群申请
func (p *ActorPlayer) ClubApplyDeal(session *cproto.Session, m *chatMsg.ClubApplyDealReq) {

	apply, err := p.dbComponent.GetClubApplyByID(m.Id)
	if err != nil {
		log.Warnf("[ClubApplyDeal] uid:%v m:%v err:%v", session.Uid, m, err)
		rpc.SendResult(&p.ActorBase, Chat023)
		return
	}

	if apply.Status != Pending {
		rpc.SendResult(&p.ActorBase, Chat034)
		return
	}

	// 检验当前用户的权限
	if user, err := p.dbComponent.GetClubMember(apply.ClubID, session.Uid); err == nil && user.Job == 0 {
		rpc.SendResult(&p.ActorBase, Chat035)
		return
	}

	if m.IsAgree == True {
		apply.Status = Agree
		err = p.dbComponent.AddClubMember(&sqlmodel.Clubmember{
			ClubID:   apply.ClubID,
			UID:      apply.UID,
			Job:      0,
			Liveness: 0,
		})
		if err != nil {
			log.Warnf("[ClubApplyDeal] AddClubMember uid:%v m:%v err:%v", session.Uid, m, err)
			rpc.SendResult(&p.ActorBase, Chat024)
			return
		}
	} else if m.IsAgree == False {
		apply.Status = Reject
	}

	if err = p.dbComponent.UpdateClubApply(apply); err != nil {
		log.Warnf("[ClubApplyDeal] UpdateClubApply uid:%v m:%v err:%v", session.Uid, m, err)
		rpc.SendResult(&p.ActorBase, Chat024)
		return
	}

	resp := &chatMsg.ClubApplyDealResp{
		Id:      m.Id,
		IsAgree: m.IsAgree,
	}
	p.NotifyTo([]int64{apply.UID, session.Uid}, resp)
}

// ClubMine 我的群
func (p *ActorPlayer) ClubMine(session *cproto.Session, m *chatMsg.ClubMineReq) {
	clubs, err := p.dbComponent.GetClubsByUid(session.Uid)
	resp := &chatMsg.ClubMineResp{
		DataArr: make([]*chatMsg.ClubInfo, 0),
	}
	if err != nil {
		log.Warnf("[ClubMine] uid:%v m:%v err:%v", session.Uid, m, err)
		rpc.SendResult(&p.ActorBase, Chat025)
		return
	}

	for _, club := range clubs {
		resp.DataArr = append(resp.DataArr, &chatMsg.ClubInfo{
			Id:        club.ID,
			Master:    club.Master,
			Builder:   club.Builder,
			CreatedAt: club.CreatedAt,
			Icon:      club.Icon,
			Mode:      club.Mode,
			Name:      club.Name,
			Notice:    club.Notice,
		})
	}
	p.SendMsg(resp)
}

// ClubApplyList 申请列表
func (p *ActorPlayer) ClubApplyList(session *cproto.Session, m *chatMsg.ClubApplyListReq) {

	applies, err := p.dbComponent.GetClubAppliesByClubId(m.ClubId)
	if err != nil {
		log.Warnf("[ClubApplyList] uid:%v m:%v err:%v", session.Uid, m, err)
		rpc.SendResult(&p.ActorBase, Chat026)
		return
	}
	resp := &chatMsg.ClubApplyListResp{
		ClubId: m.ClubId,
		Mans:   make([]*commMsg.ClubApplyInfo, 0),
	}
	for _, apply := range applies {
		info := p.getUserInfo(apply.UID)
		log.Infof("[ClubApplyList] apply:%v info:%v", apply, info)
		resp.Mans = append(resp.Mans, &commMsg.ClubApplyInfo{
			Id:     apply.ID,
			Status: apply.Status,
			Info:   info,
		})
	}
	log.Infof("[ClubApplyList] req:%v res:%v", m, resp)
	p.SendMsg(resp)
}

// 删除成员
func (p *ActorPlayer) ClubDelMember(session *cproto.Session, m *chatMsg.ClubDelMemberReq) {
	// 判断当前权限
	mem, err := p.dbComponent.GetClubMember(m.ClubId, session.Uid)
	if err != nil || mem.Job == INVALID {
		log.Warnf("[ClubDelMember]  GetClubMember err:%v", err)
		rpc.SendResult(&p.ActorBase, Chat036)
		return
	}
	err = p.dbComponent.DeleteClubOneMember(m.ClubId, m.Uid)
	if err != nil {
		log.Warnf("[ClubDelMember]  GetClubMember err:%v", err)
		rpc.SendResult(&p.ActorBase, Chat041)
		return
	}
	resp := &chatMsg.ClubDelMemberResp{
		ClubId: m.ClubId,
		Uid:    m.Uid,
	}
	p.NotifyTo([]int64{session.Uid, m.Uid}, resp)
}
