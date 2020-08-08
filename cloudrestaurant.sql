/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 80011
 Source Host           : localhost:3306
 Source Schema         : cloudrestaurant

 Target Server Type    : MySQL
 Target Server Version : 80011
 File Encoding         : 65001

 Date: 18/03/2020 14:32:58
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for food_category
-- ----------------------------
DROP TABLE IF EXISTS `food_category`;
CREATE TABLE `food_category` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `title` varchar(20) DEFAULT NULL,
  `description` varchar(30) DEFAULT NULL,
  `image_url` varchar(255) DEFAULT NULL,
  `link_url` varchar(255) DEFAULT NULL,
  `is_in_serving` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of food_category
-- ----------------------------
BEGIN;
INSERT INTO `food_category` VALUES (1, NULL, '预定早餐', NULL, NULL, 0);
INSERT INTO `food_category` VALUES (2, NULL, '果蔬生鲜', NULL, NULL, 0);
INSERT INTO `food_category` VALUES (3, NULL, '鲜花蛋糕', NULL, NULL, 0);
INSERT INTO `food_category` VALUES (4, NULL, '商超便利', NULL, NULL, 0);
INSERT INTO `food_category` VALUES (5, NULL, '美食', NULL, NULL, 0);
INSERT INTO `food_category` VALUES (6, NULL, '甜品饮品', NULL, NULL, 0);
INSERT INTO `food_category` VALUES (7, NULL, '土豪推荐', NULL, NULL, 0);
INSERT INTO `food_category` VALUES (8, NULL, '准时达', NULL, NULL, 0);
INSERT INTO `food_category` VALUES (9, NULL, '简餐', NULL, NULL, 0);
INSERT INTO `food_category` VALUES (10, NULL, '汉堡薯条', NULL, NULL, 0);
INSERT INTO `food_category` VALUES (11, NULL, '日韩料理', NULL, NULL, 0);
INSERT INTO `food_category` VALUES (12, NULL, '麻辣烫', NULL, NULL, 0);
INSERT INTO `food_category` VALUES (13, NULL, '披萨意面', NULL, NULL, 0);
INSERT INTO `food_category` VALUES (14, NULL, '川湘菜', NULL, NULL, 0);
INSERT INTO `food_category` VALUES (15, NULL, '包子粥铺', NULL, NULL, 0);
INSERT INTO `food_category` VALUES (16, NULL, '新店特惠', NULL, NULL, 0);
COMMIT;

