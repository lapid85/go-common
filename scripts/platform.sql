-- 所有表时间 created/updated 均为 UTC 微秒
SET time_zone = '+00:00';

-- 平台表
drop table if exists platforms;
create table if not exists platforms (
    id bigint not null auto_increment primary key,
    name varchar(128) not null default '' comment '平台名称',
    status tinyint not null default 0 comment '状态 0:禁用; 1:可用;',
    sort int not null default 0 comment '排序',
    remark varchar(255) not null default '' comment '备注',
    created bigint not null default 0 comment '创建时间 UTC微秒',
    updated bigint not null default 0 comment '更新时间 UTC微秒',
    index(status),
    index(created),
    index(sort),
    unique index(name)
) auto_increment = 1001;

INSERT INTO platforms (name, status, sort, remark, created, updated)
VALUES
    ('Platform A', 1, 10, 'Test Remark 1', UNIX_TIMESTAMP(NOW(6)) * 1000000, UNIX_TIMESTAMP(NOW(6)) * 1000000),
    ('Platform B', 0, 5, 'Test Remark 2', UNIX_TIMESTAMP(NOW(6)) * 1000000, UNIX_TIMESTAMP(NOW(6)) * 1000000);

-- 站点表
drop table if exists sites;
create table if not exists sites (
    id bigint not null auto_increment primary key,
    name varchar(128) not null default '' comment '站点名称',
    platform_id bigint not null default 0 comment '平台ID',
    platform_name varchar(128) not null default '' comment '平台名称',
    status tinyint not null default 0 comment '状态 0:禁用; 1:可用;',
    remark varchar(255) not null default '' comment '备注',
    sort int not null default 0 comment '排序',
    code varchar(4) not null default '' comment '站点编码',
    created bigint not null default 0 comment '创建时间 UTC微秒',
    updated bigint not null default 0 comment '更新时间 UTC微秒',
    index(status),
    index(created),
    index(sort),
    unique index(name),
    index(platform_id),
    unique index(code)
) auto_increment=10001;

INSERT INTO sites (name, platform_id, platform_name, status, remark, sort, code, created, updated)
VALUES
    ('Site A', 1001, 'Platform A', 1, 'Test Remark 1', 1, 'S001', UNIX_TIMESTAMP(NOW(6)) * 1000000, UNIX_TIMESTAMP(NOW(6)) * 1000000),
    ('Site B', 1002, 'Platform B', 0, 'Test Remark 2', 5, 'S002', UNIX_TIMESTAMP(NOW(6)) * 1000000, UNIX_TIMESTAMP(NOW(6)) * 1000000),
    ('Site C', 1001, 'Platform A', 1, 'Test Remark 3', 3, 'S003', UNIX_TIMESTAMP(NOW(6)) * 1000000, UNIX_TIMESTAMP(NOW(6)) * 1000000),
    ('Site D', 1002, 'Platform B', 0, 'Test Remark 4', 8, 'S004', UNIX_TIMESTAMP(NOW(6)) * 1000000, UNIX_TIMESTAMP(NOW(6)) * 1000000);

-- 配置表
drop table if exists configs;
create table if not exists configs (
    id bigint not null auto_increment primary key,
    platform_id bigint not null default 0 comment '平台ID',
    platform_name varchar(128) not null default '' comment '平台名称',
    site_id bigint not null default 0 comment '站点ID',
    site_name varchar(128) not null default '' comment '站点名称',
    name varchar(128) not null default '' comment '配置名称',
    value varchar(255) not null default '' comment '配置值',
    status tinyint not null default 0 comment '状态 0:禁用; 1:可用;',
    sort int not null default 0 comment '排序',
    remark varchar(255) not null default '' comment '备注',
    created bigint not null default 0 comment '创建时间 UTC微秒',
    updated bigint not null default 0 comment '更新时间 UTC微秒',
    index(platform_id),
    index(platform_name),
    index(site_id),
    index(site_name),
    index(created),
    index(sort),
    index(status),
    unique index(name)
) auto_increment = 100001;

