
CREATE TABLE IF NOT EXISTS "book" (
  "id" serial PRIMARY KEY,
  "available" integer,
  "added_at" timestamp,
  "title" varchar NOT NULL,
  "author" varchar NOT NULL,
  "publication" varchar NOT NULL,
  "isbn" bigint UNIQUE NOT NULL
);

COMMENT ON COLUMN "book"."available" IS 'only positive';