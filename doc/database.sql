SET NAMES UTF8;
  create database if not exists cms default charset utf8 collate utf8_general_ci;
  grant select,update,delete,insert,alter,create,drop on cms.* to "cms"@"%" identified by "cms";
  grant select,update,delete,insert,alter,create,drop on cms.* to "cms"@"localhost" identified by "cms";
USE cms;


/*
Navicat MySQL Data Transfer

Source Server         : root
Source Server Version : 50629
Source Host           : localhost:3306
Source Database       : cms

Target Server Type    : MYSQL
Target Server Version : 50629
File Encoding         : 65001

Date: 2016-09-04 17:49:48
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for t_user
-- ----------------------------
DROP TABLE IF EXISTS `t_user`;
CREATE TABLE `t_user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `mail` varchar(255) NOT NULL DEFAULT '',
  `realname` varchar(255) NOT NULL DEFAULT '',
  `phone` varchar(255) NOT NULL DEFAULT '',
  `department` varchar(255) NOT NULL DEFAULT '',
  `passwd` varchar(255) NOT NULL DEFAULT '',
  `created` datetime NOT NULL,
  `updated` datetime NOT NULL,
  `is_del` tinyint(1) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of t_user
-- ----------------------------
INSERT INTO `t_user` VALUES ('1', 'admin', 'admin@admin.com', 'admin', 'admin', '系统', 'SddggEssSsmtsJgFUdFgInnFGsEfgBs2Sd==', NOW(), NOW(), '0');
INSERT INTO `t_user` VALUES ('2', 'test', 'test@test.com', 'test', 'test', 'test', 'SddggEssSsmt1dDGdsdgsdGfdDgRdfgsDd==', NOW(), NOW(), '0');

-- ----------------------------
-- Table structure for t_user_group
-- ----------------------------
DROP TABLE IF EXISTS `t_user_group`;
CREATE TABLE `t_user_group` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `group_name` varchar(255) NOT NULL DEFAULT '',
  `remarks` varchar(255) NOT NULL DEFAULT '',
  `created` datetime NOT NULL,
  `updated` datetime NOT NULL,
  `is_del` tinyint(1) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of t_user_group
-- ----------------------------
INSERT INTO `t_user_group` VALUES ('1', '超级管理员', '超级管理员', NOW(), NOW(), '0');
INSERT INTO `t_user_group` VALUES ('2', 'test', 'test', NOW(), NOW(), '0');
INSERT INTO `t_user_group` VALUES ('3', 'test2', 'test2', NOW(), NOW(), '1');

-- ----------------------------
-- Table structure for t_group_role_rel
-- ----------------------------
DROP TABLE IF EXISTS `t_group_role_rel`;
CREATE TABLE `t_group_role_rel` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `group_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `role_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `is_del` tinyint(1) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=310 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of t_group_role_rel
-- ----------------------------
INSERT INTO `t_group_role_rel` VALUES ('291', '2', '1', '0');
INSERT INTO `t_group_role_rel` VALUES ('292', '2', '18', '0');
INSERT INTO `t_group_role_rel` VALUES ('293', '2', '19', '0');
INSERT INTO `t_group_role_rel` VALUES ('294', '2', '20', '0');
INSERT INTO `t_group_role_rel` VALUES ('295', '2', '21', '0');
INSERT INTO `t_group_role_rel` VALUES ('296', '2', '2', '0');
INSERT INTO `t_group_role_rel` VALUES ('297', '2', '3', '0');
INSERT INTO `t_group_role_rel` VALUES ('298', '2', '6', '0');
INSERT INTO `t_group_role_rel` VALUES ('299', '2', '13', '0');
INSERT INTO `t_group_role_rel` VALUES ('300', '2', '15', '0');
INSERT INTO `t_group_role_rel` VALUES ('301', '2', '16', '0');
INSERT INTO `t_group_role_rel` VALUES ('302', '2', '4', '0');
INSERT INTO `t_group_role_rel` VALUES ('303', '2', '25', '0');
INSERT INTO `t_group_role_rel` VALUES ('304', '2', '26', '0');
INSERT INTO `t_group_role_rel` VALUES ('305', '2', '27', '0');
INSERT INTO `t_group_role_rel` VALUES ('306', '2', '28', '0');
INSERT INTO `t_group_role_rel` VALUES ('307', '2', '29', '0');
INSERT INTO `t_group_role_rel` VALUES ('308', '2', '30', '0');
INSERT INTO `t_group_role_rel` VALUES ('309', '2', '42', '0');

-- ----------------------------
-- Table structure for t_role
-- ----------------------------
DROP TABLE IF EXISTS `t_role`;
CREATE TABLE `t_role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `pid` int(11) unsigned DEFAULT '0',
  `name` varchar(255) NOT NULL DEFAULT '',
  `role_url` varchar(255) NOT NULL DEFAULT '',
  `is_menu` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `remarks` varchar(255) NOT NULL DEFAULT '',
  `module` varchar(50) NOT NULL DEFAULT '',
  `action` varchar(50) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=43 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of t_role
-- ----------------------------
INSERT INTO `t_role` VALUES ('0', null, 'Root', '3123', '1', '根节点', '', '');
INSERT INTO `t_role` VALUES ('1', '0', '公共权限', '', '1', '公共权限 所有账号都应该有', '', '');
INSERT INTO `t_role` VALUES ('2', '0', '账号管理', '', '0', '账号管理目录', '', '');
INSERT INTO `t_role` VALUES ('3', '2', '管理员管理', 'user/listview', '0', '', 'UserController', 'ListView');
INSERT INTO `t_role` VALUES ('4', '2', '管理员组管理', 'usergroup/listview', '0', '', 'UserGroupController', 'ListView');
INSERT INTO `t_role` VALUES ('5', '2', '权限管理', 'role/listview', '0', '', 'RoleController', 'ListView');
INSERT INTO `t_role` VALUES ('6', '3', '获取管理员列表', 'user/listview', '1', '', 'UserController', 'GridList');
INSERT INTO `t_role` VALUES ('13', '3', '查看所有管理员', 'user/gridlist', '1', '', 'UserController', 'GridList');
INSERT INTO `t_role` VALUES ('15', '3', '进入添加管理员', 'user/addview', '1', '进入添加管理员页面', 'UserController', 'AddView');
INSERT INTO `t_role` VALUES ('16', '3', '添加管理员', 'user/adduser', '1', '执行添加管理员操作', 'UserController', 'Adduser');
INSERT INTO `t_role` VALUES ('18', '1', '进入欢迎页', '/welcome', '1', '进入欢迎页', 'MainController', 'Welcome');
INSERT INTO `t_role` VALUES ('19', '1', '展示导航页面', '/leftMenu', '1', '展示导航页面', 'MainController', 'LeftMenu');
INSERT INTO `t_role` VALUES ('20', '1', '展示头部信息', '/header', '1', '展示头部信息', 'MainController', 'Header');
INSERT INTO `t_role` VALUES ('21', '1', '获取菜单数据', '/loadMenu', '1', '获取菜单数据', 'MainController', 'LoadMenu');
INSERT INTO `t_role` VALUES ('25', '4', '进入添加页面', '', '1', '进入添加页面', 'UserGroupController', 'AddView');
INSERT INTO `t_role` VALUES ('26', '4', '添加管理员组', '', '1', '添加管理员组', 'UserGroupController', 'AddUserGroup');
INSERT INTO `t_role` VALUES ('27', '4', '进入修改页面', '', '1', '进入修改页面', 'UserGroupController', 'UpdateView');
INSERT INTO `t_role` VALUES ('28', '4', '修改管理员组', '', '1', '修改管理员组', 'UserGroupController', 'ModifyUserGroup');
INSERT INTO `t_role` VALUES ('29', '4', '删除管理员组', '', '1', '删除管理员组', 'UserGroupController', 'Delete');
INSERT INTO `t_role` VALUES ('30', '4', '获取权限树', '', '1', '添加管理员组的时候展示权限树', 'UserGroupController', 'LoadTreeChecked');
INSERT INTO `t_role` VALUES ('32', '5', '查询', '', '1', '查询列表', 'RoleController', 'GridList');
INSERT INTO `t_role` VALUES ('33', '5', '加载左侧树', '', '1', '加载左侧树', 'RoleController', 'ListTree');
INSERT INTO `t_role` VALUES ('34', '5', '进入添加页面', '', '1', '进入添加页面', 'RoleController', 'AddView');
INSERT INTO `t_role` VALUES ('35', '5', '添加权限', '', '1', '添加权限', 'RoleController', 'AddRole');
INSERT INTO `t_role` VALUES ('36', '5', '进入修改页面', '', '1', '进入修改页面', 'RoleController', 'UpdateView');
INSERT INTO `t_role` VALUES ('37', '5', '修改权限', '', '1', '修改权限', 'RoleController', 'Modify');
INSERT INTO `t_role` VALUES ('38', '5', '删除权限', '', '1', '删除权限', 'RoleController', 'DeleteRole');
INSERT INTO `t_role` VALUES ('42', '4', '查询', '', '1', '查询列表', 'UserGroupController', 'GridList');

-- ----------------------------
-- Table structure for t_user_group_rel
-- ----------------------------
DROP TABLE IF EXISTS `t_user_group_rel`;
CREATE TABLE `t_user_group_rel` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `group_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `is_del` tinyint(1) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of t_user_group_rel
-- ----------------------------
INSERT INTO `t_user_group_rel` VALUES ('4', '1', '1', '0');
INSERT INTO `t_user_group_rel` VALUES ('7', '2', '2', '0');
