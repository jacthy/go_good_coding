/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : SQLite
 Source Server Version : 3030001
 Source Schema         : main

 Target Server Type    : SQLite
 Target Server Version : 3030001
 File Encoding         : 65001

 Date: 15/04/2022 09:33:42
*/

PRAGMA foreign_keys = false;

-- ----------------------------
-- Table structure for products
-- ----------------------------
DROP TABLE IF EXISTS "products";
CREATE TABLE `products` (`id` integer,`created_at` datetime,`updated_at` datetime,`deleted_at` datetime,`code` text,`price` integer,PRIMARY KEY (`id`));

-- ----------------------------
-- Indexes structure for table products
-- ----------------------------
CREATE INDEX "main"."idx_products_deleted_at"
ON "products" (
  "deleted_at" ASC
);

PRAGMA foreign_keys = true;
