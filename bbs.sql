CREATE DATABASE IF NOT EXISTS `bbs_db`;

USE `bbs_db`;

CREATE TABLE IF NOT EXISTS `t_user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(31) NOT NULL DEFAULT '' COMMENT '姓名',
  `username` varchar(31) NOT NULL COMMENT '用户名',
  `email` varchar(63) NOT NULL DEFAULT '' COMMENT '邮箱',
  `password` varchar(127) NOT NULL DEFAULT '' COMMENT '密码',
  `avatar` varchar(127) NOT NULL DEFAULT '' COMMENT '头像',
  `status` tinyint unsigned NOT NULL DEFAULT 0 COMMENT '状态。0-默认；1-禁用；',
  `roles` text NOT NULL COMMENT '角色',
  `type` tinyint unsigned NOT NULL DEFAULT 0 COMMENT '类型。0-普通；1-公众号；',
  `description` text NOT NULL COMMENT '签名，独白',
  `city` varchar(10) NOT NULL DEFAULT '' COMMENT '居住地',
  `company` varchar(63) NOT NULL DEFAULT '' COMMENT '公司',
  `createTime` bigint NOT NULL COMMENT '创建时间',
  `updateTime` bigint NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY (`username`),
  UNIQUE KEY (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `t_third_user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `userId` int unsigned NOT NULL COMMENT '本站用户 ID',
  `thirdId` int unsigned NOT NULL COMMENT '第三方用户 ID',
  `type` tinyint NOT NULL DEFAULT 0 COMMENT '类型。0-github；',
  `name` varchar(31) NOT NULL DEFAULT '' COMMENT '姓名',
  `username` varchar(31) NOT NULL DEFAULT '' COMMENT '第三方用户名',
  `email` varchar(31) NOT NULL DEFAULT '' COMMENT '邮箱',
  `avatar` varchar(127) NOT NULL DEFAULT '' COMMENT '头像',
  `url` varchar(127) NOT NULL DEFAULT '' COMMENT '第三方用户接口地址',
  `htmlUrl` varchar(127) NOT NULL DEFAULT '' COMMENT 'github 地址',
  `createTime` bigint NOT NULL COMMENT '创建时间',
  `updateTime` bigint NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE uniq_user_type (`username`, `type`),
  KEY index_uid (`userId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `t_category` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(31) NOT NULL DEFAULT '' COMMENT '分类名',
  `description` text NOT NULL COMMENT '分类描述',
  `status` tinyint unsigned NOT NULL DEFAULT 0 COMMENT '状态。0-默认；1-禁用；',
  `createTime` bigint NOT NULL COMMENT '创建时间',
  `updateTime` bigint NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `t_tag` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(31) NOT NULL DEFAULT '' COMMENT '标签名',
  `description` text NOT NULL COMMENT '标签描述',
  `status` tinyint unsigned NOT NULL DEFAULT 0 COMMENT '状态。0-默认；1-禁用；',
  `createTime` bigint NOT NULL COMMENT '创建时间',
  `updateTime` bigint NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `t_article` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `categoryId` int unsigned NOT NULL COMMENT '分类 ID',
  `userId` int unsigned NOT NULL COMMENT '用户 ID',
  `title` varchar(127) NOT NULL DEFAULT '' COMMENT '标题',
  `summary` text NOT NULL COMMENT '概要',
  `content` longtext NOT NULL COMMENT '内容',
  `contentType` varchar(31) NOT NULL COMMENT '类型',
  `status` tinyint unsigned NOT NULL DEFAULT 0 COMMENT '状态。0-发布；1-删除；2-草稿；',
  `type` tinyint unsigned NOT NULL DEFAULT 0 COMMENT '状态。0-原创；1-分享；',
  `sourceUrl` varchar(127) NOT NULL DEFAULT '' COMMENT '原文链接',
  `createTime` bigint NOT NULL COMMENT '创建时间',
  `updateTime` bigint NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY (`title`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `t_article_tag` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `articleId` int unsigned NOT NULL COMMENT '文章 ID',
  `tagId` int unsigned NOT NULL COMMENT '标签 ID',
  `createTime` bigint NOT NULL COMMENT '创建时间',
  `updateTime` bigint NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY (`articleId`,`tagId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `t_comment` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `userId` int unsigned NOT NULL COMMENT '用户 ID',
  `entityType` varchar(31) NOT NULL COMMENT '评论实体类型',
  `entityId` int unsigned NOT NULL COMMENT '评论实体 ID',
  `quoteId` int unsigned NOT NULL COMMENT '引用的评论编号',
  `content` longtext NOT NULL COMMENT '内容',
  `status` tinyint unsigned NOT NULL DEFAULT 0 COMMENT '状态。0-发布；1-删除',
  `createTime` bigint NOT NULL COMMENT '创建时间',
  `updateTime` bigint NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `t_favorite` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `userId` int unsigned NOT NULL COMMENT '用户 ID',
  `entityType` varchar(31) NOT NULL COMMENT '实体类型',
  `entityId` int unsigned NOT NULL COMMENT '实体 ID',
  `createTime` bigint NOT NULL COMMENT '创建时间',
  `updateTime` bigint NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `t_topic` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `userId` int unsigned NOT NULL COMMENT '用户 ID',
  `title` varchar(127) NOT NULL DEFAULT '' COMMENT '标题',
  `content` longtext NOT NULL COMMENT '内容',
  `status` tinyint unsigned NOT NULL DEFAULT 0 COMMENT '状态。0-发布；1-删除；2-草稿；',
  `viewCount` int unsigned NOT NULL DEFAULT 0 COMMENT '查看人数',
  `lastCommentTime` int unsigned NOT NULL COMMENT '最后回复时间',
  `createTime` bigint NOT NULL COMMENT '创建时间',
  `updateTime` bigint NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY (`title`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `t_topic_tag` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `topicId` int unsigned NOT NULL COMMENT '话题 ID',
  `tagId` int unsigned NOT NULL COMMENT '标签 ID',
  `createTime` bigint NOT NULL COMMENT '创建时间',
  `updateTime` bigint NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY (`topicId`,`tagId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `t_message` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `fromId` int unsigned NOT NULL COMMENT '消息发送人',
  `userId` int unsigned NOT NULL COMMENT '消息接收人',
  `content` text NOT NULL COMMENT '内容',
  `quoteContent` text NOT NULL COMMENT '引用内容',
  `extraData` text NOT NULL COMMENT '扩展数据',
  `status` tinyint unsigned NOT NULL DEFAULT 0 COMMENT '状态。0-未读；1-已读；',
  `type` tinyint unsigned NOT NULL DEFAULT 0 COMMENT '类型。0-评论；1-系统通知；',
  `createTime` bigint NOT NULL COMMENT '创建时间',
  `updateTime` bigint NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY (`fromId`),
  KEY (`userId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `t_system_config` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `key` varchar(31) NOT NULL DEFAULT '' COMMENT '配置',
  `value` text NOT NULL COMMENT '配置值',
  `name` varchar(31) NOT NULL COMMENT '配置名称',
  `description` varchar(127) NOT NULL COMMENT '配置描述',
  `createTime` bigint NOT NULL COMMENT '创建时间',
  `updateTime` bigint NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY (`key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;