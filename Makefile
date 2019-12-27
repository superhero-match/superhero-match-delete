prepare:
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/gin-gonic/gin
	go get -u golang.org/x/sys/unix
	go get -u github.com/jinzhu/configor
	go get -u go.uber.org/zap
	go get -u github.com/segmentio/kafka-go
	go get -u golang.org/x/net/context

run:
	go build -o bin/main cmd/api/main.go
	./bin/main

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o bin/main cmd/api/main.go
	chmod +x bin/main

deps:
	dep ensure -v

dkb:
	docker build -t superhero-match-delete .

dkr:
	docker run --rm -p "4300:4300" superhero-match-delete

launch: dkb dkr

api-log:
	docker logs superhero-match-delete -f

es-log:
	docker logs es -f

rmc:
	docker rm -f $$(docker ps -a -q)

rmi:
	docker rmi -f $$(docker images -a -q)

clear: rmc rmi

api-ssh:
	docker exec -it superhero-match-delete /bin/bash

es-ssh:
	docker exec -it es /bin/bash

PHONY: prepare build dkb dkr launch api-log es-log api-ssh es-ssh rmc rmi clear