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
      go build
      go install

# References
* [mock](https://github.com/golang/mock)
* [mux](http://www.gorillatoolkit.org/pkg/mux) from the Gorilla Web Toolkit.

# TODO
1. Create a makefile or equivalent.
    * Determine how to run mockgen automatically.
