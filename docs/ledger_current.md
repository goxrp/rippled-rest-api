# ledger_current Current Ledger API

The `ledger_current` API returns the value of `ledger_current_index` and not take any parameters.

**Note:** For the Current Ledger information, the [`ledger_data`](ledger_data/) can also be used by setting `ledger_index` to `current`.

## Request

The API can be accessed via the following:

### GET Request

```
$ curl 'http://{RIPPLED_REST_URL}/api/ledger/v1.0/ledger_current'
```

### POST Request

```
$ curl 'http://{RIPPLED_REST_URL}/api/ledger/v1.0/ledger_current' -d '{}'
```

### Rippled Request

The `rippled` request is provided here for reference and comparison.

```
$ curl -H 'Content-Type: application/json' -d '{"method":"ledger_current","params":[{}]}' https://s1.ripple.com:51234/
```

## Response

The API returns the current Ledger Version as `ledger_current_index`. This can be used with the `ledger_data` API to get infomration on the Ledger Version.

### REST API Response

```json
{
  "ledger_current_index": 62988820,
  "status": "success"
}
```

### Rippled Response

The API returns the current Ledger Version as `ledger_current_index`. This can be used with the `ledger_data` API to get infomration on the Ledger Version.

```json
{
  "result":{
    "ledger_current_index":62990274,
    "status":"success"
  }
}
```