#!/bin/bash
# insert some dummy data

source credentials

mysql -u $DBUSER -p$DBPASS -h $DBHOST $DBNAME < dummy_data.sql
