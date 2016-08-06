[Load Balancer]

[HTTP API & Routing]

[Transport]

[

	[Service] -> [Database]

	[Service] -> [Database]

]


- [x] Static message
- [x] Random from list
- [ ] Generate list from data

## Workflow

`mkdir $GOPATH/build && cd $GOPATH/build`

`touch Dockerfile docker-compose.yaml`

Setup your Dockerfile for the application and docker-compse.yaml

Build application using `go build -i -o $GOPATH/build/bin`

Build image using `docker-compose build`

Run with `docker-compose up -d`
