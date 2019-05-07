# Postgres Master And Slave Setup

## Install Postgres
-- install postgres packages
sudo apt-get install postgresql postgresql-client postgresql-contrib

-- change postgres password
sudo passwd postgres

## Master Server

-- create replica user
CREATE USER replica REPLICATION LOGIN ENCRYPTED PASSWORD 'Nr:15/5/95';

-- verify the new replica user
\du

-- postgresql.conf and pg_hba.conf
run postgres-master.sh

## Slave Server

-- postgresql.conf
run postgres-slave.sh

-- stop postgres service
sudo systemctl stop postgresql

-- switch to postgres user
sudo -i -u postgres

-- remove current data
rm -rf /var/lib/postgresql/9.5/main

-- copy data from master
pg_basebackup -h 193.70.3.25 -D /var/lib/postgresql/9.5/main -U replica -v -P

-- recovery.conf
run postgres-recovery.sh

-- start postgres service
sudo systemctl start postgresql

## Test Replication (Master Server)

-- switch to postgres user
sudo -i -u postgres

psql -x -c "select * from pg_stat_replication;"
