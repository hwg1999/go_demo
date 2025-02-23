CREATE TABLE community (
    id int auto_increment PRIMARY KEY,
    community_id int UNSIGNED NOT NULL,
    community_name varchar(128) NOT NULL,
    introduction varchar(256) NOT NULL,
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT idx_community_id UNIQUE (community_id),
    CONSTRAINT idx_community_name UNIQUE (community_name)
) COLLATE = utf8mb4_general_ci;

INSERT INTO
    bluebell.community (
        id,
        community_id,
        community_name,
        introduction,
        create_time,
        update_time
    )
VALUES (
        1,
        1,
        'Go',
        'Golang',
        '2016-11-01 08:10:10',
        '2016-11-01 08:10:10'
    );

INSERT INTO
    bluebell.community (
        id,
        community_id,
        community_name,
        introduction,
        create_time,
        update_time
    )
VALUES (
        2,
        2,
        'leetcode',
        '刷题刷题刷题',
        '2020-01-01 08:00:00',
        '2020-01-01 08:00:00'
    );

INSERT INTO
    bluebell.community (
        id,
        community_id,
        community_name,
        introduction,
        create_time,
        update_time
    )
VALUES (
        3,
        3,
        'CS:GO',
        'Rush B。。。',
        '2018-08-07 08:30:00',
        '2018-08-07 08:30:00'
    );

INSERT INTO
    bluebell.community (
        id,
        community_id,
        community_name,
        introduction,
        create_time,
        update_time
    )
VALUES (
        4,
        4,
        'LOL',
        '欢迎来到英雄联盟!',
        '2016-01-01 08:00:00',
        '2016-01-01 08:00:00'
    );