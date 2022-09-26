# 商城 mall

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
    address: "localhost:3306"
    username: "root"
    password: "abc123"
    database: "mall"
  redis_service:
    address: 'localhost:6379'
    password: ''
    db: 0

  ```

### config.yaml

  ```yaml
app_config:
  log_config:
    name: "mall"
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
  local_cache_config:
    default_expiration_sec: 60
    session_expiration_sec: 1209600
  ```

## 目錄結構

---
- cmd ---- :啟動點
- conf.d ---- :設定檔案
- doc
    - DDL ---- :資料定義 sql
    - DML ---- :資料操作 sql
- service
    - internal
        - app
            - web ---- :設定路由
        - binder ---- :綁定 dig 提供的 interface
        - config. ---- :設定參數定義
        - constant ---- :常數定義
        - controller ---- :Controller 層
            - handler ---- :協助 controller bind 或處理請求
            - middleware. ---- :給 app 用的 middleware
        - core ---- :定義 usecase 的地方
            - common ---- :共用邏輯 (不可以接 controller
            - usecase ---- :使用場景 (可外接
        - errs ---- :error 定義
        - model ---- :只能在指定層級使用 model 為輸入輸出值
            - bo ---- :商業邏輯 model (core/common, core/usecase 只能在這做傳入，傳出值
            - dto ---- :資料轉換(api, in/out) model (controller/* 只能在這做傳入，傳出值
            - po ---- :資料存取層 model (repository 只能在這做傳入，傳出值
        - repository ---- :資料存取層，碰DB, Redis, .....
        - sandbox
            - tools ---- :docker-compose 啟動點
                - sql ---- :docker 第一次起 container 會執行這裡所有 sql 以初始話
        - thirdparty ---- :第三方提供套件
        - utils ---- :工具包
---