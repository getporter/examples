#!/bin/bash
set -euo pipefail

set-config() {
    echo "Creating otel-jaeger-config volume"
    docker volume create otel-jaeger-config || true
    
    echo "Copying otel config into volume..."
    docker rm otel-jaeger-helper 2&> /dev/null || true
    docker run -v otel-jaeger-config:/config --name otel-jaeger-helper busybox true
    trap "docker rm otel-jaeger-helper" EXIT

    docker cp config.yaml otel-jaeger-helper:/config
}

remove-config() {
    docker volume rm otel-jaeger-config
}

# Call the requested function and pass the arguments as-is
"$@"