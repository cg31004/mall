DROP DATABASE IF EXISTS `mall`;
CREATE DATABASE `mall` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `mall`;

DROP TABLE IF EXISTS `member`;
CREATE TABLE `member` (
    `id` BIGINT NOT NULL COMMENT 'PK',
    `account` VARCHAR(50) NOT NULL COMMENT '帳號',
    `name` VARCHAR(50) NOT NULL COMMENT '姓名',
    `password` VARCHAR(64) NOT NULL COMMENT '密碼',
    `salt` VARCHAR(64) NOT NULL COMMENT '鹽',
    `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '建立時間',
    `updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '更新時間',
    PRIMARY KEY (`id`),
    KEY `idx_account` (`account`)
) COMMENT = '會員資料';

DROP TABLE IF EXISTS `member_chart`;
CREATE TABLE `member_chart` (
    `id` INT NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `member_id` BIGINT NOT NULL COMMENT '會員id',
    `product_id` BIGINT NOT NULL COMMENT '產品id',
    `quantity` INT NOT NULL COMMENT '數量',
    `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '建立時間',
    `updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '更新時間',
    PRIMARY KEY (`id`),
    KEY `idx_member_id` (`member_id`)
) COMMENT = '會員購物車';

DROP TABLE IF EXISTS `transaction`;
CREATE TABLE `transaction` (
    `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `member_id` BIGINT NOT NULL COMMENT '會員id',
    `payment_number` VARCHAR(200) NOT NULL COMMENT '付款資訊',
    `amount` INT NOT NULL COMMENT '總價',
    `discount` INT NOT NULL COMMENT '優惠',
    `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '建立時間',
    `updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '更新時間',
    PRIMARY KEY (`id`),
    KEY `idx_member_id` (`member_id`)
) COMMENT = '會員購物車';


DROP TABLE IF EXISTS `otp_personal`;
CREATE TABLE `otp_personal` (
    `id` VARCHAR(36) NOT NULL COMMENT 'PK',
    `security` VARCHAR(64) NOT NULL COMMENT 'otp security',
    `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '建立時間',
    `updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '更新時間',
    `operator_id` VARCHAR(36) NOT NULL COMMENT '操作者',
    PRIMARY KEY (`id`),
    KEY `idx_operator_id` (`operator_id`)
) COMMENT = '個人OTP';

DROP TABLE IF EXISTS `otp_group`;
CREATE TABLE `otp_group` (
    `id` VARCHAR(36) NOT NULL COMMENT 'PK',
    `name` VARCHAR(64) NOT NULL COMMENT '群組名稱',
    `security` VARCHAR(64) NOT NULL COMMENT 'otp security',
    `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '建立時間',
    `updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '更新時間',
    `operator_id` VARCHAR(36) NOT NULL COMMENT '操作者',
    PRIMARY KEY (`id`),
    KEY `idx_operator_id` (`operator_id`)
) COMMENT = '群組OTP';

-- init user
INSERT INTO `otp_personal` (`id`, `security`,`created_at`,`updated_at`,`operator_id`)
VALUES('for_admin_otp', 'Q2NRTYZBSBIXSLHCYNTL4X33GLKBJKO6', '2022-08-16 16:15:44.486', '2022-08-16 16:15:44.486', 'your daddy');

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
    `id` VARCHAR(32) NOT NULL COMMENT 'PK',
    `account` VARCHAR(50) NOT NULL COMMENT '登入帳號',
    `password` CHAR(128) NOT NULL COMMENT '密碼Hash',
    `name` VARCHAR(50) NOT NULL COMMENT '登入名稱',
    `otp_group_id` VARCHAR(36) DEFAULT NULL COMMENT 'group otp id',
    `otp_personal_id` VARCHAR(36) DEFAULT NULL COMMENT 'personal otp id',
    `role_id` VARCHAR(36) DEFAULT NULL COMMENT 'auth_role.id',
    `status` TINYINT DEFAULT '1' NOT NULL COMMENT '用戶啟用 (0: 禁用, 1: 啟用)',
    `lock_status` TINYINT DEFAULT '1' NOT NULL COMMENT '用户狀態 (0: 鎖定, 1: 未鎖定)',
    `last_login_date_time` DATETIME(6) DEFAULT NULL COMMENT '最後更新時間',
    `failed_times` TINYINT DEFAULT 0 COMMENT '錯誤登入次數',
    `failed_date_time` DATETIME(6) DEFAULT NULL COMMENT '錯誤登入時間',
    `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '建立時間',
    `created_id` VARCHAR(32) NOT NULL COMMENT '建立時間',
    `updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '更新時間',
    `updated_id` VARCHAR(32) NOT NULL COMMENT '更新時間',
    PRIMARY KEY (`id`),
    KEY `idx_account` (`account`),
    KEY `idx_status` (`status`),
    KEY `idx_lock_status` (`lock_status`)
) COMMENT = '後台使用者表';

INSERT INTO `user`(`account`,`name`, `otp_personal_id`, `otp_group_id`, `role_id`,`password`,`status`,`lock_status`,`last_login_date_time`,`failed_times`,`failed_date_time`,`created_at`,`created_id`,`updated_at`,`updated_id`,`id`)
VALUES('admin','admin', 'for_admin_otp', NULL, '15e7df5b-8aae-4b56-8800-40c9fe862425', 'd82494f05d6917ba02f7aaa29689ccb444bb73f20380876cb05d1f37537b7892',1,1,NULL,0,NULL,'2022-08-16 16:15:44.486','0','2022-08-16 16:15:44.486','0','1559453842737729536');


DROP TABLE IF EXISTS `brand`;
CREATE TABLE `brand`
(
    `code` VARCHAR(30) NOT NULL COMMENT '品牌代碼',
    `name` VARCHAR(50) NOT NULL COMMENT '品牌名稱',
    `sync_status` TINYINT NOT NULL DEFAULT 0 COMMENT '是否啟用同步',
    `status` TINYINT NOT NULL DEFAULT 0 COMMENT '是否啟用',
    `url` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '更新品牌url',
    `created_at` DATETIME NOT NULL COMMENT '建立時間',
    `created_id` VARCHAR(32) NOT NULL COMMENT '建立id',
    `updated_at` DATETIME NOT NULL COMMENT '更新時間',
    `updated_id` VARCHAR(32) NOT NULL COMMENT '更新id',
    PRIMARY KEY (`code`)
) COMMENT ='品牌列表';


DROP TABLE IF EXISTS `sync_log`;
CREATE TABLE `sync_log` (
    `id` VARCHAR(50) NOT NULL COMMENT 'ID',
    `type` TINYINT NOT NULL DEFAULT 0  COMMENT '數據類型(0:銀行管理)',
    `from` VARCHAR(50) NOT NULL COMMENT '同步品牌',
    `to` VARCHAR(50) NOT NULL COMMENT '接收品牌',
    `sync_type` TINYINT NOT NULL DEFAULT 0 COMMENT '同步類型(0:總控後台->品牌後台 1:品牌後台->總控後台)',
    `status` TINYINT NOT NULL DEFAULT 0 COMMENT '同步結果(0:失敗 1:成功 2:處理中)',
    `re_sync_status` TINYINT NOT NULL DEFAULT 0 COMMENT '重新同步狀態(0:未重新同步 1:已重新同步)',
    `created_at` DATETIME NOT NULL COMMENT '建立時間',
    `synced_at` DATETIME NOT NULL COMMENT  '同步時間',
    `note` VARCHAR(200) NOT NULL COMMENT '備註',
    PRIMARY KEY (`id`)
) COMMENT ='同步日誌列表';


DROP TABLE IF EXISTS `sync_bank`;
CREATE TABLE `sync_bank` (
    `sync_log_id` VARCHAR(50) NOT NULL COMMENT 'SyncLogID',
    `code` VARCHAR(20) NOT NULL COMMENT '銀行代碼',
    `field` LONGTEXT NOT NULL COMMENT '同步資訊'
) COMMENT ='銀行管理詳情';


DROP TABLE IF EXISTS `bank`;
CREATE TABLE `bank` (
    `code` VARCHAR(20) NOT NULL COMMENT '銀行代碼',
    `name` VARCHAR(50) NOT NULL COMMENT '銀行名稱',
    `url` VARCHAR(1024) NOT NULL COMMENT '銀行官方網址',
    `status` TINYINT NOT NULL DEFAULT 0 COMMENT '狀態(0:關閉 1:啟用)',
    `recommend` TINYINT NOT NULL DEFAULT 0 COMMENT '推薦(0:不推薦 1:推薦)',
    `sort` INT NOT NULL DEFAULT 0 COMMENT '排序',
    `type` TINYINT NOT NULL DEFAULT 0 COMMENT '1:銀行 2:三方支付 4:交易所',
    `imgpath` VARCHAR(100) NOT NULL COMMENT '銀行圖片路徑',
    `brand` VARCHAR(20) NOT NULL COMMENT '品牌代碼',
    `created_at` DATETIME NOT NULL COMMENT '建立時間',
    `created_id` VARCHAR(32) NOT NULL COMMENT '建立id',
    `updated_at` DATETIME NOT NULL COMMENT '更新時間',
    `updated_id` VARCHAR(32) NOT NULL COMMENT '更新id',
    PRIMARY KEY (`code`, `brand`)
) COMMENT ='銀行列表';

