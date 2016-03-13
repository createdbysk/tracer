# tracer
Use the tracer API to trace an entity for the lifetime of its existence. Just like a radioactive tracer that allows scientists to trace the mechanism of chemical reactions or the path of a fluid, use the tracer API to generate identifiers to track software entities through their lifetime.

# Steps
## Install dependencies
    go get


## Install and use mockgen for tests
* Follow instructions to install [mock](https://github.com/golang/mock) for GO.
* Generate the mock route
      mkdir -p mock/mux
      mockgen -source=mux/router.go > mock/mux/mock_router.go
      mkdir -p mock/tracers/metadata
      mockgen -source=tracers/metadata/storage.go > mock/tracers/metadata/mock_storage.go
      go build
      go install

# References
* [mock](https://github.com/golang/mock)
* [mux](http://www.gorillatoolkit.org/pkg/mux) from the Gorilla Web Toolkit.

## REST
* http://thenewstack.io/make-a-restful-json-api-go/
* http://grokbase.com/t/gg/golang-nuts/144j758twr/go-nuts-is-the-full-absolute-url-of-an-http-request-directly-available
* Set Headers including Access-Control-Allow-Origin - http://stackoverflow.com/questions/12830095/setting-http-headers-in-golang

## Testing
* http://nathanleclaire.com/blog/2015/10/10/interfaces-and-composition-for-effective-unit-testing-in-golang/

# TODO
1. Create a makefile or equivalent.
    * Determine how to run mockgen automatically.
