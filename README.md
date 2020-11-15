# Instagram Stories Generator

## Quick Start Guide

### Compile

```bash
make
```

### Run Server

By default, server will run in `0.0.0.0:8080`, this can be customized by providing the
`--host` and `--port` parameters.

```bash
./instagram-storires server
```

### Create a Story

Send an `HTTP POST` request to the `http://localhost:8080/api/v1/stories/` endpoint, you can do
this manually using a tool like `curl` or via a `webhook` automation from `IFTTT` or some other
automation tool, ie:

```bash
title="Come with me and you'll be, In a world of pure imagination, Take a look and you'll see, Into your imagination" \
curl -X POST http://localhost:8080/api/v1/stories/ -d \
"{\"title\": \"$title\", \"source\": \"Pure Imagination\", \"url\": \"https://www.youtube.com/watch?v=nKhGnRIlaro\"}" \
-H "Content-Type: application/json"
```

Images will be generated under `$project/assets/stories/`, ie:
![image info](https://github.com/Alevsk/Instagram-stories/blob/master/assets/stories/a2fbdf34a845457fa8dbb6e4a0e3882d.png?raw=true)



### Help

```bash
./instagram-stories --help
NAME:
  instagram-stories - instagram-stories

DESCRIPTION:


USAGE:
  instagram-stories [FLAGS] COMMAND [ARGS...]

COMMANDS:
  server, srv  starts instagram-stories server

FLAGS:
  --help, -h     show help
  --version, -v  print the version

VERSION:
  (dev)
```