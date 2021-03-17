CREATE TABLE `insert_log` (
  `id` int NOT NULL AUTO_INCREMENT,
  `type` int DEFAULT '0',
  `goroutine` int DEFAULT '0',
  `task_start` int DEFAULT '0',
  `task_end` int DEFAULT '0',
  `description` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci