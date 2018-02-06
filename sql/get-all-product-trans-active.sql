select all_table.* ,product.product_type from (
     select created_at,product_id,qty,price,average_cost,"flag",i_d,'receive_sub' as tb 
     from receive_sub where active  and (select active from receive where receive_sub.doc_no = receive.doc_no limit 1) and product_id = 1 
    union all 
     select created_at,product_id,qty,price,average_cost,"flag",i_d,'order_sub' as tb 
     from order_sub where active  and (select active from "order" where order_sub.doc_no = "order".doc_no limit 1)  and product_id = 1
    union all 
     select created_at,product_id,qty,price,average_cost,"flag",i_d,'pick_up_sub' as tb 
     from pick_up_sub where active  and (select active from pick_up where pick_up_sub.doc_no = pick_up.doc_no limit 1) and product_id = 1
    union all 
     select created_at,product_id,qty,price,average_cost,"flag",i_d,'stock_count_sub' as tb 
     from stock_count_sub where active  and (select active from stock_count where stock_count_sub.doc_no = stock_count.doc_no limit 1) and product_id = 1
) as all_table JOIN product on all_table.product_id = product.i_d

ORDER BY all_table.created_at