CREATE TABLE `feedback`  (
                        `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
                        `bulletin_id` bigint(20) UNSIGNED NOT NULL COMMENT '关联公告id',
                        `user_id` bigint(20) NULL DEFAULT NULL COMMENT '用户id',
                        `company` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '公司名称',
                        `contact_person` char(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '联系人',
                        `contact_information` char(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '联系方式',
                        `release_time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '发布时间',
                        `remarks` TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
                        `attachment` TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '附件',
                        `reject` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT ''COMMENT '驳回留言',
                        `status` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '反馈状态;0:待审核,1:通过，2:驳回',
                        `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
                        PRIMARY KEY (`id`) USING BTREE,
                        INDEX `bulletin_id` (`bulletin_id`) USING BTREE
) ENGINE = InnoDB  CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '招标信息' ROW_FORMAT = DYNAMIC;
alter table feedback add `user_id` bigint(20) NULL DEFAULT NULL COMMENT '用户id';
alter table feedback MODIFY `reject` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT ''COMMENT '驳回留言';
alter table feedback add `user_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名';
