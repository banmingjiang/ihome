-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        5.5.53 - MySQL Community Server (GPL)
-- 服务器操作系统:                      Win32
-- HeidiSQL 版本:                  9.3.0.4984
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

-- 导出 ihome 的数据库结构
CREATE DATABASE IF NOT EXISTS `ihome` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `ihome`;


-- 导出  表 ihome.area 结构
CREATE TABLE IF NOT EXISTS `area` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;

-- 正在导出表  ihome.area 的数据：15 rows
DELETE FROM `area`;
/*!40000 ALTER TABLE `area` DISABLE KEYS */;
INSERT INTO `area` (`id`, `name`) VALUES
	(1, '东城区'),
	(2, '西城区'),
	(3, '朝阳区'),
	(4, '海定区'),
	(5, '昌平区'),
	(6, '丰台区'),
	(7, '房山区'),
	(8, '通州区'),
	(9, '顺义区'),
	(10, '大兴区'),
	(11, '怀柔区'),
	(12, '平谷区'),
	(13, '密云区'),
	(14, '延庆区'),
	(15, '石景山区');
/*!40000 ALTER TABLE `area` ENABLE KEYS */;


-- 导出  表 ihome.facility 结构
CREATE TABLE IF NOT EXISTS `facility` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=26 DEFAULT CHARSET=utf8;

-- 正在导出表  ihome.facility 的数据：25 rows
DELETE FROM `facility`;
/*!40000 ALTER TABLE `facility` DISABLE KEYS */;
INSERT INTO `facility` (`id`, `name`) VALUES
	(1, '无线网络'),
	(2, '热水沐浴'),
	(3, '空调'),
	(4, '暖气'),
	(5, '允许抽烟'),
	(6, '饮水设备'),
	(7, '牙具'),
	(8, '香皂'),
	(9, '拖鞋'),
	(10, '手纸'),
	(11, '毛巾'),
	(12, '沐浴露，洗发露'),
	(13, '冰箱'),
	(14, '洗衣机'),
	(15, '电梯'),
	(16, '允许做饭'),
	(17, '允许带宠物'),
	(18, '允许聚会'),
	(19, '门禁系统'),
	(20, '停车位'),
	(21, '有线网络'),
	(22, '电视'),
	(23, '浴缸'),
	(24, '吃鸡'),
	(25, '打台球');
/*!40000 ALTER TABLE `facility` ENABLE KEYS */;


-- 导出  表 ihome.facility_houses 结构
CREATE TABLE IF NOT EXISTS `facility_houses` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `facility_id` int(11) NOT NULL,
  `house_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- 正在导出表  ihome.facility_houses 的数据：0 rows
DELETE FROM `facility_houses`;
/*!40000 ALTER TABLE `facility_houses` DISABLE KEYS */;
/*!40000 ALTER TABLE `facility_houses` ENABLE KEYS */;


-- 导出  表 ihome.house 结构
CREATE TABLE IF NOT EXISTS `house` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `area_id` int(11) NOT NULL,
  `title` varchar(64) NOT NULL DEFAULT '',
  `price` int(11) NOT NULL DEFAULT '0',
  `address` varchar(512) NOT NULL DEFAULT '',
  `room_count` int(11) NOT NULL DEFAULT '1',
  `acreage` int(11) NOT NULL DEFAULT '0',
  `unit` varchar(32) NOT NULL DEFAULT '',
  `capacity` int(11) NOT NULL DEFAULT '1',
  `beds` varchar(64) NOT NULL DEFAULT '',
  `deposit` int(11) NOT NULL DEFAULT '0',
  `min_days` int(11) NOT NULL DEFAULT '0',
  `max_days` int(11) NOT NULL DEFAULT '0',
  `order_count` int(11) NOT NULL DEFAULT '0',
  `index_image_url` varchar(256) NOT NULL DEFAULT '',
  `ctime` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- 正在导出表  ihome.house 的数据：0 rows
DELETE FROM `house`;
/*!40000 ALTER TABLE `house` DISABLE KEYS */;
/*!40000 ALTER TABLE `house` ENABLE KEYS */;


-- 导出  表 ihome.house_image 结构
CREATE TABLE IF NOT EXISTS `house_image` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `url` varchar(256) NOT NULL DEFAULT '',
  `house_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- 正在导出表  ihome.house_image 的数据：0 rows
DELETE FROM `house_image`;
/*!40000 ALTER TABLE `house_image` DISABLE KEYS */;
/*!40000 ALTER TABLE `house_image` ENABLE KEYS */;


-- 导出  表 ihome.order_house 结构
CREATE TABLE IF NOT EXISTS `order_house` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `house_id` int(11) NOT NULL,
  `begin_date` datetime NOT NULL,
  `end_date` datetime NOT NULL,
  `days` int(11) NOT NULL DEFAULT '0',
  `house_price` int(11) NOT NULL DEFAULT '0',
  `amount` int(11) NOT NULL DEFAULT '0',
  `status` varchar(255) NOT NULL DEFAULT 'WAIT_ACCEPT',
  `coment` varchar(512) NOT NULL DEFAULT '',
  `ctime` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- 正在导出表  ihome.order_house 的数据：0 rows
DELETE FROM `order_house`;
/*!40000 ALTER TABLE `order_house` DISABLE KEYS */;
/*!40000 ALTER TABLE `order_house` ENABLE KEYS */;


-- 导出  表 ihome.user 结构
CREATE TABLE IF NOT EXISTS `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL DEFAULT '',
  `password_hash` varchar(128) NOT NULL DEFAULT '',
  `mobile` varchar(11) NOT NULL DEFAULT '',
  `real_name` varchar(32) NOT NULL DEFAULT '',
  `id_card` varchar(20) NOT NULL DEFAULT '',
  `avatar_url` varchar(256) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;

-- 正在导出表  ihome.user 的数据：8 rows
DELETE FROM `user`;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` (`id`, `name`, `password_hash`, `mobile`, `real_name`, `id_card`, `avatar_url`) VALUES
	(1, '', 'f5557d4fcf727a981a3c315aca733eefa2996f7c7cdae1fa7e0de28522820bb0', 'iii', 'iii', '', ''),
	(2, '', 'a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3', '456', '123', '', ''),
	(3, '', '1ad25d0002690dc02e2708a297d8c9df1f160d376f663309cc261c7c921367e7', 'ooo', 'ooo', '', ''),
	(4, '', 'a8c23cc814179578e3a774418ac5fc4702a66eb3b78c876df81b290465e6e334', 'oo', 'oo', '', ''),
	(5, '', '5afab9a620f6f11284505be2fb9a975b4dccfdd30970dffc7ed875490160e4d0', 'uu', 'uu', '', ''),
	(6, '', '5afab9a620f6f11284505be2fb9a975b4dccfdd30970dffc7ed875490160e4d0', 'uuu', 'uuu', '', ''),
	(7, '', '95e17a94411dbc199b8ac626d722f42a63d4521f62ddd78bf3dffa615471097b', 'uio', 'uio', '', ''),
	(8, '', 'a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3', '123456', '123456', '', '');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
