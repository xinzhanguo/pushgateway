# pushgateway


new promethues pushgateway


replace [prom/pushgateway](https://github.com/prometheus/pushgateway)


low Memory, don't OOM

## config of promethues
```
- job_name: pushgateway
  scrape_interval: 1m
  scrape_timeout: 10s
  metrics_path: /metrics
  scheme: http
  static_configs:
  - targets:
    - localhost:9091
```

## Instructions
```
# set jobname and nodeip
jobname=test
nodeip=1.1.1.1
# push metrics
curl localhost/metrics | curl --data-binary @- http://localhost:9091/metrics/job/$jobname/instance/$nodeip

# or you can put metrics by file
cat > metrics.txt <<EOF
test_1_total{"action":"test"} 20
test_2_total{"action":"test"} 21
EOF
cat metrics.txt | curl --data-binary @- http://localhost:9091/metrics/job/$jobname/instance/$nodeip

# get metrics
http://localhost:9091/metrics

```
