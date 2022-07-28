-- reverse: modify "directories" table
ALTER TABLE "directories" DROP COLUMN "parent_id", ADD COLUMN "directory_children" bigint NULL;
