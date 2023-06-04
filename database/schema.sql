CREATE DATABASE cinema
USE cinema

CREATE TABLE USERS(
    id  uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL
)

CREATE TABLE ROLES(
     id  uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(255) NOT NULL
)
CREATE TABLE MOVIES (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    description varchar(255) not null,
    name VARCHAR(255) NOT NULL,
    price decimal NOT NULL,
    created_by uuid NOT NULL,
    creation_date date DEFAULT CURRENT_DATE,
FOREIGN KEY (created_by) REFERENCES USERS(id)
)
CREATE TABLE ROLES_USERS(
     id  uuid DEFAULT gen_random_uuid() PRIMARY KEY,
     user_id uuid NOT NULL,
     rol_id uuid NOT NULL,
     FOREIGN KEY (user_id) REFERENCES USERS(id),
     FOREIGN KEY (rol_id) REFERENCES ROLES(id),    
)

INSERT INTO ROLES(NAME) VALUES ('ADMIN')
INSERT INTO ROLES(NAME) VALUES ('SELLER')
INSERT INTO ROLES(NAME) VALUES ('VIEWER')   	 