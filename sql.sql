CREATE DATABASE IF NOT EXISTS cleango;

USE cleango;

DROP TABLE IF EXISTS `order`;
DROP TABLE IF EXISTS item;
DROP TABLE IF EXISTS order_item;

CREATE TABLE `order`(
	id int primary key auto_increment,
	payment_method varchar(50) not null,

	created_at timestamp default current_timestamp(),
	updated_at timestamp default current_timestamp()
)ENGINE = INNODB;

CREATE TABLE item(
	id int primary key auto_increment,
	name varchar(255) not null,
	price int not null,

	created_at timestamp default current_timestamp(),
	updated_at timestamp default current_timestamp()
)ENGINE = INNODB;

CREATE TABLE order_item (
	order_id int not null,
	item_id int not null,
	quantity int not null,
	price int not null,

	foreign key (order_id) references `order`(id),
	foreign key (item_id) references item(id)
) ENGINE = INNODB;