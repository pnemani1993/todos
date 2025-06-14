package dbutils

const ZERO_TIME string = "1970-01-01 00:00:00"

const CREATE_TABLE_LIST = `CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		task TEXT NOT NULL,
		description TEXT,
		done INTEGER DEFAULT 0, 
		high_priority INTEGER DEFAULT 0,
		creation_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		alert TIMESTAMP
	);`

const GET_NOT_DONE_TASKS = `SELECT * FROM tasks WHERE done = 0;`
const GET_DONE_TASKS = `SELECT * FROM tasks WHERE done = 1;`

const TRIAL_DATA string = `INSERT INTO tasks (task, description, done, high_priority, alert) VALUES
('Task 1', 'Description for task 1', 0, 0, '2001-01-01 01:00:00'),
('Task 2', 'Description for task 2', 1, 1, '2025-06-14 10:00:00'),
('Task 3', 'Description for task 3', 0, 1, '2025-06-15 09:00:00'),
('Task 4', 'Description for task 4', 1, 0, '2001-01-01 01:00:00'),
('Task 5', 'Description for task 5', 0, 0, '2025-06-16 14:30:00'),
('Task 6', 'Description for task 6', 0, 1, '2001-01-01 01:00:00'),
('Task 7', 'Description for task 7', 1, 1, '2025-06-17 12:00:00'),
('Task 8', 'Description for task 8', 0, 0, '2001-01-01 01:00:00'),
('Task 9', 'Description for task 9', 1, 0, '2025-06-18 08:00:00'),
('Task 10', 'Description for task 10', 0, 1,'2001-01-01 01:00:00');`

const DELETE_ALL_TRIAL string = "DELETE FROM tasks"
