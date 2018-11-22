DROP SCHEMA IF EXISTS `book_borrow_system`;

CREATE SCHEMA book_borrow_system
    CHARACTER SET = 'utf32';

USE book_borrow_system;

DROP TABLE IF EXISTS `user_message`;
DROP TABLE IF EXISTS `applicant`;
DROP TABLE IF EXISTS `book`;
DROP TABLE IF EXISTS `user_profile`;
DROP TABLE IF EXISTS `login_authen`;

CREATE TABLE user_profile (
    pid int not null auto_increment,
    username VARCHAR(32) NOT NULL,
    nickname varchar(32),
    email VARCHAR(32) ,
    campus VARCHAR(32) ,
    student_id VARCHAR(16),
    avatar TEXT,

    lend_count INT NOT NULL DEFAULT 0,
    borrow_count INT NOT NULL DEFAULT 0,
    post_count INT NOT NULL DEFAULT 0,
    request_count INT NOT NULL DEFAULT 0,
    score INT NOT NULL DEFAULT 50,
    signup_date timestamp,

    badge VARCHAR(32),
    primary key(pid)
);

CREATE TABLE User (
    uid int not null auto_increment,
    username VARCHAR(32) NOT NULL,
    password VARCHAR(64) NOT NULL,
    email VARCHAR(32),
    PRIMARY KEY(uid)
);

create table campus(
    campus_id int not null auto_increment,
    campus_name varchar(100) not null,
    primary key(campus_id)
);


CREATE TABLE applicant_list (
    book_id INT NOT NULL,
    applicant VARCHAR(32) NOT NULL,
    applied_time TIMESTAMP NOT NULL,
    primary key(book_id)
);

CREATE TABLE user_message (
    msg_id int not null auto_increment,
    sender VARCHAR(32) NOT NULL,
    receiver VARCHAR(32) NOT NULL,
    content TEXT NOT NULL,
    sending_time timestamp not null,
    primary key(msg_id)    
);

CREATE TABLE bug_report (
    report_id int not null auto_increment,
    reporter VARCHAR(32) NOT NULL,
    content TEXT NOT NULL,
    response text,
    primary key(report_id)
 );
create table book_transaction(
	book_id int not null auto_increment,
	book_name VARCHAR(64) NOT NULL,
    book_author VARCHAR(64),
    book_description TEXT,
    book_cover TEXT,
    book_owner VARCHAR(32),
    book_borrower VARCHAR(32),
    campus VARCHAR(32) not null,

    post_expiration date not null,
    expect_return_time date not null,
    actual_return_time date ,

    post_date timestamp,
    	
    owner_rating int,
    borrower_rating int,
    owner_comment text,
    borrower_comment text,
    

    book_status VARCHAR(32) NOT NULL,
	primary key(book_id)
);
