# Host: localhost
# Database: glow


#
# Structure for table "articles"
#

DROP TABLE IF EXISTS `articles`;
CREATE TABLE `articles` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '系统id',
  `title` varchar(150) NOT NULL COMMENT '文章标题',
  `cid` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '文档分类',
  `summary` varchar(255) DEFAULT NULL COMMENT '文章摘要',
  `content` text NOT NULL COMMENT '文章内容',
  `html_content` text,
  `author` varchar(20) DEFAULT NULL COMMENT '文章作者',
  `username` varchar(30) DEFAULT NULL COMMENT '添加文章的用户名',
  `view_count` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '文章浏览次数',
  `status` tinyint(2) unsigned NOT NULL DEFAULT '1' COMMENT '文章发布状态 0-未发布，1-已发布，2-暂存',
  `create_time` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `update_time` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

#
# Data for table "articles"
#


#
# Structure for table "category"
#

DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `id` smallint(5) unsigned NOT NULL AUTO_INCREMENT,
  `cat_name` varchar(50) NOT NULL COMMENT '分类名',
  `cat_info` varchar(100) DEFAULT '' COMMENT '分类说明',
  `cat_key` varchar(20) DEFAULT NULL COMMENT '分类名唯一标识',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '添加时间',
  `update_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;

#
# Data for table "category"
#

INSERT INTO `category` VALUES (3,'nginx','','nginx',1543919897,0),(6,'lnmp','好好的来分类','lnmp',1543934208,0),(7,'haohaohao','是的呢','php',1543934236,1543934236),(8,'linux','linux','linux',1543934678,0);

#
# Structure for table "comments"
#

DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `nickname` varchar(30) NOT NULL COMMENT '昵称',
  `site_url` varchar(150) DEFAULT '' COMMENT '网站url',
  `email` varchar(100) DEFAULT '' COMMENT '邮箱',
  `article_id` int(10) unsigned NOT NULL COMMENT '文章id',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0',
  `update_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

#
# Data for table "comments"
#


#
# Structure for table "link"
#

DROP TABLE IF EXISTS `link`;
CREATE TABLE `link` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `link_name` varchar(50) NOT NULL COMMENT '连接名称',
  `link_url` varchar(150) DEFAULT NULL COMMENT '链接url',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

#
# Data for table "link"
#


#
# Structure for table "sys_config"
#

DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `site_config` text,
  `create_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '记录添加时间',
  `update_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '0',
  `site_name` varchar(100) DEFAULT '' COMMENT '网站名称',
  `key_word` varchar(250) DEFAULT '' COMMENT '网站关键字',
  `site_desc` varchar(255) DEFAULT '' COMMENT '网站描述',
  `is_open` tinyint(2) unsigned NOT NULL DEFAULT '1' COMMENT '0-关闭 1-开启',
  `close_desc` varchar(255) DEFAULT '' COMMENT '关站描叙',
  `admin_email` varchar(100) DEFAULT '' COMMENT '网站联系人email',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

#
# Data for table "sys_config"
#

INSERT INTO `sys_config` VALUES (2,'',1543933613,0,'云鹤blog','blog|golang','',1,'升级中。。。','123@163.com');

#
# Structure for table "tag"
#

DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `article_id` int(10) unsigned NOT NULL COMMENT '文章id',
  `tag_name` varchar(20) NOT NULL COMMENT 'tag名',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0',
  `update_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

#
# Data for table "tag"
#

