CREATE DATABASE IF NOT EXISTS `game_engine` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE `game_engine`;

CREATE TABLE IF NOT EXISTS `t_category` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL COMMENT '分类名',
    `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_name` (`name`)
) ENGINE=InnoDB COMMENT='分类表';

CREATE TABLE IF NOT EXISTS `t_tag` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL COMMENT '标签名',
    `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_name` (`name`)
) ENGINE=InnoDB COMMENT='标签表';

CREATE TABLE IF NOT EXISTS `t_game` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL COMMENT '游戏名称',
    `distribute_type` TINYINT(1) NOT NULL COMMENT '游戏分发类型(链接/APK文件)',
    `developer` VARCHAR(255) NOT NULL COMMENT '开发商',
    `publisher` VARCHAR(255) NOT NULL COMMENT '发行商',
    `description` TEXT COMMENT '游戏描述',
    `details` TEXT COMMENT '游戏详情',

    `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '游戏状态(1:未上架 2:已上架 3:已下架)',
    `publish_time` DATETIME NULL COMMENT '发布时间',
    `reserve_count` BIGINT(20) DEFAULT 0 COMMENT '预约次数',

    `rating_score` BIGINT(20) DEFAULT 0 COMMENT '评分总分',
    `rating_count` BIGINT(20) DEFAULT 0 COMMENT '评分次数',
    `favorite_count` BIGINT(20) DEFAULT 0 COMMENT '收藏次数',
    `download_count` BIGINT(20) DEFAULT 0 COMMENT '下载次数',
    `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `update_time` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_name` (`name`),
    KEY `idx_status_publish_time` (`status`, `publish_time`)
) ENGINE=InnoDB COMMENT='游戏表';

CREATE TABLE IF NOT EXISTS `t_game_media_info` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `game_id` BIGINT(20) NOT NULL COMMENT '游戏ID',
    `file_id` VARCHAR(255) NOT NULL COMMENT '文件ID',
    `media_type` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '媒体类型',
    `media_url` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '媒体URL',
    `status` TINYINT(1) NOT NULL COMMENT '状态',
    `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_game_id_type_url` (`game_id`, `media_type`, `media_url`),
    UNIQUE KEY `idx_file_id` (`file_id`)
) ENGINE=InnoDB COMMENT='游戏媒体信息表';

CREATE TABLE IF NOT EXISTS `t_game_category` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `game_id` BIGINT(20) NOT NULL COMMENT '游戏ID',
    `category_id` BIGINT(20) NOT NULL COMMENT '分类ID',
    `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_game_id_category_id` (`game_id`, `category_id`)
) ENGINE=InnoDB COMMENT='游戏分类表';

CREATE TABLE IF NOT EXISTS `t_game_tag` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `game_id` BIGINT(20) NOT NULL COMMENT '游戏ID',
    `tag_id` BIGINT(20) NOT NULL COMMENT '标签ID',
    `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_game_id_tag_id` (`game_id`, `tag_id`)
) ENGINE=InnoDB COMMENT='游戏标签表';

CREATE TABLE IF NOT EXISTS `t_game_favorite` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `game_id` BIGINT(20) NOT NULL COMMENT '游戏ID',
    `user_id` BIGINT(20) NOT NULL COMMENT '用户ID',
    `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_game_id_user_id` (`game_id`, `user_id`),
    KEY `idx_user_id` (`user_id`),
    KEY `idx_game_id` (`game_id`)
) ENGINE=InnoDB COMMENT='游戏收藏表';

CREATE TABLE IF NOT EXISTS `t_game_rating` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `game_id` BIGINT(20) NOT NULL COMMENT '游戏ID',
    `user_id` BIGINT(20) NOT NULL COMMENT '用户ID',
    `score` BIGINT(20) NOT NULL COMMENT '评分总分',
    `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_game_id_user_id` (`game_id`, `user_id`)
) ENGINE=InnoDB COMMENT='游戏评分表';


INSERT INTO `t_game` (`name`, `distribute_type`, `developer`, `publisher`, `description`, `details`) VALUES ('测试游戏', 1, '测试开发商', '测试发行商', '测试描述', '测试详情');

INSERT INTO `t_category` (`name`) VALUES ('战争策略');
INSERT INTO `t_category` (`name`) VALUES ('动作枪战');
INSERT INTO `t_category` (`name`) VALUES ('赛车体育');
INSERT INTO `t_category` (`name`) VALUES ('棋牌桌游');
INSERT INTO `t_category` (`name`) VALUES ('格斗快打');
INSERT INTO `t_category` (`name`) VALUES ('儿童益智');
INSERT INTO `t_category` (`name`) VALUES ('休闲创意');
INSERT INTO `t_category` (`name`) VALUES ('模拟经营');

INSERT INTO `t_tag` (`name`) VALUES ('单机');
INSERT INTO `t_tag` (`name`) VALUES ('联机');
INSERT INTO `t_tag` (`name`) VALUES ('氪金');
INSERT INTO `t_tag` (`name`) VALUES ('免费');
INSERT INTO `t_tag` (`name`) VALUES ('RPG');
INSERT INTO `t_tag` (`name`) VALUES ('卡牌');
INSERT INTO `t_tag` (`name`) VALUES ('独立');
INSERT INTO `t_tag` (`name`) VALUES ('组队');

INSERT INTO `t_game_category` (`game_id`, `category_id`) VALUES (1, 1);
INSERT INTO `t_game_category` (`game_id`, `category_id`) VALUES (1, 3);

INSERT INTO `t_game_tag` (`game_id`, `tag_id`) VALUES (1, 1);







-- 游戏预约表
CREATE TABLE IF NOT EXISTS `t_game_reserve` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `game_id` BIGINT(20) NOT NULL COMMENT '游戏ID',
    `user_id` BIGINT(20) NOT NULL COMMENT '用户ID',
    `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_game_id_user_id` (`game_id`, `user_id`),
    KEY `idx_user_id` (`user_id`),
    KEY `idx_game_id` (`game_id`)
) ENGINE=InnoDB COMMENT='游戏预约表';


-- 搜索历史表（仅登录用户）
CREATE TABLE IF NOT EXISTS `t_search_history` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT(20) NOT NULL COMMENT '用户ID',
    `search_keyword` VARCHAR(255) NOT NULL COMMENT '搜索关键词',
    `search_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `result_count` INT DEFAULT 0 COMMENT '搜索结果数量',
    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`user_id`),
    KEY `idx_search_time` (`search_time`)
) ENGINE=InnoDB COMMENT='搜索历史表';

-- 游戏行为记录表（仅登录用户）
CREATE TABLE IF NOT EXISTS `t_game_behavior` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT(20) NOT NULL COMMENT '用户ID',
    `game_id` BIGINT(20) NOT NULL COMMENT '游戏ID',
    `behavior_type` TINYINT(1) NOT NULL COMMENT '行为类型(1:查看 2:下载 3:收藏 4:评分)',
    `behavior_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `ip_address` VARCHAR(45) COMMENT 'IP地址',
    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`user_id`),
    KEY `idx_game_id` (`game_id`),
    KEY `idx_behavior_time` (`behavior_time`)
) ENGINE=InnoDB COMMENT='游戏行为记录表';