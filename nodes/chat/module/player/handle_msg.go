package player

import (
	pb "superman/internal/protocol/go_file/common"
	"superman/nodes/chat/db"
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

// FriendApply 申请加好友
func (p *ActorPlayer) FriendApply(session *cproto.Session, m *pb.FriendApplyReq) {
	apply := &db.FriendApply{
		SenderUid: session.Uid,
		TargetUid: m.TargetUid,
		Cont:      m.Cont,
		ApplyTime: time.Now().Unix(),
		Status:    0, // 0: 待处理
	}
	err := p.dbComponent.AddFriendApply(apply)

	resp := &pb.FriendApplyResp{
		ApplyData: &pb.FriendApplyReq{
			SenderUid: apply.SenderUid,
			TargetUid: apply.TargetUid,
			Cont:      apply.Cont,
		},
	}
	if err != nil {
		p.Response(session, &pb.ErrorResp{
			ErrorCode: 1,
			ErrorStr:  "申请加好友失败: " + err.Error(),
		})
		return
	}
	p.Response(session, resp)
}

// 邀请信息查询
func (p *ActorPlayer) ClubInviteList(session *cproto.Session, m *pb.ClubInviteListReq) {
	invites, err := p.dbComponent.GetClubInvitesByInviteType(session.Uid, m.InviteType)

	resp := &pb.ClubInviteListResp{
		InviteType: m.InviteType,
		DataArr:    make([]*pb.ClubInviteInfo, 0),
	}
	if err != nil {
		p.Response(session, &pb.ErrorResp{
			ErrorCode: 1,
			ErrorStr:  "获取邀请信息失败: " + err.Error(),
		})
		return
	}

	for _, invite := range invites {
		resp.DataArr = append(resp.DataArr, &pb.ClubInviteInfo{
			ClubId: invite.ClubId,
			// 只保留协议中存在的字段
		})
	}
	p.Response(session, resp)
}

// 邀请对战
func (p *ActorPlayer) ChatSgxInvite(session *cproto.Session, m *pb.ChatSgxInviteReq) {
	chat := &db.Chat{
		//Channel:   m.Channel,
		SenderUid: session.Uid,
		TargetUid: m.TargetUid,
		GameEid:   m.GameEid,
		Cont:      m.Cont,
		TimeStamp: time.Now().Unix(),
		MsgType:   1, // 假设1表示邀请对战
	}
	err := p.dbComponent.AddChat(chat)

	resp := &pb.ChatSgxInviteResp{
		SenderUid:  chat.SenderUid,
		TargetUid:  chat.TargetUid,
		GameEid:    chat.GameEid,
		Cont:       chat.Cont,
		SenderData: p.getUserBaseInfo(chat.SenderUid), // 获取发送者基本信息
	}
	if err != nil {
		p.Response(session, &pb.ErrorResp{
			ErrorCode: 1,
			ErrorStr:  "邀请对战失败: " + err.Error(),
		})
		return
	}
	p.Response(session, resp)
}

// 获取用户基本信息
func (p *ActorPlayer) getUserInfo(uid int64) *pb.UserInfo {
	user, err := p.dbComponent.GetUserByID(uid) // 假设有一个方法可以根据用户ID获取用户信息
	if err != nil {
		return nil // 如果获取失败，返回nil
	}

	return user
}

// 获取用户基本信息
func (p *ActorPlayer) getUserBaseInfo(uid int64) *pb.UserBaseInfo {
	user, err := p.dbComponent.GetUserBaseByID(uid) // 假设有一个方法可以根据用户ID获取用户信息
	if err != nil {
		return nil // 如果获取失败，返回nil
	}

	return user
}

// 群成员列表
func (p *ActorPlayer) ClubMembers(session *cproto.Session, m *pb.ClubMembersReq) {
	members, err := p.dbComponent.GetClubMembersByClubId(m.ClubId)

	resp := &pb.ClubMembersResp{
		ClubId:  m.ClubId,
		DataArr: make([]*pb.ClubMemberInfo, 0),
	}
	if err != nil {
		p.Response(session, &pb.ErrorResp{
			ErrorCode: 1,
			ErrorStr:  "获取群成员列表失败: " + err.Error(),
		})
		return
	}

	for _, member := range members {
		resp.DataArr = append(resp.DataArr, &pb.ClubMemberInfo{
			Uid:           member.Uid,
			Job:           int32(member.Job),
			Liveness:      int32(member.Liveness),
			TotalLiveness: int32(member.TotalLiveness),
			Score:         member.Score,
		})
	}
	p.Response(session, resp)
}

// 批量处理好友申请
func (p *ActorPlayer) FriendApplyBatchDeal(session *cproto.Session, m *pb.FriendApplyBatchDealReq) {
	resp := &pb.FriendApplyBatchDealResp{
		TargetUid:    session.Uid,
		SenderUidArr: m.SenderUidArr,
		IsAgree:      m.IsAgree,
		FriendArr:    make([]*pb.UserInfo, 0),
	}

	for _, senderUid := range m.SenderUidArr {
		apply, err := p.dbComponent.GetFriendApplyBySenderAndTarget(senderUid, session.Uid)
		if err != nil {
			continue
		}

		apply.Status = int(m.IsAgree)
		err = p.dbComponent.UpdateFriendApply(apply)

		if m.IsAgree == 1 {
			friendInfo := p.getUserInfo(senderUid)
			if friendInfo != nil {
				resp.FriendArr = append(resp.FriendArr, friendInfo)
			}
		}
	}

	p.Response(session, resp)
}

// 好友列表
func (p *ActorPlayer) FriendList(session *cproto.Session, m *pb.FriendListReq) {
	friends, err := p.dbComponent.GetFriendsByUid(session.Uid)

	resp := &pb.FriendListResp{
		DataArr: make([]*pb.UserInfo, 0),
	}
	if err != nil {
		p.Response(session, &pb.ErrorResp{
			ErrorCode: 1,
			ErrorStr:  "获取好友列表失败: " + err.Error(),
		})
		return
	}

	for _, friend := range friends {
		userInfo := p.getUserInfo(friend.FriendUid)
		if userInfo != nil {
			resp.DataArr = append(resp.DataArr, userInfo)
		}
	}
	p.Response(session, resp)
}

// 建群
func (p *ActorPlayer) ClubNew(session *cproto.Session, m *pb.ClubNewReq) {
	club := &db.Club{
		Master:    session.Uid,
		Builder:   session.Uid,
		CreatedAt: time.Now().Unix(),
		Icon:      int(m.Icon),
		Name:      m.Name,
		Notice:    m.Notice,
	}
	err := p.dbComponent.AddClub(club)

	resp := &pb.ClubNewResp{
		Data: &pb.ClubInfo{
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
		p.Response(session, &pb.ErrorResp{
			ErrorCode: 1,
			ErrorStr:  "建群失败: " + err.Error(),
		})
		return
	}
	p.Response(session, resp)
}

// 处理邀请
func (p *ActorPlayer) ClubInviteDeal(session *cproto.Session, m *pb.ClubInviteDealReq) {
	invite, err := p.dbComponent.GetClubInviteByClubIdAndTargetUid(m.ClubId, session.Uid)
	if err != nil {
		p.Response(session, &pb.ErrorResp{
			ErrorCode: 1,
			ErrorStr:  "处理邀请失败: 邀请不存在",
		})
		return
	}

	invite.Status = int(m.IsAgree)
	err = p.dbComponent.UpdateClubInvite(invite)

	resp := &pb.ClubInviteDealResp{
		ClubId:  m.ClubId,
		IsAgree: m.IsAgree,
	}
	if err != nil {
		p.Response(session, &pb.ErrorResp{
			ErrorCode: 1,
			ErrorStr:  "处理邀请失败: " + err.Error(),
		})
		return
	}
	p.Response(session, resp)
}

// 删除好友
func (p *ActorPlayer) FriendDel(session *cproto.Session, m *pb.FriendDelReq) {
	err := p.dbComponent.DeleteFriend(session.Uid, m.FriendUid)

	resp := &pb.FriendDelResp{
		FriendUid: m.FriendUid,
	}
	if err != nil {
		p.Response(session, &pb.ErrorResp{
			ErrorCode: 1,
			ErrorStr:  "删除好友失败: " + err.Error(),
		})
		return
	}
	p.Response(session, resp)
}

// 好友在线情况
func (p *ActorPlayer) FriendOnline(session *cproto.Session, m *pb.FriendOnlineReq) {
	onlineStatus, err := p.dbComponent.GetFriendOnlineStatus(session.Uid)

	resp := &pb.FriendOnlineResp{
		UidList: make([]int64, 0),
	}
	if err != nil {
		p.Response(session, &pb.ErrorResp{
			ErrorCode: 1,
			ErrorStr:  "获取好友在线状态失败: " + err.Error(),
		})
		return
	}

	for _, status := range onlineStatus {
		if status.Status == 1 { // 假设1表示在线
			resp.UidList = append(resp.UidList, status.Uid)
		}
	}
	p.Response(session, resp)
}

// 好友申请列表
func (p *ActorPlayer) FriendApplyList(session *cproto.Session, m *pb.FriendApplyListReq) {
	applies, err := p.dbComponent.GetFriendAppliesByTargetUid(session.Uid)

	resp := &pb.FriendApplyListResp{
		DataArr: make([]*pb.UserInfo, 0),
	}
	if err != nil {
		p.Response(session, &pb.ErrorResp{
			ErrorCode: 1,
			ErrorStr:  "获取好友申请列表失败: " + err.Error(),
		})
		return
	}

	for _, apply := range applies {
		resp.DataArr = append(resp.DataArr, &pb.UserInfo{
			UserID: apply.SenderUid,
			// .....

		})
	}
	p.Response(session, resp)
}

// 申请入群
func (p *ActorPlayer) ClubApply(session *cproto.Session, m *pb.ClubApplyReq) {
	apply := &db.ClubApply{
		ClubId:    m.ClubId,
		Uid:       session.Uid,
		ApplyTime: time.Now().Unix(),
		Status:    0, // 0: 待处理
	}
	err := p.dbComponent.AddClubApply(apply)

	resp := &pb.ClubApplyResp{
		ClubId:   m.ClubId,
		ClubName: m.ClubName,
		ApplyMan: p.getUserInfo(session.Uid), // 获取申请人信息
	}
	if err != nil {
		p.Response(session, &pb.ErrorResp{
			ErrorCode: 1,
			ErrorStr:  "申请入群失败: " + err.Error(),
		})
		return
	}
	p.Response(session, resp)
}

// 解散群
func (p *ActorPlayer) ClubDissolve(session *cproto.Session, m *pb.ClubDissolveReq) {
	err := p.dbComponent.DeleteClub(m.ClubId)

	resp := &pb.ClubDissolveResp{
		ClubId: m.ClubId,
	}
	if err != nil {
		p.Response(session, &pb.ErrorResp{
			ErrorCode: 1,
			ErrorStr:  "解散群失败: " + err.Error(),
		})
		return
	}
	p.Response(session, resp)
}

// 为成员分配积分
func (p *ActorPlayer) ClubGive(session *cproto.Session, m *pb.ClubGiveReq) {
	member, err := p.dbComponent.GetClubMember(m.ClubId, m.TargetUid)
	if err != nil {
		p.Response(session, &pb.ErrorResp{
			ErrorCode: 1,
			ErrorStr:  "分配积分失败: 成员不存在",
		})
		return
	}

	member.Score += m.Count // 假设Count是要分配的积分
	err = p.dbComponent.UpdateClubMember(member)

	resp := &pb.ClubGiveResp{
		ReqData:   m,
		SenderUid: session.Uid,
	}
	if err != nil {
		p.Response(session, &pb.ErrorResp{
			ErrorCode: 1,
			ErrorStr:  "分配积分失败: " + err.Error(),
		})
		return
	}
	p.Response(session, resp)
}

// 批量加好友
func (p *ActorPlayer) FriendApplyBatch(session *cproto.Session, m *pb.FriendApplyBatchReq) {
	resp := &pb.FriendApplyBatchResp{
		ApplyData: &pb.FriendApplyBatchReq{},
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
	//	applyInfo := &pb.FriendApplyInfo{
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

	p.Response(session, resp)
}

// 处理好友申请
func (p *ActorPlayer) FriendApplyDeal(session *cproto.Session, m *pb.FriendApplyDealReq) {
	apply, err := p.dbComponent.GetFriendApplyByID(m.SenderUid)
	if err != nil {
		p.Response(session, &pb.ErrorResp{
			ErrorCode: 1,
			ErrorStr:  "处理好友申请失败: 申请不存在",
		})
		return
	}

	//apply.Status = int(m.Status)
	err = p.dbComponent.UpdateFriendApply(apply)

	resp := &pb.FriendApplyDealResp{

		//Msg: "处理成功",
	}
	if err != nil {
		p.Response(session, &pb.ErrorResp{
			ErrorCode: 1,
			ErrorStr:  "处理好友申请失败: " + err.Error(),
		})
		return
	}
	p.Response(session, resp)
}

// 邀请入群
func (p *ActorPlayer) ClubInvite(session *cproto.Session, m *pb.ClubInviteReq) {
	invite := &db.ClubInvite{
		ClubId:     m.ClubId,
		SenderUid:  session.Uid,
		TargetUid:  m.TargetUid,
		InviteTime: time.Now().Unix(),
		Status:     0, // 0: 待处理
	}
	err := p.dbComponent.AddClubInvite(invite)

	resp := &pb.ClubInviteResp{
		//Msg: "邀请成功",
		///
	}
	if err != nil {
		p.Response(session, &pb.ErrorResp{
			ErrorCode: 1,
			ErrorStr:  "邀请入群失败: " + err.Error(),
		})
		return
	}
	p.Response(session, resp)
}

// 群列表
func (p *ActorPlayer) ClubList(session *cproto.Session, m *pb.ClubListReq) {
	clubs, err := p.dbComponent.GetAllClubs()

	resp := &pb.ClubListResp{
		DataArr: make([]*pb.ClubInfo, 0),
	}
	if err != nil {
		p.Response(session, &pb.ErrorResp{
			ErrorCode: 1,
			ErrorStr:  "获取群列表失败: " + err.Error(),
		})
		return
	}

	for _, club := range clubs {
		resp.DataArr = append(resp.DataArr, &pb.ClubInfo{
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
	p.Response(session, resp)
}

// 纯文本聊天消息
func (p *ActorPlayer) ChatText(session *cproto.Session, m *pb.ChatTextReq) {
	chat := &db.Chat{
		Channel:   int(m.Channel),
		SenderUid: session.Uid,
		TargetUid: m.TargetUid,
		ClubId:    m.ClubId,
		TimeStamp: time.Now().Unix(),
		Cont:      m.Cont,
		MsgType:   0, // 假设0表示普通文本消息
	}
	err := p.dbComponent.AddChat(chat)

	resp := &pb.ChatTextResp{
		Channel:   m.Channel,
		SenderUid: chat.SenderUid,
		TargetUid: chat.TargetUid,
		ClubId:    chat.ClubId,
		TimeStamp: chat.TimeStamp,
		Cont:      chat.Cont,
	}
	if err != nil {
		p.Response(session, &pb.ErrorResp{
			ErrorCode: 1,
			ErrorStr:  "发送文本消息失败: " + err.Error(),
		})
		return
	}
	p.Response(session, resp)
}

// 切换群成员职务
func (p *ActorPlayer) ClubJob(session *cproto.Session, m *pb.ClubJobReq) {

	member, err := p.dbComponent.GetClubMember(m.ClubId, m.Uid)
	if err != nil {
		p.Response(session, &pb.ErrorResp{
			ErrorCode: 1,
			ErrorStr:  "切换职务失败: 成员不存在",
		})
		return
	}

	member.Job = int(m.Job)
	err = p.dbComponent.UpdateClubMember(member)

	resp := &pb.ClubJobResp{
		//Msg: "切换成功",
	}
	if err != nil {
		p.Response(session, &pb.ErrorResp{
			ErrorCode: 1,
			ErrorStr:  "切换职务失败: " + err.Error(),
		})
		return
	}
	p.Response(session, resp)
}

// 聊天记录
func (p *ActorPlayer) ChatHis(session *cproto.Session, m *pb.ChatHisReq) {

	//chats, err := p.dbComponent.GetChatHistory(m.Channel, m.SenderUid, m.TargetUid, m.ClubId, m.StartTime, m.EndTime)
	//
	resp := &pb.ChatHisResp{
		//DataArr: make([]*pb.ChatInfo, 0),
	}
	//if err != nil {
	//	p.Response(session, &pb.ErrorResp{
	//		ErrorCode: 1,
	//		ErrorStr:  "获取聊天记录失败: " + err.Error(),
	//	})
	//	return
	//}
	//
	//for _, chat := range chats {
	//	resp.DataArr = append(resp.DataArr, &pb.ChatInfo{
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
	p.Response(session, resp)
}

// 处理入群申请
func (p *ActorPlayer) ClubApplyDeal(session *cproto.Session, m *pb.ClubApplyDealReq) {

	apply, err := p.dbComponent.GetClubApplyByID(session.Uid)
	if err != nil {
		p.Response(session, &pb.ErrorResp{
			ErrorCode: 1,
			ErrorStr:  "处理入群申请失败: 申请不存在",
		})
		return
	}

	//apply.Status = int(m.Status)
	err = p.dbComponent.UpdateClubApply(apply)

	resp := &pb.ClubApplyDealResp{
		//Msg: "处理成功",
	}
	if err != nil {
		p.Response(session, &pb.ErrorResp{
			ErrorCode: 1,
			ErrorStr:  "处理入群申请失败: " + err.Error(),
		})
		return
	}
	p.Response(session, resp)
}

// 我的群
func (p *ActorPlayer) ClubMine(session *cproto.Session, m *pb.ClubMineReq) {

	clubs, err := p.dbComponent.GetClubsByUid(session.Uid)

	resp := &pb.ClubMineResp{
		//Uid:     m.Uid,
		DataArr: make([]*pb.ClubInfo, 0),
	}
	if err != nil {
		p.Response(session, &pb.ErrorResp{
			ErrorCode: 1,
			ErrorStr:  "获取我的群失败: " + err.Error(),
		})
		return
	}

	for _, club := range clubs {
		resp.DataArr = append(resp.DataArr, &pb.ClubInfo{
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
	p.Response(session, resp)
}

// 申请列表
func (p *ActorPlayer) ClubApplyList(session *cproto.Session, m *pb.ClubApplyListReq) {

	applies, err := p.dbComponent.GetClubAppliesByClubId(m.ClubId)

	resp := &pb.ClubApplyListResp{
		ClubId: m.ClubId,
		Mans:   make([]*pb.UserInfo, 0),
	}
	if err != nil {
		p.Response(session, &pb.ErrorResp{
			ErrorCode: 1,
			ErrorStr:  "获取申请列表失败: " + err.Error(),
		})
		return
	}

	for _, apply := range applies {
		resp.Mans = append(resp.Mans, &pb.UserInfo{
			UserID: apply.ID,
		})
	}
	p.Response(session, resp)
}
