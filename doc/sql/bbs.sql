// 发布论坛表
CREATE TABLE `bbs`  (
                                     `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
                                     `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '标题',
                                     `abstract` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '摘要',
                                     `classification` char(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '所属分类',
                                     `review_message` char(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '审核留言',
                                     `review_status` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '审核状态;0:待审核,1:通过，2:未通过',
                                     `views` int(11) NOT NULL DEFAULT 0 COMMENT '浏览量',
                                     `reply_count` int(11) NOT NULL DEFAULT 0 COMMENT '回复量',
                                     `like_count` int(11) NOT NULL DEFAULT 0 COMMENT '点赞量',
                                     `user_id` bigint(20) NULL DEFAULT NULL COMMENT '用户id',
                                     `created_at` datetime NULL DEFAULT NULL COMMENT '发布时间',
                                     `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
                                     `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
                                     PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB  CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '发布论坛' ROW_FORMAT = COMPACT;
alter table bbs add `classification`  char(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '所属分类';
ALTER TABLE bbs ADD FULLTEXT INDEX title_content_classification(title,content,classification);
//status
`review_status` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '审核状态;0:待审核,1:通过，2:未通过',
alter table bbs add `status`  tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '审核状态;0:软删除,1:正常，2:永久删除';

alter table bbs MODIFY  title varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '标题';
alter table bbs MODIFY `review_message` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '审核留言';
alter table bbs MODIFY `abstract` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '摘要';
alter table bbs add `rank`  varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '排名';

alter table bbs MODIFY  title varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '标题';
alter table bbs MODIFY `content` LONGTEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '内容';
alter table bbs MODIFY `classification` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '所属分类';



CREATE TABLE `bbs_content`  (
                        `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
                        `bbs_id` bigint(20) NULL DEFAULT NULL COMMENT '论坛id',
                        `content` LONGTEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '内容',
                        PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB  CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '论坛内容' ROW_FORMAT = COMPACT;
ALTER TABLE bbs_content ADD FULLTEXT INDEX content(content);

CREATE TABLE `bbs_reply`  (
                        `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
                        `bbs_id` bigint(20) NULL DEFAULT NULL COMMENT '论坛id',
                        `reply_id` bigint(20) NULL DEFAULT NULL COMMENT '回复id',
                        `content` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '回复内容',
                        `like_count` int(11) NOT NULL DEFAULT 0 COMMENT '点赞量',
                        `user_id` bigint(20) NULL DEFAULT NULL COMMENT '用户id',
                        `created_at` datetime NULL DEFAULT NULL COMMENT '发布时间',
                        `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
                        `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
                        PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB  CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '论坛回复表' ROW_FORMAT = COMPACT;

CREATE TABLE `bbs_like`  (
                              `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
                              `bbs_id` bigint(20) NULL DEFAULT NULL COMMENT '论坛id',
                              `reply_id` bigint(20) NULL DEFAULT NULL COMMENT '回复id',
                              `user_id` bigint(20) NULL DEFAULT NULL COMMENT '用户id',
                              `created_at` datetime NULL DEFAULT NULL COMMENT '发布时间',
                              `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
                              `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
                              PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB  CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '论坛评论点赞表' ROW_FORMAT = COMPACT;