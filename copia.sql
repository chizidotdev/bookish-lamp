CREATE TABLE "items" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v1()),
  "title" varchar NOT NULL,
  "buying_price" float4 NOT NULL,
  "selling_price" float4 NOT NULL,
  "quantity" bigint NOT NULL DEFAULT 1,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "items" ("title");
