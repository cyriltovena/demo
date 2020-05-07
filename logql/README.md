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


```logql
{compose_service="nginx"}
| regexp
"\"(?P<method>\\w+ )(?P<path>.*) HTTP\\/\\d\\.\\d\" (?P<status_code>\\d{3}) (?P<value>\\d{1,})"
```

```logql
max by (path)
  (
    avg_over_time(
    {compose_service="nginx"}
    | regexp
    "\"(?P<method>\\w+ )(?P<path>.*) HTTP\\/\\d\\.\\d\" (?P<status_code>\\d{3}) (?P<value>\\d{1,})"
    [1m]
  )
) / 1000
```
