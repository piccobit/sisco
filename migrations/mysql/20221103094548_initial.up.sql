-- create "areas" table
CREATE TABLE `areas` (`id` bigint NOT NULL AUTO_INCREMENT, `name` varchar(255) NOT NULL, `description` varchar(255) NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `name` (`name`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- create "services" table
CREATE TABLE `services` (`id` bigint NOT NULL AUTO_INCREMENT, `name` varchar(255) NOT NULL, `description` varchar(255) NULL DEFAULT '', `protocol` varchar(255) NOT NULL, `host` varchar(255) NOT NULL, `port` varchar(255) NOT NULL, `area_services` bigint NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- create "tags" table
CREATE TABLE `tags` (`id` bigint NOT NULL AUTO_INCREMENT, `name` varchar(255) NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `name` (`name`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- create "tokens" table
CREATE TABLE `tokens` (`id` bigint NOT NULL AUTO_INCREMENT, `user` varchar(255) NOT NULL, `token` varchar(255) NOT NULL, `created` timestamp NULL, `admin` bool NOT NULL DEFAULT false, PRIMARY KEY (`id`), UNIQUE INDEX `user` (`user`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- create "service_tags" table
CREATE TABLE `service_tags` (`service_id` bigint NOT NULL, `tag_id` bigint NOT NULL, PRIMARY KEY (`service_id`, `tag_id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
