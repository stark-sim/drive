-- reverse: modify "directories" table
ALTER TABLE "directories" DROP COLUMN "parent_id", DROP COLUMN "is_public";
