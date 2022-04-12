CREATE TABLE users
(
    id   BIGINT UNSIGNED     NOT NULL AUTO_INCREMENT COMMENT 'user table id',
    name VARCHAR(255)        NOT NULL DEFAULT '' COMMENT 'user name',
    age  TINYINT(3) UNSIGNED NOT NULL DEFAULT '0' COMMENT 'user age',
    PRIMARY KEY (id)
)
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4
    COMMENT = 'user table';

INSERT INTO users (id, name, age)
VALUES (1, 'user_name_1', 1),
       (2, 'user_name_2', 2),
       (3, 'user_name_3', 3),
       (4, 'user_name_4', 4),
       (5, 'user_name_5', 5),
       (6, 'user_name_6', 6),
       (7, 'user_name_7', 7),
       (8, 'user_name_8', 8),
       (9, 'user_name_9', 9);