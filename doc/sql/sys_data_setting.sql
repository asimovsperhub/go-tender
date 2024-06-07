package sql

CREATE TABLE `sys_dataset`  (
                                   `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
                                   `consultation` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '咨询服务非营利类目',
                                   `purchase` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '采购咨询分类',
                                   `bid` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '投标咨询分类',
                                   `industry` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '行业咨询分类',
                                   `market` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '市场研究分类',
                                   `operation_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '操作管理员id',
                                   `created_at` datetime NULL DEFAULT NULL COMMENT '创建日期',
                                   `updated_at` datetime NULL DEFAULT NULL COMMENT '修改日期',
                                   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB  CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '数据中心设置' ROW_FORMAT = COMPACT;

alter table sys_dataset MODIFY purchase LONGTEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '采购咨询分类';
alter table sys_dataset MODIFY industry LONGTEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '行业咨询分类';
alter table sys_dataset MODIFY market LONGTEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '市场研究分类';
alter table sys_dataset add keywords varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '关键词';

alter table sys_dataset MODIFY `consultation` TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '咨询服务非营利类目';
alter table sys_dataset MODIFY `purchase` TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '采购咨询分类';
alter table sys_dataset MODIFY `bid` TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '投标咨询分类';
alter table sys_dataset MODIFY `industry` TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '行业咨询分类';
alter table sys_dataset MODIFY `market` TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '市场研究分类';

alter table sys_dataset add purchase_type TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '采购咨询类型';
alter table sys_dataset add industry_type TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '行业咨询类型';
alter table sys_dataset add market_type TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '市场研究类型';





