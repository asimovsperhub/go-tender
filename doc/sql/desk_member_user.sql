





// 用户表
CREATE TABLE `member_user`  (
                             `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
                             `user_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
                             `mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '中国手机不带国家代码，国际手机号格式为：国家代码-手机号',
                             `user_nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户昵称',
                             `birthday` int(11) NOT NULL DEFAULT 0 COMMENT '生日',
                             `user_password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '登录密码;cmf_password加密',
                             `user_salt` char(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '加密盐',
                             `user_status` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '用户状态;0:禁用,1:正常,2:未验证',
                             `disabled_day` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '禁用天数',
                             `user_email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户登录邮箱',
                             `sex` tinyint(2) NOT NULL DEFAULT 0 COMMENT '性别;0:保密,1:男,2:女',
                             `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户头像',
                             `dept_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '部门id',
                             `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
                             `member_level` tinyint(4) NOT NULL DEFAULT 0 COMMENT ' 1月卡 2季卡 3 年卡   0 否',
                             `integral` int(11) NOT NULL DEFAULT 0 COMMENT '积分',
                             `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '联系地址',
                             `describe` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT ' 描述信息',
                             `last_login_ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '最后登录ip',
                             `last_login_time` datetime NULL DEFAULT NULL COMMENT '最后登录时间',
                             `operation_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '操作管理员id',
                             `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
                             `disabled_at` datetime NULL DEFAULT NULL COMMENT '禁用时间',
                             `release_at` datetime NULL DEFAULT NULL COMMENT '解禁时间',
                             `maturity_at` datetime NULL DEFAULT NULL COMMENT '到期时间',
                             `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
                             `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
                             PRIMARY KEY (`id`) USING BTREE,
                             UNIQUE INDEX `user_login`(`user_name`, `deleted_at`) USING BTREE,
                             UNIQUE INDEX `mobile`(`mobile`, `deleted_at`) USING BTREE,
                             INDEX `user_nickname`(`user_nickname`) USING BTREE
) ENGINE = InnoDB  CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '会员用户表' ROW_FORMAT = COMPACT;

alter table member_user add `disabled_day` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '禁用天数';
alter table Address MODIFY addr VARCHAR(100)
alter table member_user MODIFY `disabled_day` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '禁用天数';
alter table member_user add `subscribe` int(11) NOT NULL DEFAULT 0 COMMENT '可订阅数';
alter table member_user add `buy_subscribe` int(11) NOT NULL DEFAULT 0 COMMENT '购买的订阅数';
alter table member_user add `subscribed` int(11) NOT NULL DEFAULT 0 COMMENT '已订阅数';
alter table member_user MODIFY `member_level` tinyint(4) NOT NULL DEFAULT 0 COMMENT ' 1月卡 2季卡 3 年卡   0 否';
INSERT INTO `member_user` VALUES (1, 'mxw', '13578342363', '蛮细', 0, 'mxw', 'f9aZTAa8yz', 1, 'yxh669@qq.com', 1, '用户头像', 101, '测试用户', 0, 9999,'联系地址', '描述信息', '最后登陆ip', '2022-10-26 03:01:52', 1, '', '','','','','');
INSERT INTO `member_user` VALUES (2, 'cc', '15399225820', '别名', 0, 'mxw', 'f9aZTAa8yz', 1, 'asimov@qq.com', 1, '用户头像1', 101, '测试用户1', 0, 9999,'联系地址1', '描述信息1', '最后登陆ip', '2022-10-26 03:01:52', 1, '', '','','','','');
INSERT INTO `member_user` VALUES (3, 'mxwzz', '13578342363', '蛮细', 0, 'mxw', 'f9aZTAa8yz', 1, 'yxh669@qq.com', 1, '用户头像', 101, '测试用户', 0, 9999,'联系地址', '描述信息', '最后登陆ip', '2022-10-26 03:01:52', 1, '', '','','','',NULL);
INSERT INTO `member_user` VALUES (4, 'cczz', '15399225820', '别名', 0, 'mxw', 'f9aZTAa8yz', 1, 'asimov@qq.com', 1, '用户头像1', 101, '测试用户1', 0, 9999,'联系地址1', '描述信息1', '最后登陆ip', '2022-10-26 03:01:52', 1, '', '','','','',NULL,0);

INSERT INTO `member_user` VALUES (5, 'od', '15399999999', '别名', 0, 'password', 'f9aZTAa8yz', 1, 'asimov1@qq.com', 1, '用户头像1', 101, '测试用户1', 1, 9999,'联系地址1', '描述信息1', '最后登陆ip', '2022-10-26 03:01:52', 1, '', '','','','',NULL,0);

// 发布知识表
CREATE TABLE `member_knowledge`  (
                                `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
                                `title` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '标题',
                                `knowledge_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '知识类型',
                                `authority` tinyint(2) UNSIGNED NOT NULL DEFAULT 0 COMMENT '阅读下载权限，0:所有,1:会员',
                                `primary_classification` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '一级分类',
                                `secondary_classification` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '二级分类',
                                `integral_setting` int(11) NOT NULL DEFAULT 0 COMMENT '积分设置',
                                `content` LONGTEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '内容',
                                `review_message` char(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '审核留言',
                                `review_status` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '审核状态;0:待审核,1:通过，2:未通过',
                                `user_id` bigint(20) NULL DEFAULT NULL COMMENT '发布id',
                                `user_name`  varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '发布作者',
                                `created_at` datetime NULL DEFAULT NULL COMMENT '发布时间',
                                `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
                                `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
                                PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB  CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '发布知识' ROW_FORMAT = COMPACT;

alter table member_knowledge add `user_name`  varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '发布作者';
alter table member_knowledge add `type` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '0:其他,1:视频';
alter table member_knowledge add `cover_url` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '封面主图';
alter table member_knowledge add `details_url` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '详情主图';
alter table member_knowledge add `video_url` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '视频链接';
alter table member_knowledge add `video_introduction` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '视频简介';
alter table member_knowledge add `views` int(11) NOT NULL DEFAULT 0 COMMENT '浏览量';
alter table member_knowledge add `attachment_url` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '附件链接';
alter table member_knowledge add `opreview_message` char(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '管理员留言';
alter table member_knowledge add `shortvideo_url` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '视频链接';
alter table member_knowledge add `abstract` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '摘要';
alter table member_knowledge add `display` tinyint(2) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否展示,0不展示1展示';
ALTER TABLE member_knowledge ADD FULLTEXT INDEX title_abstract(title,abstract);

alter table member_knowledge MODIFY `opreview_message` char(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '管理员留言';
alter table member_knowledge MODIFY `review_message` char(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '审核留言',

alter table member_knowledge MODIFY  title varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '标题';
alter table member_knowledge MODIFY `review_message` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '审核留言';
alter table member_knowledge MODIFY `abstract` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '摘要';
alter table member_knowledge MODIFY `details_url` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '详情主图';

alter table member_knowledge MODIFY  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '标题';
alter table member_knowledge MODIFY  `details_url` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '详情主图';
alter table member_knowledge MODIFY  `video_introduction` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '视频简介';
alter table member_knowledge MODIFY  `opreview_message` VARCHAR(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '管理员留言';
alter table member_knowledge MODIFY  `review_message` VARCHAR(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '审核留言';
alter table member_knowledge MODIFY `cover_url` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '封面主图';

alter table member_knowledge MODIFY `attachment_url` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '附件链接';
alter table member_knowledge MODIFY `video_url` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '视频链接';
alter table member_knowledge MODIFY `shortvideo_url` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '视频链接';
alter table member_knowledge add `cut_status` tinyint(2) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否展示,0进行中 1 已完成';

// 知识下载/购买记录表

CREATE TABLE `his_knowledge`  (
                                     `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
                                     `knowledge_id` bigint(20) NULL DEFAULT NULL COMMENT '知识id',
                                     `user_id` bigint(20) NULL DEFAULT NULL COMMENT '用户id',
                                     `created_at` datetime NULL DEFAULT NULL COMMENT '发布时间',
                                     `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
                                     `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
                                     PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB  CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '发布知识' ROW_FORMAT = COMPACT;



//collect 收藏夹
CREATE TABLE `member_collect`  (
                                     `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
                                     `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '标题',
                                     `type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '类型',
                                     `location` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '所在地',
                                     `industry` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '行业',
                                     `user_id` int(10) NULL DEFAULT NULL COMMENT '用户id',
                                     `article_id` int(10) NULL DEFAULT NULL COMMENT '文章id',
                                     `created_at` datetime NULL DEFAULT NULL COMMENT '收藏时间',
                                     `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
                                     `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
                                     PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB  CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '收藏夹' ROW_FORMAT = COMPACT;
alter table member_collect add  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '收藏url';
alter table member_collect MODIFY `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '收藏url';
// 订阅表 subscribe

CREATE TABLE `member_subscribe`  (
                                   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
                                   `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '方案名称',
                                   `type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '信息类型',
                                   `location` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '所在地',
                                   `keywords` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '关键字',
                                   `user_id` int(10) NULL DEFAULT NULL COMMENT '用户id',
                                   `created_at` datetime NULL DEFAULT NULL COMMENT '收藏时间',
                                   `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
                                   `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
                                   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB  CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '收藏夹' ROW_FORMAT = COMPACT;

// industry
alter table member_subscribe add  `industry` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '行业类型';