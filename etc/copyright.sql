drop database if exists copyrightp1;
create database copyrightp1 character set utf8;
use copyrightp1

drop table if exists assetok;
drop table if exists account_content;
drop table if exists aution;
drop table if exists account;
drop table if exists content;

create table account
(
   account_id           int not null primary key auto_increment,
   email                 varchar(50),
   username             varchar(30),
   identity_id          varchar(100),
   address              varchar(256)
);
CREATE UNIQUE INDEX account_email_uindex ON copyrightp1.account (email);
CREATE UNIQUE INDEX account_name_uindex ON copyrightp1.account (username);
alter table account comment '账户表';


create table content
(
   content_id           int not null primary key auto_increment,
   title                varchar(100),
   content              varchar(256),
   content_hash         varchar(100),
   ts                   timestamp
);

create table account_content
(
   content_hash         varchar(100),
   token_id             int,
   address              varchar(100),
   ts                   timestamp
);


create table auction
(
   content_hash         varchar(256),
   token_id             int,
   status               int,
   ts                   timestamp
);

CREATE table assetok
(
	content_hash         varchar(256),
	token_id             int,
	address              varchar(256)
)

