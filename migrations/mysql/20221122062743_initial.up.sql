-- create "areas" table
CREATE TABLE `areas` (`id` bigint NOT NULL AUTO_INCREMENT, `name` varchar(255) NOT NULL, `description` varchar(255) NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `name` (`name`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- create "tokens" table
CREATE TABLE `tokens` (`id` bigint NOT NULL AUTO_INCREMENT, `user` varchar(255) NOT NULL, `token` varchar(255) NOT NULL, `created` timestamp NULL, `permissions` bigint unsigned NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `user` (`user`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- create "services" table
CREATE TABLE `services` (`id` bigint NOT NULL AUTO_INCREMENT, `name` varchar(255) NOT NULL, `description` varchar(255) NULL DEFAULT '', `protocol` varchar(255) NOT NULL, `host` varchar(255) NOT NULL, `port` varchar(255) NOT NULL, `available` bool NOT NULL DEFAULT false, `heartbeat` timestamp NULL, `area_services` bigint NULL, PRIMARY KEY (`id`), CONSTRAINT `services_areas_services` FOREIGN KEY (`area_services`) REFERENCES `areas` (`id`) ON DELETE SET NULL) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- create "tags" table
CREATE TABLE `tags` (`id` bigint NOT NULL AUTO_INCREMENT, `name` varchar(255) NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `name` (`name`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- create "service_tags" table
CREATE TABLE `service_tags` (`service_id` bigint NOT NULL, `tag_id` bigint NOT NULL, PRIMARY KEY (`service_id`, `tag_id`), CONSTRAINT `service_tags_service_id` FOREIGN KEY (`service_id`) REFERENCES `services` (`id`) ON DELETE CASCADE, CONSTRAINT `service_tags_tag_id` FOREIGN KEY (`tag_id`) REFERENCES `tags` (`id`) ON DELETE CASCADE) CHARSET utf8mb4 COLLATE utf8mb4_bin;
