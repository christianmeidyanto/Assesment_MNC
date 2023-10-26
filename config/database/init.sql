CREATE DATABASE payment_api;

CREATE TABLE m_customer(
    customer_id varchar(100) primary key not null,
    username varchar(100) unique not null ,
    password varchar(100) not null,
    customer_name varchar(100) not null,
    customer_address varchar(200) not null ,
    customer_amount bigint default 0
)ENGINE=InnoDb;

CREATE TABLE t_transaction(
    transaction_id varchar(100) primary key not null,
    transaction_date date not null,
    sender_id varchar(100) not null ,
    reciever_id varchar(100) not null ,
    transaction_amount bigint not null ,
    FOREIGN KEY (sender_id) REFERENCES m_customer(customer_id)
    FOREIGN KEY (reciever_id) REFERENCES m_customer(customer_id)
)ENGINE=InnoDb;