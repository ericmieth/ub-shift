#!/bin/bash
# reset the database (drop, create scheme, insert dummy data)

./drop_database.sh
./create_database.sh
./insert_persistent_data.sh
./insert_dummy_data.sh
