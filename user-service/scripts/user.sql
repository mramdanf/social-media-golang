DROP DATABASE IF EXISTS `social_media`;

CREATE DATABASE `social_media` ;

/*Table structure for table `userinfo` */

USE `social_media`;

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `uid` int(10) NOT NULL auto_increment,
  `fullName` varchar(64) default NULL,
  `email` varchar(64) default NULL,
  `password` varchar(64) default NULL,
  `created` date default NULL,
  PRIMARY KEY  (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;