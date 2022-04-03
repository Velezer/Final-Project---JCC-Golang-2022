-- MariaDB dump 10.19  Distrib 10.4.24-MariaDB, for Win64 (AMD64)
--
-- Host: localhost    Database: db_final_project
-- ------------------------------------------------------
-- Server version	10.4.24-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `cart_items`
--

DROP TABLE IF EXISTS `cart_items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `cart_items` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `product_id` bigint(20) unsigned DEFAULT NULL,
  `count` bigint(20) unsigned DEFAULT NULL,
  `cart_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `product_cart_pair` (`product_id`,`cart_id`),
  KEY `idx_cart_items_deleted_at` (`deleted_at`),
  KEY `fk_carts_cart_items` (`cart_id`),
  CONSTRAINT `fk_cart_items_product` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
  CONSTRAINT `fk_carts_cart_items` FOREIGN KEY (`cart_id`) REFERENCES `carts` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cart_items`
--

LOCK TABLES `cart_items` WRITE;
/*!40000 ALTER TABLE `cart_items` DISABLE KEYS */;
INSERT INTO `cart_items` VALUES (1,'2022-04-03 16:49:56.363','2022-04-03 16:49:56.360',NULL,1,3,1),(2,'2022-04-03 16:50:04.642','2022-04-03 16:50:04.639',NULL,2,5,1);
/*!40000 ALTER TABLE `cart_items` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `carts`
--

DROP TABLE IF EXISTS `carts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `carts` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext DEFAULT NULL,
  `is_checkout` tinyint(1) DEFAULT 0,
  `total_price` bigint(20) unsigned DEFAULT NULL,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_carts_deleted_at` (`deleted_at`),
  KEY `fk_carts_user` (`user_id`),
  CONSTRAINT `fk_carts_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `carts`
--

