CREATE TABLE members (
    id INT PRIMARY KEY AUTO_INCREMENT,
    email VARCHAR(320) NOT NULL,
    provider VARCHAR(20) NOT NULL,
    avatar_id INT NOT NULL
);

CREATE TABLE avatars (
    id INT PRIMARY KEY AUTO_INCREMENT,
    nick VARCHAR(50) NOT NULL,
    profile LONGBLOB
);

CREATE TABLE games (
    id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(50) NOT NULL,
    left_option VARCHAR(50) NOT NULL,
    right_option VARCHAR(50) NOT NULL,
    left_desc VARCHAR(50),
    right_desc VARCHAR(50),
    avatar_id INT NOT NULL
);

CREATE TABLE votes (
    id INT PRIMARY KEY AUTO_INCREMENT,
    game_id INT NOT NULL,
    avatar_id INT NOT NULL,
    pick BOOLEAN NOT NULL
);

CREATE TABLE comments (
    id INT PRIMARY KEY AUTO_INCREMENT,
    game_id INT NOT NULL,
    avatar_id INT NOT NULL,
    parent_id INT,
    content VARCHAR(100),
    deleted BOOLEAN NOT NULL
);

CREATE TABLE likes (
    id INT PRIMARY KEY AUTO_INCREMENT,
    avatar_id INT NOT NULL,
    game_id INT,
    comment_id INT
);
