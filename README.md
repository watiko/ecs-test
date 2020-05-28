## on macOS

```
$ brew cask install xquartz
$ sudo reboot
$ echo $DISPLAY
$ socat TCP-LISTEN:6000,reuseaddr,fork UNIX-CLIENT:\"$DISPLAY\"
```

```
$ docker-compose run --rm -e DISPLAY=host.docker.internal:0 debug bash
docker> wireshark
```

