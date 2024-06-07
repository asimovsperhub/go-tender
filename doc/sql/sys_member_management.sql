


CREATE TABLE `member_fee`  (
                                `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
                                `monthlycard_original` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '月卡原价',
                                `monthlycard_current` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '月卡现价',
                                `quartercard_original` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '季卡原价',
                                `quartercard_current` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '季卡现价',
                                `annualcard_original` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '年卡原价',
                                `annualcard_current` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '年卡现价',
                                `download_knowledge` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '下载知识',
                                `download_video` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '下载视频',
                                PRIMARY KEY (`id`) USING BTREE
                                ) ENGINE = InnoDB  CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '会员管理表' ROW_FORMAT = COMPACT;

CREATE TABLE `member_integral`  (
                                      `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
                                      `monthlycard_integral` int(11) NOT NULL DEFAULT 0 COMMENT '月卡积分',
                                      `quartercard_integral` int(11) NOT NULL DEFAULT 0 COMMENT '季卡积分',
                                      `annualcard_integral` int(11) NOT NULL DEFAULT 0 COMMENT '年卡积分',
                                      `knowledge_integral` int(11) NOT NULL DEFAULT 0 COMMENT '发布知识积分',
                                      `video_integral` int(11) NOT NULL DEFAULT 0 COMMENT '发布视频积分',
                                      `issue_integral` int(11) NOT NULL DEFAULT 0 COMMENT '发行积分',
                                      `ordinary` int(11) NOT NULL DEFAULT 0 COMMENT '普通',
                                      `select` int(11) NOT NULL DEFAULT 0 COMMENT '精选',
                                      PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB  CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '会员管理表' ROW_FORMAT = COMPACT;

alter table member_integral add  `ratio` int(11) NOT NULL DEFAULT 0 COMMENT '普通积分比例';
alter table member_integral add  `monthlycard_ratio` int(11) NOT NULL DEFAULT 0 COMMENT '月卡积分比例';
alter table member_integral add  `quartercard_ratio` int(11) NOT NULL DEFAULT 0 COMMENT '季卡积分比例';
alter table member_integral add  `annualcard_ratio` int(11) NOT NULL DEFAULT 0 COMMENT '年卡积分比例';




CREATE TABLE `member_subscription`  (
                                      `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
                                      `monthlycard_subscription` int(11) NOT NULL DEFAULT 0 COMMENT '月卡可订阅数',
                                      `quartercard_subscription` int(11) NOT NULL DEFAULT 0 COMMENT '季卡可订阅数',
                                      `annualcard_subscription` int(11) NOT NULL DEFAULT 0 COMMENT '年卡可订阅数',
                                      `monthlycard_subscription_price` varchar(60) NOT NULL DEFAULT '' COMMENT '月卡新增订阅单价',
                                      `quartercard_subscription_price` varchar(60) NOT NULL DEFAULT '' COMMENT '季卡新增订阅单价',
                                      `annualcard_subscription_price` varchar(60) NOT NULL DEFAULT '' COMMENT '年卡新增订阅单价',
                                      PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB  CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '会员管理表' ROW_FORMAT = COMPACT;