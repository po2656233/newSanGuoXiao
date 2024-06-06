/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

goog.exportSymbol('proto.pb.AllopatricResp', null, global);
goog.exportSymbol('proto.pb.BarterReq', null, global);
goog.exportSymbol('proto.pb.BarterResp', null, global);
goog.exportSymbol('proto.pb.BuyGoodsReq', null, global);
goog.exportSymbol('proto.pb.BuyGoodsResp', null, global);
goog.exportSymbol('proto.pb.CheckHeroReq', null, global);
goog.exportSymbol('proto.pb.CheckHeroResp', null, global);
goog.exportSymbol('proto.pb.CheckInReq', null, global);
goog.exportSymbol('proto.pb.CheckInResp', null, global);
goog.exportSymbol('proto.pb.CheckKnapsackReq', null, global);
goog.exportSymbol('proto.pb.CheckKnapsackResp', null, global);
goog.exportSymbol('proto.pb.ChooseClassReq', null, global);
goog.exportSymbol('proto.pb.ChooseClassResp', null, global);
goog.exportSymbol('proto.pb.ChooseGameReq', null, global);
goog.exportSymbol('proto.pb.ChooseGameResp', null, global);
goog.exportSymbol('proto.pb.ChooseHeroReq', null, global);
goog.exportSymbol('proto.pb.ChooseHeroResp', null, global);
goog.exportSymbol('proto.pb.ClaimReq', null, global);
goog.exportSymbol('proto.pb.ClaimResp', null, global);
goog.exportSymbol('proto.pb.ClassItem', null, global);
goog.exportSymbol('proto.pb.ClassList', null, global);
goog.exportSymbol('proto.pb.DownHeroReq', null, global);
goog.exportSymbol('proto.pb.DownHeroReqResp', null, global);
goog.exportSymbol('proto.pb.DrawHeroReq', null, global);
goog.exportSymbol('proto.pb.DrawHeroResp', null, global);
goog.exportSymbol('proto.pb.EmailDelReq', null, global);
goog.exportSymbol('proto.pb.EmailDelResp', null, global);
goog.exportSymbol('proto.pb.EmailInfo', null, global);
goog.exportSymbol('proto.pb.EmailReadReq', null, global);
goog.exportSymbol('proto.pb.EmailReadResp', null, global);
goog.exportSymbol('proto.pb.EmailReq', null, global);
goog.exportSymbol('proto.pb.EmailResp', null, global);
goog.exportSymbol('proto.pb.GameInfo', null, global);
goog.exportSymbol('proto.pb.GameItem', null, global);
goog.exportSymbol('proto.pb.GameList', null, global);
goog.exportSymbol('proto.pb.GameScene', null, global);
goog.exportSymbol('proto.pb.GameType', null, global);
goog.exportSymbol('proto.pb.GetAllGoodsReq', null, global);
goog.exportSymbol('proto.pb.GetAllGoodsResp', null, global);
goog.exportSymbol('proto.pb.GetAllHeroReq', null, global);
goog.exportSymbol('proto.pb.GetAllHeroResp', null, global);
goog.exportSymbol('proto.pb.GetCheckInReq', null, global);
goog.exportSymbol('proto.pb.GetCheckInResp', null, global);
goog.exportSymbol('proto.pb.GetGoodsReq', null, global);
goog.exportSymbol('proto.pb.GetGoodsResp', null, global);
goog.exportSymbol('proto.pb.GetMyHeroReq', null, global);
goog.exportSymbol('proto.pb.GetMyHeroResp', null, global);
goog.exportSymbol('proto.pb.GetRechargesReq', null, global);
goog.exportSymbol('proto.pb.GetRechargesResp', null, global);
goog.exportSymbol('proto.pb.GoodsInfo', null, global);
goog.exportSymbol('proto.pb.GoodsList', null, global);
goog.exportSymbol('proto.pb.HeroInfo', null, global);
goog.exportSymbol('proto.pb.HeroType', null, global);
goog.exportSymbol('proto.pb.KnapsackInfo', null, global);
goog.exportSymbol('proto.pb.LoginReq', null, global);
goog.exportSymbol('proto.pb.LoginRequest', null, global);
goog.exportSymbol('proto.pb.LoginResp', null, global);
goog.exportSymbol('proto.pb.LoginResponse', null, global);
goog.exportSymbol('proto.pb.MasterInfo', null, global);
goog.exportSymbol('proto.pb.PingReq', null, global);
goog.exportSymbol('proto.pb.PongResp', null, global);
goog.exportSymbol('proto.pb.RechargeReq', null, global);
goog.exportSymbol('proto.pb.RechargeResp', null, global);
goog.exportSymbol('proto.pb.ReconnectReq', null, global);
goog.exportSymbol('proto.pb.ReconnectResp', null, global);
goog.exportSymbol('proto.pb.RegisterReq', null, global);
goog.exportSymbol('proto.pb.RegisterResp', null, global);
goog.exportSymbol('proto.pb.ResultPopResp', null, global);
goog.exportSymbol('proto.pb.ResultResp', null, global);
goog.exportSymbol('proto.pb.SettingTableReq', null, global);
goog.exportSymbol('proto.pb.SettingTableResp', null, global);
goog.exportSymbol('proto.pb.SuggestReq', null, global);
goog.exportSymbol('proto.pb.SuggestResp', null, global);
goog.exportSymbol('proto.pb.TableInfo', null, global);
goog.exportSymbol('proto.pb.TableItem', null, global);
goog.exportSymbol('proto.pb.TableList', null, global);
goog.exportSymbol('proto.pb.TableState', null, global);
goog.exportSymbol('proto.pb.TaskItem', null, global);
goog.exportSymbol('proto.pb.TaskList', null, global);
goog.exportSymbol('proto.pb.ToShoppingResp', null, global);
goog.exportSymbol('proto.pb.UserInfo', null, global);
goog.exportSymbol('proto.pb.WeaponInfo', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.LoginRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.LoginRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.LoginRequest.displayName = 'proto.pb.LoginRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.LoginResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.LoginResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.LoginResponse.displayName = 'proto.pb.LoginResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.UserInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.UserInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.UserInfo.displayName = 'proto.pb.UserInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.HeroInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.pb.HeroInfo.repeatedFields_, null);
};
goog.inherits(proto.pb.HeroInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.HeroInfo.displayName = 'proto.pb.HeroInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.WeaponInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.WeaponInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.WeaponInfo.displayName = 'proto.pb.WeaponInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.GoodsInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.GoodsInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.GoodsInfo.displayName = 'proto.pb.GoodsInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.GoodsList = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.pb.GoodsList.repeatedFields_, null);
};
goog.inherits(proto.pb.GoodsList, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.GoodsList.displayName = 'proto.pb.GoodsList';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.KnapsackInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.pb.KnapsackInfo.repeatedFields_, null);
};
goog.inherits(proto.pb.KnapsackInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.KnapsackInfo.displayName = 'proto.pb.KnapsackInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.EmailInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.EmailInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.EmailInfo.displayName = 'proto.pb.EmailInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.TableInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.TableInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.TableInfo.displayName = 'proto.pb.TableInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.GameInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.GameInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.GameInfo.displayName = 'proto.pb.GameInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.TaskItem = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.TaskItem, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.TaskItem.displayName = 'proto.pb.TaskItem';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.ClassItem = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.ClassItem, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.ClassItem.displayName = 'proto.pb.ClassItem';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.GameItem = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.GameItem, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.GameItem.displayName = 'proto.pb.GameItem';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.TableItem = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.TableItem, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.TableItem.displayName = 'proto.pb.TableItem';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.TaskList = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.pb.TaskList.repeatedFields_, null);
};
goog.inherits(proto.pb.TaskList, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.TaskList.displayName = 'proto.pb.TaskList';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.ClassList = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.pb.ClassList.repeatedFields_, null);
};
goog.inherits(proto.pb.ClassList, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.ClassList.displayName = 'proto.pb.ClassList';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.GameList = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.pb.GameList.repeatedFields_, null);
};
goog.inherits(proto.pb.GameList, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.GameList.displayName = 'proto.pb.GameList';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.TableList = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.pb.TableList.repeatedFields_, null);
};
goog.inherits(proto.pb.TableList, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.TableList.displayName = 'proto.pb.TableList';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.MasterInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.MasterInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.MasterInfo.displayName = 'proto.pb.MasterInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.RegisterReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.RegisterReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.RegisterReq.displayName = 'proto.pb.RegisterReq';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.RegisterResp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.RegisterResp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.RegisterResp.displayName = 'proto.pb.RegisterResp';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.LoginReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.LoginReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.LoginReq.displayName = 'proto.pb.LoginReq';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.LoginResp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.LoginResp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.LoginResp.displayName = 'proto.pb.LoginResp';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.AllopatricResp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.AllopatricResp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.AllopatricResp.displayName = 'proto.pb.AllopatricResp';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.ReconnectReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.ReconnectReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.ReconnectReq.displayName = 'proto.pb.ReconnectReq';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.ReconnectResp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.ReconnectResp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.ReconnectResp.displayName = 'proto.pb.ReconnectResp';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.ChooseClassReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.ChooseClassReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.ChooseClassReq.displayName = 'proto.pb.ChooseClassReq';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.ChooseClassResp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.ChooseClassResp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.ChooseClassResp.displayName = 'proto.pb.ChooseClassResp';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.ChooseGameReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.ChooseGameReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.ChooseGameReq.displayName = 'proto.pb.ChooseGameReq';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.ChooseGameResp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.ChooseGameResp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.ChooseGameResp.displayName = 'proto.pb.ChooseGameResp';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.SettingTableReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.SettingTableReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.SettingTableReq.displayName = 'proto.pb.SettingTableReq';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.SettingTableResp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.SettingTableResp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.SettingTableResp.displayName = 'proto.pb.SettingTableResp';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.CheckInReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.CheckInReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.CheckInReq.displayName = 'proto.pb.CheckInReq';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.CheckInResp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.CheckInResp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.CheckInResp.displayName = 'proto.pb.CheckInResp';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.GetCheckInReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.GetCheckInReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.GetCheckInReq.displayName = 'proto.pb.GetCheckInReq';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.GetCheckInResp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.pb.GetCheckInResp.repeatedFields_, null);
};
goog.inherits(proto.pb.GetCheckInResp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.GetCheckInResp.displayName = 'proto.pb.GetCheckInResp';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.DrawHeroReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.DrawHeroReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.DrawHeroReq.displayName = 'proto.pb.DrawHeroReq';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.DrawHeroResp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.pb.DrawHeroResp.repeatedFields_, null);
};
goog.inherits(proto.pb.DrawHeroResp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.DrawHeroResp.displayName = 'proto.pb.DrawHeroResp';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.GetMyHeroReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.GetMyHeroReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.GetMyHeroReq.displayName = 'proto.pb.GetMyHeroReq';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.GetMyHeroResp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.pb.GetMyHeroResp.repeatedFields_, null);
};
goog.inherits(proto.pb.GetMyHeroResp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.GetMyHeroResp.displayName = 'proto.pb.GetMyHeroResp';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.ChooseHeroReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.ChooseHeroReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.ChooseHeroReq.displayName = 'proto.pb.ChooseHeroReq';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.ChooseHeroResp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.ChooseHeroResp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.ChooseHeroResp.displayName = 'proto.pb.ChooseHeroResp';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.DownHeroReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.DownHeroReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.DownHeroReq.displayName = 'proto.pb.DownHeroReq';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.DownHeroReqResp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.DownHeroReqResp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.DownHeroReqResp.displayName = 'proto.pb.DownHeroReqResp';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.GetAllHeroReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.GetAllHeroReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.GetAllHeroReq.displayName = 'proto.pb.GetAllHeroReq';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.GetAllHeroResp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.pb.GetAllHeroResp.repeatedFields_, null);
};
goog.inherits(proto.pb.GetAllHeroResp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.GetAllHeroResp.displayName = 'proto.pb.GetAllHeroResp';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.CheckHeroReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.pb.CheckHeroReq.repeatedFields_, null);
};
goog.inherits(proto.pb.CheckHeroReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.CheckHeroReq.displayName = 'proto.pb.CheckHeroReq';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.CheckHeroResp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.pb.CheckHeroResp.repeatedFields_, null);
};
goog.inherits(proto.pb.CheckHeroResp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.CheckHeroResp.displayName = 'proto.pb.CheckHeroResp';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.RechargeReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.RechargeReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.RechargeReq.displayName = 'proto.pb.RechargeReq';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.RechargeResp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.RechargeResp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.RechargeResp.displayName = 'proto.pb.RechargeResp';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.GetRechargesReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.GetRechargesReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.GetRechargesReq.displayName = 'proto.pb.GetRechargesReq';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.GetRechargesResp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.pb.GetRechargesResp.repeatedFields_, null);
};
goog.inherits(proto.pb.GetRechargesResp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.GetRechargesResp.displayName = 'proto.pb.GetRechargesResp';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.GetGoodsReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.GetGoodsReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.GetGoodsReq.displayName = 'proto.pb.GetGoodsReq';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.GetGoodsResp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.GetGoodsResp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.GetGoodsResp.displayName = 'proto.pb.GetGoodsResp';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.GetAllGoodsReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.GetAllGoodsReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.GetAllGoodsReq.displayName = 'proto.pb.GetAllGoodsReq';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.GetAllGoodsResp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.pb.GetAllGoodsResp.repeatedFields_, null);
};
goog.inherits(proto.pb.GetAllGoodsResp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.GetAllGoodsResp.displayName = 'proto.pb.GetAllGoodsResp';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.BuyGoodsReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.BuyGoodsReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.BuyGoodsReq.displayName = 'proto.pb.BuyGoodsReq';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.BuyGoodsResp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.BuyGoodsResp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.BuyGoodsResp.displayName = 'proto.pb.BuyGoodsResp';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.CheckKnapsackReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.CheckKnapsackReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.CheckKnapsackReq.displayName = 'proto.pb.CheckKnapsackReq';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.CheckKnapsackResp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.CheckKnapsackResp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.CheckKnapsackResp.displayName = 'proto.pb.CheckKnapsackResp';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.BarterReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.BarterReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.BarterReq.displayName = 'proto.pb.BarterReq';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.BarterResp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.BarterResp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.BarterResp.displayName = 'proto.pb.BarterResp';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.ToShoppingResp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.ToShoppingResp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.ToShoppingResp.displayName = 'proto.pb.ToShoppingResp';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.EmailReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.EmailReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.EmailReq.displayName = 'proto.pb.EmailReq';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.EmailResp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.pb.EmailResp.repeatedFields_, null);
};
goog.inherits(proto.pb.EmailResp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.EmailResp.displayName = 'proto.pb.EmailResp';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.ClaimReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.ClaimReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.ClaimReq.displayName = 'proto.pb.ClaimReq';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.ClaimResp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.ClaimResp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.ClaimResp.displayName = 'proto.pb.ClaimResp';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.SuggestReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.SuggestReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.SuggestReq.displayName = 'proto.pb.SuggestReq';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.SuggestResp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.SuggestResp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.SuggestResp.displayName = 'proto.pb.SuggestResp';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.EmailReadReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.EmailReadReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.EmailReadReq.displayName = 'proto.pb.EmailReadReq';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.EmailReadResp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.EmailReadResp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.EmailReadResp.displayName = 'proto.pb.EmailReadResp';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.EmailDelReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.EmailDelReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.EmailDelReq.displayName = 'proto.pb.EmailDelReq';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.EmailDelResp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.EmailDelResp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.EmailDelResp.displayName = 'proto.pb.EmailDelResp';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.ResultResp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.ResultResp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.ResultResp.displayName = 'proto.pb.ResultResp';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.ResultPopResp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.ResultPopResp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.ResultPopResp.displayName = 'proto.pb.ResultPopResp';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.PingReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.PingReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.PingReq.displayName = 'proto.pb.PingReq';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.PongResp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.PongResp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.PongResp.displayName = 'proto.pb.PongResp';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.LoginRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.LoginRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.LoginRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.LoginRequest.toObject = function(includeInstance, msg) {
  var obj = {
    serverid: jspb.Message.getFieldWithDefault(msg, 1, 0),
    token: jspb.Message.getFieldWithDefault(msg, 2, ""),
    paramsMap: (f = msg.getParamsMap()) ? f.toObject(includeInstance, undefined) : []
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.LoginRequest}
 */
proto.pb.LoginRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.LoginRequest;
  return proto.pb.LoginRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.LoginRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.LoginRequest}
 */
proto.pb.LoginRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setServerid(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setToken(value);
      break;
    case 3:
      var value = msg.getParamsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readInt32, jspb.BinaryReader.prototype.readString, null, 0);
         });
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.LoginRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.LoginRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.LoginRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.LoginRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getServerid();
  if (f !== 0) {
    writer.writeInt32(
      1,
      f
    );
  }
  f = message.getToken();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getParamsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(3, writer, jspb.BinaryWriter.prototype.writeInt32, jspb.BinaryWriter.prototype.writeString);
  }
};


/**
 * optional int32 serverId = 1;
 * @return {number}
 */
proto.pb.LoginRequest.prototype.getServerid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.LoginRequest.prototype.setServerid = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional string token = 2;
 * @return {string}
 */
proto.pb.LoginRequest.prototype.getToken = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.pb.LoginRequest.prototype.setToken = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * map<int32, string> params = 3;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<number,string>}
 */
proto.pb.LoginRequest.prototype.getParamsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<number,string>} */ (
      jspb.Message.getMapField(this, 3, opt_noLazyCreate,
      null));
};


/**
 * Clears values from the map. The map will be non-null.
 */
proto.pb.LoginRequest.prototype.clearParamsMap = function() {
  this.getParamsMap().clear();
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.LoginResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.LoginResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.LoginResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.LoginResponse.toObject = function(includeInstance, msg) {
  var obj = {
    uid: jspb.Message.getFieldWithDefault(msg, 1, 0),
    pid: jspb.Message.getFieldWithDefault(msg, 2, 0),
    openid: jspb.Message.getFieldWithDefault(msg, 3, ""),
    paramsMap: (f = msg.getParamsMap()) ? f.toObject(includeInstance, undefined) : []
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.LoginResponse}
 */
proto.pb.LoginResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.LoginResponse;
  return proto.pb.LoginResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.LoginResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.LoginResponse}
 */
proto.pb.LoginResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setUid(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setPid(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setOpenid(value);
      break;
    case 4:
      var value = msg.getParamsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readInt32, jspb.BinaryReader.prototype.readString, null, 0);
         });
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.LoginResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.LoginResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.LoginResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.LoginResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUid();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getPid();
  if (f !== 0) {
    writer.writeInt32(
      2,
      f
    );
  }
  f = message.getOpenid();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getParamsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(4, writer, jspb.BinaryWriter.prototype.writeInt32, jspb.BinaryWriter.prototype.writeString);
  }
};


/**
 * optional int64 uid = 1;
 * @return {number}
 */
proto.pb.LoginResponse.prototype.getUid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.LoginResponse.prototype.setUid = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional int32 pid = 2;
 * @return {number}
 */
proto.pb.LoginResponse.prototype.getPid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {number} value */
proto.pb.LoginResponse.prototype.setPid = function(value) {
  jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional string openId = 3;
 * @return {string}
 */
proto.pb.LoginResponse.prototype.getOpenid = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.pb.LoginResponse.prototype.setOpenid = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * map<int32, string> params = 4;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<number,string>}
 */
proto.pb.LoginResponse.prototype.getParamsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<number,string>} */ (
      jspb.Message.getMapField(this, 4, opt_noLazyCreate,
      null));
};


/**
 * Clears values from the map. The map will be non-null.
 */
