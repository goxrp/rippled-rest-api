# Rippled REST API

[![Build Status][build-status-svg]][build-status-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![License][license-svg]][license-url]
[![Swagger UI][swaggerui-svg]][swaggerui-url]
[![Heroku][heroku-svg]][heroku-url]

Rippled REST API provides a REST-like API proxy for the [`rippled` server](https://github.com/ripple/rippled). It is more REST-like than the existing interfaces and should be easier to use for those with REST API experience, while remaining close enough to the [`rippled` server's API](https://xrpl.org/rippled-api.html) to ensure maintenance, support and documenation should be eaiser than if a larger change was introduced.

A primary design goal of Rippled REST API is to support the [OpenAPI Specification](spec_rippled-rest_openapi3.yaml) and tools such as [Swagger UI](https://goxrp.github.io/rippled-rest-api/api-explorer/).

> Note, if you just want to use Postman, see the [`rippled-postman`](https://github.com/goxrp/rippled-postman) Postman Collection which can connect to rippled servers directly.

Providing a a REST-like interface to Rippled requires thoughtful API design. Increasing REST-likeness will be introduced in several phases with  different versions of the Rippled REST API will have design guidelines for that version.

It can run as any of the following:

1. Server using `net/http` or `fasthttp`
2. AWS Lambda function behind AWS API Gateway
3. Heroku server

A test server is available at:

https://rippled-rest-api.herokuapp.com 

## Installation

### Environment Variables

1. `HTTP_ENGINE`: set to `nethttp`, `fasthttp` or `awslambda`.
2. `PORT`: TCP port to listen on

### Server

```
$ go get github.com/goxrp/rippled-rest-api
$ rippled-rest-api
```

### AWS Lambda

1. Use the `aws-package.sh` script to create a zip file to upload to AWS Lambda.
1. Change "Runtime settings"  > "Handler" to `main` from `hello`

### Heroku

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

or

```
$ heroku create
$ git push heroku master
$ heroku open
```

Heroku will automatically set the `PORT` environment variable.

### Testing

Test with:

* `curl -XGET 'https://{gatewayId}.execute-api.{awsZone}.amazonaws.com/{stage}/api'`
* `curl -XPOST 'https://{gatewayId}.execute-api.{awsZone}.amazonaws.com/{stage}/api/v1/account_info' -H 'Content-Type: application/json' -d @docs/examples/endpoint_account_info_request.json`

## API Spec

The API Spec is a work in progress and available at:

* [`spec_rippled-rest_openapi3.yaml`](spec_rippled-rest_openapi3.yaml)
* [`spec_rippled-rest_postman2.json`](spec_rippled-rest_postman2.json)

## Example Request

### API Request: HTTP API

```
POST /api/account_info

{
  "account": "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn",
  "strict": true,
  "ledger_index": "current",
  "queue": true
}
```

```
curl -XPOST 'http://localhost:8080/api/account_info' -d '{"account": "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn", "strict": true, "ledger_index": "current", "queue": true}'
```

### API Request: Websockets

```json
{
  "id": 2,
  "command": "account_info",
  "account": "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
  "strict": true,
  "ledger_index": "current",
  "queue": true
}
```

### API Request: JSON-RPC

```
POST /api/account_info

{
  "method": "account_info",
  "params": [
    {
      "account": "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn",
      "strict": true,
      "ledger_index": "current",
      "queue": true
    }
  ]
}
```

## Credits

1. Heroku support provided via [`github.com/grokify/goheroku`](https://github.com/grokify/goheroku).

 [build-status-svg]: https://github.com/goxrp/rippled-rest-api/workflows/go%20build/badge.svg?branch=master
 [build-status-url]: https://github.com/goxrp/rippled-rest-api/actions
 [goreport-svg]: https://goreportcard.com/badge/github.com/goxrp/rippled-rest-api
 [goreport-url]: https://goreportcard.com/report/github.com/goxrp/rippled-rest-api
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/goxrp/rippled-rest-api
 [docs-godoc-url]: https://pkg.go.dev/github.com/goxrp/rippled-rest-api
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/goxrp/rippled-rest-api/blob/master/LICENSE
 [heroku-svg]: https://img.shields.io/badge/%E2%86%91_deploy-Heroku-7056bf.svg?style=flat
 [heroku-url]: https://heroku.com/deploy
 [swaggerui-svg]: https://shields.io/badge/Swagger%20UI-OpenAPI%203-blue?logo=swagger
 [swaggerui-url]: https://goxrp.github.io/rippled-rest-api/api-explorer/