
CREATE TABLE IF NOT EXISTS "book" (
  "id" integer PRIMARY KEY,
  "available" integer,
  "created_at" timestamp,
  "title" varchar NOT NULL,
  "author" varchar NOT NULL,
  "publication" varchar NOT NULL,
  "isbn" varchar(13) UNIQUE
);

COMMENT ON COLUMN "book"."available" IS 'only positive';