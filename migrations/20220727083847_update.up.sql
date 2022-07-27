-- modify "directories" table
ALTER TABLE "directories" ADD COLUMN "is_public" boolean NOT NULL DEFAULT true, ADD COLUMN "parent_id" bigint NOT NULL DEFAULT 0;
