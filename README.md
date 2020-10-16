# JWT Utils

JWT cli utils.

## Install

```
$ go get ./...
```

## Usage

### Decode

_*Note:*_ This does not verify the token. It is only useful for reading a JWT.

```
$ jwtdecode yourjwthere

{"email":"arthur@example.com","exp":1610601876,"iat":1602825876,"sub":"42"}
```

## Copyright

The MIT License (MIT)

Copyright (c) 2020 Scott Barr

See the [LICENSE](LICENSE.md).
