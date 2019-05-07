set -ex

MASTER_SERVER_IP="193.70.3.25"
REPLICA_USER_PASSWORD="Nr:15/5/95"

cat >/var/lib/postgresql/9.5/main/recovery.conf << EOF
standby_mode = 'on'
primary_conninfo = 'host=$MASTER_SERVER_IP port=5432 user=replica password=$REPLICA_USER_PASSWORD'
restore_command = 'cp /postgres/archive/%f %p'
trigger_file = '/tmp/postgresql.trigger.5432'
EOF