.PHONY: build-docker

build-docker:
	docker build -t igor .
	docker tag igor:latest host.docker.internal:5000/igor:0.0.1
	docker push host.docker.internal:5000/igor:0.0.1

build:
	go build

clean:
	rm ./igor
