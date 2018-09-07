-- --------------------------------------------------------
-- 主机:                           43.247.70.20
-- 服务器版本:                        5.5.33-log - MySQL Community Server (GPL)
-- 服务器操作系统:                      Linux
-- HeidiSQL 版本:                  9.3.0.4984
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

-- 导出 credit 的数据库结构
DROP DATABASE IF EXISTS `credit`;
CREATE DATABASE IF NOT EXISTS `credit` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `credit`;


-- 导出  表 credit.t_common_category 结构
DROP TABLE IF EXISTS `t_common_category`;
CREATE TABLE IF NOT EXISTS `t_common_category` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `type` int(11) NOT NULL COMMENT '类型1: 评级者；2：项目',
  `category` varchar(50) NOT NULL COMMENT '分类名称',
  `adminUid` bigint(20) DEFAULT NULL COMMENT '管理操作员',
  `updateTime` datetime DEFAULT NULL COMMENT '最后修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COMMENT='分类';

-- 正在导出表  credit.t_common_category 的数据：~4 rows (大约)
/*!40000 ALTER TABLE `t_common_category` DISABLE KEYS */;
REPLACE INTO `t_common_category` (`id`, `type`, `category`, `adminUid`, `updateTime`) VALUES
	(1, 1, '架构', 0, '2018-09-07 01:14:52'),
	(2, 1, '安全', 0, '2018-09-07 01:15:26'),
	(3, 2, '公链', 0, '2018-09-07 01:15:44'),
	(4, 2, '应用', 0, '2018-09-07 01:16:07');
/*!40000 ALTER TABLE `t_common_category` ENABLE KEYS */;


-- 导出  表 credit.t_common_item 结构
DROP TABLE IF EXISTS `t_common_item`;
CREATE TABLE IF NOT EXISTS `t_common_item` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL COMMENT '指标名称',
  `value` double NOT NULL COMMENT '指标数值（默认为0）',
  `level` int(11) NOT NULL COMMENT '维度等级 1：重要指标；2：相关指标',
  `weight` double NOT NULL COMMENT '维度的权重，0~1之间的数字',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 COMMENT='维度';

-- 正在导出表  credit.t_common_item 的数据：~5 rows (大约)
/*!40000 ALTER TABLE `t_common_item` DISABLE KEYS */;
REPLACE INTO `t_common_item` (`id`, `name`, `value`, `level`, `weight`) VALUES
	(1, '架构清晰性', 100, 1, 0.8),
	(2, '可用性', 100, 1, 0.8),
	(3, '安全性', 100, 1, 1),
	(4, '选举公平性', 100, 2, 0.3),
	(5, '挖矿公平性', 100, 2, 0.6);
/*!40000 ALTER TABLE `t_common_item` ENABLE KEYS */;


