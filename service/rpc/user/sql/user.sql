CREATE DATABASE IF NOT EXISTS `microblog`;
USE `microblog`;
CREATE TABLE IF NOT EXISTS `user`(
  `id` bigint NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `phone` varchar(15) NOT NULL DEFAULT '',
  `email` char(255) NOT NULL DEFAULT '',
  `gender` ENUM('male', 'female', 'unknown') NOT NULL DEFAULT 'unknown' COMMENT '0:男, 1:女, 2:未知',
  `nickname` varchar(255) NOT NULL DEFAULT 'unknown',
  `birthday` date NOT NULL DEFAULT '1900-01-01',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY microblog_username_index(`username`),
  UNIQUE KEY microblog_phone_index(`phone`),
  UNIQUE KEY microblog_email_index(`email`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;