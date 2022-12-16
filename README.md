# icndev-server

## Run

```$xslt
make all
./icndev-server
```

## Run with docker

You can run application with docker.

```$xslt
docker build -t icndev-server .
```

then

```$xslt
docker run --name icndev-server -p 9060:9060 icndev-server \&
-v /mnt/data/icndev-server/logs:/root/go/src/github.com/bianjieai/icndev-server/logs \&
icndev-server
```