package player

import (
	log "github.com/po2656233/superplace/logger"
	"github.com/po2656233/superplace/net/parser/simple"
	"google.golang.org/protobuf/proto"
	. "superman/internal/constant"
	commMsg "superman/internal/protocol/go_file/common"
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
}

func (p *ActorPlayer) Response(v interface{}) error {
	mid, data, err := rpc.ParseProto(v.(proto.Message))
	if err != nil {
		return err
	}
	simple.Response(p, p.session, mid, data)
	return nil
}

// FriendApply 申请加好友
func (p *ActorPlayer) FriendApply(session *cproto.Session, m *chatMsg.FriendApplyReq) {
	apply := &sqlmodel.Friendapply{
		SenderUID: session.Uid,
		TargetUID: m.TargetUid,
		Cont:      m.Cont,
		ApplyTime: time.Now().Unix(),
		Status:    Pending, // 0: 待处理
	}
	if err := p.dbComponent.AddFriendApply(apply); err != nil {
		log.Warnf("[FriendApply] uid:%v m:%v  err:%v", session.Uid, m, err)
		p.Response(&gateMsg.ResultResp{
			State: FAILED,
			Hints: StatusText[Chat001],
		})
		return
	}
	resp := &chatMsg.FriendApplyResp{
		ApplyData: &chatMsg.FriendApplyReq{
			SenderUid: apply.SenderUID,
			TargetUid: apply.TargetUID,
			Cont:      apply.Cont,
		},
	}
	p.Response(resp)
}

// 邀请信息查询
func (p *ActorPlayer) ClubInviteList(session *cproto.Session, m *chatMsg.ClubInviteListReq) {
	invites, err := p.dbComponent.GetClubInvitesByInviteType(session.Uid, m.InviteType)

	resp := &chatMsg.ClubInviteListResp{
		InviteType: m.InviteType,
		DataArr:    make([]*chatMsg.ClubInviteInfo, 0),
	}
	if err != nil {
		log.Warnf("[ClubInviteList] uid:%v m:%v  err:%v", session.Uid, m, err)
		p.Response(&gateMsg.ResultResp{
			State: FAILED,
			Hints: StatusText[Chat002],
		})
		return
	}

	for _, invite := range invites {
		resp.DataArr = append(resp.DataArr, &chatMsg.ClubInviteInfo{
			ClubId: invite.ClubID,
			// 只保留协议中存在的字段
		})
	}
	p.Response(resp)
}

// 邀请对战
func (p *ActorPlayer) ChatSgxInvite(session *cproto.Session, m *chatMsg.ChatSgxInviteReq) {
	chat := &sqlmodel.Chat{
		//Channel:   m.Channel,
		SenderUID: session.Uid,
		TargetUID: m.TargetUid,
		GameEid:   m.GameEid,
		Cont:      m.Cont,
		TimeStamp: time.Now().Unix(),
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
		p.Response(&gateMsg.ResultResp{
			State: FAILED,
			Hints: StatusText[Chat003],
		})
		return
	}
	p.Response(resp)
}

