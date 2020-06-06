/*
 Navicat Premium Data Transfer

 Source Server         : zhiliao
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : localhost:3306
 Source Schema         : info

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 17/05/2020 21:03:19
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for assets
-- ----------------------------
DROP TABLE IF EXISTS `assets`;
CREATE TABLE `assets`  (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `assetsName` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `assetsAddress` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `execStatus` int(1) NULL DEFAULT 0,
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 358 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '资产' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for buginfo
-- ----------------------------
DROP TABLE IF EXISTS `buginfo`;
CREATE TABLE `buginfo`  (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `ipdomain` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `urladdress` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `bugname` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `bugpoc` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `startTime` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `endTime` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1245 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for buginfobase
-- ----------------------------
DROP TABLE IF EXISTS `buginfobase`;
CREATE TABLE `buginfobase`  (
  `assetsName` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `assetsAddress` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `startTime` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `endTime` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL
) ENGINE = MyISAM CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pocinfo
-- ----------------------------
DROP TABLE IF EXISTS `pocinfo`;
CREATE TABLE `pocinfo`  (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `pocname` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '插件名称',
  `pocclass` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '插件类别',
  `poccontent` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '插件描述',
  `pocvulnreport` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '漏洞危害',
  `pocfilePath` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '文件路径',
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 27 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for primaryinfocol
-- ----------------------------
DROP TABLE IF EXISTS `primaryinfocol`;
CREATE TABLE `primaryinfocol`  (
  `assetsName` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `assetsAddress` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `ipdomain` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `port` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `urladdress` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `dir` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `startTime` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `endTime` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL
) ENGINE = MyISAM CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for singleonbuginfo
-- ----------------------------
DROP TABLE IF EXISTS `singleonbuginfo`;
CREATE TABLE `singleonbuginfo`  (
  `pluginname` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `pluginpath` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `urladdress` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `returninfo` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL
) ENGINE = MyISAM CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`Id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
