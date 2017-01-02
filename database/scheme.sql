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
