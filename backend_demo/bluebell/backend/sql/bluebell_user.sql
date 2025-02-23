CREATE TABLE user (
    id bigint auto_increment PRIMARY KEY,
    user_id bigint NOT NULL,
    username varchar(64) NOT NULL,
    password varchar(64) NOT NULL,
    email varchar(64) NULL,
    gender tinyint DEFAULT 0 NOT NULL,
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT idx_user_id UNIQUE (user_id),
    CONSTRAINT idx_username UNIQUE (username)
) COLLATE = utf8mb4_general_ci;

INSERT INTO
    bluebell.user (
        id,
        user_id,
        username,
        password,
        email,
        gender,
        create_time,
        update_time
    )
VALUES (
        1,
        28018727488323585,
        'q1mi',
        '313233343536639a9119599647d841b1bef6ce5ea293',
        NULL,
        0,
        '2020-07-12 07:01:03',
        '2020-07-12 07:01:03'
    );

INSERT INTO
    bluebell.user (
        id,
        user_id,
        username,
        password,
        email,
        gender,
        create_time,
        update_time
    )
VALUES (
        2,
        4183532125556736,
        '七米',
        '313233639a9119599647d841b1bef6ce5ea293',
        NULL,
        0,
        '2020-07-12 13:03:51',
        '2020-07-12 13:03:51'
    );