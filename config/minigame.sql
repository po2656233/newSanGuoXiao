/*
 Navicat Premium Data Transfer

 Source Server         : 本地数据库
 Source Server Type    : MySQL
 Source Server Version : 80032
 Source Host           : localhost:3306
 Source Schema         : minigame

 Target Server Type    : MySQL
 Target Server Version : 80032
 File Encoding         : 65001

 Date: 20/09/2024 16:08:14
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for activity
-- ----------------------------
DROP TABLE IF EXISTS `activity`;
CREATE TABLE `activity`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id ',
  `name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '名字',
  `announcement` varchar(512) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '公告',
  `startTime` datetime(6) NULL DEFAULT NULL COMMENT '起始时间',
  `endTime` datetime(6) NULL DEFAULT NULL COMMENT '结束时间',
  `intentionality` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '目的',
  `prize` varchar(525) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT 'pd字节格式物品奖励',
  `remark` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '备注',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `update_by` bigint NULL DEFAULT 0,
  `create_by` bigint NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of activity
-- ----------------------------
INSERT INTO `activity` VALUES (1, '国庆', '庆祝国庆，全民大抽奖', '2019-10-01 00:00:00.000000', '2019-12-31 23:10:59.000000', '庆祝活动', '', '备注', NULL, NULL, NULL, 0, 0);
INSERT INTO `activity` VALUES (2, '双十一', '减价大促销', '2019-10-01 00:00:00.000000', '2019-12-31 23:10:59.000000', '活动促销', '', '促卖国货', NULL, NULL, NULL, 0, 0);
INSERT INTO `activity` VALUES (3, '圣诞', '礼物大派送', '2019-10-01 00:00:00.000000', '2019-12-31 23:10:59.000000', '节日活动', '', '拉高人气', NULL, NULL, NULL, 0, 0);

-- ----------------------------
-- Table structure for assetgoods
-- ----------------------------
DROP TABLE IF EXISTS `assetgoods`;
CREATE TABLE `assetgoods`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '资产表id',
  `uid` bigint NOT NULL COMMENT '所得的人',
  `goodsid` bigint NOT NULL COMMENT '商品id',
  `amount` int NULL DEFAULT NULL COMMENT '当前拥有数量',
  `spending` int NULL DEFAULT NULL COMMENT '已花费数量',
  `count` bigint NULL DEFAULT NULL COMMENT '累计总数量',
  `totalprice` bigint NULL DEFAULT NULL COMMENT '总价(每次消费的购买)',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `update_by` bigint NULL DEFAULT 0,
  `create_by` bigint NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uniqueid`(`id` ASC) USING BTREE,
  INDEX `normaluid`(`uid` ASC) USING BTREE,
  INDEX `normalgoods`(`goodsid` ASC) USING BTREE,
  INDEX `idx_asset_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of assetgoods
-- ----------------------------
INSERT INTO `assetgoods` VALUES (1, 1, 1, 8800, 188, 8800, 269640, '2024-02-29 17:00:02', '2024-02-29 23:23:46', NULL, 0, 0);

-- ----------------------------
-- Table structure for assethero
-- ----------------------------
DROP TABLE IF EXISTS `assethero`;
CREATE TABLE `assethero`  (
  `id` bigint NOT NULL COMMENT 'assetheros',
  `uid` bigint NOT NULL COMMENT '所得的人',
  `heroid` bigint NOT NULL COMMENT '英雄id',
  `amount` int NULL DEFAULT NULL COMMENT '当前拥有数量',
  `spending` int NULL DEFAULT NULL COMMENT '已花费数量',
  `count` bigint NULL DEFAULT NULL COMMENT '累计总数量',
  `fragments` int NULL DEFAULT 0 COMMENT '英雄碎片',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uniqueid`(`uid` ASC, `heroid` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of assethero
-- ----------------------------

-- ----------------------------
-- Table structure for assetweapon
-- ----------------------------
DROP TABLE IF EXISTS `assetweapon`;
CREATE TABLE `assetweapon`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'assetweapon武器',
  `uid` bigint NOT NULL COMMENT '所得的人',
  `weaponid` bigint NOT NULL COMMENT '武器id',
  `amount` int NULL DEFAULT NULL COMMENT '当前拥有数量',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uniqueid`(`uid` ASC, `weaponid` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of assetweapon
-- ----------------------------

-- ----------------------------
-- Table structure for categories
-- ----------------------------
DROP TABLE IF EXISTS `categories`;
CREATE TABLE `categories`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '游戏分类表',
  `name` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `type` int NULL DEFAULT NULL,
  `kind` int NULL DEFAULT NULL,
  `level` int NULL DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `update_by` bigint NULL DEFAULT 0,
  `create_by` bigint NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of categories
-- ----------------------------
INSERT INTO `categories` VALUES (1, '三国消', 1, 1010, 0, '', '2023-06-06 02:00:43', '2023-06-06 02:00:43', NULL, 0, 0);

-- ----------------------------
-- Table structure for chat_group
-- ----------------------------
DROP TABLE IF EXISTS `chat_group`;
CREATE TABLE `chat_group`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '群名称',
  `hostid` bigint NULL DEFAULT 0 COMMENT '群主ID',
  `setuptime` bigint NULL DEFAULT 0 COMMENT '创建时间',
  `explain` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '群简介',
  `notice` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '群公告',
  `adminlist` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '管理者者',
  `memberlist` varchar(4096) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '成员列表',
  `applylist` varchar(1024) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '申请列表',
  `bannedlist` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '禁言者',
  `robotid` bigint NULL DEFAULT 0 COMMENT '机器人ID',
  `robotkey` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '机器人密钥',
  `robotcontrol` int NULL DEFAULT 0 COMMENT '机器人控制(0:关闭 1:开启 2:停止 3:维护 4:销毁)',
  `remark` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '缓存key',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `update_by` bigint NULL DEFAULT 0,
  `create_by` bigint NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_chat_group_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of chat_group
-- ----------------------------
INSERT INTO `chat_group` VALUES (1, 'gamegrp3324', 3324, 1608659191, '', '', '', '3324', '2023', '', 0, '', 0, '9ff27580-447d-11eb-8000-436861745365', NULL, NULL, NULL, 0, 0);
INSERT INTO `chat_group` VALUES (2, 'gamegrp2023', 2023, 1608670338, '', '', '', '2023', '', '', 0, '', 0, '9413ad00-4497-11eb-8000-436861745365', NULL, NULL, NULL, 0, 0);
INSERT INTO `chat_group` VALUES (3, 'gamegrp23', 23, 1608745383, '', '', '', '23', '', '', 0, '', 0, '4e61ed80-4546-11eb-8000-436861745365', NULL, NULL, NULL, 0, 0);

-- ----------------------------
-- Table structure for checkin
-- ----------------------------
DROP TABLE IF EXISTS `checkin`;
CREATE TABLE `checkin`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `uid` bigint NOT NULL COMMENT '用户',
  `timestamp` bigint NULL DEFAULT NULL COMMENT '时间戳',
  `remark` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '备注',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `update_by` bigint NULL DEFAULT 0,
  `create_by` bigint NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `userid`(`id` ASC, `uid` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of checkin
-- ----------------------------

-- ----------------------------
-- Table structure for email
-- ----------------------------
DROP TABLE IF EXISTS `email`;
CREATE TABLE `email`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `accepter` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '接收者账号',
  `sender` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '发送者账号',
  `carboncopy` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '抄送给',
  `topic` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '主题',
  `content` varchar(1024) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '内容',
  `goods` varchar(525) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT 'pd字节格式物品奖励',
  `isread` int NULL DEFAULT 0 COMMENT '0:未读 1:已读',
  `timestamp` bigint NULL DEFAULT 0 COMMENT '时间',
  `remark` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '备注',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `update_by` bigint NULL DEFAULT 0,
  `create_by` bigint NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id`(`id` ASC, `sender` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 22931 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of email
-- ----------------------------
INSERT INTO `email` VALUES (1, '', '系统', '', '注册奖励', '        欢迎您成为指尖家族的一员！作为指尖的成员，请您遵守所在地方的法律法规，切勿肆意枉法。为了您的健康和幸福，也请您务必适度游戏，量力消费。<br/>        为了欢迎您的到来，指尖家族特意准备了一份微薄的见面礼，麻烦您点击<color=#FF0000> 领取</color>。<br/>        在此，谨代表指尖家族全体同仁，祝您游戏快乐，心情愉悦，开心每一天！', '10 21 8 233 7 18 12 233 128 154 231 148 168 230 136 191 229 141 161 24 1 64 10 10 18 8 232 7 18 9 232 161 168 230 131 133 229 140 133 24 10 64 3', NULL, 0, NULL, NULL, NULL, NULL, 0, 0);

-- ----------------------------
-- Table structure for game
-- ----------------------------
DROP TABLE IF EXISTS `game`;
CREATE TABLE `game`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `kid` int NOT NULL DEFAULT 0 COMMENT '种类',
  `en_name` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '英文名，用来匹配生成游戏',
  `name` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '游戏名称',
  `lessscore` bigint NULL DEFAULT 0 COMMENT '底分',
  `state` int NULL DEFAULT 0 COMMENT '状态(0未开放 1正常 2维护 3关闭)',
  `max_player` int NULL DEFAULT -1 COMMENT '最大人数(-1:无限制)',
  `remark` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '备注',
  `how_to_play` varchar(2048) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '玩法介绍',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `update_by` bigint NULL DEFAULT 0,
  `create_by` bigint NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uniqueid`(`id` ASC) USING BTREE,
  INDEX `normalkid`(`kid` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 84 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of game
-- ----------------------------
INSERT INTO `game` VALUES (1, 10, 'ChineseChess', '中国象棋', 0, 1, 2, '', '棋盘：\r\n中国象棋的棋盘有九条平行的竖线和十条平行的横线相交而组成，其中共有九十个交叉点，棋子就摆在这些交叉点上。中间部分（棋盘的第五，第六两横线之间末画竖线的空白地带）称为 河界 。两端的中间（两端第四条到第六条竖线之间的正方形部位），以斜交叉线构成中文米字形方格的地方，叫作 九宫 （它恰好有九个交叉点）。\r\n\r\n界河：\r\n这是一个和国际象棋不同的地方，及对垒双方的中间有一条河界，通常称其为楚河，也就是说棋子过河才能攻打对方的首领。这些规则都是和中国古战场的一些场景类似，并且逐渐演变而来成为中国象棋的行棋基本规则。\r\n\r\n九宫：\r\n双方的底线中心处，也就是纵向中心线分别向两边外移一条线（第四条到第六条竖线）之间的正方形部位，以斜交叉线构成 米 字方格的地方，叫作 九宫 （它恰好有九个交叉点），知道这个区域，对理解一些走棋规则很有帮助。\r\n\r\n棋子：\r\n所有的棋子共有三十二个，其中又分为红、黑两组（分别代表对垒的一方），每组共有十六个棋子（为了区别双方的棋子，不仅在颜色中有所区别，而且还会使用同音不同字的棋子），其中又各分为七种棋子，其名称和数目如下：\r\n红棋子：帅一个，车、马、炮、相、士各两个，兵五个。\r\n黑棋子：将一个，车、马、炮、象、士各两个，卒五个。\r\n\r\n将（帅）：\r\n虽然名称不同，但它们这都是对垒双方的最高统帅，对垒的目的就是通过运用各自的棋子，想方设法将对方的首领将死，方为己方胜利。\r\n这两位棋子的最高统帅，只能在九宫内行走，不得走出九宫外。行走的步法为：左、右横走，上、下竖走都行，但每次只能行走一格。将和帅不准在同一直线上直接对面（中间无棋子），如一方已先占据位置，则另一方必须回避，否则就算输了。\r\n\r\n士（仕）：\r\n每行一步棋，只许沿着 九宫 中的斜线行走一步（方格的对角线），行走方位可进、可退，其最终目的也是为了护卫各自的最高将领（帅、将）。\r\n\r\n象（相）：\r\n此棋不能越过 河界走入对方的领地，其走法为：只能斜走（两步），可以使用汉字中的田字形象地表述：田字格的对角线，即俗称象（相）走田字。\r\n行走方位可进、可退，但是，当象（相）行走的路线中，及田字中心有棋子时（无论己方或者是对方的棋子），则不允许走过去，俗称：塞象（相）眼。\r\n\r\n车（車）：\r\n此棋是中国象棋中棋力最强的棋子，每行一步棋可以上、下直线行走（进、退）；左、右横走（中间不隔棋子），且行棋步数不限。\r\n\r\n炮（砲）：\r\n此棋的行棋规则和车（車）类似，横平、竖直，只要前方没有棋子的地方都能行走。\r\n但是，它的吃棋规则很特别，必须跳过一个棋子（无论是己方的还是对方的）去吃掉对方的一个棋子。俗称：隔山打炮。\r\n\r\n马（馬）：\r\n走棋规则：使用中国的日字来形容马的行走方式比较贴切，俗称：马走日字（斜对角线）。\r\n但是，这里有一个行走规则，可以将马走日分解为：\r\n先一步直走（或一横）再一步斜走，如果在要去的方向，第一步直行处（或者横行）有别的棋子挡住，则不许走过去（俗称：蹩马腿）。\r\n行走范围不限，可以进、也可以退。\r\n\r\n卒（兵）：\r\n在没有过河界前，此棋每走一步棋只许向前直走一步（不能后退）；过了河界之后，每行一步棋可以向前直走，或者横走（左、右）一步，但也是不能后退的。\r\n根据此规则，卒（兵）走到对方的底线只能左右横走，俗称：兵走老了！\r\n\r\n吃子：\r\n①无论什么棋子，通常只要根据行棋规则能走到的部位有对方的棋子就能吃掉对方的棋子。\r\n②而唯一列外的是炮的吃棋方法，比较特殊，需要中间隔有旗子（无论是己方的还是对方的棋子）才能吃掉对方的棋子。\r\n将死和困毙：\r\n①一方的棋子攻击对方的将（帅），并在下一步要把它吃掉，称为照将，或简称将。照将不必声明。\r\n②被照将的一方必须立即应将，即用自己的着法去化解被将的状态（而不能应将不顾，而走其它的棋子）。\r\n③如果被照将而无法应将，就算被将死（一方胜棋）。\r\n④轮到走棋的一方，无子可走，就算被困毙（无棋可走这方为输棋）。\r\n\r\n行棋规则：\r\n对局时，由执红棋的一方先走，双方轮流各走一着（双方各走一着，称为一个回合），直至分出胜、负、和，对局即算终了。', NULL, NULL, NULL, 0, 0);
INSERT INTO `game` VALUES (2, 10, 'Chess', '国际象棋', 0, 1, 2, '', '', NULL, NULL, NULL, 0, 0);
INSERT INTO `game` VALUES (3, 10, 'SanGuoXiao', '三国消', 0, 1, 2, '', '', NULL, NULL, NULL, 0, 0);
INSERT INTO `game` VALUES (4, 10, 'ThreeKingdoms', '三国杀', 0, 1, 5, '', '', NULL, NULL, NULL, 0, 0);
INSERT INTO `game` VALUES (5, 1, 'Landlord', '斗地主', 0, 1, 3, '', '', NULL, NULL, NULL, 0, 0);
INSERT INTO `game` VALUES (6, 11, 'GuobiaoMahjong', '国标麻将', 0, 1, 4, '', '', NULL, NULL, NULL, 0, 0);
INSERT INTO `game` VALUES (7, 11, 'GuangDongMahjong', '广东麻将', 0, 1, 4, '', '', NULL, NULL, NULL, 0, 0);
INSERT INTO `game` VALUES (8, 11, 'SiChuanMahjong', '四川麻将', 0, 1, 4, '', '', NULL, NULL, NULL, 0, 0);
INSERT INTO `game` VALUES (9, 11, 'XLCHMahjong', '血流成河', 0, 1, 4, '', '', NULL, NULL, NULL, 0, 0);
INSERT INTO `game` VALUES (10, 11, 'XZDDMahjong', '血战到底', 0, 1, 4, '', '', NULL, NULL, NULL, 0, 0);
INSERT INTO `game` VALUES (11, 11, 'ChangShaMahjong', '长沙麻将', 0, 1, 4, '', '', NULL, NULL, NULL, 0, 0);
INSERT INTO `game` VALUES (12, 1, 'RunFast', '跑得快', 0, 1, 3, '', '', NULL, NULL, NULL, 0, 0);
INSERT INTO `game` VALUES (13, 1, 'Paohuzi', '跑胡子', 0, 1, 3, '', '', NULL, NULL, NULL, 0, 0);
INSERT INTO `game` VALUES (14, 4, 'FruitGame', '水果机', 0, 1, -1, '', '', NULL, NULL, NULL, 0, 0);
INSERT INTO `game` VALUES (15, 10, 'KingOfFighters', '拳皇', 0, 1, 2, '', '', NULL, NULL, NULL, 0, 0);
INSERT INTO `game` VALUES (16, 9, 'SuperMario', '超级玛丽', 0, 1, 1, '', '', NULL, NULL, NULL, 0, 0);
INSERT INTO `game` VALUES (17, 1, 'CatchFish', '捕鱼', 0, 1, 8, '', '', NULL, NULL, NULL, 0, 0);
INSERT INTO `game` VALUES (19, 4, 'SicBo', '骰宝', 0, 1, -1, '', '', NULL, NULL, NULL, 0, 0);
INSERT INTO `game` VALUES (20, 4, 'TenMinutesOfHappiness', '十分快乐', 0, 1, -1, '', '', NULL, NULL, NULL, 0, 0);
INSERT INTO `game` VALUES (21, 11, 'CatchChickenInGuiyang', '贵阳捉鸡', 0, 1, 4, '', '', NULL, NULL, NULL, 0, 0);
INSERT INTO `game` VALUES (22, 1, 'SanGong', '三公', 0, 1, 4, '', '', NULL, NULL, NULL, 0, 0);
INSERT INTO `game` VALUES (23, 2, 'Baccarat', '百佳乐', 0, 1, -1, '', '', NULL, NULL, NULL, 0, 0);
INSERT INTO `game` VALUES (24, 2, 'BrCowcow', '百人牛妞', 0, 1, -1, '', '', NULL, NULL, NULL, 0, 0);
INSERT INTO `game` VALUES (25, 2, 'TigerXdragon', '龙虎逗', 0, 1, -1, '', '', NULL, NULL, NULL, 0, 0);
INSERT INTO `game` VALUES (26, 2, 'toubao', '骰宝', 0, 1, -1, '', '', NULL, NULL, NULL, 0, 0);
INSERT INTO `game` VALUES (27, 2, 'tuitongzi', '推筒子', 0, 1, -1, '', '', NULL, NULL, NULL, 0, 0);

-- ----------------------------
-- Table structure for goods
-- ----------------------------
DROP TABLE IF EXISTS `goods`;
CREATE TABLE `goods`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '物品ID',
  `name` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '物品名称',
  `kind` int NULL DEFAULT 0 COMMENT '物品类别(作为房卡时，等同于kindID)',
  `level` int NULL DEFAULT 0 COMMENT '级别',
  `price` bigint NULL DEFAULT 0 COMMENT '单价',
  `store` bigint NULL DEFAULT 0 COMMENT '库存',
  `sold` bigint NULL DEFAULT NULL COMMENT '已销',
  `remark` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '备注',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `update_by` bigint NULL DEFAULT 0,
  `create_by` bigint NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `goodsid`(`id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1002 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of goods
-- ----------------------------
INSERT INTO `goods` VALUES (1, '三国消房卡', 1010, 0, 30, 1000000, 500, '三国消房卡', NULL, NULL, NULL, 0, 0);

-- ----------------------------
-- Table structure for hero
-- ----------------------------
DROP TABLE IF EXISTS `hero`;
CREATE TABLE `hero`  (
  `id` bigint NOT NULL COMMENT '英雄[武将]ID',
  `headid` int NULL DEFAULT NULL COMMENT '头像ID',
  `name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '名称',
  `sex` int NULL DEFAULT NULL COMMENT '性别',
  `country` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '国家',
  `faction` int NULL DEFAULT NULL COMMENT '所属阵营',
  `health` bigint NULL DEFAULT NULL COMMENT '体力值',
  `attack` bigint NULL DEFAULT NULL COMMENT '攻击力',
  `armor` bigint NULL DEFAULT NULL COMMENT '防御力',
  `strength` bigint NULL DEFAULT NULL COMMENT '力量',
  `agility` bigint NULL DEFAULT NULL COMMENT '敏捷',
  `intelligence` bigint NULL DEFAULT NULL COMMENT '智力',
  `spellpower` bigint NULL DEFAULT NULL COMMENT '法强',
  `skills` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '技能列表',
  `rarity` bigint NULL DEFAULT NULL COMMENT '稀有度',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of hero
-- ----------------------------
INSERT INTO `hero` VALUES (1, 1, '曹操', 1, '曹魏', 1, 10000, 1000, 20, 100, 200, 100, 110, '1|2|3', 100);
INSERT INTO `hero` VALUES (2, 2, '刘备', 1, '蜀汉', 2, 10000, 1000, 20, 100, 200, 100, 110, '1|2|3', 100);
INSERT INTO `hero` VALUES (3, 3, '孙权', 1, '东吴', 3, 10000, 1000, 20, 100, 200, 100, 110, '1|2|3', 100);
INSERT INTO `hero` VALUES (4, 4, '董卓', 1, '群', 4, 10000, 1000, 20, 100, 200, 100, 110, '1|2|3', 100);
INSERT INTO `hero` VALUES (5, 5, '袁绍', 1, '群', 5, 10000, 1000, 20, 100, 200, 100, 110, '1|2|3', 100);
INSERT INTO `hero` VALUES (6, 6, '吕布', 1, '群', 6, 10000, 1000, 20, 100, 200, 100, 110, '1|2|3', 100);
INSERT INTO `hero` VALUES (7, 7, '关羽', 1, '蜀汉', 2, 10000, 1000, 20, 100, 200, 100, 110, '1|2|3', 100);
INSERT INTO `hero` VALUES (8, 8, '张飞', 1, '蜀汉', 2, 10000, 1000, 20, 100, 200, 100, 110, '1|2|3', 100);
INSERT INTO `hero` VALUES (9, 9, '诸葛亮', 1, '蜀汉', 2, 10000, 1000, 20, 100, 200, 100, 110, '1|2|3', 100);
INSERT INTO `hero` VALUES (10, 10, '郭嘉', 1, '曹魏', 1, 10000, 1000, 20, 100, 200, 100, 110, '1|2|3', 100);
INSERT INTO `hero` VALUES (11, 11, '典韦', 1, '曹魏', 1, 10000, 1000, 20, 100, 200, 100, 110, '1|2|3', 100);
INSERT INTO `hero` VALUES (12, 12, '夏侯惇', 1, '曹魏', 1, 10000, 1000, 20, 100, 200, 100, 110, '1|2|3', 100);
INSERT INTO `hero` VALUES (13, 13, '庞统', 1, '蜀汉', 2, 10000, 1000, 20, 100, 200, 100, 110, '1|2|3', 100);
INSERT INTO `hero` VALUES (14, 14, '周瑜', 1, '东吴', 3, 10000, 1000, 20, 100, 200, 100, 110, '1|2|3', 100);

-- ----------------------------
-- Table structure for kindinfo
-- ----------------------------
DROP TABLE IF EXISTS `kindinfo`;
CREATE TABLE `kindinfo`  (
  `id` int NOT NULL COMMENT '种类ID',
  `name` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '种类名称',
  `en_name` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '种类英文名称',
  `remark` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '备注',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `update_by` bigint NULL DEFAULT 0,
  `create_by` bigint NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uniqueid`(`id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of kindinfo
-- ----------------------------
INSERT INTO `kindinfo` VALUES (1, '对战类', 'fight', '两个或两个以上玩家同桌博弈', NULL, NULL, NULL, 0, 0);
INSERT INTO `kindinfo` VALUES (2, '百人类', 'multiplayer', '多个玩家同时下注', NULL, NULL, NULL, 0, 0);
INSERT INTO `kindinfo` VALUES (3, '房卡类', 'sameTable', '该模式下的游戏，不对金币进行结算', NULL, NULL, NULL, 0, 0);
INSERT INTO `kindinfo` VALUES (4, '竞猜类', 'guess', '竞猜游戏结果，如足彩', NULL, NULL, NULL, 0, 0);
INSERT INTO `kindinfo` VALUES (5, '电玩城', 'gamesCity', '电玩城，如水果机', NULL, NULL, NULL, 0, 0);
INSERT INTO `kindinfo` VALUES (6, '电竞类', 'dualMeet ', '电竞，如dota,lol', NULL, NULL, NULL, 0, 0);
INSERT INTO `kindinfo` VALUES (7, '体育类', 'sport ', '体育赛事，如篮球，足球', NULL, NULL, NULL, 0, 0);
INSERT INTO `kindinfo` VALUES (8, '益智类', 'smart', '智力比拼，如象棋，迷宫', NULL, NULL, NULL, 0, 0);
INSERT INTO `kindinfo` VALUES (9, 'RPG类', 'RolePlayingGame', '角色扮演', NULL, NULL, NULL, 0, 0);
INSERT INTO `kindinfo` VALUES (10, 'PK类', 'PlayerKilling', '两人以上对抗', NULL, NULL, NULL, 0, 0);
INSERT INTO `kindinfo` VALUES (11, '麻将类', 'mahjong', '麻将牌相关玩法', NULL, NULL, NULL, 0, 0);

-- ----------------------------
-- Table structure for notice
-- ----------------------------
DROP TABLE IF EXISTS `notice`;
CREATE TABLE `notice`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '通知ID',
  `type` int NULL DEFAULT 0 COMMENT '通知类型',
  `platid` bigint NULL DEFAULT NULL COMMENT '平台ID',
  `gameids` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '指定游戏',
  `kindid` bigint NULL DEFAULT 0 COMMENT '关联游戏',
  `level` int NULL DEFAULT 0 COMMENT '游戏级别',
  `title` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '标题',
  `content` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '内容',
  `start` bigint NULL DEFAULT 0 COMMENT '起始时间',
  `end` bigint NULL DEFAULT 0 COMMENT '结束时间',
  `remark` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '备注',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `update_by` bigint NULL DEFAULT 0,
  `create_by` bigint NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of notice
-- ----------------------------
INSERT INTO `notice` VALUES (1, 1, 1, '', 1004, 0, '喜迎2024', '<color=#FF0000>喜迎2024!</c><color=#FFFF00>     元旦期间,麻将房房卡免费大馈赠，欢迎您和熟人前来霸房!</color>', 1609802145, 1640906144, '测试', NULL, NULL, NULL, 0, 0);

-- ----------------------------
-- Table structure for platform
-- ----------------------------
DROP TABLE IF EXISTS `platform`;
CREATE TABLE `platform`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'id',
  `name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '名字',
  `usercount` int NULL DEFAULT NULL COMMENT '在线人数',
  `allow_kinds` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '允许运行的游戏种类。以逗号分隔',
  `activities` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '日常活动',
  `code` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '与用户表里的平台ID对应',
  `servers` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '服务器地址',
  `remark` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '备注',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `update_by` bigint NULL DEFAULT 0,
  `create_by` bigint NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uniqueid`(`id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of platform
-- ----------------------------
INSERT INTO `platform` VALUES (1, '大中华', 2, '1,2,3,4,5,6,7,8,9,10,11,12,13', '1,2', 'AC32147568', '192.168.1.108:9650', '用户房间优先,如果用户没有指定房间列表,再按平台房间', NULL, NULL, NULL, 0, 0);
INSERT INTO `platform` VALUES (2, '大地利宝', 10000, '1,2,3,4,5,6,7,8,9,10,11,12,13', '1,2', 'AC32147568', '192.168.1.108:9650', '平台没有房间时，将视为不存在该平台', NULL, NULL, NULL, 0, 0);
INSERT INTO `platform` VALUES (3, '小天鹅', 600000, '1,2,3,4,5,6,7,8,9,10,11,12,13', '1,2', ' ', '192.168.1.108:9650', '平台没有房间时，将视为不存在该平台', NULL, NULL, NULL, 0, 0);
INSERT INTO `platform` VALUES (4, '华东一区', 2000, '1,2,3,4,5,6,7,8,9,10,11,12,13', '1,2', ' ', '192.168.1.108:9650', '江苏省、浙江省、安徽省、福建省、江西省、山东省和上海市、台湾省。', NULL, NULL, NULL, 0, 0);
INSERT INTO `platform` VALUES (5, '华南一区', 20000, '1,2,3,4,5,6,7,8,9,10,11,12,13', '1,2', ' ', '192.168.1.108:9650', '广东省（包括东沙群岛）、广西省、海南省（包括南海诸岛）、香港和澳门特区。', NULL, NULL, NULL, 0, 0);
INSERT INTO `platform` VALUES (6, '华北一区', 2000, '1,2,3,4,5,6,7,8,9,10,11,12,13', '1,2', ' ', '192.168.1.108:9650', '河北省、山西省、北京市、天津市和内蒙古自治区的部分地区。', NULL, NULL, NULL, 0, 0);
INSERT INTO `platform` VALUES (7, '华中一区', 2000, '1,2,3,4,5,6,7,8,9,10,11,12,13', '1,2', ' ', '192.168.1.108:9650', '湖北省，湖南省，河南省，华中地区的区域中心城市为武汉。', NULL, NULL, NULL, 0, 0);
INSERT INTO `platform` VALUES (8, '东北一区', 2000, '1,2,3,4,5,6,7,8,9,10,11,12,13', '1,2', ' ', '192.168.1.108:9650', '辽宁省、吉林省、黑龙江省，或说东北四省区（包括内蒙古东部）', NULL, NULL, NULL, 0, 0);
INSERT INTO `platform` VALUES (9, '西南一区', 2000, '1,2,3,4,5,6,7,8,9,10,11,12,13', '1,2', ' ', '192.168.1.108:9650', '四川省、云南省、贵州省、重庆市、西藏自治区的大部以及陕西省南部（陕南地区）。', NULL, NULL, NULL, 0, 0);
INSERT INTO `platform` VALUES (10, '西北一区', 2000, '1,2,3,4,5,6,7,8,9,10,11,12,13', '1,2', ' ', '192.168.1.108:9650', '宁夏回族自治区 、新疆维吾尔自治区及青海 、陕西 、甘肃三省之地。', NULL, NULL, NULL, 0, 0);

-- ----------------------------
-- Table structure for recharge
-- ----------------------------
DROP TABLE IF EXISTS `recharge`;
CREATE TABLE `recharge`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '充值ID',
  `uid` bigint NULL DEFAULT NULL COMMENT '接受充值的人',
  `byid` bigint NULL DEFAULT NULL COMMENT '代充者',
  `payment` bigint NULL DEFAULT NULL COMMENT '支付费用',
  `premoney` bigint NULL DEFAULT NULL COMMENT '充值前',
  `money` bigint NULL DEFAULT NULL COMMENT '充值后',
  `code` int NULL DEFAULT NULL COMMENT '充值码[充值方式]',
  `order` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '订单号',
  `timestamp` bigint NULL DEFAULT NULL COMMENT '充值时间',
  `remark` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '备注',
  `switch` int NULL DEFAULT 0 COMMENT '0:(余额)不转换 1:转元宝 2:转铜钱',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `update_by` bigint NULL DEFAULT 0,
  `create_by` bigint NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uniqueid`(`id` ASC) USING BTREE,
  INDEX `normaluid`(`uid` ASC) USING BTREE,
  INDEX `normalbyid`(`byid` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of recharge
-- ----------------------------
INSERT INTO `recharge` VALUES (1, 44, 35, 500, 4300, 9100, 0, '172312113144fTEx', 1723121131, '', 0, '2024-08-08 20:45:31', '2024-08-08 20:45:31', NULL, 0, 0);
INSERT INTO `recharge` VALUES (2, 44, 35, 500, 9100, 18700, 0, '172312177144Sskv', 1723121771, '', 0, '2024-08-08 20:56:12', '2024-08-08 20:56:12', NULL, 0, 0);
INSERT INTO `recharge` VALUES (3, 44, 35, 500, 18700, 37900, 0, '1723121853MiXG476856517111316544', 1723121853, '', 0, '2024-08-08 20:57:33', '2024-08-08 20:57:33', NULL, 0, 0);
INSERT INTO `recharge` VALUES (4, 44, 35, 500, 37900, 76300, 0, '1723122895Tnsf476860889589350464', 1723122895, '', 0, '2024-08-08 21:14:56', '2024-08-08 21:14:56', NULL, 0, 0);
INSERT INTO `recharge` VALUES (5, 44, 35, 500, 76300, 153100, 0, '1723122903luON476860923210891328', 1723122903, '', 0, '2024-08-08 21:15:04', '2024-08-08 21:15:04', NULL, 0, 0);

-- ----------------------------
-- Table structure for record
-- ----------------------------
DROP TABLE IF EXISTS `record`;
CREATE TABLE `record`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '金币变化记录',
  `uid` bigint NULL DEFAULT 0 COMMENT '充值者ID',
  `tid` bigint NULL DEFAULT 0 COMMENT '牌桌ID',
  `pergold` bigint NULL DEFAULT 0 COMMENT '支付之前',
  `payment` bigint NULL DEFAULT 0 COMMENT '支付',
  `gold` bigint NULL DEFAULT 0 COMMENT '金币',
  `code` int NULL DEFAULT 0 COMMENT '操作码',
  `order` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '订单号(牌局号)',
  `result` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '开奖结果',
  `remark` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '备注',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `update_by` bigint NULL DEFAULT 0,
  `create_by` bigint NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uniqueid`(`id` ASC) USING BTREE,
  INDEX `normaluid`(`uid` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 602 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of record
-- ----------------------------
INSERT INTO `record` VALUES (113, 47, 42, 2657, 138, 2795, 0, 'F9F454A2ED6E2DB32F9EEF8413A25B8D', '闲:♥10,♦9,	庄:♦K,♠5, 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-08 14:28:08', '2024-08-08 14:28:08', NULL, 0, 0);
INSERT INTO `record` VALUES (114, 48, 42, 4489, 4, 4493, 0, 'F9F454A2ED6E2DB32F9EEF8413A25B8D', '闲:♥10,♦9,	庄:♦K,♠5, 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-08 14:28:08', '2024-08-08 14:28:08', NULL, 0, 0);
INSERT INTO `record` VALUES (115, 47, 42, 2795, 83, 2878, 0, 'A251E8D7B3769DDBE2F0DFCA3EDE3BE1', '闲:♥1,♣7,	庄:♣K,♦6, 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-08 14:28:29', '2024-08-08 14:28:29', NULL, 0, 0);
INSERT INTO `record` VALUES (116, 48, 42, 4493, -141, 4352, 0, 'A251E8D7B3769DDBE2F0DFCA3EDE3BE1', '闲:♥1,♣7,	庄:♣K,♦6, 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-08 14:28:29', '2024-08-08 14:28:29', NULL, 0, 0);
INSERT INTO `record` VALUES (117, 47, 42, 2878, -88, 2790, 0, '071830617747E6A7627C744734F275F4', '闲:♥1,♣1,♣8,	庄:♠7,♠6, 中奖区域:[0 0 1 0 0 0 1 0]', '百佳乐', '2024-08-08 14:28:50', '2024-08-08 14:28:50', NULL, 0, 0);
INSERT INTO `record` VALUES (118, 48, 42, 4352, -340, 4012, 0, '071830617747E6A7627C744734F275F4', '闲:♥1,♣1,♣8,	庄:♠7,♠6, 中奖区域:[0 0 1 0 0 0 1 0]', '百佳乐', '2024-08-08 14:28:50', '2024-08-08 14:28:50', NULL, 0, 0);
INSERT INTO `record` VALUES (119, 47, 42, 2790, -246, 2544, 0, '4159CBF3B95846B1FEE6B206933DCE9F', '闲:♠10,♠4,♥7,	庄:♥7,♦5,♣8, 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-08 14:29:11', '2024-08-08 14:29:11', NULL, 0, 0);
INSERT INTO `record` VALUES (120, 48, 42, 4012, -454, 3558, 0, '4159CBF3B95846B1FEE6B206933DCE9F', '闲:♠10,♠4,♥7,	庄:♥7,♦5,♣8, 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-08 14:29:11', '2024-08-08 14:29:11', NULL, 0, 0);
INSERT INTO `record` VALUES (121, 47, 42, 2544, -255, 2289, 0, 'F4AABF9410C90C62AAC327E7B32EDD45', '闲:♥8,♠7,	庄:♣J,♦9, 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-08 14:29:32', '2024-08-08 14:29:32', NULL, 0, 0);
INSERT INTO `record` VALUES (122, 48, 42, 3558, -585, 2973, 0, 'F4AABF9410C90C62AAC327E7B32EDD45', '闲:♥8,♠7,	庄:♣J,♦9, 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-08 14:29:32', '2024-08-08 14:29:32', NULL, 0, 0);
INSERT INTO `record` VALUES (123, 48, 42, 2973, -761, 2212, 0, 'AFB6F688BEB69E3F5B3B1CE30832A1C0', '闲:♠1,♥J,♠9,	庄:♥K,♠1,♣Q, 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-08 14:29:53', '2024-08-08 14:29:53', NULL, 0, 0);
INSERT INTO `record` VALUES (124, 47, 42, 2289, -255, 2034, 0, 'AFB6F688BEB69E3F5B3B1CE30832A1C0', '闲:♠1,♥J,♠9,	庄:♥K,♠1,♣Q, 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-08 14:29:53', '2024-08-08 14:29:53', NULL, 0, 0);
INSERT INTO `record` VALUES (125, 47, 44, 2034, 311, 2345, 0, '1FB84BCAB3642CA460F061F54D23B799', 'Banker[5] T[2] D[7] X[3] H[0] 中奖区域:[0 1 0 0]', '百人牛妞', '2024-08-08 21:43:27', '2024-08-08 21:43:27', NULL, 0, 0);
INSERT INTO `record` VALUES (126, 48, 43, 2212, -192, 2020, 0, '1FB84BCAB3642CA460F061F54D23B799', 'Banker[5] T[2] D[7] X[3] H[0] 中奖区域:[0 1 0 0]', '百人牛妞', '2024-08-08 21:43:27', '2024-08-08 21:43:27', NULL, 0, 0);
INSERT INTO `record` VALUES (127, 47, 44, 2345, 304, 2649, 0, '374C3C9134647F52C3FFFE28BBA7D897', 'Banker[0] T[0] D[0] X[10] H[7] 中奖区域:[0 0 1 1]', '百人牛妞', '2024-08-08 21:43:48', '2024-08-08 21:43:48', NULL, 0, 0);
INSERT INTO `record` VALUES (128, 48, 43, 2020, -192, 1828, 0, '374C3C9134647F52C3FFFE28BBA7D897', 'Banker[0] T[0] D[0] X[10] H[7] 中奖区域:[0 0 1 1]', '百人牛妞', '2024-08-08 21:43:48', '2024-08-08 21:43:48', NULL, 0, 0);
INSERT INTO `record` VALUES (129, 47, 44, 2649, 297, 2946, 0, 'D199D799791C70350D13E678EE57EAE1', 'Banker[2] T[0] D[0] X[6] H[5] 中奖区域:[0 0 1 1]', '百人牛妞', '2024-08-08 21:44:09', '2024-08-08 21:44:09', NULL, 0, 0);
INSERT INTO `record` VALUES (130, 48, 43, 1828, -192, 1636, 0, 'D199D799791C70350D13E678EE57EAE1', 'Banker[2] T[0] D[0] X[6] H[5] 中奖区域:[0 0 1 1]', '百人牛妞', '2024-08-08 21:44:09', '2024-08-08 21:44:09', NULL, 0, 0);
INSERT INTO `record` VALUES (131, 47, 44, 2946, 291, 3237, 0, 'AA17FB3ACAEBCE826724BE17B3EE81B6', 'Banker[0] T[10] D[5] X[8] H[4] 中奖区域:[1 1 1 1]', '百人牛妞', '2024-08-08 21:44:30', '2024-08-08 21:44:30', NULL, 0, 0);
INSERT INTO `record` VALUES (132, 48, 43, 1636, -192, 1444, 0, 'AA17FB3ACAEBCE826724BE17B3EE81B6', 'Banker[0] T[10] D[5] X[8] H[4] 中奖区域:[1 1 1 1]', '百人牛妞', '2024-08-08 21:44:30', '2024-08-08 21:44:30', NULL, 0, 0);
INSERT INTO `record` VALUES (133, 48, 43, 1444, -192, 1252, 0, '6AAA45807BA99DCB5AB144BE79635B4B', 'Banker[0] T[7] D[5] X[0] H[0] 中奖区域:[1 1 0 0]', '百人牛妞', '2024-08-08 21:44:51', '2024-08-08 21:44:51', NULL, 0, 0);
INSERT INTO `record` VALUES (134, 47, 44, 3237, 355, 3592, 0, '6AAA45807BA99DCB5AB144BE79635B4B', 'Banker[0] T[7] D[5] X[0] H[0] 中奖区域:[1 1 0 0]', '百人牛妞', '2024-08-08 21:44:51', '2024-08-08 21:44:51', NULL, 0, 0);
INSERT INTO `record` VALUES (135, 47, 44, 3592, 347, 3939, 0, '357E00047295CDEA17E67452BF5C1296', 'Banker[1] T[0] D[4] X[10] H[0] 中奖区域:[0 1 1 0]', '百人牛妞', '2024-08-08 21:45:12', '2024-08-08 21:45:12', NULL, 0, 0);
INSERT INTO `record` VALUES (136, 48, 43, 1252, -192, 1060, 0, 'FF668552A72BFCD5DDEFB776D2C6B7A8', 'Banker[1] T[0] D[4] X[10] H[0] 中奖区域:[0 1 1 0]', '百人牛妞', '2024-08-08 21:45:12', '2024-08-08 21:45:12', NULL, 0, 0);
INSERT INTO `record` VALUES (137, 47, 44, 3939, 340, 4279, 0, '046CCE0357895216CFAB2BF5F6391604', 'Banker[0] T[2] D[0] X[4] H[8] 中奖区域:[1 1 1 1]', '百人牛妞', '2024-08-08 21:45:33', '2024-08-08 21:45:33', NULL, 0, 0);
INSERT INTO `record` VALUES (138, 48, 43, 1060, -192, 868, 0, '046CCE0357895216CFAB2BF5F6391604', 'Banker[8] T[4] D[9] X[9] H[4] 中奖区域:[0 1 1 0]', '百人牛妞', '2024-08-08 21:45:33', '2024-08-08 21:45:33', NULL, 0, 0);
INSERT INTO `record` VALUES (139, 47, 44, 4279, 333, 4612, 0, 'F064AD4279E4BCF6EA3DE1C0FA5195A0', 'Banker[2] T[10] D[10] X[1] H[0] 中奖区域:[1 1 0 0]', '百人牛妞', '2024-08-08 21:45:54', '2024-08-08 21:45:54', NULL, 0, 0);
INSERT INTO `record` VALUES (140, 48, 43, 868, -192, 676, 0, 'F064AD4279E4BCF6EA3DE1C0FA5195A0', 'Banker[2] T[10] D[10] X[1] H[0] 中奖区域:[1 1 0 0]', '百人牛妞', '2024-08-08 21:45:54', '2024-08-08 21:45:54', NULL, 0, 0);
INSERT INTO `record` VALUES (141, 47, 44, 4612, 326, 4938, 0, 'D654EDA44D92E00D5C4FAC1D17B318F2', 'Banker[9] T[9] D[5] X[0] H[9] 中奖区域:[1 0 0 1]', '百人牛妞', '2024-08-08 21:46:15', '2024-08-08 21:46:15', NULL, 0, 0);
INSERT INTO `record` VALUES (142, 48, 43, 676, -192, 484, 0, 'D654EDA44D92E00D5C4FAC1D17B318F2', 'Banker[10] T[0] D[6] X[0] H[0] 中奖区域:[0 0 0 0]', '百人牛妞', '2024-08-08 21:46:15', '2024-08-08 21:46:15', NULL, 0, 0);
INSERT INTO `record` VALUES (143, 48, 43, 484, -192, 292, 0, 'DDC6CBAF9A50438DC7BB2769BEDA85B2', 'Banker[0] T[0] D[7] X[1] H[4] 中奖区域:[0 1 1 1]', '百人牛妞', '2024-08-08 21:46:36', '2024-08-08 21:46:36', NULL, 0, 0);
INSERT INTO `record` VALUES (144, 47, 44, 4938, 319, 5257, 0, 'AC65A636539ED3D3BF631E9DF1D5C96C', 'Banker[0] T[0] D[7] X[1] H[4] 中奖区域:[0 1 1 1]', '百人牛妞', '2024-08-08 21:46:36', '2024-08-08 21:46:36', NULL, 0, 0);
INSERT INTO `record` VALUES (145, 47, 44, 5257, 312, 5569, 0, 'E7F17C5BB31FD05762CECE7131BDDABB', 'Banker[1] T[8] D[0] X[1] H[0] 中奖区域:[1 0 1 0]', '百人牛妞', '2024-08-08 21:46:58', '2024-08-08 21:46:58', NULL, 0, 0);
INSERT INTO `record` VALUES (146, 48, 43, 292, -192, 100, 0, 'E7F17C5BB31FD05762CECE7131BDDABB', 'Banker[1] T[8] D[0] X[1] H[0] 中奖区域:[1 0 1 0]', '百人牛妞', '2024-08-08 21:46:58', '2024-08-08 21:46:58', NULL, 0, 0);
INSERT INTO `record` VALUES (147, 47, 44, 5569, 305, 5874, 0, '663E2637E3710B221828AE817CA6D2CE', 'Banker[6] T[10] D[0] X[4] H[2] 中奖区域:[1 0 0 0]', '百人牛妞', '2024-08-08 21:47:19', '2024-08-08 21:47:19', NULL, 0, 0);
INSERT INTO `record` VALUES (148, 48, 43, 100, -192, -92, 0, '663E2637E3710B221828AE817CA6D2CE', 'Banker[6] T[10] D[0] X[4] H[2] 中奖区域:[1 0 0 0]', '百人牛妞', '2024-08-08 21:47:19', '2024-08-08 21:47:19', NULL, 0, 0);
INSERT INTO `record` VALUES (149, 47, 44, 5874, 298, 6172, 0, '98C3380E36C7FD8C660DDC7804734F8C', 'Banker[3] T[9] D[0] X[2] H[2] 中奖区域:[1 0 0 0]', '百人牛妞', '2024-08-08 21:47:40', '2024-08-08 21:47:40', NULL, 0, 0);
INSERT INTO `record` VALUES (150, 48, 43, -92, -192, -284, 0, 'A6CA0344F546F886009C3A9D20D399C9', 'Banker[0] T[4] D[1] X[0] H[7] 中奖区域:[1 1 1 1]', '百人牛妞', '2024-08-08 21:47:40', '2024-08-08 21:47:40', NULL, 0, 0);
INSERT INTO `record` VALUES (151, 47, 44, 6172, 290, 6462, 0, '72B43F76D4834FF8BB4F361EF709EB9D', 'Banker[6] T[4] D[1] X[0] H[10] 中奖区域:[0 0 0 1]', '百人牛妞', '2024-08-08 21:48:01', '2024-08-08 21:48:01', NULL, 0, 0);
INSERT INTO `record` VALUES (152, 48, 43, -284, -192, -476, 0, '62DB20B28755377DF6B00DE06264F59E', 'Banker[6] T[4] D[1] X[0] H[10] 中奖区域:[0 0 0 1]', '百人牛妞', '2024-08-08 21:48:01', '2024-08-08 21:48:01', NULL, 0, 0);
INSERT INTO `record` VALUES (153, 48, 43, -476, -192, -668, 0, '721BC7703D7B8C26AE0679AD3DDC7C49', 'Banker[1] T[5] D[2] X[3] H[9] 中奖区域:[1 1 1 1]', '百人牛妞', '2024-08-08 21:48:22', '2024-08-08 21:48:22', NULL, 0, 0);
INSERT INTO `record` VALUES (154, 47, 44, 6462, 284, 6746, 0, '53578E39AABE2412F77A0E6F04F2DF87', 'Banker[1] T[5] D[2] X[3] H[9] 中奖区域:[1 1 1 1]', '百人牛妞', '2024-08-08 21:48:22', '2024-08-08 21:48:22', NULL, 0, 0);
INSERT INTO `record` VALUES (155, 48, 43, -668, -192, -860, 0, '7AEF64DDCBAA3781F92931A8FDAFD121', 'Banker[8] T[0] D[3] X[0] H[4] 中奖区域:[0 0 0 0]', '百人牛妞', '2024-08-08 21:48:43', '2024-08-08 21:48:43', NULL, 0, 0);
INSERT INTO `record` VALUES (156, 47, 44, 6746, 278, 7024, 0, '2F110488DFE9BC04FC2C69C83F3E1AE0', 'Banker[5] T[0] D[5] X[0] H[0] 中奖区域:[0 1 0 0]', '百人牛妞', '2024-08-08 21:48:43', '2024-08-08 21:48:43', NULL, 0, 0);
INSERT INTO `record` VALUES (157, 48, 43, -860, -192, -1052, 0, 'B04780F0C82DB9759BEC2DEF6B0BF026', 'Banker[3] T[0] D[0] X[0] H[0] 中奖区域:[0 0 0 0]', '百人牛妞', '2024-08-08 21:49:04', '2024-08-08 21:49:04', NULL, 0, 0);
INSERT INTO `record` VALUES (158, 47, 44, 7024, 272, 7296, 0, 'B04780F0C82DB9759BEC2DEF6B0BF026', 'Banker[3] T[0] D[0] X[0] H[0] 中奖区域:[0 0 0 0]', '百人牛妞', '2024-08-08 21:49:04', '2024-08-08 21:49:04', NULL, 0, 0);
INSERT INTO `record` VALUES (159, 48, 43, -1052, -192, -1244, 0, '241EEEEA3D7E18CA7478B448DECD52B1', 'Banker[0] T[8] D[8] X[0] H[6] 中奖区域:[1 1 0 1]', '百人牛妞', '2024-08-08 21:49:25', '2024-08-08 21:49:25', NULL, 0, 0);
INSERT INTO `record` VALUES (160, 47, 44, 7296, 266, 7562, 0, '241EEEEA3D7E18CA7478B448DECD52B1', 'Banker[0] T[8] D[8] X[0] H[6] 中奖区域:[1 1 0 1]', '百人牛妞', '2024-08-08 21:49:25', '2024-08-08 21:49:25', NULL, 0, 0);
INSERT INTO `record` VALUES (161, 47, 44, 7562, 260, 7822, 0, '3CCDF9F5882A23EEBA3101CC93DB269A', 'Banker[3] T[0] D[8] X[6] H[0] 中奖区域:[0 1 1 0]', '百人牛妞', '2024-08-08 21:49:46', '2024-08-08 21:49:46', NULL, 0, 0);
INSERT INTO `record` VALUES (162, 48, 43, -1244, -192, -1436, 0, '3CCDF9F5882A23EEBA3101CC93DB269A', 'Banker[3] T[0] D[8] X[6] H[0] 中奖区域:[0 1 1 0]', '百人牛妞', '2024-08-08 21:49:46', '2024-08-08 21:49:46', NULL, 0, 0);
INSERT INTO `record` VALUES (163, 48, 43, -1436, -192, -1628, 0, 'C1F71E536130EC2C5E326605AFD6F6D4', 'Banker[2] T[1] D[1] X[1] H[0] 中奖区域:[0 0 0 0]', '百人牛妞', '2024-08-08 21:50:07', '2024-08-08 21:50:07', NULL, 0, 0);
INSERT INTO `record` VALUES (164, 47, 44, 7822, 254, 8076, 0, '91E2F66ADA57803F7C0EC04AD005D397', 'Banker[2] T[1] D[1] X[1] H[0] 中奖区域:[0 0 0 0]', '百人牛妞', '2024-08-08 21:50:07', '2024-08-08 21:50:07', NULL, 0, 0);
INSERT INTO `record` VALUES (165, 44, 42, 5000000, 109, 5000109, 0, 'DC2363E7712F455307FB02B2CC97E81C', '闲:[♥8,♦7,♣Q,] 庄:[♠2,♠9,♦4,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-09 10:26:00', '2024-08-09 10:26:00', NULL, 0, 0);
INSERT INTO `record` VALUES (166, 47, 43, 8076, 176, 8252, 0, 'B517208631B781F0FA1CE92D2179934F', 'Banker[0] T[0] D[6] X[6] H[6] 中奖区域:[0 1 1 1]', '百人牛妞', '2024-08-09 10:26:10', '2024-08-09 10:26:10', NULL, 0, 0);
INSERT INTO `record` VALUES (167, 44, 42, 5000109, 143, 5000252, 0, '050AD04FE3705EB4A99467942B153FDB', '闲:[♦7,♥1,] 庄:[♦9,♥4,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-09 10:26:21', '2024-08-09 10:26:21', NULL, 0, 0);
INSERT INTO `record` VALUES (168, 47, 43, 8252, 172, 8424, 0, '38FD9759B5B218AA04BF346FCB42A3B4', 'Banker[3] T[9] D[8] X[10] H[1] 中奖区域:[1 1 1 0]', '百人牛妞', '2024-08-09 10:26:31', '2024-08-09 10:26:31', NULL, 0, 0);
INSERT INTO `record` VALUES (169, 44, 42, 5000252, 46, 5000298, 0, '04A682D74401344A6932569533159543', '闲:[♣K,♣2,] 庄:[♦9,♥J,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-09 10:26:42', '2024-08-09 10:26:42', NULL, 0, 0);
INSERT INTO `record` VALUES (170, 47, 43, 8424, 168, 8592, 0, '5636637020FDE1E6ED09018EF9640A2E', 'Banker[3] T[0] D[8] X[2] H[9] 中奖区域:[0 1 0 1]', '百人牛妞', '2024-08-09 10:26:52', '2024-08-09 10:26:52', NULL, 0, 0);
INSERT INTO `record` VALUES (171, 44, 42, 5000298, 171, 5000469, 0, 'F5D76EA237888678EFA6CC137B63631D', '闲:[♥6,♦J,] 庄:[♥1,♥10,♥3,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-09 10:27:03', '2024-08-09 10:27:03', NULL, 0, 0);
INSERT INTO `record` VALUES (172, 47, 43, 8592, 164, 8756, 0, 'A9A081602CED8C4F49A4CCE80726FC89', 'Banker[0] T[3] D[6] X[0] H[0] 中奖区域:[1 1 1 0]', '百人牛妞', '2024-08-09 10:27:13', '2024-08-09 10:27:13', NULL, 0, 0);
INSERT INTO `record` VALUES (173, 44, 42, 5000469, 164, 5000633, 0, '90A9DF82D1F20481D22B1313CF0046A3', '闲:[♣10,♦J,♥3,] 庄:[♥3,♦9,♥7,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-09 10:27:24', '2024-08-09 10:27:24', NULL, 0, 0);
INSERT INTO `record` VALUES (174, 47, 43, 8756, 160, 8916, 0, '89AC21BD775017883EE8B6F2E422BD3C', 'Banker[1] T[8] D[4] X[3] H[7] 中奖区域:[1 1 1 1]', '百人牛妞', '2024-08-09 10:27:34', '2024-08-09 10:27:34', NULL, 0, 0);
INSERT INTO `record` VALUES (175, 44, 42, 5000633, 17, 5000650, 0, '7F3F997527944D40C7FBC7C9D113F4D4', '闲:[♦8,♥2,♥3,] 庄:[♦Q,♣4,♠3,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-09 10:27:45', '2024-08-09 10:27:45', NULL, 0, 0);
INSERT INTO `record` VALUES (176, 47, 43, 8916, 156, 9072, 0, '185D109BA9890AD0A593FF3E37C1C5FE', 'Banker[10] T[0] D[2] X[2] H[8] 中奖区域:[0 0 0 0]', '百人牛妞', '2024-08-09 10:27:55', '2024-08-09 10:27:55', NULL, 0, 0);
INSERT INTO `record` VALUES (177, 44, 42, 5000650, -9, 5000641, 0, 'A45DA799E1205C5000A9E21801270C76', '闲:[♠K,♥Q,♥8,] 庄:[♠5,♥5,] 中奖区域:[1 0 0 1 0 0 0 1]', '百佳乐', '2024-08-09 10:28:06', '2024-08-09 10:28:06', NULL, 0, 0);
INSERT INTO `record` VALUES (178, 47, 43, 9072, 152, 9224, 0, '492B8E04F0927FBDFD1447DF00FFE4DB', 'Banker[6] T[8] D[0] X[2] H[7] 中奖区域:[1 0 0 1]', '百人牛妞', '2024-08-09 10:28:16', '2024-08-09 10:28:16', NULL, 0, 0);
INSERT INTO `record` VALUES (179, 44, 42, 5000641, -90, 5000551, 0, 'DC6E26216D3B6CAF7A75025AA57A7364', '闲:[♦4,♣Q,♠5,] 庄:[♣2,♠4,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-09 10:28:27', '2024-08-09 10:28:27', NULL, 0, 0);
INSERT INTO `record` VALUES (180, 47, 43, 9224, 148, 9372, 0, 'F9FF129D150CDEDDD3CE34F229DF0EFE', 'Banker[6] T[9] D[0] X[0] H[7] 中奖区域:[1 0 0 1]', '百人牛妞', '2024-08-09 10:28:37', '2024-08-09 10:28:37', NULL, 0, 0);
INSERT INTO `record` VALUES (181, 44, 42, 5000551, -94, 5000457, 0, '85DD61E44A19255302B689F138D88666', '闲:[♣1,♠3,♠8,] 庄:[♠J,♣Q,♠9,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-09 10:28:48', '2024-08-09 10:28:48', NULL, 0, 0);
INSERT INTO `record` VALUES (182, 47, 43, 9372, 145, 9517, 0, 'D09AF4B0B2DB090AD6703602C2E9528D', 'Banker[6] T[10] D[0] X[3] H[3] 中奖区域:[1 0 0 0]', '百人牛妞', '2024-08-09 10:28:58', '2024-08-09 10:28:58', NULL, 0, 0);
INSERT INTO `record` VALUES (183, 44, 42, 5000457, -287, 5000170, 0, 'BE5130C25506C2A1D6810D0720B8BD1C', '闲:[♥2,♣7,] 庄:[♥Q,♠Q,] 中奖区域:[1 0 0 1 0 0 0 1]', '百佳乐', '2024-08-09 10:29:09', '2024-08-09 10:29:09', NULL, 0, 0);
INSERT INTO `record` VALUES (184, 47, 43, 9517, 142, 9659, 0, 'BB57F57A30C9CB0D79DC7E758F82C7AA', 'Banker[4] T[0] D[0] X[0] H[6] 中奖区域:[0 0 0 1]', '百人牛妞', '2024-08-09 10:29:19', '2024-08-09 10:29:19', NULL, 0, 0);
INSERT INTO `record` VALUES (185, 44, 42, 5000170, -481, 4999689, 0, '69149D175E25DECEF0B94A335547221E', '闲:[♠9,♥4,♥3,] 庄:[♠7,♣Q,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-09 10:29:30', '2024-08-09 10:29:30', NULL, 0, 0);
INSERT INTO `record` VALUES (186, 47, 43, 9659, 139, 9798, 0, '74B5EBBAC8A2497F344263D0449C8437', 'Banker[0] T[5] D[1] X[0] H[7] 中奖区域:[1 1 0 1]', '百人牛妞', '2024-08-09 10:29:40', '2024-08-09 10:29:40', NULL, 0, 0);
INSERT INTO `record` VALUES (187, 44, 42, 4999689, -528, 4999161, 0, '3F3DCB80256641D8DF1516253BA62FE7', '闲:[♣4,♣J,♦K,] 庄:[♣1,♦5,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-09 10:29:51', '2024-08-09 10:29:51', NULL, 0, 0);
INSERT INTO `record` VALUES (188, 47, 43, 9798, 136, 9934, 0, '48917ABCA6E52FF5CF043EF7BDFFA609', 'Banker[5] T[0] D[8] X[2] H[0] 中奖区域:[0 1 0 0]', '百人牛妞', '2024-08-09 10:30:01', '2024-08-09 10:30:01', NULL, 0, 0);
INSERT INTO `record` VALUES (189, 44, 42, 4999161, -560, 4998601, 0, '0FA1C0FE08014C55108FC73837B15ED4', '闲:[♥5,♥4,] 庄:[♦8,♠5,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-09 10:30:12', '2024-08-09 10:30:12', NULL, 0, 0);
INSERT INTO `record` VALUES (190, 47, 43, 9934, 133, 10067, 0, 'B101A1F467BEC4A31AEF08E133789DDC', 'Banker[7] T[3] D[6] X[5] H[2] 中奖区域:[0 0 0 0]', '百人牛妞', '2024-08-09 10:30:22', '2024-08-09 10:30:22', NULL, 0, 0);
INSERT INTO `record` VALUES (191, 44, 42, 4998601, -707, 4997894, 0, '85DE321DEB50E05AA84F3DE4B9A635FE', '闲:[♠K,♥K,] 庄:[♣4,♥5,] 中奖区域:[0 0 1 0 1 0 1 0]', '百佳乐', '2024-08-09 10:30:33', '2024-08-09 10:30:33', NULL, 0, 0);
INSERT INTO `record` VALUES (192, 47, 43, 10067, 130, 10197, 0, '7C1491BC2F2D4155796E8B870902DA6F', 'Banker[7] T[0] D[0] X[0] H[5] 中奖区域:[0 0 0 0]', '百人牛妞', '2024-08-09 10:30:43', '2024-08-09 10:30:43', NULL, 0, 0);
INSERT INTO `record` VALUES (193, 44, 42, 4997894, -858, 4997036, 0, 'FF0A0ED405CE229CCF117A56A1D02EB8', '闲:[♠8,♦4,♣1,] 庄:[♥4,♦10,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-09 10:30:55', '2024-08-09 10:30:55', NULL, 0, 0);
INSERT INTO `record` VALUES (194, 47, 43, 10197, 127, 10324, 0, '83594C26702B80513606FC5A212AA3F8', 'Banker[0] T[0] D[0] X[5] H[2] 中奖区域:[0 0 1 1]', '百人牛妞', '2024-08-09 10:31:04', '2024-08-09 10:31:04', NULL, 0, 0);
INSERT INTO `record` VALUES (195, 44, 42, 4997036, -878, 4996158, 0, '8E2B4E436BDFE0E30C61A465B39A3FAC', '闲:[♣9,♦4,] 庄:[♣9,♥Q,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-09 10:31:16', '2024-08-09 10:31:16', NULL, 0, 0);
INSERT INTO `record` VALUES (196, 47, 43, 10324, 124, 10448, 0, 'AB8C18683984D2C9D8F5780B49980E56', 'Banker[7] T[3] D[5] X[1] H[3] 中奖区域:[0 0 0 0]', '百人牛妞', '2024-08-09 10:31:26', '2024-08-09 10:31:26', NULL, 0, 0);
INSERT INTO `record` VALUES (197, 44, 42, 4996158, -958, 4995200, 0, 'A4BDA0E89F8147D23D6DE285CF30ABF0', '闲:[♠K,♣1,♠10,] 庄:[♥10,♦4,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-09 10:31:37', '2024-08-09 10:31:37', NULL, 0, 0);
INSERT INTO `record` VALUES (198, 47, 43, 10448, 121, 10569, 0, '2E3F00D87224F6F1135A08393393325E', 'Banker[0] T[7] D[1] X[0] H[0] 中奖区域:[1 1 0 0]', '百人牛妞', '2024-08-09 10:31:47', '2024-08-09 10:31:47', NULL, 0, 0);
INSERT INTO `record` VALUES (199, 44, 42, 4995200, -1127, 4994073, 0, 'D9D1E99CFCBCF2E28D08299F641447B2', '闲:[♣7,♦8,] 庄:[♥8,♠J,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-09 10:31:58', '2024-08-09 10:31:58', NULL, 0, 0);
INSERT INTO `record` VALUES (200, 47, 43, 10569, 118, 10687, 0, '0BE15061A49D0D5061F0D08C295AC46D', 'Banker[4] T[6] D[0] X[6] H[2] 中奖区域:[1 0 1 0]', '百人牛妞', '2024-08-09 10:32:08', '2024-08-09 10:32:08', NULL, 0, 0);
INSERT INTO `record` VALUES (201, 44, 42, 4994073, -983, 4993090, 0, '456BE95E6E8CD56434FF2F3CD6760E7A', '闲:[♥3,♦1,♦9,] 庄:[♦1,♣J,♥9,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-09 10:32:19', '2024-08-09 10:32:19', NULL, 0, 0);
INSERT INTO `record` VALUES (202, 47, 43, 10687, 115, 10802, 0, 'E7B930B600716561C317241A55AC31BE', 'Banker[5] T[2] D[6] X[5] H[0] 中奖区域:[0 1 1 0]', '百人牛妞', '2024-08-09 10:32:29', '2024-08-09 10:32:29', NULL, 0, 0);
INSERT INTO `record` VALUES (203, 44, 42, 4993090, -830, 4992260, 0, 'B0363768290FBCDDD71AD9A44E6B9224', '闲:[♠10,♥J,♣4,] 庄:[♥K,♠J,♠K,] 中奖区域:[1 0 0 0 0 0 0 1]', '百佳乐', '2024-08-09 10:32:40', '2024-08-09 10:32:40', NULL, 0, 0);
INSERT INTO `record` VALUES (204, 47, 43, 10802, 112, 10914, 0, '6CCE018A475BE6976ECD60C90D806EA9', 'Banker[8] T[3] D[5] X[1] H[10] 中奖区域:[0 0 0 1]', '百人牛妞', '2024-08-09 10:32:50', '2024-08-09 10:32:50', NULL, 0, 0);
INSERT INTO `record` VALUES (205, 47, 44, 10914, -178, 10736, 0, '2BFF7798B8CE5931AC3EED07B4F85657', 'Banker[9] T[6] D[7] X[7] H[5] 中奖区域:[0 0 0 0]', '百人牛妞', '2024-08-09 11:54:32', '2024-08-09 11:54:32', NULL, 0, 0);
INSERT INTO `record` VALUES (206, 47, 44, 10736, -178, 10558, 0, '0AC855E9412D167CA7B4074605B5416E', 'Banker[0] T[4] D[9] X[10] H[5] 中奖区域:[1 1 1 1]', '百人牛妞', '2024-08-09 11:54:53', '2024-08-09 11:54:53', NULL, 0, 0);
INSERT INTO `record` VALUES (207, 47, 44, 10558, -191, 10367, 0, '811B9A8A0EBC87997E4794DF7F7FADC6', 'Banker[4] T[0] D[8] X[5] H[2] 中奖区域:[0 1 1 0]', '百人牛妞', '2024-08-09 11:55:14', '2024-08-09 11:55:14', NULL, 0, 0);
INSERT INTO `record` VALUES (208, 47, 44, 10367, -191, 10176, 0, '01C9C409BF8DB9DC08511AF57235BA49', 'Banker[0] T[6] D[4] X[6] H[6] 中奖区域:[1 1 1 1]', '百人牛妞', '2024-08-09 11:55:35', '2024-08-09 11:55:35', NULL, 0, 0);
INSERT INTO `record` VALUES (209, 47, 44, 10176, -191, 9985, 0, 'C4AB21AE2D68295CA35861ADBE1E126C', 'Banker[0] T[0] D[10] X[3] H[0] 中奖区域:[1 1 1 1]', '百人牛妞', '2024-08-09 11:55:56', '2024-08-09 11:55:56', NULL, 0, 0);
INSERT INTO `record` VALUES (210, 47, 44, 9985, -191, 9794, 0, 'D04FF78072EAE53EC4C1FA1E3A31C963', 'Banker[8] T[9] D[0] X[2] H[0] 中奖区域:[1 0 0 0]', '百人牛妞', '2024-08-09 11:56:17', '2024-08-09 11:56:17', NULL, 0, 0);
INSERT INTO `record` VALUES (211, 47, 44, 9794, -191, 9603, 0, '297908A61AC45D917053223ECB2ABA05', 'Banker[0] T[9] D[10] X[10] H[0] 中奖区域:[1 1 1 1]', '百人牛妞', '2024-08-09 11:56:38', '2024-08-09 11:56:38', NULL, 0, 0);
INSERT INTO `record` VALUES (212, 47, 44, 9603, -191, 9412, 0, '31D86DE4C7B4B042C3ED6689881A67AB', 'Banker[0] T[0] D[0] X[0] H[3] 中奖区域:[1 1 1 1]', '百人牛妞', '2024-08-09 11:56:59', '2024-08-09 11:56:59', NULL, 0, 0);
INSERT INTO `record` VALUES (213, 47, 44, 9412, -191, 9221, 0, 'A69DA8FEBD0D842C82156695CF2F79F8', 'Banker[0] T[8] D[0] X[0] H[4] 中奖区域:[1 0 0 1]', '百人牛妞', '2024-08-09 11:57:20', '2024-08-09 11:57:20', NULL, 0, 0);
INSERT INTO `record` VALUES (214, 47, 44, 9221, -191, 9030, 0, '2AD9B766302767D6E8D7D297DB692D29', 'Banker[1] T[0] D[0] X[2] H[0] 中奖区域:[0 0 1 0]', '百人牛妞', '2024-08-09 11:57:41', '2024-08-09 11:57:41', NULL, 0, 0);
INSERT INTO `record` VALUES (215, 47, 44, 9030, -191, 8839, 0, '3F8C877AF98EB374F680DB14F3024F9A', 'Banker[0] T[7] D[6] X[1] H[8] 中奖区域:[1 1 1 1]', '百人牛妞', '2024-08-09 11:58:02', '2024-08-09 11:58:02', NULL, 0, 0);
INSERT INTO `record` VALUES (216, 47, 44, 8839, -191, 8648, 0, '5AA232DDA902CB4D38C0580DCD85B3A7', 'Banker[0] T[0] D[6] X[5] H[7] 中奖区域:[0 1 1 1]', '百人牛妞', '2024-08-09 11:58:23', '2024-08-09 11:58:23', NULL, 0, 0);
INSERT INTO `record` VALUES (217, 47, 44, 8648, -191, 8457, 0, 'B9B8ED3273D86FFBFA9491C83842CA92', 'Banker[6] T[7] D[10] X[0] H[9] 中奖区域:[1 1 0 1]', '百人牛妞', '2024-08-09 11:58:45', '2024-08-09 11:58:45', NULL, 0, 0);
INSERT INTO `record` VALUES (218, 47, 44, 8457, -191, 8266, 0, '8D9B4BDFC86B5B2C00B1BA5171ED2202', 'Banker[10] T[3] D[0] X[0] H[1] 中奖区域:[0 0 0 0]', '百人牛妞', '2024-08-09 11:59:06', '2024-08-09 11:59:06', NULL, 0, 0);
INSERT INTO `record` VALUES (219, 47, 44, 8266, -191, 8075, 0, 'C7B4607BD917578E00C784CD7A508B4A', 'Banker[4] T[3] D[1] X[10] H[6] 中奖区域:[0 0 1 1]', '百人牛妞', '2024-08-09 11:59:27', '2024-08-09 11:59:27', NULL, 0, 0);
INSERT INTO `record` VALUES (220, 47, 44, 8075, -191, 7884, 0, '69D9B8E0AA2156E2E8CCB6D917533AAF', 'Banker[3] T[3] D[7] X[0] H[6] 中奖区域:[1 1 0 1]', '百人牛妞', '2024-08-09 11:59:48', '2024-08-09 11:59:48', NULL, 0, 0);
INSERT INTO `record` VALUES (221, 47, 44, 7884, -191, 7693, 0, '7C2A58E55C3CB48AEDC54CC917A7C840', 'Banker[0] T[8] D[0] X[2] H[2] 中奖区域:[1 0 1 1]', '百人牛妞', '2024-08-09 12:00:09', '2024-08-09 12:00:09', NULL, 0, 0);
INSERT INTO `record` VALUES (222, 47, 44, 7693, -191, 7502, 0, '06608BED1B53A1C1C5564909F987D2AE', 'Banker[10] T[1] D[4] X[0] H[10] 中奖区域:[0 0 0 0]', '百人牛妞', '2024-08-09 12:00:30', '2024-08-09 12:00:30', NULL, 0, 0);
INSERT INTO `record` VALUES (223, 47, 44, 7502, -191, 7311, 0, '9D63B1E2FD75243213A562670F1B421A', 'Banker[3] T[0] D[5] X[1] H[0] 中奖区域:[0 1 0 0]', '百人牛妞', '2024-08-09 12:00:51', '2024-08-09 12:00:51', NULL, 0, 0);
INSERT INTO `record` VALUES (224, 47, 44, 7311, -191, 7120, 0, '4E3BF67150B9C992A912831E997AE6D6', 'Banker[9] T[0] D[1] X[10] H[3] 中奖区域:[0 0 1 0]', '百人牛妞', '2024-08-09 12:01:12', '2024-08-09 12:01:12', NULL, 0, 0);
INSERT INTO `record` VALUES (225, 44, 42, 4992260, 115, 4992375, 0, '1481C340677A59641BC1EED1540BC514', '闲:[♠5,♥3,] 庄:[♥1,♣10,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-10 13:02:50', '2024-08-10 13:02:50', NULL, 0, 0);
INSERT INTO `record` VALUES (226, 44, 42, 4992375, 27, 4992402, 0, '5C6EE29C9F289A9AC16AF417897F67A2', '闲:[♠10,♣5,♦2,] 庄:[♣J,♦K,♥Q,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-10 17:54:06', '2024-08-10 17:54:06', NULL, 0, 0);
INSERT INTO `record` VALUES (227, 44, 42, 4992402, -166, 4992236, 0, '2D07AEADA72E994F59827407D73D63EB', '闲:[♠9,♥Q,] 庄:[♥6,♥1,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-12 19:45:51', '2024-08-12 19:45:51', NULL, 0, 0);
INSERT INTO `record` VALUES (228, 44, 42, 10000, -49, 9951, 0, 'F8A679676511BD5E1A34B7DAB090372F', '闲:[♣2,♦2,♠Q,] 庄:[♥Q,♠K,♣5,] 中奖区域:[1 0 0 0 0 0 1 0]', '百佳乐', '2024-08-27 16:53:11', '2024-08-27 16:53:11', NULL, 0, 0);
INSERT INTO `record` VALUES (229, 44, 42, 9951, -184, 9767, 0, '7DE4A6D50BB96EB15890BF93348553B8', '闲:[♥Q,♦K,] 庄:[♠Q,♦9,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 16:53:32', '2024-08-27 16:53:32', NULL, 0, 0);
INSERT INTO `record` VALUES (230, 44, 42, 9767, -193, 9574, 0, 'FB0407221656DF77D07BB932096A90E9', '闲:[♦K,♥5,♥5,] 庄:[♥7,♦10,] 中奖区域:[0 0 1 0 0 0 1 0]', '百佳乐', '2024-08-27 16:54:14', '2024-08-27 16:54:14', NULL, 0, 0);
INSERT INTO `record` VALUES (231, 44, 42, 9574, -71, 9503, 0, '46E1F9E9F27B70EA1A5FA99F4C9F7EE9', '闲:[♦5,♣2,] 庄:[♣8,♠8,] 中奖区域:[1 0 0 0 0 0 0 1]', '百佳乐', '2024-08-27 16:54:35', '2024-08-27 16:54:35', NULL, 0, 0);
INSERT INTO `record` VALUES (232, 44, 42, 9503, -228, 9275, 0, '3CBAFB6EFB179051ED40D049FDB8A0F9', '闲:[♣10,♣1,] 庄:[♣K,♥9,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 16:54:56', '2024-08-27 16:54:56', NULL, 0, 0);
INSERT INTO `record` VALUES (233, 44, 42, 9275, -293, 8982, 0, '878C219954EDBD388077803431104F2D', '闲:[♠3,♥9,♠8,] 庄:[♥3,♠2,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 16:55:17', '2024-08-27 16:55:17', NULL, 0, 0);
INSERT INTO `record` VALUES (234, 44, 42, 8982, -474, 8508, 0, '7F820AC19CB885BD6CBDE3F295A8EEC2', '闲:[♠3,♥4,] 庄:[♥6,♠K,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 16:55:38', '2024-08-27 16:55:38', NULL, 0, 0);
INSERT INTO `record` VALUES (235, 44, 42, 8508, -529, 7979, 0, '3566930340337C6DCFEBF461F5CC0F88', '闲:[♠K,♥Q,♦9,] 庄:[♠1,♦9,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 16:55:59', '2024-08-27 16:55:59', NULL, 0, 0);
INSERT INTO `record` VALUES (236, 44, 42, 7979, -690, 7289, 0, '9FF317C9CAF25E70B5D011C5DFFBD1B7', '闲:[♣6,♥9,♠5,] 庄:[♣4,♥9,♥1,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 16:56:21', '2024-08-27 16:56:21', NULL, 0, 0);
INSERT INTO `record` VALUES (237, 44, 42, 7289, -704, 6585, 0, 'E439B81F95FFDA31A378599D99E453DB', '闲:[♥6,♠K,] 庄:[♥4,♦6,♣K,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 16:56:42', '2024-08-27 16:56:42', NULL, 0, 0);
INSERT INTO `record` VALUES (238, 44, 42, 6585, -876, 5709, 0, '4E5942CD7ACF43DF39B6A9A9834049F3', '闲:[♣8,♠8,] 庄:[♣Q,♠K,♦3,] 中奖区域:[1 0 0 0 0 0 1 0]', '百佳乐', '2024-08-27 16:57:03', '2024-08-27 16:57:03', NULL, 0, 0);
INSERT INTO `record` VALUES (239, 44, 42, 5709, -915, 4794, 0, '22FBF2C44F77894197A4DC4A0AC02AB6', '闲:[♦Q,♠7,] 庄:[♦8,♠4,♠K,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 16:57:24', '2024-08-27 16:57:24', NULL, 0, 0);
INSERT INTO `record` VALUES (240, 44, 42, 4794, -778, 4016, 0, 'C39E94911039080E8BAA0737C774E3FB', '闲:[♣9,♦4,♦5,] 庄:[♥Q,♦6,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 16:57:45', '2024-08-27 16:57:45', NULL, 0, 0);
INSERT INTO `record` VALUES (241, 44, 42, 4016, -941, 3075, 0, 'FAFB7D3A830BEC77E6EE0576FFFC9DAA', '闲:[♣4,♠4,] 庄:[♣6,♣3,] 中奖区域:[0 0 1 0 1 0 1 0]', '百佳乐', '2024-08-27 16:58:06', '2024-08-27 16:58:06', NULL, 0, 0);
INSERT INTO `record` VALUES (242, 44, 42, 3075, -987, 2088, 0, '7A835375F483F4EC0497902330341554', '闲:[♥4,♣1,♥7,] 庄:[♠J,♦7,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 16:58:27', '2024-08-27 16:58:27', NULL, 0, 0);
INSERT INTO `record` VALUES (243, 44, 42, 2088, -1122, 966, 0, 'CCB7A58606CA3622B0E9AD1A55179240', '闲:[♦6,♠6,♠7,] 庄:[♠10,♠2,] 中奖区域:[1 0 0 1 0 0 1 0]', '百佳乐', '2024-08-27 16:58:48', '2024-08-27 16:58:48', NULL, 0, 0);
INSERT INTO `record` VALUES (244, 44, 42, 966, -1267, -301, 0, 'D02F28D36AB6FFC627ED864867555B4A', '闲:[♦8,♥K,] 庄:[♣Q,♣9,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 16:59:09', '2024-08-27 16:59:09', NULL, 0, 0);
INSERT INTO `record` VALUES (245, 44, 42, -301, -1301, -1602, 0, 'A403732345F1818D9C00F4A115321147', '闲:[♥5,♠K,♣9,] 庄:[♠K,♠Q,♠7,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 16:59:30', '2024-08-27 16:59:30', NULL, 0, 0);
INSERT INTO `record` VALUES (246, 44, 42, -1602, -1463, -3065, 0, 'A09B3CD2B303D9A3C800818C17A761F5', '闲:[♦6,♠Q,] 庄:[♥7,♦10,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 16:59:51', '2024-08-27 16:59:51', NULL, 0, 0);
INSERT INTO `record` VALUES (247, 44, 42, -3065, -1590, -4655, 0, 'F15A47EF2ACEE6AF23A673F213F919CC', '闲:[♦5,♦8,♣1,] 庄:[♠3,♦7,♥Q,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 17:00:12', '2024-08-27 17:00:12', NULL, 0, 0);
INSERT INTO `record` VALUES (248, 44, 42, -4655, -1731, -6386, 0, 'C305B2908B2C6C7C3928446397997C7C', '闲:[♣K,♣3,♦J,] 庄:[♦10,♣4,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 17:00:33', '2024-08-27 17:00:33', NULL, 0, 0);
INSERT INTO `record` VALUES (249, 44, 42, -6386, -1546, -7932, 0, '3E66575CEAEE2B25470419C0C15D360B', '闲:[♦9,♦J,] 庄:[♦K,♦7,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 17:00:54', '2024-08-27 17:00:54', NULL, 0, 0);
INSERT INTO `record` VALUES (250, 44, 42, -7932, -1399, -9331, 0, 'FD2BD41C04E4E062F13B890CD96CD354', '闲:[♦10,♣3,♦3,] 庄:[♥Q,♥1,♠1,] 中奖区域:[1 0 0 0 0 0 1 1]', '百佳乐', '2024-08-27 17:01:15', '2024-08-27 17:01:15', NULL, 0, 0);
INSERT INTO `record` VALUES (251, 44, 42, -9331, -1584, -10915, 0, '5319032AA81956B8056E29CAAA06B983', '闲:[♣8,♣6,♥4,] 庄:[♥K,♥6,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 17:01:36', '2024-08-27 17:01:36', NULL, 0, 0);
INSERT INTO `record` VALUES (252, 44, 42, -10915, -1477, -12392, 0, 'A2D1AE8B73F75437A6F6E879EA075AED', '闲:[♣10,♣6,] 庄:[♥J,♣10,♠2,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 17:01:57', '2024-08-27 17:01:57', NULL, 0, 0);
INSERT INTO `record` VALUES (253, 44, 42, -12392, -1654, -14046, 0, '5916D0AE2EAF08B60D6CE7A37CD603AD', '闲:[♦2,♥10,♦Q,] 庄:[♥2,♥4,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 17:02:19', '2024-08-27 17:02:19', NULL, 0, 0);
INSERT INTO `record` VALUES (254, 44, 42, -14046, -1695, -15741, 0, '8AA11F522838C47921CDCF3E9643DB98', '闲:[♥K,♠6,] 庄:[♦3,♦1,♥7,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 17:02:40', '2024-08-27 17:02:40', NULL, 0, 0);
INSERT INTO `record` VALUES (255, 44, 42, -15741, -1673, -17414, 0, 'FF93D30991FE0EDEC6B16F4886FBB09E', '闲:[♣7,♦K,] 庄:[♣4,♠K,♣8,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 17:03:01', '2024-08-27 17:03:01', NULL, 0, 0);
INSERT INTO `record` VALUES (256, 44, 42, -17414, -1851, -19265, 0, '3C6C75FF705B5CF8F11CEB15B88CCA2F', '闲:[♥2,♣1,♣9,] 庄:[♣9,♥4,♣J,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 17:03:22', '2024-08-27 17:03:22', NULL, 0, 0);
INSERT INTO `record` VALUES (257, 44, 42, -19265, -1723, -20988, 0, 'ABB13252F61D47D6B3F0B1B79906C61B', '闲:[♥4,♣4,] 庄:[♣2,♦Q,] 中奖区域:[1 0 0 1 0 0 1 0]', '百佳乐', '2024-08-27 17:03:43', '2024-08-27 17:03:43', NULL, 0, 0);
INSERT INTO `record` VALUES (258, 44, 42, -20988, -1769, -22757, 0, 'B007C5B318645FE391635C16640B10DA', '闲:[♦7,♣6,♣6,] 庄:[♠4,♦J,] 中奖区域:[1 0 0 1 0 0 1 0]', '百佳乐', '2024-08-27 17:04:04', '2024-08-27 17:04:04', NULL, 0, 0);
INSERT INTO `record` VALUES (259, 44, 42, -22757, -1769, -24526, 0, '8FB1B15805C9FA06E0A1B8914E2A26CB', '闲:[♥8,♠1,] 庄:[♦Q,♦9,] 中奖区域:[0 1 0 0 0 0 0 0]', '百佳乐', '2024-08-27 17:04:25', '2024-08-27 17:04:25', NULL, 0, 0);
INSERT INTO `record` VALUES (260, 44, 42, -24526, -337, -24863, 0, 'EFF361E0D17AF780E2DF55BEC58DC654', '闲:[♠6,♥8,♥8,] 庄:[♠3,♦9,♠1,] 中奖区域:[0 1 0 0 0 0 1 0]', '百佳乐', '2024-08-27 17:04:46', '2024-08-27 17:04:46', NULL, 0, 0);
INSERT INTO `record` VALUES (261, 44, 42, -24863, -421, -25284, 0, '82CE371FE9C96ADF89FCA45C4608A0A2', '闲:[♠5,♥7,♠6,] 庄:[♣2,♠1,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 17:05:07', '2024-08-27 17:05:07', NULL, 0, 0);
INSERT INTO `record` VALUES (262, 44, 42, -25284, -441, -25725, 0, '800460DD00C2A1B8030571DB7DF69FB5', '闲:[♠2,♠4,] 庄:[♦7,♦J,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 17:05:28', '2024-08-27 17:05:28', NULL, 0, 0);
INSERT INTO `record` VALUES (263, 44, 42, -25725, -474, -26199, 0, '868D9ED9C40B0E0EB92B7FA5923486F6', '闲:[♣6,♦Q,] 庄:[♣5,♥4,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 17:05:49', '2024-08-27 17:05:49', NULL, 0, 0);
INSERT INTO `record` VALUES (264, 44, 42, -26199, -661, -26860, 0, '3AF18A6A4D74E8756345D05F26B95585', '闲:[♥4,♠3,] 庄:[♥4,♥9,♦2,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 17:06:10', '2024-08-27 17:06:10', NULL, 0, 0);
INSERT INTO `record` VALUES (265, 44, 42, -26860, -661, -27521, 0, '9E48BE2E1A72B5052CC6CF881F3A6876', '闲:[♥8,♣8,] 庄:[♥K,♥6,] 中奖区域:[0 1 0 0 0 0 1 0]', '百佳乐', '2024-08-27 17:06:31', '2024-08-27 17:06:31', NULL, 0, 0);
INSERT INTO `record` VALUES (266, 44, 42, -27521, -599, -28120, 0, '7E3B58D113788EBFF96C72C660C2609E', '闲:[♦2,♥6,] 庄:[♦10,♠4,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 17:06:52', '2024-08-27 17:06:52', NULL, 0, 0);
INSERT INTO `record` VALUES (267, 44, 42, -28120, -490, -28610, 0, '8DD138BA14B8FB7F479AC6FBE4C6FA18', '闲:[♥J,♥9,] 庄:[♦5,♣2,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 17:07:13', '2024-08-27 17:07:13', NULL, 0, 0);
INSERT INTO `record` VALUES (268, 44, 42, -28610, -571, -29181, 0, 'FBAA406144F898EDC929B62599DB092D', '闲:[♠7,♥6,♥3,] 庄:[♣1,♦J,♦Q,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 17:07:34', '2024-08-27 17:07:34', NULL, 0, 0);
INSERT INTO `record` VALUES (269, 44, 42, -29181, -578, -29759, 0, '96AB02DC91B70A4E824E9A66BCD22DAD', '闲:[♥3,♥3,] 庄:[♦J,♠K,♦3,] 中奖区域:[1 0 0 0 0 0 1 0]', '百佳乐', '2024-08-27 17:07:55', '2024-08-27 17:07:55', NULL, 0, 0);
INSERT INTO `record` VALUES (270, 44, 42, -29759, -712, -30471, 0, '4B4965D0F04CBA78686562AEA8C254B6', '闲:[♠9,♣6,] 庄:[♣8,♠Q,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 17:08:16', '2024-08-27 17:08:16', NULL, 0, 0);
INSERT INTO `record` VALUES (271, 44, 42, -30471, -722, -31193, 0, '2DDBBBD2E0E6A204E32578E62E564DAA', '闲:[♣6,♦5,♠1,] 庄:[♠Q,♦6,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 17:08:38', '2024-08-27 17:08:38', NULL, 0, 0);
INSERT INTO `record` VALUES (272, 44, 42, -31193, -687, -31880, 0, '0F66B184B03D74A93E9FA8CEBEA8FAB3', '闲:[♣6,♠2,] 庄:[♦K,♥2,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 17:08:59', '2024-08-27 17:08:59', NULL, 0, 0);
INSERT INTO `record` VALUES (273, 44, 42, -31880, -681, -32561, 0, '3AA931D451C730565E46AE5E1CB38F47', '闲:[♣9,♥K,] 庄:[♣6,♣10,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 17:09:20', '2024-08-27 17:09:20', NULL, 0, 0);
INSERT INTO `record` VALUES (274, 44, 42, -32561, -185, -32746, 0, 'D5221D572CF89641CABED81BA18E28ED', '闲:[♠1,♦4,♥J,] 庄:[♠K,♣5,] 中奖区域:[0 1 0 0 0 0 0 0]', '百佳乐', '2024-08-27 17:09:41', '2024-08-27 17:09:41', NULL, 0, 0);
INSERT INTO `record` VALUES (275, 44, 42, -32746, -291, -33037, 0, 'E656FA21CDB5CB5591CAEC9AE322B30E', '闲:[♣2,♣Q,♠Q,] 庄:[♦K,♥7,] 中奖区域:[0 0 1 0 0 0 1 0]', '百佳乐', '2024-08-27 17:10:02', '2024-08-27 17:10:02', NULL, 0, 0);
INSERT INTO `record` VALUES (276, 44, 42, -33037, -414, -33451, 0, 'C21422FEF843F113CB7E447A718B9C80', '闲:[♥5,♦9,] 庄:[♣10,♣9,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 17:10:23', '2024-08-27 17:10:23', NULL, 0, 0);
INSERT INTO `record` VALUES (277, 44, 42, -33451, -450, -33901, 0, '05920F810BB9D68703DBEA44D4451257', '闲:[♣7,♣1,] 庄:[♠9,♥6,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 17:10:44', '2024-08-27 17:10:44', NULL, 0, 0);
INSERT INTO `record` VALUES (278, 44, 42, -33901, -575, -34476, 0, '934826AE5E924B51BA7F60E262E41D8D', '闲:[♣8,♦K,] 庄:[♣10,♥4,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 17:11:05', '2024-08-27 17:11:05', NULL, 0, 0);
INSERT INTO `record` VALUES (279, 44, 42, -34476, 128, -34348, 0, '1C6B9AB52DB4228E714730FAA7F316DC', '闲:[♦4,♣1,♠J,] 庄:[♣5,♥Q,] 中奖区域:[0 1 0 0 0 0 0 0]', '百佳乐', '2024-08-27 17:11:26', '2024-08-27 17:11:26', NULL, 0, 0);
INSERT INTO `record` VALUES (280, 44, 42, -34348, 1042, -33306, 0, '0449E7E7F37BEBBAA9E0A3BEC3876807', '闲:[♦2,♣1,♣3,] 庄:[♦8,♠8,] 中奖区域:[0 1 0 0 0 0 0 1]', '百佳乐', '2024-08-27 17:11:47', '2024-08-27 17:11:47', NULL, 0, 0);
INSERT INTO `record` VALUES (281, 44, 42, -33306, 1203, -32103, 0, '8AEEB47529197C51DD279FDE3B05635A', '闲:[♥1,♥Q,♦6,] 庄:[♠9,♣1,♣3,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 17:12:08', '2024-08-27 17:12:08', NULL, 0, 0);
INSERT INTO `record` VALUES (282, 44, 42, -32103, 1059, -31044, 0, 'C78CA092175C6472C1438B983226D839', '闲:[♠Q,♦10,♥J,] 庄:[♣J,♥7,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 17:12:29', '2024-08-27 17:12:29', NULL, 0, 0);
INSERT INTO `record` VALUES (283, 44, 42, -31044, 1042, -30002, 0, 'A55EE8AC7A53379B6A7340B6151BD34D', '闲:[♠4,♦2,] 庄:[♦8,♦Q,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 17:12:50', '2024-08-27 17:12:50', NULL, 0, 0);
INSERT INTO `record` VALUES (284, 44, 42, -30002, 1003, -28999, 0, '4B4210960AFB772FDC646C50D1084A52', '闲:[♣J,♦2,] 庄:[♥8,♥Q,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 17:13:11', '2024-08-27 17:13:11', NULL, 0, 0);
INSERT INTO `record` VALUES (285, 44, 42, -28999, 822, -28177, 0, 'C723461F3962346ADBF0B676A377B7C9', '闲:[♠Q,♣6,] 庄:[♥K,♣8,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 17:13:32', '2024-08-27 17:13:32', NULL, 0, 0);
INSERT INTO `record` VALUES (286, 44, 42, -28177, 762, -27415, 0, 'FF1579228881F87F2527A03F05AC5DA0', '闲:[♦10,♣J,♠10,] 庄:[♥2,♣1,♦6,] 中奖区域:[0 0 1 0 0 0 1 0]', '百佳乐', '2024-08-27 17:13:53', '2024-08-27 17:13:53', NULL, 0, 0);
INSERT INTO `record` VALUES (287, 44, 42, -27415, 796, -26619, 0, 'CA4601FA3DF49D132A7E54FF1B76D8A4', '闲:[♦5,♣2,] 庄:[♦9,♣5,♠Q,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 17:14:14', '2024-08-27 17:14:14', NULL, 0, 0);
INSERT INTO `record` VALUES (288, 44, 42, -26619, 792, -25827, 0, 'B5242C9B1E332E6BFF4BADB8C66EDCC3', '闲:[♥1,♥1,♥3,] 庄:[♥7,♥8,] 中奖区域:[0 1 0 0 0 0 1 0]', '百佳乐', '2024-08-27 17:14:36', '2024-08-27 17:14:36', NULL, 0, 0);
INSERT INTO `record` VALUES (289, 44, 42, -25827, 867, -24960, 0, '7F679898368E100841680BD7878171B6', '闲:[♥7,♠2,] 庄:[♣6,♦1,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 17:14:57', '2024-08-27 17:14:57', NULL, 0, 0);
INSERT INTO `record` VALUES (290, 44, 42, -24960, 861, -24099, 0, '830CA6D870E584FB207BD4B28ECDA9FA', '闲:[♣10,♣3,] 庄:[♦8,♦Q,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 17:15:18', '2024-08-27 17:15:18', NULL, 0, 0);
INSERT INTO `record` VALUES (291, 44, 42, -24099, 819, -23280, 0, 'BAF7DD1A9891770D69AE97F1766FE82C', '闲:[♦9,♣2,♦J,] 庄:[♣10,♣6,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 17:15:39', '2024-08-27 17:15:39', NULL, 0, 0);
INSERT INTO `record` VALUES (292, 44, 42, -23280, 683, -22597, 0, '3058DDC3F2A74C7DB0A381752650EF47', '闲:[♣4,♠7,♠2,] 庄:[♠K,♣4,♥6,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 17:16:00', '2024-08-27 17:16:00', NULL, 0, 0);
INSERT INTO `record` VALUES (293, 44, 42, -22597, 517, -22080, 0, 'DAE8457360C8B0FE1F2830146F61E2BE', '闲:[♣5,♥3,] 庄:[♠4,♣5,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 17:16:21', '2024-08-27 17:16:21', NULL, 0, 0);
INSERT INTO `record` VALUES (294, 44, 42, -22080, 460, -21620, 0, '45D54B642453B4EE9D102B45E58907D9', '闲:[♣10,♠5,] 庄:[♦8,♥1,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 17:16:42', '2024-08-27 17:16:42', NULL, 0, 0);
INSERT INTO `record` VALUES (295, 44, 42, -21620, 345, -21275, 0, 'D0104D21221E7926AAC2190108CF717A', '闲:[♠2,♠6,] 庄:[♣Q,♣7,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 17:17:03', '2024-08-27 17:17:03', NULL, 0, 0);
INSERT INTO `record` VALUES (296, 44, 42, -21275, 231, -21044, 0, '42C199BC1BF9C9EFCF35DE2787E52505', '闲:[♦3,♣K,] 庄:[♣3,♥6,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 17:17:24', '2024-08-27 17:17:24', NULL, 0, 0);
INSERT INTO `record` VALUES (297, 44, 42, -21044, 48, -20996, 0, 'AB3F9AFB793E90714DF616D70EC3B0E6', '闲:[♣J,♣3,♠2,] 庄:[♠9,♠3,♣1,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 17:17:45', '2024-08-27 17:17:45', NULL, 0, 0);
INSERT INTO `record` VALUES (298, 44, 42, -20996, 11, -20985, 0, '2CD3F96011A5A462D24B0CA84CB96B76', '闲:[♥J,♠1,♣8,] 庄:[♦10,♥10,] 中奖区域:[1 0 0 1 0 0 0 1]', '百佳乐', '2024-08-27 17:18:06', '2024-08-27 17:18:06', NULL, 0, 0);
INSERT INTO `record` VALUES (299, 44, 42, -20985, 0, -20985, 0, 'EBE3796B939B77E2227C10FA61432338', '闲:[♦6,♠6,] 庄:[♥3,♣5,] 中奖区域:[0 0 1 0 1 0 1 0]', '百佳乐', '2024-08-27 17:18:27', '2024-08-27 17:18:27', NULL, 0, 0);
INSERT INTO `record` VALUES (300, 44, 42, -20985, 163, -20822, 0, '37C017124CE8F6740B48B89FF0A2BAB5', '闲:[♣J,♠3,♦5,] 庄:[♣9,♠7,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 17:18:48', '2024-08-27 17:18:48', NULL, 0, 0);
INSERT INTO `record` VALUES (301, 44, 42, -20822, 151, -20671, 0, '13A44E5ED0B392983E7E9FFDA915F741', '闲:[♥6,♠2,] 庄:[♦9,♠Q,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 17:19:09', '2024-08-27 17:19:09', NULL, 0, 0);
INSERT INTO `record` VALUES (302, 44, 42, -20671, -17, -20688, 0, '6CB58E1194E167B8BD977D494AAD0C8A', '闲:[♦K,♣J,♠5,] 庄:[♦Q,♦Q,♠1,] 中奖区域:[1 0 0 0 0 0 0 1]', '百佳乐', '2024-08-27 17:19:30', '2024-08-27 17:19:30', NULL, 0, 0);
INSERT INTO `record` VALUES (303, 44, 42, -20688, 8, -20680, 0, 'DBBD18DC2B2BD9D1F56F84FB938B59B0', '闲:[♥8,♠5,♠5,] 庄:[♠3,♦3,] 中奖区域:[1 0 0 1 0 0 1 1]', '百佳乐', '2024-08-27 17:19:51', '2024-08-27 17:19:51', NULL, 0, 0);
INSERT INTO `record` VALUES (304, 44, 42, -20680, 819, -19861, 0, 'BB3C18D9A6920AF57FBE0DF98290E64E', '闲:[♦K,♣3,♥Q,] 庄:[♥4,♠9,♠J,] 中奖区域:[0 1 0 0 0 0 0 0]', '百佳乐', '2024-08-27 17:20:12', '2024-08-27 17:20:12', NULL, 0, 0);
INSERT INTO `record` VALUES (305, 44, 42, -19861, 979, -18882, 0, 'C1BF1FA8401F5D255D13CFA7E87500BB', '闲:[♠7,♥Q,] 庄:[♣K,♥5,♦6,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 17:20:33', '2024-08-27 17:20:33', NULL, 0, 0);
INSERT INTO `record` VALUES (306, 44, 42, -18882, 974, -17908, 0, '8A65D59A7369C91282A7E6E654BB3CB1', '闲:[♠K,♥3,♣K,] 庄:[♣8,♠5,♣8,] 中奖区域:[0 1 0 0 0 0 1 1]', '百佳乐', '2024-08-27 17:20:55', '2024-08-27 17:20:55', NULL, 0, 0);
INSERT INTO `record` VALUES (307, 44, 42, -17908, 818, -17090, 0, '964C3BF198C437EA2B59C6B6623AA7C8', '闲:[♣6,♠2,] 庄:[♥1,♥9,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 17:21:16', '2024-08-27 17:21:16', NULL, 0, 0);
INSERT INTO `record` VALUES (308, 44, 42, -17090, 780, -16310, 0, 'E678842290B903267627CA68DCBD7499', '闲:[♣6,♣4,♣Q,] 庄:[♥J,♠2,♥7,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 17:21:37', '2024-08-27 17:21:37', NULL, 0, 0);
INSERT INTO `record` VALUES (309, 44, 42, 200, -91, 109, 0, '5BD2C050B084C7BBA535664226F6DEF9', '闲:[♥6,♥7,♠9,] 庄:[♦5,♥9,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 19:15:20', '2024-08-27 19:15:20', NULL, 0, 0);
INSERT INTO `record` VALUES (310, 44, 42, 109, -149, -40, 0, '397B22349C74D43A7CDD194ED14151FC', '闲:[♥2,♣7,] 庄:[♦2,♠3,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 19:15:48', '2024-08-27 19:15:48', NULL, 0, 0);
INSERT INTO `record` VALUES (311, 44, 42, 200, -17, 183, 0, '2AC21C82686B9939FF1C31922DC3630E', '闲:[♣J,♦7,] 庄:[♣Q,♦10,♣2,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 19:19:43', '2024-08-27 19:19:43', NULL, 0, 0);
INSERT INTO `record` VALUES (312, 44, 42, 183, -173, 10, 0, 'CBADC85AE461E7CC8A7D8A722CB5ED13', '闲:[♦6,♠7,♥2,] 庄:[♥10,♥1,♣10,] 中奖区域:[1 0 0 0 0 0 0 1]', '百佳乐', '2024-08-27 21:11:40', '2024-08-27 21:11:40', NULL, 0, 0);
INSERT INTO `record` VALUES (313, 44, 42, 10, -5, 5, 0, '1F61E379FF9744E09372E57F3454A4D3', '闲:[♦J,♦9,] 庄:[♠4,♦4,] 中奖区域:[1 0 0 1 0 0 0 1]', '百佳乐', '2024-08-27 21:15:45', '2024-08-27 21:15:45', NULL, 0, 0);
INSERT INTO `record` VALUES (314, 44, 42, 5, -5, 0, 0, '3380391D36FC0F612141836E04570301', '闲:[♥Q,♠3,♠J,] 庄:[♥2,♦4,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 21:16:06', '2024-08-27 21:16:06', NULL, 0, 0);
INSERT INTO `record` VALUES (315, 44, 42, 0, -5, -5, 0, '829BB5278C57667D20D38F7E7BA71BE2', '闲:[♥7,♥Q,] 庄:[♠7,♠K,] 中奖区域:[0 1 0 0 0 0 0 0]', '百佳乐', '2024-08-27 21:16:27', '2024-08-27 21:16:27', NULL, 0, 0);
INSERT INTO `record` VALUES (316, 44, 42, -5, -5, -10, 0, '55C44083FD2046D4C70D5C984F116414', '闲:[♣8,♦Q,] 庄:[♣3,♥K,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 21:16:48', '2024-08-27 21:16:48', NULL, 0, 0);
INSERT INTO `record` VALUES (317, 44, 42, -10, -5, -15, 0, '79725214753FCC33B8C2A7DE039B224C', '闲:[♥J,♣2,♠5,] 庄:[♠4,♠2,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 21:17:09', '2024-08-27 21:17:09', NULL, 0, 0);
INSERT INTO `record` VALUES (318, 44, 42, -15, -5, -20, 0, '5B0D63A5621C8F5B4B8E4CFD259FA517', '闲:[♥3,♣J,♦Q,] 庄:[♥8,♣9,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 21:17:30', '2024-08-27 21:17:30', NULL, 0, 0);
INSERT INTO `record` VALUES (319, 44, 42, -20, -5, -25, 0, 'E2676A6ED2F14A3A53FFE21FA2343A29', '闲:[♣3,♦3,] 庄:[♥J,♠3,♦K,] 中奖区域:[1 0 0 0 0 0 1 0]', '百佳乐', '2024-08-27 21:17:51', '2024-08-27 21:17:51', NULL, 0, 0);
INSERT INTO `record` VALUES (320, 44, 42, -25, -5, -30, 0, 'CD012A697F5DA5D7F8B81BF4D1D3FE1F', '闲:[♣4,♣2,] 庄:[♥K,♥Q,♥3,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 21:18:12', '2024-08-27 21:18:12', NULL, 0, 0);
INSERT INTO `record` VALUES (321, 44, 42, -30, -5, -35, 0, 'BF58D8EB28D75B0CE19E629405E93ED1', '闲:[♦10,♣3,♦7,] 庄:[♦6,♥4,♦6,] 中奖区域:[0 1 0 0 0 0 0 1]', '百佳乐', '2024-08-27 21:18:33', '2024-08-27 21:18:33', NULL, 0, 0);
INSERT INTO `record` VALUES (322, 44, 42, -35, -5, -40, 0, '193C365F89840CF1F730A2EC1575E03B', '闲:[♣K,♥10,] 庄:[♣1,♥8,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 21:18:54', '2024-08-27 21:18:54', NULL, 0, 0);
INSERT INTO `record` VALUES (323, 44, 42, -40, -5, -45, 0, 'C73E6AFC2497AB5C25FCF444B075182E', '闲:[♦5,♣8,♣Q,] 庄:[♣6,♠Q,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 21:19:15', '2024-08-27 21:19:15', NULL, 0, 0);
INSERT INTO `record` VALUES (324, 44, 42, -45, -5, -50, 0, '4DA1E99AFF30D3BFF33D5CB96090D5A4', '闲:[♠3,♦6,] 庄:[♥6,♣4,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 21:19:36', '2024-08-27 21:19:36', NULL, 0, 0);
INSERT INTO `record` VALUES (325, 44, 42, -50, -5, -55, 0, '9AFDD6C85E2FA425BACFD312E74DAE01', '闲:[♠8,♣9,] 庄:[♠5,♣2,] 中奖区域:[0 1 0 0 0 0 0 0]', '百佳乐', '2024-08-27 21:19:57', '2024-08-27 21:19:57', NULL, 0, 0);
INSERT INTO `record` VALUES (326, 44, 42, -55, -5, -60, 0, 'E28B03C0C86F25EC0ABFEA7543942365', '闲:[♥1,♠1,♥6,] 庄:[♣9,♥7,] 中奖区域:[1 0 0 1 0 0 1 0]', '百佳乐', '2024-08-27 21:20:18', '2024-08-27 21:20:18', NULL, 0, 0);
INSERT INTO `record` VALUES (327, 44, 42, -60, -5, -65, 0, 'C10BF4774DD72C2FB183C35A96C9A691', '闲:[♥10,♣3,] 庄:[♠9,♥9,] 中奖区域:[0 0 1 0 1 0 0 1]', '百佳乐', '2024-08-27 21:20:39', '2024-08-27 21:20:39', NULL, 0, 0);
INSERT INTO `record` VALUES (328, 44, 42, -65, -5, -70, 0, '2061BFFBC949A90C297D3948B97165F2', '闲:[♣5,♥K,♠1,] 庄:[♥7,♣6,♣8,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 21:21:00', '2024-08-27 21:21:00', NULL, 0, 0);
INSERT INTO `record` VALUES (329, 44, 42, -70, -5, -75, 0, '6F54622EC0C2E75801B2188DAAB5B676', '闲:[♠6,♦8,♦Q,] 庄:[♦8,♦4,♠K,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 21:21:22', '2024-08-27 21:21:22', NULL, 0, 0);
INSERT INTO `record` VALUES (330, 44, 42, -75, -5, -80, 0, 'D46A6957965D29EE014FDAC237AEA0EC', '闲:[♥Q,♥J,♥Q,] 庄:[♥5,♦10,] 中奖区域:[0 0 1 0 0 0 1 0]', '百佳乐', '2024-08-27 21:21:43', '2024-08-27 21:21:43', NULL, 0, 0);
INSERT INTO `record` VALUES (331, 44, 42, -80, -5, -85, 0, 'E5627B7D13F02E099FCB1CC691D9C32A', '闲:[♠10,♦2,♦5,] 庄:[♥4,♦2,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 21:22:04', '2024-08-27 21:22:04', NULL, 0, 0);
INSERT INTO `record` VALUES (332, 44, 42, -85, -5, -90, 0, '23374E1F85B48EDFE504A42D8502A0B6', '闲:[♥10,♦J,♦8,] 庄:[♦3,♣9,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 21:22:25', '2024-08-27 21:22:25', NULL, 0, 0);
INSERT INTO `record` VALUES (333, 44, 42, -90, -5, -95, 0, '131BC81E0D0BA9EC824A8EAA95B4E47F', '闲:[♣2,♣K,] 庄:[♠K,♥9,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 21:22:46', '2024-08-27 21:22:46', NULL, 0, 0);
INSERT INTO `record` VALUES (334, 44, 42, -95, -5, -100, 0, '440175B53780B710D61ECDFBF81A2730', '闲:[♥1,♠10,♠6,] 庄:[♥J,♦1,♦9,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 21:23:07', '2024-08-27 21:23:07', NULL, 0, 0);
INSERT INTO `record` VALUES (335, 44, 42, -100, -5, -105, 0, '9750A8D6A7E8D21F34DD81770F323842', '闲:[♥4,♣7,♥10,] 庄:[♠K,♠2,♥2,] 中奖区域:[0 0 1 0 0 0 0 1]', '百佳乐', '2024-08-27 21:23:28', '2024-08-27 21:23:28', NULL, 0, 0);
INSERT INTO `record` VALUES (336, 44, 42, -105, -5, -110, 0, '498C9142EAC18BE5AC345F94731EDCE7', '闲:[♣3,♦4,] 庄:[♥4,♥1,♥2,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 21:23:49', '2024-08-27 21:23:49', NULL, 0, 0);
INSERT INTO `record` VALUES (337, 44, 42, -110, -5, -115, 0, 'B35AF298050FE2B01FF25E56CECCC19D', '闲:[♥8,♠J,] 庄:[♠7,♦6,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 21:24:10', '2024-08-27 21:24:10', NULL, 0, 0);
INSERT INTO `record` VALUES (338, 44, 42, -115, -5, -120, 0, 'C6FE64C4EC37D0244A95A6B79D6B39C1', '闲:[♥3,♣9,♣9,] 庄:[♦2,♣10,♣1,] 中奖区域:[0 0 1 0 0 0 1 0]', '百佳乐', '2024-08-27 21:24:31', '2024-08-27 21:24:31', NULL, 0, 0);
INSERT INTO `record` VALUES (339, 44, 42, -120, -5, -125, 0, '5CF2AD5809BF8577825375F261199C8C', '闲:[♥8,♥1,] 庄:[♥3,♦2,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 21:24:52', '2024-08-27 21:24:52', NULL, 0, 0);
INSERT INTO `record` VALUES (340, 44, 42, -125, -5, -130, 0, '949B6AD29A3A8EC34BF21CD07950DCB3', '闲:[♣10,♦7,] 庄:[♠8,♠9,] 中奖区域:[0 1 0 0 0 0 0 0]', '百佳乐', '2024-08-27 21:25:13', '2024-08-27 21:25:13', NULL, 0, 0);
INSERT INTO `record` VALUES (341, 44, 42, -130, -5, -135, 0, '56E52548713DB1A0F77AAB58D12C7B0B', '闲:[♥4,♦J,♠6,] 庄:[♦5,♣9,♦K,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 21:25:34', '2024-08-27 21:25:34', NULL, 0, 0);
INSERT INTO `record` VALUES (342, 44, 42, -135, -5, -140, 0, 'E8137CC334D82FD37A0C73CFBDD46588', '闲:[♠Q,♥6,] 庄:[♣3,♦1,♦6,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 21:25:55', '2024-08-27 21:25:55', NULL, 0, 0);
INSERT INTO `record` VALUES (343, 44, 42, -140, -5, -145, 0, 'D9B3CAE9EF482B39CFBDDB5EAEFBA0B6', '闲:[♦Q,♠9,] 庄:[♦Q,♠6,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 21:26:16', '2024-08-27 21:26:16', NULL, 0, 0);
INSERT INTO `record` VALUES (344, 44, 42, -145, -5, -150, 0, '4E71F66881ADC6EA991FE9AABAA5197F', '闲:[♠2,♥5,] 庄:[♥9,♦9,] 中奖区域:[0 0 1 0 1 0 0 1]', '百佳乐', '2024-08-27 21:26:37', '2024-08-27 21:26:37', NULL, 0, 0);
INSERT INTO `record` VALUES (345, 44, 42, -150, -5, -155, 0, '20FEAC36728F564501068DFFC6DBC131', '闲:[♠10,♥5,♣K,] 庄:[♣8,♠5,♠2,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 21:26:58', '2024-08-27 21:26:58', NULL, 0, 0);
INSERT INTO `record` VALUES (346, 44, 42, -155, -5, -160, 0, '0664737B83074157ADDC1594FA538BC5', '闲:[♥2,♠3,♣9,] 庄:[♠7,♥K,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 21:27:19', '2024-08-27 21:27:19', NULL, 0, 0);
INSERT INTO `record` VALUES (347, 44, 42, -160, -5, -165, 0, 'ECCB5C1F314748C2DAA60B31C5FCBA2E', '闲:[♣4,♠1,♣8,] 庄:[♦Q,♦6,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 21:27:40', '2024-08-27 21:27:40', NULL, 0, 0);
INSERT INTO `record` VALUES (348, 44, 42, -165, -5, -170, 0, '82F5CE892C85575F2B1A6C8044C18A31', '闲:[♦6,♥8,] 庄:[♣9,♣K,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 21:28:02', '2024-08-27 21:28:02', NULL, 0, 0);
INSERT INTO `record` VALUES (349, 44, 42, -170, -5, -175, 0, 'CF60BBE4CBFF37B445BDCF10251D6C33', '闲:[♣8,♣9,] 庄:[♦4,♦8,♣J,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 21:28:23', '2024-08-27 21:28:23', NULL, 0, 0);
INSERT INTO `record` VALUES (350, 44, 42, -175, -5, -180, 0, 'C7B53641709ABDEAA082241723205842', '闲:[♠5,♥3,] 庄:[♠4,♣7,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 21:28:44', '2024-08-27 21:28:44', NULL, 0, 0);
INSERT INTO `record` VALUES (351, 44, 42, -180, -5, -185, 0, '0877AC7951E923F5481503BBF21B751A', '闲:[♥8,♦K,] 庄:[♣10,♣4,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 21:29:05', '2024-08-27 21:29:05', NULL, 0, 0);
INSERT INTO `record` VALUES (352, 44, 42, -185, -5, -190, 0, 'C854457550C779C9D23304547BDFCD02', '闲:[♦2,♠7,] 庄:[♦10,♠3,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 21:29:26', '2024-08-27 21:29:26', NULL, 0, 0);
INSERT INTO `record` VALUES (353, 44, 42, -190, -5, -195, 0, '9F1B0EE4B48365013D197BD51ECC150D', '闲:[♥4,♣9,♦5,] 庄:[♠7,♣K,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 21:29:47', '2024-08-27 21:29:47', NULL, 0, 0);
INSERT INTO `record` VALUES (354, 44, 42, -195, -5, -200, 0, 'A7A99C6F4E2F55C43B0608A445C312DC', '闲:[♣10,♠7,] 庄:[♦2,♥3,♥10,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 21:30:08', '2024-08-27 21:30:08', NULL, 0, 0);
INSERT INTO `record` VALUES (355, 44, 42, -200, -5, -205, 0, '34BB4D2EE8B7B9FC99E3F1638FC02B94', '闲:[♣4,♣8,♣2,] 庄:[♣6,♠7,♠10,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 21:30:29', '2024-08-27 21:30:29', NULL, 0, 0);
INSERT INTO `record` VALUES (356, 44, 42, -205, -5, -210, 0, 'B649801C066AB3118A957262C8680F22', '闲:[♦K,♣9,] 庄:[♦2,♦9,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 21:30:50', '2024-08-27 21:30:50', NULL, 0, 0);
INSERT INTO `record` VALUES (357, 44, 42, -210, -5, -215, 0, 'DAF0A888810E010D053A96A1D08A766C', '闲:[♣9,♠K,] 庄:[♥8,♥K,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 21:31:11', '2024-08-27 21:31:11', NULL, 0, 0);
INSERT INTO `record` VALUES (358, 44, 42, -215, -5, -220, 0, '4A97EC1ED23251A502A97470C5F88C58', '闲:[♦6,♠J,] 庄:[♣J,♦6,] 中奖区域:[0 1 0 0 0 0 0 0]', '百佳乐', '2024-08-27 21:31:32', '2024-08-27 21:31:32', NULL, 0, 0);
INSERT INTO `record` VALUES (359, 44, 42, -220, -5, -225, 0, 'FD2ED9D348221CF062B205262F54F3C6', '闲:[♣6,♠J,] 庄:[♠7,♣5,♠K,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 21:31:53', '2024-08-27 21:31:53', NULL, 0, 0);
INSERT INTO `record` VALUES (360, 44, 42, -225, -5, -230, 0, '72E04794E9F6C04D6A85F499AB2BEF01', '闲:[♥10,♣6,] 庄:[♦3,♣1,♥7,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 21:32:14', '2024-08-27 21:32:14', NULL, 0, 0);
INSERT INTO `record` VALUES (361, 44, 42, -230, -5, -235, 0, '2D368D18162171DDD3A660249E540BC4', '闲:[♦10,♠2,♥10,] 庄:[♥3,♠Q,♠Q,] 中奖区域:[0 0 1 0 0 0 1 1]', '百佳乐', '2024-08-27 21:32:35', '2024-08-27 21:32:35', NULL, 0, 0);
INSERT INTO `record` VALUES (362, 44, 42, -235, -5, -240, 0, '840BBEC86B4FBDD43D2E4BC7D516DE18', '闲:[♣8,♠6,] 庄:[♥7,♦2,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 21:32:56', '2024-08-27 21:32:56', NULL, 0, 0);
INSERT INTO `record` VALUES (363, 44, 42, -240, -5, -245, 0, '14400338F44A37FAB27EE19188C62A93', '闲:[♥9,♦9,] 庄:[♣6,♦1,] 中奖区域:[1 0 0 1 0 0 1 0]', '百佳乐', '2024-08-27 21:33:17', '2024-08-27 21:33:17', NULL, 0, 0);
INSERT INTO `record` VALUES (364, 44, 42, -245, -5, -250, 0, '9EA92BCD9BA941AC2785347701907A9D', '闲:[♥2,♦6,] 庄:[♦Q,♦9,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 21:33:38', '2024-08-27 21:33:38', NULL, 0, 0);
INSERT INTO `record` VALUES (365, 44, 42, -250, -5, -255, 0, 'E035AA0E8EB0C35CE8E66B79E163050B', '闲:[♣3,♣7,♥4,] 庄:[♦4,♦J,♣1,] 中奖区域:[0 1 0 0 0 0 0 0]', '百佳乐', '2024-08-27 21:33:59', '2024-08-27 21:33:59', NULL, 0, 0);
INSERT INTO `record` VALUES (366, 44, 42, -255, -5, -260, 0, 'C2B86836F8C444AA3378874D90C4F12A', '闲:[♦J,♣Q,] 庄:[♥9,♠Q,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 21:34:20', '2024-08-27 21:34:20', NULL, 0, 0);
INSERT INTO `record` VALUES (367, 44, 42, -260, -5, -265, 0, '8D762528F48CBC0AC230E7153108D5B6', '闲:[♠8,♦J,] 庄:[♥1,♦J,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 21:34:42', '2024-08-27 21:34:42', NULL, 0, 0);
INSERT INTO `record` VALUES (368, 44, 42, -265, -5, -270, 0, '59A03A7F98E5907B3582807EE4596C07', '闲:[♦6,♦5,] 庄:[♣6,♦3,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 21:35:03', '2024-08-27 21:35:03', NULL, 0, 0);
INSERT INTO `record` VALUES (369, 44, 42, -270, -5, -275, 0, '8F073EED4CC9E8A548A143FCFD954971', '闲:[♥6,♦2,] 庄:[♣K,♦7,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 21:35:24', '2024-08-27 21:35:24', NULL, 0, 0);
INSERT INTO `record` VALUES (370, 44, 42, -275, -5, -280, 0, 'ECBAC6D3D6BDC9DEACF31620ADF955A3', '闲:[♦Q,♦J,♣2,] 庄:[♠Q,♥J,♥1,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 21:35:45', '2024-08-27 21:35:45', NULL, 0, 0);
INSERT INTO `record` VALUES (371, 44, 42, -280, -5, -285, 0, 'B6F09EC890C15BC19AC3FF596C0FC3CC', '闲:[♦5,♠1,] 庄:[♥10,♦Q,♥8,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 21:36:06', '2024-08-27 21:36:06', NULL, 0, 0);
INSERT INTO `record` VALUES (372, 44, 42, -285, -5, -290, 0, 'CF3628EAB51AA2FB8A258FD2484DC29C', '闲:[♦1,♦9,♥4,] 庄:[♥K,♦5,♣2,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 21:36:27', '2024-08-27 21:36:27', NULL, 0, 0);
INSERT INTO `record` VALUES (373, 44, 42, -290, -5, -295, 0, 'F9B9435003AECADA132F64EC33B979C0', '闲:[♠8,♣2,] 庄:[♦9,♠K,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 21:36:48', '2024-08-27 21:36:48', NULL, 0, 0);
INSERT INTO `record` VALUES (374, 44, 42, -295, -5, -300, 0, 'E3E5C98CF3BA45A719743E8F07C58DF5', '闲:[♦6,♦10,] 庄:[♠K,♥3,♣7,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 21:37:09', '2024-08-27 21:37:09', NULL, 0, 0);
INSERT INTO `record` VALUES (375, 44, 42, -300, -5, -305, 0, 'A026F8852AB19A7997748C4EB3199F09', '闲:[♥Q,♥10,♠Q,] 庄:[♥Q,♦3,♦4,] 中奖区域:[0 0 1 0 0 0 1 0]', '百佳乐', '2024-08-27 21:37:30', '2024-08-27 21:37:30', NULL, 0, 0);
INSERT INTO `record` VALUES (376, 44, 42, -305, -5, -310, 0, 'CAF9694D072F636D0F56BE607BC15898', '闲:[♣9,♣9,] 庄:[♣3,♣1,] 中奖区域:[1 0 0 1 0 0 1 0]', '百佳乐', '2024-08-27 21:37:51', '2024-08-27 21:37:51', NULL, 0, 0);
INSERT INTO `record` VALUES (377, 44, 42, -310, -5, -315, 0, '14BD63B303FA978281403009B9D5410A', '闲:[♥1,♥4,♣5,] 庄:[♣9,♥6,♥2,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 21:38:12', '2024-08-27 21:38:12', NULL, 0, 0);
INSERT INTO `record` VALUES (378, 44, 42, -315, -5, -320, 0, '0474F4CE50886E89A08978663C222376', '闲:[♠8,♠Q,] 庄:[♦9,♥J,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 21:38:33', '2024-08-27 21:38:33', NULL, 0, 0);
INSERT INTO `record` VALUES (379, 44, 42, -320, -5, -325, 0, 'D9ABB9AE1C8A98A57FB415B2948E4140', '闲:[♥8,♦10,] 庄:[♦1,♦2,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 21:38:54', '2024-08-27 21:38:54', NULL, 0, 0);
INSERT INTO `record` VALUES (380, 44, 42, -325, -5, -330, 0, 'D3C54D4384B0E339F0CD2BE875E8391C', '闲:[♣5,♥6,♥4,] 庄:[♥4,♣2,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 21:39:15', '2024-08-27 21:39:15', NULL, 0, 0);
INSERT INTO `record` VALUES (381, 44, 42, -330, -5, -335, 0, '22450A806F0AFA032A8313F574D19BD0', '闲:[♦3,♥9,♣10,] 庄:[♦K,♦6,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 21:39:36', '2024-08-27 21:39:36', NULL, 0, 0);
INSERT INTO `record` VALUES (382, 44, 42, -335, -5, -340, 0, '5F990C20315574BDF5D878715AEA3F1C', '闲:[♠9,♣10,] 庄:[♥2,♦J,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 21:39:57', '2024-08-27 21:39:57', NULL, 0, 0);
INSERT INTO `record` VALUES (383, 44, 42, -340, -5, -345, 0, 'BCAA97136759CEDA4B5D18376CF155C6', '闲:[♣6,♦9,♠J,] 庄:[♣6,♣6,♥5,] 中奖区域:[1 0 0 0 0 0 0 1]', '百佳乐', '2024-08-27 21:40:18', '2024-08-27 21:40:18', NULL, 0, 0);
INSERT INTO `record` VALUES (384, 44, 42, -345, -5, -350, 0, '3447FDDD6B528952BCCC80B5BDA0F7EA', '闲:[♣5,♦4,] 庄:[♠J,♠10,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 21:40:39', '2024-08-27 21:40:39', NULL, 0, 0);
INSERT INTO `record` VALUES (385, 44, 42, -350, -5, -355, 0, 'CE0B0BB76B020F97FC181BEF75E0070C', '闲:[♦Q,♥6,] 庄:[♦Q,♠9,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 21:41:01', '2024-08-27 21:41:01', NULL, 0, 0);
INSERT INTO `record` VALUES (386, 44, 42, -355, -5, -360, 0, 'AE4E892F6C08CB2B576E2EBE03B99419', '闲:[♣6,♦7,♠9,] 庄:[♦3,♥Q,♦9,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 21:41:22', '2024-08-27 21:41:22', NULL, 0, 0);
INSERT INTO `record` VALUES (387, 44, 42, -360, -5, -365, 0, '054BAE156DBBEB5B2C91180E460237B7', '闲:[♣K,♦3,♣1,] 庄:[♣10,♦J,♠Q,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 21:41:43', '2024-08-27 21:41:43', NULL, 0, 0);
INSERT INTO `record` VALUES (388, 44, 42, -365, -5, -370, 0, 'FCBDB3AAC8CDA8EDB1EEBD8686A08C87', '闲:[♠4,♦4,] 庄:[♣7,♣9,] 中奖区域:[1 0 0 1 0 0 1 0]', '百佳乐', '2024-08-27 21:42:04', '2024-08-27 21:42:04', NULL, 0, 0);
INSERT INTO `record` VALUES (389, 44, 42, -370, -5, -375, 0, '20A5E88D5029EDC6FBE4EBE430C43A25', '闲:[♠7,♦2,] 庄:[♥4,♦7,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 21:42:25', '2024-08-27 21:42:25', NULL, 0, 0);
INSERT INTO `record` VALUES (390, 44, 42, -375, -5, -380, 0, '501AD1A9A2124045D689C5ABCBDEFA43', '闲:[♠8,♣2,♥1,] 庄:[♦5,♥K,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 21:42:46', '2024-08-27 21:42:46', NULL, 0, 0);
INSERT INTO `record` VALUES (391, 44, 42, -380, -5, -385, 0, '5E7AA6B2C55180430A42B028BC37A8F0', '闲:[♣8,♣Q,] 庄:[♠10,♠7,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 21:43:07', '2024-08-27 21:43:07', NULL, 0, 0);
INSERT INTO `record` VALUES (392, 44, 42, -385, -5, -390, 0, 'F66E8EDDD2E9B07D0D6CECA22062CAD4', '闲:[♥Q,♠Q,♣2,] 庄:[♥7,♦4,♣8,] 中奖区域:[1 0 0 0 0 0 1 0]', '百佳乐', '2024-08-27 21:43:28', '2024-08-27 21:43:28', NULL, 0, 0);
INSERT INTO `record` VALUES (393, 44, 42, -390, -5, -395, 0, '3C14DBEC43EA96CAA5A305FEC998DCFB', '闲:[♥K,♥9,] 庄:[♠K,♥9,] 中奖区域:[0 1 0 0 0 1 0 0]', '百佳乐', '2024-08-27 21:43:49', '2024-08-27 21:43:49', NULL, 0, 0);
INSERT INTO `record` VALUES (394, 44, 42, -395, -5, -400, 0, '05C411B46F04754A53957AC28807D362', '闲:[♥K,♠2,♠3,] 庄:[♦1,♠6,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 21:44:10', '2024-08-27 21:44:10', NULL, 0, 0);
INSERT INTO `record` VALUES (395, 44, 42, -400, -5, -405, 0, 'E8DBF98C39C504DE17841E94FC55CFA9', '闲:[♦10,♦K,♣1,] 庄:[♦3,♥1,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 21:44:31', '2024-08-27 21:44:31', NULL, 0, 0);
INSERT INTO `record` VALUES (396, 44, 42, -405, -5, -410, 0, '760C3C9A96C7C697710DF735490FB9AE', '闲:[♣8,♠Q,] 庄:[♦3,♥10,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 21:44:52', '2024-08-27 21:44:52', NULL, 0, 0);
INSERT INTO `record` VALUES (397, 44, 42, -410, -5, -415, 0, '229D74908D0F1967E5A54026B8FBB7FF', '闲:[♦J,♥Q,♥7,] 庄:[♥Q,♥J,♣3,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 21:45:13', '2024-08-27 21:45:13', NULL, 0, 0);
INSERT INTO `record` VALUES (398, 44, 42, -415, -5, -420, 0, '8E21FC9C2EDC91A8C8A10BEC3DEE1D86', '闲:[♥K,♥9,] 庄:[♦3,♣5,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 21:45:34', '2024-08-27 21:45:34', NULL, 0, 0);
INSERT INTO `record` VALUES (399, 44, 42, -420, -5, -425, 0, '1EDC195F561C446ECB0C70BAA2567FF7', '闲:[♣9,♣K,] 庄:[♥10,♠Q,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 21:45:55', '2024-08-27 21:45:55', NULL, 0, 0);
INSERT INTO `record` VALUES (400, 44, 42, -425, -5, -430, 0, '56CDF49391F69A68DF604020BE97A2B0', '闲:[♦10,♠4,♦5,] 庄:[♠6,♦6,] 中奖区域:[1 0 0 1 0 0 0 1]', '百佳乐', '2024-08-27 21:46:16', '2024-08-27 21:46:16', NULL, 0, 0);
INSERT INTO `record` VALUES (401, 44, 42, -430, -5, -435, 0, '3EBF9D63199EFA174ADDD4B86A7C1285', '闲:[♥3,♣5,] 庄:[♠J,♣8,] 中奖区域:[0 1 0 0 0 0 0 0]', '百佳乐', '2024-08-27 21:46:37', '2024-08-27 21:46:37', NULL, 0, 0);
INSERT INTO `record` VALUES (402, 44, 42, -435, -5, -440, 0, '3B1C4DB735744FE0313E074DC87374BC', '闲:[♠Q,♥K,♠9,] 庄:[♦K,♣6,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 21:46:58', '2024-08-27 21:46:58', NULL, 0, 0);
INSERT INTO `record` VALUES (403, 44, 42, -440, -5, -445, 0, '36DA51266DD325B4F183E4240D2CD2A6', '闲:[♣3,♣8,♣5,] 庄:[♣8,♠5,♦3,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 21:47:20', '2024-08-27 21:47:20', NULL, 0, 0);
INSERT INTO `record` VALUES (404, 44, 42, -445, -5, -450, 0, '15935E9240ED29F6793579C4A33080CB', '闲:[♥K,♠10,] 庄:[♦7,♥1,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 21:47:41', '2024-08-27 21:47:41', NULL, 0, 0);
INSERT INTO `record` VALUES (405, 44, 42, -450, -5, -455, 0, '6C093FEDE3A4C0CC1CB0D0A53755F58B', '闲:[♦K,♣5,♥9,] 庄:[♣5,♥5,♥Q,] 中奖区域:[1 0 0 0 0 0 0 1]', '百佳乐', '2024-08-27 21:48:02', '2024-08-27 21:48:02', NULL, 0, 0);
INSERT INTO `record` VALUES (406, 44, 42, -455, -5, -460, 0, '9FFBCB174C85D0B8EC34E077BD9E51B0', '闲:[♦7,♦4,♥J,] 庄:[♠9,♥5,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 21:48:23', '2024-08-27 21:48:23', NULL, 0, 0);
INSERT INTO `record` VALUES (407, 44, 42, -460, -5, -465, 0, '2954B7746E29176AA83AE05642A8227B', '闲:[♣J,♠3,] 庄:[♣J,♣8,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 21:48:44', '2024-08-27 21:48:44', NULL, 0, 0);
INSERT INTO `record` VALUES (408, 44, 42, -465, -5, -470, 0, 'F3D8B109212D4FCF09250DCD21F48C44', '闲:[♥9,♥4,♦J,] 庄:[♥2,♣8,♠3,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 21:49:05', '2024-08-27 21:49:05', NULL, 0, 0);
INSERT INTO `record` VALUES (409, 44, 42, -470, -5, -475, 0, '6E732DEDEA0BB3F71A629C856DA39533', '闲:[♠10,♠4,] 庄:[♥3,♣5,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 21:49:26', '2024-08-27 21:49:26', NULL, 0, 0);
INSERT INTO `record` VALUES (410, 44, 42, -475, -5, -480, 0, '16796790F63832B8101D580E423807E6', '闲:[♣8,♠J,] 庄:[♣J,♥9,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 21:49:47', '2024-08-27 21:49:47', NULL, 0, 0);
INSERT INTO `record` VALUES (411, 44, 42, -480, -5, -485, 0, 'A028FCE9C2068F50B0A1C3907AEBB676', '闲:[♦4,♦9,♦9,] 庄:[♠2,♥K,♠8,] 中奖区域:[0 1 0 0 0 0 1 0]', '百佳乐', '2024-08-27 21:50:08', '2024-08-27 21:50:08', NULL, 0, 0);
INSERT INTO `record` VALUES (412, 44, 42, -485, -5, -490, 0, '791B5C679CED78DA17401D33526CCE6B', '闲:[♠3,♥5,] 庄:[♦5,♥3,] 中奖区域:[0 1 0 0 0 0 0 0]', '百佳乐', '2024-08-27 21:50:29', '2024-08-27 21:50:29', NULL, 0, 0);
INSERT INTO `record` VALUES (413, 44, 42, -490, -5, -495, 0, '12BCBA87E2826831649FFE7CDAF53501', '闲:[♠9,♣8,] 庄:[♦9,♥4,♣8,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 21:50:50', '2024-08-27 21:50:50', NULL, 0, 0);
INSERT INTO `record` VALUES (414, 44, 42, -495, -5, -500, 0, '3B190D0575924C38ADD3166116D7446B', '闲:[♥3,♣9,♥7,] 庄:[♣3,♠K,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 21:51:11', '2024-08-27 21:51:11', NULL, 0, 0);
INSERT INTO `record` VALUES (415, 44, 42, -500, -5, -505, 0, '463A24DFCCBE10B1B512B2B605583B68', '闲:[♣9,♥K,] 庄:[♣8,♣Q,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 21:51:32', '2024-08-27 21:51:32', NULL, 0, 0);
INSERT INTO `record` VALUES (416, 44, 42, -505, -5, -510, 0, '3377CA9350E3E5DF341E2F94096E3A37', '闲:[♥6,♠4,] 庄:[♥9,♠K,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 21:51:53', '2024-08-27 21:51:53', NULL, 0, 0);
INSERT INTO `record` VALUES (417, 44, 42, -510, -5, -515, 0, '484473FDC37624F91E1E6C976BE68211', '闲:[♥4,♠Q,♠5,] 庄:[♠Q,♣1,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 21:52:14', '2024-08-27 21:52:14', NULL, 0, 0);
INSERT INTO `record` VALUES (418, 44, 42, -515, -5, -520, 0, 'F1C697160BD2A5508ABCCFD772A48F5E', '闲:[♦7,♠5,♣K,] 庄:[♠J,♥K,♥Q,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 21:52:35', '2024-08-27 21:52:35', NULL, 0, 0);
INSERT INTO `record` VALUES (419, 44, 42, -520, -5, -525, 0, 'A57920BBB5793746A9D1ACDA08631F4C', '闲:[♦6,♦3,] 庄:[♥1,♦8,] 中奖区域:[0 1 0 0 0 0 0 0]', '百佳乐', '2024-08-27 21:52:56', '2024-08-27 21:52:56', NULL, 0, 0);
INSERT INTO `record` VALUES (420, 44, 42, -525, -5, -530, 0, 'C3AFA8B853E16D0B80F506BCB206C228', '闲:[♣6,♠3,] 庄:[♣3,♦4,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 21:53:17', '2024-08-27 21:53:17', NULL, 0, 0);
INSERT INTO `record` VALUES (421, 44, 42, -530, -5, -535, 0, '5EC391B2746E67026AA663E8DDD42466', '闲:[♣4,♠4,] 庄:[♠2,♣7,] 中奖区域:[0 0 1 0 1 0 1 0]', '百佳乐', '2024-08-27 21:53:38', '2024-08-27 21:53:38', NULL, 0, 0);
INSERT INTO `record` VALUES (422, 44, 42, -535, -5, -540, 0, 'CF2C550476240172387A59FDCE106F06', '闲:[♦4,♣K,♦8,] 庄:[♦10,♠1,♦2,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 21:54:00', '2024-08-27 21:54:00', NULL, 0, 0);
INSERT INTO `record` VALUES (423, 44, 42, -540, -5, -545, 0, 'C71D13997E84300EE6F09879BA909E20', '闲:[♥10,♠5,♥8,] 庄:[♠5,♦8,] 中奖区域:[0 1 0 0 0 0 0 0]', '百佳乐', '2024-08-27 21:54:21', '2024-08-27 21:54:21', NULL, 0, 0);
INSERT INTO `record` VALUES (424, 44, 42, -545, -5, -550, 0, '72A7027D7470CAB385A88B5F6906FE20', '闲:[♥9,♦J,] 庄:[♣8,♦7,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 21:54:42', '2024-08-27 21:54:42', NULL, 0, 0);
INSERT INTO `record` VALUES (425, 44, 42, -550, -5, -555, 0, '02C5859EB11F840C8E1920F82EA1AB2F', '闲:[♣8,♦10,] 庄:[♥Q,♦1,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 21:55:03', '2024-08-27 21:55:03', NULL, 0, 0);
INSERT INTO `record` VALUES (426, 44, 42, -555, -5, -560, 0, '1304D7F2F453FD40415C16583B10DA8D', '闲:[♥5,♣9,♠3,] 庄:[♠K,♦Q,♣K,] 中奖区域:[1 0 0 0 0 0 0 1]', '百佳乐', '2024-08-27 21:55:24', '2024-08-27 21:55:24', NULL, 0, 0);
INSERT INTO `record` VALUES (427, 44, 42, -560, -5, -565, 0, '7C2B6105BBA51DBBE975879CCFB03956', '闲:[♦Q,♠3,♣Q,] 庄:[♦4,♣9,♣5,] 中奖区域:[0 1 0 0 0 0 1 0]', '百佳乐', '2024-08-27 21:55:45', '2024-08-27 21:55:45', NULL, 0, 0);
INSERT INTO `record` VALUES (428, 44, 42, -565, -5, -570, 0, '35778F0D8C33263F0E0F26EF4C6F0D4F', '闲:[♥Q,♥6,] 庄:[♠1,♥K,♣K,] 中奖区域:[1 0 0 0 0 0 0 1]', '百佳乐', '2024-08-27 21:56:06', '2024-08-27 21:56:06', NULL, 0, 0);
INSERT INTO `record` VALUES (429, 44, 42, -570, -5, -575, 0, '7537BC1B4EA6059906141F4475AF6F5C', '闲:[♦8,♣1,] 庄:[♣K,♦K,] 中奖区域:[1 0 0 1 0 0 0 1]', '百佳乐', '2024-08-27 21:56:27', '2024-08-27 21:56:27', NULL, 0, 0);
INSERT INTO `record` VALUES (430, 44, 42, -575, -5, -580, 0, '9FA32C20E1ACB57240F69897DBB8E19B', '闲:[♠3,♥K,♦10,] 庄:[♥J,♠1,♦5,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 21:56:48', '2024-08-27 21:56:48', NULL, 0, 0);
INSERT INTO `record` VALUES (431, 44, 42, -580, -5, -585, 0, 'AEBA982BCDC69B31266F103C01E3C719', '闲:[♣Q,♥2,♣7,] 庄:[♣3,♥Q,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 21:57:09', '2024-08-27 21:57:09', NULL, 0, 0);
INSERT INTO `record` VALUES (432, 44, 42, -585, -5, -590, 0, 'B968BE51B506F45C3005EBED1CC8FB55', '闲:[♠7,♣9,] 庄:[♣3,♥J,♠3,] 中奖区域:[1 0 0 0 0 0 0 1]', '百佳乐', '2024-08-27 21:57:30', '2024-08-27 21:57:30', NULL, 0, 0);
INSERT INTO `record` VALUES (433, 44, 42, -590, -5, -595, 0, '477FD092EA066B1126B7A1DE04E0C574', '闲:[♥7,♦Q,] 庄:[♥Q,♦9,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 21:57:51', '2024-08-27 21:57:51', NULL, 0, 0);
INSERT INTO `record` VALUES (434, 44, 42, -595, -5, -600, 0, '9DC6846A5048CA0630044DCABDAB3006', '闲:[♦4,♦7,♥9,] 庄:[♠2,♣J,♠10,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 21:58:12', '2024-08-27 21:58:12', NULL, 0, 0);
INSERT INTO `record` VALUES (435, 44, 42, -600, -5, -605, 0, 'B0A7A50102C9C028639A6091E139312E', '闲:[♣7,♦10,] 庄:[♥Q,♠4,♠J,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 21:58:33', '2024-08-27 21:58:33', NULL, 0, 0);
INSERT INTO `record` VALUES (436, 44, 42, -605, -5, -610, 0, 'BABC5479E6B0813C51E15BDECE318B8B', '闲:[♥2,♥2,♣8,] 庄:[♠J,♦3,] 中奖区域:[0 0 1 0 0 0 1 0]', '百佳乐', '2024-08-27 21:58:54', '2024-08-27 21:58:54', NULL, 0, 0);
INSERT INTO `record` VALUES (437, 44, 42, -610, -5, -615, 0, 'A013142B06BA4620E130CE3810E07F27', '闲:[♦Q,♥7,] 庄:[♠J,♦9,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 21:59:15', '2024-08-27 21:59:15', NULL, 0, 0);
INSERT INTO `record` VALUES (438, 44, 42, -615, -5, -620, 0, 'E8F3A73565D58E3380F48CD90F7D76AB', '闲:[♦5,♦6,♣2,] 庄:[♠9,♥8,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 21:59:36', '2024-08-27 21:59:36', NULL, 0, 0);
INSERT INTO `record` VALUES (439, 44, 42, -620, -5, -625, 0, '5D0CA47866A73CFFEDFF0D2976B82F93', '闲:[♠8,♥10,] 庄:[♠3,♣J,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 21:59:57', '2024-08-27 21:59:57', NULL, 0, 0);
INSERT INTO `record` VALUES (440, 44, 42, -625, -5, -630, 0, 'FFF798DDAC11C7FEA3A1C5B1BB54158C', '闲:[♥5,♦4,] 庄:[♥4,♠K,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:00:18', '2024-08-27 22:00:18', NULL, 0, 0);
INSERT INTO `record` VALUES (441, 44, 42, -630, -5, -635, 0, '1B2E76FB47B395DCD38FC1683277C0AE', '闲:[♠6,♣8,♥4,] 庄:[♥1,♣6,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:00:39', '2024-08-27 22:00:39', NULL, 0, 0);
INSERT INTO `record` VALUES (442, 44, 42, -635, -5, -640, 0, 'C340AB5103F37F32FB2D0C5406F376D8', '闲:[♠1,♣Q,] 庄:[♥5,♣4,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 22:01:01', '2024-08-27 22:01:01', NULL, 0, 0);
INSERT INTO `record` VALUES (443, 44, 42, -640, -5, -645, 0, 'B204FEED44ABCA538EE89A1B0E4E79CE', '闲:[♦10,♦6,] 庄:[♣J,♠7,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 22:01:22', '2024-08-27 22:01:22', NULL, 0, 0);
INSERT INTO `record` VALUES (444, 44, 42, -645, -5, -650, 0, '8348E2EF6C1DBBB18DB4B8B88EE3AD4B', '闲:[♠9,♦1,♥Q,] 庄:[♠2,♦2,] 中奖区域:[0 0 1 0 0 0 0 1]', '百佳乐', '2024-08-27 22:01:43', '2024-08-27 22:01:43', NULL, 0, 0);
INSERT INTO `record` VALUES (445, 44, 42, -650, -5, -655, 0, '52BAAF88A473EFD2BFD6105DD88BC124', '闲:[♦9,♥K,] 庄:[♠2,♥4,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:02:04', '2024-08-27 22:02:04', NULL, 0, 0);
INSERT INTO `record` VALUES (446, 44, 42, -655, -5, -660, 0, '4AA34D3C2D46FD9B85F6C3B4E18BF51A', '闲:[♣3,♣1,♦J,] 庄:[♦Q,♣3,♥10,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:02:25', '2024-08-27 22:02:25', NULL, 0, 0);
INSERT INTO `record` VALUES (447, 44, 42, -660, -5, -665, 0, 'BEF019ACCC4D7362790C73EC011FCADF', '闲:[♠6,♦6,♥2,] 庄:[♥8,♠6,♠J,] 中奖区域:[0 1 0 0 0 0 1 0]', '百佳乐', '2024-08-27 22:02:46', '2024-08-27 22:02:46', NULL, 0, 0);
INSERT INTO `record` VALUES (448, 44, 42, -665, -5, -670, 0, '2CCAD303C149613DF6676DC27FC64E5D', '闲:[♥9,♥4,♥10,] 庄:[♠9,♣1,♥4,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:03:07', '2024-08-27 22:03:07', NULL, 0, 0);
INSERT INTO `record` VALUES (449, 44, 42, -670, -5, -675, 0, 'D928DB4F61E05FFB513497DDEF730B5E', '闲:[♦6,♣10,] 庄:[♠8,♦4,♦1,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:03:28', '2024-08-27 22:03:28', NULL, 0, 0);
INSERT INTO `record` VALUES (450, 44, 42, -675, -5, -680, 0, '364A5344484416D491C8BEB480DEC712', '闲:[♠8,♠5,] 庄:[♦8,♥10,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 22:03:49', '2024-08-27 22:03:49', NULL, 0, 0);
INSERT INTO `record` VALUES (451, 44, 42, -680, -5, -685, 0, '860E0E684EF084103B0D6AB1EDB6CE15', '闲:[♦4,♣Q,♣1,] 庄:[♥J,♠10,♦K,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:04:10', '2024-08-27 22:04:10', NULL, 0, 0);
INSERT INTO `record` VALUES (452, 44, 42, -685, -5, -690, 0, '33B073AF4C33D54E954F7E6183B222BE', '闲:[♦K,♣4,♣4,] 庄:[♠6,♥9,] 中奖区域:[1 0 0 1 0 0 1 0]', '百佳乐', '2024-08-27 22:04:31', '2024-08-27 22:04:31', NULL, 0, 0);
INSERT INTO `record` VALUES (453, 44, 42, -690, -5, -695, 0, '49E829BF841CDF6965647CB210FE7315', '闲:[♣3,♣1,♥9,] 庄:[♥8,♣4,♠3,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:04:52', '2024-08-27 22:04:52', NULL, 0, 0);
INSERT INTO `record` VALUES (454, 44, 42, -695, -5, -700, 0, 'BF53F70C89F6C97DABF29E73D3850BFF', '闲:[♦7,♦J,] 庄:[♠2,♦4,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:05:13', '2024-08-27 22:05:13', NULL, 0, 0);
INSERT INTO `record` VALUES (455, 44, 42, -700, -5, -705, 0, 'C437729F9D6485ED90802AF0C539CDF9', '闲:[♥9,♠1,♥10,] 庄:[♣1,♣4,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 22:05:34', '2024-08-27 22:05:34', NULL, 0, 0);
INSERT INTO `record` VALUES (456, 44, 42, -705, -5, -710, 0, '3C9FF5EB3219D192D650CBFB331137B3', '闲:[♣5,♥7,♦6,] 庄:[♦5,♠1,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:05:55', '2024-08-27 22:05:55', NULL, 0, 0);
INSERT INTO `record` VALUES (457, 44, 42, -710, -5, -715, 0, 'E187DD82289B6DE1C4EA552491B1158D', '闲:[♥Q,♣K,♦8,] 庄:[♠1,♣K,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:06:16', '2024-08-27 22:06:16', NULL, 0, 0);
INSERT INTO `record` VALUES (458, 44, 42, -715, -5, -720, 0, 'A3B3211E6409F0293D773A35C0CC84F9', '闲:[♦Q,♥3,♣K,] 庄:[♥7,♣9,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 22:06:37', '2024-08-27 22:06:37', NULL, 0, 0);
INSERT INTO `record` VALUES (459, 44, 42, -720, -5, -725, 0, 'BF55494B298D8982E2FF5E32EA07C057', '闲:[♣5,♦1,] 庄:[♠4,♣8,♥7,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:06:58', '2024-08-27 22:06:58', NULL, 0, 0);
INSERT INTO `record` VALUES (460, 44, 42, -725, -5, -730, 0, 'C05EC94059ABE977F10F7C850FA3ABA0', '闲:[♦Q,♥1,♣7,] 庄:[♠10,♦5,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:07:19', '2024-08-27 22:07:19', NULL, 0, 0);
INSERT INTO `record` VALUES (461, 44, 42, -730, -5, -735, 0, '9780D616E51E09EC94131E160D825898', '闲:[♣10,♠6,] 庄:[♣2,♦10,♣Q,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:07:40', '2024-08-27 22:07:40', NULL, 0, 0);
INSERT INTO `record` VALUES (462, 44, 42, -735, -5, -740, 0, '3581283402007DCFA051FFD9F177936E', '闲:[♣2,♥9,♠8,] 庄:[♥7,♣K,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:08:02', '2024-08-27 22:08:02', NULL, 0, 0);
INSERT INTO `record` VALUES (463, 44, 42, -740, -5, -745, 0, '99ED1D28BB3A942E70C2CF87B72EAF41', '闲:[♠8,♠10,] 庄:[♦3,♣Q,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:08:23', '2024-08-27 22:08:23', NULL, 0, 0);
INSERT INTO `record` VALUES (464, 44, 42, -745, -5, -750, 0, 'DCC9BB3F1418B531A510977387A653CC', '闲:[♠9,♣3,♣8,] 庄:[♣1,♦5,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 22:08:44', '2024-08-27 22:08:44', NULL, 0, 0);
INSERT INTO `record` VALUES (465, 44, 42, -750, -5, -755, 0, '51DE46DD630BDCB7BC41E0D8E3B8FA8F', '闲:[♠4,♦7,♣7,] 庄:[♦K,♣K,] 中奖区域:[1 0 0 1 0 0 1 1]', '百佳乐', '2024-08-27 22:09:05', '2024-08-27 22:09:05', NULL, 0, 0);
INSERT INTO `record` VALUES (466, 44, 42, -755, -5, -760, 0, '4F83E53C69B0387149726B1FF9DB6C28', '闲:[♣2,♠8,♣6,] 庄:[♦K,♦7,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 22:09:26', '2024-08-27 22:09:26', NULL, 0, 0);
INSERT INTO `record` VALUES (467, 44, 42, -760, -5, -765, 0, 'F2C959C65E6043C03F6B30CD9BE8FFFE', '闲:[♠J,♦10,♣1,] 庄:[♣1,♣6,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 22:09:47', '2024-08-27 22:09:47', NULL, 0, 0);
INSERT INTO `record` VALUES (468, 44, 42, -765, -5, -770, 0, '7E232F526EDF59D4330B498881AEAF7A', '闲:[♥2,♠9,♣10,] 庄:[♠3,♠10,♥4,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 22:10:08', '2024-08-27 22:10:08', NULL, 0, 0);
INSERT INTO `record` VALUES (469, 44, 42, -770, -5, -775, 0, '21403182D2929D198148F2910C3BC7B3', '闲:[♥2,♥Q,♥J,] 庄:[♣J,♦3,♣3,] 中奖区域:[0 0 1 0 0 0 0 1]', '百佳乐', '2024-08-27 22:10:29', '2024-08-27 22:10:29', NULL, 0, 0);
INSERT INTO `record` VALUES (470, 44, 42, -775, -5, -780, 0, '2FBAF36D851A659F20D273A2964FBC81', '闲:[♦6,♥8,♦5,] 庄:[♦1,♦J,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:10:50', '2024-08-27 22:10:50', NULL, 0, 0);
INSERT INTO `record` VALUES (471, 44, 42, -780, -5, -785, 0, '6616096DAEB160FDF04CD688B3BADD49', '闲:[♠4,♠1,♥8,] 庄:[♦Q,♣5,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 22:11:11', '2024-08-27 22:11:11', NULL, 0, 0);
INSERT INTO `record` VALUES (472, 44, 42, -785, -5, -790, 0, 'CDBF2C2D764969B670F0EA2E74144183', '闲:[♦3,♣8,♣3,] 庄:[♠6,♠Q,] 中奖区域:[0 0 1 0 0 0 1 0]', '百佳乐', '2024-08-27 22:11:32', '2024-08-27 22:11:32', NULL, 0, 0);
INSERT INTO `record` VALUES (473, 44, 42, -790, -5, -795, 0, '519FE9822D59EA87581B3E0DE08F27FA', '闲:[♠3,♦4,] 庄:[♦Q,♦2,♠5,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:11:53', '2024-08-27 22:11:53', NULL, 0, 0);
INSERT INTO `record` VALUES (474, 44, 42, -795, -5, -800, 0, 'A2C12B0FCFAFA96DA5469B4CB0A37BBC', '闲:[♠5,♣2,] 庄:[♥9,♦8,] 中奖区域:[0 1 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:12:14', '2024-08-27 22:12:14', NULL, 0, 0);
INSERT INTO `record` VALUES (475, 44, 42, -800, -5, -805, 0, 'CC1C7666BB0DD522142464D56C0B1021', '闲:[♦J,♦J,] 庄:[♣4,♥4,] 中奖区域:[0 0 1 0 1 0 1 1]', '百佳乐', '2024-08-27 22:12:35', '2024-08-27 22:12:35', NULL, 0, 0);
INSERT INTO `record` VALUES (476, 44, 42, -805, -5, -810, 0, '35D559F966B57B0A1ACCD794344AEAA3', '闲:[♣K,♣2,♣6,] 庄:[♠Q,♥3,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:12:56', '2024-08-27 22:12:56', NULL, 0, 0);
INSERT INTO `record` VALUES (477, 44, 42, -810, -5, -815, 0, '6FA8B74CBF1F916576A294779DB3758F', '闲:[♣K,♥5,♦10,] 庄:[♦K,♠3,♥7,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:13:17', '2024-08-27 22:13:17', NULL, 0, 0);
INSERT INTO `record` VALUES (478, 44, 42, -815, -5, -820, 0, 'DAD23AA91EB22626129464AAF26346D4', '闲:[♠K,♥2,♠Q,] 庄:[♣3,♣7,♣5,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:13:38', '2024-08-27 22:13:38', NULL, 0, 0);
INSERT INTO `record` VALUES (479, 44, 42, -820, -5, -825, 0, 'FEDD0753D68091129F355B09E8BBD8CC', '闲:[♠8,♣K,] 庄:[♥2,♥6,] 中奖区域:[0 1 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:13:59', '2024-08-27 22:13:59', NULL, 0, 0);
INSERT INTO `record` VALUES (480, 44, 42, -825, -5, -830, 0, 'CA9E5B686F9D3F75A7B6089A05955041', '闲:[♦10,♦K,♣K,] 庄:[♠8,♥2,♣4,] 中奖区域:[0 1 0 0 0 0 1 0]', '百佳乐', '2024-08-27 22:14:20', '2024-08-27 22:14:20', NULL, 0, 0);
INSERT INTO `record` VALUES (481, 44, 42, -830, -5, -835, 0, '793E68EAFC887C313612B94C5A17DCCD', '闲:[♥7,♣Q,] 庄:[♠10,♠J,♦5,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:14:41', '2024-08-27 22:14:41', NULL, 0, 0);
INSERT INTO `record` VALUES (482, 44, 42, -835, -5, -840, 0, '25485D59232DC0F7C066AD0E330F19E7', '闲:[♣9,♥Q,] 庄:[♦2,♠5,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:15:02', '2024-08-27 22:15:02', NULL, 0, 0);
INSERT INTO `record` VALUES (483, 44, 42, -840, -5, -845, 0, '633CAFEB0D2148ABC06B4FBBA6C06041', '闲:[♣1,♠8,] 庄:[♥3,♠4,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:15:24', '2024-08-27 22:15:24', NULL, 0, 0);
INSERT INTO `record` VALUES (484, 44, 42, -845, -5, -850, 0, 'B17618BAE894C4477DA3FABD18FC4C9F', '闲:[♦6,♥3,] 庄:[♥7,♥9,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:15:45', '2024-08-27 22:15:45', NULL, 0, 0);
INSERT INTO `record` VALUES (485, 44, 42, -850, -5, -855, 0, '2260A30C4205E15866DDFB0C4B83E0F8', '闲:[♠5,♣3,] 庄:[♦6,♠Q,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:16:06', '2024-08-27 22:16:06', NULL, 0, 0);
INSERT INTO `record` VALUES (486, 44, 42, -855, -5, -860, 0, '71AFE6E96C59327CA8BB38236C28443C', '闲:[♠5,♠2,] 庄:[♣10,♣5,♣2,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:16:27', '2024-08-27 22:16:27', NULL, 0, 0);
INSERT INTO `record` VALUES (487, 44, 42, -860, -5, -865, 0, '5118BEF224140662318E09C17222F65C', '闲:[♠J,♥5,♥7,] 庄:[♠Q,♣10,♦5,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:16:48', '2024-08-27 22:16:48', NULL, 0, 0);
INSERT INTO `record` VALUES (488, 44, 42, -865, -5, -870, 0, 'E41174A29CBAB00B2DAAC2FAAF7F9436', '闲:[♦6,♣10,] 庄:[♣2,♣4,] 中奖区域:[0 1 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:17:09', '2024-08-27 22:17:09', NULL, 0, 0);
INSERT INTO `record` VALUES (489, 44, 42, -870, -5, -875, 0, 'BF9E9CC0C4A071AC8912AE9293EBB8B9', '闲:[♠Q,♣8,] 庄:[♥5,♦5,] 中奖区域:[1 0 0 1 0 0 0 1]', '百佳乐', '2024-08-27 22:17:30', '2024-08-27 22:17:30', NULL, 0, 0);
INSERT INTO `record` VALUES (490, 44, 42, -875, -5, -880, 0, '4E0C49EA4E19F8970D2F2AC56E02826F', '闲:[♣10,♦6,] 庄:[♥10,♣6,] 中奖区域:[0 1 0 0 0 1 0 0]', '百佳乐', '2024-08-27 22:17:51', '2024-08-27 22:17:51', NULL, 0, 0);
INSERT INTO `record` VALUES (491, 44, 42, -880, -5, -885, 0, '3614C7CD0A33C122C58AF21D5181268B', '闲:[♥K,♠Q,♦1,] 庄:[♠6,♥1,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 22:18:12', '2024-08-27 22:18:12', NULL, 0, 0);
INSERT INTO `record` VALUES (492, 44, 42, -885, -5, -890, 0, '6E84AAE9E16307ABA0F269CF5D1B4DDB', '闲:[♦10,♣9,] 庄:[♦6,♣4,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:18:33', '2024-08-27 22:18:33', NULL, 0, 0);
INSERT INTO `record` VALUES (493, 44, 42, -890, -5, -895, 0, 'E187B9E69B7626991ACA696141A1CCC6', '闲:[♣8,♠6,♥J,] 庄:[♣6,♦Q,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 22:18:54', '2024-08-27 22:18:54', NULL, 0, 0);
INSERT INTO `record` VALUES (494, 44, 42, -895, -5, -900, 0, 'DE0C50B943AFDACF259A3F9292F50B6C', '闲:[♠3,♦6,] 庄:[♥8,♠9,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:19:15', '2024-08-27 22:19:15', NULL, 0, 0);
INSERT INTO `record` VALUES (495, 44, 42, -900, -5, -905, 0, 'C4EFA9C7EB9EBC4B4D90C8B6C0E0AEAF', '闲:[♠J,♥10,♣4,] 庄:[♥10,♥5,♠5,] 中奖区域:[0 0 1 0 0 0 0 1]', '百佳乐', '2024-08-27 22:19:36', '2024-08-27 22:19:36', NULL, 0, 0);
INSERT INTO `record` VALUES (496, 44, 42, -905, -5, -910, 0, 'CA565F250AEF0EFF22AA57E05DE1FCB9', '闲:[♦7,♣2,] 庄:[♣6,♦9,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:19:57', '2024-08-27 22:19:57', NULL, 0, 0);
INSERT INTO `record` VALUES (497, 44, 42, -910, -5, -915, 0, '9F7691390EE451F63E7D9DA8317990CE', '闲:[♥5,♣3,] 庄:[♣1,♦3,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:20:18', '2024-08-27 22:20:18', NULL, 0, 0);
INSERT INTO `record` VALUES (498, 44, 42, -915, -5, -920, 0, '5615029AB337BF6BCC12FCFD25D57D16', '闲:[♣2,♥6,] 庄:[♣5,♣J,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:20:39', '2024-08-27 22:20:39', NULL, 0, 0);
INSERT INTO `record` VALUES (499, 44, 42, -920, -5, -925, 0, 'F85307CE1497423E301F2C42E7076A2C', '闲:[♣2,♥Q,♠6,] 庄:[♥2,♦K,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:21:00', '2024-08-27 22:21:00', NULL, 0, 0);
INSERT INTO `record` VALUES (500, 44, 42, -925, -5, -930, 0, '0183E1063E5BC9CD38F6F273AC3C5B5A', '闲:[♣K,♣10,♠9,] 庄:[♦7,♣3,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:21:21', '2024-08-27 22:21:21', NULL, 0, 0);
INSERT INTO `record` VALUES (501, 44, 42, -930, -5, -935, 0, '08995FD3311E0A317AA31AAFF23E0BC6', '闲:[♥6,♠K,] 庄:[♦10,♠6,] 中奖区域:[0 1 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:21:42', '2024-08-27 22:21:42', NULL, 0, 0);
INSERT INTO `record` VALUES (502, 44, 42, -935, -5, -940, 0, '6192B4047622DEB58E9DBA100413F2B6', '闲:[♣5,♥5,♥7,] 庄:[♥3,♣4,] 中奖区域:[0 1 0 0 0 0 1 0]', '百佳乐', '2024-08-27 22:22:04', '2024-08-27 22:22:04', NULL, 0, 0);
INSERT INTO `record` VALUES (503, 44, 42, -940, -5, -945, 0, 'A63F67EA8FCCF23A84FAD84B08DEC6F1', '闲:[♣4,♦J,] 庄:[♥J,♠8,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 22:22:25', '2024-08-27 22:22:25', NULL, 0, 0);
INSERT INTO `record` VALUES (504, 44, 42, -945, -5, -950, 0, '5CA999D93C727DCA6B73543E88C6FF62', '闲:[♦J,♣3,] 庄:[♠9,♠10,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 22:22:46', '2024-08-27 22:22:46', NULL, 0, 0);
INSERT INTO `record` VALUES (505, 44, 42, -950, -5, -955, 0, 'D251F70F5973237B43245500DCF93BA2', '闲:[♠6,♠10,] 庄:[♦K,♣J,♥1,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:23:07', '2024-08-27 22:23:07', NULL, 0, 0);
INSERT INTO `record` VALUES (506, 44, 42, -955, -5, -960, 0, '3723C37F6FA725E56C0FFE379245547E', '闲:[♦2,♣Q,♥J,] 庄:[♥K,♥5,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 22:23:28', '2024-08-27 22:23:28', NULL, 0, 0);
INSERT INTO `record` VALUES (507, 44, 42, -960, -5, -965, 0, 'F4F9B3C01C93E3041F5281F967BA58F7', '闲:[♠8,♠10,] 庄:[♦J,♥2,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:23:49', '2024-08-27 22:23:49', NULL, 0, 0);
INSERT INTO `record` VALUES (508, 44, 42, -965, -5, -970, 0, '9CB7F38B5BAC1785A44C56E263101B82', '闲:[♦8,♣4,♣Q,] 庄:[♠4,♥J,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 22:24:10', '2024-08-27 22:24:10', NULL, 0, 0);
INSERT INTO `record` VALUES (509, 44, 42, -970, -5, -975, 0, '8EB576BDA0B293C55EB5B7D768F298B7', '闲:[♥8,♣10,] 庄:[♠6,♣2,] 中奖区域:[0 1 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:24:31', '2024-08-27 22:24:31', NULL, 0, 0);
INSERT INTO `record` VALUES (510, 44, 42, -975, -5, -980, 0, 'B6AFC9D6F1CCB5A408080B198BB3631E', '闲:[♣9,♦7,] 庄:[♣6,♥J,] 中奖区域:[0 1 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:24:52', '2024-08-27 22:24:52', NULL, 0, 0);
INSERT INTO `record` VALUES (511, 44, 42, -980, -5, -985, 0, '7AACB43CB59976967632F7FC767FD903', '闲:[♦1,♣K,♠6,] 庄:[♣3,♦1,♣3,] 中奖区域:[1 0 0 0 0 0 0 1]', '百佳乐', '2024-08-27 22:25:13', '2024-08-27 22:25:13', NULL, 0, 0);
INSERT INTO `record` VALUES (512, 44, 42, -985, -5, -990, 0, 'C80E7591C7473DD562D0405F7FF2F12C', '闲:[♥2,♠1,♥6,] 庄:[♣3,♠J,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:25:34', '2024-08-27 22:25:34', NULL, 0, 0);
INSERT INTO `record` VALUES (513, 44, 42, -990, -5, -995, 0, 'CBD270A76116A77A2A8910414B8B5B8B', '闲:[♠7,♥8,] 庄:[♦4,♦5,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 22:25:55', '2024-08-27 22:25:55', NULL, 0, 0);
INSERT INTO `record` VALUES (514, 44, 42, -995, -5, -1000, 0, '7EED3B33D41357F7BA9EDC956D44EC7E', '闲:[♦9,♦3,] 庄:[♥9,♠K,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 22:26:16', '2024-08-27 22:26:16', NULL, 0, 0);
INSERT INTO `record` VALUES (515, 44, 42, -1000, -5, -1005, 0, 'AF05BC9A2FB1D05EBA8E0C09C9162B58', '闲:[♥J,♥9,] 庄:[♣8,♥K,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:26:37', '2024-08-27 22:26:37', NULL, 0, 0);
INSERT INTO `record` VALUES (516, 44, 42, -1005, -5, -1010, 0, 'DCC8DA896E3B9B787034C7FE9CE8A9F4', '闲:[♣Q,♣J,♠8,] 庄:[♠7,♣Q,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:26:58', '2024-08-27 22:26:58', NULL, 0, 0);
INSERT INTO `record` VALUES (517, 44, 42, -1010, -5, -1015, 0, '795EA38BC936D88B9ED76D394EEF2E77', '闲:[♦10,♦K,♠4,] 庄:[♦5,♠K,♠8,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 22:27:19', '2024-08-27 22:27:19', NULL, 0, 0);
INSERT INTO `record` VALUES (518, 44, 42, -1015, -5, -1020, 0, 'CB771558089337B14A48A4CD72992DC8', '闲:[♦9,♣5,♥5,] 庄:[♥6,♣4,] 中奖区域:[1 0 0 1 0 0 1 0]', '百佳乐', '2024-08-27 22:27:40', '2024-08-27 22:27:40', NULL, 0, 0);
INSERT INTO `record` VALUES (519, 44, 42, -1020, -5, -1025, 0, '234D130185B96846A69AB39C51445212', '闲:[♠9,♦8,] 庄:[♠5,♦8,♦4,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:28:01', '2024-08-27 22:28:01', NULL, 0, 0);
INSERT INTO `record` VALUES (520, 44, 42, -1025, -5, -1030, 0, 'D6A3D6B7527070A23BE7C33BF312802A', '闲:[♦7,♦7,♣K,] 庄:[♦K,♠J,♣4,] 中奖区域:[1 0 0 0 0 0 1 0]', '百佳乐', '2024-08-27 22:28:22', '2024-08-27 22:28:22', NULL, 0, 0);
INSERT INTO `record` VALUES (521, 44, 42, -1030, -5, -1035, 0, '091661D5F9E09244F86102A7BF58B5AD', '闲:[♥8,♥3,♥10,] 庄:[♣4,♥6,♠Q,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:28:43', '2024-08-27 22:28:43', NULL, 0, 0);
INSERT INTO `record` VALUES (522, 44, 42, -1035, -5, -1040, 0, '9984CFB957CA9C7FD8A82B2076B9E66C', '闲:[♥Q,♠10,] 庄:[♥5,♦3,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 22:29:04', '2024-08-27 22:29:04', NULL, 0, 0);
INSERT INTO `record` VALUES (523, 44, 42, -1040, -5, -1045, 0, '5783CD1894B38B49792F03711E5E6653', '闲:[♦8,♣7,♠1,] 庄:[♥8,♦7,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:29:26', '2024-08-27 22:29:26', NULL, 0, 0);
INSERT INTO `record` VALUES (524, 44, 42, -1045, -5, -1050, 0, 'A38DB7DD5455AFD955D1DB01A45467E6', '闲:[♠8,♣3,♦4,] 庄:[♣5,♦K,♠9,] 中奖区域:[0 1 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:29:47', '2024-08-27 22:29:47', NULL, 0, 0);
INSERT INTO `record` VALUES (525, 44, 42, -1050, -5, -1055, 0, 'B0586D755B0BB6A5F1495B824F6C54A0', '闲:[♣4,♣J,] 庄:[♣9,♣Q,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 22:30:08', '2024-08-27 22:30:08', NULL, 0, 0);
INSERT INTO `record` VALUES (526, 44, 42, -1055, -5, -1060, 0, '1F02D99B2DA08E906B69584FDCA21F52', '闲:[♦3,♥2,♦1,] 庄:[♦2,♠8,♦6,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:30:29', '2024-08-27 22:30:29', NULL, 0, 0);
INSERT INTO `record` VALUES (527, 44, 42, -1060, -5, -1065, 0, 'DC6E705075514A885516C98D16B761B6', '闲:[♣2,♥4,] 庄:[♠7,♠K,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 22:30:50', '2024-08-27 22:30:50', NULL, 0, 0);
INSERT INTO `record` VALUES (528, 44, 42, -1065, -5, -1070, 0, '38919E157D0583EB89B479753FF6D4EF', '闲:[♠10,♥5,♣5,] 庄:[♣10,♦6,] 中奖区域:[0 0 1 0 0 0 1 0]', '百佳乐', '2024-08-27 22:31:11', '2024-08-27 22:31:11', NULL, 0, 0);
INSERT INTO `record` VALUES (529, 44, 42, -1070, -5, -1075, 0, '7D33130D11FAB0DB8B49229A47780581', '闲:[♠10,♣7,] 庄:[♣J,♠K,♣3,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:31:32', '2024-08-27 22:31:32', NULL, 0, 0);
INSERT INTO `record` VALUES (530, 44, 42, -1075, -5, -1080, 0, 'BCC7C1BDF02B9CB9B981335D69BD9895', '闲:[♦J,♦6,] 庄:[♣7,♣9,] 中奖区域:[0 1 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:31:53', '2024-08-27 22:31:53', NULL, 0, 0);
INSERT INTO `record` VALUES (531, 44, 42, -1080, -5, -1085, 0, '94DCDACABDFFD5AE730EF6ED9876C5EC', '闲:[♥1,♣8,] 庄:[♣6,♥8,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:32:14', '2024-08-27 22:32:14', NULL, 0, 0);
INSERT INTO `record` VALUES (532, 44, 42, -1085, -5, -1090, 0, '770745257CA9DD1502C23C7A55D31B61', '闲:[♥K,♠7,] 庄:[♦Q,♦J,♣K,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:32:35', '2024-08-27 22:32:35', NULL, 0, 0);
INSERT INTO `record` VALUES (533, 44, 42, -1090, -5, -1095, 0, 'D0695FC13148E7AB531795927ACFAA54', '闲:[♠4,♥2,] 庄:[♥4,♣5,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 22:32:56', '2024-08-27 22:32:56', NULL, 0, 0);
INSERT INTO `record` VALUES (534, 44, 42, -1095, -5, -1100, 0, 'C088685642B5CA5835F30A8B784C730D', '闲:[♥9,♠9,] 庄:[♦2,♠2,] 中奖区域:[1 0 0 1 0 0 1 1]', '百佳乐', '2024-08-27 22:33:17', '2024-08-27 22:33:17', NULL, 0, 0);
INSERT INTO `record` VALUES (535, 44, 42, -1100, -5, -1105, 0, '7F6EF81062BEE3C890AF358F9EAE5B2D', '闲:[♠10,♥4,♥10,] 庄:[♠K,♣4,] 中奖区域:[0 1 0 0 0 0 1 0]', '百佳乐', '2024-08-27 22:33:38', '2024-08-27 22:33:38', NULL, 0, 0);
INSERT INTO `record` VALUES (536, 44, 42, -1105, -5, -1110, 0, 'DC82576786107561E7C8BFC071DA9791', '闲:[♦Q,♣J,♣8,] 庄:[♣9,♠1,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:33:59', '2024-08-27 22:33:59', NULL, 0, 0);
INSERT INTO `record` VALUES (537, 44, 42, -1110, -5, -1115, 0, '5B9A2F8D6445D8798B804196BCFF7A26', '闲:[♥10,♣J,♥2,] 庄:[♥Q,♥K,♥4,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:34:20', '2024-08-27 22:34:20', NULL, 0, 0);
INSERT INTO `record` VALUES (538, 44, 42, -1115, -5, -1120, 0, '7452F8429AEF52875626460C48C13821', '闲:[♠9,♥7,] 庄:[♣8,♠9,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 22:34:41', '2024-08-27 22:34:41', NULL, 0, 0);
INSERT INTO `record` VALUES (539, 44, 42, -1120, -5, -1125, 0, 'B0A8039E6250C728EC92D3CFE5F2E895', '闲:[♠J,♠5,♠3,] 庄:[♣7,♣3,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:35:02', '2024-08-27 22:35:02', NULL, 0, 0);
INSERT INTO `record` VALUES (540, 44, 42, -1125, -5, -1130, 0, '72B7703C54FEBA1223244B98AF28624F', '闲:[♦5,♠Q,] 庄:[♠10,♠8,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 22:35:23', '2024-08-27 22:35:23', NULL, 0, 0);
INSERT INTO `record` VALUES (541, 44, 42, -1130, -5, -1135, 0, '7F4A55E85CB7F73C049D94F27A6FF3CE', '闲:[♦6,♣Q,] 庄:[♠5,♠5,♦2,] 中奖区域:[1 0 0 0 0 0 0 1]', '百佳乐', '2024-08-27 22:35:44', '2024-08-27 22:35:44', NULL, 0, 0);
INSERT INTO `record` VALUES (542, 44, 42, -1135, -5, -1140, 0, '467A304B62E8AB34F3905E9B19D8A58D', '闲:[♦1,♣4,♠4,] 庄:[♦1,♦5,] 中奖区域:[1 0 0 1 0 0 1 0]', '百佳乐', '2024-08-27 22:36:06', '2024-08-27 22:36:06', NULL, 0, 0);
INSERT INTO `record` VALUES (543, 44, 42, -1140, -5, -1145, 0, '3D62DC3DF73B9A3AAE20AC8C0F5CFBF1', '闲:[♠J,♣Q,] 庄:[♦1,♦8,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 22:36:27', '2024-08-27 22:36:27', NULL, 0, 0);
INSERT INTO `record` VALUES (544, 44, 42, -1145, -5, -1150, 0, '1FAA447A9B473B86C4E6B96D4702ED72', '闲:[♠8,♠3,♠7,] 庄:[♣K,♦5,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:36:48', '2024-08-27 22:36:48', NULL, 0, 0);
INSERT INTO `record` VALUES (545, 44, 42, -1150, -5, -1155, 0, '77FC3B5FCB3635EC083751AB0B434C1B', '闲:[♥7,♠8,♠K,] 庄:[♠6,♥K,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 22:37:09', '2024-08-27 22:37:09', NULL, 0, 0);
INSERT INTO `record` VALUES (546, 44, 42, -1155, -5, -1160, 0, '982F161AE19F4BE9BC6AB9A973993E7F', '闲:[♣J,♥8,] 庄:[♠J,♠2,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:37:30', '2024-08-27 22:37:30', NULL, 0, 0);
INSERT INTO `record` VALUES (547, 44, 42, -1160, -5, -1165, 0, 'C62BEC0770A15644CEFC3314AE17D735', '闲:[♥K,♥9,] 庄:[♥J,♦4,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:37:51', '2024-08-27 22:37:51', NULL, 0, 0);
INSERT INTO `record` VALUES (548, 44, 42, -1165, -5, -1170, 0, 'F845A2473836D7FD0734C6CBE86AF03A', '闲:[♥3,♥5,] 庄:[♥7,♥1,] 中奖区域:[0 1 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:38:12', '2024-08-27 22:38:12', NULL, 0, 0);
INSERT INTO `record` VALUES (549, 44, 42, -1170, -5, -1175, 0, 'E2667B9ED14D86A2CCA09F719E55A919', '闲:[♦K,♥J,♣9,] 庄:[♥K,♠5,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:38:33', '2024-08-27 22:38:33', NULL, 0, 0);
INSERT INTO `record` VALUES (550, 44, 42, -1175, -5, -1180, 0, 'AC4759EE3D40BF4EC9A227C5158CFA73', '闲:[♣9,♦10,] 庄:[♠1,♣Q,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:38:54', '2024-08-27 22:38:54', NULL, 0, 0);
INSERT INTO `record` VALUES (551, 44, 42, -1180, -5, -1185, 0, 'CE3CF6EB71694A6604517EAB1DE8F4A8', '闲:[♦3,♣1,♥5,] 庄:[♦Q,♠6,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:39:15', '2024-08-27 22:39:15', NULL, 0, 0);
INSERT INTO `record` VALUES (552, 44, 42, -1185, -5, -1190, 0, '75E800077CE02925A86B0F971DD0E9E7', '闲:[♠Q,♣6,] 庄:[♦3,♣10,♠8,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:39:36', '2024-08-27 22:39:36', NULL, 0, 0);
INSERT INTO `record` VALUES (553, 44, 42, -1190, -5, -1195, 0, '2D7C483A34E47ADE9EF23D597CADC2D1', '闲:[♥4,♠Q,] 庄:[♠8,♥J,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 22:39:57', '2024-08-27 22:39:57', NULL, 0, 0);
INSERT INTO `record` VALUES (554, 44, 42, -1195, -5, -1200, 0, '80B7EA4DFB2FF11ECD026CCAE22A30E5', '闲:[♠5,♣8,♦2,] 庄:[♥9,♦5,♦K,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:40:18', '2024-08-27 22:40:18', NULL, 0, 0);
INSERT INTO `record` VALUES (555, 44, 42, -1200, -5, -1205, 0, 'E0DE41B821BA92774457FCF448AE898C', '闲:[♣5,♣K,] 庄:[♠5,♣4,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 22:40:39', '2024-08-27 22:40:39', NULL, 0, 0);
INSERT INTO `record` VALUES (556, 44, 42, -1205, -5, -1210, 0, '5BDA90202938AD842B58C6EEE1D4FD49', '闲:[♣4,♣2,] 庄:[♠8,♠9,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 22:41:00', '2024-08-27 22:41:00', NULL, 0, 0);
INSERT INTO `record` VALUES (557, 44, 42, -1210, -5, -1215, 0, '5CEDE8F474385039DD53FD28A568DCCB', '闲:[♥5,♥1,] 庄:[♣4,♥10,♥J,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:41:21', '2024-08-27 22:41:21', NULL, 0, 0);
INSERT INTO `record` VALUES (558, 44, 42, -1215, -5, -1220, 0, 'C1185B9DF5FAA78D608486AB1492E630', '闲:[♦5,♠J,♦Q,] 庄:[♥Q,♣Q,♣3,] 中奖区域:[1 0 0 0 0 0 0 1]', '百佳乐', '2024-08-27 22:41:42', '2024-08-27 22:41:42', NULL, 0, 0);
INSERT INTO `record` VALUES (559, 44, 42, -1220, -5, -1225, 0, 'E69DFEAE1218A69031B19589844EB2B2', '闲:[♠5,♥8,♦5,] 庄:[♦10,♦4,] 中奖区域:[1 0 0 1 0 0 1 0]', '百佳乐', '2024-08-27 22:42:03', '2024-08-27 22:42:03', NULL, 0, 0);
INSERT INTO `record` VALUES (560, 44, 42, -1225, -5, -1230, 0, '88362AB0A603363D692646B5354AAD1F', '闲:[♥10,♣9,] 庄:[♠J,♠Q,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:42:24', '2024-08-27 22:42:24', NULL, 0, 0);
INSERT INTO `record` VALUES (561, 44, 42, -1230, -5, -1235, 0, '3A88B84EC3F64CCFE01E49E6833F242A', '闲:[♥7,♥7,♣K,] 庄:[♠10,♥4,] 中奖区域:[0 1 0 0 0 0 1 0]', '百佳乐', '2024-08-27 22:42:46', '2024-08-27 22:42:46', NULL, 0, 0);
INSERT INTO `record` VALUES (562, 44, 42, -1235, -5, -1240, 0, '385C6F1CF529C129B95EFAB0726408CD', '闲:[♠J,♦5,♥7,] 庄:[♠Q,♣3,♣1,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 22:43:07', '2024-08-27 22:43:07', NULL, 0, 0);
INSERT INTO `record` VALUES (563, 44, 42, -1240, -5, -1245, 0, '3E1AF7E2439072AD9B20175F60D22F26', '闲:[♣K,♥1,♦8,] 庄:[♥3,♥8,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:43:28', '2024-08-27 22:43:28', NULL, 0, 0);
INSERT INTO `record` VALUES (564, 44, 42, -1245, -5, -1250, 0, '34D4FDE7EEF686C865583DDB611F47F8', '闲:[♥2,♦K,♣3,] 庄:[♣Q,♥1,♦9,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:43:49', '2024-08-27 22:43:49', NULL, 0, 0);
INSERT INTO `record` VALUES (565, 44, 42, -1250, -5, -1255, 0, 'F9F66CB24A361EA830D17806F9B7EA31', '闲:[♥3,♠K,♠K,] 庄:[♠5,♠10,] 中奖区域:[0 0 1 0 0 0 1 0]', '百佳乐', '2024-08-27 22:44:10', '2024-08-27 22:44:10', NULL, 0, 0);
INSERT INTO `record` VALUES (566, 44, 42, -1255, -5, -1260, 0, '09CD31AF46302E156320A5D611DC5710', '闲:[♦7,♥4,♦Q,] 庄:[♦10,♠1,♦4,] 中奖区域:[0 1 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:44:31', '2024-08-27 22:44:31', NULL, 0, 0);
INSERT INTO `record` VALUES (567, 44, 42, -1260, -5, -1265, 0, 'F13B7B9677516613993060D9F783468F', '闲:[♦Q,♦4,♠1,] 庄:[♥8,♣3,♦6,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:44:52', '2024-08-27 22:44:52', NULL, 0, 0);
INSERT INTO `record` VALUES (568, 44, 42, -1265, -5, -1270, 0, 'CA30440A4987348CA7F3E0550418E560', '闲:[♦4,♣6,♠7,] 庄:[♣4,♣2,♥1,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:45:13', '2024-08-27 22:45:13', NULL, 0, 0);
INSERT INTO `record` VALUES (569, 44, 42, -1270, -5, -1275, 0, '38A6B968FFEB92A806DCE35BD751A292', '闲:[♥8,♦8,] 庄:[♥Q,♦Q,♣3,] 中奖区域:[1 0 0 0 0 0 1 1]', '百佳乐', '2024-08-27 22:45:34', '2024-08-27 22:45:34', NULL, 0, 0);
INSERT INTO `record` VALUES (570, 44, 42, -1275, -5, -1280, 0, '4538336FF3F0AE07C1244D22F41EF8E4', '闲:[♦2,♦K,♦9,] 庄:[♦K,♦6,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 22:45:55', '2024-08-27 22:45:55', NULL, 0, 0);
INSERT INTO `record` VALUES (571, 44, 42, -1280, -5, -1285, 0, 'B6B0045098658FF87BA18941C8951B9D', '闲:[♦2,♠K,] 庄:[♦10,♣9,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 22:46:16', '2024-08-27 22:46:16', NULL, 0, 0);
INSERT INTO `record` VALUES (572, 44, 42, -1285, -5, -1290, 0, '5C7A72E011CEA6C76EA7169CE315B3A5', '闲:[♠3,♦J,♥5,] 庄:[♦7,♦6,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:46:37', '2024-08-27 22:46:37', NULL, 0, 0);
INSERT INTO `record` VALUES (573, 44, 42, -1290, -5, -1295, 0, '174FDB9A904A1EAC55DDE1D5FC7ABC15', '闲:[♦8,♣6,♠Q,] 庄:[♠4,♠9,♦5,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:46:58', '2024-08-27 22:46:58', NULL, 0, 0);
INSERT INTO `record` VALUES (574, 44, 42, -1295, -5, -1300, 0, '319D035C57F362E62C99690B24BCF64D', '闲:[♠3,♥2,♥9,] 庄:[♠5,♥10,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 22:47:19', '2024-08-27 22:47:19', NULL, 0, 0);
INSERT INTO `record` VALUES (575, 44, 42, -1300, -5, -1305, 0, '934C56D245843355A98884A4693D844D', '闲:[♠6,♣Q,] 庄:[♥5,♠7,♣2,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:47:40', '2024-08-27 22:47:40', NULL, 0, 0);
INSERT INTO `record` VALUES (576, 44, 42, -1305, -5, -1310, 0, '0494FD706042FAD38B8B686028095BB1', '闲:[♠4,♥8,] 庄:[♥Q,♦9,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 22:48:01', '2024-08-27 22:48:01', NULL, 0, 0);
INSERT INTO `record` VALUES (577, 44, 42, -1310, -5, -1315, 0, '9E35E8C9B2CB389AC838F131AD1124B9', '闲:[♠6,♥8,♣2,] 庄:[♠3,♦K,♥8,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:48:22', '2024-08-27 22:48:22', NULL, 0, 0);
INSERT INTO `record` VALUES (578, 44, 42, -1315, -5, -1320, 0, '95217B91545C35A15FF0457A9C331313', '闲:[♣Q,♦9,] 庄:[♥5,♠10,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:48:43', '2024-08-27 22:48:43', NULL, 0, 0);
INSERT INTO `record` VALUES (579, 44, 42, -1320, -5, -1325, 0, 'DF859042ABAF42631C8F8EBAA39ACCF3', '闲:[♠7,♣9,] 庄:[♥8,♥1,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 22:49:04', '2024-08-27 22:49:04', NULL, 0, 0);
INSERT INTO `record` VALUES (580, 44, 42, -1325, -5, -1330, 0, 'DB62A54CFE570FF7DC7143B7AD315431', '闲:[♦K,♦4,] 庄:[♠9,♠J,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 22:49:25', '2024-08-27 22:49:25', NULL, 0, 0);
INSERT INTO `record` VALUES (581, 44, 42, -1330, -5, -1335, 0, 'EFB394B0DFB3CD88AA66C6CE9810015E', '闲:[♥2,♥8,♦8,] 庄:[♥2,♦9,] 中奖区域:[1 0 0 1 0 0 1 0]', '百佳乐', '2024-08-27 22:49:47', '2024-08-27 22:49:47', NULL, 0, 0);
INSERT INTO `record` VALUES (582, 44, 42, -1335, -5, -1340, 0, '6A2CE33247C19BA474BF0B63D6986014', '闲:[♣10,♦5,♥J,] 庄:[♥5,♦9,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:50:08', '2024-08-27 22:50:08', NULL, 0, 0);
INSERT INTO `record` VALUES (583, 44, 42, -1340, -5, -1345, 0, 'E6DF6E9B3089B7A6B1002C29EC9BC4AE', '闲:[♠4,♠J,♠3,] 庄:[♦Q,♦4,♦10,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:50:29', '2024-08-27 22:50:29', NULL, 0, 0);
INSERT INTO `record` VALUES (584, 44, 42, -1345, -5, -1350, 0, '9D98760FADD23C8398E4256ACE7E9521', '闲:[♥9,♥10,] 庄:[♠J,♥K,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:50:50', '2024-08-27 22:50:50', NULL, 0, 0);
INSERT INTO `record` VALUES (585, 44, 42, -1350, -5, -1355, 0, 'EF93CDB10ACDFC55D24A2022E70764C2', '闲:[♥4,♣6,] 庄:[♥5,♠3,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 22:51:11', '2024-08-27 22:51:11', NULL, 0, 0);
INSERT INTO `record` VALUES (586, 44, 42, -1355, -5, -1360, 0, '488F817AA59A6EA44F89EDA821FBACF1', '闲:[♦7,♥K,] 庄:[♣6,♦Q,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:51:32', '2024-08-27 22:51:32', NULL, 0, 0);
INSERT INTO `record` VALUES (587, 44, 42, -1360, -5, -1365, 0, '2F7D52EA1854F3802B474DDEF6AD3333', '闲:[♦9,♥5,] 庄:[♠2,♥7,] 中奖区域:[0 0 1 0 1 0 0 0]', '百佳乐', '2024-08-27 22:51:53', '2024-08-27 22:51:53', NULL, 0, 0);
INSERT INTO `record` VALUES (588, 44, 42, -1365, -5, -1370, 0, '30E92977D1ACF5B32B2CF8BEA022C818', '闲:[♠7,♥5,♦J,] 庄:[♥5,♥Q,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 22:52:14', '2024-08-27 22:52:14', NULL, 0, 0);
INSERT INTO `record` VALUES (589, 44, 42, -1370, -5, -1375, 0, '2CC46D92B5A946D8F89D543CDA86F1C2', '闲:[♣7,♦7,♠6,] 庄:[♥K,♠3,♦5,] 中奖区域:[0 0 1 0 0 0 1 0]', '百佳乐', '2024-08-27 22:52:35', '2024-08-27 22:52:35', NULL, 0, 0);
INSERT INTO `record` VALUES (590, 44, 42, -1375, -5, -1380, 0, '2643335BE47FA01F322BD21C6E5C4467', '闲:[♠8,♥10,] 庄:[♠3,♥Q,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:52:56', '2024-08-27 22:52:56', NULL, 0, 0);
INSERT INTO `record` VALUES (591, 44, 42, -1380, -5, -1385, 0, 'DD789A86DC4058A88189B9C3451974B3', '闲:[♦Q,♦7,] 庄:[♥5,♠9,♠6,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:53:17', '2024-08-27 22:53:17', NULL, 0, 0);
INSERT INTO `record` VALUES (592, 44, 42, -1385, -5, -1390, 0, '1504245746FA40C77396DBD0576BA21A', '闲:[♣6,♣2,] 庄:[♦9,♣8,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:53:38', '2024-08-27 22:53:38', NULL, 0, 0);
INSERT INTO `record` VALUES (593, 44, 42, -1390, -5, -1395, 0, '492FCCA15FB76800FA5C9DA4A6D16D53', '闲:[♠9,♥9,] 庄:[♠Q,♦10,] 中奖区域:[1 0 0 1 0 0 1 0]', '百佳乐', '2024-08-27 22:53:59', '2024-08-27 22:53:59', NULL, 0, 0);
INSERT INTO `record` VALUES (594, 44, 42, -1395, -5, -1400, 0, '0808E649254C485FD2BFEE7E430E1A45', '闲:[♥1,♥K,♣3,] 庄:[♣4,♠9,♠10,] 中奖区域:[1 0 0 0 0 0 0 0]', '百佳乐', '2024-08-27 22:54:20', '2024-08-27 22:54:20', NULL, 0, 0);
INSERT INTO `record` VALUES (595, 44, 42, -1400, -5, -1405, 0, '3A2746C7B71C17E7787A60755AF9EE50', '闲:[♠10,♣J,♣Q,] 庄:[♣3,♦7,♣3,] 中奖区域:[0 1 0 0 0 0 0 1]', '百佳乐', '2024-08-27 22:54:41', '2024-08-27 22:54:41', NULL, 0, 0);
INSERT INTO `record` VALUES (596, 44, 42, -1405, -5, -1410, 0, '37C60640E53E8C3E5601657B989E4D52', '闲:[♣9,♦K,] 庄:[♥1,♥10,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:55:02', '2024-08-27 22:55:02', NULL, 0, 0);
INSERT INTO `record` VALUES (597, 44, 42, -1410, -5, -1415, 0, '6B46291C85EBFE63DB39D7552504DC22', '闲:[♣1,♠Q,♠1,] 庄:[♦J,♦6,] 中奖区域:[0 0 1 0 0 0 1 0]', '百佳乐', '2024-08-27 22:55:23', '2024-08-27 22:55:23', NULL, 0, 0);
INSERT INTO `record` VALUES (598, 44, 42, -1415, -5, -1420, 0, '118387270507E30412D25B9492366952', '闲:[♦1,♦K,♦1,] 庄:[♣6,♣5,♦K,] 中奖区域:[1 0 0 0 0 0 1 0]', '百佳乐', '2024-08-27 22:55:44', '2024-08-27 22:55:44', NULL, 0, 0);
INSERT INTO `record` VALUES (599, 44, 42, -1420, -5, -1425, 0, '6BA5B0775344C80DF25F31C3DD79F388', '闲:[♠8,♦K,] 庄:[♥1,♦6,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:56:05', '2024-08-27 22:56:05', NULL, 0, 0);
INSERT INTO `record` VALUES (600, 44, 42, -1425, -5, -1430, 0, 'DC40381C1B6EACF49157F144552B07C4', '闲:[♣3,♣2,♦4,] 庄:[♠4,♠3,] 中奖区域:[1 0 0 1 0 0 0 0]', '百佳乐', '2024-08-27 22:56:26', '2024-08-27 22:56:26', NULL, 0, 0);
INSERT INTO `record` VALUES (601, 44, 42, -1430, -5, -1435, 0, '412B3FE7787A38984EB0CE6472104078', '闲:[♥9,♥1,♥2,] 庄:[♦J,♦6,] 中奖区域:[0 0 1 0 0 0 0 0]', '百佳乐', '2024-08-27 22:56:47', '2024-08-27 22:56:47', NULL, 0, 0);

-- ----------------------------
-- Table structure for room
-- ----------------------------
DROP TABLE IF EXISTS `room`;
CREATE TABLE `room`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '房间标识[一个房间只能绑定一个游戏]',
  `hostid` bigint NULL DEFAULT 0 COMMENT '房主ID',
  `level` int NULL DEFAULT 0 COMMENT '房间级别(类型:=1普通级别20人 =2中等级别100人)',
  `name` varchar(16) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '房间名称',
  `roomkey` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '房间钥匙',
  `taxation` bigint NULL DEFAULT 0 COMMENT '固定台费',
  `enterscore` bigint NULL DEFAULT -1 COMMENT '准入分数(=-1不受限制)',
  `table_count` int NULL DEFAULT 0 COMMENT '当前牌桌数',
  `max_person` int NULL DEFAULT -1 COMMENT '最大人数(=-1时不受限，仅限于系统房)',
  `max_table` int NULL DEFAULT 20 COMMENT '最大牌桌数(=-1时不受限，仅限于系统房)',
  `remark` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '备注',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `update_by` bigint NULL DEFAULT 0,
  `create_by` bigint NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 19 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of room
-- ----------------------------
INSERT INTO `room` VALUES (16, 35, 1, '我是菜鸟我骄傲', '6a9edcb7b63821802aa44d35d531c9fc', 0, 2000, 0, -1, 10, '', '2024-07-17 14:00:46', '2024-07-17 14:00:46', NULL, 0, 35);
INSERT INTO `room` VALUES (17, 35, 1, '测试房', '6a9edcb7b63821802aa44d35d531c9fc', 0, 2000, 11, 50, 20, '', '2024-07-20 17:19:51', '2024-07-20 17:19:51', NULL, 0, 35);
INSERT INTO `room` VALUES (18, 35, 1, '测试房1', '93d1fcc1f4a06a8c5a7c25257ced0f97', 0, 2000, 0, 50, 10, '', '2024-07-28 00:57:05', '2024-07-28 00:57:05', NULL, 0, 35);

-- ----------------------------
-- Table structure for skill
-- ----------------------------
DROP TABLE IF EXISTS `skill`;
CREATE TABLE `skill`  (
  `id` bigint NOT NULL COMMENT '技能ID',
  `name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '名称',
  `attack` bigint NULL DEFAULT NULL COMMENT '攻击力',
  `spellpower` bigint NULL DEFAULT NULL COMMENT '法强',
  `coefficient` bigint NULL DEFAULT NULL COMMENT '系数',
  `casttime` int NULL DEFAULT NULL COMMENT '施法前摇',
  `introduce` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '技能介绍',
  `counterattack` bigint NULL DEFAULT NULL COMMENT '反击',
  `stun` bigint NULL DEFAULT NULL COMMENT '击晕',
  `stuntime` int NULL DEFAULT NULL COMMENT '眩晕时长',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of skill
-- ----------------------------

-- ----------------------------
-- Table structure for sys_api
-- ----------------------------
DROP TABLE IF EXISTS `sys_api`;
CREATE TABLE `sys_api`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `handle` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'handle',
  `title` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '标题',
  `path` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '地址',
  `type` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '接口类型',
  `action` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '请求类型',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `create_by` bigint NULL DEFAULT NULL COMMENT '创建者',
  `update_by` bigint NULL DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_api_deleted_at`(`deleted_at` ASC) USING BTREE,
  INDEX `idx_sys_api_create_by`(`create_by` ASC) USING BTREE,
  INDEX `idx_sys_api_update_by`(`update_by` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 159 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_api
-- ----------------------------
INSERT INTO `sys_api` VALUES (5, 'go-admin/app/admin/apis.SysLoginLog.Get-fm', '登录日志通过id获取', '/api/v1/sys-login-log/:id', 'BUS', 'GET', '2021-05-13 19:59:00.728', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (6, 'go-admin/app/admin/apis.SysOperaLog.GetPage-fm', '操作日志列表', '/api/v1/sys-opera-log', 'BUS', 'GET', '2021-05-13 19:59:00.778', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (7, 'go-admin/app/admin/apis.SysOperaLog.Get-fm', '操作日志通过id获取', '/api/v1/sys-opera-log/:id', 'BUS', 'GET', '2021-05-13 19:59:00.821', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (8, 'go-admin/common/actions.IndexAction.func1', '分类列表', '/api/v1/syscategory', 'BUS', 'GET', '2021-05-13 19:59:00.870', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (9, 'go-admin/common/actions.ViewAction.func1', '分类通过id获取', '/api/v1/syscategory/:id', 'BUS', 'GET', '2021-05-13 19:59:00.945', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (10, 'go-admin/common/actions.IndexAction.func1', '内容列表', '/api/v1/syscontent', 'BUS', 'GET', '2021-05-13 19:59:01.005', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (11, 'go-admin/common/actions.ViewAction.func1', '内容通过id获取', '/api/v1/syscontent/:id', 'BUS', 'GET', '2021-05-13 19:59:01.056', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (15, 'go-admin/common/actions.IndexAction.func1', 'job列表', '/api/v1/sysjob', 'BUS', 'GET', '2021-05-13 19:59:01.248', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (16, 'go-admin/common/actions.ViewAction.func1', 'job通过id获取', '/api/v1/sysjob/:id', 'BUS', 'GET', '2021-05-13 19:59:01.298', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (21, 'go-admin/app/admin/apis.SysDictType.GetPage-fm', '字典类型列表', '/api/v1/dict/type', 'BUS', 'GET', '2021-05-13 19:59:01.525', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (22, 'go-admin/app/admin/apis.SysDictType.GetAll-fm', '字典类型查询【代码生成】', '/api/v1/dict/type-option-select', 'SYS', 'GET', '2021-05-13 19:59:01.582', '2021-06-13 20:53:48.388', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (23, 'go-admin/app/admin/apis.SysDictType.Get-fm', '字典类型通过id获取', '/api/v1/dict/type/:id', 'BUS', 'GET', '2021-05-13 19:59:01.632', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (24, 'go-admin/app/admin/apis.SysDictData.GetPage-fm', '字典数据列表', '/api/v1/dict/data', 'BUS', 'GET', '2021-05-13 19:59:01.684', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (25, 'go-admin/app/admin/apis.SysDictData.Get-fm', '字典数据通过code获取', '/api/v1/dict/data/:dictCode', 'BUS', 'GET', '2021-05-13 19:59:01.732', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (26, 'go-admin/app/admin/apis.SysDictData.GetSysDictDataAll-fm', '数据字典根据key获取', '/api/v1/dict-data/option-select', 'SYS', 'GET', '2021-05-13 19:59:01.832', '2021-06-17 11:48:40.732', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (27, 'go-admin/app/admin/apis.SysDept.GetPage-fm', '部门列表', '/api/v1/dept', 'BUS', 'GET', '2021-05-13 19:59:01.940', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (28, 'go-admin/app/admin/apis.SysDept.Get-fm', '部门通过id获取', '/api/v1/dept/:id', 'BUS', 'GET', '2021-05-13 19:59:02.009', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (29, 'go-admin/app/admin/apis.SysDept.Get2Tree-fm', '查询部门下拉树【角色权限-自定权限】', '/api/v1/deptTree', 'SYS', 'GET', '2021-05-13 19:59:02.050', '2021-06-17 11:48:40.732', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (30, 'go-admin/app/admin/apis/tools.(*Gen).GetDBTableList-fm', '数据库表列表', '/api/v1/db/tables/page', 'SYS', 'GET', '2021-05-13 19:59:02.098', '2021-06-13 20:53:48.730', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (31, 'go-admin/app/admin/apis/tools.(*Gen).GetDBColumnList-fm', '数据表列列表', '/api/v1/db/columns/page', 'SYS', 'GET', '2021-05-13 19:59:02.140', '2021-06-13 20:53:48.771', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (32, 'go-admin/app/admin/apis/tools.Gen.GenCode-fm', '数据库表生成到项目', '/api/v1/gen/toproject/:tableId', 'SYS', 'GET', '2021-05-13 19:59:02.183', '2021-06-13 20:53:48.812', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (33, 'go-admin/app/admin/apis/tools.Gen.GenMenuAndApi-fm', '数据库表生成到DB', '/api/v1/gen/todb/:tableId', 'SYS', 'GET', '2021-05-13 19:59:02.227', '2021-06-13 20:53:48.853', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (34, 'go-admin/app/admin/apis/tools.(*SysTable).GetSysTablesTree-fm', '关系表数据【代码生成】', '/api/v1/gen/tabletree', 'SYS', 'GET', '2021-05-13 19:59:02.271', '2021-06-13 20:53:48.893', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (35, 'go-admin/app/admin/apis/tools.Gen.Preview-fm', '生成预览通过id获取', '/api/v1/gen/preview/:tableId', 'SYS', 'GET', '2021-05-13 19:59:02.315', '2021-06-13 20:53:48.935', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (36, 'go-admin/app/admin/apis/tools.Gen.GenApiToFile-fm', '生成api带文件', '/api/v1/gen/apitofile/:tableId', 'SYS', 'GET', '2021-05-13 19:59:02.357', '2021-06-13 20:53:48.977', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (37, 'go-admin/app/admin/apis.System.GenerateCaptchaHandler-fm', '验证码获取', '/api/v1/getCaptcha', 'SYS', 'GET', '2021-05-13 19:59:02.405', '2021-06-13 20:53:49.020', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (38, 'go-admin/app/admin/apis.SysUser.GetInfo-fm', '用户信息获取', '/api/v1/getinfo', 'SYS', 'GET', '2021-05-13 19:59:02.447', '2021-06-13 20:53:49.065', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (39, 'go-admin/app/admin/apis.SysMenu.GetPage-fm', '菜单列表', '/api/v1/menu', 'BUS', 'GET', '2021-05-13 19:59:02.497', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (40, 'go-admin/app/admin/apis.SysMenu.GetMenuTreeSelect-fm', '查询菜单下拉树结构【废弃】', '/api/v1/menuTreeselect', 'SYS', 'GET', '2021-05-13 19:59:02.542', '2021-06-03 22:37:21.857', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (41, 'go-admin/app/admin/apis.SysMenu.Get-fm', '菜单通过id获取', '/api/v1/menu/:id', 'BUS', 'GET', '2021-05-13 19:59:02.584', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (42, 'go-admin/app/admin/apis.SysMenu.GetMenuRole-fm', '角色菜单【顶部左侧菜单】', '/api/v1/menurole', 'SYS', 'GET', '2021-05-13 19:59:02.630', '2021-06-13 20:53:49.574', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (43, 'go-admin/app/admin/apis.SysMenu.GetMenuIDS-fm', '获取角色对应的菜单id数组【废弃】', '/api/v1/menuids', 'SYS', 'GET', '2021-05-13 19:59:02.675', '2021-06-03 22:39:52.500', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (44, 'go-admin/app/admin/apis.SysRole.GetPage-fm', '角色列表', '/api/v1/role', 'BUS', 'GET', '2021-05-13 19:59:02.720', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (45, 'go-admin/app/admin/apis.SysMenu.GetMenuTreeSelect-fm', '菜单权限列表【角色配菜单使用】', '/api/v1/roleMenuTreeselect/:roleId', 'SYS', 'GET', '2021-05-13 19:59:02.762', '2021-06-17 11:48:40.732', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (46, 'go-admin/app/admin/apis.SysDept.GetDeptTreeRoleSelect-fm', '角色部门结构树【自定义数据权限】', '/api/v1/roleDeptTreeselect/:roleId', 'SYS', 'GET', '2021-05-13 19:59:02.809', '2021-06-17 11:48:40.732', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (47, 'go-admin/app/admin/apis.SysRole.Get-fm', '角色通过id获取', '/api/v1/role/:id', 'BUS', 'GET', '2021-05-13 19:59:02.850', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (48, 'github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth.(*GinJWTMiddleware).RefreshHandler-fm', '刷新token', '/api/v1/refresh_token', 'SYS', 'GET', '2021-05-13 19:59:02.892', '2021-06-13 20:53:49.278', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (53, 'go-admin/app/admin/apis.SysConfig.GetPage-fm', '参数列表', '/api/v1/config', 'BUS', 'GET', '2021-05-13 19:59:03.116', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (54, 'go-admin/app/admin/apis.SysConfig.Get-fm', '参数通过id获取', '/api/v1/config/:id', 'BUS', 'GET', '2021-05-13 19:59:03.157', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (55, 'go-admin/app/admin/apis.SysConfig.GetSysConfigByKEYForService-fm', '参数通过键名搜索【基础默认配置】', '/api/v1/configKey/:configKey', 'SYS', 'GET', '2021-05-13 19:59:03.198', '2021-06-13 20:53:49.745', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (57, 'go-admin/app/jobs/apis.SysJob.RemoveJobForService-fm', 'job移除', '/api/v1/job/remove/:id', 'BUS', 'GET', '2021-05-13 19:59:03.295', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (58, 'go-admin/app/jobs/apis.SysJob.StartJobForService-fm', 'job启动', '/api/v1/job/start/:id', 'BUS', 'GET', '2021-05-13 19:59:03.339', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (59, 'go-admin/app/admin/apis.SysPost.GetPage-fm', '岗位列表', '/api/v1/post', 'BUS', 'GET', '2021-05-13 19:59:03.381', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (60, 'go-admin/app/admin/apis.SysPost.Get-fm', '岗位通过id获取', '/api/v1/post/:id', 'BUS', 'GET', '2021-05-13 19:59:03.433', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (62, 'go-admin/app/admin/apis.SysConfig.GetSysConfigBySysApp-fm', '系统前端参数', '/api/v1/app-config', 'SYS', 'GET', '2021-05-13 19:59:03.526', '2021-06-13 20:53:49.994', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (63, 'go-admin/app/admin/apis.SysUser.GetProfile-fm', '*用户信息获取', '/api/v1/user/profile', 'SYS', 'GET', '2021-05-13 19:59:03.567', '2021-06-13 20:53:50.038', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (66, 'github.com/go-admin-team/go-admin-core/sdk/pkg/ws.(*Manager).WsClient-fm', '链接ws【定时任务log】', '/ws/:id/:channel', 'BUS', 'GET', '2021-05-13 19:59:03.705', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (67, 'github.com/go-admin-team/go-admin-core/sdk/pkg/ws.(*Manager).UnWsClient-fm', '退出ws【定时任务log】', '/wslogout/:id/:channel', 'BUS', 'GET', '2021-05-13 19:59:03.756', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (68, 'go-admin/common/middleware/handler.Ping', '*用户基本信息', '/info', 'SYS', 'GET', '2021-05-13 19:59:03.800', '2021-06-13 20:53:50.251', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (72, 'go-admin/common/actions.CreateAction.func1', '分类创建', '/api/v1/syscategory', 'BUS', 'POST', '2021-05-13 19:59:03.982', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (73, 'go-admin/common/actions.CreateAction.func1', '内容创建', '/api/v1/syscontent', 'BUS', 'POST', '2021-05-13 19:59:04.027', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (76, 'go-admin/common/actions.CreateAction.func1', 'job创建', '/api/v1/sysjob', 'BUS', 'POST', '2021-05-13 19:59:04.164', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (80, 'go-admin/app/admin/apis.SysDictData.Insert-fm', '字典数据创建', '/api/v1/dict/data', 'BUS', 'POST', '2021-05-13 19:59:04.347', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (81, 'go-admin/app/admin/apis.SysDictType.Insert-fm', '字典类型创建', '/api/v1/dict/type', 'BUS', 'POST', '2021-05-13 19:59:04.391', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (82, 'go-admin/app/admin/apis.SysDept.Insert-fm', '部门创建', '/api/v1/dept', 'BUS', 'POST', '2021-05-13 19:59:04.435', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (85, 'github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth.(*GinJWTMiddleware).LoginHandler-fm', '*登录', '/api/v1/login', 'SYS', 'POST', '2021-05-13 19:59:04.597', '2021-06-13 20:53:50.784', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (86, 'go-admin/common/middleware/handler.LogOut', '*退出', '/api/v1/logout', 'SYS', 'POST', '2021-05-13 19:59:04.642', '2021-06-13 20:53:50.824', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (87, 'go-admin/app/admin/apis.SysConfig.Insert-fm', '参数创建', '/api/v1/config', 'BUS', 'POST', '2021-05-13 19:59:04.685', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (88, 'go-admin/app/admin/apis.SysMenu.Insert-fm', '菜单创建', '/api/v1/menu', 'BUS', 'POST', '2021-05-13 19:59:04.777', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (89, 'go-admin/app/admin/apis.SysPost.Insert-fm', '岗位创建', '/api/v1/post', 'BUS', 'POST', '2021-05-13 19:59:04.886', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (90, 'go-admin/app/admin/apis.SysRole.Insert-fm', '角色创建', '/api/v1/role', 'BUS', 'POST', '2021-05-13 19:59:04.975', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (91, 'go-admin/app/admin/apis.SysUser.InsetAvatar-fm', '*用户头像编辑', '/api/v1/user/avatar', 'SYS', 'POST', '2021-05-13 19:59:05.058', '2021-06-13 20:53:51.079', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (92, 'go-admin/app/admin/apis.SysApi.Update-fm', '接口编辑', '/api/v1/sys-api/:id', 'BUS', 'PUT', '2021-05-13 19:59:05.122', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (95, 'go-admin/common/actions.UpdateAction.func1', '分类编辑', '/api/v1/syscategory/:id', 'BUS', 'PUT', '2021-05-13 19:59:05.255', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (96, 'go-admin/common/actions.UpdateAction.func1', '内容编辑', '/api/v1/syscontent/:id', 'BUS', 'PUT', '2021-05-13 19:59:05.299', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (97, 'go-admin/common/actions.UpdateAction.func1', 'job编辑', '/api/v1/sysjob', 'BUS', 'PUT', '2021-05-13 19:59:05.343', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (101, 'go-admin/app/admin/apis.SysDictData.Update-fm', '字典数据编辑', '/api/v1/dict/data/:dictCode', 'BUS', 'PUT', '2021-05-13 19:59:05.519', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (102, 'go-admin/app/admin/apis.SysDictType.Update-fm', '字典类型编辑', '/api/v1/dict/type/:id', 'BUS', 'PUT', '2021-05-13 19:59:05.569', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (103, 'go-admin/app/admin/apis.SysDept.Update-fm', '部门编辑', '/api/v1/dept/:id', 'BUS', 'PUT', '2021-05-13 19:59:05.613', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (104, 'go-admin/app/other/apis.SysFileDir.Update-fm', '文件夹编辑', '/api/v1/file-dir/:id', 'BUS', 'PUT', '2021-05-13 19:59:05.662', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (105, 'go-admin/app/other/apis.SysFileInfo.Update-fm', '文件编辑', '/api/v1/file-info/:id', 'BUS', 'PUT', '2021-05-13 19:59:05.709', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (106, 'go-admin/app/admin/apis.SysRole.Update-fm', '角色编辑', '/api/v1/role/:id', 'BUS', 'PUT', '2021-05-13 19:59:05.752', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (107, 'go-admin/app/admin/apis.SysRole.Update2DataScope-fm', '角色数据权限修改', '/api/v1/roledatascope', 'BUS', 'PUT', '2021-05-13 19:59:05.803', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (108, 'go-admin/app/admin/apis.SysConfig.Update-fm', '参数编辑', '/api/v1/config/:id', 'BUS', 'PUT', '2021-05-13 19:59:05.848', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (109, 'go-admin/app/admin/apis.SysMenu.Update-fm', '编辑菜单', '/api/v1/menu/:id', 'BUS', 'PUT', '2021-05-13 19:59:05.891', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (110, 'go-admin/app/admin/apis.SysPost.Update-fm', '岗位编辑', '/api/v1/post/:id', 'BUS', 'PUT', '2021-05-13 19:59:05.934', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (111, 'go-admin/app/admin/apis.SysUser.UpdatePwd-fm', '*用户修改密码', '/api/v1/user/pwd', 'SYS', 'PUT', '2021-05-13 19:59:05.987', '2021-06-13 20:53:51.724', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (112, 'go-admin/common/actions.DeleteAction.func1', '分类删除', '/api/v1/syscategory', 'BUS', 'DELETE', '2021-05-13 19:59:06.030', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (113, 'go-admin/common/actions.DeleteAction.func1', '内容删除', '/api/v1/syscontent', 'BUS', 'DELETE', '2021-05-13 19:59:06.076', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (114, 'go-admin/app/admin/apis.SysLoginLog.Delete-fm', '登录日志删除', '/api/v1/sys-login-log', 'BUS', 'DELETE', '2021-05-13 19:59:06.118', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (115, 'go-admin/app/admin/apis.SysOperaLog.Delete-fm', '操作日志删除', '/api/v1/sys-opera-log', 'BUS', 'DELETE', '2021-05-13 19:59:06.162', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (116, 'go-admin/common/actions.DeleteAction.func1', 'job删除', '/api/v1/sysjob', 'BUS', 'DELETE', '2021-05-13 19:59:06.206', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (117, 'go-admin/app/other/apis.SysChinaAreaData.Delete-fm', '行政区删除', '/api/v1/sys-area-data', 'BUS', 'DELETE', '2021-05-13 19:59:06.249', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (120, 'go-admin/app/admin/apis.SysDictData.Delete-fm', '字典数据删除', '/api/v1/dict/data', 'BUS', 'DELETE', '2021-05-13 19:59:06.387', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (121, 'go-admin/app/admin/apis.SysDictType.Delete-fm', '字典类型删除', '/api/v1/dict/type', 'BUS', 'DELETE', '2021-05-13 19:59:06.432', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (122, 'go-admin/app/admin/apis.SysDept.Delete-fm', '部门删除', '/api/v1/dept/:id', 'BUS', 'DELETE', '2021-05-13 19:59:06.475', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (123, 'go-admin/app/other/apis.SysFileDir.Delete-fm', '文件夹删除', '/api/v1/file-dir/:id', 'BUS', 'DELETE', '2021-05-13 19:59:06.520', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (124, 'go-admin/app/other/apis.SysFileInfo.Delete-fm', '文件删除', '/api/v1/file-info/:id', 'BUS', 'DELETE', '2021-05-13 19:59:06.566', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (125, 'go-admin/app/admin/apis.SysConfig.Delete-fm', '参数删除', '/api/v1/config', 'BUS', 'DELETE', '2021-05-13 19:59:06.612', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (126, 'go-admin/app/admin/apis.SysMenu.Delete-fm', '删除菜单', '/api/v1/menu', 'BUS', 'DELETE', '2021-05-13 19:59:06.654', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (127, 'go-admin/app/admin/apis.SysPost.Delete-fm', '岗位删除', '/api/v1/post/:id', 'BUS', 'DELETE', '2021-05-13 19:59:06.702', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (128, 'go-admin/app/admin/apis.SysRole.Delete-fm', '角色删除', '/api/v1/role', 'BUS', 'DELETE', '2021-05-13 19:59:06.746', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (131, 'github.com/go-admin-team/go-admin-core/tools/transfer.Handler.func1', '系统指标', '/api/v1/metrics', 'SYS', 'GET', '2021-05-17 22:13:55.933', '2021-06-13 20:53:49.614', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (132, 'go-admin/app/other/router.registerMonitorRouter.func1', '健康状态', '/api/v1/health', 'SYS', 'GET', '2021-05-17 22:13:56.285', '2021-06-13 20:53:49.951', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (133, 'go-admin/app/admin/apis.HelloWorld', '项目默认接口', '/', 'SYS', 'GET', '2021-05-24 10:30:44.553', '2021-06-13 20:53:47.406', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (134, 'go-admin/app/other/apis.ServerMonitor.ServerInfo-fm', '服务器基本状态', '/api/v1/server-monitor', 'SYS', 'GET', '2021-05-24 10:30:44.937', '2021-06-13 20:53:48.255', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (135, 'go-admin/app/admin/apis.SysApi.GetPage-fm', '接口列表', '/api/v1/sys-api', 'BUS', 'GET', '2021-05-24 11:37:53.303', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (136, 'go-admin/app/admin/apis.SysApi.Get-fm', '接口通过id获取', '/api/v1/sys-api/:id', 'BUS', 'GET', '2021-05-24 11:37:53.359', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (137, 'go-admin/app/admin/apis.SysLoginLog.GetPage-fm', '登录日志列表', '/api/v1/sys-login-log', 'BUS', 'GET', '2021-05-24 11:47:30.397', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (138, 'go-admin/app/other/apis.File.UploadFile-fm', '文件上传', '/api/v1/public/uploadFile', 'SYS', 'POST', '2021-05-25 19:16:18.493', '2021-06-13 20:53:50.866', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (139, 'go-admin/app/admin/apis.SysConfig.Update2Set-fm', '参数信息修改【参数配置】', '/api/v1/set-config', 'BUS', 'PUT', '2021-05-27 09:45:14.853', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (140, 'go-admin/app/admin/apis.SysConfig.Get2Set-fm', '参数获取全部【配置使用】', '/api/v1/set-config', 'BUS', 'GET', '2021-05-27 11:54:14.384', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (141, 'go-admin/app/admin/apis.SysUser.GetPage-fm', '管理员列表', '/api/v1/sys-user', 'BUS', 'GET', '2021-06-13 19:24:57.111', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (142, 'go-admin/app/admin/apis.SysUser.Get-fm', '管理员通过id获取', '/api/v1/sys-user/:id', 'BUS', 'GET', '2021-06-13 19:24:57.188', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (143, 'go-admin/app/admin/apis/tools.(*SysTable).GetSysTablesInfo-fm', '', '/api/v1/sys/tables/info', '', 'GET', '2021-06-13 19:24:57.437', '2021-06-13 20:53:48.047', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (144, 'go-admin/app/admin/apis/tools.(*SysTable).GetSysTables-fm', '', '/api/v1/sys/tables/info/:tableId', '', 'GET', '2021-06-13 19:24:57.510', '2021-06-13 20:53:48.088', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (145, 'go-admin/app/admin/apis/tools.(*SysTable).GetSysTableList-fm', '', '/api/v1/sys/tables/page', '', 'GET', '2021-06-13 19:24:57.582', '2021-06-13 20:53:48.128', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (146, 'github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1', '', '/static/*filepath', '', 'GET', '2021-06-13 19:24:59.641', '2021-06-13 20:53:50.081', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (147, 'github.com/swaggo/gin-swagger.CustomWrapHandler.func1', '', '/swagger/*any', '', 'GET', '2021-06-13 19:24:59.713', '2021-06-13 20:53:50.123', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (148, 'github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1', '', '/form-generator/*filepath', '', 'GET', '2021-06-13 19:24:59.914', '2021-06-13 20:53:50.295', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (149, 'go-admin/app/admin/apis/tools.(*SysTable).InsertSysTable-fm', '', '/api/v1/sys/tables/info', '', 'POST', '2021-06-13 19:25:00.163', '2021-06-13 20:53:50.539', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (150, 'go-admin/app/admin/apis.SysUser.Insert-fm', '管理员创建', '/api/v1/sys-user', 'BUS', 'POST', '2021-06-13 19:25:00.233', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (151, 'go-admin/app/admin/apis.SysUser.Update-fm', '管理员编辑', '/api/v1/sys-user', 'BUS', 'PUT', '2021-06-13 19:25:00.986', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (152, 'go-admin/app/admin/apis/tools.(*SysTable).UpdateSysTable-fm', '', '/api/v1/sys/tables/info', '', 'PUT', '2021-06-13 19:25:01.149', '2021-06-13 20:53:51.375', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (153, 'go-admin/app/admin/apis.SysRole.Update2Status-fm', '', '/api/v1/role-status', '', 'PUT', '2021-06-13 19:25:01.446', '2021-06-13 20:53:51.636', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (154, 'go-admin/app/admin/apis.SysUser.ResetPwd-fm', '', '/api/v1/user/pwd/reset', '', 'PUT', '2021-06-13 19:25:01.601', '2021-06-13 20:53:51.764', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (155, 'go-admin/app/admin/apis.SysUser.UpdateStatus-fm', '', '/api/v1/user/status', '', 'PUT', '2021-06-13 19:25:01.671', '2021-06-13 20:53:51.806', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (156, 'go-admin/app/admin/apis.SysUser.Delete-fm', '管理员删除', '/api/v1/sys-user', 'BUS', 'DELETE', '2021-06-13 19:25:02.043', '2023-07-23 17:14:39.697', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (157, 'go-admin/app/admin/apis/tools.(*SysTable).DeleteSysTables-fm', '', '/api/v1/sys/tables/info/:tableId', '', 'DELETE', '2021-06-13 19:25:02.283', '2021-06-13 20:53:52.367', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (158, 'github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1', '', '/static/*filepath', '', 'HEAD', '2021-06-13 19:25:02.734', '2021-06-13 20:53:52.791', NULL, 0, 0);
INSERT INTO `sys_api` VALUES (159, 'github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1', '', '/form-generator/*filepath', '', 'HEAD', '2021-06-13 19:25:02.808', '2021-06-13 20:53:52.838', NULL, 0, 0);

-- ----------------------------
-- Table structure for sys_casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `sys_casbin_rule`;
CREATE TABLE `sys_casbin_rule`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v6` varchar(25) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v7` varchar(25) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_sys_casbin_rule`(`ptype` ASC, `v0` ASC, `v1` ASC, `v2` ASC, `v3` ASC, `v4` ASC, `v5` ASC, `v6` ASC, `v7` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 152 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_casbin_rule
-- ----------------------------
INSERT INTO `sys_casbin_rule` VALUES (63, 'p', 'platform', '/api/v1/config', 'DELETE', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (20, 'p', 'platform', '/api/v1/config', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (34, 'p', 'platform', '/api/v1/config', 'POST', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (21, 'p', 'platform', '/api/v1/config/:id', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (49, 'p', 'platform', '/api/v1/config/:id', 'PUT', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (14, 'p', 'platform', '/api/v1/dept', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (33, 'p', 'platform', '/api/v1/dept', 'POST', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (60, 'p', 'platform', '/api/v1/dept/:id', 'DELETE', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (15, 'p', 'platform', '/api/v1/dept/:id', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (44, 'p', 'platform', '/api/v1/dept/:id', 'PUT', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (58, 'p', 'platform', '/api/v1/dict/data', 'DELETE', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (12, 'p', 'platform', '/api/v1/dict/data', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (31, 'p', 'platform', '/api/v1/dict/data', 'POST', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (13, 'p', 'platform', '/api/v1/dict/data/:dictCode', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (42, 'p', 'platform', '/api/v1/dict/data/:dictCode', 'PUT', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (59, 'p', 'platform', '/api/v1/dict/type', 'DELETE', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (10, 'p', 'platform', '/api/v1/dict/type', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (32, 'p', 'platform', '/api/v1/dict/type', 'POST', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (11, 'p', 'platform', '/api/v1/dict/type/:id', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (43, 'p', 'platform', '/api/v1/dict/type/:id', 'PUT', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (61, 'p', 'platform', '/api/v1/file-dir/:id', 'DELETE', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (45, 'p', 'platform', '/api/v1/file-dir/:id', 'PUT', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (62, 'p', 'platform', '/api/v1/file-info/:id', 'DELETE', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (46, 'p', 'platform', '/api/v1/file-info/:id', 'PUT', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (22, 'p', 'platform', '/api/v1/job/remove/:id', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (23, 'p', 'platform', '/api/v1/job/start/:id', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (64, 'p', 'platform', '/api/v1/menu', 'DELETE', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (16, 'p', 'platform', '/api/v1/menu', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (35, 'p', 'platform', '/api/v1/menu', 'POST', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (17, 'p', 'platform', '/api/v1/menu/:id', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (50, 'p', 'platform', '/api/v1/menu/:id', 'PUT', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (24, 'p', 'platform', '/api/v1/post', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (36, 'p', 'platform', '/api/v1/post', 'POST', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (65, 'p', 'platform', '/api/v1/post/:id', 'DELETE', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (25, 'p', 'platform', '/api/v1/post/:id', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (51, 'p', 'platform', '/api/v1/post/:id', 'PUT', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (66, 'p', 'platform', '/api/v1/role', 'DELETE', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (18, 'p', 'platform', '/api/v1/role', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (37, 'p', 'platform', '/api/v1/role', 'POST', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (19, 'p', 'platform', '/api/v1/role/:id', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (47, 'p', 'platform', '/api/v1/role/:id', 'PUT', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (48, 'p', 'platform', '/api/v1/roledatascope', 'PUT', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (71, 'p', 'platform', '/api/v1/set-config', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (70, 'p', 'platform', '/api/v1/set-config', 'PUT', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (67, 'p', 'platform', '/api/v1/sys-api', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (68, 'p', 'platform', '/api/v1/sys-api/:id', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (38, 'p', 'platform', '/api/v1/sys-api/:id', 'PUT', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (57, 'p', 'platform', '/api/v1/sys-area-data', 'DELETE', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (54, 'p', 'platform', '/api/v1/sys-login-log', 'DELETE', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (69, 'p', 'platform', '/api/v1/sys-login-log', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (1, 'p', 'platform', '/api/v1/sys-login-log/:id', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (55, 'p', 'platform', '/api/v1/sys-opera-log', 'DELETE', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (2, 'p', 'platform', '/api/v1/sys-opera-log', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (3, 'p', 'platform', '/api/v1/sys-opera-log/:id', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (76, 'p', 'platform', '/api/v1/sys-user', 'DELETE', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (72, 'p', 'platform', '/api/v1/sys-user', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (74, 'p', 'platform', '/api/v1/sys-user', 'POST', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (75, 'p', 'platform', '/api/v1/sys-user', 'PUT', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (73, 'p', 'platform', '/api/v1/sys-user/:id', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (52, 'p', 'platform', '/api/v1/syscategory', 'DELETE', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (4, 'p', 'platform', '/api/v1/syscategory', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (28, 'p', 'platform', '/api/v1/syscategory', 'POST', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (5, 'p', 'platform', '/api/v1/syscategory/:id', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (39, 'p', 'platform', '/api/v1/syscategory/:id', 'PUT', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (53, 'p', 'platform', '/api/v1/syscontent', 'DELETE', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (6, 'p', 'platform', '/api/v1/syscontent', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (29, 'p', 'platform', '/api/v1/syscontent', 'POST', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (7, 'p', 'platform', '/api/v1/syscontent/:id', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (40, 'p', 'platform', '/api/v1/syscontent/:id', 'PUT', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (56, 'p', 'platform', '/api/v1/sysjob', 'DELETE', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (8, 'p', 'platform', '/api/v1/sysjob', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (30, 'p', 'platform', '/api/v1/sysjob', 'POST', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (41, 'p', 'platform', '/api/v1/sysjob', 'PUT', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (9, 'p', 'platform', '/api/v1/sysjob/:id', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (26, 'p', 'platform', '/ws/:id/:channel', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (27, 'p', 'platform', '/wslogout/:id/:channel', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (139, 'p', 'user', '/api/v1/config', 'DELETE', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (96, 'p', 'user', '/api/v1/config', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (110, 'p', 'user', '/api/v1/config', 'POST', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (97, 'p', 'user', '/api/v1/config/:id', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (125, 'p', 'user', '/api/v1/config/:id', 'PUT', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (90, 'p', 'user', '/api/v1/dept', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (109, 'p', 'user', '/api/v1/dept', 'POST', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (136, 'p', 'user', '/api/v1/dept/:id', 'DELETE', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (91, 'p', 'user', '/api/v1/dept/:id', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (120, 'p', 'user', '/api/v1/dept/:id', 'PUT', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (134, 'p', 'user', '/api/v1/dict/data', 'DELETE', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (88, 'p', 'user', '/api/v1/dict/data', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (107, 'p', 'user', '/api/v1/dict/data', 'POST', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (89, 'p', 'user', '/api/v1/dict/data/:dictCode', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (118, 'p', 'user', '/api/v1/dict/data/:dictCode', 'PUT', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (135, 'p', 'user', '/api/v1/dict/type', 'DELETE', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (86, 'p', 'user', '/api/v1/dict/type', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (108, 'p', 'user', '/api/v1/dict/type', 'POST', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (87, 'p', 'user', '/api/v1/dict/type/:id', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (119, 'p', 'user', '/api/v1/dict/type/:id', 'PUT', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (137, 'p', 'user', '/api/v1/file-dir/:id', 'DELETE', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (121, 'p', 'user', '/api/v1/file-dir/:id', 'PUT', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (138, 'p', 'user', '/api/v1/file-info/:id', 'DELETE', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (122, 'p', 'user', '/api/v1/file-info/:id', 'PUT', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (98, 'p', 'user', '/api/v1/job/remove/:id', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (99, 'p', 'user', '/api/v1/job/start/:id', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (140, 'p', 'user', '/api/v1/menu', 'DELETE', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (92, 'p', 'user', '/api/v1/menu', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (111, 'p', 'user', '/api/v1/menu', 'POST', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (93, 'p', 'user', '/api/v1/menu/:id', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (126, 'p', 'user', '/api/v1/menu/:id', 'PUT', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (100, 'p', 'user', '/api/v1/post', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (112, 'p', 'user', '/api/v1/post', 'POST', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (141, 'p', 'user', '/api/v1/post/:id', 'DELETE', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (101, 'p', 'user', '/api/v1/post/:id', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (127, 'p', 'user', '/api/v1/post/:id', 'PUT', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (142, 'p', 'user', '/api/v1/role', 'DELETE', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (94, 'p', 'user', '/api/v1/role', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (113, 'p', 'user', '/api/v1/role', 'POST', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (95, 'p', 'user', '/api/v1/role/:id', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (123, 'p', 'user', '/api/v1/role/:id', 'PUT', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (124, 'p', 'user', '/api/v1/roledatascope', 'PUT', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (147, 'p', 'user', '/api/v1/set-config', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (146, 'p', 'user', '/api/v1/set-config', 'PUT', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (143, 'p', 'user', '/api/v1/sys-api', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (144, 'p', 'user', '/api/v1/sys-api/:id', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (114, 'p', 'user', '/api/v1/sys-api/:id', 'PUT', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (133, 'p', 'user', '/api/v1/sys-area-data', 'DELETE', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (130, 'p', 'user', '/api/v1/sys-login-log', 'DELETE', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (145, 'p', 'user', '/api/v1/sys-login-log', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (77, 'p', 'user', '/api/v1/sys-login-log/:id', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (131, 'p', 'user', '/api/v1/sys-opera-log', 'DELETE', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (78, 'p', 'user', '/api/v1/sys-opera-log', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (79, 'p', 'user', '/api/v1/sys-opera-log/:id', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (152, 'p', 'user', '/api/v1/sys-user', 'DELETE', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (148, 'p', 'user', '/api/v1/sys-user', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (150, 'p', 'user', '/api/v1/sys-user', 'POST', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (151, 'p', 'user', '/api/v1/sys-user', 'PUT', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (149, 'p', 'user', '/api/v1/sys-user/:id', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (128, 'p', 'user', '/api/v1/syscategory', 'DELETE', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (80, 'p', 'user', '/api/v1/syscategory', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (104, 'p', 'user', '/api/v1/syscategory', 'POST', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (81, 'p', 'user', '/api/v1/syscategory/:id', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (115, 'p', 'user', '/api/v1/syscategory/:id', 'PUT', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (129, 'p', 'user', '/api/v1/syscontent', 'DELETE', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (82, 'p', 'user', '/api/v1/syscontent', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (105, 'p', 'user', '/api/v1/syscontent', 'POST', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (83, 'p', 'user', '/api/v1/syscontent/:id', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (116, 'p', 'user', '/api/v1/syscontent/:id', 'PUT', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (132, 'p', 'user', '/api/v1/sysjob', 'DELETE', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (84, 'p', 'user', '/api/v1/sysjob', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (106, 'p', 'user', '/api/v1/sysjob', 'POST', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (117, 'p', 'user', '/api/v1/sysjob', 'PUT', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (85, 'p', 'user', '/api/v1/sysjob/:id', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (102, 'p', 'user', '/ws/:id/:channel', 'GET', '', '', '', '', '');
INSERT INTO `sys_casbin_rule` VALUES (103, 'p', 'user', '/wslogout/:id/:channel', 'GET', '', '', '', '', '');

-- ----------------------------
-- Table structure for sys_columns
-- ----------------------------
DROP TABLE IF EXISTS `sys_columns`;
CREATE TABLE `sys_columns`  (
  `column_id` bigint NOT NULL AUTO_INCREMENT,
  `table_id` bigint NULL DEFAULT NULL,
  `column_name` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `column_comment` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `column_type` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `go_type` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `go_field` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `json_field` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `is_pk` varchar(4) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `is_increment` varchar(4) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `is_required` varchar(4) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `is_insert` varchar(4) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `is_edit` varchar(4) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `is_list` varchar(4) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `is_query` varchar(4) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `query_type` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `html_type` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `dict_type` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `sort` bigint NULL DEFAULT NULL,
  `list` varchar(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `pk` tinyint(1) NULL DEFAULT NULL,
  `required` tinyint(1) NULL DEFAULT NULL,
  `super_column` tinyint(1) NULL DEFAULT NULL,
  `usable_column` tinyint(1) NULL DEFAULT NULL,
  `increment` tinyint(1) NULL DEFAULT NULL,
  `insert` tinyint(1) NULL DEFAULT NULL,
  `edit` tinyint(1) NULL DEFAULT NULL,
  `query` tinyint(1) NULL DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `fk_table_name` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `fk_table_name_class` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `fk_table_name_package` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `fk_label_id` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `fk_label_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `create_by` bigint NULL DEFAULT NULL COMMENT '创建者',
  `update_by` bigint NULL DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`column_id`) USING BTREE,
  INDEX `idx_sys_columns_create_by`(`create_by` ASC) USING BTREE,
  INDEX `idx_sys_columns_update_by`(`update_by` ASC) USING BTREE,
  INDEX `idx_sys_columns_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 243 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_columns
-- ----------------------------
INSERT INTO `sys_columns` VALUES (1, 1, 'id', '资产表id', 'bigint', 'int', 'Id', 'id', '1', '', '1', '1', '', '', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.057', '2023-07-21 16:08:30.057', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (2, 1, 'uid', '用户ID', 'bigint', 'int64', 'Uid', 'uid', '0', '', '0', '1', '', '1', '1', 'EQ', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.064', '2023-07-21 16:16:11.829', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (3, 1, 'goodsid', '商品id', 'bigint', 'int64', 'Goodsid', 'goodsid', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.066', '2023-07-21 16:16:11.832', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (4, 1, 'amount', '当前拥有数量', 'int', 'int64', 'Amount', 'amount', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.069', '2023-07-21 16:16:11.836', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (5, 1, 'spending', '已花费数量', 'int', 'int64', 'Spending', 'spending', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.070', '2023-07-21 16:16:11.841', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (6, 1, 'count', '累计总数量', 'bigint', 'int64', 'Count', 'count', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.072', '2023-07-21 16:16:11.844', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (7, 1, 'totalprice', '总价', 'bigint', 'int64', 'Totalprice', 'totalprice', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.074', '2023-07-21 16:16:11.847', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (8, 1, 'code', '最近操作码', 'int', 'int64', 'Code', 'code', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.077', '2023-07-21 16:16:11.850', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (9, 1, 'time', '最近消费时间', 'int', 'int64', 'Time', 'time', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 9, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.080', '2023-07-21 16:16:11.852', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (10, 1, 'remark', '备注', 'varchar(255)', 'string', 'Remark', 'remark', '0', '', '0', '1', '', '1', '', 'LIKE', 'input', '', 10, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.083', '2023-07-21 16:16:11.856', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (11, 1, 'created_at', '', 'datetime', 'time.Time', 'CreatedAt', 'createdAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 11, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.084', '2023-07-21 16:08:30.084', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (12, 1, 'updated_at', '', 'datetime', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 12, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.086', '2023-07-21 16:08:30.086', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (13, 1, 'deleted_at', '', 'datetime', 'time.Time', 'DeletedAt', 'deletedAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 13, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.089', '2023-07-21 16:08:30.089', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (14, 1, 'update_by', '', 'bigint', 'string', 'UpdateBy', 'updateBy', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 14, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.091', '2023-07-21 16:08:30.091', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (15, 1, 'create_by', '', 'bigint', 'string', 'CreateBy', 'createBy', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 15, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.093', '2023-07-21 16:08:30.093', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (16, 2, 'id', 'id ', 'bigint', 'int', 'Id', 'id', '1', '', '1', '1', '', '', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.104', '2023-07-21 16:08:30.104', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (17, 2, 'name', '名字', 'varchar(255)', 'string', 'Name', 'name', '0', '', '0', '1', '', '1', '1', 'LIKE', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.106', '2023-07-21 19:13:18.216', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (18, 2, 'announcement', '公告', 'varchar(512)', 'string', 'Announcement', 'announcement', '0', '', '0', '1', '', '1', '1', 'LIKE', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.109', '2023-07-21 19:13:18.221', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (19, 2, 'startTime', '起始时间', 'datetime(6)', 'time.Time', 'StartTime', 'startTime', '0', '', '0', '1', '', '1', '0', 'LTE', 'datetime', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.113', '2023-07-21 19:13:18.224', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (20, 2, 'endTime', '结束时间', 'datetime(6)', 'time.Time', 'EndTime', 'endTime', '0', '', '0', '1', '', '1', '0', 'GTE', 'datetime', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.116', '2023-07-21 19:13:18.229', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (21, 2, 'intentionality', '目的', 'varchar(255)', 'string', 'Intentionality', 'intentionality', '0', '', '0', '1', '', '1', '0', 'LIKE', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.117', '2023-07-21 19:13:18.234', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (22, 2, 'remark', '备注', 'varchar(255)', 'string', 'Remark', 'remark', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.120', '2023-07-21 19:13:18.239', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (23, 2, 'goods', 'pd字节格式物品奖励', 'varchar(525)', 'string', 'Goods', 'goods', '0', '', '0', '0', '', '', '', 'EQ', 'input', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.122', '2023-07-21 19:13:18.242', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (24, 2, 'created_at', '', 'datetime', 'time.Time', 'CreatedAt', 'createdAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 9, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.125', '2023-07-21 16:08:30.125', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (25, 2, 'updated_at', '', 'datetime', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 10, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.128', '2023-07-21 16:08:30.128', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (26, 2, 'deleted_at', '', 'datetime', 'time.Time', 'DeletedAt', 'deletedAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 11, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.130', '2023-07-21 16:08:30.130', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (27, 2, 'update_by', '', 'bigint', 'string', 'UpdateBy', 'updateBy', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 12, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.132', '2023-07-21 16:08:30.132', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (28, 2, 'create_by', '', 'bigint', 'string', 'CreateBy', 'createBy', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 13, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.134', '2023-07-21 16:08:30.134', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (29, 3, 'id', '游戏分类表', 'bigint', 'int', 'Id', 'id', '1', '', '1', '1', '', '', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.141', '2023-07-21 16:08:30.141', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (30, 3, 'name', '名称', 'varchar(255)', 'string', 'Name', 'name', '0', '', '0', '1', '', '1', '1', 'LIKE', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.144', '2023-07-21 20:27:20.792', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (31, 3, 'type', '类型', 'int', 'int64', 'Type', 'type', '0', '', '0', '1', '', '1', '1', 'EQ', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.147', '2023-07-21 20:27:20.795', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (32, 3, 'kind', '种类', 'int', 'int64', 'Kind', 'kind', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.151', '2023-07-21 20:27:20.799', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (33, 3, 'level', '级别', 'int', 'int64', 'Level', 'level', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.153', '2023-07-21 20:27:20.801', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (34, 3, 'remark', '备注', 'varchar(255)', 'string', 'Remark', 'remark', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.155', '2023-07-21 20:27:20.803', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (35, 3, 'created_at', '', 'datetime', 'time.Time', 'CreatedAt', 'createdAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.156', '2023-07-21 16:08:30.156', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (36, 3, 'updated_at', '', 'datetime', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.159', '2023-07-21 16:08:30.159', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (37, 3, 'deleted_at', '', 'datetime', 'time.Time', 'DeletedAt', 'deletedAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 9, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.162', '2023-07-21 16:08:30.162', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (38, 3, 'update_by', '', 'bigint', 'string', 'UpdateBy', 'updateBy', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 10, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.164', '2023-07-21 16:08:30.164', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (39, 3, 'create_by', '', 'bigint', 'string', 'CreateBy', 'createBy', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 11, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.167', '2023-07-21 16:08:30.167', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (40, 4, 'id', '', 'bigint', 'int', 'Id', 'id', '1', '', '1', '1', '', '', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.174', '2023-07-21 16:08:30.174', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (41, 4, 'name', '群名称', 'varchar(255)', 'string', 'Name', 'name', '0', '', '0', '1', '', '1', '1', 'LIKE', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.175', '2023-07-21 23:58:33.510', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (42, 4, 'hostid', '群主ID', 'bigint', 'int64', 'Hostid', 'hostid', '0', '', '0', '1', '', '1', '1', 'EQ', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.179', '2023-07-21 23:58:33.515', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (43, 4, 'setuptime', '创建时间', 'int', 'int64', 'Setuptime', 'setuptime', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.182', '2023-07-21 23:58:33.518', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (44, 4, 'explain', '群简介', 'varchar(255)', 'string', 'Explain', 'explain', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.184', '2023-07-21 23:58:33.521', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (45, 4, 'notice', '群公告', 'varchar(255)', 'string', 'Notice', 'notice', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.186', '2023-07-21 23:58:33.524', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (46, 4, 'adminlist', '管理者者', 'varchar(255)', 'string', 'Adminlist', 'adminlist', '0', '', '0', '1', '', '1', '1', 'LIKE', 'input', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.188', '2023-07-21 23:58:33.527', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (47, 4, 'memberlist', '成员列表', 'varchar(4096)', 'string', 'Memberlist', 'memberlist', '0', '', '0', '1', '', '1', '1', 'LIKE', 'input', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.190', '2023-07-21 23:58:33.532', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (48, 4, 'applylist', '申请列表', 'varchar(1024)', 'string', 'Applylist', 'applylist', '0', '', '0', '1', '', '1', '1', 'LIKE', 'input', '', 9, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.192', '2023-07-21 23:58:33.536', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (49, 4, 'bannedlist', '禁言者', 'varchar(255)', 'string', 'Bannedlist', 'bannedlist', '0', '', '0', '1', '', '1', '1', 'LIKE', 'input', '', 10, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.196', '2023-07-21 23:58:33.538', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (50, 4, 'robotid', '机器人ID', 'bigint', 'int64', 'Robotid', 'robotid', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 11, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.199', '2023-07-21 23:58:33.542', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (51, 4, 'robotkey', '机器人密钥', 'varchar(255)', 'string', 'Robotkey', 'robotkey', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 12, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.201', '2023-07-21 23:58:33.545', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (52, 4, 'robotcontrol', '机器人控制', 'int', 'int64', 'Robotcontrol', 'robotcontrol', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 13, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.203', '2023-07-21 23:58:33.550', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (53, 4, 'remark', '缓存key', 'varchar(255)', 'string', 'Remark', 'remark', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 14, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.205', '2023-07-21 23:58:33.554', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (54, 4, 'created_at', '', 'datetime', 'time.Time', 'CreatedAt', 'createdAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 15, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.207', '2023-07-21 16:08:30.207', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (55, 4, 'updated_at', '', 'datetime', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 16, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.209', '2023-07-21 16:08:30.209', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (56, 4, 'deleted_at', '', 'datetime', 'time.Time', 'DeletedAt', 'deletedAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 17, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.213', '2023-07-21 16:08:30.213', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (57, 4, 'update_by', '', 'bigint', 'string', 'UpdateBy', 'updateBy', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 18, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.215', '2023-07-21 16:08:30.215', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (58, 4, 'create_by', '', 'bigint', 'string', 'CreateBy', 'createBy', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 19, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.216', '2023-07-21 16:08:30.216', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (59, 5, 'id', '', 'bigint', 'int', 'Id', 'id', '1', '', '1', '1', '', '', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.225', '2023-07-21 16:08:30.225', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (60, 5, 'uid', '用户', 'bigint', 'int', 'Uid', 'uid', '1', '', '1', '1', '', '1', '1', 'EQ', 'input', '', 2, '', 1, 1, 0, 0, 0, 1, 0, 0, '', 'user', 'User', 'user', 'account', 'account', '2023-07-21 16:08:30.227', '2023-07-21 18:28:24.800', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (61, 5, 'timestamp', '时间戳', 'int', 'int64', 'Timestamp', 'timestamp', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.230', '2023-07-21 18:28:24.805', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (62, 5, 'remark', '备注', 'varchar(255)', 'string', 'Remark', 'remark', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.232', '2023-07-21 18:28:24.809', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (63, 5, 'created_at', '', 'datetime', 'time.Time', 'CreatedAt', 'createdAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.234', '2023-07-21 16:08:30.234', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (64, 5, 'updated_at', '', 'datetime', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.236', '2023-07-21 16:08:30.236', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (65, 5, 'deleted_at', '', 'datetime', 'time.Time', 'DeletedAt', 'deletedAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.237', '2023-07-21 16:08:30.237', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (66, 5, 'update_by', '', 'bigint', 'string', 'UpdateBy', 'updateBy', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.238', '2023-07-21 16:08:30.238', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (67, 5, 'create_by', '', 'bigint', 'string', 'CreateBy', 'createBy', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 9, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.240', '2023-07-21 16:08:30.240', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (68, 6, 'id', '', 'bigint', 'int', 'Id', 'id', '1', '', '1', '1', '', '', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.249', '2023-07-21 16:08:30.249', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (69, 6, 'accepter', '接收者账号', 'varchar(255)', 'string', 'Accepter', 'accepter', '0', '', '0', '1', '', '1', '1', 'LIKE', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.252', '2023-07-21 23:48:04.599', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (70, 6, 'sender', '发送者账号', 'varchar(255)', 'string', 'Sender', 'sender', '0', '', '0', '1', '', '1', '1', 'LIKE', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.254', '2023-07-21 23:48:04.602', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (71, 6, 'carboncopy', '抄送给', 'varchar(255)', 'string', 'Carboncopy', 'carboncopy', '0', '', '0', '1', '', '1', '', 'LIKE', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.256', '2023-07-21 23:48:04.607', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (72, 6, 'topic', '主题', 'varchar(255)', 'string', 'Topic', 'topic', '0', '', '0', '1', '', '1', '1', 'LIKE', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.258', '2023-07-21 23:48:04.611', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (73, 6, 'content', '内容', 'varchar(1024)', 'string', 'Content', 'content', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.261', '2023-07-21 23:48:04.614', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (74, 6, 'goods', 'pd字节格式物品奖励', 'varchar(525)', 'string', 'Goods', 'goods', '0', '', '0', '0', '', '', '', 'EQ', 'input', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.263', '2023-07-21 23:48:04.616', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (75, 6, 'isread', '是否已读', 'int', 'int64', 'Isread', 'isread', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.265', '2023-07-21 23:48:04.618', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (76, 6, 'timestamp', '时间', 'int', 'int64', 'Timestamp', 'timestamp', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 9, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.266', '2023-07-21 23:48:04.620', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (77, 6, 'remark', '备注', 'varchar(255)', 'string', 'Remark', 'remark', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 10, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.268', '2023-07-21 23:48:04.624', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (78, 6, 'created_at', '', 'datetime', 'time.Time', 'CreatedAt', 'createdAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 11, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.270', '2023-07-21 16:08:30.270', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (79, 6, 'updated_at', '', 'datetime', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 12, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.271', '2023-07-21 16:08:30.271', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (80, 6, 'deleted_at', '', 'datetime', 'time.Time', 'DeletedAt', 'deletedAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 13, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.273', '2023-07-21 16:08:30.273', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (81, 6, 'update_by', '', 'bigint', 'string', 'UpdateBy', 'updateBy', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 14, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.275', '2023-07-21 16:08:30.275', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (82, 6, 'create_by', '', 'bigint', 'string', 'CreateBy', 'createBy', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 15, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.277', '2023-07-21 16:08:30.277', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (83, 7, 'id', '', 'bigint', 'int', 'Id', 'id', '1', '', '1', '1', '', '', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.286', '2023-07-21 16:08:30.286', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (84, 7, 'type', '游戏类型', 'int', 'int64', 'Type', 'type', '0', '', '1', '1', '', '', '', 'EQ', 'input', '', 2, '', 0, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.289', '2023-07-21 20:43:13.751', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (85, 7, 'kindid', '游戏种类', 'int', 'int64', 'Kindid', 'kindid', '0', '', '1', '1', '', '', '', 'EQ', 'input', '', 3, '', 0, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.291', '2023-07-21 20:43:13.754', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (86, 7, 'hostid', '房主ID', 'bigint', 'int64', 'Hostid', 'hostid', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.292', '2023-07-21 20:43:13.756', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (87, 7, 'level', '游戏等级', 'int', 'int64', 'Level', 'level', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.296', '2023-07-21 20:43:13.759', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (88, 7, 'name', '桌子名称', 'varchar(255)', 'string', 'Name', 'name', '0', '', '0', '1', '', '1', '1', 'LIKE', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.298', '2023-07-21 20:43:13.762', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (89, 7, 'password', '游戏密钥', 'varchar(255)', 'string', 'Password', 'password', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.300', '2023-07-21 20:43:13.765', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (90, 7, 'maxonline', '在线人数', 'int', 'int64', 'Maxonline', 'maxonline', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.302', '2023-07-21 20:43:13.768', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (91, 7, 'amount', '剩余场次', 'int', 'int64', 'Amount', 'amount', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 9, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.305', '2023-07-21 20:43:13.771', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (92, 7, 'enterscore', '房间进场分', 'bigint', 'int64', 'Enterscore', 'enterscore', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 10, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.308', '2023-07-21 20:43:13.773', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (93, 7, 'lessscore', '房间坐下分', 'bigint', 'int64', 'Lessscore', 'lessscore', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 11, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.311', '2023-07-21 20:43:13.776', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (94, 7, 'playscore', '初始积分', 'bigint', 'int64', 'Playscore', 'playscore', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 12, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.315', '2023-07-21 20:43:13.779', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (95, 7, 'state', '状态', 'int', 'int64', 'State', 'state', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 13, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.317', '2023-07-21 20:43:13.783', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (96, 7, 'commission', '税收', 'int', 'int64', 'Commission', 'commission', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 14, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.319', '2023-07-21 20:43:13.786', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (97, 7, 'maxchair', '最大座位数', 'int', 'int64', 'Maxchair', 'maxchair', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 15, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.321', '2023-07-21 20:43:13.789', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (98, 7, 'robot_count', '机器人数量', 'int', 'int64', 'RobotCount', 'robotCount', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 16, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.323', '2023-07-21 20:43:13.791', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (99, 7, 'remark', '备注', 'varchar(255)', 'string', 'Remark', 'remark', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 17, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.324', '2023-07-21 20:43:13.794', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (100, 7, 'how_to_play', '玩法介绍', 'varchar(1024)', 'string', 'HowToPlay', 'howToPlay', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 18, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.327', '2023-07-21 20:43:13.799', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (101, 7, 'created_at', '', 'datetime', 'time.Time', 'CreatedAt', 'createdAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 19, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.330', '2023-07-21 16:08:30.330', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (102, 7, 'updated_at', '', 'datetime', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 20, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.332', '2023-07-21 16:08:30.332', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (103, 7, 'deleted_at', '', 'datetime', 'time.Time', 'DeletedAt', 'deletedAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 21, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.335', '2023-07-21 16:08:30.335', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (104, 7, 'update_by', '', 'bigint', 'string', 'UpdateBy', 'updateBy', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 22, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.337', '2023-07-21 16:08:30.337', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (105, 7, 'create_by', '', 'bigint', 'string', 'CreateBy', 'createBy', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 23, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.339', '2023-07-21 16:08:30.339', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (106, 8, 'id', '物品ID', 'bigint', 'int', 'Id', 'id', '1', '', '1', '1', '', '', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.348', '2023-07-21 16:08:30.348', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (107, 8, 'name', '物品名称', 'varchar(255)', 'string', 'Name', 'name', '0', '', '0', '1', '', '1', '1', 'LIKE', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.350', '2023-07-21 23:39:44.150', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (108, 8, 'kind', '物品类别', 'bigint', 'int64', 'Kind', 'kind', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.352', '2023-07-21 23:39:44.153', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (109, 8, 'level', '物品等级', 'int', 'int64', 'Level', 'level', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.354', '2023-07-21 23:39:44.158', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (110, 8, 'price', '单价', 'bigint', 'int64', 'Price', 'price', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.356', '2023-07-21 23:39:44.162', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (111, 8, 'store', '库存', 'int', 'int64', 'Store', 'store', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.358', '2023-07-21 23:39:44.167', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (112, 8, 'sold', '已销', 'int', 'int64', 'Sold', 'sold', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.359', '2023-07-21 23:39:44.171', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (113, 8, 'remark', '备注', 'varchar(255)', 'string', 'Remark', 'remark', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.362', '2023-07-21 23:39:44.175', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (114, 8, 'created_at', '', 'datetime', 'time.Time', 'CreatedAt', 'createdAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 9, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.365', '2023-07-21 16:08:30.365', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (115, 8, 'updated_at', '', 'datetime', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 10, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.367', '2023-07-21 16:08:30.367', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (116, 8, 'deleted_at', '', 'datetime', 'time.Time', 'DeletedAt', 'deletedAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 11, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.368', '2023-07-21 16:08:30.368', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (117, 8, 'update_by', '', 'bigint', 'string', 'UpdateBy', 'updateBy', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 12, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.370', '2023-07-21 16:08:30.370', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (118, 8, 'create_by', '', 'bigint', 'string', 'CreateBy', 'createBy', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 13, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.372', '2023-07-21 16:08:30.372', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (119, 9, 'id', 'ID', 'bigint', 'int', 'Id', 'id', '1', '', '1', '1', '', '', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.382', '2023-07-21 16:08:30.382', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (120, 9, 'name', '种类名称', 'varchar(255)', 'string', 'Name', 'name', '0', '', '0', '1', '', '1', '1', 'LIKE', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.384', '2023-07-21 20:44:20.721', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (121, 9, 'type', '游戏种类', 'int', 'int64', 'Type', 'type', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.387', '2023-07-21 20:44:20.725', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (122, 9, 'game', '游戏名称', 'varchar(255)', 'string', 'Game', 'game', '0', '', '0', '1', '', '1', '1', 'LIKE', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.390', '2023-07-21 20:44:20.729', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (123, 9, 'remark', '备注', 'varchar(255)', 'string', 'Remark', 'remark', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.394', '2023-07-21 20:44:20.733', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (124, 9, 'created_at', '', 'datetime', 'time.Time', 'CreatedAt', 'createdAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.398', '2023-07-21 16:08:30.398', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (125, 9, 'updated_at', '', 'datetime', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.401', '2023-07-21 16:08:30.401', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (126, 9, 'deleted_at', '', 'datetime', 'time.Time', 'DeletedAt', 'deletedAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.403', '2023-07-21 16:08:30.403', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (127, 9, 'update_by', '', 'bigint', 'string', 'UpdateBy', 'updateBy', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 9, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.405', '2023-07-21 16:08:30.405', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (128, 9, 'create_by', '', 'bigint', 'string', 'CreateBy', 'createBy', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 10, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:30.408', '2023-07-21 16:08:30.408', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (129, 10, 'id', '通知ID', 'int', 'int', 'Id', 'id', '1', '', '1', '1', '', '', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.541', '2023-07-21 16:08:45.541', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (130, 10, 'type', '通知类型', 'int', 'int64', 'Type', 'type', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.544', '2023-07-21 23:55:30.233', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (131, 10, 'platid', '平台ID', 'bigint', 'int64', 'Platid', 'platid', '0', '', '0', '1', '', '1', '1', 'EQ', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.546', '2023-07-21 23:55:30.237', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (132, 10, 'gameids', '指定游戏', 'varchar(255)', 'string', 'Gameids', 'gameids', '0', '', '0', '1', '', '1', '1', 'LIKE', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.548', '2023-07-21 23:55:30.241', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (133, 10, 'kindid', '关联游戏', 'bigint', 'int64', 'Kindid', 'kindid', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.552', '2023-07-21 23:55:30.246', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (134, 10, 'level', '游戏级别', 'int', 'int64', 'Level', 'level', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.554', '2023-07-21 23:55:30.249', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (135, 10, 'title', '标题', 'varchar(255)', 'string', 'Title', 'title', '0', '', '0', '1', '', '1', '1', 'LIKE', 'input', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.556', '2023-07-21 23:55:30.252', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (136, 10, 'content', '内容', 'varchar(255)', 'string', 'Content', 'content', '0', '', '0', '1', '', '1', '1', 'LIKE', 'input', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.557', '2023-07-21 23:55:30.255', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (137, 10, 'start', '起始时间', 'int', 'int64', 'Start', 'start', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 9, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.559', '2023-07-21 23:55:30.258', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (138, 10, 'end', '结束时间', 'int', 'int64', 'End', 'end', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 10, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.561', '2023-07-21 23:55:30.263', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (139, 10, 'remark', '备注', 'varchar(255)', 'string', 'Remark', 'remark', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 11, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.563', '2023-07-21 23:55:30.266', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (140, 10, 'created_at', '', 'datetime', 'time.Time', 'CreatedAt', 'createdAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 12, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.564', '2023-07-21 16:08:45.564', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (141, 10, 'updated_at', '', 'datetime', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 13, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.568', '2023-07-21 16:08:45.568', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (142, 10, 'deleted_at', '', 'datetime', 'time.Time', 'DeletedAt', 'deletedAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 14, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.570', '2023-07-21 16:08:45.570', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (143, 10, 'update_by', '', 'bigint', 'string', 'UpdateBy', 'updateBy', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 15, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.572', '2023-07-21 16:08:45.572', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (144, 10, 'create_by', '', 'bigint', 'string', 'CreateBy', 'createBy', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 16, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.574', '2023-07-21 16:08:45.574', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (145, 11, 'id', 'id', 'int', 'int', 'Id', 'id', '1', '', '1', '1', '', '', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.582', '2023-07-21 16:08:45.582', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (146, 11, 'name', '名字', 'varchar(255)', 'string', 'Name', 'name', '0', '', '0', '1', '', '1', '1', 'LIKE', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.586', '2023-07-21 18:43:50.705', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (147, 11, 'usercount', '在线人数', 'int', 'int64', 'Usercount', 'usercount', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.589', '2023-07-21 18:43:50.711', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (148, 11, 'rooms', '房间编号(以逗号分隔)', 'varchar(255)', 'string', 'Rooms', 'rooms', '0', '', '0', '1', '', '1', '', 'LIKE', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.590', '2023-07-21 18:43:50.715', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (149, 11, 'activities', '日常活动', 'varchar(255)', 'string', 'Activities', 'activities', '0', '', '0', '1', '', '1', '', 'LIKE', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.592', '2023-07-21 18:43:50.718', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (150, 11, 'remark', '备注', 'varchar(255)', 'string', 'Remark', 'remark', '0', '', '0', '1', '', '1', '', 'LIKE', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.593', '2023-07-21 18:43:50.723', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (151, 11, 'code', '与用户表里的平台ID对应', 'varchar(255)', 'string', 'Code', 'code', '0', '', '0', '1', '', '1', '', 'LIKE', 'input', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.595', '2023-07-21 18:43:50.729', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (152, 11, 'servers', '服务器地址', 'varchar(255)', 'string', 'Servers', 'servers', '0', '', '0', '1', '', '1', '', 'LIKE', 'input', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.596', '2023-07-21 18:43:50.732', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (153, 11, 'created_at', '', 'datetime', 'time.Time', 'CreatedAt', 'createdAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 9, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.598', '2023-07-21 16:08:45.598', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (154, 11, 'updated_at', '', 'datetime', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 10, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.601', '2023-07-21 16:08:45.601', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (155, 11, 'deleted_at', '', 'datetime', 'time.Time', 'DeletedAt', 'deletedAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 11, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.604', '2023-07-21 16:08:45.604', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (156, 11, 'update_by', '', 'bigint', 'string', 'UpdateBy', 'updateBy', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 12, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.606', '2023-07-21 16:08:45.606', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (157, 11, 'create_by', '', 'bigint', 'string', 'CreateBy', 'createBy', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 13, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.609', '2023-07-21 16:08:45.609', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (158, 12, 'id', '充值ID', 'bigint', 'int', 'Id', 'id', '1', '', '1', '1', '', '', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.615', '2023-07-21 16:08:45.615', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (159, 12, 'uid', '充值者', 'bigint', 'int64', 'Uid', 'uid', '0', '', '0', '1', '', '1', '1', 'LIKE', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', 'user', 'User', 'user', 'account', 'account', '2023-07-21 16:08:45.619', '2023-07-21 18:35:16.700', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (160, 12, 'byid', '代充者', 'bigint', 'int64', 'Byid', 'byid', '0', '', '0', '1', '', '1', '1', 'LIKE', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', 'user', 'User', 'user', 'account', 'account', '2023-07-21 16:08:45.620', '2023-07-21 18:35:16.704', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (161, 12, 'payment', '支付费用', 'bigint', 'int64', 'Payment', 'payment', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.622', '2023-07-21 18:35:16.708', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (162, 12, 'premoney', '充值前', 'bigint', 'int64', 'Premoney', 'premoney', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.624', '2023-07-21 18:35:16.713', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (163, 12, 'money', '充值后', 'bigint', 'int64', 'Money', 'money', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.626', '2023-07-21 18:35:16.716', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (164, 12, 'code', '充值码', 'int', 'int64', 'Code', 'code', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.627', '2023-07-21 18:35:16.720', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (165, 12, 'order', '订单号', 'varchar(255)', 'string', 'Order', 'order', '0', '', '0', '1', '', '1', '', 'LIKE', 'input', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.629', '2023-07-21 18:35:16.723', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (166, 12, 'timestamp', '充值时间', 'bigint', 'int64', 'Timestamp', 'timestamp', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 9, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.630', '2023-07-21 18:35:16.729', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (167, 12, 'remark', '备注', 'varchar(255)', 'string', 'Remark', 'remark', '0', '', '0', '1', '', '1', '', 'LIKE', 'input', '', 10, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.632', '2023-07-21 18:35:16.733', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (168, 12, 'success', '状态', 'int', 'string', 'Success', 'success', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 11, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.636', '2023-07-21 18:35:16.736', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (169, 12, 'created_at', '', 'datetime', 'time.Time', 'CreatedAt', 'createdAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 12, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.639', '2023-07-21 16:08:45.639', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (170, 12, 'updated_at', '', 'datetime', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 13, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.641', '2023-07-21 16:08:45.641', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (171, 12, 'deleted_at', '', 'datetime', 'time.Time', 'DeletedAt', 'deletedAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 14, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.642', '2023-07-21 16:08:45.642', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (172, 12, 'update_by', '', 'bigint', 'string', 'UpdateBy', 'updateBy', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 15, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.645', '2023-07-21 16:08:45.645', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (173, 12, 'create_by', '', 'bigint', 'string', 'CreateBy', 'createBy', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 16, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.647', '2023-07-21 16:08:45.647', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (174, 13, 'id', '', 'bigint', 'int', 'Id', 'id', '1', '', '1', '1', '', '', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.657', '2023-07-21 16:08:45.657', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (175, 13, 'uid', '目标用户ID', 'bigint', 'int64', 'Uid', 'uid', '0', '', '0', '1', '', '1', '1', 'EQ', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.659', '2023-07-21 22:29:47.449', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (176, 13, 'byid', '充值者ID', 'bigint', 'int64', 'Byid', 'byid', '0', '', '0', '1', '', '1', '1', 'EQ', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.661', '2023-07-21 22:29:47.453', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (177, 13, 'hostID', '房主ID', 'bigint', 'int64', 'HostID', 'hostID', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.664', '2023-07-21 22:29:47.456', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (178, 13, 'gid', '游戏ID', 'bigint', 'int64', 'Gid', 'gid', '0', '', '0', '1', '', '1', '1', 'EQ', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.667', '2023-07-21 22:29:47.460', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (179, 13, 'kid', '种类ID', 'bigint', 'int64', 'Kid', 'kid', '0', '', '0', '1', '', '1', '0', 'EQ', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.670', '2023-07-21 22:29:47.463', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (180, 13, 'level', '等级', 'int', 'int64', 'Level', 'level', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.673', '2023-07-21 22:29:47.466', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (181, 13, 'pergold', '支付之前', 'bigint', 'int64', 'Pergold', 'pergold', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.675', '2023-07-21 22:29:47.470', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (182, 13, 'payment', '支付', 'bigint', 'int64', 'Payment', 'payment', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 9, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.677', '2023-07-21 22:29:47.473', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (183, 13, 'gold', '金币', 'bigint', 'int64', 'Gold', 'gold', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 10, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.679', '2023-07-21 22:29:47.478', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (184, 13, 'code', '操作码', 'int', 'int64', 'Code', 'code', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 11, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.681', '2023-07-21 22:29:47.481', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (185, 13, 'time', '时间戳', 'int', 'int64', 'Time', 'time', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 12, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.683', '2023-07-21 22:29:47.485', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (186, 13, 'order', '订单号', 'varchar(255)', 'string', 'Order', 'order', '0', '', '0', '1', '', '1', '1', 'LIKE', 'input', '', 13, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.686', '2023-07-21 22:29:47.488', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (187, 13, 'success', '是否成功', 'int', 'int64', 'Success', 'success', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 14, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.687', '2023-07-21 22:29:47.492', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (188, 13, 'remark', '备注', 'varchar(255)', 'string', 'Remark', 'remark', '0', '', '0', '1', '', '1', '', 'LIKE', 'input', '', 15, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.688', '2023-07-21 22:29:47.495', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (189, 13, 'created_at', '', 'datetime', 'time.Time', 'CreatedAt', 'createdAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 16, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.690', '2023-07-21 16:08:45.690', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (190, 13, 'updated_at', '', 'datetime', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 17, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.691', '2023-07-21 16:08:45.691', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (191, 13, 'deleted_at', '', 'datetime', 'time.Time', 'DeletedAt', 'deletedAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 18, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.693', '2023-07-21 16:08:45.693', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (192, 13, 'update_by', '', 'bigint', 'string', 'UpdateBy', 'updateBy', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 19, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.694', '2023-07-21 16:08:45.694', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (193, 13, 'create_by', '', 'bigint', 'string', 'CreateBy', 'createBy', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 20, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.696', '2023-07-21 16:08:45.696', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (194, 14, 'id', '房间标识', 'int', 'int', 'Id', 'id', '1', '', '1', '1', '', '', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.706', '2023-07-21 16:08:45.706', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (195, 14, 'type', '房间类型', 'int', 'int64', 'Type', 'type', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.708', '2023-07-21 19:30:57.148', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (196, 14, 'num', '房间号码==roomId', 'int', 'int', 'Num', 'num', '1', '', '1', '1', '', '1', '1', 'EQ', 'input', '', 3, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.710', '2023-07-21 19:30:57.153', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (197, 14, 'name', '房间名称', 'varchar(255)', 'string', 'Name', 'name', '0', '', '0', '1', '', '1', '1', 'LIKE', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.712', '2023-07-21 19:30:57.157', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (198, 14, 'games', '游戏列表(填写游戏ID)', 'varchar(2048)', 'string', 'Games', 'games', '0', '', '0', '1', '', '1', '', 'LIKE', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.714', '2023-07-21 19:30:57.160', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (199, 14, 'roomkey', '房间钥匙', 'varchar(255)', 'string', 'Roomkey', 'roomkey', '0', '', '0', '1', '', '1', '', 'LIKE', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.717', '2023-07-21 19:30:57.164', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (200, 14, 'mark', '备注', 'varchar(255)', 'string', 'Mark', 'mark', '0', '', '0', '1', '', '1', '', 'LIKE', 'input', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.721', '2023-07-21 19:30:57.168', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (201, 14, 'created_at', '', 'datetime', 'time.Time', 'CreatedAt', 'createdAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.723', '2023-07-21 16:08:45.723', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (202, 14, 'updated_at', '', 'datetime', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 9, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.725', '2023-07-21 16:08:45.725', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (203, 14, 'deleted_at', '', 'datetime', 'time.Time', 'DeletedAt', 'deletedAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 10, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.728', '2023-07-21 16:08:45.728', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (204, 14, 'update_by', '', 'bigint', 'string', 'UpdateBy', 'updateBy', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 11, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.730', '2023-07-21 16:08:45.730', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (205, 14, 'create_by', '', 'bigint', 'string', 'CreateBy', 'createBy', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 12, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:08:45.732', '2023-07-21 16:08:45.732', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (206, 15, 'id', 'ID身份', 'bigint', 'int', 'Id', 'id', '1', '', '1', '1', '', '', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:09:02.185', '2023-07-21 16:09:02.185', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (207, 15, 'name', '姓名', 'varchar(20)', 'string', 'Name', 'name', '0', '', '0', '1', '', '1', '0', 'EQ', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:09:02.188', '2023-07-21 17:31:18.632', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (208, 15, 'account', '账号', 'varchar(255)', 'string', 'Account', 'account', '0', '', '0', '1', '', '1', '1', 'LIKE', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:09:02.192', '2023-07-21 17:31:18.635', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (209, 15, 'password', '密码', 'varchar(255)', 'string', 'Password', 'password', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:09:02.194', '2023-07-21 17:31:18.637', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (210, 15, 'head', '头像', 'varchar(255)', 'string', 'Head', 'head', '0', '', '0', '0', '', '', '', 'EQ', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:09:02.196', '2023-07-21 17:31:18.642', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (211, 15, 'face', '头像ID', 'int', 'int64', 'Face', 'face', '0', '', '0', '0', '', '', '', 'EQ', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:09:02.197', '2023-07-21 17:31:18.645', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (212, 15, 'signature', '签名', 'varchar(255)', 'string', 'Signature', 'signature', '0', '', '0', '0', '', '', '', 'EQ', 'input', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:09:02.199', '2023-07-21 17:31:18.648', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (213, 15, 'gender', '性别', 'int', 'int64', 'Gender', 'gender', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:09:02.201', '2023-07-21 17:31:18.651', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (214, 15, 'age', '年龄', 'int', 'int64', 'Age', 'age', '0', '', '0', '0', '', '', '', 'EQ', 'input', '', 9, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:09:02.203', '2023-07-21 17:31:18.654', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (215, 15, 'vip', 'VIP级别', 'int', 'string', 'Vip', 'vip', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 10, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:09:02.204', '2023-07-21 17:31:18.660', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (216, 15, 'level', '级别', 'tinyint', 'int64', 'Level', 'level', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 11, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:09:02.207', '2023-07-21 17:31:18.663', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (217, 15, 'money', '总存款', 'bigint', 'int64', 'Money', 'money', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 12, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:09:02.210', '2023-07-21 17:31:18.666', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (218, 15, 'passport', '身份ID', 'varchar(255)', 'string', 'Passport', 'passport', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 13, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:09:02.212', '2023-07-21 17:31:18.670', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (219, 15, 'realname', '真实名字', 'varchar(255)', 'string', 'Realname', 'realname', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 14, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:09:02.215', '2023-07-21 17:31:18.673', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (220, 15, 'phone', '手机', 'varchar(255)', 'string', 'Phone', 'phone', '0', '', '0', '1', '', '1', '1', 'LIKE', 'input', '', 15, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:09:02.216', '2023-07-21 17:31:18.677', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (221, 15, 'address', '住址', 'varchar(255)', 'string', 'Address', 'address', '0', '', '0', '1', '', '0', '', 'EQ', 'input', '', 16, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:09:02.218', '2023-07-21 17:31:18.679', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (222, 15, 'email', '邮箱', 'varchar(255)', 'string', 'Email', 'email', '0', '', '0', '1', '', '0', '', 'EQ', 'input', '', 17, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:09:02.220', '2023-07-21 17:31:18.681', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (223, 15, 'identity', '识别码', 'varchar(255)', 'string', 'Identity', 'identity', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 18, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:09:02.223', '2023-07-21 17:31:18.683', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (224, 15, 'agentid', '代理标识(代理人ID)', 'bigint', 'int64', 'Agentid', 'agentid', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 19, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:09:02.227', '2023-07-21 17:31:18.687', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (225, 15, 'referralcode', '推荐标识(由邀请码生成)', 'varchar(255)', 'string', 'Referralcode', 'referralcode', '0', '', '0', '0', '', '', '', 'EQ', 'input', '', 20, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:09:02.229', '2023-07-21 17:31:18.692', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (226, 15, 'serveraddr', '服务器地址(由平台指定)', 'varchar(255)', 'string', 'Serveraddr', 'serveraddr', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 21, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:09:02.231', '2023-07-21 17:31:18.695', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (227, 15, 'clientaddr', '连接地址', 'varchar(255)', 'string', 'Clientaddr', 'clientaddr', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 22, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:09:02.232', '2023-07-21 17:31:18.699', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (228, 15, 'machinecode', '机器序列', 'varchar(255)', 'string', 'Machinecode', 'machinecode', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 23, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:09:02.234', '2023-07-21 17:31:18.702', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (229, 15, 'deposit', '现有存款', 'bigint', 'int64', 'Deposit', 'deposit', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 24, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:09:02.236', '2023-07-21 17:31:18.705', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (230, 15, 'withdraw', '总提款数目', 'bigint', 'int64', 'Withdraw', 'withdraw', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 25, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:09:02.237', '2023-07-21 17:31:18.710', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (231, 15, 'signintime', '注册时间', 'bigint', 'int64', 'Signintime', 'signintime', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 26, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:09:02.240', '2023-07-21 17:31:18.712', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (232, 15, 'logintime', '登陆时间', 'bigint', 'int64', 'Logintime', 'logintime', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 27, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:09:02.242', '2023-07-21 17:31:18.714', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (233, 15, 'leavetime', '离开时间', 'bigint', 'int64', 'Leavetime', 'leavetime', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 28, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:09:02.244', '2023-07-21 17:31:18.718', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (234, 15, 'platformid', '平台ID', 'int', 'int64', 'Platformid', 'platformid', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 29, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:09:02.246', '2023-07-21 17:31:18.721', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (235, 15, 'roomnums', '房间列表(VIP可以指定房间)', 'varchar(255)', 'string', 'Roomnums', 'roomnums', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 30, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:09:02.248', '2023-07-21 17:31:18.726', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (236, 15, 'friends', '朋友列表', 'varchar(2048)', 'string', 'Friends', 'friends', '0', '', '0', '0', '', '', '', 'EQ', 'input', '', 31, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:09:02.250', '2023-07-21 17:31:18.729', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (237, 15, 'applylist', '请求通过的好友列表', 'varchar(2048)', 'string', 'Applylist', 'applylist', '0', '', '0', '0', '', '', '', 'EQ', 'input', '', 32, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:09:02.251', '2023-07-21 17:31:18.732', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (238, 15, 'remark', '备注', 'varchar(255)', 'string', 'Remark', 'remark', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 33, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:09:02.252', '2023-07-21 17:31:18.735', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (239, 15, 'created_at', '', 'datetime', 'time.Time', 'CreatedAt', 'createdAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 34, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:09:02.253', '2023-07-21 16:09:02.253', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (240, 15, 'updated_at', '', 'datetime', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 35, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:09:02.256', '2023-07-21 16:09:02.256', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (241, 15, 'deleted_at', '', 'datetime', 'time.Time', 'DeletedAt', 'deletedAt', '0', '', '0', '1', '', '', '', 'EQ', 'datetime', '', 36, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:09:02.259', '2023-07-21 16:09:02.259', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (242, 15, 'update_by', '', 'bigint', 'string', 'UpdateBy', 'updateBy', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 37, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:09:02.261', '2023-07-21 16:09:02.261', NULL, 0, 0);
INSERT INTO `sys_columns` VALUES (243, 15, 'create_by', '', 'bigint', 'string', 'CreateBy', 'createBy', '0', '', '0', '1', '', '', '', 'EQ', 'input', '', 38, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', '2023-07-21 16:09:02.262', '2023-07-21 16:09:02.262', NULL, 0, 0);

-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `config_name` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'ConfigName',
  `config_key` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'ConfigKey',
  `config_value` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'ConfigValue',
  `config_type` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'ConfigType',
  `is_frontend` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '是否前台',
  `remark` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'Remark',
  `create_by` bigint NULL DEFAULT NULL COMMENT '创建者',
  `update_by` bigint NULL DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_config_create_by`(`create_by` ASC) USING BTREE,
  INDEX `idx_sys_config_update_by`(`update_by` ASC) USING BTREE,
  INDEX `idx_sys_config_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_config
-- ----------------------------
INSERT INTO `sys_config` VALUES (1, '皮肤样式', 'sys_index_skinName', 'skin-green', 'Y', '0', '主框架页-默认皮肤样式名称:蓝色 skin-blue、绿色 skin-green、紫色 skin-purple、红色 skin-red、黄色 skin-yellow', 1, 1, '2021-05-13 19:56:37.913', '2021-06-05 13:50:13.123', NULL);
INSERT INTO `sys_config` VALUES (2, '初始密码', 'sys_user_initPassword', '123456', 'Y', '0', '用户管理-账号初始密码:123456', 1, 1, '2021-05-13 19:56:37.913', '2021-05-13 19:56:37.913', NULL);
INSERT INTO `sys_config` VALUES (3, '侧栏主题', 'sys_index_sideTheme', 'theme-dark', 'Y', '0', '主框架页-侧边栏主题:深色主题theme-dark，浅色主题theme-light', 1, 1, '2021-05-13 19:56:37.913', '2021-05-13 19:56:37.913', NULL);
INSERT INTO `sys_config` VALUES (4, '系统名称', 'sys_app_name', '宏海添航管理系统', 'Y', '1', '', 1, 0, '2021-03-17 08:52:06.067', '2021-05-28 10:08:25.248', NULL);
INSERT INTO `sys_config` VALUES (5, '系统logo', 'sys_app_logo', 'https://doc-image.zhangwj.com/img/go-admin.png', 'Y', '1', '', 1, 0, '2021-03-17 08:53:19.462', '2021-03-17 08:53:19.462', NULL);

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept`  (
  `dept_id` bigint NOT NULL AUTO_INCREMENT,
  `parent_id` bigint NULL DEFAULT NULL,
  `dept_path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `dept_name` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `sort` tinyint NULL DEFAULT NULL,
  `leader` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `phone` varchar(11) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `email` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `status` tinyint NULL DEFAULT NULL,
  `create_by` bigint NULL DEFAULT NULL COMMENT '创建者',
  `update_by` bigint NULL DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`dept_id`) USING BTREE,
  INDEX `idx_sys_dept_update_by`(`update_by` ASC) USING BTREE,
  INDEX `idx_sys_dept_deleted_at`(`deleted_at` ASC) USING BTREE,
  INDEX `idx_sys_dept_create_by`(`create_by` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
INSERT INTO `sys_dept` VALUES (1, 0, '/0/1/', '爱拓科技', 0, 'aituo', '13782218188', 'atuo@aituo.com', 2, 1, 1, '2021-05-13 19:56:37.913', '2021-06-05 17:06:44.960', NULL);
INSERT INTO `sys_dept` VALUES (7, 1, '/0/1/7/', '研发部', 1, 'aituo', '13782218188', 'atuo@aituo.com', 2, 1, 1, '2021-05-13 19:56:37.913', '2021-06-16 21:35:00.109', NULL);
INSERT INTO `sys_dept` VALUES (8, 1, '/0/1/8/', '运维部', 0, 'aituo', '13782218188', 'atuo@aituo.com', 2, 1, 1, '2021-05-13 19:56:37.913', '2021-06-16 21:41:39.747', NULL);
INSERT INTO `sys_dept` VALUES (9, 1, '/0/1/9/', '客服部', 0, 'aituo', '13782218188', 'atuo@aituo.com', 2, 1, 1, '2021-05-13 19:56:37.913', '2021-06-05 17:07:05.993', NULL);
INSERT INTO `sys_dept` VALUES (10, 1, '/0/1/10/', '人力资源', 3, 'aituo', '13782218188', 'atuo@aituo.com', 1, 1, 1, '2021-05-13 19:56:37.913', '2021-06-05 17:07:08.503', NULL);
INSERT INTO `sys_dept` VALUES (11, 0, '/0/11/', '海航企业管理', 10, 'killy', '13500000022', '2437854120@qq.com', 2, 0, 0, '2023-07-23 16:17:09.371', '2023-07-23 16:17:37.568', NULL);
INSERT INTO `sys_dept` VALUES (12, 11, '/0/11/12/', '技术部', 10, 'majini', '13500000033', '2437854120@qq.com', 2, 0, 0, '2023-07-23 16:18:07.647', '2023-07-23 16:18:07.650', NULL);

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_data`;
CREATE TABLE `sys_dict_data`  (
  `dict_code` bigint NOT NULL AUTO_INCREMENT,
  `dict_sort` bigint NULL DEFAULT NULL,
  `dict_label` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `dict_value` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `dict_type` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `css_class` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `list_class` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `is_default` varchar(8) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `status` tinyint NULL DEFAULT NULL,
  `default` varchar(8) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `create_by` bigint NULL DEFAULT NULL COMMENT '创建者',
  `update_by` bigint NULL DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`dict_code`) USING BTREE,
  INDEX `idx_sys_dict_data_create_by`(`create_by` ASC) USING BTREE,
  INDEX `idx_sys_dict_data_update_by`(`update_by` ASC) USING BTREE,
  INDEX `idx_sys_dict_data_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 33 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dict_data
-- ----------------------------
INSERT INTO `sys_dict_data` VALUES (1, 0, '正常', '2', 'sys_normal_disable', '', '', '', 2, '', '系统正常', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:40.168', NULL);
INSERT INTO `sys_dict_data` VALUES (2, 0, '停用', '1', 'sys_normal_disable', '', '', '', 2, '', '系统停用', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_data` VALUES (3, 0, '男', '0', 'sys_user_sex', '', '', '', 2, '', '性别男', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_data` VALUES (4, 0, '女', '1', 'sys_user_sex', '', '', '', 2, '', '性别女', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_data` VALUES (5, 0, '未知', '2', 'sys_user_sex', '', '', '', 2, '', '性别未知', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_data` VALUES (6, 0, '显示', '0', 'sys_show_hide', '', '', '', 2, '', '显示菜单', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_data` VALUES (7, 0, '隐藏', '1', 'sys_show_hide', '', '', '', 2, '', '隐藏菜单', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_data` VALUES (8, 0, '是', 'Y', 'sys_yes_no', '', '', '', 2, '', '系统默认是', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_data` VALUES (9, 0, '否', 'N', 'sys_yes_no', '', '', '', 2, '', '系统默认否', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_data` VALUES (10, 0, '正常', '2', 'sys_job_status', '', '', '', 2, '', '正常状态', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_data` VALUES (11, 0, '停用', '1', 'sys_job_status', '', '', '', 2, '', '停用状态', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_data` VALUES (12, 0, '默认', 'DEFAULT', 'sys_job_group', '', '', '', 2, '', '默认分组', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_data` VALUES (13, 0, '系统', 'SYSTEM', 'sys_job_group', '', '', '', 2, '', '系统分组', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_data` VALUES (14, 0, '通知', '1', 'sys_notice_type', '', '', '', 2, '', '通知', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_data` VALUES (15, 0, '公告', '2', 'sys_notice_type', '', '', '', 2, '', '公告', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_data` VALUES (16, 0, '正常', '2', 'sys_common_status', '', '', '', 2, '', '正常状态', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_data` VALUES (17, 0, '关闭', '1', 'sys_common_status', '', '', '', 2, '', '关闭状态', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_data` VALUES (18, 0, '新增', '1', 'sys_oper_type', '', '', '', 2, '', '新增操作', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_data` VALUES (19, 0, '修改', '2', 'sys_oper_type', '', '', '', 2, '', '修改操作', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_data` VALUES (20, 0, '删除', '3', 'sys_oper_type', '', '', '', 2, '', '删除操作', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_data` VALUES (21, 0, '授权', '4', 'sys_oper_type', '', '', '', 2, '', '授权操作', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_data` VALUES (22, 0, '导出', '5', 'sys_oper_type', '', '', '', 2, '', '导出操作', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_data` VALUES (23, 0, '导入', '6', 'sys_oper_type', '', '', '', 2, '', '导入操作', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_data` VALUES (24, 0, '强退', '7', 'sys_oper_type', '', '', '', 2, '', '强退操作', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_data` VALUES (25, 0, '生成代码', '8', 'sys_oper_type', '', '', '', 2, '', '生成操作', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_data` VALUES (26, 0, '清空数据', '9', 'sys_oper_type', '', '', '', 2, '', '清空操作', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_data` VALUES (27, 0, '成功', '0', 'sys_notice_status', '', '', '', 2, '', '成功状态', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_data` VALUES (28, 0, '失败', '1', 'sys_notice_status', '', '', '', 2, '', '失败状态', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_data` VALUES (29, 0, '登录', '10', 'sys_oper_type', '', '', '', 2, '', '登录操作', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_data` VALUES (30, 0, '退出', '11', 'sys_oper_type', '', '', '', 2, '', '', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_data` VALUES (31, 0, '获取验证码', '12', 'sys_oper_type', '', '', '', 2, '', '获取验证码', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_data` VALUES (32, 0, '正常', '1', 'sys_content_status', '', '', '', 1, '', '', 1, 1, '2021-05-13 19:56:40.845', '2021-05-13 19:56:40.845', NULL);
INSERT INTO `sys_dict_data` VALUES (33, 1, '禁用', '2', 'sys_content_status', '', '', '', 1, '', '', 1, 1, '2021-05-13 19:56:40.845', '2021-05-13 19:56:40.845', NULL);

-- ----------------------------
-- Table structure for sys_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_type`;
CREATE TABLE `sys_dict_type`  (
  `dict_id` bigint NOT NULL AUTO_INCREMENT,
  `dict_name` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `dict_type` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `status` tinyint NULL DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `create_by` bigint NULL DEFAULT NULL COMMENT '创建者',
  `update_by` bigint NULL DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`dict_id`) USING BTREE,
  INDEX `idx_sys_dict_type_create_by`(`create_by` ASC) USING BTREE,
  INDEX `idx_sys_dict_type_update_by`(`update_by` ASC) USING BTREE,
  INDEX `idx_sys_dict_type_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dict_type
-- ----------------------------
INSERT INTO `sys_dict_type` VALUES (1, '系统开关', 'sys_normal_disable', 2, '系统开关列表', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_type` VALUES (2, '用户性别', 'sys_user_sex', 2, '用户性别列表', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_type` VALUES (3, '菜单状态', 'sys_show_hide', 2, '菜单状态列表', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_type` VALUES (4, '系统是否', 'sys_yes_no', 2, '系统是否列表', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_type` VALUES (5, '任务状态', 'sys_job_status', 2, '任务状态列表', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_type` VALUES (6, '任务分组', 'sys_job_group', 2, '任务分组列表', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_type` VALUES (7, '通知类型', 'sys_notice_type', 2, '通知类型列表', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_type` VALUES (8, '系统状态', 'sys_common_status', 2, '登录状态列表', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_type` VALUES (9, '操作类型', 'sys_oper_type', 2, '操作类型列表', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_type` VALUES (10, '通知状态', 'sys_notice_status', 2, '通知状态列表', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:37.914', NULL);
INSERT INTO `sys_dict_type` VALUES (11, '内容状态', 'sys_content_status', 2, '', 1, 1, '2021-05-13 19:56:40.813', '2021-05-13 19:56:40.813', NULL);

-- ----------------------------
-- Table structure for sys_job
-- ----------------------------
DROP TABLE IF EXISTS `sys_job`;
CREATE TABLE `sys_job`  (
  `job_id` bigint NOT NULL AUTO_INCREMENT,
  `job_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `job_group` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `job_type` tinyint NULL DEFAULT NULL,
  `cron_expression` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `invoke_target` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `args` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `misfire_policy` bigint NULL DEFAULT NULL,
  `concurrent` tinyint NULL DEFAULT NULL,
  `status` tinyint NULL DEFAULT NULL,
  `entry_id` smallint NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `create_by` bigint NULL DEFAULT NULL COMMENT '创建者',
  `update_by` bigint NULL DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`job_id`) USING BTREE,
  INDEX `idx_sys_job_create_by`(`create_by` ASC) USING BTREE,
  INDEX `idx_sys_job_update_by`(`update_by` ASC) USING BTREE,
  INDEX `idx_sys_job_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_job
-- ----------------------------
INSERT INTO `sys_job` VALUES (1, '接口测试', 'DEFAULT', 1, '0/5 * * * * ', 'http://localhost:8000', '', 1, 1, 1, 0, '2021-05-13 19:56:37.914', '2021-06-14 20:59:55.417', NULL, 1, 1);
INSERT INTO `sys_job` VALUES (2, '函数测试', 'DEFAULT', 2, '0/5 * * * * ', 'ExamplesOne', '参数', 1, 1, 1, 0, '2021-05-13 19:56:37.914', '2021-05-31 23:55:37.221', NULL, 1, 1);

-- ----------------------------
-- Table structure for sys_login_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_login_log`;
CREATE TABLE `sys_login_log`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `username` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户名',
  `status` varchar(4) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '状态',
  `ipaddr` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'ip地址',
  `login_location` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '归属地',
  `browser` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '浏览器',
  `os` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '系统',
  `platform` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '固件',
  `login_time` timestamp NULL DEFAULT NULL COMMENT '登录时间',
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  `msg` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '信息',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '最后更新时间',
  `create_by` bigint NULL DEFAULT NULL COMMENT '创建者',
  `update_by` bigint NULL DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_login_log_create_by`(`create_by` ASC) USING BTREE,
  INDEX `idx_sys_login_log_update_by`(`update_by` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_login_log
-- ----------------------------

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu`  (
  `menu_id` bigint NOT NULL AUTO_INCREMENT,
  `menu_name` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `title` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `icon` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `path` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `paths` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `menu_type` varchar(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `action` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `permission` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `parent_id` smallint NULL DEFAULT NULL,
  `no_cache` tinyint(1) NULL DEFAULT NULL,
  `breadcrumb` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `component` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `sort` tinyint NULL DEFAULT NULL,
  `visible` varchar(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `is_frame` varchar(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0',
  `create_by` bigint NULL DEFAULT NULL COMMENT '创建者',
  `update_by` bigint NULL DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`menu_id`) USING BTREE,
  INDEX `idx_sys_menu_deleted_at`(`deleted_at` ASC) USING BTREE,
  INDEX `idx_sys_menu_create_by`(`create_by` ASC) USING BTREE,
  INDEX `idx_sys_menu_update_by`(`update_by` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 659 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO `sys_menu` VALUES (2, 'Admin', '系统管理', 'api-server', '/admin', '/0/2', 'M', '无', '', 0, 1, '', 'Layout', 10, '0', '1', 0, 1, '2021-05-20 21:58:45.679', '2023-07-21 15:39:04.145', NULL);
INSERT INTO `sys_menu` VALUES (3, 'SysUserManage', '用户管理', 'user', '/admin/sys-user', '/0/2/3', 'C', '无', 'admin:sysUser:list', 2, 0, '', '/admin/sys-user/index', 10, '0', '1', 0, 1, '2021-05-20 22:08:44.526', '2023-07-21 15:39:04.149', NULL);
INSERT INTO `sys_menu` VALUES (43, '', '新增管理员', 'app-group-fill', '', '/0/2/3/43', 'F', 'POST', 'admin:sysUser:add', 3, 0, '', '', 10, '0', '1', 0, 1, '2021-05-20 22:08:44.526', '2023-07-21 15:39:04.161', NULL);
INSERT INTO `sys_menu` VALUES (44, '', '查询管理员', 'app-group-fill', '', '/0/2/3/44', 'F', 'GET', 'admin:sysUser:query', 3, 0, '', '', 40, '0', '1', 0, 1, '2021-05-20 22:08:44.526', '2023-07-21 15:39:04.162', NULL);
INSERT INTO `sys_menu` VALUES (45, '', '修改管理员', 'app-group-fill', '', '/0/2/3/45', 'F', 'PUT', 'admin:sysUser:edit', 3, 0, '', '', 30, '0', '1', 0, 1, '2021-05-20 22:08:44.526', '2023-07-21 15:39:04.163', NULL);
INSERT INTO `sys_menu` VALUES (46, '', '删除管理员', 'app-group-fill', '', '/0/2/3/46', 'F', 'DELETE', 'admin:sysUser:remove', 3, 0, '', '', 20, '0', '1', 0, 1, '2021-05-20 22:08:44.526', '2023-07-21 15:39:04.164', NULL);
INSERT INTO `sys_menu` VALUES (51, 'SysMenuManage', '菜单管理', 'tree-table', '/admin/sys-menu', '/0/2/51', 'C', '无', 'admin:sysMenu:list', 2, 1, '', '/admin/sys-menu/index', 30, '0', '1', 0, 1, '2021-05-20 22:08:44.526', '2023-07-21 15:39:04.151', NULL);
INSERT INTO `sys_menu` VALUES (52, 'SysRoleManage', '角色管理', 'peoples', '/admin/sys-role', '/0/2/52', 'C', '无', 'admin:sysRole:list', 2, 1, '', '/admin/sys-role/index', 20, '0', '1', 0, 1, '2021-05-20 22:08:44.526', '2023-07-21 15:39:04.152', NULL);
INSERT INTO `sys_menu` VALUES (56, 'SysDeptManage', '部门管理', 'tree', '/admin/sys-dept', '/0/2/56', 'C', '无', 'admin:sysDept:list', 2, 0, '', '/admin/sys-dept/index', 40, '0', '1', 0, 1, '2021-05-20 22:08:44.526', '2023-07-21 15:39:04.152', NULL);
INSERT INTO `sys_menu` VALUES (57, 'SysPostManage', '岗位管理', 'pass', '/admin/sys-post', '/0/2/57', 'C', '无', 'admin:sysPost:list', 2, 0, '', '/admin/sys-post/index', 50, '0', '1', 0, 1, '2021-05-20 22:08:44.526', '2023-07-21 15:39:04.154', NULL);
INSERT INTO `sys_menu` VALUES (58, 'Dict', '字典管理', 'education', '/admin/dict', '/0/2/58', 'C', '无', 'admin:sysDictType:list', 2, 0, '', '/admin/dict/index', 60, '0', '1', 0, 1, '2021-05-20 22:08:44.526', '2023-07-21 15:39:04.155', NULL);
INSERT INTO `sys_menu` VALUES (59, 'SysDictDataManage', '字典数据', 'education', '/admin/dict/data/:dictId', '/0/2/59', 'C', '无', 'admin:sysDictData:list', 2, 0, '', '/admin/dict/data', 100, '1', '1', 0, 1, '2021-05-20 22:08:44.526', '2023-07-21 15:39:04.156', NULL);
INSERT INTO `sys_menu` VALUES (60, 'Tools', '开发工具', 'dev-tools', '/dev-tools', '/0/60', 'M', '无', '', 0, 0, '', 'Layout', 40, '0', '1', 1, 1, '2020-04-11 15:52:48.000', '2023-07-21 15:39:04.146', NULL);
INSERT INTO `sys_menu` VALUES (61, 'Swagger', '系统接口', 'guide', '/dev-tools/swagger', '/0/60/61', 'C', '无', '', 60, 0, '', '/dev-tools/swagger/index', 1, '0', '1', 1, 1, '2020-04-11 15:52:48.000', '2023-07-21 15:39:04.190', NULL);
INSERT INTO `sys_menu` VALUES (62, 'SysConfigManage', '参数管理', 'swagger', '/admin/sys-config', '/0/2/62', 'C', '无', 'admin:sysConfig:list', 2, 0, '', '/admin/sys-config/index', 70, '0', '1', 0, 1, '2021-05-20 22:08:44.526', '2023-07-21 15:39:04.157', NULL);
INSERT INTO `sys_menu` VALUES (211, 'Log', '日志管理', 'log', '/log', '/0/2/211', 'M', '', '', 2, 0, '', '/log/index', 80, '0', '1', 0, 1, '2021-05-20 22:08:44.526', '2023-07-21 15:39:04.158', NULL);
INSERT INTO `sys_menu` VALUES (212, 'SysLoginLogManage', '登录日志', 'logininfor', '/admin/sys-login-log', '/0/2/211/212', 'C', '', 'admin:sysLoginLog:list', 211, 0, '', '/admin/sys-login-log/index', 1, '0', '1', 0, 1, '2021-05-20 22:08:44.526', '2023-07-21 15:39:04.197', NULL);
INSERT INTO `sys_menu` VALUES (216, 'OperLog', '操作日志', 'skill', '/admin/sys-oper-log', '/0/2/211/216', 'C', '', 'admin:sysOperLog:list', 211, 0, '', '/admin/sys-oper-log/index', 1, '0', '1', 0, 1, '2021-05-20 22:08:44.526', '2023-07-21 15:39:04.197', NULL);
INSERT INTO `sys_menu` VALUES (220, '', '新增菜单', 'app-group-fill', '', '/0/2/51/220', 'F', '', 'admin:sysMenu:add', 51, 0, '', '', 1, '0', '1', 1, 1, '2020-04-11 15:52:48.000', '2023-07-21 15:39:04.165', NULL);
INSERT INTO `sys_menu` VALUES (221, '', '修改菜单', 'app-group-fill', '', '/0/2/51/221', 'F', '', 'admin:sysMenu:edit', 51, 0, '', '', 1, '0', '1', 1, 1, '2020-04-11 15:52:48.000', '2023-07-21 15:39:04.166', NULL);
INSERT INTO `sys_menu` VALUES (222, '', '查询菜单', 'app-group-fill', '', '/0/2/51/222', 'F', '', 'admin:sysMenu:query', 51, 0, '', '', 1, '0', '1', 1, 1, '2020-04-11 15:52:48.000', '2023-07-21 15:39:04.166', NULL);
INSERT INTO `sys_menu` VALUES (223, '', '删除菜单', 'app-group-fill', '', '/0/2/51/223', 'F', '', 'admin:sysMenu:remove', 51, 0, '', '', 1, '0', '1', 1, 1, '2020-04-11 15:52:48.000', '2023-07-21 15:39:04.168', NULL);
INSERT INTO `sys_menu` VALUES (224, '', '新增角色', 'app-group-fill', '', '/0/2/52/224', 'F', '', 'admin:sysRole:add', 52, 0, '', '', 1, '0', '1', 1, 1, '2020-04-11 15:52:48.000', '2023-07-21 15:39:04.169', NULL);
INSERT INTO `sys_menu` VALUES (225, '', '查询角色', 'app-group-fill', '', '/0/2/52/225', 'F', '', 'admin:sysRole:query', 52, 0, '', '', 1, '0', '1', 1, 1, '2020-04-11 15:52:48.000', '2023-07-21 15:39:04.169', NULL);
INSERT INTO `sys_menu` VALUES (226, '', '修改角色', 'app-group-fill', '', '/0/2/52/226', 'F', '', 'admin:sysRole:update', 52, 0, '', '', 1, '0', '1', 1, 1, '2020-04-11 15:52:48.000', '2023-07-21 15:39:04.170', NULL);
INSERT INTO `sys_menu` VALUES (227, '', '删除角色', 'app-group-fill', '', '/0/2/52/227', 'F', '', 'admin:sysRole:remove', 52, 0, '', '', 1, '0', '1', 1, 1, '2020-04-11 15:52:48.000', '2023-07-21 15:39:04.172', NULL);
INSERT INTO `sys_menu` VALUES (228, '', '查询部门', 'app-group-fill', '', '/0/2/56/228', 'F', '', 'admin:sysDept:query', 56, 0, '', '', 40, '0', '1', 0, 1, '2021-05-20 22:08:44.526', '2023-07-21 15:39:04.173', NULL);
INSERT INTO `sys_menu` VALUES (229, '', '新增部门', 'app-group-fill', '', '/0/2/56/229', 'F', '', 'admin:sysDept:add', 56, 0, '', '', 10, '0', '1', 0, 1, '2021-05-20 22:08:44.526', '2023-07-21 15:39:04.174', NULL);
INSERT INTO `sys_menu` VALUES (230, '', '修改部门', 'app-group-fill', '', '/0/2/56/230', 'F', '', 'admin:sysDept:edit', 56, 0, '', '', 30, '0', '1', 0, 1, '2021-05-20 22:08:44.526', '2023-07-21 15:39:04.174', NULL);
INSERT INTO `sys_menu` VALUES (231, '', '删除部门', 'app-group-fill', '', '/0/2/56/231', 'F', '', 'admin:sysDept:remove', 56, 0, '', '', 20, '0', '1', 0, 1, '2021-05-20 22:08:44.526', '2023-07-21 15:39:04.176', NULL);
INSERT INTO `sys_menu` VALUES (232, '', '查询岗位', 'app-group-fill', '', '/0/2/57/232', 'F', '', 'admin:sysPost:query', 57, 0, '', '', 0, '0', '1', 1, 1, '2020-04-11 15:52:48.000', '2023-07-21 15:39:04.177', NULL);
INSERT INTO `sys_menu` VALUES (233, '', '新增岗位', 'app-group-fill', '', '/0/2/57/233', 'F', '', 'admin:sysPost:add', 57, 0, '', '', 0, '0', '1', 1, 1, '2020-04-11 15:52:48.000', '2023-07-21 15:39:04.178', NULL);
INSERT INTO `sys_menu` VALUES (234, '', '修改岗位', 'app-group-fill', '', '/0/2/57/234', 'F', '', 'admin:sysPost:edit', 57, 0, '', '', 0, '0', '1', 1, 1, '2020-04-11 15:52:48.000', '2023-07-21 15:39:04.180', NULL);
INSERT INTO `sys_menu` VALUES (235, '', '删除岗位', 'app-group-fill', '', '/0/2/57/235', 'F', '', 'admin:sysPost:remove', 57, 0, '', '', 0, '0', '1', 1, 1, '2020-04-11 15:52:48.000', '2023-07-21 15:39:04.181', NULL);
INSERT INTO `sys_menu` VALUES (236, '', '查询字典', 'app-group-fill', '', '/0/2/58/236', 'F', '', 'admin:sysDictType:query', 58, 0, '', '', 0, '0', '1', 1, 1, '2020-04-11 15:52:48.000', '2023-07-21 15:39:04.181', NULL);
INSERT INTO `sys_menu` VALUES (237, '', '新增类型', 'app-group-fill', '', '/0/2/58/237', 'F', '', 'admin:sysDictType:add', 58, 0, '', '', 0, '0', '1', 1, 1, '2020-04-11 15:52:48.000', '2023-07-21 15:39:04.182', NULL);
INSERT INTO `sys_menu` VALUES (238, '', '修改类型', 'app-group-fill', '', '/0/2/58/238', 'F', '', 'admin:sysDictType:edit', 58, 0, '', '', 0, '0', '1', 1, 1, '2020-04-11 15:52:48.000', '2023-07-21 15:39:04.183', NULL);
INSERT INTO `sys_menu` VALUES (239, '', '删除类型', 'app-group-fill', '', '/0/2/58/239', 'F', '', 'system:sysdicttype:remove', 58, 0, '', '', 0, '0', '1', 1, 1, '2020-04-11 15:52:48.000', '2023-07-21 15:39:04.184', NULL);
INSERT INTO `sys_menu` VALUES (240, '', '查询数据', 'app-group-fill', '', '/0/2/59/240', 'F', '', 'admin:sysDictData:query', 59, 0, '', '', 0, '0', '1', 1, 1, '2020-04-11 15:52:48.000', '2023-07-21 15:39:04.185', NULL);
INSERT INTO `sys_menu` VALUES (241, '', '新增数据', 'app-group-fill', '', '/0/2/59/241', 'F', '', 'admin:sysDictData:add', 59, 0, '', '', 0, '0', '1', 1, 1, '2020-04-11 15:52:48.000', '2023-07-21 15:39:04.186', NULL);
INSERT INTO `sys_menu` VALUES (242, '', '修改数据', 'app-group-fill', '', '/0/2/59/242', 'F', '', 'admin:sysDictData:edit', 59, 0, '', '', 0, '0', '1', 1, 1, '2020-04-11 15:52:48.000', '2023-07-21 15:39:04.187', NULL);
INSERT INTO `sys_menu` VALUES (243, '', '删除数据', 'app-group-fill', '', '/0/2/59/243', 'F', '', 'admin:sysDictData:remove', 59, 0, '', '', 0, '0', '1', 1, 1, '2020-04-11 15:52:48.000', '2023-07-21 15:39:04.188', NULL);
INSERT INTO `sys_menu` VALUES (244, '', '查询参数', 'app-group-fill', '', '/0/2/62/244', 'F', '', 'admin:sysConfig:query', 62, 0, '', '', 0, '0', '1', 1, 1, '2020-04-11 15:52:48.000', '2023-07-21 15:39:04.193', NULL);
INSERT INTO `sys_menu` VALUES (245, '', '新增参数', 'app-group-fill', '', '/0/2/62/245', 'F', '', 'admin:sysConfig:add', 62, 0, '', '', 0, '0', '1', 1, 1, '2020-04-11 15:52:48.000', '2023-07-21 15:39:04.194', NULL);
INSERT INTO `sys_menu` VALUES (246, '', '修改参数', 'app-group-fill', '', '/0/2/62/246', 'F', '', 'admin:sysConfig:edit', 62, 0, '', '', 0, '0', '1', 1, 1, '2020-04-11 15:52:48.000', '2023-07-21 15:39:04.195', NULL);
INSERT INTO `sys_menu` VALUES (247, '', '删除参数', 'app-group-fill', '', '/0/2/62/247', 'F', '', 'admin:sysConfig:remove', 62, 0, '', '', 0, '0', '1', 1, 1, '2020-04-11 15:52:48.000', '2023-07-21 15:39:04.195', NULL);
INSERT INTO `sys_menu` VALUES (248, '', '查询登录日志', 'app-group-fill', '', '/0/2/211/212/248', 'F', '', 'admin:sysLoginLog:query', 212, 0, '', '', 0, '0', '1', 1, 1, '2020-04-11 15:52:48.000', '2023-07-21 15:39:04.198', NULL);
INSERT INTO `sys_menu` VALUES (249, '', '删除登录日志', 'app-group-fill', '', '/0/2/211/212/249', 'F', '', 'admin:sysLoginLog:remove', 212, 0, '', '', 0, '0', '1', 1, 1, '2020-04-11 15:52:48.000', '2023-07-21 15:39:04.199', NULL);
INSERT INTO `sys_menu` VALUES (250, '', '查询操作日志', 'app-group-fill', '', '/0/2/211/216/250', 'F', '', 'admin:sysOperLog:query', 216, 0, '', '', 0, '0', '1', 1, 1, '2020-04-11 15:52:48.000', '2023-07-21 15:39:04.200', NULL);
INSERT INTO `sys_menu` VALUES (251, '', '删除操作日志', 'app-group-fill', '', '/0/2/211/216/251', 'F', '', 'admin:sysOperLog:remove', 216, 0, '', '', 0, '0', '1', 1, 1, '2020-04-11 15:52:48.000', '2023-07-21 15:39:04.201', NULL);
INSERT INTO `sys_menu` VALUES (261, 'Gen', '代码生成', 'code', '/dev-tools/gen', '/0/60/261', 'C', '', '', 60, 0, '', '/dev-tools/gen/index', 2, '0', '1', 1, 1, '2020-04-11 15:52:48.000', '2023-07-21 15:39:04.190', NULL);
INSERT INTO `sys_menu` VALUES (262, 'EditTable', '代码生成修改', 'build', '/dev-tools/editTable', '/0/60/262', 'C', '', '', 60, 0, '', '/dev-tools/gen/editTable', 100, '1', '1', 1, 1, '2020-04-11 15:52:48.000', '2023-07-21 15:39:04.191', NULL);
INSERT INTO `sys_menu` VALUES (264, 'Build', '表单构建', 'build', '/dev-tools/build', '/0/60/264', 'C', '', '', 60, 0, '', '/dev-tools/build/index', 1, '0', '1', 1, 1, '2020-04-11 15:52:48.000', '2023-07-21 15:39:04.192', NULL);
INSERT INTO `sys_menu` VALUES (269, 'ServerMonitor', '服务监控', 'druid', '/sys-tools/monitor', '/0/537/269', 'C', '', 'sysTools:serverMonitor:list', 537, 0, '', '/sys-tools/monitor', 0, '0', '1', 1, 1, '2020-04-14 00:28:19.000', '2023-07-21 15:39:04.210', NULL);
INSERT INTO `sys_menu` VALUES (459, 'Schedule', '定时任务', 'time-range', '/schedule', '/0/459', 'M', '无', '', 0, 0, '', 'Layout', 20, '0', '1', 1, 1, '2020-08-03 09:17:37.000', '2023-07-21 15:39:04.147', NULL);
INSERT INTO `sys_menu` VALUES (460, 'ScheduleManage', 'Schedule', 'job', '/schedule/manage', '/0/459/460', 'C', '无', 'job:sysJob:list', 459, 0, '', '/schedule/index', 0, '0', '1', 1, 1, '2020-08-03 09:17:37.000', '2023-07-21 15:39:04.202', NULL);
INSERT INTO `sys_menu` VALUES (461, 'sys_job', '分页获取定时任务', 'app-group-fill', '', '/0/459/460/461', 'F', '无', 'job:sysJob:query', 460, 0, '', '', 0, '0', '1', 1, 1, '2020-08-03 09:17:37.000', '2023-07-21 15:39:04.204', NULL);
INSERT INTO `sys_menu` VALUES (462, 'sys_job', '创建定时任务', 'app-group-fill', '', '/0/459/460/462', 'F', '无', 'job:sysJob:add', 460, 0, '', '', 0, '0', '1', 1, 1, '2020-08-03 09:17:37.000', '2023-07-21 15:39:04.205', NULL);
INSERT INTO `sys_menu` VALUES (463, 'sys_job', '修改定时任务', 'app-group-fill', '', '/0/459/460/463', 'F', '无', 'job:sysJob:edit', 460, 0, '', '', 0, '0', '1', 1, 1, '2020-08-03 09:17:37.000', '2023-07-21 15:39:04.206', NULL);
INSERT INTO `sys_menu` VALUES (464, 'sys_job', '删除定时任务', 'app-group-fill', '', '/0/459/460/464', 'F', '无', 'job:sysJob:remove', 460, 0, '', '', 0, '0', '1', 1, 1, '2020-08-03 09:17:37.000', '2023-07-21 15:39:04.207', NULL);
INSERT INTO `sys_menu` VALUES (471, 'JobLog', '日志', 'bug', '/schedule/log', '/0/459/471', 'C', '', '', 459, 0, '', '/schedule/log', 0, '1', '1', 1, 1, '2020-08-05 21:24:46.000', '2023-07-21 15:39:04.203', NULL);
INSERT INTO `sys_menu` VALUES (528, 'SysApiManage', '接口管理', 'api-doc', '/admin/sys-api', '/0/2/528', 'C', '无', 'admin:sysApi:list', 2, 0, '', '/admin/sys-api/index', 0, '0', '0', 0, 1, '2021-05-20 22:08:44.526', '2023-07-21 15:39:04.159', NULL);
INSERT INTO `sys_menu` VALUES (529, '', '查询接口', 'app-group-fill', '', '/0/2/528/529', 'F', '无', 'admin:sysApi:query', 528, 0, '', '', 40, '0', '0', 0, 1, '2021-05-20 22:08:44.526', '2023-07-21 15:39:04.209', NULL);
INSERT INTO `sys_menu` VALUES (531, '', '修改接口', 'app-group-fill', '', '/0/2/528/531', 'F', '无', 'admin:sysApi:edit', 528, 0, '', '', 30, '0', '0', 0, 1, '2021-05-20 22:08:44.526', '2023-07-21 15:39:04.210', NULL);
INSERT INTO `sys_menu` VALUES (537, 'SysTools', '系统工具', 'system-tools', '/sys-tools', '/0/537', 'M', '', '', 0, 0, '', 'Layout', 30, '0', '1', 1, 1, '2021-05-21 11:13:32.166', '2023-07-21 15:39:04.148', NULL);
INSERT INTO `sys_menu` VALUES (540, 'SysConfigSet', '参数设置', 'system-tools', '/admin/sys-config/set', '/0/2/540', 'C', '', 'admin:sysConfigSet:list', 2, 0, '', '/admin/sys-config/set', 0, '0', '1', 1, 1, '2021-05-25 16:06:52.560', '2023-07-21 15:39:04.160', NULL);
INSERT INTO `sys_menu` VALUES (542, '', '修改', 'upload', '', '/0/2/540/542', 'F', '', 'admin:sysConfigSet:update', 540, 0, '', '', 0, '0', '1', 1, 1, '2021-06-13 11:45:48.670', '2023-07-21 15:39:04.211', NULL);
INSERT INTO `sys_menu` VALUES (543, '', '用户', 'peoples', '/user', '/0/543', 'M', '无', '', 0, 0, '', 'Layout', 1, '0', '1', 1, 1, '2023-07-21 17:12:35.422', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (544, 'AssetManage', '资产列表', 'network', '/user/asset', '/0/543/544', 'C', '无', 'user:asset:list', 543, 0, '', '/user/asset/index', 2, '0', '1', 1, 1, '2023-07-21 17:12:35.426', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (545, '', '分页获取Asset', '', 'asset', '/0/543/544/545', 'F', '无', 'user:asset:query', 544, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 17:12:35.431', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (546, '', '创建Asset', '', 'asset', '/0/543/544/546', 'F', '无', 'user:asset:add', 544, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 17:12:35.434', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (547, '', '修改Asset', '', 'asset', '/0/543/544/547', 'F', '无', 'user:asset:edit', 544, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 17:12:35.438', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (548, '', '删除Asset', '', 'asset', '/0/543/544/548', 'F', '无', 'user:asset:remove', 544, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 17:12:35.444', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (549, 'UserManage', '用户列表', 'project-group', '/user/user', '/0/543/549', 'C', '无', 'user:user:list', 543, 0, '', '/user/user/index', 1, '0', '1', 1, 1, '2023-07-21 17:12:35.426', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (550, '', '分页获取User', '', 'user', '/0/543/549/550', 'F', '无', 'user:user:query', 549, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 17:12:35.431', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (551, '', '创建User', '', 'user', '/0/543/549/551', 'F', '无', 'user:user:add', 549, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 17:12:35.434', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (552, '', '修改User', '', 'user', '/0/543/549/552', 'F', '无', 'user:user:edit', 549, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 17:12:35.438', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (553, '', '删除User', '', 'user', '/0/543/549/553', 'F', '无', 'user:user:remove', 549, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 17:12:35.444', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (554, 'RechargeManage', '充值记录', 'archived', '/user/recharge', '/0/543/554', 'C', '无', 'user:recharge:list', 543, 0, '', '/user/recharge/index', 3, '0', '1', 1, 1, '2023-07-21 17:12:35.426', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (555, '', '分页获取Recharge', '', 'recharge', '/0/543/554/555', 'F', '无', 'user:recharge:query', 554, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 17:12:35.431', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (556, '', '创建Recharge', '', 'recharge', '/0/543/554/556', 'F', '无', 'user:recharge:add', 554, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 17:12:35.434', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (557, '', '修改Recharge', '', 'recharge', '/0/543/554/557', 'F', '无', 'user:recharge:edit', 554, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 17:12:35.438', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (558, '', '删除Recharge', '', 'recharge', '/0/543/554/558', 'F', '无', 'user:recharge:remove', 554, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 17:12:35.444', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (559, 'CheckinManage', '签到记录', 'calendar', '/user/checkin', '/0/543/559', 'C', '无', 'user:checkin:list', 543, 0, '', '/user/checkin/index', 4, '0', '1', 1, 1, '2023-07-21 17:12:35.426', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (560, '', '分页获取Checkin', '', 'checkin', '/0/543/559/560', 'F', '无', 'user:checkin:query', 559, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 17:12:35.431', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (561, '', '创建Checkin', '', 'checkin', '/0/543/559/561', 'F', '无', 'user:checkin:add', 559, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 17:12:35.434', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (562, '', '修改Checkin', '', 'checkin', '/0/543/559/562', 'F', '无', 'user:checkin:edit', 559, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 17:12:35.438', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (563, '', '删除Checkin', '', 'checkin', '/0/543/559/563', 'F', '无', 'user:checkin:remove', 559, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 17:12:35.444', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (564, '', '平台', 'example', '/platform', '/0/564', 'M', '无', '', 0, 0, '', 'Layout', 2, '0', '1', 1, 1, '2023-07-21 18:44:03.021', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (565, 'PlatformManage', '平台列表', 'list', '/platform/platform', '/0/564/565', 'C', '无', 'platform:platform:list', 564, 0, '', '/platform/platform/index', 1, '0', '1', 1, 1, '2023-07-21 18:44:03.039', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (566, '', '分页获取平台列表', '', 'platform', '/0/564/565/566', 'F', '无', 'platform:platform:query', 565, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 18:44:03.048', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (567, '', '创建平台列表', '', 'platform', '/0/564/565/567', 'F', '无', 'platform:platform:add', 565, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 18:44:03.053', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (568, '', '修改平台列表', '', 'platform', '/0/564/565/568', 'F', '无', 'platform:platform:edit', 565, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 18:44:03.065', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (569, '', '删除平台列表', '', 'platform', '/0/564/565/569', 'F', '无', 'platform:platform:remove', 565, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 18:44:03.070', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (576, '', '活动列表', 'pass', '/activity', '/0/576', 'M', '无', '', 0, 0, '', 'Layout', 0, '0', '0', 1, 0, '2023-07-21 18:56:34.514', '2023-07-21 18:56:34.515', '2023-07-21 18:58:09.363');
INSERT INTO `sys_menu` VALUES (577, 'ActivityManage', '活动列表', 'alarm-settings', '/platform/activity', '/0/576/577', 'C', '无', 'platform:activity:list', 564, 0, '', '/platform/activity/index', 4, '0', '1', 1, 1, '2023-07-21 18:56:34.517', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (578, '', '分页获取活动列表', '', 'activity', '/0/576/577/578', 'F', '无', 'platform:activity:query', 577, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 18:56:34.520', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (579, '', '创建活动列表', '', 'activity', '/0/576/577/579', 'F', '无', 'platform:activity:add', 577, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 18:56:34.523', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (580, '', '修改活动列表', '', 'activity', '/0/576/577/580', 'F', '无', 'platform:activity:edit', 577, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 18:56:34.528', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (581, '', '删除活动列表', '', 'activity', '/0/576/577/581', 'F', '无', 'platform:activity:remove', 577, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 18:56:34.532', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (582, '', '游戏', 'dashboard', '/categories', '/0/582', 'M', '无', '', 0, 0, '', 'Layout', 3, '0', '1', 1, 1, '2023-07-21 19:27:37.188', '2023-07-21 19:28:11.509', '2023-07-21 20:32:59.299');
INSERT INTO `sys_menu` VALUES (588, '', '房间列表', 'pass', '/room', '/0/588', 'M', '无', '', 0, 0, '', 'Layout', 0, '0', '0', 1, 0, '2023-07-21 19:31:03.419', '2023-07-21 19:31:03.421', '2023-07-21 19:33:36.350');
INSERT INTO `sys_menu` VALUES (589, 'RoomManage', '房间列表', 'number', '/platform/room', '/0/588/589', 'C', '无', 'platform:room:list', 564, 0, '', '/platform/room/index', 2, '0', '1', 1, 1, '2023-07-21 19:31:03.434', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (590, '', '分页获取房间列表', '', 'room', '/0/588/589/590', 'F', '无', 'platform:room:query', 589, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 19:31:03.440', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (591, '', '创建房间列表', '', 'room', '/0/588/589/591', 'F', '无', 'platform:room:add', 589, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 19:31:03.444', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (592, '', '修改房间列表', '', 'room', '/0/588/589/592', 'F', '无', 'platform:room:edit', 589, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 19:31:03.446', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (593, '', '删除房间列表', '', 'room', '/0/588/589/593', 'F', '无', 'platform:room:remove', 589, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 19:31:03.449', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (606, '', '游戏', 'dashboard', '/categories', '/0/606', 'M', '无', '', 0, 0, '', 'Layout', 3, '0', '1', 1, 1, '2023-07-21 20:32:28.864', '2023-07-21 20:33:15.615', '2023-07-21 20:40:19.642');
INSERT INTO `sys_menu` VALUES (607, 'CategoriesManage', '游戏类别', 'root-addr', '/game/categories', '/0/606/607', 'C', '无', 'game:categories:list', 618, 0, '', '/game/categories/index', 3, '0', '1', 1, 1, '2023-07-21 20:32:28.880', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (608, '', '分页获取游戏类别', '', 'categories', '/0/606/607/608', 'F', '无', 'game:categories:query', 607, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 20:32:28.887', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (609, '', '创建游戏类别', '', 'categories', '/0/606/607/609', 'F', '无', 'game:categories:add', 607, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 20:32:28.892', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (610, '', '修改游戏类别', '', 'categories', '/0/606/607/610', 'F', '无', 'game:categories:edit', 607, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 20:32:28.897', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (611, '', '删除游戏类别', '', 'categories', '/0/606/607/611', 'F', '无', 'game:categories:remove', 607, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 20:32:28.902', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (618, '', '游戏', 'table', '/game', '/0/618', 'M', '无', '', 0, 0, '', 'Layout', 3, '0', '1', 1, 1, '2023-07-21 20:39:18.449', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (619, 'GameManage', '游戏列表', 'build', '/game/game', '/0/618/619', 'C', '无', 'game:game:list', 618, 0, '', '/game/game/index', 1, '0', '1', 1, 1, '2023-07-21 20:39:18.463', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (620, '', '分页获取Game', '', 'game', '/0/618/619/620', 'F', '无', 'game:game:query', 619, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 20:39:18.468', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (621, '', '创建Game', '', 'game', '/0/618/619/621', 'F', '无', 'game:game:add', 619, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 20:39:18.478', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (622, '', '修改Game', '', 'game', '/0/618/619/622', 'F', '无', 'game:game:edit', 619, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 20:39:18.484', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (623, '', '删除Game', '', 'game', '/0/618/619/623', 'F', '无', 'game:game:remove', 619, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 20:39:18.491', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (624, '', '游戏种类', 'pass', '/kindinfo', '/0/624', 'M', '无', '', 0, 0, '', 'Layout', 0, '0', '0', 1, 0, '2023-07-21 20:44:27.847', '2023-07-21 20:44:27.848', '2023-07-21 20:45:44.776');
INSERT INTO `sys_menu` VALUES (625, 'KindinfoManage', '游戏种类', 'http-err', '/game/kindinfo', '/0/624/625', 'C', '无', 'game:kindinfo:list', 618, 0, '', '/game/kindinfo/index', 4, '0', '1', 1, 1, '2023-07-21 20:44:27.862', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (626, '', '分页获取游戏种类', '', 'kindinfo', '/0/624/625/626', 'F', '无', 'game:kindinfo:query', 625, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 20:44:27.868', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (627, '', '创建游戏种类', '', 'kindinfo', '/0/624/625/627', 'F', '无', 'game:kindinfo:add', 625, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 20:44:27.878', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (628, '', '修改游戏种类', '', 'kindinfo', '/0/624/625/628', 'F', '无', 'game:kindinfo:edit', 625, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 20:44:27.883', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (629, '', '删除游戏种类', '', 'kindinfo', '/0/624/625/629', 'F', '无', 'game:kindinfo:remove', 625, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 20:44:27.888', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (630, '', '游戏记录', 'pass', '/record', '/0/630', 'M', '无', '', 0, 0, '', 'Layout', 0, '0', '0', 1, 0, '2023-07-21 22:29:56.135', '2023-07-21 22:29:56.137', '2023-07-21 22:31:11.591');
INSERT INTO `sys_menu` VALUES (631, 'RecordManage', '游戏记录', 'date-range', '/game/record', '/0/630/631', 'C', '无', 'game:record:list', 618, 0, '', '/game/record/index', 2, '0', '1', 1, 1, '2023-07-21 22:29:56.141', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (632, '', '分页获取游戏记录', '', 'record', '/0/630/631/632', 'F', '无', 'game:record:query', 631, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 22:29:56.147', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (633, '', '创建游戏记录', '', 'record', '/0/630/631/633', 'F', '无', 'game:record:add', 631, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 22:29:56.152', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (634, '', '修改游戏记录', '', 'record', '/0/630/631/634', 'F', '无', 'game:record:edit', 631, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 22:29:56.156', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (635, '', '删除游戏记录', '', 'record', '/0/630/631/635', 'F', '无', 'game:record:remove', 631, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 22:29:56.161', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (636, '', '商品列表', 'pass', '/goods', '/0/636', 'M', '无', '', 0, 0, '', 'Layout', 0, '0', '0', 1, 0, '2023-07-21 23:39:51.224', '2023-07-21 23:39:51.226', '2023-07-21 23:41:43.417');
INSERT INTO `sys_menu` VALUES (637, 'GoodsManage', '商品列表', 'shopping', '/platform/goods', '/0/636/637', 'C', '无', 'platform:goods:list', 564, 0, '', '/platform/goods/index', 3, '0', '1', 1, 1, '2023-07-21 23:39:51.239', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (638, '', '分页获取商品列表', '', 'goods', '/0/636/637/638', 'F', '无', 'platform:goods:query', 637, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 23:39:51.244', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (639, '', '创建商品列表', '', 'goods', '/0/636/637/639', 'F', '无', 'platform:goods:add', 637, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 23:39:51.253', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (640, '', '修改商品列表', '', 'goods', '/0/636/637/640', 'F', '无', 'platform:goods:edit', 637, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 23:39:51.258', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (641, '', '删除商品列表', '', 'goods', '/0/636/637/641', 'F', '无', 'platform:goods:remove', 637, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 23:39:51.263', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (642, '', '邮箱列表', 'pass', '/email', '/0/642', 'M', '无', '', 0, 0, '', 'Layout', 0, '0', '0', 1, 0, '2023-07-21 23:48:13.045', '2023-07-21 23:48:13.057', '2023-07-21 23:49:12.459');
INSERT INTO `sys_menu` VALUES (643, 'EmailManage', '邮箱列表', 'checkbox', '/user/email', '/0/642/643', 'C', '无', 'user:email:list', 543, 0, '', '/user/email/index', 5, '0', '1', 1, 1, '2023-07-21 23:48:13.061', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (644, '', '分页获取邮箱列表', '', 'email', '/0/642/643/644', 'F', '无', 'user:email:query', 643, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 23:48:13.066', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (645, '', '创建邮箱列表', '', 'email', '/0/642/643/645', 'F', '无', 'user:email:add', 643, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 23:48:13.069', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (646, '', '修改邮箱列表', '', 'email', '/0/642/643/646', 'F', '无', 'user:email:edit', 643, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 23:48:13.072', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (647, '', '删除邮箱列表', '', 'email', '/0/642/643/647', 'F', '无', 'user:email:remove', 643, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 23:48:13.077', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (648, '', '公告消息', 'pass', '/notice', '/0/648', 'M', '无', '', 0, 0, '', 'Layout', 0, '0', '0', 1, 0, '2023-07-21 23:55:37.537', '2023-07-21 23:55:37.538', '2023-07-21 23:56:15.871');
INSERT INTO `sys_menu` VALUES (649, 'NoticeManage', '公告消息', 'guide', '/platform/notice', '/0/648/649', 'C', '无', 'platform:notice:list', 564, 0, '', '/platform/notice/index', 5, '0', '1', 1, 1, '2023-07-21 23:55:37.542', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (650, '', '分页获取公告消息', '', 'notice', '/0/648/649/650', 'F', '无', 'platform:notice:query', 649, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 23:55:37.549', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (651, '', '创建公告消息', '', 'notice', '/0/648/649/651', 'F', '无', 'platform:notice:add', 649, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 23:55:37.554', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (652, '', '修改公告消息', '', 'notice', '/0/648/649/652', 'F', '无', 'platform:notice:edit', 649, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 23:55:37.558', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (653, '', '删除公告消息', '', 'notice', '/0/648/649/653', 'F', '无', 'platform:notice:remove', 649, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 23:55:37.565', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (654, '', '社交群组', 'pass', '/chat-group', '/0/654', 'M', '无', '', 0, 0, '', 'Layout', 0, '0', '0', 1, 0, '2023-07-21 23:58:39.173', '2023-07-21 23:58:39.174', '2023-07-22 00:04:49.079');
INSERT INTO `sys_menu` VALUES (655, 'ChatGroupManage', '社交群组', 'wechat', '/user/chat-group', '/0/654/655', 'C', '无', 'user:chatGroup:list', 543, 0, '', '/user/chat-group/index', 6, '0', '1', 1, 1, '2023-07-21 23:58:39.204', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (656, '', '分页获取社交群组', '', 'chat_group', '/0/654/655/656', 'F', '无', 'user:chatGroup:query', 655, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 23:58:39.210', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (657, '', '创建社交群组', '', 'chat_group', '/0/654/655/657', 'F', '无', 'user:chatGroup:add', 655, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 23:58:39.217', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (658, '', '修改社交群组', '', 'chat_group', '/0/654/655/658', 'F', '无', 'user:chatGroup:edit', 655, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 23:58:39.222', '2023-07-23 17:14:39.690', NULL);
INSERT INTO `sys_menu` VALUES (659, '', '删除社交群组', '', 'chat_group', '/0/654/655/659', 'F', '无', 'user:chatGroup:remove', 655, 0, '', '', 0, '0', '0', 1, 1, '2023-07-21 23:58:39.225', '2023-07-23 17:14:39.690', NULL);

-- ----------------------------
-- Table structure for sys_menu_api_rule
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu_api_rule`;
CREATE TABLE `sys_menu_api_rule`  (
  `sys_menu_menu_id` bigint NOT NULL,
  `sys_api_id` bigint NOT NULL COMMENT '主键编码',
  PRIMARY KEY (`sys_menu_menu_id`, `sys_api_id`) USING BTREE,
  INDEX `fk_sys_menu_api_rule_sys_api`(`sys_api_id` ASC) USING BTREE,
  CONSTRAINT `fk_sys_menu_api_rule_sys_api` FOREIGN KEY (`sys_api_id`) REFERENCES `sys_api` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_sys_menu_api_rule_sys_menu` FOREIGN KEY (`sys_menu_menu_id`) REFERENCES `sys_menu` (`menu_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_menu_api_rule
-- ----------------------------
INSERT INTO `sys_menu_api_rule` VALUES (550, 5);
INSERT INTO `sys_menu_api_rule` VALUES (216, 6);
INSERT INTO `sys_menu_api_rule` VALUES (250, 6);
INSERT INTO `sys_menu_api_rule` VALUES (550, 6);
INSERT INTO `sys_menu_api_rule` VALUES (550, 7);
INSERT INTO `sys_menu_api_rule` VALUES (550, 8);
INSERT INTO `sys_menu_api_rule` VALUES (550, 9);
INSERT INTO `sys_menu_api_rule` VALUES (550, 10);
INSERT INTO `sys_menu_api_rule` VALUES (550, 11);
INSERT INTO `sys_menu_api_rule` VALUES (550, 15);
INSERT INTO `sys_menu_api_rule` VALUES (550, 16);
INSERT INTO `sys_menu_api_rule` VALUES (58, 21);
INSERT INTO `sys_menu_api_rule` VALUES (236, 21);
INSERT INTO `sys_menu_api_rule` VALUES (550, 21);
INSERT INTO `sys_menu_api_rule` VALUES (238, 23);
INSERT INTO `sys_menu_api_rule` VALUES (550, 23);
INSERT INTO `sys_menu_api_rule` VALUES (59, 24);
INSERT INTO `sys_menu_api_rule` VALUES (240, 24);
INSERT INTO `sys_menu_api_rule` VALUES (550, 24);
INSERT INTO `sys_menu_api_rule` VALUES (242, 25);
INSERT INTO `sys_menu_api_rule` VALUES (550, 25);
INSERT INTO `sys_menu_api_rule` VALUES (58, 26);
INSERT INTO `sys_menu_api_rule` VALUES (236, 26);
INSERT INTO `sys_menu_api_rule` VALUES (56, 27);
INSERT INTO `sys_menu_api_rule` VALUES (228, 27);
INSERT INTO `sys_menu_api_rule` VALUES (550, 27);
INSERT INTO `sys_menu_api_rule` VALUES (230, 28);
INSERT INTO `sys_menu_api_rule` VALUES (550, 28);
INSERT INTO `sys_menu_api_rule` VALUES (226, 29);
INSERT INTO `sys_menu_api_rule` VALUES (51, 39);
INSERT INTO `sys_menu_api_rule` VALUES (222, 39);
INSERT INTO `sys_menu_api_rule` VALUES (550, 39);
INSERT INTO `sys_menu_api_rule` VALUES (221, 41);
INSERT INTO `sys_menu_api_rule` VALUES (550, 41);
INSERT INTO `sys_menu_api_rule` VALUES (52, 44);
INSERT INTO `sys_menu_api_rule` VALUES (225, 44);
INSERT INTO `sys_menu_api_rule` VALUES (550, 44);
INSERT INTO `sys_menu_api_rule` VALUES (226, 45);
INSERT INTO `sys_menu_api_rule` VALUES (226, 46);
INSERT INTO `sys_menu_api_rule` VALUES (226, 47);
INSERT INTO `sys_menu_api_rule` VALUES (550, 47);
INSERT INTO `sys_menu_api_rule` VALUES (62, 53);
INSERT INTO `sys_menu_api_rule` VALUES (244, 53);
INSERT INTO `sys_menu_api_rule` VALUES (550, 53);
INSERT INTO `sys_menu_api_rule` VALUES (246, 54);
INSERT INTO `sys_menu_api_rule` VALUES (550, 54);
INSERT INTO `sys_menu_api_rule` VALUES (550, 57);
INSERT INTO `sys_menu_api_rule` VALUES (550, 58);
INSERT INTO `sys_menu_api_rule` VALUES (57, 59);
INSERT INTO `sys_menu_api_rule` VALUES (232, 59);
INSERT INTO `sys_menu_api_rule` VALUES (550, 59);
INSERT INTO `sys_menu_api_rule` VALUES (234, 60);
INSERT INTO `sys_menu_api_rule` VALUES (550, 60);
INSERT INTO `sys_menu_api_rule` VALUES (550, 66);
INSERT INTO `sys_menu_api_rule` VALUES (550, 67);
INSERT INTO `sys_menu_api_rule` VALUES (550, 72);
INSERT INTO `sys_menu_api_rule` VALUES (550, 73);
INSERT INTO `sys_menu_api_rule` VALUES (550, 76);
INSERT INTO `sys_menu_api_rule` VALUES (241, 80);
INSERT INTO `sys_menu_api_rule` VALUES (550, 80);
INSERT INTO `sys_menu_api_rule` VALUES (237, 81);
INSERT INTO `sys_menu_api_rule` VALUES (550, 81);
INSERT INTO `sys_menu_api_rule` VALUES (229, 82);
INSERT INTO `sys_menu_api_rule` VALUES (550, 82);
INSERT INTO `sys_menu_api_rule` VALUES (245, 87);
INSERT INTO `sys_menu_api_rule` VALUES (550, 87);
INSERT INTO `sys_menu_api_rule` VALUES (220, 88);
INSERT INTO `sys_menu_api_rule` VALUES (550, 88);
INSERT INTO `sys_menu_api_rule` VALUES (233, 89);
INSERT INTO `sys_menu_api_rule` VALUES (550, 89);
INSERT INTO `sys_menu_api_rule` VALUES (224, 90);
INSERT INTO `sys_menu_api_rule` VALUES (550, 90);
INSERT INTO `sys_menu_api_rule` VALUES (531, 92);
INSERT INTO `sys_menu_api_rule` VALUES (550, 92);
INSERT INTO `sys_menu_api_rule` VALUES (550, 95);
INSERT INTO `sys_menu_api_rule` VALUES (550, 96);
INSERT INTO `sys_menu_api_rule` VALUES (550, 97);
INSERT INTO `sys_menu_api_rule` VALUES (242, 101);
INSERT INTO `sys_menu_api_rule` VALUES (550, 101);
INSERT INTO `sys_menu_api_rule` VALUES (238, 102);
INSERT INTO `sys_menu_api_rule` VALUES (550, 102);
INSERT INTO `sys_menu_api_rule` VALUES (230, 103);
INSERT INTO `sys_menu_api_rule` VALUES (550, 103);
INSERT INTO `sys_menu_api_rule` VALUES (550, 104);
INSERT INTO `sys_menu_api_rule` VALUES (550, 105);
INSERT INTO `sys_menu_api_rule` VALUES (226, 106);
INSERT INTO `sys_menu_api_rule` VALUES (550, 106);
INSERT INTO `sys_menu_api_rule` VALUES (226, 107);
INSERT INTO `sys_menu_api_rule` VALUES (550, 107);
INSERT INTO `sys_menu_api_rule` VALUES (246, 108);
INSERT INTO `sys_menu_api_rule` VALUES (550, 108);
INSERT INTO `sys_menu_api_rule` VALUES (221, 109);
INSERT INTO `sys_menu_api_rule` VALUES (550, 109);
INSERT INTO `sys_menu_api_rule` VALUES (234, 110);
INSERT INTO `sys_menu_api_rule` VALUES (550, 110);
INSERT INTO `sys_menu_api_rule` VALUES (550, 112);
INSERT INTO `sys_menu_api_rule` VALUES (550, 113);
INSERT INTO `sys_menu_api_rule` VALUES (249, 114);
INSERT INTO `sys_menu_api_rule` VALUES (550, 114);
INSERT INTO `sys_menu_api_rule` VALUES (251, 115);
INSERT INTO `sys_menu_api_rule` VALUES (550, 115);
INSERT INTO `sys_menu_api_rule` VALUES (550, 116);
INSERT INTO `sys_menu_api_rule` VALUES (550, 117);
INSERT INTO `sys_menu_api_rule` VALUES (243, 120);
INSERT INTO `sys_menu_api_rule` VALUES (550, 120);
INSERT INTO `sys_menu_api_rule` VALUES (239, 121);
INSERT INTO `sys_menu_api_rule` VALUES (550, 121);
INSERT INTO `sys_menu_api_rule` VALUES (231, 122);
INSERT INTO `sys_menu_api_rule` VALUES (550, 122);
INSERT INTO `sys_menu_api_rule` VALUES (550, 123);
INSERT INTO `sys_menu_api_rule` VALUES (550, 124);
INSERT INTO `sys_menu_api_rule` VALUES (247, 125);
INSERT INTO `sys_menu_api_rule` VALUES (550, 125);
INSERT INTO `sys_menu_api_rule` VALUES (223, 126);
INSERT INTO `sys_menu_api_rule` VALUES (550, 126);
INSERT INTO `sys_menu_api_rule` VALUES (235, 127);
INSERT INTO `sys_menu_api_rule` VALUES (550, 127);
INSERT INTO `sys_menu_api_rule` VALUES (227, 128);
INSERT INTO `sys_menu_api_rule` VALUES (550, 128);
INSERT INTO `sys_menu_api_rule` VALUES (51, 135);
INSERT INTO `sys_menu_api_rule` VALUES (528, 135);
INSERT INTO `sys_menu_api_rule` VALUES (529, 135);
INSERT INTO `sys_menu_api_rule` VALUES (550, 135);
INSERT INTO `sys_menu_api_rule` VALUES (531, 136);
INSERT INTO `sys_menu_api_rule` VALUES (550, 136);
INSERT INTO `sys_menu_api_rule` VALUES (212, 137);
INSERT INTO `sys_menu_api_rule` VALUES (248, 137);
INSERT INTO `sys_menu_api_rule` VALUES (550, 137);
INSERT INTO `sys_menu_api_rule` VALUES (542, 139);
INSERT INTO `sys_menu_api_rule` VALUES (550, 139);
INSERT INTO `sys_menu_api_rule` VALUES (540, 140);
INSERT INTO `sys_menu_api_rule` VALUES (550, 140);
INSERT INTO `sys_menu_api_rule` VALUES (3, 141);
INSERT INTO `sys_menu_api_rule` VALUES (44, 141);
INSERT INTO `sys_menu_api_rule` VALUES (550, 141);
INSERT INTO `sys_menu_api_rule` VALUES (45, 142);
INSERT INTO `sys_menu_api_rule` VALUES (550, 142);
INSERT INTO `sys_menu_api_rule` VALUES (43, 150);
INSERT INTO `sys_menu_api_rule` VALUES (550, 150);
INSERT INTO `sys_menu_api_rule` VALUES (45, 151);
INSERT INTO `sys_menu_api_rule` VALUES (550, 151);
INSERT INTO `sys_menu_api_rule` VALUES (46, 156);
INSERT INTO `sys_menu_api_rule` VALUES (550, 156);

-- ----------------------------
-- Table structure for sys_migration
-- ----------------------------
DROP TABLE IF EXISTS `sys_migration`;
CREATE TABLE `sys_migration`  (
  `version` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `apply_time` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`version`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_migration
-- ----------------------------
INSERT INTO `sys_migration` VALUES ('1599190683659', '2023-07-21 15:39:04.140');
INSERT INTO `sys_migration` VALUES ('1653638869132', '2023-07-21 15:39:04.212');

-- ----------------------------
-- Table structure for sys_opera_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_opera_log`;
CREATE TABLE `sys_opera_log`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '操作模块',
  `business_type` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '操作类型',
  `business_types` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'BusinessTypes',
  `method` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '函数',
  `request_method` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '请求方式: GET POST PUT DELETE',
  `operator_type` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '操作类型',
  `oper_name` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '操作者',
  `dept_name` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '部门名称',
  `oper_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '访问地址',
  `oper_ip` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '客户端ip',
  `oper_location` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '访问位置',
  `oper_param` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '请求参数',
  `status` varchar(4) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '操作状态 1:正常 2:关闭',
  `oper_time` timestamp NULL DEFAULT NULL COMMENT '操作时间',
  `json_result` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '返回数据',
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  `latency_time` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '耗时',
  `user_agent` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'ua',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '最后更新时间',
  `create_by` bigint NULL DEFAULT NULL COMMENT '创建者',
  `update_by` bigint NULL DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_opera_log_create_by`(`create_by` ASC) USING BTREE,
  INDEX `idx_sys_opera_log_update_by`(`update_by` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_opera_log
-- ----------------------------

-- ----------------------------
-- Table structure for sys_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post`  (
  `post_id` bigint NOT NULL AUTO_INCREMENT,
  `post_name` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `post_code` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `sort` tinyint NULL DEFAULT NULL,
  `status` tinyint NULL DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `create_by` bigint NULL DEFAULT NULL COMMENT '创建者',
  `update_by` bigint NULL DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`post_id`) USING BTREE,
  INDEX `idx_sys_post_create_by`(`create_by` ASC) USING BTREE,
  INDEX `idx_sys_post_update_by`(`update_by` ASC) USING BTREE,
  INDEX `idx_sys_post_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_post
-- ----------------------------
INSERT INTO `sys_post` VALUES (1, '首席执行官', 'CEO', 2, 2, '首席执行官', 1, 1, '2021-05-13 19:56:37.913', '2023-07-21 21:50:37.337', NULL);
INSERT INTO `sys_post` VALUES (2, '首席技术执行官', 'CTO', 2, 2, '首席技术执行官', 1, 1, '2021-05-13 19:56:37.913', '2021-05-13 19:56:37.913', NULL);
INSERT INTO `sys_post` VALUES (3, '首席运营官', 'COO', 3, 2, '测试工程师', 1, 1, '2021-05-13 19:56:37.913', '2021-05-13 19:56:37.913', NULL);
INSERT INTO `sys_post` VALUES (4, '普通员工', 'Employee', 4, 2, '', 1, 1, '2023-07-23 16:20:55.448', '2023-07-23 16:21:20.185', NULL);

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `role_id` bigint NOT NULL AUTO_INCREMENT,
  `role_name` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `status` varchar(4) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `role_key` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `role_sort` bigint NULL DEFAULT NULL,
  `flag` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `admin` tinyint(1) NULL DEFAULT NULL,
  `data_scope` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `create_by` bigint NULL DEFAULT NULL COMMENT '创建者',
  `update_by` bigint NULL DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`role_id`) USING BTREE,
  INDEX `idx_sys_role_deleted_at`(`deleted_at` ASC) USING BTREE,
  INDEX `idx_sys_role_create_by`(`create_by` ASC) USING BTREE,
  INDEX `idx_sys_role_update_by`(`update_by` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES (1, '系统管理员', '2', 'admin', 1, '', '', 1, '', 1, 1, '2021-05-13 19:56:37.913', '2021-05-13 19:56:37.913', NULL);
INSERT INTO `sys_role` VALUES (3, '平台管理员', '2', 'platform', 2, '', '专门对代理商进行开放', 0, '1', 0, 0, '2023-07-23 16:13:19.480', '2023-07-23 17:14:39.689', NULL);
INSERT INTO `sys_role` VALUES (4, '工作人员', '2', 'person', 3, '', '', 0, '5', 0, 0, '2023-07-23 16:16:09.619', '2023-07-23 16:16:19.141', NULL);
INSERT INTO `sys_role` VALUES (5, '用户', '2', 'user', 1, '', '', 0, '1', 0, 0, '2023-07-23 17:15:20.166', '2023-07-23 17:15:27.291', NULL);

-- ----------------------------
-- Table structure for sys_role_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_dept`;
CREATE TABLE `sys_role_dept`  (
  `role_id` smallint NOT NULL,
  `dept_id` smallint NOT NULL,
  PRIMARY KEY (`role_id`, `dept_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role_dept
-- ----------------------------

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu`  (
  `role_id` bigint NOT NULL,
  `menu_id` bigint NOT NULL,
  PRIMARY KEY (`role_id`, `menu_id`) USING BTREE,
  INDEX `fk_sys_role_menu_sys_menu`(`menu_id` ASC) USING BTREE,
  CONSTRAINT `fk_sys_role_menu_sys_menu` FOREIGN KEY (`menu_id`) REFERENCES `sys_menu` (`menu_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_sys_role_menu_sys_role` FOREIGN KEY (`role_id`) REFERENCES `sys_role` (`role_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
INSERT INTO `sys_role_menu` VALUES (4, 269);
INSERT INTO `sys_role_menu` VALUES (4, 459);
INSERT INTO `sys_role_menu` VALUES (4, 460);
INSERT INTO `sys_role_menu` VALUES (4, 461);
INSERT INTO `sys_role_menu` VALUES (4, 462);
INSERT INTO `sys_role_menu` VALUES (4, 463);
INSERT INTO `sys_role_menu` VALUES (4, 464);
INSERT INTO `sys_role_menu` VALUES (4, 471);
INSERT INTO `sys_role_menu` VALUES (4, 537);
INSERT INTO `sys_role_menu` VALUES (3, 543);
INSERT INTO `sys_role_menu` VALUES (4, 543);
INSERT INTO `sys_role_menu` VALUES (3, 544);
INSERT INTO `sys_role_menu` VALUES (4, 544);
INSERT INTO `sys_role_menu` VALUES (3, 545);
INSERT INTO `sys_role_menu` VALUES (4, 545);
INSERT INTO `sys_role_menu` VALUES (3, 546);
INSERT INTO `sys_role_menu` VALUES (4, 546);
INSERT INTO `sys_role_menu` VALUES (3, 547);
INSERT INTO `sys_role_menu` VALUES (4, 547);
INSERT INTO `sys_role_menu` VALUES (3, 548);
INSERT INTO `sys_role_menu` VALUES (4, 548);
INSERT INTO `sys_role_menu` VALUES (3, 549);
INSERT INTO `sys_role_menu` VALUES (4, 549);
INSERT INTO `sys_role_menu` VALUES (5, 549);
INSERT INTO `sys_role_menu` VALUES (3, 550);
INSERT INTO `sys_role_menu` VALUES (4, 550);
INSERT INTO `sys_role_menu` VALUES (5, 550);
INSERT INTO `sys_role_menu` VALUES (3, 551);
INSERT INTO `sys_role_menu` VALUES (4, 551);
INSERT INTO `sys_role_menu` VALUES (5, 551);
INSERT INTO `sys_role_menu` VALUES (3, 552);
INSERT INTO `sys_role_menu` VALUES (4, 552);
INSERT INTO `sys_role_menu` VALUES (5, 552);
INSERT INTO `sys_role_menu` VALUES (3, 553);
INSERT INTO `sys_role_menu` VALUES (4, 553);
INSERT INTO `sys_role_menu` VALUES (5, 553);
INSERT INTO `sys_role_menu` VALUES (3, 554);
INSERT INTO `sys_role_menu` VALUES (4, 554);
INSERT INTO `sys_role_menu` VALUES (3, 555);
INSERT INTO `sys_role_menu` VALUES (4, 555);
INSERT INTO `sys_role_menu` VALUES (3, 556);
INSERT INTO `sys_role_menu` VALUES (4, 556);
INSERT INTO `sys_role_menu` VALUES (3, 557);
INSERT INTO `sys_role_menu` VALUES (4, 557);
INSERT INTO `sys_role_menu` VALUES (3, 558);
INSERT INTO `sys_role_menu` VALUES (4, 558);
INSERT INTO `sys_role_menu` VALUES (3, 559);
INSERT INTO `sys_role_menu` VALUES (4, 559);
INSERT INTO `sys_role_menu` VALUES (3, 560);
INSERT INTO `sys_role_menu` VALUES (4, 560);
INSERT INTO `sys_role_menu` VALUES (3, 561);
INSERT INTO `sys_role_menu` VALUES (4, 561);
INSERT INTO `sys_role_menu` VALUES (3, 562);
INSERT INTO `sys_role_menu` VALUES (4, 562);
INSERT INTO `sys_role_menu` VALUES (3, 563);
INSERT INTO `sys_role_menu` VALUES (4, 563);
INSERT INTO `sys_role_menu` VALUES (3, 564);
INSERT INTO `sys_role_menu` VALUES (4, 564);
INSERT INTO `sys_role_menu` VALUES (3, 565);
INSERT INTO `sys_role_menu` VALUES (4, 565);
INSERT INTO `sys_role_menu` VALUES (3, 566);
INSERT INTO `sys_role_menu` VALUES (4, 566);
INSERT INTO `sys_role_menu` VALUES (3, 567);
INSERT INTO `sys_role_menu` VALUES (4, 567);
INSERT INTO `sys_role_menu` VALUES (3, 568);
INSERT INTO `sys_role_menu` VALUES (4, 568);
INSERT INTO `sys_role_menu` VALUES (3, 569);
INSERT INTO `sys_role_menu` VALUES (4, 569);
INSERT INTO `sys_role_menu` VALUES (3, 577);
INSERT INTO `sys_role_menu` VALUES (4, 577);
INSERT INTO `sys_role_menu` VALUES (3, 578);
INSERT INTO `sys_role_menu` VALUES (4, 578);
INSERT INTO `sys_role_menu` VALUES (3, 579);
INSERT INTO `sys_role_menu` VALUES (4, 579);
INSERT INTO `sys_role_menu` VALUES (3, 580);
INSERT INTO `sys_role_menu` VALUES (4, 580);
INSERT INTO `sys_role_menu` VALUES (3, 581);
INSERT INTO `sys_role_menu` VALUES (4, 581);
INSERT INTO `sys_role_menu` VALUES (3, 589);
INSERT INTO `sys_role_menu` VALUES (4, 589);
INSERT INTO `sys_role_menu` VALUES (3, 590);
INSERT INTO `sys_role_menu` VALUES (4, 590);
INSERT INTO `sys_role_menu` VALUES (3, 591);
INSERT INTO `sys_role_menu` VALUES (4, 591);
INSERT INTO `sys_role_menu` VALUES (3, 592);
INSERT INTO `sys_role_menu` VALUES (4, 592);
INSERT INTO `sys_role_menu` VALUES (3, 593);
INSERT INTO `sys_role_menu` VALUES (4, 593);
INSERT INTO `sys_role_menu` VALUES (3, 607);
INSERT INTO `sys_role_menu` VALUES (4, 607);
INSERT INTO `sys_role_menu` VALUES (3, 608);
INSERT INTO `sys_role_menu` VALUES (4, 608);
INSERT INTO `sys_role_menu` VALUES (3, 609);
INSERT INTO `sys_role_menu` VALUES (4, 609);
INSERT INTO `sys_role_menu` VALUES (3, 610);
INSERT INTO `sys_role_menu` VALUES (4, 610);
INSERT INTO `sys_role_menu` VALUES (3, 611);
INSERT INTO `sys_role_menu` VALUES (4, 611);
INSERT INTO `sys_role_menu` VALUES (3, 618);
INSERT INTO `sys_role_menu` VALUES (4, 618);
INSERT INTO `sys_role_menu` VALUES (3, 619);
INSERT INTO `sys_role_menu` VALUES (4, 619);
INSERT INTO `sys_role_menu` VALUES (3, 620);
INSERT INTO `sys_role_menu` VALUES (4, 620);
INSERT INTO `sys_role_menu` VALUES (3, 621);
INSERT INTO `sys_role_menu` VALUES (4, 621);
INSERT INTO `sys_role_menu` VALUES (3, 622);
INSERT INTO `sys_role_menu` VALUES (4, 622);
INSERT INTO `sys_role_menu` VALUES (3, 623);
INSERT INTO `sys_role_menu` VALUES (4, 623);
INSERT INTO `sys_role_menu` VALUES (3, 625);
INSERT INTO `sys_role_menu` VALUES (4, 625);
INSERT INTO `sys_role_menu` VALUES (3, 626);
INSERT INTO `sys_role_menu` VALUES (4, 626);
INSERT INTO `sys_role_menu` VALUES (3, 627);
INSERT INTO `sys_role_menu` VALUES (4, 627);
INSERT INTO `sys_role_menu` VALUES (3, 628);
INSERT INTO `sys_role_menu` VALUES (4, 628);
INSERT INTO `sys_role_menu` VALUES (3, 629);
INSERT INTO `sys_role_menu` VALUES (4, 629);
INSERT INTO `sys_role_menu` VALUES (3, 631);
INSERT INTO `sys_role_menu` VALUES (4, 631);
INSERT INTO `sys_role_menu` VALUES (3, 632);
INSERT INTO `sys_role_menu` VALUES (4, 632);
INSERT INTO `sys_role_menu` VALUES (3, 633);
INSERT INTO `sys_role_menu` VALUES (4, 633);
INSERT INTO `sys_role_menu` VALUES (3, 634);
INSERT INTO `sys_role_menu` VALUES (4, 634);
INSERT INTO `sys_role_menu` VALUES (3, 635);
INSERT INTO `sys_role_menu` VALUES (4, 635);
INSERT INTO `sys_role_menu` VALUES (3, 637);
INSERT INTO `sys_role_menu` VALUES (4, 637);
INSERT INTO `sys_role_menu` VALUES (3, 638);
INSERT INTO `sys_role_menu` VALUES (4, 638);
INSERT INTO `sys_role_menu` VALUES (3, 639);
INSERT INTO `sys_role_menu` VALUES (4, 639);
INSERT INTO `sys_role_menu` VALUES (3, 640);
INSERT INTO `sys_role_menu` VALUES (4, 640);
INSERT INTO `sys_role_menu` VALUES (3, 641);
INSERT INTO `sys_role_menu` VALUES (4, 641);
INSERT INTO `sys_role_menu` VALUES (3, 643);
INSERT INTO `sys_role_menu` VALUES (4, 643);
INSERT INTO `sys_role_menu` VALUES (3, 644);
INSERT INTO `sys_role_menu` VALUES (4, 644);
INSERT INTO `sys_role_menu` VALUES (3, 645);
INSERT INTO `sys_role_menu` VALUES (4, 645);
INSERT INTO `sys_role_menu` VALUES (3, 646);
INSERT INTO `sys_role_menu` VALUES (4, 646);
INSERT INTO `sys_role_menu` VALUES (3, 647);
INSERT INTO `sys_role_menu` VALUES (4, 647);
INSERT INTO `sys_role_menu` VALUES (3, 649);
INSERT INTO `sys_role_menu` VALUES (4, 649);
INSERT INTO `sys_role_menu` VALUES (3, 650);
INSERT INTO `sys_role_menu` VALUES (4, 650);
INSERT INTO `sys_role_menu` VALUES (3, 651);
INSERT INTO `sys_role_menu` VALUES (4, 651);
INSERT INTO `sys_role_menu` VALUES (3, 652);
INSERT INTO `sys_role_menu` VALUES (4, 652);
INSERT INTO `sys_role_menu` VALUES (3, 653);
INSERT INTO `sys_role_menu` VALUES (4, 653);
INSERT INTO `sys_role_menu` VALUES (3, 655);
INSERT INTO `sys_role_menu` VALUES (4, 655);
INSERT INTO `sys_role_menu` VALUES (3, 656);
INSERT INTO `sys_role_menu` VALUES (4, 656);
INSERT INTO `sys_role_menu` VALUES (3, 657);
INSERT INTO `sys_role_menu` VALUES (4, 657);
INSERT INTO `sys_role_menu` VALUES (3, 658);
INSERT INTO `sys_role_menu` VALUES (4, 658);
INSERT INTO `sys_role_menu` VALUES (3, 659);
INSERT INTO `sys_role_menu` VALUES (4, 659);

-- ----------------------------
-- Table structure for sys_tables
-- ----------------------------
DROP TABLE IF EXISTS `sys_tables`;
CREATE TABLE `sys_tables`  (
  `table_id` bigint NOT NULL AUTO_INCREMENT,
  `table_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `table_comment` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `class_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `tpl_category` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `package_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `module_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `module_front_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '前端文件名',
  `business_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `function_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `function_author` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `pk_column` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `pk_go_field` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `pk_json_field` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `options` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `tree_code` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `tree_parent_code` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `tree_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `tree` tinyint(1) NULL DEFAULT 0,
  `crud` tinyint(1) NULL DEFAULT 1,
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `is_data_scope` tinyint NULL DEFAULT NULL,
  `is_actions` tinyint NULL DEFAULT NULL,
  `is_auth` tinyint NULL DEFAULT NULL,
  `is_logical_delete` varchar(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `logical_delete` tinyint(1) NULL DEFAULT NULL,
  `logical_delete_column` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `create_by` bigint NULL DEFAULT NULL COMMENT '创建者',
  `update_by` bigint NULL DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`table_id`) USING BTREE,
  INDEX `idx_sys_tables_update_by`(`update_by` ASC) USING BTREE,
  INDEX `idx_sys_tables_deleted_at`(`deleted_at` ASC) USING BTREE,
  INDEX `idx_sys_tables_create_by`(`create_by` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_tables
-- ----------------------------
INSERT INTO `sys_tables` VALUES (1, 'asset', '资产列表', 'Asset', 'crud', 'user', 'asset', '', 'asset', '资产列表信息', 'xiaoshuai', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-07-21 16:08:30.054', '2023-07-21 16:16:11.815', NULL, 0, 0);
INSERT INTO `sys_tables` VALUES (2, 'activity', '活动列表', 'Activity', 'crud', 'platform', 'activity', '', 'activity', '活动详情', 'xiaoshuai', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-07-21 16:08:30.101', '2023-07-21 19:13:18.200', NULL, 0, 0);
INSERT INTO `sys_tables` VALUES (3, 'categories', '游戏类别', 'Categories', 'crud', 'game', 'categories', '', 'categories', '游戏类别信息', 'xiaoshuai', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-07-21 16:08:30.139', '2023-07-21 20:27:20.777', NULL, 0, 0);
INSERT INTO `sys_tables` VALUES (4, 'chat_group', '社交群组', 'ChatGroup', 'crud', 'user', 'chat-group', '', 'chatGroup', 'ChatGroup', 'xiaoshuai', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-07-21 16:08:30.172', '2023-07-21 23:58:33.506', NULL, 0, 0);
INSERT INTO `sys_tables` VALUES (5, 'checkin', '签到记录', 'Checkin', 'crud', 'user', 'checkin', '', 'checkin', 'Checkin', 'xiaoshuai', 'uid', 'Uid', 'uid', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-07-21 16:08:30.222', '2023-07-21 18:28:24.784', NULL, 0, 0);
INSERT INTO `sys_tables` VALUES (6, 'email', '邮箱列表', 'Email', 'crud', 'user', 'email', '', 'email', '邮箱列表', 'xiaoshuai', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-07-21 16:08:30.247', '2023-07-21 23:48:04.584', NULL, 0, 0);
INSERT INTO `sys_tables` VALUES (7, 'game', '游戏', 'Game', 'crud', 'game', 'game', '', 'game', '游戏', 'xiaoshuai', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-07-21 16:08:30.284', '2023-07-21 20:43:13.738', NULL, 0, 0);
INSERT INTO `sys_tables` VALUES (8, 'goods', '商品列表', 'Goods', 'crud', 'platform', 'goods', '', 'goods', '商品列表信息', 'xiaoshuai', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-07-21 16:08:30.345', '2023-07-21 23:39:44.134', NULL, 0, 0);
INSERT INTO `sys_tables` VALUES (9, 'kindinfo', '游戏种类', 'Kindinfo', 'crud', 'game', 'kindinfo', '', 'kindinfo', '游戏种类列表', 'xiaoshuai', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-07-21 16:08:30.379', '2023-07-21 20:44:20.705', NULL, 0, 0);
INSERT INTO `sys_tables` VALUES (10, 'notice', '公告消息', 'Notice', 'crud', 'platform', 'notice', '', 'notice', '公告通知', 'xiaoshuai', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-07-21 16:08:45.527', '2023-07-21 23:55:30.218', NULL, 0, 0);
INSERT INTO `sys_tables` VALUES (11, 'platform', '平台列表', 'Platform', 'crud', 'platform', 'platform', '', 'platform', '平台信息', 'xiaoshuai', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-07-21 16:08:45.580', '2023-07-21 18:43:50.690', NULL, 0, 0);
INSERT INTO `sys_tables` VALUES (12, 'recharge', '充值列表', 'Recharge', 'crud', 'user', 'recharge', '', 'recharge', '充值记录', 'xiaoshuai', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-07-21 16:08:45.613', '2023-07-21 18:35:16.687', NULL, 0, 0);
INSERT INTO `sys_tables` VALUES (13, 'record', '游戏记录', 'Record', 'crud', 'game', 'record', '', 'record', '游戏记录', 'xiaoshuai', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-07-21 16:08:45.654', '2023-07-21 22:29:47.446', NULL, 0, 0);
INSERT INTO `sys_tables` VALUES (14, 'room', '房间列表', 'Room', 'crud', 'platform', 'room', '', 'room', '房间列表', 'xiaoshuai', 'num', 'Num', 'num', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-07-21 16:08:45.704', '2023-07-21 19:30:57.144', NULL, 0, 0);
INSERT INTO `sys_tables` VALUES (15, 'user', '用户列表', 'User', 'crud', 'user', 'user', '', 'user', '用户列表', 'xiaoshuai', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-07-21 16:09:02.172', '2023-07-21 17:31:18.618', NULL, 0, 0);

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `user_id` bigint NOT NULL AUTO_INCREMENT COMMENT '编码',
  `username` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户名',
  `password` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '密码',
  `nick_name` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '昵称',
  `phone` varchar(11) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '手机号',
  `role_id` bigint NULL DEFAULT NULL COMMENT '角色ID',
  `salt` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '加盐',
  `avatar` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '头像',
  `sex` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '性别',
  `email` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  `dept_id` bigint NULL DEFAULT NULL COMMENT '部门',
  `post_id` bigint NULL DEFAULT NULL COMMENT '岗位',
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  `status` varchar(4) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '状态',
  `create_by` bigint NULL DEFAULT NULL COMMENT '创建者',
  `update_by` bigint NULL DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`user_id`) USING BTREE,
  INDEX `idx_sys_user_create_by`(`create_by` ASC) USING BTREE,
  INDEX `idx_sys_user_update_by`(`update_by` ASC) USING BTREE,
  INDEX `idx_sys_user_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1, 'admin', '$2a$10$/Glr4g9Svr6O0kvjsRJCXu3f0W8/dsP3XZyVNi1019ratWpSPMyw.', 'zhangwj', '13818888888', 1, '', '', '1', '1@qq.com', 1, 1, '', '2', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:40.205', NULL);
INSERT INTO `sys_user` VALUES (2, '哒哒', '$2a$10$y2m2I7bnid1suwfuxFkSiuCFe8J8KuCxVV.Ym7luzbmbQI7JM1Saq', 'killy', '13500022201', 3, '', '', '0', '2437854120@qq.com', 11, 1, '', '2', 1, 0, '2023-07-23 16:23:04.394', '2023-07-23 16:23:04.394', NULL);
INSERT INTO `sys_user` VALUES (3, 'xiaoxiao', '$2a$10$4ekzLwQnvj/MSztASOHz4uuMt2xl1sT8Zwes6R4bKxGlWrslwg2Qq', 'hiki', '13555000044', 4, '', '', '1', 'pitter19890809@outlook.com', 11, 4, '', '2', 1, 0, '2023-07-23 16:23:50.013', '2023-07-23 16:48:48.809', NULL);
INSERT INTO `sys_user` VALUES (4, 'sa', '$2a$10$6YVJlgY9WNrEG6mIPmIWde.Ecm7lbJAQ9EPPirlhTAVLHo4Ur/jK2', 'hejil', '14555000444', 5, '', '', '1', '2437854120@qq.com', 12, 4, '', '2', 1, 0, '2023-07-23 16:24:34.459', '2023-07-23 17:15:43.121', NULL);

-- ----------------------------
-- Table structure for table
-- ----------------------------
DROP TABLE IF EXISTS `table`;
CREATE TABLE `table`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '桌子ID',
  `name` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '名称',
  `rid` bigint NOT NULL DEFAULT 0 COMMENT '房间ID',
  `gid` bigint NOT NULL DEFAULT 0 COMMENT '游戏ID',
  `opentime` bigint NULL DEFAULT 0 COMMENT '开桌时间(时间戳)',
  `maxround` int NULL DEFAULT NULL COMMENT '游戏最大轮次(=-1不受限)',
  `commission` int NULL DEFAULT 0 COMMENT '税收(仅针对赢者收取,千分之一)',
  `remain` int NULL DEFAULT 0 COMMENT '剩余场次(=-1不受限)',
  `playscore` bigint NULL DEFAULT 0 COMMENT '初始积分(携带的积分=-1不受限)',
  `max_sitter` int NULL DEFAULT -1 COMMENT '客人(即:可容纳玩家数量 =-1时,不受限)',
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `update_by` bigint NULL DEFAULT 0,
  `create_by` bigint NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 48 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of table
-- ----------------------------
INSERT INTO `table` VALUES (21, '测试三国消', 17, 3, 0, -1, 5, -1, 200, 2, '', '2024-07-20 17:20:44', '2024-07-20 17:20:44', NULL, 0, 0);
INSERT INTO `table` VALUES (22, '测试象棋', 17, 1, 0, -1, 5, -1, 200, 2, '', '2024-07-22 12:43:55', '2024-07-22 12:43:55', NULL, 0, 0);
INSERT INTO `table` VALUES (24, '测试国际象棋', 17, 2, 0, -1, 5, -1, 200, 2, '', '2024-07-27 12:42:26', '2024-07-27 12:42:26', NULL, 0, 0);
INSERT INTO `table` VALUES (25, '测试国际象棋322', 17, 2, 0, -1, 5, -1, 200, 2, '', '2024-07-27 12:52:39', '2024-07-27 12:52:39', NULL, 0, 0);
INSERT INTO `table` VALUES (33, '测试三国消324', 17, 3, 0, -1, 5, -1, 200, 2, '', '2024-07-28 15:12:51', '2024-07-28 15:12:51', NULL, 0, 0);
INSERT INTO `table` VALUES (42, '测试百佳乐', 17, 23, 0, 100, 5, 95, -1, -1, '', '2024-08-05 22:53:40', '2024-08-05 22:53:40', NULL, 0, 0);
INSERT INTO `table` VALUES (43, '试试龙虎', 17, 24, 10000, -1, 20, -1, 200, -1, '', '2024-08-08 18:37:33', '2024-08-08 18:37:33', NULL, 0, 0);
INSERT INTO `table` VALUES (44, '试试牛牛', 17, 24, 10000, -1, 20, -1, 200, -1, '', '2024-08-08 18:38:20', '2024-08-08 18:38:20', NULL, 0, 0);
INSERT INTO `table` VALUES (45, '试试骰宝', 17, 26, 10000, -1, 5, -1, 200, -1, '', '2024-08-08 18:39:04', '2024-08-08 18:39:04', NULL, 0, 0);
INSERT INTO `table` VALUES (46, '试试推童子', 17, 27, 10000, -1, 5, -1, 200, -1, '', '2024-08-08 18:45:41', '2024-08-08 18:45:41', NULL, 0, 0);
INSERT INTO `table` VALUES (47, '试试推童子', 17, 25, 10000, -1, 5, -1, 200, -1, '', '2024-08-09 10:24:58', '2024-08-09 10:24:58', NULL, 0, 0);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID身份',
  `name` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '姓名',
  `account` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '账号',
  `password` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '密码',
  `head` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '头像',
  `face` int NULL DEFAULT 0 COMMENT '头像ID',
  `signature` varchar(128) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '签名',
  `gender` int NULL DEFAULT 0 COMMENT '性别',
  `age` int NULL DEFAULT 0 COMMENT '年龄',
  `empiric` int NULL DEFAULT 0 COMMENT '经验值',
  `vip` int NULL DEFAULT 0 COMMENT 'VIP级别',
  `level` tinyint NULL DEFAULT 0 COMMENT '级别',
  `yuanbao` bigint NULL DEFAULT 0 COMMENT '元宝(游戏中不使用money,1:100转成元宝,1:100转成金币)',
  `coin` bigint NULL DEFAULT 0 COMMENT '铜钱(金币)',
  `money` bigint NULL DEFAULT NULL COMMENT '余额(游戏中禁止直接使用)',
  `passport` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '身份ID',
  `realname` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '真实名字',
  `phone` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '手机',
  `address` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '住址',
  `email` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '邮箱',
  `identity` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '识别码',
  `agentid` bigint NULL DEFAULT 0 COMMENT '代理标识(代理人ID)',
  `referralcode` varchar(8) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '推荐标识(由邀请码生成)',
  `serveraddr` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '服务器地址(由平台指定)',
  `clientaddr` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '连接地址',
  `machinecode` varchar(128) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '机器序列',
  `signintime` bigint NULL DEFAULT 0 COMMENT '注册时间',
  `logintime` bigint NULL DEFAULT 0 COMMENT '登陆时间',
  `leavetime` bigint NULL DEFAULT 0 COMMENT '离开时间',
  `remark` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '备注',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `update_by` bigint NULL DEFAULT 0,
  `create_by` bigint NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uniqueid`(`id` ASC) USING BTREE,
  INDEX `normalpsw`(`password` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 51 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (35, 'xiaoming', 'xiaoming', '8d70d8ab2768f232ebe874175065ead3', '', 0, '', 0, 0, 0, 0, 0, 0, 10000, 100, '', '', '', '127.0.0.1', '', '44bad3ad-be8d-4703-a465-2c6c33cbc88c', 0, '', '', '', '', 1720866644, 1725726436, 0, '', '2024-07-13 18:30:44', '2024-07-13 18:30:44', NULL, 0, 0);
INSERT INTO `user` VALUES (36, 'your_username', 'your_username', '466994507d172def0dce9ab607700131', '', 0, '', 0, 0, 0, 0, 0, 0, 10000, 100, '', '', '', '127.0.0.1', '', 'd3742489-ac74-4b53-9ab4-460312422404', 0, '', '', '', '', 1720960474, 0, 0, '', '2024-07-14 20:34:35', '2024-07-14 20:34:35', NULL, 0, 0);
INSERT INTO `user` VALUES (43, 'name000', 'name000', 'cf1f30408ae1fd689ddb8fbd3a3688f4', '', 0, '', 0, 0, 0, 0, 0, 0, 10000, 100, '', '', '', '127.0.0.1', '', 'ce2268b2-3033-4cce-97fb-bac61631b8a5', 0, '', '', '', '', 1722076980, 0, 0, '', '2024-07-27 18:43:01', '2024-07-27 18:43:01', NULL, 0, 0);
INSERT INTO `user` VALUES (44, 'name002', 'name002', 'cf1f30408ae1fd689ddb8fbd3a3688f4', '', 0, '', 0, 0, 0, 0, 0, 0, -1435, 100, '', '', '', '127.0.0.1', '', '95066c3d-7299-47b4-93a0-36653fd6339d', 0, '', '', '', '', 1722076991, 1726808907, 1726808948, '', '2024-07-27 18:43:11', '2024-07-27 18:43:11', NULL, 0, 0);
INSERT INTO `user` VALUES (45, 'name003', 'name003', 'cf1f30408ae1fd689ddb8fbd3a3688f4', '', 0, '', 0, 0, 0, 0, 0, 0, 10000, 100, '', '', '', '127.0.0.1', '', '3b522c79-a4b8-4ed7-9ad2-f9edb4f85cd3', 0, '', '', '', '', 1722076991, 0, 0, '', '2024-07-27 18:43:11', '2024-07-27 18:43:11', NULL, 0, 0);
INSERT INTO `user` VALUES (46, 'name001', 'name001', 'cf1f30408ae1fd689ddb8fbd3a3688f4', '', 0, '', 0, 0, 0, 0, 0, 0, 10000, 100, '', '', '', '127.0.0.1', '', '6ab69d10-4abd-4170-8a45-80546df3ffa9', 0, '', '', '', '', 1722091909, 0, 0, '', '2024-07-27 22:51:50', '2024-07-27 22:51:50', NULL, 0, 0);
INSERT INTO `user` VALUES (47, 'name004', 'name004', 'cf1f30408ae1fd689ddb8fbd3a3688f4', '', 0, '', 0, 0, 0, 0, 0, 0, 10000, 100, '', '', '', '127.0.0.1', '', '329e6ce9-3d63-4d0e-8d1b-f1ded1e713c9', 0, '', '', '', '', 1722433376, 0, 0, '', '2024-07-31 21:42:57', '2024-07-31 21:42:57', NULL, 0, 0);
INSERT INTO `user` VALUES (48, 'name005', 'name005', 'cf1f30408ae1fd689ddb8fbd3a3688f4', '', 0, '', 0, 0, 0, 0, 0, 0, 10000, 100, '', '', '', '127.0.0.1', '', 'cd1f1307-0aa6-4f24-be51-758e768ff5ae', 0, '', '', '', '', 1722433377, 0, 0, '', '2024-07-31 21:42:57', '2024-07-31 21:42:57', NULL, 0, 0);
INSERT INTO `user` VALUES (49, 'name006', 'name006', 'cf1f30408ae1fd689ddb8fbd3a3688f4', '', 0, '', 0, 0, 0, 0, 0, 0, 10000, 100, '', '', '', '127.0.0.1', '', '0e71b24a-3066-45ef-9a05-d4a641049e99', 0, '', '', '', '', 1722433377, 0, 0, '', '2024-07-31 21:42:57', '2024-07-31 21:42:57', NULL, 0, 0);
INSERT INTO `user` VALUES (50, 'xiaomingp', 'xiaomingp', '8d70d8ab2768f232ebe874175065ead3', '', 0, '', 0, 0, 0, 0, 0, 0, 0, 0, '', '', '', '127.0.0.1', '', 'a5025782-7a40-47f5-bd23-41b958eab172', 0, '', '', '', '', 1725724557, 0, 0, '', '2024-09-07 23:55:58', '2024-09-07 23:55:58', NULL, 0, 0);

-- ----------------------------
-- Table structure for user_bind
-- ----------------------------
DROP TABLE IF EXISTS `user_bind`;
CREATE TABLE `user_bind`  (
  `uid` bigint NOT NULL COMMENT '用户唯一id',
  `sdk_id` int NULL DEFAULT 0 COMMENT 'sdk配置id',
  `pid` int NULL DEFAULT NULL COMMENT '平台id',
  `open_id` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '平台帐号open_id',
  `bind_time` bigint NULL DEFAULT NULL COMMENT '绑定时间',
  `up_time` timestamp NOT NULL COMMENT '最后一次更新时间',
  PRIMARY KEY (`uid`) USING BTREE,
  UNIQUE INDEX `pid_open_id_key`(`pid` ASC, `open_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_bind
-- ----------------------------

-- ----------------------------
-- Table structure for weapon
-- ----------------------------
DROP TABLE IF EXISTS `weapon`;
CREATE TABLE `weapon`  (
  `id` bigint NOT NULL COMMENT '武器id',
  `name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '名称',
  `type` int NULL DEFAULT NULL COMMENT '类型',
  `level` int NULL DEFAULT NULL COMMENT '等级',
  `damage` bigint NULL DEFAULT NULL COMMENT '伤害值',
  `prob` bigint NULL DEFAULT NULL COMMENT '掉落概率',
  `rarity` int NULL DEFAULT NULL COMMENT '稀有程度',
  `count` int NULL DEFAULT NULL COMMENT '总数',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of weapon
-- ----------------------------

-- ----------------------------
-- Table structure for whitelist
-- ----------------------------
DROP TABLE IF EXISTS `whitelist`;
CREATE TABLE `whitelist`  (
  `id` bigint NOT NULL COMMENT '白名单',
  `ip` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'ip地址',
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `starttime` bigint NULL DEFAULT 0 COMMENT '起始时间',
  `endtime` bigint NULL DEFAULT 0 COMMENT '结束时间(0即为永久)',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of whitelist
-- ----------------------------

-- ----------------------------
-- Procedure structure for More20
-- ----------------------------
DROP PROCEDURE IF EXISTS `More20`;
delimiter ;;
CREATE PROCEDURE `More20`()
BEGIN
	#Routine body goes here...

DECLARE i INT DEFAULT 1;
DECLARE uname VARCHAR(20) DEFAULT ''; 
WHILE i<=40 DO
SET uname = CONCAT(randomname(),'麻将房');
INSERT INTO `minigame`.`game`( `type`, `kindid`, `hostid`, `level`, `name`, `password`, `maxonline`, `amount`, `enterscore`, `lessscore`, `playscore`, `state`, `commission`, `maxchair`) VALUES ( 3, 3002, 1, 0, uname, '000', 0, 100, 100*FLOOR( 10 + RAND() * (100 - 1)) , 500, 9900, 0, 1, 4);


SET i = i+1;

END WHILE;

END
;;
delimiter ;

-- ----------------------------
-- Procedure structure for proc_initData
-- ----------------------------
DROP PROCEDURE IF EXISTS `proc_initData`;
delimiter ;;
CREATE PROCEDURE `proc_initData`()
BEGIN

DECLARE i INT DEFAULT 1005; 
DECLARE uname CHAR(6) DEFAULT '';  
WHILE i<=1030 DO

SET uname = CONCAT(randomname(),'麻将房');
	INSERT INTO `minigame`.`game`(`id`, `type`, `kindid`, `hostid`, `level`, `name`, `password`, `maxonline`, `amount`, `enterscore`, `lessscore`, `playscore`, `state`, `commission`, `maxchair`) VALUES (i, 3, 3002, 1, 0, uname, '000', 0, 0, 9900, 600*FLOOR( 10 + RAND() * (100 - 1)), 9900, 0, 1, 5);

SET i = i+1;

END WHILE;

END
;;
delimiter ;

-- ----------------------------
-- Function structure for randomname
-- ----------------------------
DROP FUNCTION IF EXISTS `randomname`;
delimiter ;;
CREATE FUNCTION `randomname`()
 RETURNS char(4) CHARSET utf8mb3
  DETERMINISTIC
BEGIN

 DECLARE xing varchar(2056) DEFAULT '赵钱孙李周郑王冯陈楮卫蒋沈韩杨朱秦尤许何吕施张孔曹严华金魏陶姜戚谢喻柏水窦章云苏潘葛奚范彭郎鲁韦昌马苗凤花方俞任袁柳酆鲍史唐费廉岑薛雷贺倪汤滕殷罗毕郝邬安常乐于时傅皮齐康伍余元卜顾孟平黄和穆萧尹姚邵湛汪祁毛禹狄米贝明臧计伏成戴谈宋茅庞熊纪舒屈项祝董梁杜阮蓝闽席季麻强贾路娄危江童颜郭梅盛林刁锺徐丘骆高夏蔡田樊胡凌霍虞万支柯昝管卢莫经裘缪干解应宗丁宣贲邓郁单杭洪包诸左石崔吉钮龚程嵇邢滑裴陆荣翁';

 
 

 DECLARE ming varchar(2056) DEFAULT '嘉懿煜城懿轩烨伟苑博伟泽熠彤鸿煊博涛烨霖烨华煜祺智宸正豪昊然明杰诚立轩立辉峻熙弘文熠彤鸿煊烨霖哲瀚鑫鹏致远俊驰雨泽烨磊晟睿天佑文昊修洁黎昕远航旭尧鸿涛伟祺轩越泽浩宇瑾瑜皓轩擎苍擎宇志泽睿渊楷瑞轩弘文哲瀚雨泽鑫磊梦琪忆之桃慕青问兰尔岚元香初夏沛菡傲珊曼文乐菱痴珊恨玉惜文香寒新柔语蓉海安夜蓉涵柏水桃醉蓝春儿语琴从彤傲晴语兰又菱碧彤元霜怜梦紫寒妙彤曼易南莲紫翠雨寒易烟如萱若南寻真晓亦向珊慕灵以蕊寻雁映易雪柳孤岚笑霜海云凝天沛珊寒云冰旋宛儿绿真盼儿晓霜碧凡夏菡曼香若烟半梦雅绿冰蓝灵槐平安书翠翠风香巧代云梦曼幼翠友巧听寒梦柏醉易访旋亦玉凌萱访卉怀亦笑蓝春翠靖柏夜蕾冰夏梦松书雪乐枫念薇靖雁寻春恨山从寒忆香觅波静曼凡旋以亦念露芷蕾千兰新波代真新蕾雁玉冷卉紫山千琴恨天傲芙盼山怀蝶冰兰山柏翠萱乐丹翠柔谷山之瑶冰露尔珍谷雪乐萱涵菡海莲傲蕾青槐冬儿易梦惜雪宛海之柔夏青亦瑶妙菡春竹修杰伟诚建辉晋鹏天磊绍辉泽洋明轩健柏煊昊强伟宸博超君浩子骞明辉鹏涛炎彬鹤轩越彬风华靖琪明诚高格光华国源宇晗昱涵润翰飞翰海昊乾浩博和安弘博鸿朗华奥华灿嘉慕坚秉建明金鑫锦程瑾瑜鹏经赋景同靖琪君昊俊明季同开济凯安康成乐语力勤良哲理群茂彦敏博明达朋义彭泽鹏举濮存溥心璞瑜浦泽奇邃祥荣轩';

 
 

 DECLARE l_xing int DEFAULT LENGTH(xing) / 3; # 这里的长度不是字符串的字数,而是此字符串的占的容量大小,一个汉字占3个字节

 DECLARE l_ming int DEFAULT LENGTH(ming) / 3;

 
 

 DECLARE return_str varchar(255) DEFAULT '';

 
 

 
 

 
 

 # 先选出姓

 SET return_str = CONCAT(return_str, SUBSTRING(xing, FLOOR(1 + RAND() * l_xing), 1));

 
 

 
 

 #再选出名

 SET return_str = CONCAT(return_str, SUBSTRING(ming, FLOOR(1 + RAND() * l_ming), 1));

 
 

 
 

 IF RAND()>0.400 THEN

 #再选出名

 SET return_str = CONCAT(return_str, SUBSTRING(ming, FLOOR(1 + RAND() * l_ming), 1));

 END IF;

 
 

 RETURN return_str;
END
;;
delimiter ;

-- ----------------------------
-- Procedure structure for u_head_and_low_pro
-- ----------------------------
DROP PROCEDURE IF EXISTS `u_head_and_low_pro`;
delimiter ;;
CREATE PROCEDURE `u_head_and_low_pro`()
BEGIN
	DECLARE
		n INT DEFAULT 1005;
	
	SET @exesql = 'INSERT INTO `minigame`.`game` (
	`id`,
	`type`,
	`kindid`,
	`hostid`,
	`level`,
	`name`,
	`password`,
	`maxonline`,
	`amount`,
	`enterscore`,
	`lessscore`,
	`playscore`,
	`state`,
	`commission`,
	`maxchair` 
	)values ';
	
	SET @exedata = '';
	WHILE
			n < 21 DO
			
			SET @exedata = concat(
				@exedata,
				"(",
				n,
				3,
				3002,
				1,
				0,
				'aaa',
				'000',
				0,
				0,
				9900,
				600,
				9900,
				0,
				1,
				5,
				")" 
			);
		IF
			n % 20 = 0 THEN
				
				SET @exesql = concat( @exesql, @exedata, ";" );
			PREPARE stmt 
			FROM
				@exesql;
			EXECUTE stmt;
			DEALLOCATE PREPARE stmt;
			COMMIT;
			
			SET @exesql = 'INSERT INTO `minigame`.`game` (
			`id`,
			`type`,
			`kindid`,
			`hostid`,
			`level`,
			`name`,
			`password`,
			`maxonline`,
			`amount`,
			`enterscore`,
			`lessscore`,
			`playscore`,
			`state`,
			`commission`,
			`maxchair` 
			)values ';
			
			SET @exedata = '';
			ELSE 
				SET @exedata = concat( @exedata, ',' );
			
		END IF;
		
		SET n = n + 1;
		
	END WHILE;
	
END
;;
delimiter ;

SET FOREIGN_KEY_CHECKS = 1;
