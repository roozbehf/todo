# ToDo List Example
    
A very simple ToDo List example using Swagger and Go.

## Server

### REST API
The API of the server is defined by the [`api/server.yml`](api/server.yml) Swagger specification. 

### Build 

Make sure you that
* you have `dep` installed. Visit https://github.com/golang/dep 
* your `GOPATH` and `GOROOT` environments are set properly.

#### Makefile
There is a [`Makefile`](Makefile) provided that offers a number of targets for preparing, building and running the service. To build the service, simply run:
```
make clean go-dep go-build
```

## License
The code is published under an [MIT license](LICENSE.md). 

