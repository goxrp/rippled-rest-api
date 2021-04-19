# ledger_closed Closed Ledger API

> The `ledger_closed` method returns the unique identifiers of the most recently closed ledger. (This ledger is not necessarily validated and immutable yet.)

**Note:** For the Closed Ledger information, the [`ledger_data`](ledger_data/) can also be used by setting `ledger_index` to `closed`.

## Request

The API can be accessed via the following:

### GET Request

```
$ curl 'http://{RIPPLED_REST_URL}/api/ledger/v1.0/ledger_closed'
```

### POST Request

```
$ curl 'http://{RIPPLED_REST_URL}/api/ledger/v1.0/ledger_closed' -d '{}'
```

### Rippled Request

The `rippled` request is provided here for reference and comparison.

```
$ curl -H 'Content-Type: application/json' -d '{"method":"ledger_closed","params":[{}]}' https://s1.ripple.com:51234/
```

## Response

The API returns the current Ledger Version as `ledger_current_index`. This can be used with the `ledger_data` API to get infomration on the Ledger Version.

### REST API Response

```json
{
  "ledger_hash": "7D3A3D86BA8613CBF47184A32A4FAC49CC9748DAA690A2E94CE7EF417CAD2C4B",
  "ledger_index": 62990218,
  "status": "success"
}
```

### Rippled Response

The `rippled` response is provided here for reference and comparison.

```json
{
  "result":{
    "ledger_hash": "7D3A3D86BA8613CBF47184A32A4FAC49CC9748DAA690A2E94CE7EF417CAD2C4B",
    "ledger_index": 62990218,
    "status": "success"
  }
}
```