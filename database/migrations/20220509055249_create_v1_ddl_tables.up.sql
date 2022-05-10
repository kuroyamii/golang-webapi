CREATE TABLE IF NOT EXISTS seat(
	table_id INT NOT NULL AUTO_INCREMENT,
	status TINYINT,
	PRIMARY KEY (table_id)
);


CREATE TABLE IF NOT EXISTS customer(
	customer_id BIGINT NOT NULL AUTO_INCREMENT,
	name VARCHAR(50),
	table_id INT,
	PRIMARY KEY (customer_id),
	FOREIGN KEY (table_id) REFERENCES seat(table_id)
	ON DELETE CASCADE
	ON UPDATE NO ACTION
);

CREATE TABLE IF NOT EXISTS orders(
	order_id BIGINT NOT NULL AUTO_INCREMENT,
	customer_id BIGINT,
	ordered_at DATETIME,
	PRIMARY KEY (order_id),
	FOREIGN KEY (customer_id) REFERENCES customer(customer_id)
	ON DELETE CASCADE
	ON UPDATE NO ACTION
);

CREATE TABLE IF NOT EXISTS food_type(
	food_type_id INT NOT NULL AUTO_INCREMENT,
	type_name VARCHAR(15),
	PRIMARY KEY (food_type_id)
);


CREATE TABLE IF NOT EXISTS food(
	food_id INT NOT NULL AUTO_INCREMENT,
	name VARCHAR(50),
	image_path TEXT,
	price INT,
	food_type_id INT,
	description TEXT,
	stock INT,
	PRIMARY KEY (food_id),
	FOREIGN KEY (food_type_id) REFERENCES food_type(food_type_id)
	ON DELETE CASCADE
	ON UPDATE NO ACTION
);

CREATE TABLE IF NOT EXISTS order_detail(
	details_id BIGINT NOT NULL AUTO_INCREMENT,
	order_id BIGINT,
	food_id INT,
	PRIMARY KEY (details_id),
	FOREIGN KEY (order_id) REFERENCES orders(order_id)
	ON DELETE CASCADE
	ON UPDATE NO ACTION,
	FOREIGN KEY (food_id) REFERENCES food(food_id)
	ON DELETE CASCADE
	ON UPDATE NO ACTION
);

CREATE TABLE IF NOT EXISTS records(
	record_id BIGINT NOT NULL AUTO_INCREMENT,
	food_id INT,
	amount INT,
	PRIMARY KEY (record_id),
	FOREIGN KEY (food_id) REFERENCES food(food_id)
	ON DELETE CASCADE
	ON UPDATE NO ACTION
);