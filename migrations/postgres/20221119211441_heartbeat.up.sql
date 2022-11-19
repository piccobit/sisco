-- modify "services" table
ALTER TABLE "services" ADD COLUMN "available" boolean NULL DEFAULT false, ADD COLUMN "heartbeat" timestamptz NULL;
