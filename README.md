# 總控後台 central-admin

## 專案起始點

- ./cmd/main.go

## 容器啟動文件

- ./service/internal/sandbox/tools/docker-compose.yaml

## 設定檔位置
- ./conf.d/app.yaml
- ./conf.d/config.yaml

## 連線DEV設定檔

-  本機使用docker建立過本地MYSQL / MONGO，則設定黨可以保持本地配置，
-  沒有配置過，則可以修改為連線DEV配置。

### app.yaml

  ```yaml
  ops_config:
  mysql_service:
    address: "10.200.252.216:3306"
    username: "rd_infra_ops"
    password: "bkinfra_ops@DEV"
    database: "central_admin"
  mongo_service:
    host: "10.200.252.156:27017"
    user: "rd1"
    passwd: "devbackend1"
    db: "central-admin"
  file_server:
    path: "https://squirrel-dev.paradise-soft.com.tw"
    host: "https://squirrel-dev.paradise-soft.com.tw:12345"
    username: "central-admin"
    password: "1qaz!QAZ"
  ```

### config.yaml

  ```yaml
  app_config:
  log_config:
    name: "central-admin"
    env: "local"
    level: "debug"
  gin_config:
    port: ":8800"
    debug_mode: true
  mysql_config:
    log_mode: true
    max_idle: 3
    max_open: 10
    conn_max_life_min: 15
  mongo_config:
    max_idle_time: 5
    max_open_conn: 50
  local_cache_config:
    default_expiration_sec: 60
  ```