CREATE TABLE `sys_enterprise`  (
                                  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
                                  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '企业名称',
                                  `location` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '所在地',
                                  `industry` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '所属行业',
                                  `contact` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '联系电话',
                                  `icon` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '图标',
                                  `introduction` varchar(800) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '简介',
                                  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '其他',
                                  `license` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '营业执照',
                                  `license_status` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '审核状态 0待审核 1审核通过 2审核未通过',
                                  `license_message` char(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '营业执照审核留言',
                                  `certificate` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '证明书',
                                  `certificate_status` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '审核状态 0待审核 1审核通过 2审核未通过',
                                  `certificate_message` char(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '证明书审核留言',
                                  `operation_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '操作管理员id',
                                  `establishment_at` datetime NULL DEFAULT NULL COMMENT '成立时间',
                                  `created_at` datetime NULL DEFAULT NULL COMMENT '创建日期',
                                  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改日期',
                                  PRIMARY KEY (`id`) USING BTREE,
                                  UNIQUE INDEX `name`(`name`) USING BTREE
) ENGINE = InnoDB  CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '企业管理' ROW_FORMAT = COMPACT;


INSERT INTO `sys_enterprise` VALUES (1, '九立咨询', '深圳', '工程建设', '15399225820', '/mnt/ico.png', '工程建设加速绅士手是是', 'remark', '/mnt/license.png', 0, '/mnt/certificate.png', 0, 1, '','','');
//修改列
alter table sys_enterprise MODIFY establishment_at VARCHAR(50) NULL DEFAULT NULL DEFAULT '' COMMENT '成立时间';
// 新增列
alter table sys_enterprise MODIFY  license_message char(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '营业执照审核留言';
alter table sys_enterprise MODIFY  `oplicense_message` char(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '营业执照管理员留言';

alter table sys_enterprise MODIFY `certificate_message` char(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '证明书审核留言';
alter table sys_enterprise MODIFY `opcertificate_message` char(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '证明书管理员留言';

//`user_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
alter table sys_enterprise add  `user_id` int(10) UNSIGNED NULL DEFAULT NULL  COMMENT '用户id';

alter table sys_enterprise add `nickname` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '企业简称',
//content
alter table sys_enterprise add `content` LONGTEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '主页内容';

// 升级完mysql遇到的问题 :Row size too large (> 8126). Changing some columns to TEXT or BLOB or using ROW_FORMAT=DYNAMIC or ROW_FORMAT=COMPRESSED may help. In current row format, BLOB prefix of 768 bytes is stored inline.

ALTER  TABLE sys_enterprise ROW_FORMAT = DYNAMIC;


alter table sys_enterprise MODIFY  license_message VARCHAR(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '营业执照审核留言';
alter table sys_enterprise MODIFY  `oplicense_message` VARCHAR(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '营业执照管理员留言';
alter table sys_enterprise MODIFY `certificate_message` VARCHAR(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '证明书审核留言';
alter table sys_enterprise MODIFY `opcertificate_message` VARCHAR(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '证明书管理员留言';
alter table sys_enterprise MODIFY `license` VARCHAR(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '营业执照';
alter table sys_enterprise MODIFY `certificate` VARCHAR(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '证明书',



// 天眼查search

CREATE TABLE `tyc`  (
                       `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
                       `name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '企业名称',
                       `type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '类型:commerce-工商信息,punishment-经营风险,qualification-资质,lawSuit-司法风险',
                       `number` int(10) UNSIGNED NOT NULL DEFAULT 1 COMMENT '页吗',
                       `size` int(10) UNSIGNED NOT NULL DEFAULT 20 COMMENT '页大小',
                       `body` TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '内容',
                       `created_at` datetime NULL DEFAULT NULL COMMENT '创建日期',
                       `updated_at` datetime NULL DEFAULT NULL COMMENT '修改日期',
                       `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
                       PRIMARY KEY (`id`) USING BTREE,
                       INDEX `name`(`name`) USING BTREE
) ENGINE = InnoDB  CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '天眼查' ROW_FORMAT = DYNAMIC;
