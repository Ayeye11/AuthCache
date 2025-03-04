CREATE TABLE IF NOT EXISTS users (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(100) UNIQUE NOT NULL,
    hash_password VARCHAR(100) NOT NULL,
    firstname VARCHAR(100) NOT NULL,
    lastname VARCHAR(100) NOT NULL,
    age INT CHECK(age >= 18),
    role_id BIGINT NOT NULL,
    FOREIGN KEY (role_id) REFERENCES ac_roles(id) ON DELETE CASCADE
);