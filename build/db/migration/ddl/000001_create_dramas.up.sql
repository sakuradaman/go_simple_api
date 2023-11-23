USE `dramas` ;

SET CHARSET utf8mb4;

CREATE TABLE
    IF NOT EXISTS dramas (
        id INT(11) AUTO_INCREMENT COMMENT "ID",
        name VARCHAR(20) NOT NULL COMMENT "name",
        year INT(11) NOT NULL COMMENT "year",
        created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT "created date",
        updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT "updated date",
        deleted_at DATETIME COMMENT "drama deleted date",
        PRIMARY KEY (id)
    ) ENGINE = INNODB COMMENT = 'dramas';