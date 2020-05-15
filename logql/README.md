# LogQL v2 demo

## Install the Loki docker-driver

This demo uses the Loki driver plugin to ship logs you need to install it first using :

```bash
docker plugin install grafana/loki-docker-driver:latest --alias loki --grant-all-permissions
```

## Run

```bash
docker-compose up
open http://localhost:3000
```

## Try

Queries to try

```logql
sum by (method,path)
(
  rate({compose_service="nginx"}
    | regexp "\"(?P<method>\\w+ )(?P<path>.*) HTTP"
  [1m])
 )
```

```logql
{compose_service="nginx"}
    | regexp
    "\"(?P<method>\\w+ )(?P<path>.*) HTTP\\/\\d\\.\\d\" (?P<status_code>\\d{3}) \\d{1,} (?P<value>\\d{1,}.\\d{1,})"
```

```logql
max by (path)
  (
    avg_over_time(
    {compose_service="nginx"}
    | regexp
   "\"(?P<method>\\w+ )(?P<path>.*) HTTP\\/\\d\\.\\d\" (?P<status_code>\\d{3}) \\d{1,} (?P<value>\\d{1,}.\\d{1,})"    [1m]
  )
)
```


```logql
quantile_over_time(0.99,
  {compose_service="nginx"}| regexp
    "\"\\w+ .* HTTP\\/\\d\\.\\d\" \\d{3} \\d{1,} (?P<value>\\d{1,}.\\d{1,})"
    [1m]
  )
```