proto.pb.LoginResponse.prototype.clearParamsMap = function() {
  this.getParamsMap().clear();
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.UserInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.UserInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.UserInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.UserInfo.toObject = function(includeInstance, msg) {
  var obj = {
    userid: jspb.Message.getFieldWithDefault(msg, 1, 0),
    name: jspb.Message.getFieldWithDefault(msg, 2, ""),
    account: jspb.Message.getFieldWithDefault(msg, 3, ""),
    password: jspb.Message.getFieldWithDefault(msg, 4, ""),
    faceid: jspb.Message.getFieldWithDefault(msg, 5, 0),
    gender: jspb.Message.getFieldWithDefault(msg, 6, 0),
    age: jspb.Message.getFieldWithDefault(msg, 7, 0),
    vip: jspb.Message.getFieldWithDefault(msg, 8, 0),
    level: jspb.Message.getFieldWithDefault(msg, 9, 0),
    yuanbao: jspb.Message.getFieldWithDefault(msg, 10, 0),
    coin: jspb.Message.getFieldWithDefault(msg, 11, 0),
    money: jspb.Message.getFieldWithDefault(msg, 12, 0),
    passportid: jspb.Message.getFieldWithDefault(msg, 13, ""),
    realname: jspb.Message.getFieldWithDefault(msg, 14, ""),
    phonenum: jspb.Message.getFieldWithDefault(msg, 15, ""),
    email: jspb.Message.getFieldWithDefault(msg, 16, ""),
    address: jspb.Message.getFieldWithDefault(msg, 17, ""),
    identity: jspb.Message.getFieldWithDefault(msg, 18, ""),
    agentid: jspb.Message.getFieldWithDefault(msg, 19, 0),
    referralcode: jspb.Message.getFieldWithDefault(msg, 20, ""),
    clientaddr: jspb.Message.getFieldWithDefault(msg, 21, ""),
    serveraddr: jspb.Message.getFieldWithDefault(msg, 22, ""),
    machinecode: jspb.Message.getFieldWithDefault(msg, 23, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.UserInfo}
 */
proto.pb.UserInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.UserInfo;
  return proto.pb.UserInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.UserInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.UserInfo}
 */
proto.pb.UserInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setUserid(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setAccount(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setPassword(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setFaceid(value);
      break;
    case 6:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setGender(value);
      break;
    case 7:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setAge(value);
      break;
    case 8:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setVip(value);
      break;
    case 9:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setLevel(value);
      break;
    case 10:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setYuanbao(value);
      break;
    case 11:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setCoin(value);
      break;
    case 12:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setMoney(value);
      break;
    case 13:
      var value = /** @type {string} */ (reader.readString());
      msg.setPassportid(value);
      break;
    case 14:
      var value = /** @type {string} */ (reader.readString());
      msg.setRealname(value);
      break;
    case 15:
      var value = /** @type {string} */ (reader.readString());
      msg.setPhonenum(value);
      break;
    case 16:
      var value = /** @type {string} */ (reader.readString());
      msg.setEmail(value);
      break;
    case 17:
      var value = /** @type {string} */ (reader.readString());
      msg.setAddress(value);
      break;
    case 18:
      var value = /** @type {string} */ (reader.readString());
      msg.setIdentity(value);
      break;
    case 19:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setAgentid(value);
      break;
    case 20:
      var value = /** @type {string} */ (reader.readString());
      msg.setReferralcode(value);
      break;
    case 21:
      var value = /** @type {string} */ (reader.readString());
      msg.setClientaddr(value);
      break;
    case 22:
      var value = /** @type {string} */ (reader.readString());
      msg.setServeraddr(value);
      break;
    case 23:
      var value = /** @type {string} */ (reader.readString());
      msg.setMachinecode(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.UserInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.UserInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.UserInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.UserInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUserid();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getAccount();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getPassword();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getFaceid();
  if (f !== 0) {
    writer.writeInt32(
      5,
      f
    );
  }
  f = message.getGender();
  if (f !== 0) {
    writer.writeInt32(
      6,
      f
    );
  }
  f = message.getAge();
  if (f !== 0) {
    writer.writeInt32(
      7,
      f
    );
  }
  f = message.getVip();
  if (f !== 0) {
    writer.writeInt32(
      8,
      f
    );
  }
  f = message.getLevel();
  if (f !== 0) {
    writer.writeInt32(
      9,
      f
    );
  }
  f = message.getYuanbao();
  if (f !== 0) {
    writer.writeInt64(
      10,
      f
    );
  }
  f = message.getCoin();
  if (f !== 0) {
    writer.writeInt64(
      11,
      f
    );
  }
  f = message.getMoney();
  if (f !== 0) {
    writer.writeInt64(
      12,
      f
    );
  }
  f = message.getPassportid();
  if (f.length > 0) {
    writer.writeString(
      13,
      f
    );
  }
  f = message.getRealname();
  if (f.length > 0) {
    writer.writeString(
      14,
      f
    );
  }
  f = message.getPhonenum();
  if (f.length > 0) {
    writer.writeString(
      15,
      f
    );
  }
  f = message.getEmail();
  if (f.length > 0) {
    writer.writeString(
      16,
      f
    );
  }
  f = message.getAddress();
  if (f.length > 0) {
    writer.writeString(
      17,
      f
    );
  }
  f = message.getIdentity();
  if (f.length > 0) {
    writer.writeString(
      18,
      f
    );
  }
  f = message.getAgentid();
  if (f !== 0) {
    writer.writeInt64(
      19,
      f
    );
  }
  f = message.getReferralcode();
  if (f.length > 0) {
    writer.writeString(
      20,
      f
    );
  }
  f = message.getClientaddr();
  if (f.length > 0) {
    writer.writeString(
      21,
      f
    );
  }
  f = message.getServeraddr();
  if (f.length > 0) {
    writer.writeString(
      22,
      f
    );
  }
  f = message.getMachinecode();
  if (f.length > 0) {
    writer.writeString(
      23,
      f
    );
  }
};


/**
 * optional int64 userID = 1;
 * @return {number}
 */
proto.pb.UserInfo.prototype.getUserid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.UserInfo.prototype.setUserid = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional string name = 2;
 * @return {string}
 */
proto.pb.UserInfo.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.pb.UserInfo.prototype.setName = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string account = 3;
 * @return {string}
 */
proto.pb.UserInfo.prototype.getAccount = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.pb.UserInfo.prototype.setAccount = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string password = 4;
 * @return {string}
 */
proto.pb.UserInfo.prototype.getPassword = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.pb.UserInfo.prototype.setPassword = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional int32 faceID = 5;
 * @return {number}
 */
proto.pb.UserInfo.prototype.getFaceid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/** @param {number} value */
proto.pb.UserInfo.prototype.setFaceid = function(value) {
  jspb.Message.setProto3IntField(this, 5, value);
};


/**
 * optional int32 gender = 6;
 * @return {number}
 */
proto.pb.UserInfo.prototype.getGender = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 6, 0));
};


/** @param {number} value */
proto.pb.UserInfo.prototype.setGender = function(value) {
  jspb.Message.setProto3IntField(this, 6, value);
};


/**
 * optional int32 age = 7;
 * @return {number}
 */
proto.pb.UserInfo.prototype.getAge = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 7, 0));
};


/** @param {number} value */
proto.pb.UserInfo.prototype.setAge = function(value) {
  jspb.Message.setProto3IntField(this, 7, value);
};


/**
 * optional int32 vIP = 8;
 * @return {number}
 */
proto.pb.UserInfo.prototype.getVip = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 8, 0));
};


/** @param {number} value */
proto.pb.UserInfo.prototype.setVip = function(value) {
  jspb.Message.setProto3IntField(this, 8, value);
};


/**
 * optional int32 level = 9;
 * @return {number}
 */
proto.pb.UserInfo.prototype.getLevel = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 9, 0));
};


/** @param {number} value */
proto.pb.UserInfo.prototype.setLevel = function(value) {
  jspb.Message.setProto3IntField(this, 9, value);
};


/**
 * optional int64 yuanBao = 10;
 * @return {number}
 */
proto.pb.UserInfo.prototype.getYuanbao = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 10, 0));
};


/** @param {number} value */
proto.pb.UserInfo.prototype.setYuanbao = function(value) {
  jspb.Message.setProto3IntField(this, 10, value);
};


/**
 * optional int64 coin = 11;
 * @return {number}
 */
proto.pb.UserInfo.prototype.getCoin = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 11, 0));
};


/** @param {number} value */
proto.pb.UserInfo.prototype.setCoin = function(value) {
  jspb.Message.setProto3IntField(this, 11, value);
};


/**
 * optional int64 money = 12;
 * @return {number}
 */
proto.pb.UserInfo.prototype.getMoney = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 12, 0));
};


/** @param {number} value */
proto.pb.UserInfo.prototype.setMoney = function(value) {
  jspb.Message.setProto3IntField(this, 12, value);
};


/**
 * optional string passPortID = 13;
 * @return {string}
 */
proto.pb.UserInfo.prototype.getPassportid = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 13, ""));
};


/** @param {string} value */
proto.pb.UserInfo.prototype.setPassportid = function(value) {
  jspb.Message.setProto3StringField(this, 13, value);
};


/**
 * optional string realName = 14;
 * @return {string}
 */
proto.pb.UserInfo.prototype.getRealname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 14, ""));
};


/** @param {string} value */
proto.pb.UserInfo.prototype.setRealname = function(value) {
  jspb.Message.setProto3StringField(this, 14, value);
};


/**
 * optional string phoneNum = 15;
 * @return {string}
 */
proto.pb.UserInfo.prototype.getPhonenum = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 15, ""));
};


/** @param {string} value */
proto.pb.UserInfo.prototype.setPhonenum = function(value) {
  jspb.Message.setProto3StringField(this, 15, value);
};


/**
 * optional string email = 16;
 * @return {string}
 */
proto.pb.UserInfo.prototype.getEmail = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 16, ""));
};


/** @param {string} value */
proto.pb.UserInfo.prototype.setEmail = function(value) {
  jspb.Message.setProto3StringField(this, 16, value);
};


/**
 * optional string address = 17;
 * @return {string}
 */
proto.pb.UserInfo.prototype.getAddress = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 17, ""));
};


/** @param {string} value */
proto.pb.UserInfo.prototype.setAddress = function(value) {
  jspb.Message.setProto3StringField(this, 17, value);
};


/**
 * optional string iDentity = 18;
 * @return {string}
 */
proto.pb.UserInfo.prototype.getIdentity = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 18, ""));
};


/** @param {string} value */
proto.pb.UserInfo.prototype.setIdentity = function(value) {
  jspb.Message.setProto3StringField(this, 18, value);
};


/**
 * optional int64 agentID = 19;
 * @return {number}
 */
proto.pb.UserInfo.prototype.getAgentid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 19, 0));
};


/** @param {number} value */
proto.pb.UserInfo.prototype.setAgentid = function(value) {
  jspb.Message.setProto3IntField(this, 19, value);
};


/**
 * optional string referralCode = 20;
 * @return {string}
 */
proto.pb.UserInfo.prototype.getReferralcode = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 20, ""));
};


/** @param {string} value */
proto.pb.UserInfo.prototype.setReferralcode = function(value) {
  jspb.Message.setProto3StringField(this, 20, value);
};


/**
 * optional string clientAddr = 21;
 * @return {string}
 */
proto.pb.UserInfo.prototype.getClientaddr = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 21, ""));
};


/** @param {string} value */
proto.pb.UserInfo.prototype.setClientaddr = function(value) {
  jspb.Message.setProto3StringField(this, 21, value);
};


/**
 * optional string serverAddr = 22;
 * @return {string}
 */
proto.pb.UserInfo.prototype.getServeraddr = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 22, ""));
};


/** @param {string} value */
proto.pb.UserInfo.prototype.setServeraddr = function(value) {
  jspb.Message.setProto3StringField(this, 22, value);
};


/**
 * optional string machineCode = 23;
 * @return {string}
 */
proto.pb.UserInfo.prototype.getMachinecode = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 23, ""));
};


/** @param {string} value */
proto.pb.UserInfo.prototype.setMachinecode = function(value) {
  jspb.Message.setProto3StringField(this, 23, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.pb.HeroInfo.repeatedFields_ = [15];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.HeroInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.HeroInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.HeroInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.HeroInfo.toObject = function(includeInstance, msg) {
  var obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, 0),
    headid: jspb.Message.getFieldWithDefault(msg, 2, 0),
    name: jspb.Message.getFieldWithDefault(msg, 3, ""),
    sex: jspb.Message.getFieldWithDefault(msg, 4, 0),
    rarity: jspb.Message.getFieldWithDefault(msg, 5, 0),
    faction: jspb.Message.getFieldWithDefault(msg, 6, 0),
    healthpoint: jspb.Message.getFieldWithDefault(msg, 7, 0),
    healthpointfull: jspb.Message.getFieldWithDefault(msg, 8, 0),
    strength: jspb.Message.getFieldWithDefault(msg, 9, 0),
    agility: jspb.Message.getFieldWithDefault(msg, 10, 0),
    intelligence: jspb.Message.getFieldWithDefault(msg, 11, 0),
    attackpoint: jspb.Message.getFieldWithDefault(msg, 12, 0),
    armorpoint: jspb.Message.getFieldWithDefault(msg, 13, 0),
    spellpower: jspb.Message.getFieldWithDefault(msg, 14, 0),
    skillsList: jspb.Message.getRepeatedField(msg, 15)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.HeroInfo}
 */
proto.pb.HeroInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.HeroInfo;
  return proto.pb.HeroInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.HeroInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.HeroInfo}
 */
proto.pb.HeroInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setHeadid(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setSex(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setRarity(value);
      break;
    case 6:
      var value = /** @type {!proto.pb.HeroType} */ (reader.readEnum());
      msg.setFaction(value);
      break;
    case 7:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setHealthpoint(value);
      break;
    case 8:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setHealthpointfull(value);
      break;
    case 9:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setStrength(value);
      break;
    case 10:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setAgility(value);
      break;
    case 11:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setIntelligence(value);
      break;
    case 12:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setAttackpoint(value);
      break;
    case 13:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setArmorpoint(value);
      break;
    case 14:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setSpellpower(value);
      break;
    case 15:
      var value = /** @type {!Array<number>} */ (reader.readPackedInt64());
      msg.setSkillsList(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.HeroInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.HeroInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.HeroInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.HeroInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getHeadid();
  if (f !== 0) {
    writer.writeInt32(
      2,
      f
    );
  }
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getSex();
  if (f !== 0) {
    writer.writeInt32(
      4,
      f
    );
  }
  f = message.getRarity();
  if (f !== 0) {
    writer.writeInt32(
      5,
      f
    );
  }
  f = message.getFaction();
  if (f !== 0.0) {
    writer.writeEnum(
      6,
      f
    );
  }
  f = message.getHealthpoint();
  if (f !== 0) {
    writer.writeInt64(
      7,
      f
    );
  }
  f = message.getHealthpointfull();
  if (f !== 0) {
    writer.writeInt64(
      8,
      f
    );
  }
  f = message.getStrength();
  if (f !== 0) {
    writer.writeInt64(
      9,
      f
    );
  }
  f = message.getAgility();
  if (f !== 0) {
    writer.writeInt64(
      10,
      f
    );
  }
  f = message.getIntelligence();
  if (f !== 0) {
    writer.writeInt64(
      11,
      f
    );
  }
  f = message.getAttackpoint();
  if (f !== 0) {
    writer.writeInt64(
      12,
      f
    );
  }
  f = message.getArmorpoint();
  if (f !== 0) {
    writer.writeInt64(
      13,
      f
    );
  }
  f = message.getSpellpower();
  if (f !== 0) {
    writer.writeInt64(
      14,
      f
    );
  }
  f = message.getSkillsList();
  if (f.length > 0) {
    writer.writePackedInt64(
      15,
      f
    );
  }
};


/**
 * optional int64 iD = 1;
 * @return {number}
 */
proto.pb.HeroInfo.prototype.getId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.HeroInfo.prototype.setId = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional int32 headId = 2;
 * @return {number}
 */
proto.pb.HeroInfo.prototype.getHeadid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {number} value */
proto.pb.HeroInfo.prototype.setHeadid = function(value) {
  jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional string name = 3;
 * @return {string}
 */
proto.pb.HeroInfo.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.pb.HeroInfo.prototype.setName = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional int32 sex = 4;
 * @return {number}
 */
proto.pb.HeroInfo.prototype.getSex = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/** @param {number} value */
proto.pb.HeroInfo.prototype.setSex = function(value) {
  jspb.Message.setProto3IntField(this, 4, value);
};


/**
 * optional int32 rarity = 5;
 * @return {number}
 */
proto.pb.HeroInfo.prototype.getRarity = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/** @param {number} value */
proto.pb.HeroInfo.prototype.setRarity = function(value) {
  jspb.Message.setProto3IntField(this, 5, value);
};


/**
 * optional HeroType faction = 6;
 * @return {!proto.pb.HeroType}
 */
proto.pb.HeroInfo.prototype.getFaction = function() {
  return /** @type {!proto.pb.HeroType} */ (jspb.Message.getFieldWithDefault(this, 6, 0));
};


/** @param {!proto.pb.HeroType} value */
proto.pb.HeroInfo.prototype.setFaction = function(value) {
  jspb.Message.setProto3EnumField(this, 6, value);
};


/**
 * optional int64 healthPoint = 7;
 * @return {number}
 */
proto.pb.HeroInfo.prototype.getHealthpoint = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 7, 0));
};


/** @param {number} value */
proto.pb.HeroInfo.prototype.setHealthpoint = function(value) {
  jspb.Message.setProto3IntField(this, 7, value);
};


/**
 * optional int64 healthPointFull = 8;
 * @return {number}
 */
proto.pb.HeroInfo.prototype.getHealthpointfull = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 8, 0));
};


/** @param {number} value */
proto.pb.HeroInfo.prototype.setHealthpointfull = function(value) {
  jspb.Message.setProto3IntField(this, 8, value);
};


/**
 * optional int64 strength = 9;
 * @return {number}
 */
proto.pb.HeroInfo.prototype.getStrength = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 9, 0));
};


/** @param {number} value */
proto.pb.HeroInfo.prototype.setStrength = function(value) {
  jspb.Message.setProto3IntField(this, 9, value);
};


/**
 * optional int64 agility = 10;
 * @return {number}
 */
proto.pb.HeroInfo.prototype.getAgility = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 10, 0));
};


/** @param {number} value */
proto.pb.HeroInfo.prototype.setAgility = function(value) {
  jspb.Message.setProto3IntField(this, 10, value);
};


/**
 * optional int64 intelligence = 11;
 * @return {number}
 */
proto.pb.HeroInfo.prototype.getIntelligence = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 11, 0));
};


/** @param {number} value */
proto.pb.HeroInfo.prototype.setIntelligence = function(value) {
  jspb.Message.setProto3IntField(this, 11, value);
};


/**
 * optional int64 attackPoint = 12;
 * @return {number}
 */
proto.pb.HeroInfo.prototype.getAttackpoint = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 12, 0));
};


/** @param {number} value */
proto.pb.HeroInfo.prototype.setAttackpoint = function(value) {
  jspb.Message.setProto3IntField(this, 12, value);
};


/**
 * optional int64 armorPoint = 13;
 * @return {number}
 */
proto.pb.HeroInfo.prototype.getArmorpoint = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 13, 0));
};


/** @param {number} value */
proto.pb.HeroInfo.prototype.setArmorpoint = function(value) {
  jspb.Message.setProto3IntField(this, 13, value);
};


/**
 * optional int64 spellPower = 14;
 * @return {number}
 */
proto.pb.HeroInfo.prototype.getSpellpower = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 14, 0));
};


/** @param {number} value */
proto.pb.HeroInfo.prototype.setSpellpower = function(value) {
  jspb.Message.setProto3IntField(this, 14, value);
};


/**
 * repeated int64 skills = 15;
 * @return {!Array<number>}
 */
proto.pb.HeroInfo.prototype.getSkillsList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 15));
};


/** @param {!Array<number>} value */
proto.pb.HeroInfo.prototype.setSkillsList = function(value) {
  jspb.Message.setField(this, 15, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 */
proto.pb.HeroInfo.prototype.addSkills = function(value, opt_index) {
  jspb.Message.addToRepeatedField(this, 15, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 */
proto.pb.HeroInfo.prototype.clearSkillsList = function() {
  this.setSkillsList([]);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.WeaponInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.WeaponInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.WeaponInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.WeaponInfo.toObject = function(includeInstance, msg) {
  var obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, 0),
    name: jspb.Message.getFieldWithDefault(msg, 2, ""),
    type: jspb.Message.getFieldWithDefault(msg, 3, 0),
    level: jspb.Message.getFieldWithDefault(msg, 4, 0),
    damage: jspb.Message.getFieldWithDefault(msg, 5, 0),
    prob: jspb.Message.getFieldWithDefault(msg, 6, 0),
    count: jspb.Message.getFieldWithDefault(msg, 7, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.WeaponInfo}
 */
proto.pb.WeaponInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.WeaponInfo;
  return proto.pb.WeaponInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.WeaponInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.WeaponInfo}
 */
proto.pb.WeaponInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setType(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setLevel(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setDamage(value);
      break;
    case 6:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setProb(value);
      break;
    case 7:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setCount(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.WeaponInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.WeaponInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.WeaponInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.WeaponInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getType();
  if (f !== 0) {
    writer.writeInt32(
      3,
      f
    );
  }
  f = message.getLevel();
  if (f !== 0) {
    writer.writeInt32(
      4,
      f
    );
  }
  f = message.getDamage();
  if (f !== 0) {
    writer.writeInt64(
      5,
      f
    );
  }
  f = message.getProb();
  if (f !== 0) {
    writer.writeInt64(
      6,
      f
    );
  }
  f = message.getCount();
  if (f !== 0) {
    writer.writeInt32(
      7,
      f
    );
  }
};


/**
 * optional int64 iD = 1;
 * @return {number}
 */
proto.pb.WeaponInfo.prototype.getId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.WeaponInfo.prototype.setId = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional string name = 2;
 * @return {string}
 */
proto.pb.WeaponInfo.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.pb.WeaponInfo.prototype.setName = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional int32 type = 3;
 * @return {number}
 */
proto.pb.WeaponInfo.prototype.getType = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/** @param {number} value */
proto.pb.WeaponInfo.prototype.setType = function(value) {
  jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional int32 level = 4;
 * @return {number}
 */
proto.pb.WeaponInfo.prototype.getLevel = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/** @param {number} value */
proto.pb.WeaponInfo.prototype.setLevel = function(value) {
  jspb.Message.setProto3IntField(this, 4, value);
};


/**
 * optional int64 damage = 5;
 * @return {number}
 */
proto.pb.WeaponInfo.prototype.getDamage = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/** @param {number} value */
proto.pb.WeaponInfo.prototype.setDamage = function(value) {
  jspb.Message.setProto3IntField(this, 5, value);
};


/**
 * optional int64 prob = 6;
 * @return {number}
 */
proto.pb.WeaponInfo.prototype.getProb = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 6, 0));
};


/** @param {number} value */
proto.pb.WeaponInfo.prototype.setProb = function(value) {
  jspb.Message.setProto3IntField(this, 6, value);
};


/**
 * optional int32 count = 7;
 * @return {number}
 */
proto.pb.WeaponInfo.prototype.getCount = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 7, 0));
};


/** @param {number} value */
proto.pb.WeaponInfo.prototype.setCount = function(value) {
  jspb.Message.setProto3IntField(this, 7, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.GoodsInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.GoodsInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.GoodsInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.GoodsInfo.toObject = function(includeInstance, msg) {
  var obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, 0),
    name: jspb.Message.getFieldWithDefault(msg, 2, ""),
    kind: jspb.Message.getFieldWithDefault(msg, 3, 0),
    level: jspb.Message.getFieldWithDefault(msg, 4, 0),
    price: jspb.Message.getFieldWithDefault(msg, 5, 0),
    store: jspb.Message.getFieldWithDefault(msg, 6, 0),
    sold: jspb.Message.getFieldWithDefault(msg, 7, 0),
    amount: jspb.Message.getFieldWithDefault(msg, 8, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.GoodsInfo}
 */
proto.pb.GoodsInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.GoodsInfo;
  return proto.pb.GoodsInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.GoodsInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.GoodsInfo}
 */
proto.pb.GoodsInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setKind(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setLevel(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setPrice(value);
      break;
    case 6:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setStore(value);
      break;
    case 7:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setSold(value);
      break;
    case 8:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setAmount(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.GoodsInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.GoodsInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.GoodsInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.GoodsInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getKind();
  if (f !== 0) {
    writer.writeInt32(
      3,
      f
    );
  }
  f = message.getLevel();
  if (f !== 0) {
    writer.writeInt32(
      4,
      f
    );
  }
  f = message.getPrice();
  if (f !== 0) {
    writer.writeInt64(
      5,
      f
    );
  }
  f = message.getStore();
  if (f !== 0) {
    writer.writeInt64(
      6,
      f
    );
  }
  f = message.getSold();
  if (f !== 0) {
    writer.writeInt64(
      7,
      f
    );
  }
  f = message.getAmount();
  if (f !== 0) {
    writer.writeInt32(
      8,
      f
    );
  }
};


/**
 * optional int64 iD = 1;
 * @return {number}
 */
proto.pb.GoodsInfo.prototype.getId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.GoodsInfo.prototype.setId = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional string name = 2;
 * @return {string}
 */
proto.pb.GoodsInfo.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.pb.GoodsInfo.prototype.setName = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional int32 kind = 3;
 * @return {number}
 */
proto.pb.GoodsInfo.prototype.getKind = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/** @param {number} value */
proto.pb.GoodsInfo.prototype.setKind = function(value) {
  jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional int32 level = 4;
 * @return {number}
 */
proto.pb.GoodsInfo.prototype.getLevel = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/** @param {number} value */
proto.pb.GoodsInfo.prototype.setLevel = function(value) {
  jspb.Message.setProto3IntField(this, 4, value);
};


/**
 * optional int64 price = 5;
 * @return {number}
 */
proto.pb.GoodsInfo.prototype.getPrice = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/** @param {number} value */
proto.pb.GoodsInfo.prototype.setPrice = function(value) {
  jspb.Message.setProto3IntField(this, 5, value);
};


/**
 * optional int64 store = 6;
 * @return {number}
 */
proto.pb.GoodsInfo.prototype.getStore = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 6, 0));
};


/** @param {number} value */
proto.pb.GoodsInfo.prototype.setStore = function(value) {
  jspb.Message.setProto3IntField(this, 6, value);
};


/**
 * optional int64 sold = 7;
 * @return {number}
 */
proto.pb.GoodsInfo.prototype.getSold = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 7, 0));
};


