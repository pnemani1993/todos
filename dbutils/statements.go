package dbutils

const CREATE_TABLE_LIST = `CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		task TEXT NOT NULL,
		description TEXT,
		done INTEGER DEFAULT 0, 
		high_priority INTEGER DEFAULT 0,
		creation_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

const CREATE_TABLE_REMINDER = `CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		task TEXT NOT NULL,
		done INTEGER DEFAULT 0, 
		alert TIMESTAMP
	);`

const SELECT_NOT_DONE_TASKS = `SELECT * FROM tasks WHERE done = 0;`
const SELECT_DONE_TASKS = `SELECT * FROM tasks WHERE done = 1;`
const SELECT_ALL_HIGH_PRIORITY_TASKS string = `SELECT * FROM tasks WHERE high_priority = 1;`

const TRIAL_DATA string = `INSERT INTO tasks (task, description, done, high_priority) VALUES
('Task 1', 'Description for task 1', 0, 0),
('Task 2', 'Description for task 2', 1, 1),
('Task 3', 'Description for task 3', 0, 1),
('Task 4', 'Description for task 4', 1, 0),
('Task 5', 'Description for task 5', 0, 0),
('Task 6', 'Description for task 6', 0, 1),
('Task 7', 'Description for task 7', 1, 1),
('Task 8', 'Description for task 8', 0, 0),
('Task 9', 'Description for task 9', 1, 0),
('Task 10', 'Description for task 10', 0, 1);`

const DELETE_ALL_TRIAL string = "DELETE FROM tasks"

const DELETE_TABLE string = "DROP TABLE tasks"
