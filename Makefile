.PHONY: build-docker

build-docker:
	./scripts/build-docker.sh

build:
	go build

clean:
	rm ./igor