/** @param {number} value */
proto.pb.GoodsInfo.prototype.setSold = function(value) {
  jspb.Message.setProto3IntField(this, 7, value);
};


/**
 * optional int32 amount = 8;
 * @return {number}
 */
proto.pb.GoodsInfo.prototype.getAmount = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 8, 0));
};


/** @param {number} value */
proto.pb.GoodsInfo.prototype.setAmount = function(value) {
  jspb.Message.setProto3IntField(this, 8, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.pb.GoodsList.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.GoodsList.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.GoodsList.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.GoodsList} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.GoodsList.toObject = function(includeInstance, msg) {
  var obj = {
    allgoodsList: jspb.Message.toObjectList(msg.getAllgoodsList(),
    proto.pb.GoodsInfo.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.GoodsList}
 */
proto.pb.GoodsList.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.GoodsList;
  return proto.pb.GoodsList.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.GoodsList} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.GoodsList}
 */
proto.pb.GoodsList.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.pb.GoodsInfo;
      reader.readMessage(value,proto.pb.GoodsInfo.deserializeBinaryFromReader);
      msg.addAllgoods(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.GoodsList.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.GoodsList.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.GoodsList} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.GoodsList.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getAllgoodsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.pb.GoodsInfo.serializeBinaryToWriter
    );
  }
};


/**
 * repeated GoodsInfo allGoods = 1;
 * @return {!Array<!proto.pb.GoodsInfo>}
 */
proto.pb.GoodsList.prototype.getAllgoodsList = function() {
  return /** @type{!Array<!proto.pb.GoodsInfo>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.pb.GoodsInfo, 1));
};


/** @param {!Array<!proto.pb.GoodsInfo>} value */
proto.pb.GoodsList.prototype.setAllgoodsList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.pb.GoodsInfo=} opt_value
 * @param {number=} opt_index
 * @return {!proto.pb.GoodsInfo}
 */
proto.pb.GoodsList.prototype.addAllgoods = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.pb.GoodsInfo, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 */
proto.pb.GoodsList.prototype.clearAllgoodsList = function() {
  this.setAllgoodsList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.pb.KnapsackInfo.repeatedFields_ = [3,4,5];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.KnapsackInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.KnapsackInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.KnapsackInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.KnapsackInfo.toObject = function(includeInstance, msg) {
  var obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, 0),
    name: jspb.Message.getFieldWithDefault(msg, 2, ""),
    myweaponryList: jspb.Message.toObjectList(msg.getMyweaponryList(),
    proto.pb.WeaponInfo.toObject, includeInstance),
    mygoodsList: jspb.Message.toObjectList(msg.getMygoodsList(),
    proto.pb.GoodsInfo.toObject, includeInstance),
    myherolistList: jspb.Message.toObjectList(msg.getMyherolistList(),
    proto.pb.HeroInfo.toObject, includeInstance),
    number: jspb.Message.getFieldWithDefault(msg, 6, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.KnapsackInfo}
 */
proto.pb.KnapsackInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.KnapsackInfo;
  return proto.pb.KnapsackInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.KnapsackInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.KnapsackInfo}
 */
proto.pb.KnapsackInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 3:
      var value = new proto.pb.WeaponInfo;
      reader.readMessage(value,proto.pb.WeaponInfo.deserializeBinaryFromReader);
      msg.addMyweaponry(value);
      break;
    case 4:
      var value = new proto.pb.GoodsInfo;
      reader.readMessage(value,proto.pb.GoodsInfo.deserializeBinaryFromReader);
      msg.addMygoods(value);
      break;
    case 5:
      var value = new proto.pb.HeroInfo;
      reader.readMessage(value,proto.pb.HeroInfo.deserializeBinaryFromReader);
      msg.addMyherolist(value);
      break;
    case 6:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setNumber(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.KnapsackInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.KnapsackInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.KnapsackInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.KnapsackInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getMyweaponryList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      3,
      f,
      proto.pb.WeaponInfo.serializeBinaryToWriter
    );
  }
  f = message.getMygoodsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      4,
      f,
      proto.pb.GoodsInfo.serializeBinaryToWriter
    );
  }
  f = message.getMyherolistList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      5,
      f,
      proto.pb.HeroInfo.serializeBinaryToWriter
    );
  }
  f = message.getNumber();
  if (f !== 0) {
    writer.writeInt32(
      6,
      f
    );
  }
};


/**
 * optional int64 iD = 1;
 * @return {number}
 */
proto.pb.KnapsackInfo.prototype.getId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.KnapsackInfo.prototype.setId = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional string name = 2;
 * @return {string}
 */
proto.pb.KnapsackInfo.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.pb.KnapsackInfo.prototype.setName = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * repeated WeaponInfo myWeaponry = 3;
 * @return {!Array<!proto.pb.WeaponInfo>}
 */
proto.pb.KnapsackInfo.prototype.getMyweaponryList = function() {
  return /** @type{!Array<!proto.pb.WeaponInfo>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.pb.WeaponInfo, 3));
};


/** @param {!Array<!proto.pb.WeaponInfo>} value */
proto.pb.KnapsackInfo.prototype.setMyweaponryList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 3, value);
};


/**
 * @param {!proto.pb.WeaponInfo=} opt_value
 * @param {number=} opt_index
 * @return {!proto.pb.WeaponInfo}
 */
proto.pb.KnapsackInfo.prototype.addMyweaponry = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 3, opt_value, proto.pb.WeaponInfo, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 */
proto.pb.KnapsackInfo.prototype.clearMyweaponryList = function() {
  this.setMyweaponryList([]);
};


/**
 * repeated GoodsInfo myGoods = 4;
 * @return {!Array<!proto.pb.GoodsInfo>}
 */
proto.pb.KnapsackInfo.prototype.getMygoodsList = function() {
  return /** @type{!Array<!proto.pb.GoodsInfo>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.pb.GoodsInfo, 4));
};


/** @param {!Array<!proto.pb.GoodsInfo>} value */
proto.pb.KnapsackInfo.prototype.setMygoodsList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 4, value);
};


/**
 * @param {!proto.pb.GoodsInfo=} opt_value
 * @param {number=} opt_index
 * @return {!proto.pb.GoodsInfo}
 */
proto.pb.KnapsackInfo.prototype.addMygoods = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 4, opt_value, proto.pb.GoodsInfo, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 */
proto.pb.KnapsackInfo.prototype.clearMygoodsList = function() {
  this.setMygoodsList([]);
};


/**
 * repeated HeroInfo myHeroList = 5;
 * @return {!Array<!proto.pb.HeroInfo>}
 */
proto.pb.KnapsackInfo.prototype.getMyherolistList = function() {
  return /** @type{!Array<!proto.pb.HeroInfo>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.pb.HeroInfo, 5));
};


/** @param {!Array<!proto.pb.HeroInfo>} value */
proto.pb.KnapsackInfo.prototype.setMyherolistList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 5, value);
};


/**
 * @param {!proto.pb.HeroInfo=} opt_value
 * @param {number=} opt_index
 * @return {!proto.pb.HeroInfo}
 */
proto.pb.KnapsackInfo.prototype.addMyherolist = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 5, opt_value, proto.pb.HeroInfo, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 */
proto.pb.KnapsackInfo.prototype.clearMyherolistList = function() {
  this.setMyherolistList([]);
};


/**
 * optional int32 number = 6;
 * @return {number}
 */
proto.pb.KnapsackInfo.prototype.getNumber = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 6, 0));
};


/** @param {number} value */
proto.pb.KnapsackInfo.prototype.setNumber = function(value) {
  jspb.Message.setProto3IntField(this, 6, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.EmailInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.EmailInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.EmailInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.EmailInfo.toObject = function(includeInstance, msg) {
  var obj = {
    emailid: jspb.Message.getFieldWithDefault(msg, 1, 0),
    acceptname: jspb.Message.getFieldWithDefault(msg, 2, ""),
    sender: jspb.Message.getFieldWithDefault(msg, 3, ""),
    cc: jspb.Message.getFieldWithDefault(msg, 4, ""),
    topic: jspb.Message.getFieldWithDefault(msg, 5, ""),
    content: jspb.Message.getFieldWithDefault(msg, 6, ""),
    isread: jspb.Message.getFieldWithDefault(msg, 7, false),
    awardlist: (f = msg.getAwardlist()) && proto.pb.GoodsList.toObject(includeInstance, f),
    timestamp: jspb.Message.getFieldWithDefault(msg, 9, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.EmailInfo}
 */
proto.pb.EmailInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.EmailInfo;
  return proto.pb.EmailInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.EmailInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.EmailInfo}
 */
proto.pb.EmailInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setEmailid(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setAcceptname(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setSender(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setCc(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setTopic(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setContent(value);
      break;
    case 7:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setIsread(value);
      break;
    case 8:
      var value = new proto.pb.GoodsList;
      reader.readMessage(value,proto.pb.GoodsList.deserializeBinaryFromReader);
      msg.setAwardlist(value);
      break;
    case 9:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setTimestamp(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.EmailInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.EmailInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.EmailInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.EmailInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getEmailid();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getAcceptname();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getSender();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getCc();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getTopic();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getContent();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
  f = message.getIsread();
  if (f) {
    writer.writeBool(
      7,
      f
    );
  }
  f = message.getAwardlist();
  if (f != null) {
    writer.writeMessage(
      8,
      f,
      proto.pb.GoodsList.serializeBinaryToWriter
    );
  }
  f = message.getTimestamp();
  if (f !== 0) {
    writer.writeInt64(
      9,
      f
    );
  }
};


/**
 * optional int64 emailID = 1;
 * @return {number}
 */
proto.pb.EmailInfo.prototype.getEmailid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.EmailInfo.prototype.setEmailid = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional string acceptName = 2;
 * @return {string}
 */
proto.pb.EmailInfo.prototype.getAcceptname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.pb.EmailInfo.prototype.setAcceptname = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string sender = 3;
 * @return {string}
 */
proto.pb.EmailInfo.prototype.getSender = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.pb.EmailInfo.prototype.setSender = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string cc = 4;
 * @return {string}
 */
proto.pb.EmailInfo.prototype.getCc = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.pb.EmailInfo.prototype.setCc = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string topic = 5;
 * @return {string}
 */
proto.pb.EmailInfo.prototype.getTopic = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/** @param {string} value */
proto.pb.EmailInfo.prototype.setTopic = function(value) {
  jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional string content = 6;
 * @return {string}
 */
proto.pb.EmailInfo.prototype.getContent = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/** @param {string} value */
proto.pb.EmailInfo.prototype.setContent = function(value) {
  jspb.Message.setProto3StringField(this, 6, value);
};


/**
 * optional bool isRead = 7;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.pb.EmailInfo.prototype.getIsread = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 7, false));
};


/** @param {boolean} value */
proto.pb.EmailInfo.prototype.setIsread = function(value) {
  jspb.Message.setProto3BooleanField(this, 7, value);
};


/**
 * optional GoodsList awardList = 8;
 * @return {?proto.pb.GoodsList}
 */
proto.pb.EmailInfo.prototype.getAwardlist = function() {
  return /** @type{?proto.pb.GoodsList} */ (
    jspb.Message.getWrapperField(this, proto.pb.GoodsList, 8));
};


/** @param {?proto.pb.GoodsList|undefined} value */
proto.pb.EmailInfo.prototype.setAwardlist = function(value) {
  jspb.Message.setWrapperField(this, 8, value);
};


/**
 * Clears the message field making it undefined.
 */
proto.pb.EmailInfo.prototype.clearAwardlist = function() {
  this.setAwardlist(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.pb.EmailInfo.prototype.hasAwardlist = function() {
  return jspb.Message.getField(this, 8) != null;
};


/**
 * optional int64 timeStamp = 9;
 * @return {number}
 */
proto.pb.EmailInfo.prototype.getTimestamp = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 9, 0));
};


/** @param {number} value */
proto.pb.EmailInfo.prototype.setTimestamp = function(value) {
  jspb.Message.setProto3IntField(this, 9, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.TableInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.TableInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.TableInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.TableInfo.toObject = function(includeInstance, msg) {
  var obj = {
    hostid: jspb.Message.getFieldWithDefault(msg, 1, 0),
    name: jspb.Message.getFieldWithDefault(msg, 2, ""),
    password: jspb.Message.getFieldWithDefault(msg, 3, ""),
    state: jspb.Message.getFieldWithDefault(msg, 4, 0),
    enterscore: jspb.Message.getFieldWithDefault(msg, 5, 0),
    lessscore: jspb.Message.getFieldWithDefault(msg, 6, 0),
    playscore: jspb.Message.getFieldWithDefault(msg, 7, 0),
    commission: jspb.Message.getFieldWithDefault(msg, 8, 0),
    maxchair: jspb.Message.getFieldWithDefault(msg, 9, 0),
    amount: jspb.Message.getFieldWithDefault(msg, 10, 0),
    maxonline: jspb.Message.getFieldWithDefault(msg, 11, 0),
    robotcount: jspb.Message.getFieldWithDefault(msg, 12, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.TableInfo}
 */
proto.pb.TableInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.TableInfo;
  return proto.pb.TableInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.TableInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.TableInfo}
 */
proto.pb.TableInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setHostid(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setPassword(value);
      break;
    case 4:
      var value = /** @type {!proto.pb.TableState} */ (reader.readEnum());
      msg.setState(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setEnterscore(value);
      break;
    case 6:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setLessscore(value);
      break;
    case 7:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setPlayscore(value);
      break;
    case 8:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setCommission(value);
      break;
    case 9:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setMaxchair(value);
      break;
    case 10:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setAmount(value);
      break;
    case 11:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setMaxonline(value);
      break;
    case 12:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setRobotcount(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.TableInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.TableInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.TableInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.TableInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getHostid();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getPassword();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getState();
  if (f !== 0.0) {
    writer.writeEnum(
      4,
      f
    );
  }
  f = message.getEnterscore();
  if (f !== 0) {
    writer.writeInt32(
      5,
      f
    );
  }
  f = message.getLessscore();
  if (f !== 0) {
    writer.writeInt32(
      6,
      f
    );
  }
  f = message.getPlayscore();
  if (f !== 0) {
    writer.writeInt64(
      7,
      f
    );
  }
  f = message.getCommission();
  if (f !== 0) {
    writer.writeInt32(
      8,
      f
    );
  }
  f = message.getMaxchair();
  if (f !== 0) {
    writer.writeInt32(
      9,
      f
    );
  }
  f = message.getAmount();
  if (f !== 0) {
    writer.writeInt32(
      10,
      f
    );
  }
  f = message.getMaxonline();
  if (f !== 0) {
    writer.writeInt32(
      11,
      f
    );
  }
  f = message.getRobotcount();
  if (f !== 0) {
    writer.writeInt32(
      12,
      f
    );
  }
};


/**
 * optional int64 hostID = 1;
 * @return {number}
 */
proto.pb.TableInfo.prototype.getHostid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.TableInfo.prototype.setHostid = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional string name = 2;
 * @return {string}
 */
proto.pb.TableInfo.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.pb.TableInfo.prototype.setName = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string password = 3;
 * @return {string}
 */
proto.pb.TableInfo.prototype.getPassword = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.pb.TableInfo.prototype.setPassword = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional TableState state = 4;
 * @return {!proto.pb.TableState}
 */
proto.pb.TableInfo.prototype.getState = function() {
  return /** @type {!proto.pb.TableState} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/** @param {!proto.pb.TableState} value */
proto.pb.TableInfo.prototype.setState = function(value) {
  jspb.Message.setProto3EnumField(this, 4, value);
};


/**
 * optional int32 enterScore = 5;
 * @return {number}
 */
proto.pb.TableInfo.prototype.getEnterscore = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/** @param {number} value */
proto.pb.TableInfo.prototype.setEnterscore = function(value) {
  jspb.Message.setProto3IntField(this, 5, value);
};


/**
 * optional int32 lessScore = 6;
 * @return {number}
 */
proto.pb.TableInfo.prototype.getLessscore = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 6, 0));
};


/** @param {number} value */
proto.pb.TableInfo.prototype.setLessscore = function(value) {
  jspb.Message.setProto3IntField(this, 6, value);
};


/**
 * optional int64 playScore = 7;
 * @return {number}
 */
proto.pb.TableInfo.prototype.getPlayscore = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 7, 0));
};


/** @param {number} value */
proto.pb.TableInfo.prototype.setPlayscore = function(value) {
  jspb.Message.setProto3IntField(this, 7, value);
};


/**
 * optional int32 commission = 8;
 * @return {number}
 */
proto.pb.TableInfo.prototype.getCommission = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 8, 0));
};


/** @param {number} value */
proto.pb.TableInfo.prototype.setCommission = function(value) {
  jspb.Message.setProto3IntField(this, 8, value);
};


/**
 * optional int32 maxChair = 9;
 * @return {number}
 */
proto.pb.TableInfo.prototype.getMaxchair = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 9, 0));
};


/** @param {number} value */
proto.pb.TableInfo.prototype.setMaxchair = function(value) {
  jspb.Message.setProto3IntField(this, 9, value);
};


/**
 * optional int32 amount = 10;
 * @return {number}
 */
proto.pb.TableInfo.prototype.getAmount = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 10, 0));
};


/** @param {number} value */
proto.pb.TableInfo.prototype.setAmount = function(value) {
  jspb.Message.setProto3IntField(this, 10, value);
};


/**
 * optional int32 maxOnline = 11;
 * @return {number}
 */
proto.pb.TableInfo.prototype.getMaxonline = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 11, 0));
};


/** @param {number} value */
proto.pb.TableInfo.prototype.setMaxonline = function(value) {
  jspb.Message.setProto3IntField(this, 11, value);
};


/**
 * optional int32 robotCount = 12;
 * @return {number}
 */
proto.pb.TableInfo.prototype.getRobotcount = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 12, 0));
};


/** @param {number} value */
proto.pb.TableInfo.prototype.setRobotcount = function(value) {
  jspb.Message.setProto3IntField(this, 12, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.GameInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.GameInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.GameInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.GameInfo.toObject = function(includeInstance, msg) {
  var obj = {
    type: jspb.Message.getFieldWithDefault(msg, 1, 0),
    kindid: jspb.Message.getFieldWithDefault(msg, 2, 0),
    level: jspb.Message.getFieldWithDefault(msg, 3, 0),
    scene: jspb.Message.getFieldWithDefault(msg, 4, 0),
    name: jspb.Message.getFieldWithDefault(msg, 5, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.GameInfo}
 */
proto.pb.GameInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.GameInfo;
  return proto.pb.GameInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.GameInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.GameInfo}
 */
proto.pb.GameInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.pb.GameType} */ (reader.readEnum());
      msg.setType(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setKindid(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setLevel(value);
      break;
    case 4:
      var value = /** @type {!proto.pb.GameScene} */ (reader.readEnum());
      msg.setScene(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.GameInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.GameInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.GameInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.GameInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getType();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getKindid();
  if (f !== 0) {
    writer.writeInt32(
      2,
      f
    );
  }
  f = message.getLevel();
  if (f !== 0) {
    writer.writeInt32(
      3,
      f
    );
  }
  f = message.getScene();
  if (f !== 0.0) {
    writer.writeEnum(
      4,
      f
    );
  }
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
};


/**
 * optional GameType type = 1;
 * @return {!proto.pb.GameType}
 */
proto.pb.GameInfo.prototype.getType = function() {
  return /** @type {!proto.pb.GameType} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {!proto.pb.GameType} value */
proto.pb.GameInfo.prototype.setType = function(value) {
  jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional int32 kindID = 2;
 * @return {number}
 */
proto.pb.GameInfo.prototype.getKindid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {number} value */
proto.pb.GameInfo.prototype.setKindid = function(value) {
  jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional int32 level = 3;
 * @return {number}
 */
proto.pb.GameInfo.prototype.getLevel = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/** @param {number} value */
proto.pb.GameInfo.prototype.setLevel = function(value) {
  jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional GameScene scene = 4;
 * @return {!proto.pb.GameScene}
 */
proto.pb.GameInfo.prototype.getScene = function() {
  return /** @type {!proto.pb.GameScene} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/** @param {!proto.pb.GameScene} value */
proto.pb.GameInfo.prototype.setScene = function(value) {
  jspb.Message.setProto3EnumField(this, 4, value);
};


/**
 * optional string name = 5;
 * @return {string}
 */
proto.pb.GameInfo.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/** @param {string} value */
proto.pb.GameInfo.prototype.setName = function(value) {
  jspb.Message.setProto3StringField(this, 5, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.TaskItem.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.TaskItem.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.TaskItem} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.TaskItem.toObject = function(includeInstance, msg) {
  var obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, 0),
    twice: jspb.Message.getFieldWithDefault(msg, 2, 0),
    hints: jspb.Message.getFieldWithDefault(msg, 3, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.TaskItem}
 */
proto.pb.TaskItem.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.TaskItem;
  return proto.pb.TaskItem.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.TaskItem} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.TaskItem}
 */
proto.pb.TaskItem.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setTwice(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setHints(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.TaskItem.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.TaskItem.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.TaskItem} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.TaskItem.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f !== 0) {
    writer.writeInt32(
      1,
      f
    );
  }
  f = message.getTwice();
  if (f !== 0) {
    writer.writeInt32(
      2,
      f
    );
  }
  f = message.getHints();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
};


/**
 * optional int32 iD = 1;
 * @return {number}
 */
proto.pb.TaskItem.prototype.getId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.TaskItem.prototype.setId = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional int32 twice = 2;
 * @return {number}
 */
proto.pb.TaskItem.prototype.getTwice = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {number} value */
proto.pb.TaskItem.prototype.setTwice = function(value) {
  jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional string hints = 3;
 * @return {string}
 */
proto.pb.TaskItem.prototype.getHints = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.pb.TaskItem.prototype.setHints = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.ClassItem.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.ClassItem.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.ClassItem} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.ClassItem.toObject = function(includeInstance, msg) {
  var obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, 0),
    name: jspb.Message.getFieldWithDefault(msg, 2, ""),
    key: jspb.Message.getFieldWithDefault(msg, 3, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.ClassItem}
 */
proto.pb.ClassItem.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.ClassItem;
  return proto.pb.ClassItem.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.ClassItem} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.ClassItem}
 */
proto.pb.ClassItem.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setKey(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.ClassItem.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.ClassItem.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.ClassItem} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.ClassItem.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getKey();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
};


/**
 * optional int64 iD = 1;
 * @return {number}
 */
proto.pb.ClassItem.prototype.getId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.ClassItem.prototype.setId = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional string name = 2;
 * @return {string}
 */
proto.pb.ClassItem.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.pb.ClassItem.prototype.setName = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string key = 3;
 * @return {string}
 */
proto.pb.ClassItem.prototype.getKey = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.pb.ClassItem.prototype.setKey = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.GameItem.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.GameItem.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.GameItem} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.GameItem.toObject = function(includeInstance, msg) {
  var obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, 0),
    info: (f = msg.getInfo()) && proto.pb.GameInfo.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.GameItem}
 */
proto.pb.GameItem.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.GameItem;
  return proto.pb.GameItem.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.GameItem} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.GameItem}
 */
