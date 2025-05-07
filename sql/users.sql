create database if not exists dawn;

use dawn;

create table if not exists users
(
    id            bigint auto_increment comment 'id' primary key,
    account       varchar(256)                                     not null comment '账号',
    user_password varchar(512)                                     not null comment '密码',
    user_name     varchar(256)                                     null comment '用户昵称',
    user_avatar   varchar(1024)                                    null comment '用户头像',
    user_profile  varchar(512)                                     null comment '用户简介',
    user_role     enum ('user', 'admin') default 'user'            not null comment '用户角色',
    create_time   datetime               default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time   datetime               default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    is_delete     tinyint                default 0                 not null comment '是否删除 0 - 未删除 1 - 已删除'
) comment '用户' collate = utf8mb4_unicode_ci;

alter table users
add column user_account varchar(256) generated always as (if(is_delete = 0, account, null)) stored;

create unique index uk_user_account on users(user_account);
