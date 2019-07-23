## Go-Kit Microservice

This is a Go-Kit service created for the sole purpose of learning go-kit

#### Services

1. Sum two numbers together.<br>
   POST - `curl -XPOST -d'{"A":13,"B": 12}' localhost:8081/sum`<br>
   Response - `{"v":25}`
   <br>
2. Concat to values together.<br>
   POST - `curl -XPOST -d'{"A":"hello","B": "world"}' localhost:8081/concat`<br>
   Response - `{"v":"helloworld"}`

#### To run

`go run main.go`

#### To test

`go test`

#### To add a new service

The main issue with Go-Kit is the repeating of boilerplate code. You'll see this a lot when you want to add a new service. However, this also means that you can use the examples in each file to recreate fairly easily. You'll need to do the following:

1. Endpoint package
   - Add ExampleRequest and ExampleResponse to datastruct.go
   - Add entries to create endpoints in set.go
   - Add entries to any current middleware in middleware.go
2. Httptransport package
   - Add new decoders for both ExampleRequest and ExampleResponse struct you've created into decoders.go
   - (optional) if response is different from the generic response, then add a new encoder to create the response, else use the generic response encoder
   - Link the new methods together by creating a new handler in http.go
3. Service package
   - Add entry to Service interface in servicve.go
   - Add entry to New method in service.go
   - Create bussiness logic in service.go
   - Add entries to any current middleware in middleware.go

#### To add middleware

The great thing about Go-Kit is that middleware is easily added, updated, or removed. Go-Kit middleware are simply just decorators that take an endpoint and returns an endpoint so that it adds behaviour without affecting its function.

_This following implementation is only consistent with my current knowledege of the different types of microservice middleware. I will not be suprised to find that there is other middleware which will need a completely different implementation._

Currently we have two types of middleware

1. Endpoint package middleware.go<br>
   Middleware defined here can act on when an endpoint action has started and finished (intercepting?)
2. Service package middleware.go<br>
   Middleware defined here can act on the results of business logic actions.

### Jenkins setup

To use in a Jenkins environment you must have the following plugins installed:

- [Go](https://wiki.jenkins.io/display/JENKINS/Go+Plugin) plugin
  installed and the global tool Go configured as:

  - Name: `Go 1.12.7`

  - Install automatically checked

  - Install from golang.org

  - Version: `Go 1.12.7`

- [OpenShift Client plugin](https://plugins.jenkins.io/openshift-client)
  installed and have global tool OpenShift Client Tools configured as:

  - Name: `oc 3.9.0`

  - Install automatically checked

  - Add installer 'Extract *.zip/*.tar.gz'

  - Download URL for binary archive: `https://github.com/openshift/origin/releases/download/v3.9.0/openshift-origin-client-tools-v3.9.0-191fece-linux-64bit.tar.gz`

  - Subdirectory of extracted archive: `openshift-origin-client-tools-v3.9.0-191fece-linux-64bit`

- [Docker Pipeline Plugin](https://wiki.jenkins.io/display/JENKINS/Docker+Pipeline+Plugin)
