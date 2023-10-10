#!/bin/bash
set -euo pipefail

setup-config() {
touch /home/nonroot/redis.conf

cat <<EOF >> /home/nonroot/redis.conf
bind 127.0.0.1
daemonize yes
EOF

echo "/home/nonroot/redis.conf"
}

start-redis() {
    redis-server $1
}

get-connection-url() {
	echo "redis://127.0.0.1:$1"
}

shutdown() {
	redis-cli shutdown
}

# Call the requested function and pass the arguments as-is
"$@"
