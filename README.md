# Go Mock

An example go library that uses [`go.uber.org/mock`](https://pkg.go.dev/go.uber.org/mock) to easily
mock an external service for unit testing.

An important `Application` method `CoolAlgorithm` needs to be unit tested.

`CoolAlgorithm` relies on an external `calculator` service and for whatever reason, that external
service cannot be used during unit testing.

To remedy this, we first define the `Calculatorer` interface:

In order to mock this interface, we first have to install `mockgen`:

```sh
go install go.uber.org/mock/mockgen@latest
```

We then use `mockgen` to generate a mock client to use in our unit test:

```sh
mockgen -source=./calculator/Calculatorer.go -destination=./mock_calculator/Client.go
```

Finally, we construct an `Application` instance that uses our mocked client. We set expectations and
return values for the methods that will be called during the `CoolAlgorithm` unit test:

_TODO: add code examples_
