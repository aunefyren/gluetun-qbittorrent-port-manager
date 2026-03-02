#!/bin/sh

# initialize the command with the binary
CMD="/app/gluetun-qbittorrent-port-manager"

# Add the flags
if [ -n "$tz" ]; then
  CMD="$CMD --tz $tz"
fi
if [ -n "$environment" ]; then
  CMD="$CMD --environment $environment"
fi
if [ -n "$interval" ]; then
  CMD="$CMD --interval $interval"
fi
if [ -n "$portfile" ]; then
  CMD="$CMD --portfile $portfile"
fi

# add qBit flags
if [ -n "$port" ]; then
  CMD="$CMD --port $port"
fi
if [ -n "$ip" ]; then
  CMD="$CMD --ip $ip"
fi
if [ -n "$https" ]; then
  CMD="$CMD --https $https"
fi
if [ -n "$username" ]; then
  CMD="$CMD --username $username"
fi
if [ -n "$password" ]; then
  CMD="$CMD --password $password"
fi

# execute the final command
exec $CMD