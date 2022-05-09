CREATE TABLE seat(
	table_id INT,
	status TINYINT,
	PRIMARY KEY (table_id)
);


CREATE TABLE customer(
	customer_id BIGINT NOT NULL AUTO_INCREMENT,
	name VARCHAR(50),
	table_id INT,
	PRIMARY KEY (customer_id),
	FOREIGN KEY (table_id) REFERENCES seat(table_id)
);

CREATE TABLE orders(
	order_id BIGINT NOT NULL AUTO_INCREMENT,
	customer_id BIGINT,
	ordered_at DATETIME,
	PRIMARY KEY (order_id),
	FOREIGN KEY (customer_id) REFERENCES customer(customer_id)
);


CREATE TABLE food(
	food_id INT NOT NULL AUTO_INCREMENT,
	name VARCHAR(50),
	price INT,
	food_type VARCHAR(10),
	description TEXT,
	stock INT,
	PRIMARY KEY (food_id)
);

CREATE TABLE order_detail(
	details_id BIGINT NOT NULL AUTO_INCREMENT,
	order_id BIGINT,
	food_id INT,
	PRIMARY KEY (details_id),
	FOREIGN KEY (order_id) REFERENCES orders(order_id),
	FOREIGN KEY (food_id) REFERENCES food(food_id)
);

CREATE TABLE records(
	record_id BIGINT NOT NULL AUTO_INCREMENT,
	food_id INT,
	amount INT,
	PRIMARY KEY (record_id),
	FOREIGN KEY (food_id) REFERENCES food(food_id)
);