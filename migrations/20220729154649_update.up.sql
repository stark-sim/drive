-- modify "objects" table
ALTER TABLE "objects" DROP COLUMN "user_objects", ADD COLUMN "user_id" bigint NOT NULL;
