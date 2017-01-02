#!/bin/bash
# drop all tables in database

source credentials

mysql -u $DBUSER -p$DBPASS -h $DBHOST -Nse 'show tables' $DBNAME | \
	while read table;
	do mysql -u $DBUSER -p$DBPASS -h $DBHOST -e "drop table $table" $DBNAME;
	done
