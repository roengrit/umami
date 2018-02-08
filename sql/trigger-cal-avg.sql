CREATE TRIGGER "order_avg" BEFORE INSERT OR UPDATE OF "flag", "product_id", "qty", "price", "active", "doc_date" OR DELETE ON "public"."order_sub"
FOR EACH ROW
EXECUTE PROCEDURE "public"."avg"();

CREATE TRIGGER "pick_up_avg" BEFORE INSERT OR UPDATE OF "product_id", "flag", "qty", "price", "active", "doc_date" OR DELETE ON "public"."pick_up_sub"
FOR EACH ROW
EXECUTE PROCEDURE "public"."avg"();

CREATE TRIGGER "receive_avg" BEFORE INSERT OR UPDATE OF "product_id", "flag", "qty", "price", "active", "doc_date" OR DELETE ON "public"."receive_sub"
FOR EACH ROW
EXECUTE PROCEDURE "public"."avg"();

CREATE TRIGGER "stock_avg" BEFORE INSERT OR UPDATE OF "product_id", "flag", "qty", "price", "active", "doc_date" OR DELETE ON "public"."stock_count_sub"
FOR EACH ROW
EXECUTE PROCEDURE "public"."avg"();