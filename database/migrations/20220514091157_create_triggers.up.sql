DROP TRIGGER IF EXISTS record_after_insert;
CREATE TRIGGER record_after_insert
AFTER INSERT ON food
FOR EACH ROW BEGIN
	INSERT INTO records(food_id, amount)
	VALUES (NEW.food_id,0);
END;