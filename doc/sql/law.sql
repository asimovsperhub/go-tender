


CREATE TABLE law(
                    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT 'UID',
                    fileid VARCHAR(50)  COMMENT '文件id',
                    publish VARCHAR(50) COMMENT '公布日期',
                    expiry VARCHAR(50) COMMENT '施行日期',
                    office VARCHAR(50) COMMENT '制定机关',
                    title VARCHAR(255) COMMENT '标题',
                    type VARCHAR(50) COMMENT '类型',
                    word  VARCHAR(255) COMMENT 'word',
                    pdf VARCHAR(255) COMMENT 'pdf',
                    url VARCHAR(255) COMMENT '外部链接'
) ENGINE = InnoDB  CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '常用法律表' ROW_FORMAT = COMPACT;