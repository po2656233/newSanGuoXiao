package home

import (
	"github.com/po2656233/superplace/const/code"
	cactor "github.com/po2656233/superplace/net/actor"
	pb "superman/internal/protocol/gofile"
	"superman/internal/rpc"
)

type (
	ActorHome struct {
		cactor.Base
	}
)

func (p *ActorHome) AliasID() string {
	return "home"
}

// OnInit center为后端节点，不直接与客户端通信，所以了一些remote函数，供RPC调用
func (p *ActorHome) OnInit() {
	p.Remote().Register(p.GetClassList)
	p.Remote().Register(p.GetRoomList)
	p.Remote().Register(p.GetTableList)
	p.Remote().Register(p.GetGameList)
	p.Remote().Register(p.GetTaskList)
	//
	p.Remote().Register(p.CreateRoom)
}

func (p *ActorHome) GetClassList() (*pb.GetClassListResp, int32) {
	//target := rpc.GetTargetPath(p.App(), ".db", rpc.CenterType)
	resp, errCode := rpc.SendData(p.App(), rpc.SourcePath, ".db", rpc.CenterType, &pb.GetClassListReq{})
	//errCode := p.App().ActorSystem().CallWait(rpc.SourcePath, target, "ClassList", &pb.GetClassListReq{}, resp)
	if resp == nil {
		return nil, errCode
	}
	return resp.(*pb.GetClassListResp), code.OK
}

// GetRoomList 获取房间列表
func (p *ActorHome) GetRoomList() (*pb.GetRoomListResp, int32) {
	resp, errCode := rpc.SendData(p.App(), rpc.SourcePath, ".db", rpc.CenterType, &pb.GetRoomListReq{})
	if resp == nil {
		return nil, errCode
	}
	return resp.(*pb.GetRoomListResp), code.OK
}

// GetTableList 获取牌桌列表
func (p *ActorHome) GetTableList() (*pb.GetTableListResp, int32) {
	resp, errCode := rpc.SendData(p.App(), rpc.SourcePath, ".db", rpc.CenterType, &pb.GetTableListReq{})
	if resp == nil {
		return nil, errCode
	}
	return resp.(*pb.GetTableListResp), code.OK
}

// GetGameList 获取游戏列表
func (p *ActorHome) GetGameList() (*pb.GetGameListResp, int32) {
	resp, errCode := rpc.SendData(p.App(), rpc.SourcePath, ".db", rpc.CenterType, &pb.GetGameListReq{})
	if resp == nil {
		return nil, errCode
	}
	return resp.(*pb.GetGameListResp), code.OK
}

// GetTaskList 获取任务列表
func (p *ActorHome) GetTaskList() (*pb.GetTaskListResp, int32) {
	resp, errCode := rpc.SendData(p.App(), rpc.SourcePath, ".db", rpc.CenterType, &pb.GetTaskListReq{})
	if resp == nil {
		return nil, errCode
	}
	return resp.(*pb.GetTaskListResp), code.OK
}
