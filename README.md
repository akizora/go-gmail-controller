# go-gmail-controller

## docker

``:sh
docker-compose exec goapp go run hello/main.go
docker-compose exec goapp go mod download google.golang.org/api
docker-compose exec goapp go mod download github.com/googleapis/enterprise-certificate-proxy
docker-compose exec goapp go mod download github.com/googleapis/gax-go/v2
docker-compose exec goapp go mod download go.opencensus.io
docker-compose exec goapp go mod download github.com/golang/groupcache
docker-compose exec goapp go mod download golang.org/x/text
docker-compose exec goapp go mod download google.golang.org/genproto
docker-compose exec goapp go mod download google.golang.org/grpc
docker-compose exec goapp go mod download golang.org/x/sys
docker-compose exec goapp go get golang.org/x/oauth2
docker-compose exec goapp go mod download cloud.google.com/go cloud.google.com/go/compute
```