// 获取用户基本信息
func (p *ActorPlayer) getUserInfo(uid int64) *commMsg.UserInfo {
	user, err := p.dbComponent.GetUserByID(uid) // 假设有一个方法可以根据用户ID获取用户信息
	if err != nil {
		log.Errorf("[getUserInfo] uid:%v  err:%v", uid, err)
		return nil // 如果获取失败，返回nil
	}
	return user
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

// 群成员列表
func (p *ActorPlayer) ClubMembers(session *cproto.Session, m *chatMsg.ClubMembersReq) {
	members, err := p.dbComponent.GetClubMembersByClubId(m.ClubId)

	resp := &chatMsg.ClubMembersResp{
		ClubId:  m.ClubId,
		DataArr: make([]*chatMsg.ClubMemberInfo, 0),
	}
	if err != nil {
		log.Warnf("[ClubMembers] uid:%v m:%v  err:%v", session.Uid, m, err)
		p.Response(&gateMsg.ResultResp{
			State: FAILED,
			Hints: StatusText[Chat004],
		})
		return
	}

	for _, member := range members {
		resp.DataArr = append(resp.DataArr, &chatMsg.ClubMemberInfo{
			Uid:           member.UID,
			Job:           int32(member.Job),
			Liveness:      int32(member.Liveness),
			TotalLiveness: int32(member.TotalLiveness),
			Score:         member.Score,
		})
	}
	p.Response(resp)
}

// 批量处理好友申请
func (p *ActorPlayer) FriendApplyBatchDeal(session *cproto.Session, m *chatMsg.FriendApplyBatchDealReq) {
	resp := &chatMsg.FriendApplyBatchDealResp{
		TargetUid:    session.Uid,
		SenderUidArr: m.SenderUidArr,
		IsAgree:      m.IsAgree,
		FriendArr:    make([]*commMsg.UserInfo, 0),
	}

	for _, senderUid := range m.SenderUidArr {
		apply, err := p.dbComponent.GetFriendApplyBySenderAndTarget(senderUid, session.Uid)
		if err != nil {
			log.Warnf("[FriendApplyBatchDeal]-->GetFriendApplyBySenderAndTarget uid:%v m:%v apply:%v err:%v", session.Uid, m, apply, err)
			continue
		}

		apply.Status = m.IsAgree
		err = p.dbComponent.UpdateFriendApply(apply)

		if m.IsAgree == 1 {
			friendInfo := p.getUserInfo(senderUid)
			if friendInfo != nil {
				resp.FriendArr = append(resp.FriendArr, friendInfo)
			}
		}
	}

	p.Response(resp)
}

// 好友列表
func (p *ActorPlayer) FriendList(session *cproto.Session, m *chatMsg.FriendListReq) {
	friends, err := p.dbComponent.GetFriendsByUid(session.Uid)

	resp := &chatMsg.FriendListResp{
		DataArr: make([]*commMsg.UserInfo, 0),
	}
	if err != nil {
		log.Warnf("[FriendList] uid:%v m:%v err:%v", session.Uid, m, err)
		p.Response(&gateMsg.ResultResp{
			State: FAILED,
			Hints: StatusText[Chat005],
		})
		return
	}

	for _, friend := range friends {
		userInfo := p.getUserInfo(friend.FriendUID)
		if userInfo != nil {
			resp.DataArr = append(resp.DataArr, userInfo)
		}
	}
	p.Response(resp)
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
	err := p.dbComponent.AddClub(club)

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
	if err != nil {
		log.Warnf("[ClubNew] uid:%v m:%v err:%v", session.Uid, m, err)
		p.Response(&gateMsg.ResultResp{
			State: FAILED,
			Hints: StatusText[Chat006],
		})
		return
	}
	p.Response(resp)
}

// 处理邀请
func (p *ActorPlayer) ClubInviteDeal(session *cproto.Session, m *chatMsg.ClubInviteDealReq) {
	invite, err := p.dbComponent.GetClubInviteByClubIdAndTargetUid(m.ClubId, session.Uid)
	if err != nil {
		log.Warnf("[ClubInviteDeal] uid:%v m:%v err:%v", session.Uid, m, err)
		p.Response(&gateMsg.ResultResp{
			State: FAILED,
			Hints: StatusText[Chat007],
		})
		return
	}

	invite.Status = m.IsAgree
	err = p.dbComponent.UpdateClubInvite(invite)

	resp := &chatMsg.ClubInviteDealResp{
		ClubId:  m.ClubId,
		IsAgree: m.IsAgree,
	}
	if err != nil {
		log.Warnf("[ClubInviteDeal] uid:%v m:%v err:%v", session.Uid, m, err)
		p.Response(&gateMsg.ResultResp{
			State: FAILED,
			Hints: StatusText[Chat008],
		})
		return
	}
	p.Response(resp)
}

// 删除好友
func (p *ActorPlayer) FriendDel(session *cproto.Session, m *chatMsg.FriendDelReq) {
	err := p.dbComponent.DeleteFriend(session.Uid, m.FriendUid)

	resp := &chatMsg.FriendDelResp{
		FriendUid: m.FriendUid,
	}
	if err != nil {
		log.Warnf("[FriendDel] uid:%v m:%v err:%v", session.Uid, m, err)
		p.Response(&gateMsg.ResultResp{
			State: FAILED,
			Hints: StatusText[Chat009],
		})
		return
	}
	p.Response(resp)
}

// 好友在线情况
func (p *ActorPlayer) FriendOnline(session *cproto.Session, m *chatMsg.FriendOnlineReq) {
	onlineStatus, err := p.dbComponent.GetFriendOnlineStatus(session.Uid)

	resp := &chatMsg.FriendOnlineResp{
		UidList: make([]int64, 0),
	}
	if err != nil {
		log.Warnf("[FriendOnline] uid:%v m:%v err:%v", session.Uid, m, err)
		p.Response(&gateMsg.ResultResp{
			State: FAILED,
			Hints: StatusText[Chat010],
		})
		return
	}

	for _, status := range onlineStatus {
		if status.Status == 1 { // 假设1表示在线
			resp.UidList = append(resp.UidList, status.Uid)
		}
	}
	p.Response(resp)
}

// 好友申请列表 todo
func (p *ActorPlayer) FriendApplyList(session *cproto.Session, m *chatMsg.FriendApplyListReq) {
	applies, err := p.dbComponent.GetFriendAppliesByTargetUid(session.Uid)

	resp := &chatMsg.FriendApplyListResp{
		DataArr: make([]*commMsg.UserInfo, 0),
	}
	if err != nil {
		log.Warnf("[FriendApplyList] uid:%v m:%v err:%v", session.Uid, m, err)
		p.Response(&gateMsg.ResultResp{
			State: FAILED,
			Hints: StatusText[Chat011],
		})
		return
	}

	for _, apply := range applies {
		resp.DataArr = append(resp.DataArr, &commMsg.UserInfo{
			UserID: apply.SenderUID,
			// .....

		})
	}
	p.Response(resp)
}

// 申请入群
func (p *ActorPlayer) ClubApply(session *cproto.Session, m *chatMsg.ClubApplyReq) {
	apply := &sqlmodel.Clubapply{
		ClubID:    m.ClubId,
		UID:       session.Uid,
		ApplyTime: time.Now().Unix(),
		Status:    Pending, // 0: 待处理
	}
	err := p.dbComponent.AddClubApply(apply)

	resp := &chatMsg.ClubApplyResp{
		ClubId:   m.ClubId,
		ClubName: m.ClubName,
		ApplyMan: p.getUserInfo(session.Uid), // 获取申请人信息
	}
	if err != nil {
		log.Warnf("[ClubApply] uid:%v m:%v err:%v", session.Uid, m, err)
		p.Response(&gateMsg.ResultResp{
			State: FAILED,
			Hints: StatusText[Chat012],
		})
		return
	}
	p.Response(resp)
}

// 解散群
func (p *ActorPlayer) ClubDissolve(session *cproto.Session, m *chatMsg.ClubDissolveReq) {
	err := p.dbComponent.DeleteClub(m.ClubId)

	resp := &chatMsg.ClubDissolveResp{
		ClubId: m.ClubId,
	}
	if err != nil {
		log.Warnf("[ClubDissolve] uid:%v m:%v err:%v", session.Uid, m, err)
		p.Response(&gateMsg.ResultResp{
			State: FAILED,
			Hints: StatusText[Chat013],
		})
		return
	}
	p.Response(resp)
}

// 为成员分配积分
func (p *ActorPlayer) ClubGive(session *cproto.Session, m *chatMsg.ClubGiveReq) {
	member, err := p.dbComponent.GetClubMember(m.ClubId, m.TargetUid)
	if err != nil {
		log.Warnf("[ClubGive] uid:%v m:%v err:%v", session.Uid, m, err)
		p.Response(&gateMsg.ResultResp{
			State: FAILED,
			Hints: StatusText[Chat014],
		})
		return
	}

	member.Score += m.Count // 假设Count是要分配的积分
	err = p.dbComponent.UpdateClubMember(member)

	resp := &chatMsg.ClubGiveResp{
		ReqData:   m,
		SenderUid: session.Uid,
	}
	if err != nil {
		log.Warnf("[ClubGive] uid:%v m:%v err:%v", session.Uid, m, err)
		p.Response(&gateMsg.ResultResp{
			State: FAILED,
			Hints: StatusText[Chat015],
		})
		return
	}
	p.Response(resp)
}

// 批量加好友 todo
func (p *ActorPlayer) FriendApplyBatch(session *cproto.Session, m *chatMsg.FriendApplyBatchReq) {
	resp := &chatMsg.FriendApplyBatchResp{
		ApplyData: &chatMsg.FriendApplyBatchReq{},
	}

	//for _, apply := range m.List {
	//	friendApply := &db.FriendApply{
	//		SenderUid: apply.SenderUid,
	//		TargetUid: apply.TargetUid,
	//		Cont:      apply.Cont,
	//		ApplyTime: time.Now().Unix(),
	//		Status:    0, // 0: 待处理
	//	}
	//	err := p.dbComponent.AddFriendApply(friendApply)
	//
	//	applyInfo := &chatMsg.FriendApplyInfo{
	//		Id:        friendApply.ID,
	//		SenderUid: friendApply.SenderUid,
	//		Cont:      friendApply.Cont,
	//		ApplyTime: friendApply.ApplyTime,
	//		Status:    int32(friendApply.Status),
	//	}
	//	if err != nil {
	//		applyInfo.Status = 2 // 假设2表示处理失败
	//		applyInfo.Msg = "申请失败: " + err.Error()
	//	}
	//	resp.DataArr = append(resp.DataArr, applyInfo)
	//}

	p.Response(resp)
}

// 处理好友申请
func (p *ActorPlayer) FriendApplyDeal(session *cproto.Session, m *chatMsg.FriendApplyDealReq) {
	apply, err := p.dbComponent.GetFriendApplyByID(m.SenderUid)
	if err != nil {
		log.Warnf("[FriendApplyDeal] uid:%v m:%v err:%v", session.Uid, m, err)
		p.Response(&gateMsg.ResultResp{
			State: FAILED,
			Hints: StatusText[Chat016],
		})
		return
	}

	//apply.Status = int(m.Status)
	err = p.dbComponent.UpdateFriendApply(apply)

	resp := &chatMsg.FriendApplyDealResp{

		//Msg: "处理成功",
	}
	if err != nil {
		log.Warnf("[FriendApplyDeal] uid:%v m:%v err:%v", session.Uid, m, err)
		p.Response(&gateMsg.ResultResp{
			State: FAILED,
			Hints: StatusText[Chat017],
		})
		return
	}
	p.Response(resp)
}

// 邀请入群
func (p *ActorPlayer) ClubInvite(session *cproto.Session, m *chatMsg.ClubInviteReq) {
	invite := &sqlmodel.Clubinvite{
		ClubID:     m.ClubId,
		SenderUID:  session.Uid,
		TargetUID:  m.TargetUid,
		InviteTime: time.Now().Unix(),
		Status:     Pending, // 0: 待处理
	}
	err := p.dbComponent.AddClubInvite(invite)

	resp := &chatMsg.ClubInviteResp{
		//Msg: "邀请成功",
		///
	}
	if err != nil {
		log.Warnf("[ClubInvite] uid:%v m:%v err:%v", session.Uid, m, err)
		p.Response(&gateMsg.ResultResp{
			State: FAILED,
			Hints: StatusText[Chat018],
		})
		return
	}
	p.Response(resp)
}

// 群列表
func (p *ActorPlayer) ClubList(session *cproto.Session, m *chatMsg.ClubListReq) {
	clubs, err := p.dbComponent.GetAllClubs()

	resp := &chatMsg.ClubListResp{
		DataArr: make([]*chatMsg.ClubInfo, 0),
	}
	if err != nil {
		log.Warnf("[ClubList] uid:%v m:%v err:%v", session.Uid, m, err)
		p.Response(&gateMsg.ResultResp{
			State: FAILED,
			Hints: StatusText[Chat019],
		})
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
	p.Response(resp)
}

// 纯文本聊天消息
func (p *ActorPlayer) ChatText(session *cproto.Session, m *chatMsg.ChatTextReq) {
	chat := &sqlmodel.Chat{
		Channel:   m.Channel,
		SenderUID: session.Uid,
		TargetUID: m.TargetUid,
		ClubID:    m.ClubId,
		TimeStamp: time.Now().Unix(),
		Cont:      m.Cont,
		MsgType:   0, // 假设0表示普通文本消息
	}
	err := p.dbComponent.AddChat(chat)

	resp := &chatMsg.ChatTextResp{
		Channel:   m.Channel,
		SenderUid: chat.SenderUID,
		TargetUid: chat.TargetUID,
		ClubId:    chat.ClubID,
		TimeStamp: chat.TimeStamp,
		Cont:      chat.Cont,
	}
	if err != nil {
		log.Warnf("[ChatText] uid:%v m:%v err:%v", session.Uid, m, err)
		p.Response(&gateMsg.ResultResp{
			State: FAILED,
			Hints: StatusText[Chat020],
		})
		return
	}
	p.Response(resp)
}

// 切换群成员职务
func (p *ActorPlayer) ClubJob(session *cproto.Session, m *chatMsg.ClubJobReq) {

	member, err := p.dbComponent.GetClubMember(m.ClubId, m.Uid)
	if err != nil {
		log.Warnf("[ClubJob] uid:%v m:%v err:%v", session.Uid, m, err)
		p.Response(&gateMsg.ResultResp{
			State: FAILED,
			Hints: StatusText[Chat021],
		})
		return
	}

	member.Job = m.Job
	err = p.dbComponent.UpdateClubMember(member)

	resp := &chatMsg.ClubJobResp{
		//Msg: "切换成功",
	}
	if err != nil {
		log.Warnf("[ClubJob] uid:%v m:%v err:%v", session.Uid, m, err)
		p.Response(&gateMsg.ResultResp{
			State: FAILED,
			Hints: StatusText[Chat022],
		})
		return
	}
	p.Response(resp)
}

// 聊天记录 todo
func (p *ActorPlayer) ChatHis(session *cproto.Session, m *chatMsg.ChatHisReq) {

	//chats, err := p.dbComponent.GetChatHistory(m.Channel, m.SenderUid, m.TargetUid, m.ClubId, m.StartTime, m.EndTime)
	//
	resp := &chatMsg.ChatHisResp{
		//DataArr: make([]*chatMsg.ChatInfo, 0),
	}
	//if err != nil {
	//	p.Response(&chatMsg.ErrorResp{
	//		ErrorCode: 1,
	//		ErrorStr:  "获取聊天记录失败: " + err.Error(),
	//	})
	//	return
	//}
	//
	//for _, chat := range chats {
	//	resp.DataArr = append(resp.DataArr, &chatMsg.ChatInfo{
	//		Channel:   int32(chat.Channel),
	//		SenderUid: chat.SenderUid,
	//		TargetUid: chat.TargetUid,
	//		ClubId:    chat.ClubId,
	//		TimeStamp: chat.TimeStamp,
	//		Cont:      chat.Cont,
	//		GameEid:   chat.GameEid,
	//		MsgType:   int32(chat.MsgType),
	//	})
	//}
	p.Response(resp)
}

// 处理入群申请
func (p *ActorPlayer) ClubApplyDeal(session *cproto.Session, m *chatMsg.ClubApplyDealReq) {

	apply, err := p.dbComponent.GetClubApplyByID(session.Uid)
	if err != nil {
		log.Warnf("[ClubApplyDeal] uid:%v m:%v err:%v", session.Uid, m, err)
		p.Response(&gateMsg.ResultResp{
			State: FAILED,
			Hints: StatusText[Chat023],
		})
		return
	}

	//apply.Status = int(m.Status)
	err = p.dbComponent.UpdateClubApply(apply)

	resp := &chatMsg.ClubApplyDealResp{
		//Msg: "处理成功",
	}
	if err != nil {
		log.Warnf("[ClubApplyDeal] uid:%v m:%v err:%v", session.Uid, m, err)
		p.Response(&gateMsg.ResultResp{
			State: FAILED,
			Hints: StatusText[Chat024],
		})
		return
	}
	p.Response(resp)
}

// 我的群
func (p *ActorPlayer) ClubMine(session *cproto.Session, m *chatMsg.ClubMineReq) {
	clubs, err := p.dbComponent.GetClubsByUid(session.Uid)
	resp := &chatMsg.ClubMineResp{
		//Uid:     m.Uid,
		DataArr: make([]*chatMsg.ClubInfo, 0),
	}
	if err != nil {
		log.Warnf("[ClubMine] uid:%v m:%v err:%v", session.Uid, m, err)
		p.Response(&gateMsg.ResultResp{
			State: FAILED,
			Hints: StatusText[Chat025],
		})
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
	p.Response(resp)
}

// 申请列表
func (p *ActorPlayer) ClubApplyList(session *cproto.Session, m *chatMsg.ClubApplyListReq) {

	applies, err := p.dbComponent.GetClubAppliesByClubId(m.ClubId)

	resp := &chatMsg.ClubApplyListResp{
		ClubId: m.ClubId,
		Mans:   make([]*commMsg.UserInfo, 0),
	}
	if err != nil {
		log.Warnf("[ClubApplyList] uid:%v m:%v err:%v", session.Uid, m, err)
		p.Response(&gateMsg.ResultResp{
			State: FAILED,
			Hints: StatusText[Chat026],
		})
		return
	}

	for _, apply := range applies {
		resp.Mans = append(resp.Mans, &commMsg.UserInfo{
			UserID: apply.ID,
		})
	}
	p.Response(resp)
}
