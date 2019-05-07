set -ex

PROJECT_ID="bookery"
PROJECT_ROOT_DIR="/bookery"
APP_DIR="$PROJECT_ROOT_DIR/app"
LOGS_DIR="/bookery/logs"
PROJECT_URL="https://storage.googleapis.com/bookery-binary/bookery.tar"
BINARY_FILE_PATH="$APP_DIR/$PROJECT_ID"
ENVIRONMENT="production"
DOMAIN_NAME="bookery"

# Sysctl Tweaks
cat >/etc/sysctl.conf << EOF
### IMPROVE SYSTEM MEMORY MANAGEMENT ###

# Increase size of file handles and inode cache
fs.file-max = 10000000

# Do less swapping
vm.swappiness = 10
vm.dirty_ratio = 60
vm.dirty_background_ratio = 2

### GENERAL NETWORK SECURITY OPTIONS ###

# Number of times SYNACKs for passive TCP connection.
net.ipv4.tcp_synack_retries = 2

# Allowed local port range
net.ipv4.ip_local_port_range = 2000 65535

# Protect Against TCP Time-Wait
net.ipv4.tcp_rfc1337 = 1

# Decrease the time default value for tcp_fin_timeout connection
net.ipv4.tcp_fin_timeout = 15

# Decrease the time default value for connections to keep alive
net.ipv4.tcp_keepalive_time = 60
net.ipv4.tcp_keepalive_intvl = 10
net.ipv4.tcp_keepalive_probes = 6

### TUNING NETWORK PERFORMANCE ###

# Default Socket Receive Buffer
net.core.rmem_default = 31457280

# Maximum Socket Receive Buffer
net.core.rmem_max = 12582912

# Default Socket Send Buffer
net.core.wmem_default = 31457280

# Maximum Socket Send Buffer
net.core.wmem_max = 12582912

# Increase number of incoming connections
net.core.somaxconn = 4096

# Increase number of incoming connections backlog
net.core.netdev_max_backlog = 65536

# Increase the maximum amount of option memory buffers
net.core.optmem_max = 25165824

# Increase the maximum total buffer-space allocatable
# This is measured in units of pages (4096 bytes)
net.ipv4.tcp_mem = 65536 131072 262144
net.ipv4.udp_mem = 65536 131072 262144

# Increase the read-buffer space allocatable
net.ipv4.tcp_rmem = 8192 87380 16777216
net.ipv4.udp_rmem_min = 16384

# Increase the write-buffer-space allocatable
net.ipv4.tcp_wmem = 8192 65536 16777216
net.ipv4.udp_wmem_min = 16384

# Increase the tcp-time-wait buckets pool size to prevent simple DOS attacks
# This may cause dropped frames with load-balancing and NATs, 
# only use this for a server that communicates only over your local network.
#net.ipv4.tcp_max_tw_buckets = 1440000
#net.ipv4.tcp_tw_recycle = 1
#net.ipv4.tcp_tw_reuse = 1
EOF
sysctl -p
cat >/etc/security/limits.conf << EOF
*         hard    nofile      1048000
*         soft    nofile      1048000
root      hard    nofile      1048000
root      soft    nofile      1048000
EOF
ulimit -n 1048000
ulimit -Hn 1048000

# Install dependencies from apt
sudo apt-get update -yq && sudo apt-get upgrade -yq
apt-get autoremove -yq
apt-get install -yq runit

# UFW Firewall
sudo ufw allow ssh
sudo ufw allow http
sudo ufw allow https
sudo ufw allow 28015
sudo ufw allow 29015
sudo ufw allow 8080
sudo ufw allow 5432

cat >/etc/ufw/before.rules << EOF
#
# rules.before
#
# Rules that should be run before the ufw command line added rules. Custom
# rules should be added to one of these chains:
#   ufw-before-input
#   ufw-before-output
#   ufw-before-forward
#

# redirect port 443 to 80
*nat
:PREROUTING ACCEPT [0:0]
-A PREROUTING -p tcp --dport 443 -j REDIRECT --to-port 80
COMMIT

# Don't delete these required lines, otherwise there will be errors
*filter
:ufw-before-input - [0:0]
:ufw-before-output - [0:0]
:ufw-before-forward - [0:0]
:ufw-not-local - [0:0]
# End required lines


# allow all on loopback
-A ufw-before-input -i lo -j ACCEPT
-A ufw-before-output -o lo -j ACCEPT

# quickly process packets for which we already have a connection
-A ufw-before-input -m conntrack --ctstate RELATED,ESTABLISHED -j ACCEPT
-A ufw-before-output -m conntrack --ctstate RELATED,ESTABLISHED -j ACCEPT
-A ufw-before-forward -m conntrack --ctstate RELATED,ESTABLISHED -j ACCEPT

