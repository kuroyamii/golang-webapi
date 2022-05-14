package cafeQuery

const (
	GET_ALL_FOOD_BY_TYPE = `
	SELECT f.food_id,f.name,f.image_path,f.price,t.type_name,f.description,f.stock
	FROM food f
	INNER JOIN food_type t ON t.food_type_id = f.food_type_id
	WHERE t.type_name = ?;
	`

	SEARCH_FOODS_BY_QUERY = `
	SELECT f.food_id,f.name,f.image_path,f.price,t.type_name,f.description,f.stock
	FROM food f
	INNER JOIN food_type t USING(food_type_id)
	WHERE f.name LIKE ?;
	`

	RESERVE_TABLE = `
	UPDATE seat
	SET status = TRUE
	WHERE table_id = ?
	`

	UNRESERVE_TABLE = `
	UPDATE seat
	SET status = FALSE
	WHERE table_id = ?
	`

	INSERT_CUSTOMER = `
	INSERT INTO customer(name,table_id)
	VALUES
	( ? , ? );
	`

	REMOVE_CUSTOMER = `
	DELETE FROM customer 
	WHERE customer_id = ?;
	`

	GET_ORDER_BY_CUSTOMER_ID = `
	SELECT order_id, customer_id, ordered_at, waiter_id
	FROM orders
	WHERE customer_id = ?;
	`

	GET_ORDER_DETAILS_BY_ORDER_ID = `
	SELECT details_id, order_id, food_id
	FROM order_detail
	WHERE order_id = ?;
	`

	INSERT_ORDER = `
	INSERT INTO orders(customer_id,ordered_at, waiter_id)
	VALUES ( ? ,now(), ? );
	`

	INSERT_ORDER_DETAIL = `
	INSERT INTO order_detail (order_id,food_id)
	VALUES(?,?);
	`

	INSERT_RECORD = `
	INSERT INTO records (food_id, amount)
	VALUES (?,?);
	`

	GET_FOOD_AND_RECORDS = `
	
	`

	GET_ALL_WAITER = `
	SELECT waiter_id, name
	FROM waiters;
	`

	GET_SEATS = `
	SELECT table_id, status
	FROM seat
	`

	GET_SUM_PEOPLE = `
	SELECT SUM(people) AS summary FROM
	(
	SELECT COUNT(c.customer_id) AS people
	FROM customer c 
	UNION ALL
	SELECT COUNT(w.waiter_id) AS people
	FROM waiters w
	) AS people;
	`

	GET_ALL_CUSTOMER = `
	SELECT customer_id, name, table_id
	FROM customer;
	`

	GET_CUSTOMER_BY_ID = `
	SELECT customer_id, name, table_id 
	FROM customer
	WHERE customer_id = ?;
	`

	GET_WAITER_BY_WAITER_ID = `
	SELECT waiter_id, name FROM waiters
	WHERE waiter_id = ?;
	`

	GET_FOOD_BY_ID = `
	SELECT f.food_id,f.name,f.image_path,f.price,t.type_name,f.description,f.stock
	FROM food f
	INNER JOIN food_type t ON t.food_type_id = f.food_type_id
	WHERE f.food_id = ?;
	`

	GET_WAITER_SUM = `
	SELECT COUNT(waiter_id) AS sum
	FROM waiters;
	`

	INSERT_TO_LOG = `
	INSERT INTO log(customer_id, customer_name, table_id, order_id, waiter_id, ordered_at, details_id, food_id)
	VALUES (?,?,?,?,?,?,?,?);
	`
)
