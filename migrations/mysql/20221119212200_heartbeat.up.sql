-- modify "services" table
ALTER TABLE `services` ADD COLUMN `available` bool NULL DEFAULT false, ADD COLUMN `heartbeat` timestamp NULL;
