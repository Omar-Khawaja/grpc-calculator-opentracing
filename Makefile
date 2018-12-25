install:
	dep version || go get -u github.com/golang/dep/cmd/dep
	dep ensure
	bash jaeger.sh
