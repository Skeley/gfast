-- phpMyAdmin SQL Dump
-- version 5.0.4
-- https://www.phpmyadmin.net/
--
-- 主机： localhost
-- 生成日期： 2025-03-22 09:27:45
-- 服务器版本： 5.7.44-log
-- PHP 版本： 7.4.33

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";

--
-- 数据库： `gfast-v32`
--

-- --------------------------------------------------------

--
-- 表的结构 `community`
--

CREATE TABLE `community` (
                             `major_id` int(10) UNSIGNED NOT NULL COMMENT '主id',
                             `minor_id` int(10) UNSIGNED NOT NULL COMMENT '次id',
                             `community_name` varchar(100) NOT NULL COMMENT '小区名'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;

-- --------------------------------------------------------

--
-- 表的结构 `community_union`
--

CREATE TABLE `community_union` (
                                   `id` int(10) UNSIGNED NOT NULL,
                                   `user_id` int(10) UNSIGNED NOT NULL COMMENT '用户id',
                                   `community_major_id` int(10) UNSIGNED NOT NULL COMMENT '小区主id',
                                   `community_minor_id` int(10) UNSIGNED NOT NULL COMMENT '小区次id'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;

-- --------------------------------------------------------

--
-- 表的结构 `material`
--

CREATE TABLE `material` (
                            `id` int(10) UNSIGNED NOT NULL COMMENT '耗材id',
                            `type` int(10) UNSIGNED NOT NULL COMMENT '耗材类型',
                            `brand` int(10) UNSIGNED NOT NULL COMMENT '耗材品牌',
                            `image` text NOT NULL COMMENT '图片url',
                            `price` int(10) UNSIGNED NOT NULL COMMENT '价格',
                            `comment` text NOT NULL COMMENT '描述'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;

-- --------------------------------------------------------

--
-- 表的结构 `material_brand`
--

CREATE TABLE `material_brand` (
                                  `brand` int(10) UNSIGNED NOT NULL COMMENT '耗材品牌',
                                  `name` varchar(100) NOT NULL COMMENT '品牌名'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;

-- --------------------------------------------------------

--
-- 表的结构 `material_type`
--

CREATE TABLE `material_type` (
                                 `type` int(10) UNSIGNED NOT NULL COMMENT '耗材类型',
                                 `name` varchar(100) NOT NULL COMMENT '类型名'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;

-- --------------------------------------------------------

--
-- 表的结构 `project`
--

CREATE TABLE `project` (
                           `id` int(10) UNSIGNED NOT NULL COMMENT '项目id',
                           `valid` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否有效',
                           `name` varchar(100) NOT NULL DEFAULT '' COMMENT '项目名',
                           `community_major_id` int(10) UNSIGNED NOT NULL COMMENT '小区主id',
                           `community_minor_id` int(10) UNSIGNED NOT NULL COMMENT '小区次id',
                           `start_date` datetime DEFAULT NULL COMMENT '开工日期',
                           `estimated_completion_date` datetime DEFAULT NULL COMMENT '预计完工日期',
                           `completion_date` datetime DEFAULT NULL COMMENT '完工日期',
                           `progress` tinyint(8) UNSIGNED DEFAULT '0' COMMENT '进度',
                           `inspection_report` text COMMENT '检查报告PDF链接',
                           `acceptance_report` text COMMENT '验收报告PDF链接',
                           `manager` int(10) UNSIGNED NOT NULL COMMENT '物业经理',
                           `associate` int(10) UNSIGNED DEFAULT NULL COMMENT '合伙人'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;

-- --------------------------------------------------------

--
-- 表的结构 `task`
--

CREATE TABLE `task` (
                        `id` bigint(20) UNSIGNED NOT NULL COMMENT '任务id',
                        `porject_id` int(10) UNSIGNED NOT NULL COMMENT '项目',
                        `name` varchar(100) NOT NULL DEFAULT '' COMMENT '项目名',
                        `type` int(10) UNSIGNED NOT NULL COMMENT '类型',
                        `start_date` datetime DEFAULT NULL COMMENT '开工日期',
                        `estimated_completion_date` datetime DEFAULT NULL COMMENT '预计完工日期',
                        `completion_date` datetime DEFAULT NULL COMMENT '完工日期',
                        `progress` tinyint(3) UNSIGNED NOT NULL DEFAULT '0' COMMENT '进度',
                        `acceptance_report` text COMMENT '验收报告PDF链接'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;

-- --------------------------------------------------------

--
-- 表的结构 `task_action`
--

CREATE TABLE `task_action` (
                               `task` bigint(20) UNSIGNED NOT NULL COMMENT '任务id',
                               `step` int(10) UNSIGNED NOT NULL COMMENT '步骤id',
                               `state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0: 未开始, 1: 进行中, 2: 已完成',
                               `modify_date` datetime DEFAULT NULL COMMENT '修改日期',
                               `comment` text NOT NULL COMMENT '描述',
                               `images` text NOT NULL COMMENT '图片url数组 {josn}'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;

-- --------------------------------------------------------

--
-- 表的结构 `task_stage`
--

CREATE TABLE `task_stage` (
                              `id` int(10) UNSIGNED NOT NULL COMMENT 'stage id',
                              `type` int(10) UNSIGNED NOT NULL COMMENT '所属任务类型',
                              `name` varchar(100) NOT NULL COMMENT '阶段名',
                              `icon` text NOT NULL COMMENT '图标',
                              `comment` text NOT NULL COMMENT '描述',
                              `order` tinyint(3) UNSIGNED NOT NULL COMMENT '任务流位置'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;

-- --------------------------------------------------------

--
-- 表的结构 `task_step`
--

CREATE TABLE `task_step` (
                             `id` int(10) UNSIGNED NOT NULL COMMENT '步骤id',
                             `stage_id` int(10) UNSIGNED NOT NULL COMMENT 'stage id',
                             `name` varchar(100) NOT NULL COMMENT '步骤名',
                             `comment` text NOT NULL COMMENT '描述',
                             `order` tinyint(3) UNSIGNED NOT NULL COMMENT '排序位置'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;

-- --------------------------------------------------------

--
-- 表的结构 `task_template`
--

CREATE TABLE `task_template` (
                                 `type` int(10) UNSIGNED NOT NULL COMMENT '类型',
                                 `name` varchar(100) NOT NULL COMMENT '默认项目名'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;

-- --------------------------------------------------------

--
-- 表的结构 `task_type`
--

CREATE TABLE `task_type` (
                             `type` int(10) UNSIGNED NOT NULL COMMENT '类型',
                             `name` varchar(100) NOT NULL COMMENT '类型名'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;

-- --------------------------------------------------------

--
-- 表的结构 `user_info`
--

CREATE TABLE `user_info` (
                             `id` int(10) UNSIGNED NOT NULL COMMENT '用户id',
                             `type` tinyint(16) UNSIGNED NOT NULL COMMENT '用户类型',
                             `community_major_id` int(10) UNSIGNED NOT NULL COMMENT '小区主id',
                             `community_minor_id` int(10) UNSIGNED NOT NULL COMMENT '小区次id',
                             `inviter` int(10) UNSIGNED NOT NULL COMMENT '邀请人id'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;

-- --------------------------------------------------------

--
-- 表的结构 `user_type`
--

CREATE TABLE `user_type` (
                             `type` tinyint(16) UNSIGNED NOT NULL COMMENT '用户类型',
                             `name` varchar(100) NOT NULL COMMENT 'name'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;

--
-- 转储表的索引
--

--
-- 表的索引 `community`
--
ALTER TABLE `community`
    ADD PRIMARY KEY (`major_id`,`minor_id`);

--
-- 表的索引 `community_union`
--
ALTER TABLE `community_union`
    ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`);

--
-- 表的索引 `material`
--
ALTER TABLE `material`
    ADD PRIMARY KEY (`id`);

--
-- 表的索引 `material_brand`
--
ALTER TABLE `material_brand`
    ADD PRIMARY KEY (`brand`);

--
-- 表的索引 `material_type`
--
ALTER TABLE `material_type`
    ADD PRIMARY KEY (`type`);

--
-- 表的索引 `project`
--
ALTER TABLE `project`
    ADD PRIMARY KEY (`id`),
  ADD KEY `manager` (`manager`),
  ADD KEY `associate` (`associate`);

--
-- 表的索引 `task`
--
ALTER TABLE `task`
    ADD PRIMARY KEY (`id`);

--
-- 表的索引 `task_stage`
--
ALTER TABLE `task_stage`
    ADD PRIMARY KEY (`id`);

--
-- 表的索引 `task_step`
--
ALTER TABLE `task_step`
    ADD PRIMARY KEY (`id`);

--
-- 表的索引 `task_template`
--
ALTER TABLE `task_template`
    ADD PRIMARY KEY (`type`);

--
-- 表的索引 `task_type`
--
ALTER TABLE `task_type`
    ADD PRIMARY KEY (`type`);

--
-- 表的索引 `user_info`
--
ALTER TABLE `user_info`
    ADD PRIMARY KEY (`id`);

--
-- 表的索引 `user_type`
--
ALTER TABLE `user_type`
    ADD PRIMARY KEY (`type`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `community_union`
--
ALTER TABLE `community_union`
    MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `material`
--
ALTER TABLE `material`
    MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '耗材id';

--
-- 使用表AUTO_INCREMENT `material_brand`
--
ALTER TABLE `material_brand`
    MODIFY `brand` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '耗材品牌';

--
-- 使用表AUTO_INCREMENT `material_type`
--
ALTER TABLE `material_type`
    MODIFY `type` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '耗材类型';

--
-- 使用表AUTO_INCREMENT `project`
--
ALTER TABLE `project`
    MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '项目id';

--
-- 使用表AUTO_INCREMENT `task`
--
ALTER TABLE `task`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '任务id';

--
-- 使用表AUTO_INCREMENT `task_stage`
--
ALTER TABLE `task_stage`
    MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'stage id';

--
-- 使用表AUTO_INCREMENT `task_step`
--
ALTER TABLE `task_step`
    MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '步骤id';

--
-- 使用表AUTO_INCREMENT `task_type`
--
ALTER TABLE `task_type`
    MODIFY `type` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '类型';

--
-- 使用表AUTO_INCREMENT `user_type`
--
ALTER TABLE `user_type`
    MODIFY `type` tinyint(16) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户类型';
COMMIT;
