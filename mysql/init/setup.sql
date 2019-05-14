USE  nuxt_vue_go_chat;

/*
Create users table. It has 'id' which has a unique identity, 
'name' with the length of 30 characters, 'session' id with the 
length of 36 characters, 'password' with the length of 64 characters,
created time and updated time. Primary key is 'id'.
*/
CREATE TABLE IF NOT EXISTS users (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
    name VARCHAR(30) NOT NULL,
    session_id VARCHAR(36) NOT NULL,
    password VARCHAR(64) NOT NULL,
    created_at DATETIME DEFAULT NULL,
    updated_at DATETIME DEFAULT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4; 

/*
Create sessions table. It has 'id' with the length 
of 36 characters, 'user id', created time, and 
updated time. Primary key is 'id'.
*/
CREATE TABLE IF NOT EXISTS sessions (
    id VARCHAR(36) NOT NULL,
    user_id INT UNSIGNED NOT NULL,
    created_at DATETIME DEFAULT NULL,
    updated_at DATETIME DEFAULT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*
Create threads table. It has 'id' which has
a unique identity, 'time' with the length of 
20 characters, user id and created time, 
updated time. Primary key is 'id'.
*/
CREATE TABLE IF NOT EXISTS threads (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
    title VARCHAR(20) NOT NULL,
    user_id INT UNSIGNED NOT NULL,
    created_at DATETIME DEFAULT NULL,
    updated_at DATETIME DEFAULT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY (title)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*
Create comments. It has 'id' which has unique 
character, thread id, user id, content with the 
length of 200 characters, created time, and 
updated time. Primary key is 'id'.
*/
CREATE TABLE IF NOT EXISTS comments (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
    thread_id INT UNSIGNED NOT NULL,
    user_id INT UNSIGNED NOT NULL,
    content VARCHAR(200) NOT NULL,
    created_at DATETIME DEFAULT NULL,
    updated_at DATETIME DEFAULT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

