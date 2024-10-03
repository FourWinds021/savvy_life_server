CREATE DATABASE IF NOT EXISTS savvy_data;

CREATE TABLE IF NOT EXISTS savvy_data.t_exam
(
    `id`                        BIGINT(20) NOT NULL AUTO_INCREMENT,
    `uuid`                      VARCHAR(64) NOT NULL DEFAULT '' COMMENT '用户唯一标识',
    `email`                     VARCHAR(256) NOT NULL DEFAULT '' COMMENT '用户邮箱',
    `phone_number`              VARCHAR(32) NOT NULL DEFAULT '' COMMENT '用户手机号',
    `nice_name`                 VARCHAR(64) NOT NULL DEFAULT '' COMMENT '用户别名',
    `password`                  VARCHAR(256) NOT NULL DEFAULT '' COMMENT '加密后用户密码',
    `sign_up_time`              timestamp NOT NULL DEFAULT '1970-01-01 08:00:01' COMMENT '签名时间',
    `create_time`               timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
    `update_time`               timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录变更时间',
)ENGINE INNODB DEFAULT CHARSET UTF8;
