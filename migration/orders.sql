-- ----------------------------
-- Table structure for orders
-- ----------------------------
DROP TABLE IF EXISTS "public"."orders";

CREATE TABLE "public"."orders" (
  "id" SERIAL,
  "user_id" int8,
  "state" VARCHAR(10),
  "item_name" VARCHAR(80),
  "price" NUMERIC,
  "frequency_update_order" NUMERIC,
  "created_at" TIMESTAMP,
  "updated_at" TIMESTAMP
);

ALTER TABLE "public"."orders" OWNER TO "uxlkrajhzodzxm";

-- ----------------------------
-- Primary Key for table orders
-- ----------------------------
ALTER TABLE "public"."orders" ADD CONSTRAINT "order_id_pkey" PRIMARY KEY ("id");
