/*
 Navicat Premium Data Transfer

 Source Server         : 本地数据库
 Source Server Type    : MySQL
 Source Server Version : 80032
 Source Host           : localhost:3306
 Source Schema         : social

 Target Server Type    : MySQL
 Target Server Version : 80032
 File Encoding         : 65001

 Date: 20/09/2024 16:08:14
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for Chat
-- ----------------------------
DROP TABLE IF EXISTS `Chat`;
CREATE TABLE `Chat` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '聊天记录唯一标识',
  `channel` int DEFAULT NULL COMMENT '聊天频道',
  `senderUid` bigint DEFAULT NULL COMMENT '发送者用户ID',
  `targetUid` bigint DEFAULT NULL COMMENT '接收者用户ID',
  `clubId` bigint DEFAULT NULL COMMENT '俱乐部ID（如果是俱乐部聊天）',
  `timeStamp` bigint DEFAULT NULL COMMENT '消息发送时间戳',
  `cont` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '消息内容',
  `gameEid` bigint DEFAULT NULL COMMENT '游戏ID（如果是游戏邀请）',
  `msgType` int DEFAULT NULL COMMENT '消息类型（0: 普通文本, 1: 邀请对战）',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_channel` (`channel`),
  KEY `idx_senderUid` (`senderUid`),
  KEY `idx_targetUid` (`targetUid`),
  KEY `idx_clubId` (`clubId`)
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '聊天记录表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for Club
-- ----------------------------
DROP TABLE IF EXISTS `Club`;
CREATE TABLE `Club` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '俱乐部唯一标识',
  `master` bigint DEFAULT NULL COMMENT '俱乐部主人用户ID',
  `builder` bigint DEFAULT NULL COMMENT '俱乐部创建者用户ID',
  `createdAt` bigint DEFAULT NULL COMMENT '俱乐部创建时间戳',
  `icon` int DEFAULT NULL COMMENT '俱乐部图标ID',
  `mode` int DEFAULT NULL COMMENT '俱乐部模式',
  `score` bigint DEFAULT NULL COMMENT '俱乐部总积分',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '俱乐部名称',
  `notice` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '俱乐部公告',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_master` (`master`),
  KEY `idx_builder` (`builder`)
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '俱乐部信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for ClubMember
-- ----------------------------
DROP TABLE IF EXISTS `ClubMember`;
CREATE TABLE `ClubMember` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '俱乐部成员记录唯一标识',
  `clubId` bigint DEFAULT NULL COMMENT '俱乐部ID',
  `uid` bigint DEFAULT NULL COMMENT '成员用户ID',
  `job` int DEFAULT NULL COMMENT '成员职位',
  `liveness` int DEFAULT NULL COMMENT '当前活跃度',
  `totalLiveness` int DEFAULT NULL COMMENT '总活跃度',
  `score` bigint DEFAULT NULL COMMENT '成员积分',
  `refereeUid` bigint DEFAULT NULL COMMENT '推荐人用户ID',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_clubId_uid` (`clubId`, `uid`),
  KEY `idx_uid` (`uid`)
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '俱乐部成员表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for ClubApply
-- ----------------------------
DROP TABLE IF EXISTS `ClubApply`;
CREATE TABLE `ClubApply` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '俱乐部申请记录唯一标识',
  `clubId` bigint DEFAULT NULL COMMENT '申请加入的俱乐部ID',
  `uid` bigint DEFAULT NULL COMMENT '申请人用户ID',
  `applyTime` bigint DEFAULT NULL COMMENT '申请时间戳',
  `status` int DEFAULT NULL COMMENT '申请状态（0: 待处理, 1: 已同意, 2: 已拒绝）',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_clubId_uid` (`clubId`, `uid`),
  KEY `idx_uid` (`uid`)
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '俱乐部申请表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for ClubInvite
-- ----------------------------
DROP TABLE IF EXISTS `ClubInvite`;
CREATE TABLE `ClubInvite` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '俱乐部邀请记录唯一标识',
  `clubId` bigint DEFAULT NULL COMMENT '邀请加入的俱乐部ID',
  `senderUid` bigint DEFAULT NULL COMMENT '邀请人用户ID',
  `targetUid` bigint DEFAULT NULL COMMENT '被邀请人用户ID',
  `inviteTime` bigint DEFAULT NULL COMMENT '邀请时间戳',
  `status` int DEFAULT NULL COMMENT '邀请状态（0: 待处理, 1: 已同意, 2: 已拒绝）',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_clubId_targetUid` (`clubId`, `targetUid`),
  KEY `idx_senderUid` (`senderUid`)
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '俱乐部邀请表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for Friend
-- ----------------------------
DROP TABLE IF EXISTS `Friend`;
CREATE TABLE `Friend` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '好友关系记录唯一标识',
  `uid` bigint DEFAULT NULL COMMENT '用户ID',
  `friendUid` bigint DEFAULT NULL COMMENT '好友用户ID',
  `addTime` bigint DEFAULT NULL COMMENT '添加好友的时间戳',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_uid_friendUid` (`uid`, `friendUid`)
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '好友关系表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for FriendApply
-- ----------------------------
DROP TABLE IF EXISTS `FriendApply`;
CREATE TABLE `FriendApply` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '好友申请记录唯一标识',
  `senderUid` bigint DEFAULT NULL COMMENT '申请人用户ID',
  `targetUid` bigint DEFAULT NULL COMMENT '被申请人用户ID',
  `cont` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '申请附言',
  `applyTime` bigint DEFAULT NULL COMMENT '申请时间戳',
  `status` int DEFAULT NULL COMMENT '申请状态（0: 待处理, 1: 已同意, 2: 已拒绝）',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_senderUid_targetUid` (`senderUid`, `targetUid`),
  KEY `idx_targetUid` (`targetUid`)
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '好友申请表' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;