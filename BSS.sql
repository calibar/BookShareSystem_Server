-- MySQL dump 10.13  Distrib 5.7.21, for Win64 (x86_64)
--
-- Host: localhost    Database: book_borrow_system
-- ------------------------------------------------------
-- Server version	5.7.21

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `applicant_list`
--

DROP TABLE IF EXISTS `applicant_list`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `applicant_list` (
  `book_id` int(11) NOT NULL,
  `applicant` varchar(32) NOT NULL,
  `applied_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`book_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf32;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `applicant_list`
--

LOCK TABLES `applicant_list` WRITE;
/*!40000 ALTER TABLE `applicant_list` DISABLE KEYS */;
/*!40000 ALTER TABLE `applicant_list` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `book_transaction`
--

DROP TABLE IF EXISTS `book_transaction`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `book_transaction` (
  `book_id` int(11) NOT NULL AUTO_INCREMENT,
  `book_name` varchar(64) NOT NULL,
  `book_author` varchar(64) DEFAULT NULL,
  `book_description` text,
  `book_cover` text,
  `book_owner` varchar(32) DEFAULT NULL,
  `book_borrower` varchar(32) DEFAULT NULL,
  `campus` varchar(32) NOT NULL,
  `post_expiration` date NOT NULL,
  `expect_return_time` date NOT NULL,
  `actual_return_time` date DEFAULT NULL,
  `post_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `owner_rating` int(11) DEFAULT NULL,
  `borrower_rating` int(11) DEFAULT NULL,
  `owner_comment` text,
  `borrower_comment` text,
  `book_status` varchar(32) NOT NULL,
  PRIMARY KEY (`book_id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf32;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `book_transaction`
--

LOCK TABLES `book_transaction` WRITE;
/*!40000 ALTER TABLE `book_transaction` DISABLE KEYS */;
INSERT INTO `book_transaction` VALUES (5,'Book1','Author1','just for testing','','mengxuan','LiuMengxuan','UofS','2018-12-08','2018-12-07','2006-01-01','2018-11-28 08:24:53',0,0,'','','borrowed'),(6,'Book1','Author1','just for testing','','','LiuMengxuan','UofS','2018-12-08','2018-12-07','2006-01-01','2018-11-28 08:34:00',0,0,'','','request'),(7,'Book1','Author1','just for testing','','','LiuMengxuan','UofS','2018-12-08','2018-12-07','2006-01-01','2018-11-28 08:34:57',0,0,'','','request'),(8,'Book1','Author1','just for testing','','LiuMengxuan','','UofS','2018-12-08','2018-12-07','2006-01-01','2018-11-28 08:36:29',0,0,'','','post');
/*!40000 ALTER TABLE `book_transaction` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `bug_report`
--

DROP TABLE IF EXISTS `bug_report`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `bug_report` (
  `report_id` int(11) NOT NULL AUTO_INCREMENT,
  `reporter` varchar(32) NOT NULL,
  `content` text NOT NULL,
  `response` text,
  PRIMARY KEY (`report_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf32;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `bug_report`
--

LOCK TABLES `bug_report` WRITE;
/*!40000 ALTER TABLE `bug_report` DISABLE KEYS */;
/*!40000 ALTER TABLE `bug_report` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `campus`
--

DROP TABLE IF EXISTS `campus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `campus` (
  `campus_id` int(11) NOT NULL AUTO_INCREMENT,
  `campus_name` varchar(100) NOT NULL,
  PRIMARY KEY (`campus_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf32;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `campus`
--

LOCK TABLES `campus` WRITE;
/*!40000 ALTER TABLE `campus` DISABLE KEYS */;
/*!40000 ALTER TABLE `campus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `uid` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(32) NOT NULL,
  `password` varchar(64) NOT NULL,
  `email` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf32;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (2,'mengxuan','a1b1c433d3c14af65c85687b6ec640da','cainmajest@gmail.com'),(6,'NickMengxuan','a1b1c433d3c14af65c85687b6ec640da','mel290@gmail.com'),(7,'LiuMengxuan','a1b1c433d3c14af65c85687b6ec640da','mengxuanliu@outlook.com');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_message`
--

DROP TABLE IF EXISTS `user_message`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_message` (
  `msg_id` int(11) NOT NULL AUTO_INCREMENT,
  `sender` varchar(32) NOT NULL,
  `receiver` varchar(32) NOT NULL,
  `content` text NOT NULL,
  `sending_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `msg_type` varchar(32) NOT NULL DEFAULT 'normal',
  `isDealed` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`msg_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf32;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_message`
--

LOCK TABLES `user_message` WRITE;
/*!40000 ALTER TABLE `user_message` DISABLE KEYS */;
INSERT INTO `user_message` VALUES (2,'mengxuan','LiuMengxuan','hello','2018-11-29 04:08:50','',0);
/*!40000 ALTER TABLE `user_message` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_profile`
--

DROP TABLE IF EXISTS `user_profile`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_profile` (
  `pid` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(32) NOT NULL,
  `nickname` varchar(32) DEFAULT NULL,
  `email` varchar(32) DEFAULT NULL,
  `campus` varchar(32) DEFAULT NULL,
  `student_id` varchar(16) DEFAULT NULL,
  `avatar` text,
  `lend_count` int(11) NOT NULL DEFAULT '0',
  `borrow_count` int(11) NOT NULL DEFAULT '0',
  `post_count` int(11) NOT NULL DEFAULT '0',
  `request_count` int(11) NOT NULL DEFAULT '0',
  `score` int(11) NOT NULL DEFAULT '50',
  `signup_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `badge` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf32;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_profile`
--

LOCK TABLES `user_profile` WRITE;
/*!40000 ALTER TABLE `user_profile` DISABLE KEYS */;
INSERT INTO `user_profile` VALUES (1,'mengxuan','Nick','cainmajesty@gmail.com','UofS','11239476','string',15,10,20,10,94,'2018-11-28 08:24:53','gold'),(2,'NickMengxuan','Nick','mel290@gmail.com','UofS','mel290','string',10,5,10,10,52,'2018-11-26 23:17:43','gold'),(3,'LiuMengxuan','Nick','mengxuanliu@outlook.com','UofS','mengxuanliu','string',10,6,11,11,49,'2018-11-28 08:36:29','gold');
/*!40000 ALTER TABLE `user_profile` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-11-28 16:12:33
