CREATE TABLE "users" (
  "id" varchar NOT NULL PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "balance" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transactions" (
  "id" varchar NOT NULL PRIMARY KEY,
  "type" varchar NOT NULL,
  "date" varchar NOT NULL,
  "product_name" varchar NOT NULL,
  "seller_name" varchar NOT NULL,
  "value" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "products" (
  "id" varchar NOT NULL PRIMARY KEY,
  "title" varchar UNIQUE NOT NULL,
  "producer_name" varchar NOT NULL,
  "value" int NOT NULL
);

ALTER TABLE "products" ADD FOREIGN KEY ("producer_name") REFERENCES "users" ("name");
ALTER TABLE "transactions" ADD FOREIGN KEY ("product_name") REFERENCES "products" ("title");
ALTER TABLE "transactions" ADD FOREIGN KEY ("seller_name") REFERENCES "users" ("name");

CREATE INDEX ON "users" ("name");
CREATE INDEX ON "products" ("title");

