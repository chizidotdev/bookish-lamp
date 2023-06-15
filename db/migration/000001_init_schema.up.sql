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

-- CREATE TABLE "dashboard" (
--   "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v1()),
--   "user_id" uuid NOT NULL,
--   "created_at" timestamptz NOT NULL DEFAULT (now()),
--   "updated_at" timestamptz NOT NULL DEFAULT (now()),
--   "total_items" bigint NOT NULL DEFAULT 0,
--   "low_stock_items" bigint NOT NULL DEFAULT 0,
--   "items_to_ship" bigint NOT NULL DEFAULT 0,
--   "recent_sales" bigint NOT NULL DEFAULT 0,
--   "sales_performance" float4 NOT NULL DEFAULT 0,
--   "pending_orders" bigint NOT NULL DEFAULT 0,
--   "notifications" varchar NOT NULL DEFAULT '',
--   "inventory_value" float4 NOT NULL DEFAULT 0,
--   "expiring_items" bigint NOT NULL DEFAULT 0
-- );

CREATE TABLE "sales" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v1()),
  "item_id" uuid NOT NULL,
  "user_id" uuid REFERENCES users(id) NOT NULL,
  "quantity_sold" bigint NOT NULL,
  "sale_price" float4 NOT NULL,
  "sale_date" timestamptz NOT NULL DEFAULT (now()),
  "customer_name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "orders" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v1()),
  "user_id" uuid REFERENCES users(id) NOT NULL,
  "order_date" timestamptz NOT NULL DEFAULT (now()),
  "total_amount" float4 NOT NULL,
  "status" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "order_items" (
  id uuid DEFAULT uuid_generate_v1() PRIMARY KEY,
  order_id uuid REFERENCES orders(id) NOT NULL,
  item_id uuid REFERENCES items(id) NOT NULL,
  quantity bigint NOT NULL,
  price float4 NOT NULL,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE "customer" (
  id uuid DEFAULT uuid_generate_v1() PRIMARY KEY,
  order_id uuid REFERENCES orders(id) NOT NULL,
  name varchar NOT NULL,
  email varchar,
  phone varchar,
  address text,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now()
);

CREATE INDEX ON "items" ("title");

CREATE UNIQUE INDEX ON "items" ("user_id", "title");

ALTER TABLE "items" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "dashboard" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "sales" ADD FOREIGN KEY ("item_id") REFERENCES "items" ("id");
