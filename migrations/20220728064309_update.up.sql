-- modify "directories" table
ALTER TABLE "directories" DROP COLUMN "parent_id", ADD COLUMN "directory_children" bigint NULL;
-- modify "objects" table
ALTER TABLE "objects" ADD COLUMN "is_public" boolean NOT NULL DEFAULT true, ADD COLUMN "directory_objects" bigint NULL;
