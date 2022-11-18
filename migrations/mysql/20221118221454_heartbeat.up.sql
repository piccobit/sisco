-- modify "services" table
ALTER TABLE `services` ADD COLUMN `available` bool NOT NULL DEFAULT false, ADD COLUMN `heartbeat` timestamp NULL, ADD CONSTRAINT `services_areas_services` FOREIGN KEY (`area_services`) REFERENCES `areas` (`id`) ON DELETE SET NULL;
-- modify "service_tags" table
ALTER TABLE `service_tags` ADD CONSTRAINT `service_tags_service_id` FOREIGN KEY (`service_id`) REFERENCES `services` (`id`) ON DELETE CASCADE, ADD CONSTRAINT `service_tags_tag_id` FOREIGN KEY (`tag_id`) REFERENCES `tags` (`id`) ON DELETE CASCADE;
