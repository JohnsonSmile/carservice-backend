/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50736
 Source Host           : localhost:3306
 Source Schema         : carservice

 Target Server Type    : MySQL
 Target Server Version : 50736
 File Encoding         : 65001

 Date: 20/11/2023 13:40:21
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_charge_fee
-- ----------------------------
DROP TABLE IF EXISTS `t_charge_fee`;
CREATE TABLE `t_charge_fee` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `fee_per_degree` int(11) NOT NULL,
  `position_id` bigint(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of t_charge_fee
-- ----------------------------
BEGIN;
INSERT INTO `t_charge_fee` (`id`, `created_at`, `updated_at`, `deleted_at`, `fee_per_degree`, `position_id`) VALUES (1, '2023-11-18 15:21:49', '2023-11-18 15:21:49', NULL, 205, 3);
COMMIT;

-- ----------------------------
-- Table structure for t_city
-- ----------------------------
DROP TABLE IF EXISTS `t_city`;
CREATE TABLE `t_city` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of t_city
-- ----------------------------
BEGIN;
INSERT INTO `t_city` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`) VALUES (1, '2023-11-17 20:07:35', '2023-11-17 20:07:35', NULL, '深圳');
INSERT INTO `t_city` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`) VALUES (2, '2023-11-17 20:07:35', '2023-11-17 20:07:35', NULL, '香港');
COMMIT;

-- ----------------------------
-- Table structure for t_high_way_fee
-- ----------------------------
DROP TABLE IF EXISTS `t_high_way_fee`;
CREATE TABLE `t_high_way_fee` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `start_position_id` bigint(20) NOT NULL,
  `end_position_id` bigint(20) NOT NULL,
  `fee` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of t_high_way_fee
-- ----------------------------
BEGIN;
INSERT INTO `t_high_way_fee` (`id`, `created_at`, `updated_at`, `deleted_at`, `start_position_id`, `end_position_id`, `fee`) VALUES (1, '2023-11-17 20:07:35', '2023-11-17 20:07:35', NULL, 1, 2, 1004);
COMMIT;

-- ----------------------------
-- Table structure for t_order
-- ----------------------------
DROP TABLE IF EXISTS `t_order`;
CREATE TABLE `t_order` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `order_type_id` bigint(20) NOT NULL,
  `order_status` tinyint(1) NOT NULL,
  `user_id` int(11) NOT NULL,
  `end_at` datetime DEFAULT NULL,
  `fee` int(11) NOT NULL,
  `start_position_id` bigint(20) NOT NULL,
  `end_position_id` bigint(20) NOT NULL,
  `order_sn` varchar(255) NOT NULL,
  `start_at` datetime DEFAULT NULL,
  `unit_count` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of t_order
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for t_order_type
-- ----------------------------
DROP TABLE IF EXISTS `t_order_type`;
CREATE TABLE `t_order_type` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `fee_per_unit` int(11) NOT NULL,
  `unit` varchar(10) NOT NULL,
  `name` varchar(10) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of t_order_type
-- ----------------------------
BEGIN;
INSERT INTO `t_order_type` (`id`, `created_at`, `updated_at`, `deleted_at`, `fee_per_unit`, `unit`, `name`) VALUES (1, '2023-11-17 20:07:35', '2023-11-17 20:07:35', NULL, 500, '公里', '高速公路');
INSERT INTO `t_order_type` (`id`, `created_at`, `updated_at`, `deleted_at`, `fee_per_unit`, `unit`, `name`) VALUES (2, '2023-11-17 20:07:35', '2023-11-17 20:07:35', NULL, 400, '度', '充电桩');
INSERT INTO `t_order_type` (`id`, `created_at`, `updated_at`, `deleted_at`, `fee_per_unit`, `unit`, `name`) VALUES (3, '2023-11-17 20:07:35', '2023-11-17 20:07:35', NULL, 1000, '小时', '停车场');
COMMIT;

