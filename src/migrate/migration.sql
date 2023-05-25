CREATE TABLE db_stori_balance_notify_dev.transactions (
	user_id BINARY(16) NOT NULL,
	product varchar(100) NULL,
	code varchar(100) NULL,
	description varchar(100) NULL,
	`date` DATE NULL,
	amount DECIMAL(10,2) NULL
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;
