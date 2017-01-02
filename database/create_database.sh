#!/bin/bash
# import the database scheme from file

source credentials

mysql -u $DBUSER -p$DBPASS -h $DBHOST $DBNAME < scheme.sql