-- ----------------------------
-- Table structure for goods
-- ----------------------------
DROP TABLE IF EXISTS `goods`;
CREATE TABLE `goods` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(12) DEFAULT NULL,
  `description` varchar(32) DEFAULT NULL,
  `icon` varchar(255) DEFAULT NULL,
  `sell_count` int(11) DEFAULT NULL,
  `price` float DEFAULT NULL,
  `old_price` float DEFAULT NULL,
  `shop_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of goods
-- ----------------------------
BEGIN;
INSERT INTO `goods` VALUES (1, '小小鲜肉包', '滑蛋牛肉粥(1份)+小小鲜肉包(4只)', '', 14, 25, 29, 1);
INSERT INTO `goods` VALUES (2, '滑蛋牛肉粥+小小鲜肉包', '滑蛋牛肉粥(1份)+小小鲜肉包(3只)', '', 6, 35, 41, 1);
INSERT INTO `goods` VALUES (3, '滑蛋牛肉粥+绿甘蓝馅饼', '滑蛋牛肉粥(1份)+绿甘蓝馅饼(1张)', '', 2, 25, 30, 1);
INSERT INTO `goods` VALUES (4, '茶香卤味蛋', '咸鸡蛋', '', 688, 2.5, 3, 1);
INSERT INTO `goods` VALUES (5, '韭菜鸡蛋馅饼(2张)', '韭菜鸡蛋馅饼', '', 381, 10, 12, 1);
INSERT INTO `goods` VALUES (6, '小小鲜肉包+豆浆套餐', '小小鲜肉包(8只)装+豆浆(1杯)', '', 335, 9.9, 11.9, 479);
INSERT INTO `goods` VALUES (7, '翠香炒素饼', '咸鲜翠香素炒饼', '', 260, 17.9, 20.9, 485);
INSERT INTO `goods` VALUES (8, '香煎鲜肉包', '咸鲜猪肉鲜肉包', '', 173, 10.9, 12.9, 486);
COMMIT;

-- ----------------------------
-- Table structure for member
-- ----------------------------
DROP TABLE IF EXISTS `member`;
CREATE TABLE `member` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(20) DEFAULT NULL,
  `mobile` varchar(11) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `register_time` bigint(20) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `balance` double DEFAULT NULL,
  `is_active` tinyint(4) DEFAULT NULL,
  `city` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of member
-- ----------------------------
BEGIN;
INSERT INTO `member` VALUES (1, '18569430588', '18569430588', '123456', 1582095444, '', 0, 0, '');
COMMIT;

-- ----------------------------
-- Table structure for service
-- ----------------------------
DROP TABLE IF EXISTS `service`;
CREATE TABLE `service` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) DEFAULT NULL,
  `description` varchar(30) DEFAULT NULL,
  `icon_name` varchar(3) DEFAULT NULL,
  `icon_color` varchar(6) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of service
-- ----------------------------
BEGIN;
INSERT INTO `service` VALUES (4, '开发票', '该商家支持开发票，请在下单时填写好发票抬头', '票', '999999');
INSERT INTO `service` VALUES (7, '外卖保', '已加入“外卖保”计划，食品安全有保障', '保', '999999');
INSERT INTO `service` VALUES (9, '准时达', '准时必达，超时秒赔', '准', '57A9FF');
COMMIT;

-- ----------------------------
-- Table structure for shop
-- ----------------------------
DROP TABLE IF EXISTS `shop`;
CREATE TABLE `shop` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(12) DEFAULT NULL,
  `promotion_info` varchar(30) DEFAULT NULL,
  `address` varchar(100) DEFAULT NULL,
  `phone` varchar(11) DEFAULT NULL,
  `status` tinyint(4) DEFAULT NULL,
  `longitude` double DEFAULT NULL,
  `latitude` double DEFAULT NULL,
  `image_path` varchar(255) DEFAULT NULL,
  `is_new` tinyint(1) DEFAULT NULL,
  `is_premium` tinyint(1) DEFAULT NULL,
  `rating` float DEFAULT NULL,
  `rating_count` int(11) DEFAULT NULL,
  `recent_order_num` int(11) DEFAULT NULL,
  `minimum_order_amount` int(11) DEFAULT NULL,
  `delivery_fee` int(11) DEFAULT NULL,
  `opening_hours` varchar(20) DEFAULT NULL,
  `supports` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=489 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of shop
-- ----------------------------
BEGIN;
INSERT INTO `shop` VALUES (1, '嘉禾一品（温都水城）', '欢迎光临，用餐高峰请提前下单，谢谢', '北京市昌平区宏福苑温都水城F1', '13437850035', 1, 116.36868, 40.10039, '', 1, 1, 4.7, 961, 106, 20, 5, '8:30/20:30', NULL);
INSERT INTO `shop` VALUES (479, '杨国福麻辣烫', '欢迎光临，用餐高峰请提前下单，谢谢', '北京市市蜀山区南二环路天鹅湖万达广场8号楼1705室', '13167583411', 1, 117.22124, 31.81948, '', 1, 1, 4.2, 167, 755, 20, 5, '8:30/20:30', NULL);
INSERT INTO `shop` VALUES (485, '好适口', '欢迎光临，用餐高峰请提前下单，谢谢', '北京市海淀区西二旗大街58号', '12345678901', 1, 120.65355, 31.26578, '', 1, 1, 4.6, 576, 58, 20, 5, '8:30/20:30', NULL);
INSERT INTO `shop` VALUES (486, '东来顺旗舰店', '老北京正宗涮羊肉,非物质文化遗产', '北京市天河区东圃镇汇彩路38号1领汇创展商务中心401', '13544323775', 1, 113.41724, 23.1127, '', 1, 1, 4.2, 372, 542, 20, 5, '09:00/21:30', NULL);
INSERT INTO `shop` VALUES (487, '北京酒家', '北京第一家传承300年酒家', '北京市海淀区上下九商业步行街内', '13257482341', 0, 113.24826, 23.11488, '', 1, 1, 4.2, 871, 923, 20, 5, '8:30/20:30', NULL);
INSERT INTO `shop` VALUES (488, '和平鸽饺子馆', '吃饺子就来和平鸽饺子馆', '北京市越秀区德政中路171', '17098764762', 1, 113.27521, 23.12092, '', 1, 1, 4.2, 273, 483, 20, 5, '8:30/20:30', NULL);
COMMIT;

-- ----------------------------
-- Table structure for shop_service
-- ----------------------------
DROP TABLE IF EXISTS `shop_service`;
CREATE TABLE `shop_service` (
  `shop_id` bigint(20) NOT NULL,
  `service_id` bigint(20) NOT NULL,
  PRIMARY KEY (`shop_id`,`service_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of shop_service
-- ----------------------------
BEGIN;
INSERT INTO `shop_service` VALUES (1, 4);
INSERT INTO `shop_service` VALUES (1, 7);
INSERT INTO `shop_service` VALUES (1, 9);
INSERT INTO `shop_service` VALUES (479, 4);
INSERT INTO `shop_service` VALUES (479, 7);
INSERT INTO `shop_service` VALUES (479, 9);
COMMIT;

-- ----------------------------
-- Table structure for sms_code
-- ----------------------------
DROP TABLE IF EXISTS `sms_code`;
CREATE TABLE `sms_code` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `phone` varchar(11) DEFAULT NULL,
  `biz_id` varchar(30) DEFAULT NULL,
  `code` varchar(6) DEFAULT NULL,
  `create_time` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sms_code
-- ----------------------------
BEGIN;
INSERT INTO `sms_code` VALUES (1, '18569430588', '470615182095419483^0', '830642', 1582095419);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
