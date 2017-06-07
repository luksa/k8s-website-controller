build:
	CGO_ENABLED=0 GOOS=linux go build -o website-controller -a pkg/website-controller.go

image: build
	docker build -t luksa/website-controller .

push: image
	docker push luksa/website-controller:latest
