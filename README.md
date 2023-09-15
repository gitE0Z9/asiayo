# asiayo

## 啟動

1. 編譯golang應用

    ```sh
    cd src
    go build . -o main
    ```

2. 賦予執行權限後，啟動伺服器 

    ```sh
    chmod +x main
    ./main
    ```

3. 查看API

    ```sh
    curl http://localhost:8080
    ```

## 路由

> `/exchange-rate`

### 參數:

- `source`: 來源匯率，支持 `USD`, `JPY`, `TWD`

- `target`: 目標匯率，支持 `USD`, `JPY`, `TWD`

- `amount`: 金額，e.g. `$100`
