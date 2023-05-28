CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v1()),
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "items" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v1()),
  "title" varchar NOT NULL,
  "buying_price" float4 NOT NULL,
  "selling_price" float4 NOT NULL,
  "quantity" bigint NOT NULL DEFAULT 1,
  "user_id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "items" ("title");

CREATE UNIQUE INDEX ON "items" ("user_id", "title");

ALTER TABLE "items" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");