#!/bin/bash
# insert some dummy data

source credentials

mysql -u $DBUSER -p$DBPASS -h $DBHOST $DBNAME < persistent_data.sql
