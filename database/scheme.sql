CREATE TABLE `weekday` (
	`id` smallint(4) unsigned NOT NULL,
	`name` varchar(10) NOT NULL,
	PRIMARY KEY (`id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

INSERT INTO `weekday` (id, name) VALUES 
(0, 'Sonntag'),
(1, 'Montag'),
(2, 'Dienstag'),
(3, 'Mittwoch'),
(4, 'Donnerstag'),
(5, 'Freitag'),
(6, 'Samstag');

CREATE TABLE `day` (
	`date` DATE NOT NULL,
	`weekday_id` smallint(4) unsigned NOT NULL,
	PRIMARY KEY (`date`),
	CONSTRAINT FOREIGN KEY (`weekday_id`) REFERENCES `weekday` (`id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE `employee` (
	`id` smallint(4) unsigned NOT NULL AUTO_INCREMENT,
	`firstname` varchar(80) NOT NULL,
	`lastname` varchar(80) NOT NULL,
	`manager` boolean NOT NULL DEFAULT 0,
	`mailaddress` varchar(80) NOT NULL,
	`active` boolean NOT NULL DEFAULT 0,
	`workinghours` DECIMAL(4,2),
	PRIMARY KEY (`id`),
	UNIQUE(`mailaddress`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;



