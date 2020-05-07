# LogQL v2 demo

## Install the loki docker-driver

docker plugin install grafana/loki-docker-driver:latest --alias loki --grant-all-permissions

## Run

```bash
docker-compose up
open http://localhost:3000
```

## Try

Queries to try

```logql
sum by (path,method)
(
  rate({compose_service="nginx"}
    | regexp "\"(?P<method>\\w+ )(?P<path>.*) HTTP"
  [1m])
 )
```
