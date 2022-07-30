-- reverse: modify "objects" table
ALTER TABLE "objects" DROP COLUMN "user_id", ADD COLUMN "user_objects" bigint NULL;
