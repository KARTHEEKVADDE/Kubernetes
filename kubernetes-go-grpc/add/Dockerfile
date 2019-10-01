FROM golang

COPY . /gitrepo/kubernetes-go-grpc/add
WORKDIR /gitrepo/kubernetes-go-grpc/add
RUN go get .
ENTRYPOINT go run main.go
EXPOSE 3000
