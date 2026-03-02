# gluetun-qbittorrent-port-manager

> [!NOTE]  
> This is a spiritual successor to [SnoringDragon/gluetun-qbittorrent-port-manager](https://github.com/SnoringDragon/gluetun-qbittorrent-port-manager).

## What?
A simple Go application that syncs the Gluetun port-forwarded port to qBittorrent.

## Why?
Gluetun can get a port which is port-forwarded, and has the ability to write this port to a file. We need qBitTorrent to use this port as a `listen port` when doing peer-to-peer connection. So this simple program synchronizes those values.

## How?
The file is read by the program, the qBitTorrent port is checked through their API. If they are not in sync, qBitTorrent is updated through their API.

## When
The file is always monitored and the port is updated when the file changes. There is also a configurable interval which the program will query qBitTorrent and verify the port is correct. There is also a verification, so the port is only changed if it is incorrect.

## Configuration?
| Config file entry | Startup flag | Environment variable | Type | Description |
|-----|-----|-----|-------|--------------|
| port | port | PORT | int | Port qBit listens on (default: `8080`) |
| ip | ip | IP | string | IP qBit listens on (default: `localhost`) |
| https | HTTPS | https | bool | qBit protocol (`true` or `false`) |
| username | username | USERNAME | string | Username for qBit login (default: `admin`) |
| password | password | PASSWORD | string | Password for qBit login |
| timezone | tz | TZ | string | Timezone of the app (default: `Europe/Paris`) |
| environment | environment | ENVIRONMENT | string | Defines program behavior (default: `production`) |
| log_level | loglevel | LOGLEVEL | string | Amount of logs (default: `info`) |
| interval | interval | INTERVAL | int | Minutes between qBit port check (default: `5`) |
| port_file | portfile | PORTFILE | string | File where Gluetun writes port (default: `/tmp/gluetun/forwarded_port`) |

## Docker example?
```yaml
services:
  gluetun-qbittorrent-port-manager:
    container_name: gluetun-qbittorrent-port-manager
    image: ghcr.io/aunefyren/gluetun-qbittorrent-port-manager:latest
    restart: unless-stopped
    environment:
      PUID: 1000 # this ensures the container is not running as root
      PGID: 1000 # this ensures the container is not running as root
      TZ: Europe/Paris
      PORTFILE: /tmp/gluetun/forwarded_port # should be similar to mounted volume
      IP: localhost
      PORT: 8080
      USERNAME: admin
      PASSWORD: secretpassword
    volumes:
      - /tmp/gluetun/forwarded_port:/tmp/gluetun/forwarded_port:rw # make sure this leads to your port file
```