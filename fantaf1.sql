CREATE TABLE user ( id INT(64) PRIMARY KEY, admin INT(1) NOT NULL DEFAULT FALSE);
CREATE TABLE fanta (code VARCHAR PRIMARY KEY, name VARCHAR NOT NULL, owner INT(64) NOT NULL, FOREIGN KEY(owner) REFERENCES user(id));
CREATE TABLE fanta_user (user INT(64) NOT NULL, fanta VARCHAR NOT NULL, PRIMARY KEY(user, fanta), FOREIGN KEY(user) REFERENCES user(id), FOREIGN KEY(fanta) REFERENCES fanta(code));