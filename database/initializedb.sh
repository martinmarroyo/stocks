#!/bin/bash
# A script that creates initial tables and schemas located in /opt/database/tables for the test database
#
# Test datatabase configs: 
# DB_USER=stocks 
# DB_PASS=stocks
# DB_NAME=coop-stocks

# If schemas exist already, don't bother running
SCHEMA_CHECK="SELECT CASE WHEN EXISTS(SELECT schema_name FROM information_schema.schemata WHERE schema_name='stocks') THEN 1 ELSE 0 END AS exist"
if [ $( psql -qtAX -d coop-stocks -U stocks -c "${SCHEMA_CHECK}") -eq 1 ]
    then
        echo "Data already loaded. Continuing..."
        exit 0
    else 
        echo "Database is empty. Initializing..."
fi
# Create schema
psql -qtAX -U stocks -d coop-stocks -c "CREATE SCHEMA IF NOT EXISTS stocks;"
# Create tables
TABLE_DDL=$( for file in $(ls /opt/database/tables/*.sql); do echo $file; done)
for table in $TABLE_DDL
    do
        psql -qtAX -U stocks -d coop-stocks -f $table
    done
echo "Database initialized. Ready for storage"