proto.pb.GameItem.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setId(value);
      break;
    case 2:
      var value = new proto.pb.GameInfo;
      reader.readMessage(value,proto.pb.GameInfo.deserializeBinaryFromReader);
      msg.setInfo(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.GameItem.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.GameItem.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.GameItem} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.GameItem.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getInfo();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.pb.GameInfo.serializeBinaryToWriter
    );
  }
};


/**
 * optional int64 iD = 1;
 * @return {number}
 */
proto.pb.GameItem.prototype.getId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.GameItem.prototype.setId = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional GameInfo info = 2;
 * @return {?proto.pb.GameInfo}
 */
proto.pb.GameItem.prototype.getInfo = function() {
  return /** @type{?proto.pb.GameInfo} */ (
    jspb.Message.getWrapperField(this, proto.pb.GameInfo, 2));
};


/** @param {?proto.pb.GameInfo|undefined} value */
proto.pb.GameItem.prototype.setInfo = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 */
proto.pb.GameItem.prototype.clearInfo = function() {
  this.setInfo(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.pb.GameItem.prototype.hasInfo = function() {
  return jspb.Message.getField(this, 2) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.TableItem.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.TableItem.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.TableItem} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.TableItem.toObject = function(includeInstance, msg) {
  var obj = {
    num: jspb.Message.getFieldWithDefault(msg, 1, 0),
    gameid: jspb.Message.getFieldWithDefault(msg, 2, 0),
    info: (f = msg.getInfo()) && proto.pb.TableInfo.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.TableItem}
 */
proto.pb.TableItem.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.TableItem;
  return proto.pb.TableItem.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.TableItem} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.TableItem}
 */
proto.pb.TableItem.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setNum(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setGameid(value);
      break;
    case 3:
      var value = new proto.pb.TableInfo;
      reader.readMessage(value,proto.pb.TableInfo.deserializeBinaryFromReader);
      msg.setInfo(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.TableItem.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.TableItem.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.TableItem} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.TableItem.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getNum();
  if (f !== 0) {
    writer.writeInt32(
      1,
      f
    );
  }
  f = message.getGameid();
  if (f !== 0) {
    writer.writeInt64(
      2,
      f
    );
  }
  f = message.getInfo();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      proto.pb.TableInfo.serializeBinaryToWriter
    );
  }
};


/**
 * optional int32 num = 1;
 * @return {number}
 */
proto.pb.TableItem.prototype.getNum = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.TableItem.prototype.setNum = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional int64 gameID = 2;
 * @return {number}
 */
proto.pb.TableItem.prototype.getGameid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {number} value */
proto.pb.TableItem.prototype.setGameid = function(value) {
  jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional TableInfo info = 3;
 * @return {?proto.pb.TableInfo}
 */
proto.pb.TableItem.prototype.getInfo = function() {
  return /** @type{?proto.pb.TableInfo} */ (
    jspb.Message.getWrapperField(this, proto.pb.TableInfo, 3));
};


/** @param {?proto.pb.TableInfo|undefined} value */
proto.pb.TableItem.prototype.setInfo = function(value) {
  jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 */
proto.pb.TableItem.prototype.clearInfo = function() {
  this.setInfo(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.pb.TableItem.prototype.hasInfo = function() {
  return jspb.Message.getField(this, 3) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.pb.TaskList.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.TaskList.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.TaskList.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.TaskList} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.TaskList.toObject = function(includeInstance, msg) {
  var obj = {
    taskList: jspb.Message.toObjectList(msg.getTaskList(),
    proto.pb.TaskItem.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.TaskList}
 */
proto.pb.TaskList.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.TaskList;
  return proto.pb.TaskList.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.TaskList} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.TaskList}
 */
proto.pb.TaskList.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.pb.TaskItem;
      reader.readMessage(value,proto.pb.TaskItem.deserializeBinaryFromReader);
      msg.addTask(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.TaskList.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.TaskList.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.TaskList} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.TaskList.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getTaskList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.pb.TaskItem.serializeBinaryToWriter
    );
  }
};


/**
 * repeated TaskItem task = 1;
 * @return {!Array<!proto.pb.TaskItem>}
 */
proto.pb.TaskList.prototype.getTaskList = function() {
  return /** @type{!Array<!proto.pb.TaskItem>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.pb.TaskItem, 1));
};


/** @param {!Array<!proto.pb.TaskItem>} value */
proto.pb.TaskList.prototype.setTaskList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.pb.TaskItem=} opt_value
 * @param {number=} opt_index
 * @return {!proto.pb.TaskItem}
 */
proto.pb.TaskList.prototype.addTask = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.pb.TaskItem, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 */
proto.pb.TaskList.prototype.clearTaskList = function() {
  this.setTaskList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.pb.ClassList.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.ClassList.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.ClassList.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.ClassList} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.ClassList.toObject = function(includeInstance, msg) {
  var obj = {
    classifyList: jspb.Message.toObjectList(msg.getClassifyList(),
    proto.pb.ClassItem.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.ClassList}
 */
proto.pb.ClassList.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.ClassList;
  return proto.pb.ClassList.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.ClassList} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.ClassList}
 */
proto.pb.ClassList.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.pb.ClassItem;
      reader.readMessage(value,proto.pb.ClassItem.deserializeBinaryFromReader);
      msg.addClassify(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.ClassList.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.ClassList.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.ClassList} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.ClassList.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getClassifyList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.pb.ClassItem.serializeBinaryToWriter
    );
  }
};


/**
 * repeated ClassItem classify = 1;
 * @return {!Array<!proto.pb.ClassItem>}
 */
proto.pb.ClassList.prototype.getClassifyList = function() {
  return /** @type{!Array<!proto.pb.ClassItem>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.pb.ClassItem, 1));
};


/** @param {!Array<!proto.pb.ClassItem>} value */
proto.pb.ClassList.prototype.setClassifyList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.pb.ClassItem=} opt_value
 * @param {number=} opt_index
 * @return {!proto.pb.ClassItem}
 */
proto.pb.ClassList.prototype.addClassify = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.pb.ClassItem, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 */
proto.pb.ClassList.prototype.clearClassifyList = function() {
  this.setClassifyList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.pb.GameList.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.GameList.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.GameList.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.GameList} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.GameList.toObject = function(includeInstance, msg) {
  var obj = {
    itemsList: jspb.Message.toObjectList(msg.getItemsList(),
    proto.pb.GameItem.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.GameList}
 */
proto.pb.GameList.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.GameList;
  return proto.pb.GameList.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.GameList} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.GameList}
 */
proto.pb.GameList.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.pb.GameItem;
      reader.readMessage(value,proto.pb.GameItem.deserializeBinaryFromReader);
      msg.addItems(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.GameList.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.GameList.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.GameList} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.GameList.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getItemsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.pb.GameItem.serializeBinaryToWriter
    );
  }
};


/**
 * repeated GameItem items = 1;
 * @return {!Array<!proto.pb.GameItem>}
 */
proto.pb.GameList.prototype.getItemsList = function() {
  return /** @type{!Array<!proto.pb.GameItem>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.pb.GameItem, 1));
};


/** @param {!Array<!proto.pb.GameItem>} value */
proto.pb.GameList.prototype.setItemsList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.pb.GameItem=} opt_value
 * @param {number=} opt_index
 * @return {!proto.pb.GameItem}
 */
proto.pb.GameList.prototype.addItems = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.pb.GameItem, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 */
proto.pb.GameList.prototype.clearItemsList = function() {
  this.setItemsList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.pb.TableList.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.TableList.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.TableList.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.TableList} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.TableList.toObject = function(includeInstance, msg) {
  var obj = {
    itemsList: jspb.Message.toObjectList(msg.getItemsList(),
    proto.pb.TableItem.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.TableList}
 */
proto.pb.TableList.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.TableList;
  return proto.pb.TableList.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.TableList} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.TableList}
 */
proto.pb.TableList.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.pb.TableItem;
      reader.readMessage(value,proto.pb.TableItem.deserializeBinaryFromReader);
      msg.addItems(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.TableList.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.TableList.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.TableList} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.TableList.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getItemsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.pb.TableItem.serializeBinaryToWriter
    );
  }
};


/**
 * repeated TableItem items = 1;
 * @return {!Array<!proto.pb.TableItem>}
 */
proto.pb.TableList.prototype.getItemsList = function() {
  return /** @type{!Array<!proto.pb.TableItem>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.pb.TableItem, 1));
};


/** @param {!Array<!proto.pb.TableItem>} value */
proto.pb.TableList.prototype.setItemsList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.pb.TableItem=} opt_value
 * @param {number=} opt_index
 * @return {!proto.pb.TableItem}
 */
proto.pb.TableList.prototype.addItems = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.pb.TableItem, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 */
proto.pb.TableList.prototype.clearItemsList = function() {
  this.setItemsList([]);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.MasterInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.MasterInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.MasterInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.MasterInfo.toObject = function(includeInstance, msg) {
  var obj = {
    userinfo: (f = msg.getUserinfo()) && proto.pb.UserInfo.toObject(includeInstance, f),
    classes: (f = msg.getClasses()) && proto.pb.ClassList.toObject(includeInstance, f),
    tasks: (f = msg.getTasks()) && proto.pb.TaskList.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.MasterInfo}
 */
proto.pb.MasterInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.MasterInfo;
  return proto.pb.MasterInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.MasterInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.MasterInfo}
 */
proto.pb.MasterInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.pb.UserInfo;
      reader.readMessage(value,proto.pb.UserInfo.deserializeBinaryFromReader);
      msg.setUserinfo(value);
      break;
    case 2:
      var value = new proto.pb.ClassList;
      reader.readMessage(value,proto.pb.ClassList.deserializeBinaryFromReader);
      msg.setClasses(value);
      break;
    case 3:
      var value = new proto.pb.TaskList;
      reader.readMessage(value,proto.pb.TaskList.deserializeBinaryFromReader);
      msg.setTasks(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.MasterInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.MasterInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.MasterInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.MasterInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUserinfo();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.pb.UserInfo.serializeBinaryToWriter
    );
  }
  f = message.getClasses();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.pb.ClassList.serializeBinaryToWriter
    );
  }
  f = message.getTasks();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      proto.pb.TaskList.serializeBinaryToWriter
    );
  }
};


/**
 * optional UserInfo userInfo = 1;
 * @return {?proto.pb.UserInfo}
 */
proto.pb.MasterInfo.prototype.getUserinfo = function() {
  return /** @type{?proto.pb.UserInfo} */ (
    jspb.Message.getWrapperField(this, proto.pb.UserInfo, 1));
};


/** @param {?proto.pb.UserInfo|undefined} value */
proto.pb.MasterInfo.prototype.setUserinfo = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 */
proto.pb.MasterInfo.prototype.clearUserinfo = function() {
  this.setUserinfo(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.pb.MasterInfo.prototype.hasUserinfo = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional ClassList classes = 2;
 * @return {?proto.pb.ClassList}
 */
proto.pb.MasterInfo.prototype.getClasses = function() {
  return /** @type{?proto.pb.ClassList} */ (
    jspb.Message.getWrapperField(this, proto.pb.ClassList, 2));
};


/** @param {?proto.pb.ClassList|undefined} value */
proto.pb.MasterInfo.prototype.setClasses = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 */
proto.pb.MasterInfo.prototype.clearClasses = function() {
  this.setClasses(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.pb.MasterInfo.prototype.hasClasses = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional TaskList tasks = 3;
 * @return {?proto.pb.TaskList}
 */
proto.pb.MasterInfo.prototype.getTasks = function() {
  return /** @type{?proto.pb.TaskList} */ (
    jspb.Message.getWrapperField(this, proto.pb.TaskList, 3));
};


/** @param {?proto.pb.TaskList|undefined} value */
proto.pb.MasterInfo.prototype.setTasks = function(value) {
  jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 */
proto.pb.MasterInfo.prototype.clearTasks = function() {
  this.setTasks(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.pb.MasterInfo.prototype.hasTasks = function() {
  return jspb.Message.getField(this, 3) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.RegisterReq.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.RegisterReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.RegisterReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.RegisterReq.toObject = function(includeInstance, msg) {
  var obj = {
    name: jspb.Message.getFieldWithDefault(msg, 1, ""),
    password: jspb.Message.getFieldWithDefault(msg, 2, ""),
    securitycode: jspb.Message.getFieldWithDefault(msg, 3, ""),
    machinecode: jspb.Message.getFieldWithDefault(msg, 4, ""),
    invitationcode: jspb.Message.getFieldWithDefault(msg, 5, ""),
    platformid: jspb.Message.getFieldWithDefault(msg, 6, 0),
    gender: jspb.Message.getFieldWithDefault(msg, 7, 0),
    age: jspb.Message.getFieldWithDefault(msg, 8, 0),
    faceid: jspb.Message.getFieldWithDefault(msg, 9, 0),
    passportid: jspb.Message.getFieldWithDefault(msg, 10, ""),
    realname: jspb.Message.getFieldWithDefault(msg, 11, ""),
    phonenum: jspb.Message.getFieldWithDefault(msg, 12, ""),
    email: jspb.Message.getFieldWithDefault(msg, 13, ""),
    address: jspb.Message.getFieldWithDefault(msg, 14, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.RegisterReq}
 */
proto.pb.RegisterReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.RegisterReq;
  return proto.pb.RegisterReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.RegisterReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.RegisterReq}
 */
proto.pb.RegisterReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setPassword(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setSecuritycode(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setMachinecode(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setInvitationcode(value);
      break;
    case 6:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setPlatformid(value);
      break;
    case 7:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setGender(value);
      break;
    case 8:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setAge(value);
      break;
    case 9:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setFaceid(value);
      break;
    case 10:
      var value = /** @type {string} */ (reader.readString());
      msg.setPassportid(value);
      break;
    case 11:
      var value = /** @type {string} */ (reader.readString());
      msg.setRealname(value);
      break;
    case 12:
      var value = /** @type {string} */ (reader.readString());
      msg.setPhonenum(value);
      break;
    case 13:
      var value = /** @type {string} */ (reader.readString());
      msg.setEmail(value);
      break;
    case 14:
      var value = /** @type {string} */ (reader.readString());
      msg.setAddress(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.RegisterReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.RegisterReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.RegisterReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.RegisterReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getPassword();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getSecuritycode();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getMachinecode();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getInvitationcode();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getPlatformid();
  if (f !== 0) {
    writer.writeInt64(
      6,
      f
    );
  }
  f = message.getGender();
  if (f !== 0) {
    writer.writeInt32(
      7,
      f
    );
  }
  f = message.getAge();
  if (f !== 0) {
    writer.writeInt32(
      8,
      f
    );
  }
  f = message.getFaceid();
  if (f !== 0) {
    writer.writeInt32(
      9,
      f
    );
  }
  f = message.getPassportid();
  if (f.length > 0) {
    writer.writeString(
      10,
      f
    );
  }
  f = message.getRealname();
  if (f.length > 0) {
    writer.writeString(
      11,
      f
    );
  }
  f = message.getPhonenum();
  if (f.length > 0) {
    writer.writeString(
      12,
      f
    );
  }
  f = message.getEmail();
  if (f.length > 0) {
    writer.writeString(
      13,
      f
    );
  }
  f = message.getAddress();
  if (f.length > 0) {
    writer.writeString(
      14,
      f
    );
  }
};


/**
 * optional string name = 1;
 * @return {string}
 */
proto.pb.RegisterReq.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.pb.RegisterReq.prototype.setName = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string password = 2;
 * @return {string}
 */
proto.pb.RegisterReq.prototype.getPassword = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.pb.RegisterReq.prototype.setPassword = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string securityCode = 3;
 * @return {string}
 */
proto.pb.RegisterReq.prototype.getSecuritycode = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.pb.RegisterReq.prototype.setSecuritycode = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string machineCode = 4;
 * @return {string}
 */
proto.pb.RegisterReq.prototype.getMachinecode = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.pb.RegisterReq.prototype.setMachinecode = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string invitationCode = 5;
 * @return {string}
 */
proto.pb.RegisterReq.prototype.getInvitationcode = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/** @param {string} value */
proto.pb.RegisterReq.prototype.setInvitationcode = function(value) {
  jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional int64 platformID = 6;
 * @return {number}
 */
proto.pb.RegisterReq.prototype.getPlatformid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 6, 0));
};


/** @param {number} value */
proto.pb.RegisterReq.prototype.setPlatformid = function(value) {
  jspb.Message.setProto3IntField(this, 6, value);
};


/**
 * optional int32 gender = 7;
 * @return {number}
 */
proto.pb.RegisterReq.prototype.getGender = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 7, 0));
};


/** @param {number} value */
proto.pb.RegisterReq.prototype.setGender = function(value) {
  jspb.Message.setProto3IntField(this, 7, value);
};


/**
 * optional int32 age = 8;
 * @return {number}
 */
proto.pb.RegisterReq.prototype.getAge = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 8, 0));
};


/** @param {number} value */
proto.pb.RegisterReq.prototype.setAge = function(value) {
  jspb.Message.setProto3IntField(this, 8, value);
};


/**
 * optional int32 faceID = 9;
 * @return {number}
 */
proto.pb.RegisterReq.prototype.getFaceid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 9, 0));
};


/** @param {number} value */
proto.pb.RegisterReq.prototype.setFaceid = function(value) {
  jspb.Message.setProto3IntField(this, 9, value);
};


/**
 * optional string passPortID = 10;
 * @return {string}
 */
proto.pb.RegisterReq.prototype.getPassportid = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 10, ""));
};


/** @param {string} value */
proto.pb.RegisterReq.prototype.setPassportid = function(value) {
  jspb.Message.setProto3StringField(this, 10, value);
};


/**
 * optional string realName = 11;
 * @return {string}
 */
proto.pb.RegisterReq.prototype.getRealname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 11, ""));
};


/** @param {string} value */
proto.pb.RegisterReq.prototype.setRealname = function(value) {
  jspb.Message.setProto3StringField(this, 11, value);
};


/**
 * optional string phoneNum = 12;
 * @return {string}
 */
proto.pb.RegisterReq.prototype.getPhonenum = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 12, ""));
};


/** @param {string} value */
proto.pb.RegisterReq.prototype.setPhonenum = function(value) {
  jspb.Message.setProto3StringField(this, 12, value);
};


/**
 * optional string email = 13;
 * @return {string}
 */
proto.pb.RegisterReq.prototype.getEmail = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 13, ""));
};


/** @param {string} value */
proto.pb.RegisterReq.prototype.setEmail = function(value) {
  jspb.Message.setProto3StringField(this, 13, value);
};


/**
 * optional string address = 14;
 * @return {string}
 */
proto.pb.RegisterReq.prototype.getAddress = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 14, ""));
};


