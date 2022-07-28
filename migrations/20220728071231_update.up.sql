-- modify "directories" table
ALTER TABLE "directories" DROP COLUMN "directory_children", ADD COLUMN "parent_id" bigint NULL;
