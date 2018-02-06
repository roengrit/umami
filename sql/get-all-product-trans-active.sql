select all_table.* ,product.product_type from (
			 select receive.created_at,product_id,qty,price,average_cost,receive."flag",doc_type,receive_sub.i_d,'receive_sub' as tb   
			 from receive_sub join receive on receive_sub.doc_no = receive.doc_no where receive.active and  receive_sub.active and product_id = 1 
    union all 
			 select "order".created_at,product_id,qty*-1 as qty,price,average_cost,"order"."flag",doc_type,order_sub.i_d,'order_sub' as tb   
			 from order_sub join "order" on order_sub.doc_no = "order".doc_no where "order".active and  order_sub.active and product_id = 1  
    union all 
			 select "pick_up".created_at,product_id,qty *-1 as qty,price,average_cost,"pick_up"."flag",doc_type,pick_up_sub.i_d,'pick_up_sub' as tb   
			 from pick_up_sub join "pick_up" on pick_up_sub.doc_no = "pick_up".doc_no where "pick_up".active and  pick_up_sub.active and product_id = 1  
    union all 
			 select "stock_count".created_at,product_id,qty,price,average_cost,"stock_count"."flag",doc_type,stock_count_sub.i_d,'stock_count_sub' as tb   
			 from stock_count_sub join "stock_count" on stock_count_sub.doc_no = "stock_count".doc_no where "stock_count".active and  stock_count_sub.active and product_id = 1  
 ) as all_table JOIN product on all_table.product_id = product.i_d

ORDER BY all_table.created_at,doc_type,all_table.i_d