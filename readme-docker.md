# Build Docker Image

Instructions to build docker image contained in file `Dockerfile`/

```bash
docker build -t simplebank:latest .
docker images | grep simplebank
docker image prune                      # remove images with <none> tag - dangling images
docker stop simplebank
docker rm simplebank

docker run --name simplebank -p 8080:8080 simplebank:latest
docker run --name simplebank -p 8080:8080 -e GIN_MODE=release simplebank:latest
docker run --name simplebank -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE=postgres://root:password@172.17.0.2:5432/simplebank?sslmode=disable simplebank:latest
docker run --name simplebank --network bank-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE=postgres://root:password@postgres13:5432/simplebank?sslmode=disable simplebank:latest

docker container inspect postgres13             # 172.17.0.2
docker container inspect simplebank             # 172.17.0.3

docker network ls
docker network inspect bridge                   # postgres13
docker network create bank-network
docker network connect bank-network postgres13  # 172.26.0.2
```

### Single Stage
Image size on a single stag build is pretty large, so best to create a multi-stage docker build Dockerfile. \
Here is the before and after image sizes. 

simplebank                               latest            5e66936b0bfc   52 seconds ago   662MB
simplebank                               latest            ea2791b24484   24 seconds ago   41.1MB