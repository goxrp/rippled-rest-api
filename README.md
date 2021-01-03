# Ripple API

Ripple API provides

https://xrpl.org/rippled-api.html

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