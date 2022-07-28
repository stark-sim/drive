-- reverse: modify "objects" table
ALTER TABLE "objects" DROP COLUMN "directory_objects", DROP COLUMN "is_public";
-- reverse: modify "directories" table
ALTER TABLE "directories" DROP COLUMN "directory_children", ADD COLUMN "parent_id" bigint NOT NULL DEFAULT 0;