LOCK TABLES `carts` WRITE;
/*!40000 ALTER TABLE `carts` DISABLE KEYS */;
INSERT INTO `carts` VALUES (1,'2022-04-03 16:49:07.334','2022-04-03 16:51:26.571',NULL,'string',0,35000,6),(2,'2022-04-03 16:50:40.967','2022-04-03 17:07:41.603',NULL,'string',1,0,6);
/*!40000 ALTER TABLE `carts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `categories`
--

DROP TABLE IF EXISTS `categories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `categories` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `idx_categories_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `categories`
--

LOCK TABLES `categories` WRITE;
/*!40000 ALTER TABLE `categories` DISABLE KEYS */;
INSERT INTO `categories` VALUES (1,'2022-04-03 16:47:27.994','2022-04-03 16:47:27.994',NULL,'string');
/*!40000 ALTER TABLE `categories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_merchant`
--

DROP TABLE IF EXISTS `order_merchant`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `order_merchant` (
  `order_id` bigint(20) unsigned NOT NULL,
  `user_id` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`order_id`,`user_id`),
  KEY `fk_order_merchant_user` (`user_id`),
  CONSTRAINT `fk_order_merchant_order` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `fk_order_merchant_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_merchant`
--

LOCK TABLES `order_merchant` WRITE;
/*!40000 ALTER TABLE `order_merchant` DISABLE KEYS */;
INSERT INTO `order_merchant` VALUES (1,1),(2,1),(2,3);
/*!40000 ALTER TABLE `order_merchant` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_statuses`
--

DROP TABLE IF EXISTS `order_statuses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `order_statuses` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `idx_order_statuses_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=326 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_statuses`
--

LOCK TABLES `order_statuses` WRITE;
/*!40000 ALTER TABLE `order_statuses` DISABLE KEYS */;
INSERT INTO `order_statuses` VALUES (1,'2022-04-03 16:44:08.188','2022-04-03 16:44:08.188',NULL,'UNPAID'),(2,'2022-04-03 16:44:08.194','2022-04-03 16:44:08.194',NULL,'PAID'),(3,'2022-04-03 16:44:08.196','2022-04-03 16:44:08.196',NULL,'CANCELLED'),(4,'2022-04-03 16:44:08.201','2022-04-03 16:44:08.201',NULL,'SHIPPING'),(5,'2022-04-03 16:44:08.203','2022-04-03 16:44:08.203',NULL,'DELIVERED');
/*!40000 ALTER TABLE `order_statuses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `orders`
--

DROP TABLE IF EXISTS `orders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `orders` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `status_id` bigint(20) unsigned DEFAULT NULL,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `cart_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_orders_deleted_at` (`deleted_at`),
  KEY `fk_orders_status` (`status_id`),
  KEY `fk_orders_user` (`user_id`),
  KEY `fk_orders_cart` (`cart_id`),
  CONSTRAINT `fk_orders_cart` FOREIGN KEY (`cart_id`) REFERENCES `carts` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `fk_orders_status` FOREIGN KEY (`status_id`) REFERENCES `order_statuses` (`id`),
  CONSTRAINT `fk_orders_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `orders`
--

LOCK TABLES `orders` WRITE;
/*!40000 ALTER TABLE `orders` DISABLE KEYS */;
INSERT INTO `orders` VALUES (1,'2022-04-03 16:51:26.584','2022-04-03 16:58:19.044',NULL,3,6,1),(2,'2022-04-03 17:07:41.606','2022-04-03 17:07:41.606',NULL,1,6,2);
/*!40000 ALTER TABLE `orders` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product_category`
--

DROP TABLE IF EXISTS `product_category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `product_category` (
  `product_id` bigint(20) unsigned NOT NULL,
  `category_id` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`product_id`,`category_id`),
  KEY `fk_product_category_category` (`category_id`),
  CONSTRAINT `fk_product_category_category` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `fk_product_category_product` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product_category`
--

LOCK TABLES `product_category` WRITE;
/*!40000 ALTER TABLE `product_category` DISABLE KEYS */;
INSERT INTO `product_category` VALUES (1,1),(2,1),(3,1);
/*!40000 ALTER TABLE `product_category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `products`
--

DROP TABLE IF EXISTS `products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `products` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext DEFAULT NULL,
  `description` longtext DEFAULT NULL,
  `count` bigint(20) unsigned DEFAULT NULL,
  `price` bigint(20) unsigned DEFAULT NULL,
  `image_url` longtext DEFAULT NULL,
  `store_id` bigint(20) unsigned DEFAULT NULL,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_products_deleted_at` (`deleted_at`),
  KEY `fk_stores_products` (`store_id`),
  KEY `fk_products_user` (`user_id`),
  CONSTRAINT `fk_products_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
  CONSTRAINT `fk_stores_products` FOREIGN KEY (`store_id`) REFERENCES `stores` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` VALUES (1,'2022-04-03 16:47:27.992','2022-04-03 16:47:27.992',NULL,'string','string',100,5000,'https://csacsa.ssa',1,1),(2,'2022-04-03 16:47:46.071','2022-04-03 16:47:46.071',NULL,'sadsi','string',100,4000,'https://csacsa.ssa',1,1),(3,'2022-04-03 16:47:55.415','2022-04-03 16:47:55.415',NULL,'ssaocksapopk','string',100,2000,'https://csacsa.ssa',1,1);
/*!40000 ALTER TABLE `products` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `roles`
--

DROP TABLE IF EXISTS `roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `roles` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `idx_roles_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=131 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `roles`
--

LOCK TABLES `roles` WRITE;
/*!40000 ALTER TABLE `roles` DISABLE KEYS */;
INSERT INTO `roles` VALUES (1,'2022-04-03 16:44:08.175','2022-04-03 16:44:08.175',NULL,'USER'),(2,'2022-04-03 16:44:08.183','2022-04-03 16:44:08.183',NULL,'MERCHANT');
/*!40000 ALTER TABLE `roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `stores`
--

DROP TABLE IF EXISTS `stores`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `stores` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) NOT NULL,
  `description` longtext DEFAULT NULL,
  `address` longtext DEFAULT NULL,
  `image_url` longtext DEFAULT NULL,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `idx_stores_deleted_at` (`deleted_at`),
  KEY `fk_stores_user` (`user_id`),
  CONSTRAINT `fk_stores_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
  CONSTRAINT `chk_stores_name` CHECK (`name` <> '')
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `stores`
--

LOCK TABLES `stores` WRITE;
/*!40000 ALTER TABLE `stores` DISABLE KEYS */;
INSERT INTO `stores` VALUES (1,'2022-04-03 16:46:36.393','2022-04-03 16:46:36.393',NULL,'string','string','string','https://cacscs/das.jpg',1),(3,'2022-04-03 16:46:46.061','2022-04-03 16:46:46.061',NULL,'strsdsading','string','string','https://cacscs/das.jpg',1),(4,'2022-04-03 16:46:51.396','2022-04-03 16:46:51.396',NULL,'stwewewrsdsading','string','string','https://cacscs/das.jpg',1);
/*!40000 ALTER TABLE `stores` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `email` varchar(191) NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` longtext DEFAULT NULL,
  `address` longtext NOT NULL,
  `role_id` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`),
  UNIQUE KEY `username` (`username`),
  KEY `idx_users_deleted_at` (`deleted_at`),
  KEY `fk_users_role` (`role_id`),
  CONSTRAINT `fk_users_role` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `chk_users_username` CHECK (`username` <> ''),
  CONSTRAINT `chk_users_address` CHECK (`address` <> '')
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'2022-04-03 16:45:04.300','2022-04-03 16:45:04.300',NULL,'string@gmail.com','merchant','$2a$10$eP2TqO5Vag5qLGhA.Snb0.k4kF6dVQ/fzIfPNmZNBcUa/2FEk.Kn2','string',2),(3,'2022-04-03 16:45:17.159','2022-04-03 16:45:17.159',NULL,'stridsang@gmail.com','merchant1','$2a$10$8whj0ME/y1XXNzZW99K3mepUNm1IeJBdUNoPDOO8ixZiZQsDPwo6y','sdtring',2),(5,'2022-04-03 16:45:25.159','2022-04-03 16:45:25.159',NULL,'stridsasdasng@gmail.com','merchant2','$2a$10$KJZw5DZr6JQXEeHKqSicK.44unSP0cg13scilkUXUJkVkJ1wl8cci','sdtring',2),(6,'2022-04-03 16:48:28.476','2022-04-03 16:48:28.476',NULL,'sdasd@gsfds.cm','string','$2a$10$D66M.DDnKgRChgfWGDZ/M.z0wc7DEyE8aRzYaBFV8r0b3He.c/Ao6','string',1);
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

-- Dump completed on 2022-04-03 22:53:29