/** @param {string} value */
proto.pb.RegisterReq.prototype.setAddress = function(value) {
  jspb.Message.setProto3StringField(this, 14, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.RegisterResp.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.RegisterResp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.RegisterResp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.RegisterResp.toObject = function(includeInstance, msg) {
  var obj = {
    sdkid: jspb.Message.getFieldWithDefault(msg, 2, 0),
    pid: jspb.Message.getFieldWithDefault(msg, 3, 0),
    openid: jspb.Message.getFieldWithDefault(msg, 4, ""),
    serverid: jspb.Message.getFieldWithDefault(msg, 5, 0),
    ip: jspb.Message.getFieldWithDefault(msg, 6, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.RegisterResp}
 */
proto.pb.RegisterResp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.RegisterResp;
  return proto.pb.RegisterResp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.RegisterResp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.RegisterResp}
 */
proto.pb.RegisterResp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 2:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setSdkid(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setPid(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setOpenid(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setServerid(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setIp(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.RegisterResp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.RegisterResp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.RegisterResp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.RegisterResp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getSdkid();
  if (f !== 0) {
    writer.writeInt32(
      2,
      f
    );
  }
  f = message.getPid();
  if (f !== 0) {
    writer.writeInt32(
      3,
      f
    );
  }
  f = message.getOpenid();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getServerid();
  if (f !== 0) {
    writer.writeInt32(
      5,
      f
    );
  }
  f = message.getIp();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
};


/**
 * optional int32 sdkId = 2;
 * @return {number}
 */
proto.pb.RegisterResp.prototype.getSdkid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {number} value */
proto.pb.RegisterResp.prototype.setSdkid = function(value) {
  jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional int32 pid = 3;
 * @return {number}
 */
proto.pb.RegisterResp.prototype.getPid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/** @param {number} value */
proto.pb.RegisterResp.prototype.setPid = function(value) {
  jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional string openId = 4;
 * @return {string}
 */
proto.pb.RegisterResp.prototype.getOpenid = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.pb.RegisterResp.prototype.setOpenid = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional int32 serverId = 5;
 * @return {number}
 */
proto.pb.RegisterResp.prototype.getServerid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/** @param {number} value */
proto.pb.RegisterResp.prototype.setServerid = function(value) {
  jspb.Message.setProto3IntField(this, 5, value);
};


/**
 * optional string ip = 6;
 * @return {string}
 */
proto.pb.RegisterResp.prototype.getIp = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/** @param {string} value */
proto.pb.RegisterResp.prototype.setIp = function(value) {
  jspb.Message.setProto3StringField(this, 6, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.LoginReq.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.LoginReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.LoginReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.LoginReq.toObject = function(includeInstance, msg) {
  var obj = {
    account: jspb.Message.getFieldWithDefault(msg, 1, ""),
    password: jspb.Message.getFieldWithDefault(msg, 2, ""),
    securitycode: jspb.Message.getFieldWithDefault(msg, 3, ""),
    machinecode: jspb.Message.getFieldWithDefault(msg, 4, ""),
    platformid: jspb.Message.getFieldWithDefault(msg, 5, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.LoginReq}
 */
proto.pb.LoginReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.LoginReq;
  return proto.pb.LoginReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.LoginReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.LoginReq}
 */
proto.pb.LoginReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setAccount(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setPassword(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setSecuritycode(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setMachinecode(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setPlatformid(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.LoginReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.LoginReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.LoginReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.LoginReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getAccount();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getPassword();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getSecuritycode();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getMachinecode();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getPlatformid();
  if (f !== 0) {
    writer.writeInt64(
      5,
      f
    );
  }
};


/**
 * optional string account = 1;
 * @return {string}
 */
proto.pb.LoginReq.prototype.getAccount = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.pb.LoginReq.prototype.setAccount = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string password = 2;
 * @return {string}
 */
proto.pb.LoginReq.prototype.getPassword = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.pb.LoginReq.prototype.setPassword = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string securityCode = 3;
 * @return {string}
 */
proto.pb.LoginReq.prototype.getSecuritycode = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.pb.LoginReq.prototype.setSecuritycode = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string machineCode = 4;
 * @return {string}
 */
proto.pb.LoginReq.prototype.getMachinecode = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.pb.LoginReq.prototype.setMachinecode = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional int64 platformID = 5;
 * @return {number}
 */
proto.pb.LoginReq.prototype.getPlatformid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/** @param {number} value */
proto.pb.LoginReq.prototype.setPlatformid = function(value) {
  jspb.Message.setProto3IntField(this, 5, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.LoginResp.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.LoginResp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.LoginResp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.LoginResp.toObject = function(includeInstance, msg) {
  var obj = {
    maininfo: (f = msg.getMaininfo()) && proto.pb.MasterInfo.toObject(includeInstance, f),
    ingameid: jspb.Message.getFieldWithDefault(msg, 2, 0),
    intablenum: jspb.Message.getFieldWithDefault(msg, 3, 0),
    token: jspb.Message.getFieldWithDefault(msg, 4, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.LoginResp}
 */
proto.pb.LoginResp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.LoginResp;
  return proto.pb.LoginResp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.LoginResp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.LoginResp}
 */
proto.pb.LoginResp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.pb.MasterInfo;
      reader.readMessage(value,proto.pb.MasterInfo.deserializeBinaryFromReader);
      msg.setMaininfo(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setIngameid(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setIntablenum(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setToken(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.LoginResp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.LoginResp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.LoginResp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.LoginResp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getMaininfo();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.pb.MasterInfo.serializeBinaryToWriter
    );
  }
  f = message.getIngameid();
  if (f !== 0) {
    writer.writeInt64(
      2,
      f
    );
  }
  f = message.getIntablenum();
  if (f !== 0) {
    writer.writeInt64(
      3,
      f
    );
  }
  f = message.getToken();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
};


/**
 * optional MasterInfo mainInfo = 1;
 * @return {?proto.pb.MasterInfo}
 */
proto.pb.LoginResp.prototype.getMaininfo = function() {
  return /** @type{?proto.pb.MasterInfo} */ (
    jspb.Message.getWrapperField(this, proto.pb.MasterInfo, 1));
};


/** @param {?proto.pb.MasterInfo|undefined} value */
proto.pb.LoginResp.prototype.setMaininfo = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 */
proto.pb.LoginResp.prototype.clearMaininfo = function() {
  this.setMaininfo(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.pb.LoginResp.prototype.hasMaininfo = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional int64 inGameID = 2;
 * @return {number}
 */
proto.pb.LoginResp.prototype.getIngameid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {number} value */
proto.pb.LoginResp.prototype.setIngameid = function(value) {
  jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional int64 inTableNum = 3;
 * @return {number}
 */
proto.pb.LoginResp.prototype.getIntablenum = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/** @param {number} value */
proto.pb.LoginResp.prototype.setIntablenum = function(value) {
  jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional string token = 4;
 * @return {string}
 */
proto.pb.LoginResp.prototype.getToken = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.pb.LoginResp.prototype.setToken = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.AllopatricResp.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.AllopatricResp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.AllopatricResp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.AllopatricResp.toObject = function(includeInstance, msg) {
  var obj = {
    userid: jspb.Message.getFieldWithDefault(msg, 1, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.AllopatricResp}
 */
proto.pb.AllopatricResp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.AllopatricResp;
  return proto.pb.AllopatricResp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.AllopatricResp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.AllopatricResp}
 */
proto.pb.AllopatricResp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setUserid(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.AllopatricResp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.AllopatricResp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.AllopatricResp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.AllopatricResp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUserid();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
};


/**
 * optional int64 userID = 1;
 * @return {number}
 */
proto.pb.AllopatricResp.prototype.getUserid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.AllopatricResp.prototype.setUserid = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.ReconnectReq.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.ReconnectReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.ReconnectReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.ReconnectReq.toObject = function(includeInstance, msg) {
  var obj = {
    account: jspb.Message.getFieldWithDefault(msg, 1, ""),
    password: jspb.Message.getFieldWithDefault(msg, 2, ""),
    machinecode: jspb.Message.getFieldWithDefault(msg, 3, ""),
    platformid: jspb.Message.getFieldWithDefault(msg, 4, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.ReconnectReq}
 */
proto.pb.ReconnectReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.ReconnectReq;
  return proto.pb.ReconnectReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.ReconnectReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.ReconnectReq}
 */
proto.pb.ReconnectReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setAccount(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setPassword(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setMachinecode(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setPlatformid(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.ReconnectReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.ReconnectReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.ReconnectReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.ReconnectReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getAccount();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getPassword();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getMachinecode();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getPlatformid();
  if (f !== 0) {
    writer.writeInt64(
      4,
      f
    );
  }
};


/**
 * optional string account = 1;
 * @return {string}
 */
proto.pb.ReconnectReq.prototype.getAccount = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.pb.ReconnectReq.prototype.setAccount = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string password = 2;
 * @return {string}
 */
proto.pb.ReconnectReq.prototype.getPassword = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.pb.ReconnectReq.prototype.setPassword = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string machineCode = 3;
 * @return {string}
 */
proto.pb.ReconnectReq.prototype.getMachinecode = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.pb.ReconnectReq.prototype.setMachinecode = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional int64 platformID = 4;
 * @return {number}
 */
proto.pb.ReconnectReq.prototype.getPlatformid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/** @param {number} value */
proto.pb.ReconnectReq.prototype.setPlatformid = function(value) {
  jspb.Message.setProto3IntField(this, 4, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.ReconnectResp.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.ReconnectResp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.ReconnectResp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.ReconnectResp.toObject = function(includeInstance, msg) {
  var obj = {
    maininfo: (f = msg.getMaininfo()) && proto.pb.MasterInfo.toObject(includeInstance, f),
    ingameid: jspb.Message.getFieldWithDefault(msg, 2, 0),
    intablenum: jspb.Message.getFieldWithDefault(msg, 3, 0),
    token: jspb.Message.getFieldWithDefault(msg, 4, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.ReconnectResp}
 */
proto.pb.ReconnectResp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.ReconnectResp;
  return proto.pb.ReconnectResp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.ReconnectResp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.ReconnectResp}
 */
proto.pb.ReconnectResp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.pb.MasterInfo;
      reader.readMessage(value,proto.pb.MasterInfo.deserializeBinaryFromReader);
      msg.setMaininfo(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setIngameid(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setIntablenum(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setToken(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.ReconnectResp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.ReconnectResp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.ReconnectResp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.ReconnectResp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getMaininfo();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.pb.MasterInfo.serializeBinaryToWriter
    );
  }
  f = message.getIngameid();
  if (f !== 0) {
    writer.writeInt64(
      2,
      f
    );
  }
  f = message.getIntablenum();
  if (f !== 0) {
    writer.writeInt64(
      3,
      f
    );
  }
  f = message.getToken();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
};


/**
 * optional MasterInfo mainInfo = 1;
 * @return {?proto.pb.MasterInfo}
 */
proto.pb.ReconnectResp.prototype.getMaininfo = function() {
  return /** @type{?proto.pb.MasterInfo} */ (
    jspb.Message.getWrapperField(this, proto.pb.MasterInfo, 1));
};


/** @param {?proto.pb.MasterInfo|undefined} value */
proto.pb.ReconnectResp.prototype.setMaininfo = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 */
proto.pb.ReconnectResp.prototype.clearMaininfo = function() {
  this.setMaininfo(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.pb.ReconnectResp.prototype.hasMaininfo = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional int64 inGameID = 2;
 * @return {number}
 */
proto.pb.ReconnectResp.prototype.getIngameid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {number} value */
proto.pb.ReconnectResp.prototype.setIngameid = function(value) {
  jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional int64 inTableNum = 3;
 * @return {number}
 */
proto.pb.ReconnectResp.prototype.getIntablenum = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/** @param {number} value */
proto.pb.ReconnectResp.prototype.setIntablenum = function(value) {
  jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional string token = 4;
 * @return {string}
 */
proto.pb.ReconnectResp.prototype.getToken = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.pb.ReconnectResp.prototype.setToken = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.ChooseClassReq.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.ChooseClassReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.ChooseClassReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.ChooseClassReq.toObject = function(includeInstance, msg) {
  var obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, 0),
    tablekey: jspb.Message.getFieldWithDefault(msg, 2, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.ChooseClassReq}
 */
proto.pb.ChooseClassReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.ChooseClassReq;
  return proto.pb.ChooseClassReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.ChooseClassReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.ChooseClassReq}
 */
proto.pb.ChooseClassReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setTablekey(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.ChooseClassReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.ChooseClassReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.ChooseClassReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.ChooseClassReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f !== 0) {
    writer.writeInt32(
      1,
      f
    );
  }
  f = message.getTablekey();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional int32 iD = 1;
 * @return {number}
 */
proto.pb.ChooseClassReq.prototype.getId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.ChooseClassReq.prototype.setId = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional string tableKey = 2;
 * @return {string}
 */
proto.pb.ChooseClassReq.prototype.getTablekey = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.pb.ChooseClassReq.prototype.setTablekey = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.ChooseClassResp.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.ChooseClassResp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.ChooseClassResp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.ChooseClassResp.toObject = function(includeInstance, msg) {
  var obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, 0),
    games: (f = msg.getGames()) && proto.pb.GameList.toObject(includeInstance, f),
    pagenum: jspb.Message.getFieldWithDefault(msg, 3, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.ChooseClassResp}
 */
proto.pb.ChooseClassResp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.ChooseClassResp;
  return proto.pb.ChooseClassResp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.ChooseClassResp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.ChooseClassResp}
 */
proto.pb.ChooseClassResp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setId(value);
      break;
    case 2:
      var value = new proto.pb.GameList;
      reader.readMessage(value,proto.pb.GameList.deserializeBinaryFromReader);
      msg.setGames(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setPagenum(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.ChooseClassResp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.ChooseClassResp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.ChooseClassResp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.ChooseClassResp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f !== 0) {
    writer.writeInt32(
      1,
      f
    );
  }
  f = message.getGames();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.pb.GameList.serializeBinaryToWriter
    );
  }
  f = message.getPagenum();
  if (f !== 0) {
    writer.writeInt32(
      3,
      f
    );
  }
};


/**
 * optional int32 iD = 1;
 * @return {number}
 */
proto.pb.ChooseClassResp.prototype.getId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.ChooseClassResp.prototype.setId = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional GameList games = 2;
 * @return {?proto.pb.GameList}
 */
proto.pb.ChooseClassResp.prototype.getGames = function() {
  return /** @type{?proto.pb.GameList} */ (
    jspb.Message.getWrapperField(this, proto.pb.GameList, 2));
};


/** @param {?proto.pb.GameList|undefined} value */
proto.pb.ChooseClassResp.prototype.setGames = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 */
proto.pb.ChooseClassResp.prototype.clearGames = function() {
  this.setGames(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.pb.ChooseClassResp.prototype.hasGames = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional int32 pageNum = 3;
 * @return {number}
 */
proto.pb.ChooseClassResp.prototype.getPagenum = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/** @param {number} value */
proto.pb.ChooseClassResp.prototype.setPagenum = function(value) {
  jspb.Message.setProto3IntField(this, 3, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.ChooseGameReq.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.ChooseGameReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.ChooseGameReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.ChooseGameReq.toObject = function(includeInstance, msg) {
  var obj = {
    info: (f = msg.getInfo()) && proto.pb.GameInfo.toObject(includeInstance, f),
    pagenum: jspb.Message.getFieldWithDefault(msg, 2, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.ChooseGameReq}
 */
proto.pb.ChooseGameReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.ChooseGameReq;
  return proto.pb.ChooseGameReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.ChooseGameReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.ChooseGameReq}
 */
proto.pb.ChooseGameReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.pb.GameInfo;
      reader.readMessage(value,proto.pb.GameInfo.deserializeBinaryFromReader);
      msg.setInfo(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setPagenum(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.ChooseGameReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.ChooseGameReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.ChooseGameReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.ChooseGameReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getInfo();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.pb.GameInfo.serializeBinaryToWriter
    );
  }
  f = message.getPagenum();
  if (f !== 0) {
    writer.writeInt32(
      2,
      f
    );
  }
};


/**
 * optional GameInfo info = 1;
 * @return {?proto.pb.GameInfo}
 */
proto.pb.ChooseGameReq.prototype.getInfo = function() {
  return /** @type{?proto.pb.GameInfo} */ (
    jspb.Message.getWrapperField(this, proto.pb.GameInfo, 1));
};


/** @param {?proto.pb.GameInfo|undefined} value */
proto.pb.ChooseGameReq.prototype.setInfo = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 */
proto.pb.ChooseGameReq.prototype.clearInfo = function() {
  this.setInfo(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.pb.ChooseGameReq.prototype.hasInfo = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional int32 pageNum = 2;
 * @return {number}
 */
proto.pb.ChooseGameReq.prototype.getPagenum = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {number} value */
proto.pb.ChooseGameReq.prototype.setPagenum = function(value) {
  jspb.Message.setProto3IntField(this, 2, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.ChooseGameResp.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.ChooseGameResp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.ChooseGameResp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.ChooseGameResp.toObject = function(includeInstance, msg) {
  var obj = {
    info: (f = msg.getInfo()) && proto.pb.GameInfo.toObject(includeInstance, f),
    pagenum: jspb.Message.getFieldWithDefault(msg, 2, 0),
    tables: (f = msg.getTables()) && proto.pb.TableList.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.ChooseGameResp}
 */
proto.pb.ChooseGameResp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.ChooseGameResp;
  return proto.pb.ChooseGameResp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.ChooseGameResp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.ChooseGameResp}
 */
proto.pb.ChooseGameResp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.pb.GameInfo;
      reader.readMessage(value,proto.pb.GameInfo.deserializeBinaryFromReader);
      msg.setInfo(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setPagenum(value);
      break;
    case 3:
      var value = new proto.pb.TableList;
      reader.readMessage(value,proto.pb.TableList.deserializeBinaryFromReader);
      msg.setTables(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.ChooseGameResp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.ChooseGameResp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.ChooseGameResp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.ChooseGameResp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getInfo();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.pb.GameInfo.serializeBinaryToWriter
    );
  }
  f = message.getPagenum();
  if (f !== 0) {
    writer.writeInt32(
      2,
      f
    );
  }
  f = message.getTables();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      proto.pb.TableList.serializeBinaryToWriter
    );
  }
};


/**
 * optional GameInfo info = 1;
 * @return {?proto.pb.GameInfo}
 */
proto.pb.ChooseGameResp.prototype.getInfo = function() {
  return /** @type{?proto.pb.GameInfo} */ (
    jspb.Message.getWrapperField(this, proto.pb.GameInfo, 1));
};


/** @param {?proto.pb.GameInfo|undefined} value */
proto.pb.ChooseGameResp.prototype.setInfo = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 */
proto.pb.ChooseGameResp.prototype.clearInfo = function() {
  this.setInfo(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.pb.ChooseGameResp.prototype.hasInfo = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional int32 pageNum = 2;
 * @return {number}
 */
proto.pb.ChooseGameResp.prototype.getPagenum = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {number} value */
proto.pb.ChooseGameResp.prototype.setPagenum = function(value) {
  jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional TableList tables = 3;
 * @return {?proto.pb.TableList}
 */
proto.pb.ChooseGameResp.prototype.getTables = function() {
  return /** @type{?proto.pb.TableList} */ (
    jspb.Message.getWrapperField(this, proto.pb.TableList, 3));
};


/** @param {?proto.pb.TableList|undefined} value */
proto.pb.ChooseGameResp.prototype.setTables = function(value) {
  jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 */
proto.pb.ChooseGameResp.prototype.clearTables = function() {
  this.setTables(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.pb.ChooseGameResp.prototype.hasTables = function() {
  return jspb.Message.getField(this, 3) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.SettingTableReq.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.SettingTableReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.SettingTableReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.SettingTableReq.toObject = function(includeInstance, msg) {
  var obj = {
    ginfo: (f = msg.getGinfo()) && proto.pb.GameInfo.toObject(includeInstance, f),
    tinfo: (f = msg.getTinfo()) && proto.pb.TableInfo.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.SettingTableReq}
 */
proto.pb.SettingTableReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.SettingTableReq;
  return proto.pb.SettingTableReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.SettingTableReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.SettingTableReq}
 */
proto.pb.SettingTableReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.pb.GameInfo;
      reader.readMessage(value,proto.pb.GameInfo.deserializeBinaryFromReader);
      msg.setGinfo(value);
      break;
    case 2:
      var value = new proto.pb.TableInfo;
      reader.readMessage(value,proto.pb.TableInfo.deserializeBinaryFromReader);
      msg.setTinfo(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.SettingTableReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.SettingTableReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.SettingTableReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.SettingTableReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getGinfo();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.pb.GameInfo.serializeBinaryToWriter
    );
  }
  f = message.getTinfo();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.pb.TableInfo.serializeBinaryToWriter
    );
  }
};


/**
 * optional GameInfo gInfo = 1;
 * @return {?proto.pb.GameInfo}
 */
proto.pb.SettingTableReq.prototype.getGinfo = function() {
  return /** @type{?proto.pb.GameInfo} */ (
    jspb.Message.getWrapperField(this, proto.pb.GameInfo, 1));
};


/** @param {?proto.pb.GameInfo|undefined} value */
proto.pb.SettingTableReq.prototype.setGinfo = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 */
proto.pb.SettingTableReq.prototype.clearGinfo = function() {
  this.setGinfo(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.pb.SettingTableReq.prototype.hasGinfo = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional TableInfo tInfo = 2;
 * @return {?proto.pb.TableInfo}
 */
proto.pb.SettingTableReq.prototype.getTinfo = function() {
  return /** @type{?proto.pb.TableInfo} */ (
    jspb.Message.getWrapperField(this, proto.pb.TableInfo, 2));
};


/** @param {?proto.pb.TableInfo|undefined} value */
proto.pb.SettingTableReq.prototype.setTinfo = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 */
proto.pb.SettingTableReq.prototype.clearTinfo = function() {
  this.setTinfo(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.pb.SettingTableReq.prototype.hasTinfo = function() {
  return jspb.Message.getField(this, 2) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.SettingTableResp.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.SettingTableResp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.SettingTableResp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.SettingTableResp.toObject = function(includeInstance, msg) {
  var obj = {
    item: (f = msg.getItem()) && proto.pb.TableItem.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.SettingTableResp}
 */
proto.pb.SettingTableResp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.SettingTableResp;
  return proto.pb.SettingTableResp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.SettingTableResp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.SettingTableResp}
 */
proto.pb.SettingTableResp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.pb.TableItem;
      reader.readMessage(value,proto.pb.TableItem.deserializeBinaryFromReader);
      msg.setItem(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.SettingTableResp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.SettingTableResp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.SettingTableResp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.SettingTableResp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getItem();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.pb.TableItem.serializeBinaryToWriter
    );
  }
};


/**
 * optional TableItem item = 1;
 * @return {?proto.pb.TableItem}
 */
proto.pb.SettingTableResp.prototype.getItem = function() {
  return /** @type{?proto.pb.TableItem} */ (
    jspb.Message.getWrapperField(this, proto.pb.TableItem, 1));
};


/** @param {?proto.pb.TableItem|undefined} value */
proto.pb.SettingTableResp.prototype.setItem = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 */
proto.pb.SettingTableResp.prototype.clearItem = function() {
  this.setItem(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.pb.SettingTableResp.prototype.hasItem = function() {
  return jspb.Message.getField(this, 1) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.CheckInReq.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.CheckInReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.CheckInReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.CheckInReq.toObject = function(includeInstance, msg) {
  var obj = {
    userid: jspb.Message.getFieldWithDefault(msg, 1, 0),
    remark: jspb.Message.getFieldWithDefault(msg, 2, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.CheckInReq}
 */
proto.pb.CheckInReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.CheckInReq;
  return proto.pb.CheckInReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.CheckInReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.CheckInReq}
 */
proto.pb.CheckInReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setUserid(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setRemark(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.CheckInReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.CheckInReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.CheckInReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.CheckInReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUserid();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getRemark();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional int64 userID = 1;
 * @return {number}
 */
proto.pb.CheckInReq.prototype.getUserid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.CheckInReq.prototype.setUserid = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional string remark = 2;
 * @return {string}
 */
proto.pb.CheckInReq.prototype.getRemark = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.pb.CheckInReq.prototype.setRemark = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.CheckInResp.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.CheckInResp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.CheckInResp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.CheckInResp.toObject = function(includeInstance, msg) {
  var obj = {
    userid: jspb.Message.getFieldWithDefault(msg, 1, 0),
    remark: jspb.Message.getFieldWithDefault(msg, 2, ""),
    timestamp: jspb.Message.getFieldWithDefault(msg, 3, 0),
    awardlist: (f = msg.getAwardlist()) && proto.pb.GoodsList.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.CheckInResp}
 */
proto.pb.CheckInResp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.CheckInResp;
  return proto.pb.CheckInResp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.CheckInResp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.CheckInResp}
 */
proto.pb.CheckInResp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setUserid(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setRemark(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setTimestamp(value);
      break;
    case 4:
      var value = new proto.pb.GoodsList;
      reader.readMessage(value,proto.pb.GoodsList.deserializeBinaryFromReader);
      msg.setAwardlist(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.CheckInResp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.CheckInResp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.CheckInResp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.CheckInResp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUserid();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getRemark();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getTimestamp();
  if (f !== 0) {
    writer.writeInt64(
      3,
      f
    );
  }
  f = message.getAwardlist();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      proto.pb.GoodsList.serializeBinaryToWriter
    );
  }
};


/**
 * optional int64 userID = 1;
 * @return {number}
 */
proto.pb.CheckInResp.prototype.getUserid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.CheckInResp.prototype.setUserid = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional string remark = 2;
 * @return {string}
 */
proto.pb.CheckInResp.prototype.getRemark = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.pb.CheckInResp.prototype.setRemark = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional int64 timestamp = 3;
 * @return {number}
 */
proto.pb.CheckInResp.prototype.getTimestamp = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/** @param {number} value */
proto.pb.CheckInResp.prototype.setTimestamp = function(value) {
  jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional GoodsList awardList = 4;
 * @return {?proto.pb.GoodsList}
 */
proto.pb.CheckInResp.prototype.getAwardlist = function() {
  return /** @type{?proto.pb.GoodsList} */ (
    jspb.Message.getWrapperField(this, proto.pb.GoodsList, 4));
};


/** @param {?proto.pb.GoodsList|undefined} value */
proto.pb.CheckInResp.prototype.setAwardlist = function(value) {
  jspb.Message.setWrapperField(this, 4, value);
};


/**
 * Clears the message field making it undefined.
 */
proto.pb.CheckInResp.prototype.clearAwardlist = function() {
  this.setAwardlist(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.pb.CheckInResp.prototype.hasAwardlist = function() {
  return jspb.Message.getField(this, 4) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.GetCheckInReq.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.GetCheckInReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.GetCheckInReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.GetCheckInReq.toObject = function(includeInstance, msg) {
  var obj = {

  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.GetCheckInReq}
 */
proto.pb.GetCheckInReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.GetCheckInReq;
  return proto.pb.GetCheckInReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.GetCheckInReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.GetCheckInReq}
 */
proto.pb.GetCheckInReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.GetCheckInReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.GetCheckInReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.GetCheckInReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.GetCheckInReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.pb.GetCheckInResp.repeatedFields_ = [2];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.GetCheckInResp.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.GetCheckInResp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.GetCheckInResp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.GetCheckInResp.toObject = function(includeInstance, msg) {
  var obj = {
    userid: jspb.Message.getFieldWithDefault(msg, 1, 0),
    allcheckinList: jspb.Message.toObjectList(msg.getAllcheckinList(),
    proto.pb.CheckInResp.toObject, includeInstance),
    pagenum: jspb.Message.getFieldWithDefault(msg, 3, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.GetCheckInResp}
 */
proto.pb.GetCheckInResp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.GetCheckInResp;
  return proto.pb.GetCheckInResp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.GetCheckInResp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.GetCheckInResp}
 */
proto.pb.GetCheckInResp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setUserid(value);
      break;
    case 2:
      var value = new proto.pb.CheckInResp;
      reader.readMessage(value,proto.pb.CheckInResp.deserializeBinaryFromReader);
      msg.addAllcheckin(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setPagenum(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.GetCheckInResp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.GetCheckInResp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.GetCheckInResp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.GetCheckInResp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUserid();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getAllcheckinList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      2,
      f,
      proto.pb.CheckInResp.serializeBinaryToWriter
    );
  }
  f = message.getPagenum();
  if (f !== 0) {
    writer.writeInt32(
      3,
      f
    );
  }
};


/**
 * optional int64 userID = 1;
 * @return {number}
 */
proto.pb.GetCheckInResp.prototype.getUserid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.GetCheckInResp.prototype.setUserid = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * repeated CheckInResp allCheckin = 2;
 * @return {!Array<!proto.pb.CheckInResp>}
 */
proto.pb.GetCheckInResp.prototype.getAllcheckinList = function() {
  return /** @type{!Array<!proto.pb.CheckInResp>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.pb.CheckInResp, 2));
};


/** @param {!Array<!proto.pb.CheckInResp>} value */
proto.pb.GetCheckInResp.prototype.setAllcheckinList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 2, value);
};


/**
 * @param {!proto.pb.CheckInResp=} opt_value
 * @param {number=} opt_index
 * @return {!proto.pb.CheckInResp}
 */
proto.pb.GetCheckInResp.prototype.addAllcheckin = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 2, opt_value, proto.pb.CheckInResp, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 */
proto.pb.GetCheckInResp.prototype.clearAllcheckinList = function() {
  this.setAllcheckinList([]);
};


/**
 * optional int32 pageNum = 3;
 * @return {number}
 */
proto.pb.GetCheckInResp.prototype.getPagenum = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/** @param {number} value */
proto.pb.GetCheckInResp.prototype.setPagenum = function(value) {
  jspb.Message.setProto3IntField(this, 3, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.DrawHeroReq.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.DrawHeroReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.DrawHeroReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.DrawHeroReq.toObject = function(includeInstance, msg) {
  var obj = {
    amount: jspb.Message.getFieldWithDefault(msg, 1, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.DrawHeroReq}
 */
proto.pb.DrawHeroReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.DrawHeroReq;
  return proto.pb.DrawHeroReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.DrawHeroReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.DrawHeroReq}
 */
proto.pb.DrawHeroReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setAmount(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.DrawHeroReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.DrawHeroReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.DrawHeroReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.DrawHeroReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getAmount();
  if (f !== 0) {
    writer.writeInt32(
      1,
      f
    );
  }
};


/**
 * optional int32 amount = 1;
 * @return {number}
 */
proto.pb.DrawHeroReq.prototype.getAmount = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.DrawHeroReq.prototype.setAmount = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.pb.DrawHeroResp.repeatedFields_ = [2];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.DrawHeroResp.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.DrawHeroResp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.DrawHeroResp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.DrawHeroResp.toObject = function(includeInstance, msg) {
  var obj = {
    userid: jspb.Message.getFieldWithDefault(msg, 1, 0),
    herolistList: jspb.Message.toObjectList(msg.getHerolistList(),
    proto.pb.HeroInfo.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.DrawHeroResp}
 */
proto.pb.DrawHeroResp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.DrawHeroResp;
  return proto.pb.DrawHeroResp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.DrawHeroResp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.DrawHeroResp}
 */
proto.pb.DrawHeroResp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setUserid(value);
      break;
    case 2:
      var value = new proto.pb.HeroInfo;
      reader.readMessage(value,proto.pb.HeroInfo.deserializeBinaryFromReader);
      msg.addHerolist(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.DrawHeroResp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.DrawHeroResp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.DrawHeroResp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.DrawHeroResp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUserid();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getHerolistList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      2,
      f,
      proto.pb.HeroInfo.serializeBinaryToWriter
    );
  }
};


/**
 * optional int64 userID = 1;
 * @return {number}
 */
proto.pb.DrawHeroResp.prototype.getUserid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.DrawHeroResp.prototype.setUserid = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * repeated HeroInfo heroList = 2;
 * @return {!Array<!proto.pb.HeroInfo>}
 */
proto.pb.DrawHeroResp.prototype.getHerolistList = function() {
  return /** @type{!Array<!proto.pb.HeroInfo>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.pb.HeroInfo, 2));
};


/** @param {!Array<!proto.pb.HeroInfo>} value */
proto.pb.DrawHeroResp.prototype.setHerolistList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 2, value);
};


/**
 * @param {!proto.pb.HeroInfo=} opt_value
 * @param {number=} opt_index
 * @return {!proto.pb.HeroInfo}
 */
proto.pb.DrawHeroResp.prototype.addHerolist = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 2, opt_value, proto.pb.HeroInfo, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 */
proto.pb.DrawHeroResp.prototype.clearHerolistList = function() {
  this.setHerolistList([]);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.GetMyHeroReq.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.GetMyHeroReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.GetMyHeroReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.GetMyHeroReq.toObject = function(includeInstance, msg) {
  var obj = {

  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.GetMyHeroReq}
 */
proto.pb.GetMyHeroReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.GetMyHeroReq;
  return proto.pb.GetMyHeroReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.GetMyHeroReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.GetMyHeroReq}
 */
proto.pb.GetMyHeroReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.GetMyHeroReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.GetMyHeroReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.GetMyHeroReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.GetMyHeroReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.pb.GetMyHeroResp.repeatedFields_ = [2];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.GetMyHeroResp.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.GetMyHeroResp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.GetMyHeroResp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.GetMyHeroResp.toObject = function(includeInstance, msg) {
  var obj = {
    userid: jspb.Message.getFieldWithDefault(msg, 1, 0),
    herolistList: jspb.Message.toObjectList(msg.getHerolistList(),
    proto.pb.HeroInfo.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.GetMyHeroResp}
 */
proto.pb.GetMyHeroResp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.GetMyHeroResp;
  return proto.pb.GetMyHeroResp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.GetMyHeroResp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.GetMyHeroResp}
 */
proto.pb.GetMyHeroResp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setUserid(value);
      break;
    case 2:
      var value = new proto.pb.HeroInfo;
      reader.readMessage(value,proto.pb.HeroInfo.deserializeBinaryFromReader);
      msg.addHerolist(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.GetMyHeroResp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.GetMyHeroResp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.GetMyHeroResp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.GetMyHeroResp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUserid();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getHerolistList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      2,
      f,
      proto.pb.HeroInfo.serializeBinaryToWriter
    );
  }
};


/**
 * optional int64 userID = 1;
 * @return {number}
 */
proto.pb.GetMyHeroResp.prototype.getUserid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.GetMyHeroResp.prototype.setUserid = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * repeated HeroInfo heroList = 2;
 * @return {!Array<!proto.pb.HeroInfo>}
 */
proto.pb.GetMyHeroResp.prototype.getHerolistList = function() {
  return /** @type{!Array<!proto.pb.HeroInfo>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.pb.HeroInfo, 2));
};


/** @param {!Array<!proto.pb.HeroInfo>} value */
proto.pb.GetMyHeroResp.prototype.setHerolistList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 2, value);
};


/**
 * @param {!proto.pb.HeroInfo=} opt_value
 * @param {number=} opt_index
 * @return {!proto.pb.HeroInfo}
 */
proto.pb.GetMyHeroResp.prototype.addHerolist = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 2, opt_value, proto.pb.HeroInfo, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 */
proto.pb.GetMyHeroResp.prototype.clearHerolistList = function() {
  this.setHerolistList([]);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.ChooseHeroReq.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.ChooseHeroReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.ChooseHeroReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.ChooseHeroReq.toObject = function(includeInstance, msg) {
  var obj = {
    position: jspb.Message.getFieldWithDefault(msg, 1, 0),
    heroid: jspb.Message.getFieldWithDefault(msg, 2, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.ChooseHeroReq}
 */
proto.pb.ChooseHeroReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.ChooseHeroReq;
  return proto.pb.ChooseHeroReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.ChooseHeroReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.ChooseHeroReq}
 */
proto.pb.ChooseHeroReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setPosition(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setHeroid(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.ChooseHeroReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.ChooseHeroReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.ChooseHeroReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.ChooseHeroReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getPosition();
  if (f !== 0) {
    writer.writeInt32(
      1,
      f
    );
  }
  f = message.getHeroid();
  if (f !== 0) {
    writer.writeInt64(
      2,
      f
    );
  }
};


/**
 * optional int32 position = 1;
 * @return {number}
 */
proto.pb.ChooseHeroReq.prototype.getPosition = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.ChooseHeroReq.prototype.setPosition = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional int64 heroID = 2;
 * @return {number}
 */
proto.pb.ChooseHeroReq.prototype.getHeroid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {number} value */
proto.pb.ChooseHeroReq.prototype.setHeroid = function(value) {
  jspb.Message.setProto3IntField(this, 2, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.ChooseHeroResp.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.ChooseHeroResp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.ChooseHeroResp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.ChooseHeroResp.toObject = function(includeInstance, msg) {
  var obj = {
    position: jspb.Message.getFieldWithDefault(msg, 1, 0),
    hero: (f = msg.getHero()) && proto.pb.HeroInfo.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.ChooseHeroResp}
 */
proto.pb.ChooseHeroResp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.ChooseHeroResp;
  return proto.pb.ChooseHeroResp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.ChooseHeroResp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.ChooseHeroResp}
 */
proto.pb.ChooseHeroResp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setPosition(value);
      break;
    case 2:
      var value = new proto.pb.HeroInfo;
      reader.readMessage(value,proto.pb.HeroInfo.deserializeBinaryFromReader);
      msg.setHero(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.ChooseHeroResp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.ChooseHeroResp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.ChooseHeroResp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.ChooseHeroResp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getPosition();
  if (f !== 0) {
    writer.writeInt32(
      1,
      f
    );
  }
  f = message.getHero();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.pb.HeroInfo.serializeBinaryToWriter
    );
  }
};


/**
 * optional int32 position = 1;
 * @return {number}
 */
proto.pb.ChooseHeroResp.prototype.getPosition = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.ChooseHeroResp.prototype.setPosition = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional HeroInfo hero = 2;
 * @return {?proto.pb.HeroInfo}
 */
proto.pb.ChooseHeroResp.prototype.getHero = function() {
  return /** @type{?proto.pb.HeroInfo} */ (
    jspb.Message.getWrapperField(this, proto.pb.HeroInfo, 2));
};


/** @param {?proto.pb.HeroInfo|undefined} value */
proto.pb.ChooseHeroResp.prototype.setHero = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 */
proto.pb.ChooseHeroResp.prototype.clearHero = function() {
  this.setHero(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.pb.ChooseHeroResp.prototype.hasHero = function() {
  return jspb.Message.getField(this, 2) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.DownHeroReq.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.DownHeroReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.DownHeroReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.DownHeroReq.toObject = function(includeInstance, msg) {
  var obj = {
    position: jspb.Message.getFieldWithDefault(msg, 1, 0),
    heroid: jspb.Message.getFieldWithDefault(msg, 2, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.DownHeroReq}
 */
proto.pb.DownHeroReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.DownHeroReq;
  return proto.pb.DownHeroReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.DownHeroReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.DownHeroReq}
 */
proto.pb.DownHeroReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setPosition(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setHeroid(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.DownHeroReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.DownHeroReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.DownHeroReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.DownHeroReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getPosition();
  if (f !== 0) {
    writer.writeInt32(
      1,
      f
    );
  }
  f = message.getHeroid();
  if (f !== 0) {
    writer.writeInt64(
      2,
      f
    );
  }
};


/**
 * optional int32 position = 1;
 * @return {number}
 */
proto.pb.DownHeroReq.prototype.getPosition = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.DownHeroReq.prototype.setPosition = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional int64 heroID = 2;
 * @return {number}
 */
proto.pb.DownHeroReq.prototype.getHeroid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {number} value */
proto.pb.DownHeroReq.prototype.setHeroid = function(value) {
  jspb.Message.setProto3IntField(this, 2, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.DownHeroReqResp.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.DownHeroReqResp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.DownHeroReqResp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.DownHeroReqResp.toObject = function(includeInstance, msg) {
  var obj = {
    position: jspb.Message.getFieldWithDefault(msg, 1, 0),
    heroid: jspb.Message.getFieldWithDefault(msg, 2, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.DownHeroReqResp}
 */
proto.pb.DownHeroReqResp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.DownHeroReqResp;
  return proto.pb.DownHeroReqResp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.DownHeroReqResp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.DownHeroReqResp}
 */
proto.pb.DownHeroReqResp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setPosition(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setHeroid(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.DownHeroReqResp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.DownHeroReqResp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.DownHeroReqResp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.DownHeroReqResp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getPosition();
  if (f !== 0) {
    writer.writeInt32(
      1,
      f
    );
  }
  f = message.getHeroid();
  if (f !== 0) {
    writer.writeInt64(
      2,
      f
    );
  }
};


/**
 * optional int32 position = 1;
 * @return {number}
 */
proto.pb.DownHeroReqResp.prototype.getPosition = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.DownHeroReqResp.prototype.setPosition = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional int64 heroID = 2;
 * @return {number}
 */
proto.pb.DownHeroReqResp.prototype.getHeroid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {number} value */
proto.pb.DownHeroReqResp.prototype.setHeroid = function(value) {
  jspb.Message.setProto3IntField(this, 2, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.GetAllHeroReq.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.GetAllHeroReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.GetAllHeroReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.GetAllHeroReq.toObject = function(includeInstance, msg) {
  var obj = {

  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.GetAllHeroReq}
 */
proto.pb.GetAllHeroReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.GetAllHeroReq;
  return proto.pb.GetAllHeroReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.GetAllHeroReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.GetAllHeroReq}
 */
proto.pb.GetAllHeroReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.GetAllHeroReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.GetAllHeroReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.GetAllHeroReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.GetAllHeroReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.pb.GetAllHeroResp.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.GetAllHeroResp.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.GetAllHeroResp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.GetAllHeroResp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.GetAllHeroResp.toObject = function(includeInstance, msg) {
  var obj = {
    herolistList: jspb.Message.toObjectList(msg.getHerolistList(),
    proto.pb.HeroInfo.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.GetAllHeroResp}
 */
proto.pb.GetAllHeroResp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.GetAllHeroResp;
  return proto.pb.GetAllHeroResp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.GetAllHeroResp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.GetAllHeroResp}
 */
proto.pb.GetAllHeroResp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.pb.HeroInfo;
      reader.readMessage(value,proto.pb.HeroInfo.deserializeBinaryFromReader);
      msg.addHerolist(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.GetAllHeroResp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.GetAllHeroResp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.GetAllHeroResp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.GetAllHeroResp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getHerolistList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.pb.HeroInfo.serializeBinaryToWriter
    );
  }
};


/**
 * repeated HeroInfo heroList = 1;
 * @return {!Array<!proto.pb.HeroInfo>}
 */
proto.pb.GetAllHeroResp.prototype.getHerolistList = function() {
  return /** @type{!Array<!proto.pb.HeroInfo>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.pb.HeroInfo, 1));
};


/** @param {!Array<!proto.pb.HeroInfo>} value */
proto.pb.GetAllHeroResp.prototype.setHerolistList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.pb.HeroInfo=} opt_value
 * @param {number=} opt_index
 * @return {!proto.pb.HeroInfo}
 */
proto.pb.GetAllHeroResp.prototype.addHerolist = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.pb.HeroInfo, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 */
proto.pb.GetAllHeroResp.prototype.clearHerolistList = function() {
  this.setHerolistList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.pb.CheckHeroReq.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.CheckHeroReq.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.CheckHeroReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.CheckHeroReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.CheckHeroReq.toObject = function(includeInstance, msg) {
  var obj = {
    heroidsList: jspb.Message.getRepeatedField(msg, 1),
    name: jspb.Message.getFieldWithDefault(msg, 2, ""),
    sex: jspb.Message.getFieldWithDefault(msg, 3, 0),
    country: jspb.Message.getFieldWithDefault(msg, 4, ""),
    faction: jspb.Message.getFieldWithDefault(msg, 5, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.CheckHeroReq}
 */
proto.pb.CheckHeroReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.CheckHeroReq;
  return proto.pb.CheckHeroReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.CheckHeroReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.CheckHeroReq}
 */
proto.pb.CheckHeroReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!Array<number>} */ (reader.readPackedInt32());
      msg.setHeroidsList(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setSex(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setCountry(value);
      break;
    case 5:
      var value = /** @type {!proto.pb.HeroType} */ (reader.readEnum());
      msg.setFaction(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.CheckHeroReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.CheckHeroReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.CheckHeroReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.CheckHeroReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getHeroidsList();
  if (f.length > 0) {
    writer.writePackedInt32(
      1,
      f
    );
  }
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getSex();
  if (f !== 0) {
    writer.writeInt32(
      3,
      f
    );
  }
  f = message.getCountry();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getFaction();
  if (f !== 0.0) {
    writer.writeEnum(
      5,
      f
    );
  }
};


/**
 * repeated int32 heroIDs = 1;
 * @return {!Array<number>}
 */
proto.pb.CheckHeroReq.prototype.getHeroidsList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 1));
};


/** @param {!Array<number>} value */
proto.pb.CheckHeroReq.prototype.setHeroidsList = function(value) {
  jspb.Message.setField(this, 1, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 */
proto.pb.CheckHeroReq.prototype.addHeroids = function(value, opt_index) {
  jspb.Message.addToRepeatedField(this, 1, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 */
proto.pb.CheckHeroReq.prototype.clearHeroidsList = function() {
  this.setHeroidsList([]);
};


/**
 * optional string name = 2;
 * @return {string}
 */
proto.pb.CheckHeroReq.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.pb.CheckHeroReq.prototype.setName = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional int32 sex = 3;
 * @return {number}
 */
proto.pb.CheckHeroReq.prototype.getSex = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/** @param {number} value */
proto.pb.CheckHeroReq.prototype.setSex = function(value) {
  jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional string country = 4;
 * @return {string}
 */
proto.pb.CheckHeroReq.prototype.getCountry = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.pb.CheckHeroReq.prototype.setCountry = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional HeroType faction = 5;
 * @return {!proto.pb.HeroType}
 */
proto.pb.CheckHeroReq.prototype.getFaction = function() {
  return /** @type {!proto.pb.HeroType} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/** @param {!proto.pb.HeroType} value */
proto.pb.CheckHeroReq.prototype.setFaction = function(value) {
  jspb.Message.setProto3EnumField(this, 5, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.pb.CheckHeroResp.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.CheckHeroResp.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.CheckHeroResp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.CheckHeroResp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.CheckHeroResp.toObject = function(includeInstance, msg) {
  var obj = {
    herolistList: jspb.Message.toObjectList(msg.getHerolistList(),
    proto.pb.HeroInfo.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.CheckHeroResp}
 */
proto.pb.CheckHeroResp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.CheckHeroResp;
  return proto.pb.CheckHeroResp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.CheckHeroResp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.CheckHeroResp}
 */
proto.pb.CheckHeroResp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.pb.HeroInfo;
      reader.readMessage(value,proto.pb.HeroInfo.deserializeBinaryFromReader);
      msg.addHerolist(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.CheckHeroResp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.CheckHeroResp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.CheckHeroResp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.CheckHeroResp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getHerolistList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.pb.HeroInfo.serializeBinaryToWriter
    );
  }
};


/**
 * repeated HeroInfo heroList = 1;
 * @return {!Array<!proto.pb.HeroInfo>}
 */
proto.pb.CheckHeroResp.prototype.getHerolistList = function() {
  return /** @type{!Array<!proto.pb.HeroInfo>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.pb.HeroInfo, 1));
};


/** @param {!Array<!proto.pb.HeroInfo>} value */
proto.pb.CheckHeroResp.prototype.setHerolistList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.pb.HeroInfo=} opt_value
 * @param {number=} opt_index
 * @return {!proto.pb.HeroInfo}
 */
proto.pb.CheckHeroResp.prototype.addHerolist = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.pb.HeroInfo, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 */
proto.pb.CheckHeroResp.prototype.clearHerolistList = function() {
  this.setHerolistList([]);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.RechargeReq.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.RechargeReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.RechargeReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.RechargeReq.toObject = function(includeInstance, msg) {
  var obj = {
    userid: jspb.Message.getFieldWithDefault(msg, 1, 0),
    byid: jspb.Message.getFieldWithDefault(msg, 2, 0),
    payment: jspb.Message.getFieldWithDefault(msg, 3, 0),
    method: jspb.Message.getFieldWithDefault(msg, 4, 0),
    pb_switch: jspb.Message.getFieldWithDefault(msg, 5, 0),
    reason: jspb.Message.getFieldWithDefault(msg, 6, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.RechargeReq}
 */
proto.pb.RechargeReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.RechargeReq;
  return proto.pb.RechargeReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.RechargeReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.RechargeReq}
 */
proto.pb.RechargeReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setUserid(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setByid(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setPayment(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setMethod(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setSwitch(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setReason(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.RechargeReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.RechargeReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.RechargeReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.RechargeReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUserid();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getByid();
  if (f !== 0) {
    writer.writeInt64(
      2,
      f
    );
  }
  f = message.getPayment();
  if (f !== 0) {
    writer.writeInt64(
      3,
      f
    );
  }
  f = message.getMethod();
  if (f !== 0) {
    writer.writeInt32(
      4,
      f
    );
  }
  f = message.getSwitch();
  if (f !== 0) {
    writer.writeInt32(
      5,
      f
    );
  }
  f = message.getReason();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
};


/**
 * optional int64 userID = 1;
 * @return {number}
 */
proto.pb.RechargeReq.prototype.getUserid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.RechargeReq.prototype.setUserid = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional int64 byiD = 2;
 * @return {number}
 */
proto.pb.RechargeReq.prototype.getByid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {number} value */
proto.pb.RechargeReq.prototype.setByid = function(value) {
  jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional int64 payment = 3;
 * @return {number}
 */
proto.pb.RechargeReq.prototype.getPayment = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/** @param {number} value */
proto.pb.RechargeReq.prototype.setPayment = function(value) {
  jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional int32 method = 4;
 * @return {number}
 */
proto.pb.RechargeReq.prototype.getMethod = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/** @param {number} value */
proto.pb.RechargeReq.prototype.setMethod = function(value) {
  jspb.Message.setProto3IntField(this, 4, value);
};


/**
 * optional int32 switch = 5;
 * @return {number}
 */
proto.pb.RechargeReq.prototype.getSwitch = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/** @param {number} value */
proto.pb.RechargeReq.prototype.setSwitch = function(value) {
  jspb.Message.setProto3IntField(this, 5, value);
};


/**
 * optional string reason = 6;
 * @return {string}
 */
proto.pb.RechargeReq.prototype.getReason = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/** @param {string} value */
proto.pb.RechargeReq.prototype.setReason = function(value) {
  jspb.Message.setProto3StringField(this, 6, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.RechargeResp.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.RechargeResp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.RechargeResp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.RechargeResp.toObject = function(includeInstance, msg) {
  var obj = {
    userid: jspb.Message.getFieldWithDefault(msg, 1, 0),
    byid: jspb.Message.getFieldWithDefault(msg, 2, 0),
    premoney: jspb.Message.getFieldWithDefault(msg, 3, 0),
    payment: jspb.Message.getFieldWithDefault(msg, 4, 0),
    money: jspb.Message.getFieldWithDefault(msg, 5, 0),
    yuanbao: jspb.Message.getFieldWithDefault(msg, 6, 0),
    coin: jspb.Message.getFieldWithDefault(msg, 7, 0),
    method: jspb.Message.getFieldWithDefault(msg, 8, 0),
    issuccess: jspb.Message.getFieldWithDefault(msg, 9, false),
    order: jspb.Message.getFieldWithDefault(msg, 10, ""),
    timestamp: jspb.Message.getFieldWithDefault(msg, 11, 0),
    reason: jspb.Message.getFieldWithDefault(msg, 12, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.RechargeResp}
 */
proto.pb.RechargeResp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.RechargeResp;
  return proto.pb.RechargeResp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.RechargeResp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.RechargeResp}
 */
proto.pb.RechargeResp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setUserid(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setByid(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setPremoney(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setPayment(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setMoney(value);
      break;
    case 6:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setYuanbao(value);
      break;
    case 7:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setCoin(value);
      break;
    case 8:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setMethod(value);
      break;
    case 9:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setIssuccess(value);
      break;
    case 10:
      var value = /** @type {string} */ (reader.readString());
      msg.setOrder(value);
      break;
    case 11:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setTimestamp(value);
      break;
    case 12:
      var value = /** @type {string} */ (reader.readString());
      msg.setReason(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.RechargeResp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.RechargeResp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.RechargeResp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.RechargeResp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUserid();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getByid();
  if (f !== 0) {
    writer.writeInt64(
      2,
      f
    );
  }
  f = message.getPremoney();
  if (f !== 0) {
    writer.writeInt64(
      3,
      f
    );
  }
  f = message.getPayment();
  if (f !== 0) {
    writer.writeInt64(
      4,
      f
    );
  }
  f = message.getMoney();
  if (f !== 0) {
    writer.writeInt64(
      5,
      f
    );
  }
  f = message.getYuanbao();
  if (f !== 0) {
    writer.writeInt64(
      6,
      f
    );
  }
  f = message.getCoin();
  if (f !== 0) {
    writer.writeInt64(
      7,
      f
    );
  }
  f = message.getMethod();
  if (f !== 0) {
    writer.writeInt32(
      8,
      f
    );
  }
  f = message.getIssuccess();
  if (f) {
    writer.writeBool(
      9,
      f
    );
  }
  f = message.getOrder();
  if (f.length > 0) {
    writer.writeString(
      10,
      f
    );
  }
  f = message.getTimestamp();
  if (f !== 0) {
    writer.writeInt64(
      11,
      f
    );
  }
  f = message.getReason();
  if (f.length > 0) {
    writer.writeString(
      12,
      f
    );
  }
};


/**
 * optional int64 userID = 1;
 * @return {number}
 */
proto.pb.RechargeResp.prototype.getUserid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.RechargeResp.prototype.setUserid = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional int64 byiD = 2;
 * @return {number}
 */
proto.pb.RechargeResp.prototype.getByid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {number} value */
proto.pb.RechargeResp.prototype.setByid = function(value) {
  jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional int64 preMoney = 3;
 * @return {number}
 */
proto.pb.RechargeResp.prototype.getPremoney = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/** @param {number} value */
proto.pb.RechargeResp.prototype.setPremoney = function(value) {
  jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional int64 payment = 4;
 * @return {number}
 */
proto.pb.RechargeResp.prototype.getPayment = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/** @param {number} value */
proto.pb.RechargeResp.prototype.setPayment = function(value) {
  jspb.Message.setProto3IntField(this, 4, value);
};


/**
 * optional int64 money = 5;
 * @return {number}
 */
proto.pb.RechargeResp.prototype.getMoney = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/** @param {number} value */
proto.pb.RechargeResp.prototype.setMoney = function(value) {
  jspb.Message.setProto3IntField(this, 5, value);
};


/**
 * optional int64 yuanBao = 6;
 * @return {number}
 */
proto.pb.RechargeResp.prototype.getYuanbao = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 6, 0));
};


/** @param {number} value */
proto.pb.RechargeResp.prototype.setYuanbao = function(value) {
  jspb.Message.setProto3IntField(this, 6, value);
};


/**
 * optional int64 coin = 7;
 * @return {number}
 */
proto.pb.RechargeResp.prototype.getCoin = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 7, 0));
};


/** @param {number} value */
proto.pb.RechargeResp.prototype.setCoin = function(value) {
  jspb.Message.setProto3IntField(this, 7, value);
};


/**
 * optional int32 method = 8;
 * @return {number}
 */
proto.pb.RechargeResp.prototype.getMethod = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 8, 0));
};


/** @param {number} value */
proto.pb.RechargeResp.prototype.setMethod = function(value) {
  jspb.Message.setProto3IntField(this, 8, value);
};


/**
 * optional bool isSuccess = 9;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.pb.RechargeResp.prototype.getIssuccess = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 9, false));
};


/** @param {boolean} value */
proto.pb.RechargeResp.prototype.setIssuccess = function(value) {
  jspb.Message.setProto3BooleanField(this, 9, value);
};


/**
 * optional string order = 10;
 * @return {string}
 */
proto.pb.RechargeResp.prototype.getOrder = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 10, ""));
};


/** @param {string} value */
proto.pb.RechargeResp.prototype.setOrder = function(value) {
  jspb.Message.setProto3StringField(this, 10, value);
};


/**
 * optional int64 timeStamp = 11;
 * @return {number}
 */
proto.pb.RechargeResp.prototype.getTimestamp = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 11, 0));
};


/** @param {number} value */
proto.pb.RechargeResp.prototype.setTimestamp = function(value) {
  jspb.Message.setProto3IntField(this, 11, value);
};


/**
 * optional string reason = 12;
 * @return {string}
 */
proto.pb.RechargeResp.prototype.getReason = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 12, ""));
};


/** @param {string} value */
proto.pb.RechargeResp.prototype.setReason = function(value) {
  jspb.Message.setProto3StringField(this, 12, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.GetRechargesReq.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.GetRechargesReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.GetRechargesReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.GetRechargesReq.toObject = function(includeInstance, msg) {
  var obj = {

  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.GetRechargesReq}
 */
proto.pb.GetRechargesReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.GetRechargesReq;
  return proto.pb.GetRechargesReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.GetRechargesReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.GetRechargesReq}
 */
proto.pb.GetRechargesReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.GetRechargesReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.GetRechargesReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.GetRechargesReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.GetRechargesReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.pb.GetRechargesResp.repeatedFields_ = [2];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.GetRechargesResp.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.GetRechargesResp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.GetRechargesResp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.GetRechargesResp.toObject = function(includeInstance, msg) {
  var obj = {
    userid: jspb.Message.getFieldWithDefault(msg, 1, 0),
    allrechargesList: jspb.Message.toObjectList(msg.getAllrechargesList(),
    proto.pb.RechargeResp.toObject, includeInstance),
    pagenum: jspb.Message.getFieldWithDefault(msg, 3, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.GetRechargesResp}
 */
proto.pb.GetRechargesResp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.GetRechargesResp;
  return proto.pb.GetRechargesResp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.GetRechargesResp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.GetRechargesResp}
 */
proto.pb.GetRechargesResp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setUserid(value);
      break;
    case 2:
      var value = new proto.pb.RechargeResp;
      reader.readMessage(value,proto.pb.RechargeResp.deserializeBinaryFromReader);
      msg.addAllrecharges(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setPagenum(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.GetRechargesResp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.GetRechargesResp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.GetRechargesResp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.GetRechargesResp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUserid();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getAllrechargesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      2,
      f,
      proto.pb.RechargeResp.serializeBinaryToWriter
    );
  }
  f = message.getPagenum();
  if (f !== 0) {
    writer.writeInt32(
      3,
      f
    );
  }
};


/**
 * optional int64 userID = 1;
 * @return {number}
 */
proto.pb.GetRechargesResp.prototype.getUserid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.GetRechargesResp.prototype.setUserid = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * repeated RechargeResp allRecharges = 2;
 * @return {!Array<!proto.pb.RechargeResp>}
 */
proto.pb.GetRechargesResp.prototype.getAllrechargesList = function() {
  return /** @type{!Array<!proto.pb.RechargeResp>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.pb.RechargeResp, 2));
};


/** @param {!Array<!proto.pb.RechargeResp>} value */
proto.pb.GetRechargesResp.prototype.setAllrechargesList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 2, value);
};


/**
 * @param {!proto.pb.RechargeResp=} opt_value
 * @param {number=} opt_index
 * @return {!proto.pb.RechargeResp}
 */
proto.pb.GetRechargesResp.prototype.addAllrecharges = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 2, opt_value, proto.pb.RechargeResp, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 */
proto.pb.GetRechargesResp.prototype.clearAllrechargesList = function() {
  this.setAllrechargesList([]);
};


/**
 * optional int32 pageNum = 3;
 * @return {number}
 */
proto.pb.GetRechargesResp.prototype.getPagenum = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/** @param {number} value */
proto.pb.GetRechargesResp.prototype.setPagenum = function(value) {
  jspb.Message.setProto3IntField(this, 3, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.GetGoodsReq.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.GetGoodsReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.GetGoodsReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.GetGoodsReq.toObject = function(includeInstance, msg) {
  var obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.GetGoodsReq}
 */
proto.pb.GetGoodsReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.GetGoodsReq;
  return proto.pb.GetGoodsReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.GetGoodsReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.GetGoodsReq}
 */
proto.pb.GetGoodsReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setId(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.GetGoodsReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.GetGoodsReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.GetGoodsReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.GetGoodsReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
};


/**
 * optional int64 iD = 1;
 * @return {number}
 */
proto.pb.GetGoodsReq.prototype.getId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.GetGoodsReq.prototype.setId = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.GetGoodsResp.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.GetGoodsResp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.GetGoodsResp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.GetGoodsResp.toObject = function(includeInstance, msg) {
  var obj = {
    userid: jspb.Message.getFieldWithDefault(msg, 1, 0),
    info: (f = msg.getInfo()) && proto.pb.GoodsInfo.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.GetGoodsResp}
 */
proto.pb.GetGoodsResp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.GetGoodsResp;
  return proto.pb.GetGoodsResp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.GetGoodsResp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.GetGoodsResp}
 */
proto.pb.GetGoodsResp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setUserid(value);
      break;
    case 2:
      var value = new proto.pb.GoodsInfo;
      reader.readMessage(value,proto.pb.GoodsInfo.deserializeBinaryFromReader);
      msg.setInfo(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.GetGoodsResp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.GetGoodsResp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.GetGoodsResp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.GetGoodsResp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUserid();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getInfo();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.pb.GoodsInfo.serializeBinaryToWriter
    );
  }
};


/**
 * optional int64 userID = 1;
 * @return {number}
 */
proto.pb.GetGoodsResp.prototype.getUserid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.GetGoodsResp.prototype.setUserid = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional GoodsInfo info = 2;
 * @return {?proto.pb.GoodsInfo}
 */
proto.pb.GetGoodsResp.prototype.getInfo = function() {
  return /** @type{?proto.pb.GoodsInfo} */ (
    jspb.Message.getWrapperField(this, proto.pb.GoodsInfo, 2));
};


/** @param {?proto.pb.GoodsInfo|undefined} value */
proto.pb.GetGoodsResp.prototype.setInfo = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 */
proto.pb.GetGoodsResp.prototype.clearInfo = function() {
  this.setInfo(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.pb.GetGoodsResp.prototype.hasInfo = function() {
  return jspb.Message.getField(this, 2) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.GetAllGoodsReq.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.GetAllGoodsReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.GetAllGoodsReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.GetAllGoodsReq.toObject = function(includeInstance, msg) {
  var obj = {

  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.GetAllGoodsReq}
 */
proto.pb.GetAllGoodsReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.GetAllGoodsReq;
  return proto.pb.GetAllGoodsReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.GetAllGoodsReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.GetAllGoodsReq}
 */
proto.pb.GetAllGoodsReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.GetAllGoodsReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.GetAllGoodsReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.GetAllGoodsReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.GetAllGoodsReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.pb.GetAllGoodsResp.repeatedFields_ = [2];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.GetAllGoodsResp.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.GetAllGoodsResp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.GetAllGoodsResp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.GetAllGoodsResp.toObject = function(includeInstance, msg) {
  var obj = {
    userid: jspb.Message.getFieldWithDefault(msg, 1, 0),
    infoList: jspb.Message.toObjectList(msg.getInfoList(),
    proto.pb.GoodsInfo.toObject, includeInstance),
    pagenum: jspb.Message.getFieldWithDefault(msg, 3, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.GetAllGoodsResp}
 */
proto.pb.GetAllGoodsResp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.GetAllGoodsResp;
  return proto.pb.GetAllGoodsResp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.GetAllGoodsResp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.GetAllGoodsResp}
 */
proto.pb.GetAllGoodsResp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setUserid(value);
      break;
    case 2:
      var value = new proto.pb.GoodsInfo;
      reader.readMessage(value,proto.pb.GoodsInfo.deserializeBinaryFromReader);
      msg.addInfo(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setPagenum(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.GetAllGoodsResp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.GetAllGoodsResp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.GetAllGoodsResp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.GetAllGoodsResp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUserid();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getInfoList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      2,
      f,
      proto.pb.GoodsInfo.serializeBinaryToWriter
    );
  }
  f = message.getPagenum();
  if (f !== 0) {
    writer.writeInt32(
      3,
      f
    );
  }
};


/**
 * optional int64 userID = 1;
 * @return {number}
 */
proto.pb.GetAllGoodsResp.prototype.getUserid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.GetAllGoodsResp.prototype.setUserid = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * repeated GoodsInfo info = 2;
 * @return {!Array<!proto.pb.GoodsInfo>}
 */
proto.pb.GetAllGoodsResp.prototype.getInfoList = function() {
  return /** @type{!Array<!proto.pb.GoodsInfo>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.pb.GoodsInfo, 2));
};


/** @param {!Array<!proto.pb.GoodsInfo>} value */
proto.pb.GetAllGoodsResp.prototype.setInfoList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 2, value);
};


/**
 * @param {!proto.pb.GoodsInfo=} opt_value
 * @param {number=} opt_index
 * @return {!proto.pb.GoodsInfo}
 */
proto.pb.GetAllGoodsResp.prototype.addInfo = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 2, opt_value, proto.pb.GoodsInfo, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 */
proto.pb.GetAllGoodsResp.prototype.clearInfoList = function() {
  this.setInfoList([]);
};


/**
 * optional int32 pageNum = 3;
 * @return {number}
 */
proto.pb.GetAllGoodsResp.prototype.getPagenum = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/** @param {number} value */
proto.pb.GetAllGoodsResp.prototype.setPagenum = function(value) {
  jspb.Message.setProto3IntField(this, 3, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.BuyGoodsReq.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.BuyGoodsReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.BuyGoodsReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.BuyGoodsReq.toObject = function(includeInstance, msg) {
  var obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, 0),
    payment: jspb.Message.getFieldWithDefault(msg, 2, 0),
    count: jspb.Message.getFieldWithDefault(msg, 3, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.BuyGoodsReq}
 */
proto.pb.BuyGoodsReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.BuyGoodsReq;
  return proto.pb.BuyGoodsReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.BuyGoodsReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.BuyGoodsReq}
 */
proto.pb.BuyGoodsReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setPayment(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setCount(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.BuyGoodsReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.BuyGoodsReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.BuyGoodsReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.BuyGoodsReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getPayment();
  if (f !== 0) {
    writer.writeInt64(
      2,
      f
    );
  }
  f = message.getCount();
  if (f !== 0) {
    writer.writeInt64(
      3,
      f
    );
  }
};


/**
 * optional int64 iD = 1;
 * @return {number}
 */
proto.pb.BuyGoodsReq.prototype.getId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.BuyGoodsReq.prototype.setId = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional int64 payment = 2;
 * @return {number}
 */
proto.pb.BuyGoodsReq.prototype.getPayment = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {number} value */
proto.pb.BuyGoodsReq.prototype.setPayment = function(value) {
  jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional int64 count = 3;
 * @return {number}
 */
proto.pb.BuyGoodsReq.prototype.getCount = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/** @param {number} value */
proto.pb.BuyGoodsReq.prototype.setCount = function(value) {
  jspb.Message.setProto3IntField(this, 3, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.BuyGoodsResp.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.BuyGoodsResp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.BuyGoodsResp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.BuyGoodsResp.toObject = function(includeInstance, msg) {
  var obj = {
    userid: jspb.Message.getFieldWithDefault(msg, 1, 0),
    info: (f = msg.getInfo()) && proto.pb.GoodsInfo.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.BuyGoodsResp}
 */
proto.pb.BuyGoodsResp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.BuyGoodsResp;
  return proto.pb.BuyGoodsResp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.BuyGoodsResp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.BuyGoodsResp}
 */
proto.pb.BuyGoodsResp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setUserid(value);
      break;
    case 2:
      var value = new proto.pb.GoodsInfo;
      reader.readMessage(value,proto.pb.GoodsInfo.deserializeBinaryFromReader);
      msg.setInfo(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.BuyGoodsResp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.BuyGoodsResp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.BuyGoodsResp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.BuyGoodsResp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUserid();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getInfo();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.pb.GoodsInfo.serializeBinaryToWriter
    );
  }
};


/**
 * optional int64 userID = 1;
 * @return {number}
 */
proto.pb.BuyGoodsResp.prototype.getUserid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.BuyGoodsResp.prototype.setUserid = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional GoodsInfo info = 2;
 * @return {?proto.pb.GoodsInfo}
 */
proto.pb.BuyGoodsResp.prototype.getInfo = function() {
  return /** @type{?proto.pb.GoodsInfo} */ (
    jspb.Message.getWrapperField(this, proto.pb.GoodsInfo, 2));
};


/** @param {?proto.pb.GoodsInfo|undefined} value */
proto.pb.BuyGoodsResp.prototype.setInfo = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 */
proto.pb.BuyGoodsResp.prototype.clearInfo = function() {
  this.setInfo(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.pb.BuyGoodsResp.prototype.hasInfo = function() {
  return jspb.Message.getField(this, 2) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.CheckKnapsackReq.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.CheckKnapsackReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.CheckKnapsackReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.CheckKnapsackReq.toObject = function(includeInstance, msg) {
  var obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, 0),
    number: jspb.Message.getFieldWithDefault(msg, 2, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.CheckKnapsackReq}
 */
proto.pb.CheckKnapsackReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.CheckKnapsackReq;
  return proto.pb.CheckKnapsackReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.CheckKnapsackReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.CheckKnapsackReq}
 */
proto.pb.CheckKnapsackReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setNumber(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.CheckKnapsackReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.CheckKnapsackReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.CheckKnapsackReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.CheckKnapsackReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getNumber();
  if (f !== 0) {
    writer.writeInt32(
      2,
      f
    );
  }
};


/**
 * optional int64 iD = 1;
 * @return {number}
 */
proto.pb.CheckKnapsackReq.prototype.getId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.CheckKnapsackReq.prototype.setId = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional int32 number = 2;
 * @return {number}
 */
proto.pb.CheckKnapsackReq.prototype.getNumber = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {number} value */
proto.pb.CheckKnapsackReq.prototype.setNumber = function(value) {
  jspb.Message.setProto3IntField(this, 2, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.CheckKnapsackResp.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.CheckKnapsackResp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.CheckKnapsackResp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.CheckKnapsackResp.toObject = function(includeInstance, msg) {
  var obj = {
    userid: jspb.Message.getFieldWithDefault(msg, 1, 0),
    info: (f = msg.getInfo()) && proto.pb.KnapsackInfo.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.CheckKnapsackResp}
 */
proto.pb.CheckKnapsackResp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.CheckKnapsackResp;
  return proto.pb.CheckKnapsackResp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.CheckKnapsackResp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.CheckKnapsackResp}
 */
proto.pb.CheckKnapsackResp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setUserid(value);
      break;
    case 2:
      var value = new proto.pb.KnapsackInfo;
      reader.readMessage(value,proto.pb.KnapsackInfo.deserializeBinaryFromReader);
      msg.setInfo(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.CheckKnapsackResp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.CheckKnapsackResp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.CheckKnapsackResp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.CheckKnapsackResp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUserid();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getInfo();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.pb.KnapsackInfo.serializeBinaryToWriter
    );
  }
};


/**
 * optional int64 userID = 1;
 * @return {number}
 */
proto.pb.CheckKnapsackResp.prototype.getUserid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.CheckKnapsackResp.prototype.setUserid = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional KnapsackInfo info = 2;
 * @return {?proto.pb.KnapsackInfo}
 */
proto.pb.CheckKnapsackResp.prototype.getInfo = function() {
  return /** @type{?proto.pb.KnapsackInfo} */ (
    jspb.Message.getWrapperField(this, proto.pb.KnapsackInfo, 2));
};


/** @param {?proto.pb.KnapsackInfo|undefined} value */
proto.pb.CheckKnapsackResp.prototype.setInfo = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 */
proto.pb.CheckKnapsackResp.prototype.clearInfo = function() {
  this.setInfo(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.pb.CheckKnapsackResp.prototype.hasInfo = function() {
  return jspb.Message.getField(this, 2) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.BarterReq.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.BarterReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.BarterReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.BarterReq.toObject = function(includeInstance, msg) {
  var obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, 0),
    toid: jspb.Message.getFieldWithDefault(msg, 2, 0),
    amount: jspb.Message.getFieldWithDefault(msg, 3, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.BarterReq}
 */
proto.pb.BarterReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.BarterReq;
  return proto.pb.BarterReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.BarterReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.BarterReq}
 */
proto.pb.BarterReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setToid(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setAmount(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.BarterReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.BarterReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.BarterReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.BarterReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getToid();
  if (f !== 0) {
    writer.writeInt64(
      2,
      f
    );
  }
  f = message.getAmount();
  if (f !== 0) {
    writer.writeInt32(
      3,
      f
    );
  }
};


/**
 * optional int64 iD = 1;
 * @return {number}
 */
proto.pb.BarterReq.prototype.getId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.BarterReq.prototype.setId = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional int64 toID = 2;
 * @return {number}
 */
proto.pb.BarterReq.prototype.getToid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {number} value */
proto.pb.BarterReq.prototype.setToid = function(value) {
  jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional int32 amount = 3;
 * @return {number}
 */
proto.pb.BarterReq.prototype.getAmount = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/** @param {number} value */
proto.pb.BarterReq.prototype.setAmount = function(value) {
  jspb.Message.setProto3IntField(this, 3, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.BarterResp.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.BarterResp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.BarterResp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.BarterResp.toObject = function(includeInstance, msg) {
  var obj = {
    userid: jspb.Message.getFieldWithDefault(msg, 1, 0),
    info: (f = msg.getInfo()) && proto.pb.KnapsackInfo.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.BarterResp}
 */
proto.pb.BarterResp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.BarterResp;
  return proto.pb.BarterResp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.BarterResp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.BarterResp}
 */
proto.pb.BarterResp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setUserid(value);
      break;
    case 2:
      var value = new proto.pb.KnapsackInfo;
      reader.readMessage(value,proto.pb.KnapsackInfo.deserializeBinaryFromReader);
      msg.setInfo(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.BarterResp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.BarterResp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.BarterResp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.BarterResp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUserid();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getInfo();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.pb.KnapsackInfo.serializeBinaryToWriter
    );
  }
};


/**
 * optional int64 userID = 1;
 * @return {number}
 */
proto.pb.BarterResp.prototype.getUserid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.BarterResp.prototype.setUserid = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional KnapsackInfo info = 2;
 * @return {?proto.pb.KnapsackInfo}
 */
proto.pb.BarterResp.prototype.getInfo = function() {
  return /** @type{?proto.pb.KnapsackInfo} */ (
    jspb.Message.getWrapperField(this, proto.pb.KnapsackInfo, 2));
};


/** @param {?proto.pb.KnapsackInfo|undefined} value */
proto.pb.BarterResp.prototype.setInfo = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 */
proto.pb.BarterResp.prototype.clearInfo = function() {
  this.setInfo(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.pb.BarterResp.prototype.hasInfo = function() {
  return jspb.Message.getField(this, 2) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.ToShoppingResp.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.ToShoppingResp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.ToShoppingResp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.ToShoppingResp.toObject = function(includeInstance, msg) {
  var obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, 0),
    count: jspb.Message.getFieldWithDefault(msg, 2, 0),
    reason: jspb.Message.getFieldWithDefault(msg, 3, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.ToShoppingResp}
 */
proto.pb.ToShoppingResp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.ToShoppingResp;
  return proto.pb.ToShoppingResp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.ToShoppingResp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.ToShoppingResp}
 */
proto.pb.ToShoppingResp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setCount(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setReason(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.ToShoppingResp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.ToShoppingResp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.ToShoppingResp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.ToShoppingResp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getCount();
  if (f !== 0) {
    writer.writeInt32(
      2,
      f
    );
  }
  f = message.getReason();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
};


/**
 * optional int64 iD = 1;
 * @return {number}
 */
proto.pb.ToShoppingResp.prototype.getId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.ToShoppingResp.prototype.setId = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional int32 count = 2;
 * @return {number}
 */
proto.pb.ToShoppingResp.prototype.getCount = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {number} value */
proto.pb.ToShoppingResp.prototype.setCount = function(value) {
  jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional string reason = 3;
 * @return {string}
 */
proto.pb.ToShoppingResp.prototype.getReason = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.pb.ToShoppingResp.prototype.setReason = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.EmailReq.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.EmailReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.EmailReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.EmailReq.toObject = function(includeInstance, msg) {
  var obj = {
    pagenum: jspb.Message.getFieldWithDefault(msg, 1, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.EmailReq}
 */
proto.pb.EmailReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.EmailReq;
  return proto.pb.EmailReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.EmailReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.EmailReq}
 */
proto.pb.EmailReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setPagenum(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.EmailReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.EmailReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.EmailReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.EmailReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getPagenum();
  if (f !== 0) {
    writer.writeInt32(
      1,
      f
    );
  }
};


/**
 * optional int32 pageNum = 1;
 * @return {number}
 */
proto.pb.EmailReq.prototype.getPagenum = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.EmailReq.prototype.setPagenum = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.pb.EmailResp.repeatedFields_ = [2];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.EmailResp.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.EmailResp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.EmailResp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.EmailResp.toObject = function(includeInstance, msg) {
  var obj = {
    userid: jspb.Message.getFieldWithDefault(msg, 1, 0),
    infosList: jspb.Message.toObjectList(msg.getInfosList(),
    proto.pb.EmailInfo.toObject, includeInstance),
    pagenum: jspb.Message.getFieldWithDefault(msg, 3, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.EmailResp}
 */
proto.pb.EmailResp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.EmailResp;
  return proto.pb.EmailResp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.EmailResp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.EmailResp}
 */
proto.pb.EmailResp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setUserid(value);
      break;
    case 2:
      var value = new proto.pb.EmailInfo;
      reader.readMessage(value,proto.pb.EmailInfo.deserializeBinaryFromReader);
      msg.addInfos(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setPagenum(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.EmailResp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.EmailResp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.EmailResp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.EmailResp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUserid();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getInfosList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      2,
      f,
      proto.pb.EmailInfo.serializeBinaryToWriter
    );
  }
  f = message.getPagenum();
  if (f !== 0) {
    writer.writeInt32(
      3,
      f
    );
  }
};


/**
 * optional int64 userID = 1;
 * @return {number}
 */
proto.pb.EmailResp.prototype.getUserid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.EmailResp.prototype.setUserid = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * repeated EmailInfo infos = 2;
 * @return {!Array<!proto.pb.EmailInfo>}
 */
proto.pb.EmailResp.prototype.getInfosList = function() {
  return /** @type{!Array<!proto.pb.EmailInfo>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.pb.EmailInfo, 2));
};


/** @param {!Array<!proto.pb.EmailInfo>} value */
proto.pb.EmailResp.prototype.setInfosList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 2, value);
};


/**
 * @param {!proto.pb.EmailInfo=} opt_value
 * @param {number=} opt_index
 * @return {!proto.pb.EmailInfo}
 */
proto.pb.EmailResp.prototype.addInfos = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 2, opt_value, proto.pb.EmailInfo, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 */
proto.pb.EmailResp.prototype.clearInfosList = function() {
  this.setInfosList([]);
};


/**
 * optional int32 pageNum = 3;
 * @return {number}
 */
proto.pb.EmailResp.prototype.getPagenum = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/** @param {number} value */
proto.pb.EmailResp.prototype.setPagenum = function(value) {
  jspb.Message.setProto3IntField(this, 3, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.ClaimReq.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.ClaimReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.ClaimReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.ClaimReq.toObject = function(includeInstance, msg) {
  var obj = {
    emailid: jspb.Message.getFieldWithDefault(msg, 1, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.ClaimReq}
 */
proto.pb.ClaimReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.ClaimReq;
  return proto.pb.ClaimReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.ClaimReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.ClaimReq}
 */
proto.pb.ClaimReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setEmailid(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.ClaimReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.ClaimReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.ClaimReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.ClaimReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getEmailid();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
};


/**
 * optional int64 emailID = 1;
 * @return {number}
 */
proto.pb.ClaimReq.prototype.getEmailid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.ClaimReq.prototype.setEmailid = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.ClaimResp.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.ClaimResp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.ClaimResp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.ClaimResp.toObject = function(includeInstance, msg) {
  var obj = {
    userid: jspb.Message.getFieldWithDefault(msg, 1, 0),
    emailid: jspb.Message.getFieldWithDefault(msg, 2, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.ClaimResp}
 */
proto.pb.ClaimResp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.ClaimResp;
  return proto.pb.ClaimResp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.ClaimResp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.ClaimResp}
 */
proto.pb.ClaimResp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setUserid(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setEmailid(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.ClaimResp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.ClaimResp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.ClaimResp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.ClaimResp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUserid();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getEmailid();
  if (f !== 0) {
    writer.writeInt64(
      2,
      f
    );
  }
};


/**
 * optional int64 userID = 1;
 * @return {number}
 */
proto.pb.ClaimResp.prototype.getUserid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.ClaimResp.prototype.setUserid = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional int64 emailID = 2;
 * @return {number}
 */
proto.pb.ClaimResp.prototype.getEmailid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {number} value */
proto.pb.ClaimResp.prototype.setEmailid = function(value) {
  jspb.Message.setProto3IntField(this, 2, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.SuggestReq.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.SuggestReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.SuggestReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.SuggestReq.toObject = function(includeInstance, msg) {
  var obj = {
    content: jspb.Message.getFieldWithDefault(msg, 1, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.SuggestReq}
 */
proto.pb.SuggestReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.SuggestReq;
  return proto.pb.SuggestReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.SuggestReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.SuggestReq}
 */
proto.pb.SuggestReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setContent(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.SuggestReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.SuggestReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.SuggestReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.SuggestReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getContent();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string content = 1;
 * @return {string}
 */
proto.pb.SuggestReq.prototype.getContent = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.pb.SuggestReq.prototype.setContent = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.SuggestResp.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.SuggestResp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.SuggestResp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.SuggestResp.toObject = function(includeInstance, msg) {
  var obj = {
    userid: jspb.Message.getFieldWithDefault(msg, 1, 0),
    feedback: (f = msg.getFeedback()) && proto.pb.EmailInfo.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.SuggestResp}
 */
proto.pb.SuggestResp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.SuggestResp;
  return proto.pb.SuggestResp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.SuggestResp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.SuggestResp}
 */
proto.pb.SuggestResp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setUserid(value);
      break;
    case 2:
      var value = new proto.pb.EmailInfo;
      reader.readMessage(value,proto.pb.EmailInfo.deserializeBinaryFromReader);
      msg.setFeedback(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.SuggestResp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.SuggestResp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.SuggestResp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.SuggestResp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUserid();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getFeedback();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.pb.EmailInfo.serializeBinaryToWriter
    );
  }
};


/**
 * optional int64 userID = 1;
 * @return {number}
 */
proto.pb.SuggestResp.prototype.getUserid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.SuggestResp.prototype.setUserid = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional EmailInfo feedback = 2;
 * @return {?proto.pb.EmailInfo}
 */
proto.pb.SuggestResp.prototype.getFeedback = function() {
  return /** @type{?proto.pb.EmailInfo} */ (
    jspb.Message.getWrapperField(this, proto.pb.EmailInfo, 2));
};


/** @param {?proto.pb.EmailInfo|undefined} value */
proto.pb.SuggestResp.prototype.setFeedback = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 */
proto.pb.SuggestResp.prototype.clearFeedback = function() {
  this.setFeedback(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.pb.SuggestResp.prototype.hasFeedback = function() {
  return jspb.Message.getField(this, 2) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.EmailReadReq.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.EmailReadReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.EmailReadReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.EmailReadReq.toObject = function(includeInstance, msg) {
  var obj = {
    emailid: jspb.Message.getFieldWithDefault(msg, 1, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.EmailReadReq}
 */
proto.pb.EmailReadReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.EmailReadReq;
  return proto.pb.EmailReadReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.EmailReadReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.EmailReadReq}
 */
proto.pb.EmailReadReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setEmailid(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.EmailReadReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.EmailReadReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.EmailReadReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.EmailReadReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getEmailid();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
};


/**
 * optional int64 emailID = 1;
 * @return {number}
 */
proto.pb.EmailReadReq.prototype.getEmailid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.EmailReadReq.prototype.setEmailid = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.EmailReadResp.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.EmailReadResp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.EmailReadResp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.EmailReadResp.toObject = function(includeInstance, msg) {
  var obj = {
    userid: jspb.Message.getFieldWithDefault(msg, 1, 0),
    emailid: jspb.Message.getFieldWithDefault(msg, 2, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.EmailReadResp}
 */
proto.pb.EmailReadResp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.EmailReadResp;
  return proto.pb.EmailReadResp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.EmailReadResp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.EmailReadResp}
 */
proto.pb.EmailReadResp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setUserid(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setEmailid(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.EmailReadResp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.EmailReadResp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.EmailReadResp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.EmailReadResp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUserid();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getEmailid();
  if (f !== 0) {
    writer.writeInt64(
      2,
      f
    );
  }
};


/**
 * optional int64 userID = 1;
 * @return {number}
 */
proto.pb.EmailReadResp.prototype.getUserid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.EmailReadResp.prototype.setUserid = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional int64 emailID = 2;
 * @return {number}
 */
proto.pb.EmailReadResp.prototype.getEmailid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {number} value */
proto.pb.EmailReadResp.prototype.setEmailid = function(value) {
  jspb.Message.setProto3IntField(this, 2, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.EmailDelReq.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.EmailDelReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.EmailDelReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.EmailDelReq.toObject = function(includeInstance, msg) {
  var obj = {
    emailid: jspb.Message.getFieldWithDefault(msg, 1, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.EmailDelReq}
 */
proto.pb.EmailDelReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.EmailDelReq;
  return proto.pb.EmailDelReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.EmailDelReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.EmailDelReq}
 */
proto.pb.EmailDelReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setEmailid(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.EmailDelReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.EmailDelReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.EmailDelReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.EmailDelReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getEmailid();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
};


/**
 * optional int64 emailID = 1;
 * @return {number}
 */
proto.pb.EmailDelReq.prototype.getEmailid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.EmailDelReq.prototype.setEmailid = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.EmailDelResp.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.EmailDelResp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.EmailDelResp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.EmailDelResp.toObject = function(includeInstance, msg) {
  var obj = {
    userid: jspb.Message.getFieldWithDefault(msg, 1, 0),
    emailid: jspb.Message.getFieldWithDefault(msg, 2, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.EmailDelResp}
 */
proto.pb.EmailDelResp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.EmailDelResp;
  return proto.pb.EmailDelResp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.EmailDelResp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.EmailDelResp}
 */
proto.pb.EmailDelResp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setUserid(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setEmailid(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.EmailDelResp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.EmailDelResp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.EmailDelResp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.EmailDelResp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUserid();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getEmailid();
  if (f !== 0) {
    writer.writeInt64(
      2,
      f
    );
  }
};


/**
 * optional int64 userID = 1;
 * @return {number}
 */
proto.pb.EmailDelResp.prototype.getUserid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.EmailDelResp.prototype.setUserid = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional int64 emailID = 2;
 * @return {number}
 */
proto.pb.EmailDelResp.prototype.getEmailid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {number} value */
proto.pb.EmailDelResp.prototype.setEmailid = function(value) {
  jspb.Message.setProto3IntField(this, 2, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.ResultResp.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.ResultResp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.ResultResp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.ResultResp.toObject = function(includeInstance, msg) {
  var obj = {
    state: jspb.Message.getFieldWithDefault(msg, 1, 0),
    hints: jspb.Message.getFieldWithDefault(msg, 2, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.ResultResp}
 */
proto.pb.ResultResp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.ResultResp;
  return proto.pb.ResultResp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.ResultResp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.ResultResp}
 */
proto.pb.ResultResp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setState(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setHints(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.ResultResp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.ResultResp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.ResultResp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.ResultResp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getState();
  if (f !== 0) {
    writer.writeInt32(
      1,
      f
    );
  }
  f = message.getHints();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional int32 state = 1;
 * @return {number}
 */
proto.pb.ResultResp.prototype.getState = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.ResultResp.prototype.setState = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional string hints = 2;
 * @return {string}
 */
proto.pb.ResultResp.prototype.getHints = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.pb.ResultResp.prototype.setHints = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.ResultPopResp.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.ResultPopResp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.ResultPopResp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.ResultPopResp.toObject = function(includeInstance, msg) {
  var obj = {
    flag: jspb.Message.getFieldWithDefault(msg, 1, 0),
    title: jspb.Message.getFieldWithDefault(msg, 2, ""),
    hints: jspb.Message.getFieldWithDefault(msg, 3, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.ResultPopResp}
 */
proto.pb.ResultPopResp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.ResultPopResp;
  return proto.pb.ResultPopResp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.ResultPopResp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.ResultPopResp}
 */
proto.pb.ResultPopResp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setFlag(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setTitle(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setHints(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.ResultPopResp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.ResultPopResp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.ResultPopResp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.ResultPopResp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getFlag();
  if (f !== 0) {
    writer.writeInt32(
      1,
      f
    );
  }
  f = message.getTitle();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getHints();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
};


/**
 * optional int32 flag = 1;
 * @return {number}
 */
proto.pb.ResultPopResp.prototype.getFlag = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.ResultPopResp.prototype.setFlag = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional string title = 2;
 * @return {string}
 */
proto.pb.ResultPopResp.prototype.getTitle = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.pb.ResultPopResp.prototype.setTitle = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string hints = 3;
 * @return {string}
 */
proto.pb.ResultPopResp.prototype.getHints = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.pb.ResultPopResp.prototype.setHints = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.PingReq.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.PingReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.PingReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.PingReq.toObject = function(includeInstance, msg) {
  var obj = {

  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.PingReq}
 */
proto.pb.PingReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.PingReq;
  return proto.pb.PingReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.PingReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.PingReq}
 */
proto.pb.PingReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.PingReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.PingReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.PingReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.PingReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.PongResp.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.PongResp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.PongResp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.PongResp.toObject = function(includeInstance, msg) {
  var obj = {

  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.PongResp}
 */
proto.pb.PongResp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.PongResp;
  return proto.pb.PongResp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.PongResp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.PongResp}
 */
proto.pb.PongResp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.PongResp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.PongResp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.PongResp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.PongResp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
};


/**
 * @enum {number}
 */
proto.pb.HeroType = {
  HERONULL: 0,
  JIN: 1,
  MU: 2,
  SHUI: 3,
  HUO: 4,
  TU: 5
};

/**
 * @enum {number}
 */
proto.pb.TableState = {
  INITTB: 0,
  OPENTB: 1,
  REPAIRTB: 2,
  CLEARTB: 3,
  STOPTB: 4,
  CLOSETB: 5
};

/**
 * @enum {number}
 */
proto.pb.GameType = {
  GENERAL: 0,
  FIGHT: 1,
  MULTIPERSON: 2,
  TABLECARD: 3,
  GUESS: 4,
  GAMESCITY: 5,
  DUALMEET: 6,
  SPORT: 7,
  SMART: 8,
  RPG: 9
};

/**
 * @enum {number}
 */
proto.pb.GameScene = {
  FREE: 0,
  START: 1,
  CALL: 2,
  DECIDE: 3,
  PLAYING: 4,
  OPENING: 5,
  OVER: 6,
  CLOSING: 7,
  SITDIRECT: 8,
  ROLLDICE: 9,
  WAITOPERATE: 10,
  CHANGETHREE: 11,
  DINGQUE: 12,
  CHECKTING: 13,
  CHECKHUAZHU: 14
};

goog.object.extend(exports, proto.pb);
