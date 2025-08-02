/*
 Navicat Premium Dump SQL

 Source Server         : MySQL本地测试
 Source Server Type    : MySQL
 Source Server Version : 50744 (5.7.44)
 Source Host           : localhost:3306
 Source Schema         : wallet-example

 Target Server Type    : MySQL
 Target Server Version : 50744 (5.7.44)
 File Encoding         : 65001

 Date: 02/08/2025 17:27:16
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for coin
-- ----------------------------
DROP TABLE IF EXISTS `coin`;
CREATE TABLE `coin`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `symbol` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '币种符号',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '币种名称',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '币种图标',
  `usdt_price` decimal(26, 8) UNSIGNED NOT NULL COMMENT 'USDT价格',
  `is_auto_price` tinyint(3) UNSIGNED NOT NULL COMMENT '是否自动获取价格（0：否，1：是）',
  `precision` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '币种精度',
  `is_transfer` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否可划转（0：否，1：手动，2：自动）',
  `transfer_rate` decimal(3, 2) NULL DEFAULT NULL COMMENT '划转费率',
  `min_transfer_fee` decimal(26, 18) NULL DEFAULT NULL COMMENT '最低划转费用',
  `min_transfer` decimal(26, 18) NULL DEFAULT NULL COMMENT '最低划转',
  `max_transfer` decimal(26, 18) NULL DEFAULT NULL COMMENT '最大划转',
  `sort` int(10) NOT NULL DEFAULT 0 COMMENT '排序（升序）',
  `status` tinyint(3) UNSIGNED NOT NULL COMMENT '状态（0：禁用，1：正常）',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `modified_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_symbol`(`symbol`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '币种表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of coin
-- ----------------------------
INSERT INTO `coin` VALUES (1, 'USDT', 'USDT', 'https://cdn.jsdelivr.net/gh/chainres/coin/usdt.png', 1.00000000, 0, 6, 2, 0.00, 0.000000000000000000, 0.000000000000000000, 0.000000000000000000, 2, 1, '2021-11-28 22:10:42', '2023-04-21 13:45:03');
INSERT INTO `coin` VALUES (2, 'USDC', 'USDC', 'https://cdn.jsdelivr.net/gh/chainres/coin/btc.png', 1.00000000, 0, 6, 2, 0.00, 0.000000000000000000, 0.000000000000000000, 0.000000000000000000, 3, 1, '2021-12-06 08:58:20', '2023-04-25 23:16:02');

-- ----------------------------
-- Table structure for coin_conf
-- ----------------------------
DROP TABLE IF EXISTS `coin_conf`;
CREATE TABLE `coin_conf`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `coin_symbol` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '币种符号',
  `network_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '网络名称',
  `decimals` int(11) UNSIGNED NOT NULL COMMENT '币种精度',
  `is_withdraw` tinyint(1) NOT NULL COMMENT '是否可提现（0：否，1：手动，2：自动）',
  `withdraw_auto` decimal(26, 18) UNSIGNED NULL DEFAULT 0.000000000000000000 COMMENT '提现自动转的阈值',
  `withdraw_rate` decimal(3, 2) NULL DEFAULT NULL COMMENT '提现费率',
  `min_withdraw_fee` decimal(26, 18) NULL DEFAULT NULL COMMENT '最低提现费用',
  `min_withdraw` decimal(26, 18) NULL DEFAULT NULL COMMENT '最低提现',
  `max_withdraw` decimal(26, 18) NULL DEFAULT NULL COMMENT '最大提现',
  `withdraw_private_key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '提现私钥',
  `min_recharge` decimal(26, 18) UNSIGNED NULL DEFAULT 0.000000000000000000 COMMENT '最小充值',
  `is_recharge` tinyint(1) NOT NULL COMMENT '是否可充值（0：否，1：自动）',
  `recharge_confirm` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '充值确认数',
  `withdraw_confirm` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '提现确认数',
  `sort` int(10) NOT NULL DEFAULT 0 COMMENT '排序（升序）',
  `status` tinyint(3) UNSIGNED NOT NULL COMMENT '状态（0：禁用，1：正常）',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `modified_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_cid_nid`(`coin_symbol`, `network_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '币种配置表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of coin_conf
-- ----------------------------
INSERT INTO `coin_conf` VALUES (1, 'USDT', 'BEP20', 18, 1, 0.000000000000000000, 0.05, 0.100000000000000000, 10.000000000000000000, 10000.000000000000000000, 'LIQBJHq6DsZQAv/F8PztIw==', 0.001000000000000000, 1, 4, 5, 0, 1, '2022-08-19 18:18:55', '2023-04-27 17:00:04');
INSERT INTO `coin_conf` VALUES (2, 'BTC', 'BEP20', 18, 1, 0.000000000000000000, 0.05, 0.000000000000000000, 0.000000000000000000, 0.000000000000000000, 'bd/Idbun+U9eO3rKolAsaA==', 5.000000000000000000, 1, 2, 10, 0, 1, '2023-04-25 22:59:49', '2023-04-27 17:00:12');
INSERT INTO `coin_conf` VALUES (3, 'USDT', 'Polygon', 18, 1, 0.000000000000000000, 0.05, 0.100000000000000000, 10.000000000000000000, 10000.000000000000000000, '', 0.001000000000000000, 1, 4, 5, 0, 1, '2022-08-19 18:18:55', '2023-04-27 17:00:04');
INSERT INTO `coin_conf` VALUES (4, 'USDT', 'ERC20', 6, 1, 0.000000000000000000, 0.05, 0.100000000000000000, 10.000000000000000000, 10000.000000000000000000, '', 0.001000000000000000, 1, 4, 5, 0, 1, '2022-08-19 18:18:55', '2023-04-27 17:00:04');
INSERT INTO `coin_conf` VALUES (5, 'USDT', 'TRC20', 6, 1, 0.000000000000000000, 0.05, 0.100000000000000000, 10.000000000000000000, 10000.000000000000000000, '', 0.001000000000000000, 1, 4, 5, 0, 1, '2022-08-19 18:18:55', '2023-04-27 17:00:04');
INSERT INTO `coin_conf` VALUES (6, 'USDC', 'Polygon-TEST', 6, 2, 0.000000000000000000, 0.05, 0.100000000000000000, 1.000000000000000000, 10000.000000000000000000, '', 0.001000000000000000, 1, 4, 5, 99, 1, '2022-08-19 18:18:55', '2023-04-27 17:00:04');

-- ----------------------------
-- Table structure for coin_network
-- ----------------------------
DROP TABLE IF EXISTS `coin_network`;
CREATE TABLE `coin_network`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID自增',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '网络名称',
  `sort` int(10) NOT NULL DEFAULT 0 COMMENT '排序（升序）',
  `status` tinyint(3) UNSIGNED NOT NULL COMMENT '状态（0：禁用，1：正常）',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `modified_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '币种网络表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of coin_network
-- ----------------------------
INSERT INTO `coin_network` VALUES (1, 'Polygon', 2, 1, '2021-11-28 22:17:24', '2022-08-20 12:59:04');
INSERT INTO `coin_network` VALUES (3, 'BEP20', 2, 1, '2021-11-28 22:17:24', '2022-08-20 12:59:43');
INSERT INTO `coin_network` VALUES (6, 'ERC20', 0, 1, '2023-04-08 22:15:40', '2023-04-08 22:15:40');
INSERT INTO `coin_network` VALUES (7, 'TRC20', 0, 1, '2023-10-11 01:35:02', '2023-10-11 01:35:04');

-- ----------------------------
-- Table structure for member
-- ----------------------------
DROP TABLE IF EXISTS `member`;
CREATE TABLE `member`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `pid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '直推上级',
  `tel` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '手机号（区号用下划线隔开）',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '邮箱号',
  `nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户昵称',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户头像',
  `pwd` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '密码',
  `pay_pwd` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '支付密码',
  `google_key` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '谷歌验证key',
  `fishing_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '防钓鱼码',
  `last_update_safe` datetime NULL DEFAULT NULL COMMENT '最后修改安全项的时间',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '个性签名',
  `status` tinyint(3) UNSIGNED NULL DEFAULT 1 COMMENT '状态（0：禁用，1：正常）',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `modified_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_tel`(`tel`) USING BTREE,
  UNIQUE INDEX `uk_email`(`email`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of member
-- ----------------------------
INSERT INTO `member` VALUES (3, 0, '86_15111111111', '44@qq.com', '86_****1111', '', '6512bd43d9caa6e02c990b0a82652dca', '', '', '', '2023-09-03 23:07:07', '', 1, '2023-09-03 18:32:28', '2023-09-03 18:32:28');
INSERT INTO `member` VALUES (4, 0, '86_15044444444', '33@qq.com', '11@qq.com', '', '6512bd43d9caa6e02c990b0a82652dca', '', 'Q5USNVIZFKLUZXQVBXU3IEM3T6PHBBMX', '88888', '2023-09-03 23:02:18', '', 1, '2023-09-03 18:32:46', '2023-09-03 18:32:46');
INSERT INTO `member` VALUES (5, 0, '11', NULL, '11', '', '6512bd43d9caa6e02c990b0a82652dca', '', '', '', '2025-08-01 11:26:09', '', 1, '2025-08-02 11:26:09', '2025-08-02 11:26:09');
INSERT INTO `member` VALUES (6, 0, '113', NULL, '113', '', '6512bd43d9caa6e02c990b0a82652dca', '', '', '', '2025-08-01 11:43:30', '', 1, '2025-08-02 11:43:30', '2025-08-02 11:43:30');
INSERT INTO `member` VALUES (7, 0, '1134', NULL, '1134', '', '6512bd43d9caa6e02c990b0a82652dca', '', '', '', '2025-08-01 11:43:55', '', 1, '2025-08-02 11:43:55', '2025-08-02 11:43:55');
INSERT INTO `member` VALUES (8, 0, '111', NULL, '111', '', '6512bd43d9caa6e02c990b0a82652dca', '', '', '', '2025-08-01 12:00:16', '', 1, '2025-08-02 12:00:16', '2025-08-02 12:00:16');

-- ----------------------------
-- Table structure for member_bill
-- ----------------------------
DROP TABLE IF EXISTS `member_bill`;
CREATE TABLE `member_bill`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `member_id` bigint(20) NOT NULL COMMENT '用户ID',
  `coin_symbol` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '币种符号',
  `from_account` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '来源账户',
  `to_account` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '接收账户',
  `mode` tinyint(3) UNSIGNED NOT NULL COMMENT '类型（0：划转，1：收入，2：支出）',
  `business_type` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '业务类型',
  `business_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '业务ID',
  `amount` decimal(48, 18) UNSIGNED NULL DEFAULT NULL COMMENT '数量',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `modified_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户账单表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of member_bill
-- ----------------------------
INSERT INTO `member_bill` VALUES (1, 5, 'USDC', '', '', 1, 'recharge', '2', 2.000000000000000000, '', '2025-08-02 16:57:19', '2025-08-02 16:57:19');
INSERT INTO `member_bill` VALUES (2, 5, 'USDC', '', '', 2, 'withdraw', '285', 1.000000000000000000, '', '2025-08-02 17:09:58', '2025-08-02 17:09:58');
INSERT INTO `member_bill` VALUES (3, 5, 'USDC', '', '', 2, 'withdraw', '286', 1.000000000000000000, '', '2025-08-02 17:13:15', '2025-08-02 17:13:15');

-- ----------------------------
-- Table structure for member_coin
-- ----------------------------
DROP TABLE IF EXISTS `member_coin`;
CREATE TABLE `member_coin`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `member_id` bigint(20) UNSIGNED NOT NULL COMMENT '用户ID',
  `coin_symbol` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '币种符号',
  `balance` decimal(48, 18) UNSIGNED NOT NULL COMMENT '可用余额',
  `frozen_balance` decimal(48, 18) NOT NULL COMMENT '冻结余额',
  `virtual_balance` decimal(48, 18) UNSIGNED NOT NULL DEFAULT 0.000000000000000000 COMMENT '虚拟余额',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `modified_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_mid_cid`(`member_id`, `coin_symbol`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户钱包表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of member_coin
-- ----------------------------
INSERT INTO `member_coin` VALUES (13, 5, 'USDC', 0.000000000000000000, 0.000000000000000000, 0.000000000000000000, '2025-08-02 16:40:00', '2025-08-02 16:40:00');
INSERT INTO `member_coin` VALUES (14, 5, 'USDT', 0.000000000000000000, 0.000000000000000000, 0.000000000000000000, '2025-08-02 16:40:00', '2025-08-02 16:40:00');

-- ----------------------------
-- Table structure for recharge
-- ----------------------------
DROP TABLE IF EXISTS `recharge`;
CREATE TABLE `recharge`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `business_id` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '业务ID',
  `member_id` bigint(20) UNSIGNED NOT NULL COMMENT '用户ID',
  `network_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '网络',
  `coin_symbol` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '币种符号',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '充值地址',
  `amount` decimal(48, 18) UNSIGNED NOT NULL DEFAULT 0.000000000000000000 COMMENT '数量',
  `max_block_high` bigint(20) NULL DEFAULT 0 COMMENT '最大区块高度',
  `block_high` bigint(20) UNSIGNED NULL DEFAULT 0 COMMENT '区块高度',
  `txid` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '区块交易哈希',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `status` tinyint(3) UNSIGNED NOT NULL COMMENT '状态（0：区块确认中，1：充值到账，2：区块确认失败）',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `modified_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_business_id`(`business_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '充值表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for recharge_address
-- ----------------------------
DROP TABLE IF EXISTS `recharge_address`;
CREATE TABLE `recharge_address`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `key` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '随机key',
  `member_id` bigint(20) NOT NULL COMMENT '用户ID',
  `network_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '网络',
  `coin_symbol` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '币种符号',
  `address` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '地址',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `modified_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_mid_nn_cs`(`member_id`, `network_name`, `coin_symbol`) USING BTREE,
  UNIQUE INDEX `uk_nn_cs_address`(`network_name`, `coin_symbol`, `address`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户充值的地址表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for withdraw
-- ----------------------------
DROP TABLE IF EXISTS `withdraw`;
CREATE TABLE `withdraw`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `member_id` bigint(20) UNSIGNED NOT NULL COMMENT '用户ID',
  `network_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '网络',
  `coin_symbol` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '币种符号',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '充值地址',
  `amount` decimal(48, 18) UNSIGNED NOT NULL DEFAULT 0.000000000000000000 COMMENT '数量',
  `fee` decimal(48, 18) UNSIGNED NOT NULL DEFAULT 0.000000000000000000 COMMENT '手续费',
  `block_high` bigint(20) UNSIGNED NULL DEFAULT 0 COMMENT '区块高度',
  `txid` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '区块交易哈希',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `status` tinyint(3) UNSIGNED NOT NULL COMMENT '状态（0：审核中，1：审核通过，2：审核不通过，3：链上打包中，4：提币成功，5：提币失败，6：手动成功）',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `modified_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 287 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '提现表' ROW_FORMAT = DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;
