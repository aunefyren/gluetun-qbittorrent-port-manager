#!/bin/sh

# initialize the command with the binary
CMD="/app/gluetun-qbittorrent-port-manager"

# Add the flags
if [ -n "$TZ" ]; then
  CMD="$CMD --tz $TZ"
fi
if [ -n "$ENVIRONMENT" ]; then
  CMD="$CMD --environment $ENVIRONMENT"
fi
if [ -n "$INTERVAL" ]; then
  CMD="$CMD --interval $INTERVAL"
fi
if [ -n "$PORTFILE" ]; then
  CMD="$CMD --portfile $PORTFILE"
fi
if [ -n "$LOGLEVEL" ]; then
  CMD="$CMD --loglevel $LOGLEVEL"
fi

# add qBit flags
if [ -n "$PORT" ]; then
  CMD="$CMD --port $PORTt"
fi
if [ -n "$IP" ]; then
  CMD="$CMD --ip $IP"
fi
if [ -n "$HTTPS" ]; then
  CMD="$CMD --https $HTTPS"
fi
if [ -n "$USERNAME" ]; then
  CMD="$CMD --username $USERNAME"
fi
if [ -n "$PASSWORD" ]; then
  CMD="$CMD --password $PASSWORD"
fi

# execute the final command
exec $CMD