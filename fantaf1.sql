CREATE TABLE user ( id INT(64) PRIMARY KEY, admin INT(1) NOT NULL DEFAULT FALSE);
CREATE TABLE fanta (code VARCHAR PRIMARY KEY, name VARCHAR NOT NULL, owner INT(64) NOT NULL, FOREIGN KEY(owner) REFERENCES user(id));
CREATE TABLE fanta_user (user INT(64) NOT NULL, fanta VARCHAR NOT NULL, PRIMARY KEY(user, fanta), FOREIGN KEY(user) REFERENCES user(id), FOREIGN KEY(fanta) REFERENCES fanta(code));
CREATE TABLE driver (id INTEGER PRIMARY KEY AUTOINCREMENT, name VARCHAR NOT NULL);
CREATE TABLE constructor (id INTEGER PRIMARY KEY AUTOINCREMENT, name VARCHAR NOT NULL);
CREATE TABLE season_driver (year INT NOT NULL, driver INT NOT NULL, constructor INT NOT NULL, PRIMARY KEY(year, driver), FOREIGN KEY (driver) REFERENCES driver(id), FOREIGN KEY (constructor) REFERENCES constructor(id));
CREATE TABLE price_driver(driver INT PRIMARY KEY, price INT NOT NULL, FOREIGN KEY (driver) REFERENCES driver(id));
CREATE TABLE race(day INT NOT NULL, month INT NOT NULL, year INT NOT NULL, name VARCHAR NOT NULL, round INT NOT NULL, PRIMARY KEY(year, round), UNIQUE(day, month, year));
CREATE TABLE race_prediction_result(race_year INT NOT NULL, race_round INT NOT NULL, vsc INT(1) NOT NULL, sc INT(1) NOT NULL, red_flag INT(1) NOT NULL, dfs INT NOT NULL, double INT(1) NOT NULL,
                poleman_win INT(1) NOT NULL, wet INT(1) NOT NULL, double_point INT NOT NULL, PRIMARY KEY (race_year, race_round), FOREIGN KEY (race_year, race_round) REFERENCES race(year, round));
CREATE TABLE team (id INTEGER PRIMARY KEY AUTOINCREMENT, name VARCHAR NOT NULL, user INT NOT NULL, fanta VARCHAR NOT NULL, UNIQUE(user, fanta), FOREIGN KEY (user) REFERENCES user(id), FOREIGN KEY (fanta) REFERENCES fanta(id));
CREATE TABLE team_driver (id INTEGER PRIMARY KEY AUTOINCREMENT, team INT NOT NULL, driver INT NOT NULL, year INT NOT NULL, round INT NOT NULL, UNIQUE (team, driver, year, round), FOREIGN KEY (team) REFERENCES team(id), FOREIGN KEY (driver) REFERENCES driver(id), FOREIGN KEY (race_year, race_round) REFERENCES race(year, round));