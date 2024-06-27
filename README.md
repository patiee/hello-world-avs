# go-operator

Go service for registering as delegation operator, reading `hello world` smart contract events and sending response for every task created.

1. Prepare .env

    ```sh
    cp .env.example .env
    ```

2. Run spam task

    ```sh
    go run cmd/spam/spamTask.go
    ```

3. Run operator

    ```sh
    go run cmd/spam/operator.go
    ```
