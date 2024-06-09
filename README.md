# pushgateway


new promethues pushgateway


replace [prom/pushgateway](https://github.com/prometheus/pushgateway)


used low memerey, don't OOM

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