-- 导出  表 credit.t_follow_social 结构
DROP TABLE IF EXISTS `t_follow_social`;
CREATE TABLE IF NOT EXISTS `t_follow_social` (
  `followeRrid` bigint(20) unsigned NOT NULL COMMENT '关注者',
  `befolloweRid` bigint(20) unsigned NOT NULL COMMENT '被关注者',
  `social` int(11) NOT NULL DEFAULT '1' COMMENT '关系1:关注,2:互相关注',
  `createTime` datetime NOT NULL COMMENT '关注时间',
  KEY `followerid_befollowerid` (`followeRrid`,`befolloweRid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='社交关系';

-- 正在导出表  credit.t_follow_social 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `t_follow_social` DISABLE KEYS */;
/*!40000 ALTER TABLE `t_follow_social` ENABLE KEYS */;


-- 导出  表 credit.t_project 结构
DROP TABLE IF EXISTS `t_project`;
CREATE TABLE IF NOT EXISTS `t_project` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '项目名称',
  `url` varchar(200) NOT NULL DEFAULT '' COMMENT '项目官网',
  `logo` varchar(200) NOT NULL DEFAULT '' COMMENT '项目LOGO',
  `token` varchar(50) NOT NULL DEFAULT '' COMMENT 'token代码',
  `tags` varchar(200) NOT NULL DEFAULT '' COMMENT '标签',
  `description` varchar(2000) NOT NULL DEFAULT '' COMMENT '项目描述',
  `score` double NOT NULL DEFAULT '0' COMMENT '项目总分',
  `rater` int(11) NOT NULL DEFAULT '0' COMMENT '评级人数',
  PRIMARY KEY (`id`),
  KEY `index_token` (`token`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='项目';

-- 正在导出表  credit.t_project 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `t_project` DISABLE KEYS */;
/*!40000 ALTER TABLE `t_project` ENABLE KEYS */;


-- 导出  表 credit.t_project_category 结构
DROP TABLE IF EXISTS `t_project_category`;
CREATE TABLE IF NOT EXISTS `t_project_category` (
  `projectId` bigint(20) unsigned NOT NULL,
  `categoryId` bigint(20) unsigned NOT NULL,
  `category` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='项目的分类信息，多对多';

-- 正在导出表  credit.t_project_category 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `t_project_category` DISABLE KEYS */;
/*!40000 ALTER TABLE `t_project_category` ENABLE KEYS */;


-- 导出  表 credit.t_project_rate 结构
DROP TABLE IF EXISTS `t_project_rate`;
CREATE TABLE IF NOT EXISTS `t_project_rate` (
  `projectId` bigint(20) unsigned NOT NULL COMMENT '项目ID',
  `itemId` bigint(20) unsigned NOT NULL COMMENT '指标ID',
  `score` double NOT NULL COMMENT '项目指标平均分',
  `raterCount` int(11) NOT NULL COMMENT '评级人数',
  KEY `projectId_itemId` (`projectId`,`itemId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 正在导出表  credit.t_project_rate 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `t_project_rate` DISABLE KEYS */;
/*!40000 ALTER TABLE `t_project_rate` ENABLE KEYS */;


-- 导出  表 credit.t_rater 结构
DROP TABLE IF EXISTS `t_rater`;
CREATE TABLE IF NOT EXISTS `t_rater` (
  `userId` bigint(20) unsigned NOT NULL COMMENT '评级者用户ID',
  `status` int(11) NOT NULL COMMENT '评级者申请的状态：',
  `creditCodes` varchar(200) DEFAULT NULL COMMENT '信誉码数组',
  `certificates` varchar(200) DEFAULT NULL COMMENT '证件照数组',
  `level` int(11) DEFAULT NULL COMMENT '评级者等级',
  `categorys` varchar(200) DEFAULT NULL COMMENT '分类',
  `auditUid` bigint(20) DEFAULT NULL COMMENT '审核者用户',
  `auditDescribe` varchar(200) DEFAULT NULL COMMENT '审核描述'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='评级者';

-- 正在导出表  credit.t_rater 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `t_rater` DISABLE KEYS */;
/*!40000 ALTER TABLE `t_rater` ENABLE KEYS */;


-- 导出  表 credit.t_rater_category 结构
DROP TABLE IF EXISTS `t_rater_category`;
CREATE TABLE IF NOT EXISTS `t_rater_category` (
  `raterId` bigint(20) unsigned DEFAULT NULL COMMENT '评级者',
  `categoryId` bigint(20) unsigned DEFAULT NULL,
  `category` varchar(50) DEFAULT NULL COMMENT '分类'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='分类';

-- 正在导出表  credit.t_rater_category 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `t_rater_category` DISABLE KEYS */;
/*!40000 ALTER TABLE `t_rater_category` ENABLE KEYS */;


-- 导出  表 credit.t_rater_credit 结构
DROP TABLE IF EXISTS `t_rater_credit`;
CREATE TABLE IF NOT EXISTS `t_rater_credit` (
  `senderRid` bigint(20) unsigned NOT NULL COMMENT '生产信誉码的评级者',
  `code` varchar(50) NOT NULL COMMENT '信誉码',
  `receiverRid` bigint(20) NOT NULL COMMENT '使用信誉码的评级者',
  `index` int(11) NOT NULL COMMENT '生产信誉码的index值',
  `sendTime` datetime NOT NULL COMMENT '生成时间',
  `receiveTime` datetime NOT NULL COMMENT '使用时间',
  KEY `sender_rid` (`senderRid`),
  KEY `receiver_rid` (`receiverRid`),
  KEY `code` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='评级者的信誉码';

-- 正在导出表  credit.t_rater_credit 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `t_rater_credit` DISABLE KEYS */;
/*!40000 ALTER TABLE `t_rater_credit` ENABLE KEYS */;


-- 导出  表 credit.t_rater_project 结构
DROP TABLE IF EXISTS `t_rater_project`;
CREATE TABLE IF NOT EXISTS `t_rater_project` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `raterId` bigint(20) DEFAULT NULL COMMENT '评审者',
  `projectId` bigint(20) DEFAULT NULL COMMENT '项目',
  `createTime` datetime DEFAULT NULL COMMENT '评审时间',
  `score` double DEFAULT NULL COMMENT '评审分数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='评级者评的项目';

-- 正在导出表  credit.t_rater_project 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `t_rater_project` DISABLE KEYS */;
/*!40000 ALTER TABLE `t_rater_project` ENABLE KEYS */;


-- 导出  表 credit.t_rater_project_item 结构
DROP TABLE IF EXISTS `t_rater_project_item`;
CREATE TABLE IF NOT EXISTS `t_rater_project_item` (
  `rpId` bigint(20) unsigned NOT NULL COMMENT 't_rater_project表的ID',
  `itemId` bigint(20) unsigned NOT NULL COMMENT 't_common_item表ID',
  `value` double NOT NULL COMMENT '具体分数',
  `createTime` double NOT NULL COMMENT '评分时间',
  `content` varchar(500) DEFAULT NULL COMMENT '评审详细内容'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='评审项目的具体分数';

-- 正在导出表  credit.t_rater_project_item 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `t_rater_project_item` DISABLE KEYS */;
/*!40000 ALTER TABLE `t_rater_project_item` ENABLE KEYS */;


-- 导出  表 credit.t_user 结构
DROP TABLE IF EXISTS `t_user`;
CREATE TABLE IF NOT EXISTS `t_user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(100) NOT NULL DEFAULT '' COMMENT '用户登陆名（手机号）',
  `salt` varchar(100) NOT NULL DEFAULT '' COMMENT '密码盐',
  `password` varchar(100) NOT NULL DEFAULT '' COMMENT '加盐之后的密码md5(md5(password) + salt)',
  `authtoken` varchar(100) NOT NULL DEFAULT '' COMMENT '验证令牌',
  `type` int(11) NOT NULL DEFAULT '1' COMMENT '1:普通用户 2:管理员用户',
  `createTime` datetime DEFAULT NULL COMMENT '创建时间',
  `loginTime` datetime DEFAULT NULL COMMENT '登陆时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表';

-- 正在导出表  credit.t_user 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `t_user` DISABLE KEYS */;
/*!40000 ALTER TABLE `t_user` ENABLE KEYS */;


-- 导出  表 credit.t_user_info 结构
DROP TABLE IF EXISTS `t_user_info`;
CREATE TABLE IF NOT EXISTS `t_user_info` (
  `userId` bigint(20) unsigned NOT NULL COMMENT 't_user表ID',
  `nikename` varchar(50) NOT NULL COMMENT '用户昵称',
  `icon` varchar(200) NOT NULL COMMENT '用户头像',
  `level` int(11) NOT NULL COMMENT '用户等级',
  `credit` double NOT NULL COMMENT '用户信用分',
  `income` double NOT NULL COMMENT '用户总收入',
  `balance` double NOT NULL COMMENT '用户余额',
  `tags` varchar(200) NOT NULL COMMENT '用户标签',
  `project` int(11) NOT NULL COMMENT '评级项目数',
  `inviteCode` int(11) NOT NULL COMMENT '我的邀请码',
  `promotionCount` int(11) NOT NULL COMMENT '邀请码推广人数'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 正在导出表  credit.t_user_info 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `t_user_info` DISABLE KEYS */;
/*!40000 ALTER TABLE `t_user_info` ENABLE KEYS */;


-- 导出  表 credit.t_user_invite 结构
DROP TABLE IF EXISTS `t_user_invite`;
CREATE TABLE IF NOT EXISTS `t_user_invite` (
  `senderUid` bigint(20) unsigned NOT NULL COMMENT '邀请者用户',
  `code` varchar(50) NOT NULL COMMENT '邀请码',
  `receiverUid` bigint(20) unsigned DEFAULT NULL COMMENT '被邀请者用户',
  `income` double unsigned DEFAULT NULL COMMENT '收入',
  `sendTime` datetime DEFAULT NULL COMMENT '邀请时间',
  `receiveTime` datetime DEFAULT NULL COMMENT '接收邀约时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户邀请';

-- 正在导出表  credit.t_user_invite 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `t_user_invite` DISABLE KEYS */;
/*!40000 ALTER TABLE `t_user_invite` ENABLE KEYS */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
