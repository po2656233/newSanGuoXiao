package manger

//
////----------房间----------------------
////房间信息
//type RoomInfo struct {
//	ID        int32   //房间ID 0表示无效
//	Key       string   //房间钥匙
//	State     uint8    //房间状态 [0:无效] [1:Open] [2:Close] [3:Other] [4:Clear]
//	TableSet  []int64 //桌子集合(等同于GameID集合)
//	MaxPeople int32   //最大承载人数
//}
//
////房间接口
//type IRoomMange interface {
//	Create(roomID int32) (*RoomInfo, bool)           //创建房间[房间ID和钥匙配对]
//	Check(roomID int32) (*RoomInfo, bool)            //查找房间
//	Delete(roomID int32, key string) bool            //删除房间[房间ID和钥匙配对成功后,才能删除]
//	Open(roomID int32, key string) (*RoomInfo, bool) //开启房间
//	Close(roomID int32, key string) bool             //关闭房间
//	Clear(roomID int32, key string) bool             //清理房间
//}
//
////管理房间
//type RoomManger struct {
//	sync.Map
//	//rooms map[int64]*RoomInfo
//}
//
/////////////////房间和子游戏管理////////////////////////
//var lock sync.Mutex
//
////由于新增平台，则这里的管理不应该再是单例模式
//func GetRoomManger() *RoomManger {
//	return &RoomManger{
//		sync.Map{},
//	}
//}
//
////新增桌子(tid == gid 即桌子ID和游戏ID共用)
//func (self *RoomInfo) AddTable(game *protoMsg.GameItem) {
//	self.TableSet = CopyInsert(self.TableSet, len(self.TableSet), game.ID).([]int64)
//	self.TableSet = SliceRemoveDuplicate(self.TableSet).([]int64)
//	self.MaxPeople += game.Info.MaxOnline
//}
//
////删除桌子
//func (self *RoomInfo) DelTable(game *protoMsg.GameItem) {
//	self.TableSet = DeleteValue(self.TableSet, game.ID).([]int64)
//	self.MaxPeople -= game.Info.MaxOnline
//}
//
////测试
//func (self *RoomManger) PrintfAll() { //打印当前所房间
//	log.Debug("所有房间号:->")
//	self.Range(func(index, value interface{}) bool {
//		log.Debug("index:%v, 房间号码:%v", index, value)
//		return true
//	})
//}
//
//func (self *RoomManger) Create(roomInfo *protoMsg.RoomInfo) (*RoomInfo, bool) { //新增
//	info := &RoomInfo{}
//	bRet := true
//	if v, ok := self.Load(roomInfo.RoomNum); !ok {
//		log.Debug("创建房间:%v", roomInfo.RoomNum)
//		info.ID = roomInfo.RoomNum
//		info.Key = roomInfo.RoomKey
//		info.TableSet = make([]int64, 0)
//		if nil != roomInfo.Games {
//			for k, v := range roomInfo.Games.Items {
//				info.MaxPeople += v.Info.MaxOnline
//				info.TableSet = CopyInsert(info.TableSet, k, v.ID).([]int64)
//			}
//		}
//		self.Store(info.ID, info)
//	} else {
//		log.Debug("房间:%v 已经存在", roomInfo.RoomNum)
//		info = v.(*RoomInfo)
//		bRet = false
//	}
//	return info, bRet
//}
//
//func (self *RoomManger) DeleteRoom(roomID int32, key string) bool { //删除
//	bRet := false
//	self.Range(func(index, value interface{}) bool {
//		if roomID == index.(int32) && key == value.(*RoomInfo).Key {
//			bRet = true
//			//for _,gameHandle :=range value.(*RoomInfo).Things.GameSet{
//			//}
//
//			self.Delete(index)
//			return false
//		}
//		return true
//	})
//	return bRet
//}
//
//func (self *RoomManger) Check(roomID int32) (*RoomInfo, bool) { //查找
//	if value, ok := self.Load(roomID); ok {
//		return value.(*RoomInfo), ok
//	}
//	return nil, false
//}
//func (self *RoomManger) Open(roomID int32, key string) (*RoomInfo, bool) { //开启
//	lock.Lock()
//	defer lock.Unlock()
//	info := &RoomInfo{}
//	return info, true
//}
//func (self *RoomManger) Close(roomID int32, key string) bool { //关闭
//	lock.Lock()
//	defer lock.Unlock()
//
//	return true
//}
//func (self *RoomManger) Clear(roomID int32, key string) bool { //清理房间
//	lock.Lock()
//	defer lock.Unlock()
//
//	return true
//}
