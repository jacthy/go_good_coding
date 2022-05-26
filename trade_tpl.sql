/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50617
 Source Host           : localhost:3306
 Source Schema         : trade_tpl

 Target Server Type    : MySQL
 Target Server Version : 50617
 File Encoding         : 65001

 Date: 15/04/2022 09:36:17
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tt_bill_type
-- ----------------------------
DROP TABLE IF EXISTS `tt_bill_type`;
CREATE TABLE `tt_bill_type` (
  `bill_type_id` char(36) NOT NULL DEFAULT '' COMMENT '单据类型ID',
  `bill_type_name` varchar(50) NOT NULL DEFAULT '' COMMENT '单据类型名称',
  `bill_type_desc` varchar(100) NOT NULL DEFAULT '' COMMENT '单据类型描述',
  `is_bind_tpl` tinyint(2) NOT NULL DEFAULT '0' COMMENT '是否绑定模板，绑定模板时会在公共模块显示　0:否 1:是',
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modified_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `seq` int(11) NOT NULL DEFAULT '0' COMMENT '模板类型自定义参数总数记录',
  `categories` varchar(300) NOT NULL DEFAULT '' COMMENT '模板类型分组信息列表 json字符串',
  PRIMARY KEY (`bill_type_id`),
  KEY `IX_modified_on` (`modified_on`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='单据类型';

-- ----------------------------
-- Table structure for tt_param
-- ----------------------------
DROP TABLE IF EXISTS `tt_param`;
CREATE TABLE `tt_param` (
  `param_id` char(36) NOT NULL COMMENT '参数ID',
  `bill_type_id` varchar(36) NOT NULL DEFAULT '' COMMENT '单据类型ID',
  `param_name` varchar(50) NOT NULL DEFAULT '' COMMENT '参数名称',
  `param_field_name` varchar(50) NOT NULL DEFAULT '' COMMENT '参数对应字段名称',
  `param_type` tinyint(2) NOT NULL DEFAULT '0' COMMENT '参数类型，默认0，0文本，1金额，2编号（废弃），3电子签章位置',
  `param_scope` tinyint(2) NOT NULL DEFAULT '0' COMMENT '参数范围，默认0，0基础参数，1其他参数',
  `preview_value` varchar(50) NOT NULL DEFAULT '' COMMENT '预览值',
  `placeholder_count` int(11) NOT NULL DEFAULT '0' COMMENT '占位符宽度（字符数）',
  `is_thousands_format` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否千分位格式化,默认0，0否，1是',
  `sort_no` int(11) NOT NULL DEFAULT '0' COMMENT '排序字段',
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `cate` varchar(30) NOT NULL DEFAULT '' COMMENT '参数分组',
  `is_custom_param` tinyint(2) NOT NULL DEFAULT '0' COMMENT '参数类型:0 内置参数, 1 用户自定义参数',
  `is_export_in_regular_order` tinyint(2) NOT NULL DEFAULT '0' COMMENT '正式订单导出该参数;1：是；0：否；默认0',
  PRIMARY KEY (`param_id`),
  KEY `IX_bill_type_id` (`bill_type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='业务参数';

-- ----------------------------
-- Table structure for tt_tpl
-- ----------------------------
DROP TABLE IF EXISTS `tt_tpl`;
CREATE TABLE `tt_tpl` (
  `tpl_id` char(36) NOT NULL COMMENT '模板ID',
  `tpl_name` varchar(50) NOT NULL DEFAULT '' COMMENT '模板名称',
  `bill_type_id` varchar(36) NOT NULL DEFAULT '' COMMENT '单据类型ID',
  `is_used` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否被引用,0:未引用，1已被引用',
  `org_code` varchar(64) NOT NULL DEFAULT '' COMMENT '租户编码',
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modified_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`tpl_id`),
  KEY `IX_org_code` (`org_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='单据模板';

-- ----------------------------
-- Table structure for tt_tpl_extend
-- ----------------------------
DROP TABLE IF EXISTS `tt_tpl_extend`;
CREATE TABLE `tt_tpl_extend` (
  `tpl_id` char(36) NOT NULL COMMENT '模板ID',
  `tpl_content` mediumtext COMMENT '模板内容',
  `used_param_ids` varchar(5000) NOT NULL DEFAULT '' COMMENT '模板内容中引用的参数集合',
  `org_code` varchar(64) NOT NULL DEFAULT '' COMMENT '租户编码',
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modified_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`tpl_id`),
  KEY `IX_org_code` (`org_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='单据模板内容';

-- ----------------------------
-- Table structure for tt_tpl_preview
-- ----------------------------
DROP TABLE IF EXISTS `tt_tpl_preview`;
CREATE TABLE `tt_tpl_preview` (
  `tpl_id` char(36) NOT NULL COMMENT '模板ID',
  `tpl_preview_content` mediumtext COMMENT '模板内容',
  `org_code` varchar(64) NOT NULL DEFAULT '' COMMENT '租户编码',
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modified_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`tpl_id`),
  KEY `IX_org_code` (`org_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='单据模板内容预览';

-- ----------------------------
-- Table structure for tt_tpl_relate
-- ----------------------------
DROP TABLE IF EXISTS `tt_tpl_relate`;
CREATE TABLE `tt_tpl_relate` (
  `relate_id` char(36) NOT NULL COMMENT '关联ID',
  `tpl_id` char(36) NOT NULL COMMENT '关联的模板ID',
  `tpl_source` tinyint(2) NOT NULL DEFAULT '1' COMMENT '模板来源 1:云客模板 2:ERP模板',
  `org_code` varchar(64) NOT NULL DEFAULT '' COMMENT '租户编码',
  `app_code` varchar(20) NOT NULL DEFAULT '' COMMENT '应用编码',
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modified_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`relate_id`),
  KEY `IX_org_app_code` (`org_code`,`app_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='模板关联表';

-- ----------------------------
-- Table structure for tt_tpl_to_param
-- ----------------------------
DROP TABLE IF EXISTS `tt_tpl_to_param`;
CREATE TABLE `tt_tpl_to_param` (
  `tpl_id` char(36) NOT NULL COMMENT '模板ID',
  `param_id` varchar(36) NOT NULL COMMENT '参数ID',
  `preview_value` varchar(50) NOT NULL DEFAULT '' COMMENT '预览值',
  `placeholder_count` int(11) NOT NULL DEFAULT '0' COMMENT '占位符宽度',
  `is_thousands_format` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否千分位格式化,默认0，0否，1是',
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modified_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `is_export_in_regular_order` tinyint(2) DEFAULT '0' COMMENT '正式订单导出该参数;1：是；0：否；默认0',
  `sort_no` int(11) NOT NULL DEFAULT '0' COMMENT '排序字段',
  PRIMARY KEY (`tpl_id`,`param_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='单据模板参数（存储参数在模板层级的信息）';

-- ----------------------------
-- Table structure for tt_used_param
-- ----------------------------
DROP TABLE IF EXISTS `tt_used_param`;
CREATE TABLE `tt_used_param` (
  `tpl_id` char(36) NOT NULL COMMENT '模板ID',
  `param_id` char(36) NOT NULL COMMENT '参数ID,包括内置参数，模版类型自定义参数，模版自定义参数',
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modified_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  UNIQUE KEY `IX_tpl_param_id` (`tpl_id`,`param_id`),
  KEY `IX_tpl_id` (`tpl_id`),
  KEY `IX_param_id` (`param_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='模板使用的参数';

SET FOREIGN_KEY_CHECKS = 1;