INSERT INTO configs (platform_id, platform_name, site_id, site_name, name, value, status, sort, remark, created, updated)
VALUES
    (1001, 'Platform A', 10001, 'Site A', 'Config 01', 'Value 01', 1, 10, 'Test Remark 01', UNIX_TIMESTAMP(NOW(6)) * 1000000, UNIX_TIMESTAMP(NOW(6)) * 1000000),
    (1001, 'Platform A', 10001, 'Site A', 'Config 02', 'Value 02', 1, 20, 'Test Remark 02', UNIX_TIMESTAMP(NOW(6)) * 1000000, UNIX_TIMESTAMP(NOW(6)) * 1000000),
    (1001, 'Platform A', 10001, 'Site A', 'Config 03', 'Value 03', 0, 15, 'Test Remark 03', UNIX_TIMESTAMP(NOW(6)) * 1000000, UNIX_TIMESTAMP(NOW(6)) * 1000000),
    (1001, 'Platform A', 10001, 'Site A', 'Config 04', 'Value 04', 0, 50, 'Test Remark 10', UNIX_TIMESTAMP(NOW(6)) * 1000000, UNIX_TIMESTAMP(NOW(6)) * 1000000),
    (1001, 'Platform A', 10001, 'Site A', 'Config 05', 'Value 05', 1, 55, 'Test Remark 11', UNIX_TIMESTAMP(NOW(6)) * 1000000, UNIX_TIMESTAMP(NOW(6)) * 1000000),
    (1002, 'Platform B', 10002, 'Site B', 'Config 11', 'Value 14', 1, 15, 'Test Remark 04', UNIX_TIMESTAMP(NOW(6)) * 1000000, UNIX_TIMESTAMP(NOW(6)) * 1000000),
    (1002, 'Platform B', 10002, 'Site B', 'Config 12', 'Value 15', 1, 25, 'Test Remark 05', UNIX_TIMESTAMP(NOW(6)) * 1000000, UNIX_TIMESTAMP(NOW(6)) * 1000000),
    (1002, 'Platform B', 10002, 'Site B', 'Config 13', 'Value 16', 0, 30, 'Test Remark 06', UNIX_TIMESTAMP(NOW(6)) * 1000000, UNIX_TIMESTAMP(NOW(6)) * 1000000),
    (1002, 'Platform B', 10002, 'Site B', 'Config 14', 'Value 17', 0, 35, 'Test Remark 07', UNIX_TIMESTAMP(NOW(6)) * 1000000, UNIX_TIMESTAMP(NOW(6)) * 1000000),
    (1002, 'Platform B', 10002, 'Site B', 'Config 15', 'Value 18', 1, 40, 'Test Remark 08', UNIX_TIMESTAMP(NOW(6)) * 1000000, UNIX_TIMESTAMP(NOW(6)) * 1000000),
    (1002, 'Platform B', 10002, 'Site B', 'Config 16', 'Value 19', 1, 45, 'Test Remark 09', UNIX_TIMESTAMP(NOW(6)) * 1000000, UNIX_TIMESTAMP(NOW(6)) * 1000000);

-- 触发器 - 更新平台名称
DELIMITER //
CREATE TRIGGER trigger_update_platform_name
AFTER UPDATE ON platforms FOR EACH ROW
BEGIN
    UPDATE configs
    SET platform_name = NEW.name
    WHERE platform_id = NEW.id;

    UPDATE sites
    SET platform_name = NEW.name
    WHERE platform_id = NEW.id;
END;
//
DELIMITER ;

-- 触发器 - 更新站点名称
DELIMITER //
CREATE TRIGGER trigger_update_site_name
AFTER UPDATE ON sites FOR EACH ROW
BEGIN
    UPDATE configs
    SET site_name = NEW.name
    WHERE site_id = NEW.id;
END;
//
DELIMITER ;
