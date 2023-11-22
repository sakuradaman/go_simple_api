USE `dramas` ;

SET CHARSET utf8mb4;

CREATE TABLE
    IF NOT EXISTS dramas (
        id INT(11) AUTO_INCREMENT COMMENT "ID",
        name VARCHAR(20) NOT NULL COMMENT "name",
        year DATE NOT NULL COMMENT "year",
        PRIMARY KEY (id)
    ) ENGINE = INNODB COMMENT = 'drama';