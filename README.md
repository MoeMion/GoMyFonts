# GoMyFont
> A Google Fonts proxier to set up your own Google Fonts Mirror.



![GoMyFont](screenshot.png)

A simple go program for setting up a Google Fonts Mirror. [简体中文](https://www.qwq.cc/)

## Features

- Easy to deploy.
- Support for Cache.
- A practice for my go learning.

## Build & Run

```shell
git clone https://github.com/MoeMion/GoMyFonts.git
cd ./GoMyFonts/
go build
#or
go run
```

## Usage example

You can build for go file or download [Releases](https://github.com/MoeMion/GoMyFonts/releases).

```shell
$ ./gomyfonts -help
Usage: gomyfonts [-h] [-p :port] [-l link] [-t title] [-c timeout of cache]
Github: https://github.com/MoeMion/GoMyFonts
Author: Mion
Options:
  -c int
        Expiration of cache,Unit:Minute. (default 10)
  -l string
        The url of your mirror site. (default "http://127.0.0.1/")
  -p string
        Bind TCP Port. (default ":2333")
  -t string
        The title of your mirror site. (default "GoMyFonts")
```

## Run as a daemon

You can run MyGoFonts as a daemon. Here are some possible methods.

### Systemd

Copy binary file to `/usr/local/bin`,  run command as follow:

```sh
cp GoMyFonts /usr/local/bin
```

Create the systemd configuration file at `/etc/systemd/system/gomyfonts.service`:

```
[Unit]
Description=System service for GoMyFonts.
After=network.target

[Service]
Type=simple
Restart=always
ExecStart=/usr/local/bin/GoMyFonts #you can specify parameters here.

[Install]
WantedBy=multi-user.target
```

Launch GoMyFonts when system startup with:

```shell
systemctl enable gomyfonts
```

Launch clashd immediately with:

```
systemctl start gomyfonts
```

Check the health and logs of Clash with:

```
systemctl status gomyfonts
journalctl -xe
```

### Supervisor

Example of Supervisor configuration:

```
[program:gomyfonts]
command=/path/to/gomyfonts #you can specify parameters here.
directory=/path/to/
autorestart=true
startsecs=3
startretries=3
user=root
priority=999
numprocs=1
```



## Use TLS and HTTPS

GoMyFonts currently does not support this feature, but you can use Nginx or Caddy to enable TLS/HTTPS connection.

## Author

Mion

Visit my blog : https://www.qwq.cc/

GoMyFonts is a practice for my go learning, if you find some bugs, please submit issue XD!