-- ----------------------------
-- Table structure for t_park_fee
-- ----------------------------
DROP TABLE IF EXISTS `t_park_fee`;
CREATE TABLE `t_park_fee` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `fee_per_hour` int(11) NOT NULL,
  `position_id` bigint(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of t_park_fee
-- ----------------------------
BEGIN;
INSERT INTO `t_park_fee` (`id`, `created_at`, `updated_at`, `deleted_at`, `fee_per_hour`, `position_id`) VALUES (1, '2023-11-18 19:57:07', '2023-11-18 19:57:07', NULL, 100, 4);
COMMIT;

-- ----------------------------
-- Table structure for t_place
-- ----------------------------
DROP TABLE IF EXISTS `t_place`;
CREATE TABLE `t_place` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `place_type` tinyint(1) NOT NULL,
  `name` varchar(255) NOT NULL,
  `region_id` bigint(20) NOT NULL,
  `longitude` decimal(10,6) NOT NULL,
  `latitude` decimal(10,6) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1003 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of t_place
-- ----------------------------
BEGIN;
INSERT INTO `t_place` (`id`, `created_at`, `updated_at`, `deleted_at`, `place_type`, `name`, `region_id`, `longitude`, `latitude`) VALUES (11, '2023-11-18 15:21:49', '2023-11-18 15:21:49', NULL, 2, '万象城-001', 1, 24.150000, 113.230000);
INSERT INTO `t_place` (`id`, `created_at`, `updated_at`, `deleted_at`, `place_type`, `name`, `region_id`, `longitude`, `latitude`) VALUES (21, '2023-11-18 19:57:07', '2023-11-18 19:57:07', NULL, 3, '前海湾-001', 1, 24.150000, 113.230000);
INSERT INTO `t_place` (`id`, `created_at`, `updated_at`, `deleted_at`, `place_type`, `name`, `region_id`, `longitude`, `latitude`) VALUES (1001, '2023-11-17 20:07:35', '2023-11-17 20:07:35', NULL, 1, '宝安中心-001', 1, 24.150000, 113.230000);
INSERT INTO `t_place` (`id`, `created_at`, `updated_at`, `deleted_at`, `place_type`, `name`, `region_id`, `longitude`, `latitude`) VALUES (1002, '2023-11-17 20:07:35', '2023-11-17 20:07:35', NULL, 1, '尖沙咀-A', 2, 23.150000, 123.230000);
COMMIT;

-- ----------------------------
-- Table structure for t_position
-- ----------------------------
DROP TABLE IF EXISTS `t_position`;
CREATE TABLE `t_position` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(255) NOT NULL,
  `order_type_id` tinyint(1) NOT NULL,
  `city_id` bigint(20) NOT NULL,
  `region_id` bigint(20) NOT NULL,
  `place_id` bigint(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of t_position
-- ----------------------------
BEGIN;
INSERT INTO `t_position` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `order_type_id`, `city_id`, `region_id`, `place_id`) VALUES (1, '2023-11-17 20:07:35', '2023-11-17 20:07:35', NULL, '宝安中心-金城路A', 1, 1, 1, 1001);
INSERT INTO `t_position` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `order_type_id`, `city_id`, `region_id`, `place_id`) VALUES (2, '2023-11-17 20:07:35', '2023-11-17 20:07:35', NULL, '九龙-尖沙咀-A', 1, 2, 2, 1011);
INSERT INTO `t_position` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `order_type_id`, `city_id`, `region_id`, `place_id`) VALUES (3, '2023-11-18 15:21:49', '2023-11-18 15:21:49', NULL, '宝安区-万象城-001', 2, 1, 1, 11);
INSERT INTO `t_position` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `order_type_id`, `city_id`, `region_id`, `place_id`) VALUES (4, '2023-11-18 19:57:07', '2023-11-18 19:57:07', NULL, '宝安区-前海湾-001', 3, 1, 1, 21);
COMMIT;

-- ----------------------------
-- Table structure for t_region
-- ----------------------------
DROP TABLE IF EXISTS `t_region`;
CREATE TABLE `t_region` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(255) NOT NULL,
  `city_id` bigint(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of t_region
-- ----------------------------
BEGIN;
INSERT INTO `t_region` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `city_id`) VALUES (1, '2023-11-17 20:07:35', '2023-11-17 20:07:35', NULL, '南山区', 1);
INSERT INTO `t_region` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `city_id`) VALUES (2, '2023-11-17 20:07:35', '2023-11-17 20:07:35', NULL, '九龙', 2);
COMMIT;

-- ----------------------------
-- Table structure for t_user
-- ----------------------------
DROP TABLE IF EXISTS `t_user`;
CREATE TABLE `t_user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `phone` varchar(50) NOT NULL,
  `username` varchar(80) NOT NULL,
  `password` varchar(80) NOT NULL,
  `avatar` varchar(2046) NOT NULL,
  `bio` varchar(2046) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_t_user_phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of t_user
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
