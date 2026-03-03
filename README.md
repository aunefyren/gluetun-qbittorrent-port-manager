# Gluetun qBitTorrent Port Manager

> [!NOTE]  
> This is a spiritual successor to [SnoringDragon/gluetun-qbittorrent-port-manager](https://github.com/SnoringDragon/gluetun-qbittorrent-port-manager).

## What?
A simple Go application that syncs your [Gluetun port](https://github.com/qdm12/gluetun-wiki/blob/main/setup/options/port-forwarding.md) to the [qBittorrent](https://github.com/qbittorrent/qBittorrent) listening port. It can run as a standalone executable, or in a Docker container.

## Why?
Gluetun can receive a port from a VPN provider which is port-forwarded. Many want to utilize this port elsewhere, and therefore Gluetun has the ability to write this port to a file. It is often desired for qBitTorrent to use this port as a `listen port` when doing peer-to-peer connections. This simple program synchronizes those values, ensuring qBitTorrent stays in sync with Gluetun.

## How?
The Gluetun file is read by the program and the qBitTorrent port is checked through their API. If they are not in sync, qBitTorrent is updated through their API. If Gluetun does not have an open port, nothing is done to qBitTorrent.

## When?
The file is always monitored and the port is updated when the file changes. There is also a configurable interval after which the program will query qBitTorrent and verify the port is correct. There is also a verification before updating, ensuring the port is only changed if it is incorrect.

## Configuration?
The `config.json` with the main configuration is automatically created and placed within the `config` directory upon startup. It can be manually edited in a text editor. You can also use start-up flags when launching the executable, which will be written to the file. Like this: 

```powershell
.\gluetun-qbittorrent-port-manager.exe -loglevel=debug
```

Or you can use environment variables on the system, which is typically done when running a Docker container. See the table below for all configuration options available:

| Config file entry | Startup flag | Environment variable | Type | Description |
|-----|-----|-----|-------|--------------|
| port | port | PORT | int | Port qBit listens on (default: `8080`) |
| ip | ip | IP | string | IP qBit listens on (default: `localhost`) |
| https | https | HTTPS | bool | qBit protocol (`true` or `false`) |
| username | username | USERNAME | string | Username for qBit login (default: `admin`) |
| password | password | PASSWORD | string | Password for qBit login |
| timezone | tz | TZ | string | Timezone of the app (default: `Europe/Paris`) |
| environment | environment | ENVIRONMENT | string | Defines program behavior (default: `production`) |
| log_level | loglevel | LOGLEVEL | string | Amount of logs (default: `info`) |
| interval | interval | INTERVAL | int | Minutes between qBit port check (default: `5`) |
| port_file | portfile | PORTFILE | string | File where Gluetun writes port (default: `/tmp/gluetun/forwarded_port`) |

## Docker Compose example?
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
      IP: localhost # qBit connection details
      HTTPS: False
      PORT: 8080
      USERNAME: admin
      PASSWORD: secretpassword
    volumes:
      - ./config/:/app/config/:rw # only needed if you want a permanent config.json file
      - /tmp/gluetun/forwarded_port:/tmp/gluetun/forwarded_port:rw # make sure this leads to your port file from Gluetun
```