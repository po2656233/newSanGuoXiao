package home

import (
	"github.com/po2656233/superplace/const/code"
	pb "superman/internal/protocol/gofile"
	"superman/internal/rpc"
)

// CreateRoom 创建房间
func (p *ActorHome) CreateRoom(req *pb.CreateRoomReq) (*pb.CreateRoomResp, int32) {
	resp, errCode := rpc.SendData(p.App(), rpc.SourcePath, ".db", rpc.CenterType, req)
	if resp == nil {
		return nil, errCode
	}
	return resp.(*pb.CreateRoomResp), code.OK
}
