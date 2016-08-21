.PHONY : image shell generate

image : Dockerfile
	@docker build -t dev-go-pbuf:latest .

shell :
	@docker run -itv $(GOPATH)/src:/go/src dev-go-pbuf:latest

generate: *.proto
	@go generate

