CREATE TABLE `user` (
                        `id` bigint(20) NOT NULL AUTO_INCREMENT,
                        `user_id` bigint(20) NOT NULL,
                        `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
                        `password` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
                        `email` varchar(64) COLLATE utf8mb4_general_ci,
                        `gender` tinyint(4) NOT NULL DEFAULT '0',
                        `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                        `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `idx_username` (`username`) USING BTREE,
                        UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


create table `community`(
    `id` int(11) not null auto_increment,
    `community_id` int(10) unsigned not null ,
    `community_name` varchar(128) collate utf8mb4_general_ci not null ,
    `introduction` varchar(256) collate utf8mb4_general_ci not null ,
    `create_time` timestamp not null default current_timestamp,
    `update_time` timestamp not null default current_timestamp on update current_timestamp,
    primary key (`id`),
    unique key `idx_community_id` (`community_id`),
    unique key `idx_community_name` (`community_name`)
)engine =InnoDB DEFAULT CHARSET = utf8mb4 collate =utf8mb4_general_ci;

-- 插入第一个记录
INSERT INTO `community` (`community_id`, `community_name`, `introduction`)
VALUES (5, 'Community A', 'This is a description of Community A.');

-- 插入第二个记录
INSERT INTO `community` (`community_id`, `community_name`, `introduction`)
VALUES (2, 'Community B', 'This is a description of Community B.');

-- 插入第三个记录
INSERT INTO `community` (`community_id`, `community_name`, `introduction`)
VALUES (3, 'Community C', 'This is a description of Community C.');

-- 插入第四个记录
INSERT INTO `community` (`community_id`, `community_name`, `introduction`)
VALUES (4, 'Community D', 'This is a description of Community D.');

create table `post`(
    `id` bigint(20) not null auto_increment,
    `post_id` bigint(20) not null comment '帖子' ,
    `title` varchar(128) collate utf8mb4_general_ci not null comment '标题',
    `content` varchar(8192) collate utf8mb4_general_ci not null comment '内容',
    `author_id` bigint(20) not null comment '作者id',
    `community_id` bigint(20) not null comment '所属社区',
    `status` tinyint(4) not null default '1' comment '帖子状态',
    `create_time` timestamp null default current_timestamp comment '创建时间',
    `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
    PRIMARY KEY (`id`),
    unique key `idx_post_id` (`post_id`),
    key `idx_author_id` (`author_id`),
    key `idx_community_id` (`community_id`)
 )engine = InnoDB default charset = utf8mb4 collate = utf8mb4_general_ci;

