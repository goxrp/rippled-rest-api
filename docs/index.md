# Rippled XRP Ledger REST API Developer Guide

## Overview

This is a Developer Guide for the XRP Ledger API as implemented by the Ripple REST API Proxy service. The goal if this service is to provide a REST API representation of the `rippled` API which is available via websocks, JSON-RPC and gRPC, but not REST.

This service comes with the following:

* Stand-alone service that can run locally
* Service that can run as an AWS Lambda function
* OpenAPI 3.0 Spec
* OpenAPI 2.0 Spec
* Postman 2.1 Collection

If you do not wish to host this or run it locally, you can use a service hosted on RapidAPI that uses this service.

## Design

The goal for this API is to provide an API that is REST-like, while accounting for the design of the underlying API.

* All JSON RPC APIs have been mapped to HTTP POST APIs
* Select APIs have been mapped to HTTP GET APIs for convenience

## Related Projects

### Rippled Go Client SDK


## Past Projects

Other REST API efforts have been discontinued. It is also desired that this work be used to create SDKs in mutiple languages. It is designed to be able to support the same goals as the following efforts.

* [https://github.com/ripple/ripple-rest](https://github.com/ripple/ripple-rest)
* [https://www.npmjs.com/package/ripple-rest](https://www.npmjs.com/package/ripple-rest)
* [https://github.com/ripple-unmaintained/ripple-lib-java](https://github.com/ripple-unmaintained/ripple-lib-java)

## Questions

If you have questions for this API, please ask on GitHub issues or Stack Overflow using the `rippled-rest` tag.