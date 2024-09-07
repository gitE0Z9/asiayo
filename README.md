# asiayo 匯率作業

## 啟動

1. 編譯golang應用

    ```sh
    go build .
    ```

2. 賦予執行權限後，啟動伺服器 

    ```sh
    chmod +x main
    ./main
    ```

3. 查看 Swagger

    ```sh
    curl http://localhost:8080/docs/index.html
    ```

## 路由

> `/api/v1/exchange-rate`

### 參數:

- `source`: 來源匯率，支持 `USD`, `JPY`, `TWD`

- `target`: 目標匯率，支持 `USD`, `JPY`, `TWD`

- `amount`: 金額，e.g. `$100`


## 測試

指令: `go test ./tests`

或者使用 `gotestfmt`: `go test -json -v ./tests | gotestfmt`
