
DROP DATABASE IF EXISTS `mall`;
CREATE DATABASE `mall`;

USE `mall`;

DROP TABLE IF EXISTS `member`;
CREATE TABLE `member` (
    `id` CHAR(36) NOT NULL COMMENT 'PK',
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
    `id` CHAR(36) NOT NULL COMMENT 'PK',
    `member_id` CHAR(36) NOT NULL COMMENT '會員id',
    `product_id` CHAR(36) NOT NULL COMMENT '產品id',
    `quantity` INT NOT NULL COMMENT '數量',
    `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '建立時間',
    `updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '更新時間',
    PRIMARY KEY (`id`),
    KEY `idx_member_id` (`member_id`)
) COMMENT = '會員購物車';

DROP TABLE IF EXISTS `product`;
CREATE TABLE IF NOT EXISTS `product` (
    `id` CHAR(36) NOT NULL COMMENT 'PK',
    `name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '產品名稱',
    `amount` INT NOT NULL  DEFAULT 0 COMMENT '產品價格',
    `quantity` INT NOT NULL COMMENT '數量',
    `image` VARCHAR(255) NOT NULL COMMENT '圖片路徑',
    `status` TINYINT NOT NULL COMMENT '狀態 0： "無庫存 1 ：尚有庫存',
    `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '建立時間',
    `updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '更新時間',
    PRIMARY KEY (`id`)
) COMMENT = '產品';

DROP TABLE IF EXISTS `transaction`;
CREATE TABLE IF NOT EXISTS `transaction` (
    `id` CHAR(36) NOT NULL COMMENT 'PK',
    `member_id` CHAR(36) NOT NULL COMMENT '會員id',
    `payment_number` VARCHAR(255) NOT NULL COMMENT  '卡號 結帳號碼',
    `amount` INT NOT NULL COMMENT '總價',
    `status` TINYINT NOT NULL COMMENT '狀態 0, : 待處理,1 : 成功,2 : 失敗',
    `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '建立時間',
    `updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '更新時間',
    PRIMARY KEY (`id`),
    KEY `idx_member_id` (`member_id`)
) COMMENT = '購物訂單';

DROP TABLE IF EXISTS `transaction_item`;
CREATE TABLE IF NOT EXISTS `transaction_item` (
    `id` INT NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `transaction_id` CHAR(36) NOT NULL COMMENT '產品id',
    `name` VARCHAR(255) NOT NULL COMMENT  '產品名稱',
    `amount` INT NOT NULL COMMENT '產品價格',
    `quantity` INT NOT NULL COMMENT '數量',
    `image` VARCHAR(255) NOT NULL COMMENT '圖片路徑',
    PRIMARY KEY (`id`),
    KEY `idx_transaction_id` (`transaction_id`)
) COMMENT = '購物訂單物品';
