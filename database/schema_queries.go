package database

var createUsersTableQuery = `
CREATE TABLE IF NOT EXISTS users (
	id int UNIQUE NOT NULL AUTO_INCREMENT,
	email varchar(255) UNIQUE NOT NULL,
	password varchar(255) NOT NULL,
	firstName varchar(255),
	lastName varchar(255),
	enabled tinyint DEFAULT '1',
	dateCreated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
)
`

var createTagsTableQuery = `
CREATE TABLE IF NOT EXISTS tags (
	id int UNIQUE NOT NULL AUTO_INCREMENT,
	name varchar(63) UNIQUE NOT NULL,
	primaryColor varchar(31) NOT NULL,
	secondaryColor varchar(31) NOT NULL,
	public tinyint DEFAULT '1',
	enabled tinyint DEFAULT '1'
)
`

var createTasksTableQuery = `
CREATE TABLE IF NOT EXISTS tasks (
	id int UNIQUE NOT NULL AUTO_INCREMENT,
	name varchar(255) NOT NULL,
	description varchar(2047),
	importance int NOT NULL,
	status int NOT NULL DEFAULT '1',
	duration int NOT NULL DEFAULT '2',
	dueDate TIMESTAMP,
	dateCreated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
)
`

var createTasksTagsTableQuery = `
CREATE TABLE IF NOT EXISTS tasks_tags (
	id int UNIQUE NOT NULL AUTO_INCREMENT,
	idTask int NOT NULL,
	idTag int NOT NULL
)
`
