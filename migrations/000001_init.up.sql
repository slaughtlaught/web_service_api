CREATE TABLE users(
	id 		char(20) PRIMARY KEY,
	name 	varchar(50) NOT NULL,
	email	varchar(50) UNIQUE NOT NULL,
	hashed_password varchar(100) NOT NULL
);

CREATE INDEX idx_users_id ON users(id);
CREATE INDEX idx_users_email ON users(email);


CREATE TABLE notes(
	id		 char(20) PRIMARY KEY,
	title    varchar(80) NOT NULL,         
	content  text NOT NULL,
	user_id  char(20) REFERENCES users(id) NOT NULL 
);

