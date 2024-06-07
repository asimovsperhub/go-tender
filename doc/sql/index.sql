
CREATE TABLE `statistics`  (
                                     `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
                                     `knowledge` int(11) NOT NULL DEFAULT 0 COMMENT '知识下载总量',
                                     `consultation` int(11) NOT NULL DEFAULT 0 COMMENT '历史咨询人数',
                                     PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB  CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '前台统计' ROW_FORMAT = COMPACT;