# drop INVALID packets (logs these in loglevel medium and higher)
-A ufw-before-input -m conntrack --ctstate INVALID -j ufw-logging-deny
-A ufw-before-input -m conntrack --ctstate INVALID -j DROP

# ok icmp codes for INPUT
-A ufw-before-input -p icmp --icmp-type destination-unreachable -j ACCEPT
-A ufw-before-input -p icmp --icmp-type source-quench -j ACCEPT
-A ufw-before-input -p icmp --icmp-type time-exceeded -j ACCEPT
-A ufw-before-input -p icmp --icmp-type parameter-problem -j ACCEPT
-A ufw-before-input -p icmp --icmp-type echo-request -j ACCEPT

# ok icmp code for FORWARD
-A ufw-before-forward -p icmp --icmp-type destination-unreachable -j ACCEPT
-A ufw-before-forward -p icmp --icmp-type source-quench -j ACCEPT
-A ufw-before-forward -p icmp --icmp-type time-exceeded -j ACCEPT
-A ufw-before-forward -p icmp --icmp-type parameter-problem -j ACCEPT
-A ufw-before-forward -p icmp --icmp-type echo-request -j ACCEPT

# allow dhcp client to work
-A ufw-before-input -p udp --sport 67 --dport 68 -j ACCEPT

#
# ufw-not-local
#
-A ufw-before-input -j ufw-not-local

# if LOCAL, RETURN
-A ufw-not-local -m addrtype --dst-type LOCAL -j RETURN

# if MULTICAST, RETURN
-A ufw-not-local -m addrtype --dst-type MULTICAST -j RETURN

# if BROADCAST, RETURN
-A ufw-not-local -m addrtype --dst-type BROADCAST -j RETURN

# all other non-local packets are dropped
-A ufw-not-local -m limit --limit 3/min --limit-burst 10 -j ufw-logging-deny
-A ufw-not-local -j DROP

# allow MULTICAST mDNS for service discovery (be sure the MULTICAST line above
# is uncommented)
-A ufw-before-input -p udp -d 224.0.0.251 --dport 5353 -j ACCEPT

# allow MULTICAST UPnP for service discovery (be sure the MULTICAST line above
# is uncommented)
-A ufw-before-input -p udp -d 239.255.255.250 --dport 1900 -j ACCEPT

# don't delete the 'COMMIT' line or these rules won't be processed
COMMIT
EOF
sudo ufw disable
sudo ufw enable

# Get the application tar from the GCS bucket.
if [ ! -d $PROJECT_ROOT_DIR ]; then
mkdir $PROJECT_ROOT_DIR
fi
curl -H 'Cache-Control: no-cache' -o "$PROJECT_ROOT_DIR/$PROJECT_ID.tar" $PROJECT_URL
if [ -d $APP_DIR ]; then
   if [ -L $APP_DIR ]; then 
      rm $APP_DIR
   else 
      rm -rf $APP_DIR
   fi
fi
mkdir $APP_DIR
if [ ! -d $LOGS_DIR ]; then
mkdir $LOGS_DIR
fi
tar -x -f "$PROJECT_ROOT_DIR/$PROJECT_ID.tar" -C $APP_DIR
chmod +x $BINARY_FILE_PATH

sed -i "s/domain-name/$DOMAIN_NAME/g" "$APP_DIR/$ENVIRONMENT.json"

# Configure runit to run the Go app.
if [ ! -d /etc/service ]; then
mkdir /etc/service
fi
if [ ! -d "/etc/service/$PROJECT_ID" ]; then
mkdir "/etc/service/$PROJECT_ID"
fi
if [ ! -d "/etc/service/$PROJECT_ID/log" ]; then
mkdir "/etc/service/$PROJECT_ID/log"
fi

cat >"/etc/service/$PROJECT_ID/run" << EOF
#!/bin/sh -e
exec 2>&1
ulimit -n 1048000
ulimit -Hn 1048000
exec chpst $BINARY_FILE_PATH --root=$APP_DIR --env=$ENVIRONMENT
EOF
chmod +x "/etc/service/$PROJECT_ID/run"
cat >"/etc/service/$PROJECT_ID/log/run" << EOF
#!/bin/sh -e
exec 2>&1
ulimit -n 1048000
ulimit -Hn 1048000
exec chpst svlogd -tt $LOGS_DIR
EOF
chmod +x "/etc/service/$PROJECT_ID/log/run"
sleep 5s
sv stop $PROJECT_ID
sv start $PROJECT_ID
# Application should now be running under runit