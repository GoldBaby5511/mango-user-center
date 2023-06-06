CREATE TABLE `user_account` (
  `sys_id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `game_id` int unsigned NOT NULL DEFAULT '0' COMMENT '游戏ID',
  `nickname` varchar(32) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '昵称',
  `account` varchar(32) NOT NULL DEFAULT '' COMMENT '账号',
  `state` tinyint unsigned NOT NULL DEFAULT '3' COMMENT '状态 1正常2冻结3未验证',
  `sex` tinyint unsigned NOT NULL DEFAULT '3' COMMENT '性别 1男2女3未知',
  `age` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '年龄',
  `channel_id` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '注册主渠道',
  `site_id` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '注册子渠道',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  `passwd` varchar(64) NOT NULL DEFAULT '' COMMENT '密码',
  `salt` varchar(10) NOT NULL DEFAULT '' COMMENT '密码盐',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `change_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`sys_id`),
  UNIQUE KEY `uk_account` (`account`)
) ENGINE=InnoDB COMMENT '用户账号' DEFAULT CHARSET=utf8;

CREATE TABLE `user_captcha` (
  `sys_id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `account` varchar(32) NOT NULL DEFAULT '' COMMENT '账号',
  `captcha` char(6) NOT NULL DEFAULT '' COMMENT '验证码',
  `is_used` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否已使用',
  `category` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '类型 1注册2更改密码',
  `ip` int unsigned NOT NULL DEFAULT '0' COMMENT 'IP地址',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `change_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`sys_id`),
  KEY `idx_account_createtime` (`account`,`create_time`),
  KEY `idx_ip_createtime` (`ip`,`create_time`)
) ENGINE=InnoDB COMMENT '用户验证码' DEFAULT CHARSET=utf8;

CREATE TABLE `user_login` (
  `sys_id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `refresh_token` varchar(32) NOT NULL DEFAULT '' COMMENT '刷新token',
  `refresh_time` datetime DEFAULT NULL COMMENT '刷新时间',
  `device_os` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '设备系统 0未知1苹果2安卓3Web',
  `device_os_version` varchar(64) NOT NULL DEFAULT '' COMMENT '设备系统版本号',
  `device_id` varchar(64) NOT NULL DEFAULT '' COMMENT '设备ID',
  `version` varchar(64) NOT NULL DEFAULT '' COMMENT '版本号',
  `ip` int NOT NULL DEFAULT '0' COMMENT 'IP地址',
  `city` varchar(16) NOT NULL DEFAULT '' COMMENT '城市',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `change_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`sys_id`),
  KEY `idx_userid_createtime` (`user_id`,`create_time`),
  KEY `idx_userid` (`user_id`)
) ENGINE=InnoDB COMMENT '用户登录记录' DEFAULT CHARSET=utf8;


CREATE TABLE `user_online` (
  `sys_id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int unsigned NOT NULL COMMENT '用户ID',
  `offline_time` int unsigned NOT NULL DEFAULT '0' COMMENT '离线时间',
  `online_second` int unsigned NOT NULL DEFAULT '0' COMMENT '在线时长',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `change_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`sys_id`),
  KEY `idx_userid_createtime` (`user_id`,`create_time`)
) ENGINE=InnoDB COMMENT '用户在线' DEFAULT CHARSET=utf8;

CREATE TABLE `game_id_normal` (
  `sys_id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `game_id` int unsigned NOT NULL,
  `user_id` int unsigned NOT NULL,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`sys_id`),
  UNIQUE KEY `uk_gameid` (`game_id`)
) ENGINE=InnoDB COMMENT '预生成游戏id-普通' DEFAULT CHARSET=utf8;

CREATE TABLE `game_id_excellent` (
  `sys_id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `game_id` int unsigned NOT NULL,
  `user_id` int unsigned NOT NULL,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`sys_id`),
  UNIQUE KEY `uk_gameid` (`game_id`)
) ENGINE=InnoDB COMMENT '预生成游戏id-靓号' DEFAULT CHARSET=utf8;
