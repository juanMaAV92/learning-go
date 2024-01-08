

## Run DB

```bash
    cd ~/database
    docker build . -t api-rest-ws-db 

    docker run -p 54321:5432 api-rest-ws-db
```