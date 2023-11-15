# Go Mock

An example go library that uses [`go.uber.org/mock`](https://pkg.go.dev/go.uber.org/mock) to easily
mock an external service for unit testing.

An important `Application` method `CoolAlgorithm` needs to be unit tested.
https://github.com/peterjamesmatthews/go-mock/blob/94b81916b1b4e541228f7cafdb95af62c86a87d4/application/CoolAlgorithm.go#L4

`CoolAlgorithm` relies on an external `calculator` service and for whatever reason, that external
service cannot be used during unit testing.

To remedy this, we first define the `Calculatorer` interface:
https://github.com/peterjamesmatthews/go-mock/blob/94b81916b1b4e541228f7cafdb95af62c86a87d4/calculator/Calculatorer.go#L4-L13

In order to mock this interface, we first have to install `mockgen`:

```sh
go install go.uber.org/mock/mockgen@latest
```

We then use `mockgen` to generate a mock client to use in our unit test:

```sh
mockgen -source=./calculator/Calculatorer.go -destination=./mock_calculator/Client.go
```

https://github.com/peterjamesmatthews/go-mock/blob/94b81916b1b4e541228f7cafdb95af62c86a87d4/mock_calculator/Client.go#L1-L94

Finally, we construct an `Application` instance that uses our mocked client. We set expectations and
return values for the methods that will be called during the `CoolAlgorithm` unit test:

https://github.com/peterjamesmatthews/go-mock/blob/94b81916b1b4e541228f7cafdb95af62c86a87d4/application/CoolAlgorithm_test.go#L11-L44
