CREATE TABLE `weekday` (
	`id` smallint(4) unsigned NOT NULL,
	`name` varchar(10) NOT NULL,
	PRIMARY KEY (`id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE `day` (
	`id` smallint(4) unsigned NOT NULL AUTO_INCREMENT,
	`date` DATE NOT NULL,
	`weekday_id` smallint(4) unsigned NOT NULL,
	PRIMARY KEY (`id`),
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

CREATE TABLE `branch` (
	`id` smallint(4) unsigned NOT NULL AUTO_INCREMENT,
	`name` varchar(200) NOT NULL,
	`location` varchar(200) NOT NULL,
	PRIMARY KEY (`id`),
	UNIQUE(`name`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE `counter` (
	`id` smallint(4) unsigned NOT NULL AUTO_INCREMENT,
	`branch_id` smallint(4) unsigned NOT NULL,
	`name` varchar(200) NOT NULL,
	PRIMARY KEY (`id`),
	CONSTRAINT FOREIGN KEY (`branch_id`) REFERENCES `branch` (`id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE `businesshour` (
	`id` smallint(4) unsigned NOT NULL AUTO_INCREMENT,
	`opening` TIME NOT NULL,
	`closing` TIME NOT NULL,
	UNIQUE(`opening`,`closing`),
	PRIMARY KEY (`id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE `counter_businesshour_weekday` (
	`counter_id` smallint(4) unsigned NOT NULL,
	`businesshour_id` smallint(4) unsigned NOT NULL,
	`weekday_id` smallint(4) unsigned NOT NULL,
	PRIMARY KEY (`counter_id`,`businesshour_id`,`weekday_id`),
	CONSTRAINT FOREIGN KEY (`counter_id`) REFERENCES `counter` (`id`),
	CONSTRAINT FOREIGN KEY (`businesshour_id`) REFERENCES `businesshour` (`id`),
	CONSTRAINT FOREIGN KEY (`weekday_id`) REFERENCES `weekday` (`id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
