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
  rate({compose_project="logql"}
  | pattern `<ip> - - [<_>] "<method> <path> HTTP<_>" <status> <_> <latency> `
  [1m])
 )
```

```logql
{compose_project="logql"}
  | pattern `<ip> - - [<_>] "<method> <path> HTTP<_>" <status> <_> <latency> `
```

```logql
quantile_over_time(0.99,{compose_project="logql"}
  | pattern `<ip> - - [<_>] "<method> <path> HTTP<_>" <status> <_> <latency> ` | latency != "" | unwrap latency [$__interval]) by (compose_project)
```
