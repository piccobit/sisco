-- reverse: modify "service_tags" table
ALTER TABLE "service_tags" DROP CONSTRAINT "service_tags_tag_id", DROP CONSTRAINT "service_tags_service_id";
-- reverse: modify "services" table
ALTER TABLE "services" DROP CONSTRAINT "services_areas_services", DROP COLUMN "heartbeat", DROP COLUMN "available";
