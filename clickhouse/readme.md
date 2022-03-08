
# clickhouse 

## docker-compose install

```
version: "3.7"
services:
  clickhouse:
    container_name: clickhouse
    image: yandex/clickhouse-server
    volumes:
      - ./data/config:/var/lib/clickhouse
    ports:
      - "8123:8123"
      - "19000:9000"
      - "19009:9009"
      - "19004:9004"
    ulimits:
      nproc: 65535
      nofile:
        soft: 262144
        hard: 262144
    healthcheck:
      test: ["CMD", "wget", "--spider", "-q", "localhost:8123/ping"]
      interval: 30s
      timeout: 5s
      retries: 3
    deploy:
      resources:
        limits:
          cpus: '4'
          memory: 4096M
        reservations:
          memory: 4096M
  ch_client:
    image: yandex/clickhouse-client
    entrypoint:
      - /bin/sleep
    command:
      - infinity
    networks:
        - ch_ntw
  tabixui:
    container_name: tabixui
    image: spoonest/clickhouse-tabix-web-client
    ports:
      - "18080:80"
    depends_on:
      - clickhouse
    deploy:
      resources:
        limits:
          cpus: '0.1'
          memory: 128M
        reservations:
          memory: 128M
networks:
  ch_ntw:
    driver: bridge
    ipam:
      config:
        - subnet: 10.222.1.0/24
```
