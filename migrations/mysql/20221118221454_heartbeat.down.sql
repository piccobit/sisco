-- reverse: modify "service_tags" table
ALTER TABLE `service_tags` DROP FOREIGN KEY `service_tags_tag_id`, DROP FOREIGN KEY `service_tags_service_id`;
-- reverse: modify "services" table
ALTER TABLE `services` DROP FOREIGN KEY `services_areas_services`, DROP COLUMN `heartbeat`, DROP COLUMN `available`;
