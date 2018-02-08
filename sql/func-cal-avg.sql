CREATE OR REPLACE FUNCTION "public"."avg"()
  RETURNS "pg_catalog"."trigger" AS $BODY$
BEGIN  
IF TG_OP = 'DELETE' THEN
    INSERT INTO stock_adj(product_id)
    VALUES(OLD.product_id);
    RETURN OLD;
ELSE 
  INSERT INTO stock_adj(product_id)
 VALUES(NEW.product_id);
 RETURN NEW;
END IF;

END;
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;

ALTER FUNCTION "public"."avg"() OWNER TO "postgres";