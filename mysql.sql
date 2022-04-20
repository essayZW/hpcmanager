CREATE DATABASE  IF NOT EXISTS `hpcmanager` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `hpcmanager`;
-- MySQL dump 10.13  Distrib 8.0.28, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: hpcmanager
-- ------------------------------------------------------
-- Server version	8.0.26

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `group`
--

DROP TABLE IF EXISTS `group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `group` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL,
  `hpc_group_id` int NOT NULL DEFAULT '0',
  `create_time` timestamp NOT NULL,
  `creater_id` int NOT NULL,
  `creater_username` varchar(32) NOT NULL,
  `creater_name` varchar(32) NOT NULL,
  `tutor_id` int NOT NULL DEFAULT '0',
  `tutor_username` varchar(32) DEFAULT NULL,
  `tutor_name` varchar(32) DEFAULT NULL,
  `balance` decimal(18,2) DEFAULT '0.00',
  `extraAttributes` varchar(2048) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_UNIQUE` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `hpc_group`
--

DROP TABLE IF EXISTS `hpc_group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `hpc_group` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL,
  `queue_name` varchar(64) NOT NULL,
  `gid` int NOT NULL,
  `extraAttributes` varchar(2048) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_UNIQUE` (`name`),
  UNIQUE KEY `queue_name_UNIQUE` (`queue_name`),
  UNIQUE KEY `gid_UNIQUE` (`gid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `hpc_usagetime`
--

DROP TABLE IF EXISTS `hpc_usagetime`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `hpc_usagetime` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `username` varchar(32) NOT NULL,
  `user_name` varchar(32) NOT NULL,
  `hpc_username` varchar(64) NOT NULL,
  `tutor_id` int NOT NULL,
  `tutor_username` varchar(32) NOT NULL,
  `tutor_user_name` varchar(32) NOT NULL,
  `hpc_group_name` varchar(64) NOT NULL,
  `queue_name` varchar(64) NOT NULL,
  `wall_time` int NOT NULL DEFAULT '0',
  `gwall_time` int NOT NULL DEFAULT '0',
  `start_time` timestamp NOT NULL,
  `end_time` timestamp NOT NULL,
  `create_time` timestamp NOT NULL,
  `extraAttributes` varchar(2048) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `hpc_user`
--

DROP TABLE IF EXISTS `hpc_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `hpc_user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `node_username` varchar(64) NOT NULL,
  `node_uid` int NOT NULL,
  `node_max_quota` int DEFAULT '0',
  `quota_start_time` timestamp NULL DEFAULT NULL,
  `quota_end_time` timestamp NULL DEFAULT NULL,
  `extraAttributes` varchar(2048) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `node_apply`
--

DROP TABLE IF EXISTS `node_apply`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `node_apply` (
  `id` int NOT NULL AUTO_INCREMENT,
  `create_time` timestamp NOT NULL,
  `creater_id` int NOT NULL,
  `creater_username` varchar(32) NOT NULL,
  `creater_name` varchar(32) NOT NULL,
  `project_id` int NOT NULL,
  `tutor_check_status` tinyint NOT NULL DEFAULT '-1',
  `manager_check_status` tinyint NOT NULL DEFAULT '-1',
  `status` tinyint NOT NULL DEFAULT '1',
  `message_tutor` varchar(300) DEFAULT NULL,
  `message_manager` varchar(300) DEFAULT NULL,
  `tutor_check_time` timestamp NULL DEFAULT NULL,
  `tutor_id` int NOT NULL,
  `tutor_name` varchar(32) NOT NULL,
  `tutor_username` varchar(32) NOT NULL,
  `manager_check_time` timestamp NULL DEFAULT NULL,
  `manager_checker_id` int DEFAULT NULL,
  `manager_checker_username` varchar(32) DEFAULT NULL,
  `manager_checker_name` varchar(32) DEFAULT NULL,
  `modify_time` timestamp NULL DEFAULT NULL,
  `modify_userid` int DEFAULT NULL,
  `modify_name` varchar(32) DEFAULT NULL,
  `modify_username` varchar(32) DEFAULT NULL,
  `node_type` varchar(64) NOT NULL,
  `node_num` int NOT NULL,
  `start_time` timestamp NOT NULL,
  `end_time` timestamp NOT NULL,
  `extraAttributes` varchar(2048) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `node_distribute`
--

DROP TABLE IF EXISTS `node_distribute`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `node_distribute` (
  `id` int NOT NULL AUTO_INCREMENT,
  `apply_id` int NOT NULL,
  `handler_flag` tinyint NOT NULL DEFAULT '0',
  `handler_userid` int DEFAULT NULL,
  `handler_username` varchar(32) DEFAULT NULL,
  `handler_user_name` varchar(32) DEFAULT NULL,
  `distribute_bill_id` int DEFAULT '0',
  `create_time` timestamp NOT NULL,
  `extraAttributes` varchar(2048) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `node_distribute_bill`
--

DROP TABLE IF EXISTS `node_distribute_bill`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `node_distribute_bill` (
  `id` int NOT NULL AUTO_INCREMENT,
  `apply_id` int NOT NULL,
  `node_distribute_id` int NOT NULL,
  `fee` decimal(18,2) NOT NULL,
  `pay_fee` decimal(18,2) NOT NULL DEFAULT '0.00',
  `pay_flag` tinyint DEFAULT '0',
  `pay_time` timestamp NULL DEFAULT NULL,
  `pay_type` tinyint DEFAULT NULL,
  `pay_message` varchar(512) DEFAULT NULL,
  `user_id` int NOT NULL,
  `user_username` varchar(32) NOT NULL,
  `user_name` varchar(32) NOT NULL,
  `user_group_id` int NOT NULL,
  `create_time` timestamp NOT NULL,
  `extraAttributes` varchar(2048) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `node_distribute_log`
--

DROP TABLE IF EXISTS `node_distribute_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `node_distribute_log` (
  `id` int NOT NULL AUTO_INCREMENT,
  `queue_name` varchar(64) NOT NULL,
  `user_id` int NOT NULL,
  `distribute_id` int NOT NULL,
  `nodes_names` varchar(512) NOT NULL,
  `operation` tinyint NOT NULL,
  `create_time` timestamp NOT NULL,
  `extraAttributes` varchar(2048) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `node_quota_bill`
--

DROP TABLE IF EXISTS `node_quota_bill`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `node_quota_bill` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `user_name` varchar(32) NOT NULL,
  `user_username` varchar(32) NOT NULL,
  `user_group_id` int NOT NULL DEFAULT '-1',
  `oper_type` tinyint NOT NULL,
  `old_size` int NOT NULL,
  `new_size` int NOT NULL,
  `old_end_time` timestamp NOT NULL,
  `new_end_time` timestamp NOT NULL,
  `fee` decimal(18,2) NOT NULL,
  `pay_flag` tinyint DEFAULT '0',
  `pay_fee` decimal(18,2) DEFAULT '0.00',
  `pay_time` timestamp NULL DEFAULT NULL,
  `pay_type` tinyint DEFAULT NULL,
  `pay_message` varchar(300) DEFAULT NULL,
  `create_time` timestamp NOT NULL,
  `extraAttributes` varchar(2048) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `paper_apply`
--

DROP TABLE IF EXISTS `paper_apply`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `paper_apply` (
  `id` int NOT NULL AUTO_INCREMENT,
  `creater_id` int NOT NULL,
  `creater_username` varchar(32) NOT NULL,
  `creater_name` varchar(32) NOT NULL,
  `create_time` timestamp NOT NULL,
  `user_group_id` int NOT NULL,
  `tutor_id` timestamp NOT NULL,
  `tutor_username` varchar(32) NOT NULL,
  `tutor_name` varchar(32) NOT NULL,
  `paper_title` varchar(256) NOT NULL,
  `paper_category` varchar(128) NOT NULL,
  `paper_partition` varchar(32) NOT NULL,
  `papaer_firstpage_img` varchar(512) NOT NULL,
  `paper_thankspage_img` varchar(512) NOT NULL,
  `remark_message` varchar(512) DEFAULT NULL,
  `check_status` tinyint DEFAULT '-1',
  `checker_id` int DEFAULT NULL,
  `checker_name` varchar(32) DEFAULT NULL,
  `checker_username` varchar(32) DEFAULT NULL,
  `check_money` decimal(18,2) DEFAULT '0.00',
  `check_message` varchar(512) DEFAULT NULL,
  `check_time` timestamp NULL DEFAULT NULL,
  `extraAttributes` varchar(2048) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `permission`
--

DROP TABLE IF EXISTS `permission`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `permission` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL,
  `level` tinyint NOT NULL,
  `description` varchar(512) NOT NULL,
  `create_time` timestamp NOT NULL,
  `extraAttributes` varchar(2048) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `level_UNIQUE` (`level`),
  UNIQUE KEY `name_UNIQUE` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `project`
--

DROP TABLE IF EXISTS `project`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `project` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(128) NOT NULL,
  `from` varchar(128) NOT NULL,
  `numbering` varchar(128) NOT NULL,
  `expenses` varchar(16) NOT NULL,
  `description` varchar(1024) DEFAULT NULL,
  `creater_user_id` int NOT NULL,
  `creater_username` varchar(32) DEFAULT NULL,
  `creater_user_name` varchar(32) DEFAULT NULL,
  `create_time` timestamp NOT NULL,
  `modify_time` timestamp NOT NULL,
  `modify_user_id` int NOT NULL,
  `modify_username` varchar(32) DEFAULT NULL,
  `modify_user_name` varchar(45) DEFAULT NULL,
  `extraAttributes` varchar(2048) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `technology_apply`
--

DROP TABLE IF EXISTS `technology_apply`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `technology_apply` (
  `id` int NOT NULL AUTO_INCREMENT,
  `creater_id` int NOT NULL,
  `create_time` timestamp NOT NULL,
  `tutor_id` int NOT NULL,
  `project_id` int NOT NULL,
  `project_name` varchar(128) NOT NULL,
  `project_description` varchar(1024) DEFAULT NULL,
  `prize_level` varchar(128) NOT NULL,
  `prize_img` varchar(512) NOT NULL,
  `remark_message` varchar(512) DEFAULT NULL,
  `check_status` tinyint DEFAULT '-1',
  `checker_id` int DEFAULT NULL,
  `checker_name` varchar(32) DEFAULT NULL,
  `check_message` varchar(512) DEFAULT NULL,
  `check_time` timestamp NULL DEFAULT NULL,
  `extraAttributes` varchar(2048) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(32) NOT NULL,
  `password` char(32) NOT NULL,
  `tel` varchar(16) DEFAULT NULL,
  `email` varchar(128) DEFAULT NULL,
  `name` varchar(32) NOT NULL,
  `pinyin_name` varchar(64) NOT NULL,
  `college_name` varchar(64) DEFAULT NULL,
  `group_id` int NOT NULL,
  `hpc_user_id` int NOT NULL DEFAULT '0',
  `create_time` timestamp NOT NULL,
  `extraAttributes` varchar(2048) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username_UNIQUE` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `user_group_apply`
--

DROP TABLE IF EXISTS `user_group_apply`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_group_apply` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `user_username` varchar(32) NOT NULL,
  `user_name` varchar(32) NOT NULL,
  `apply_group_id` int NOT NULL,
  `tutor_id` int NOT NULL,
  `tutor_username` varchar(32) NOT NULL,
  `tutor_name` varchar(32) NOT NULL,
  `tutor_check_status` tinyint NOT NULL DEFAULT '-1',
  `manager_check_status` tinyint NOT NULL DEFAULT '-1',
  `status` tinyint NOT NULL DEFAULT '1',
  `message_tutor` varchar(300) DEFAULT NULL,
  `message_manager` varchar(300) DEFAULT NULL,
  `tutor_check_time` timestamp NULL DEFAULT NULL,
  `manager_check_time` timestamp NULL DEFAULT NULL,
  `manager_checker_id` int DEFAULT NULL,
  `manager_checker_username` varchar(32) DEFAULT NULL,
  `manager_checker_name` varchar(32) DEFAULT NULL,
  `create_time` timestamp NOT NULL,
  `extraAttributes` varchar(2048) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `user_permission`
--

DROP TABLE IF EXISTS `user_permission`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_permission` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `permission_id` int NOT NULL,
  `create_time` timestamp NOT NULL,
  `extraAttributes` varchar(2048) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_level` (`user_id`,`permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `week_usage_bill`
--

DROP TABLE IF EXISTS `week_usage_bill`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `week_usage_bill` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `user_username` varchar(32) NOT NULL,
  `user_name` varchar(32) NOT NULL,
  `wall_time` int NOT NULL,
  `gwall_time` int NOT NULL,
  `fee` decimal(18,2) NOT NULL,
  `pay_fee` decimal(18,2) NOT NULL DEFAULT '0.00',
  `start_time` timestamp NOT NULL,
  `end_time` timestamp NOT NULL,
  `pay_flag` tinyint NOT NULL DEFAULT '0',
  `pay_time` timestamp NULL DEFAULT NULL,
  `pay_type` tinyint DEFAULT NULL,
  `pay_message` varchar(256) DEFAULT NULL,
  `user_group_id` int NOT NULL,
  `create_time` timestamp NOT NULL,
  `extraAttributes` varchar(2048) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-04-20 20:01:08
