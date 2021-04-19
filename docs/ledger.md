# ledger Rippled REST API

> Retrieve information about the public [ledger](https://xrpl.org/ledgers.html).

## Request

### Parameters

| Parameter | Type | Required | Description |
|-----------|------|--------- | ------------|
| `ledger_hash` | `string` | no | A 20-byte hex string for the ledger version to use. (See Specifying Ledgers) |
| `ledger_index` | `string` or `uint` | no | The ledger index of the ledger to use, or a shortcut string to choose a ledger automatically. In addition to specifying a specific numerical `ledger_index` version, three shortcuts can be used: `validated`, `closed`, and `current`. (See [Specifying Ledgers](https://xrpl.org/basic-data-types.html#specifying-ledgers)) |
| `full` | `boolean` | no | Admin required If true, return full information on the entire ledger. Ignored if you did not specify a ledger version. Defaults to false. (Equivalent to enabling transactions, accounts, and expand.) Caution: This is a very large amount of data -- on the order of several hundred megabytes! |
| `accounts` | `boolean` | no | Admin required. If true, return information on accounts in the ledger. Ignored if you did not specify a ledger version. Defaults to false. Caution: This returns a very large amount of data! |
| `transactions` | `boolean` | no | If true, return information on transactions in the specified ledger version. Defaults to false. Ignored if you did not specify a ledger version. |
| `expand` | `boolean` | no | Provide full JSON-formatted information for transaction/account information instead of only hashes. Defaults to false. Ignored unless you request transactions, accounts, or both. |
| `owner_funds` | `boolean` | no | If true, include owner_funds field in the metadata of OfferCreate transactions in the response. Defaults to false. Ignored unless transactions are included and expand is true. |

### GET Example

```
GET /api/ledger/v1.0/ledger?ledger_index=62977748
```

### POST Example

```
POST /api/ledger/v1.0/ledger

{
  "ledger_index: "62977748"
}
```

### Rippled Example

```
$ curl -H 'Content-Type: application/json' -d '{"method":"ledger","params":[{"ledger_index":"validated","full":false,"accounts":false,"transactions":false,"expand":false,"owner_funds":false}]}' https://s1.ripple.com:51234/
```



## Response


```json
{
  "ledger":{
    "accepted":true,
    "account_hash":"B258A8BB4743FB74CBBD6E9F67E4A56C4432EA09E5805E4CC2DA26F2DBE8F3D1",
    "close_flags":0,
    "close_time":638329271,
    "close_time_human":"2020-Mar-24 01:41:11.000000000 UTC",
    "close_time_resolution":10,
    "closed":true,
    "hash":"3652D7FD0576BC452C0D2E9B747BDD733075971D1A9A1D98125055DEF428721A",
    "ledger_hash":"3652D7FD0576BC452C0D2E9B747BDD733075971D1A9A1D98125055DEF428721A",
    "ledger_index":"54300940",
    "parent_close_time":638329270,
    "parent_hash":"AE996778246BC81F85D5AF051241DAA577C23BCA04C034A7074F93700194520D",
    "seqNum":"54300940",
    "totalCoins":"99991024049618156",
    "total_coins":"99991024049618156",
    "transaction_hash":"FC6FFCB71B2527DDD630EE5409D38913B4D4C026AA6C3B14A3E9D4ED45CFE30D"
  },
  "ledger_hash":"3652D7FD0576BC452C0D2E9B747BDD733075971D1A9A1D98125055DEF428721A",
  "ledger_index":54300940,
  "status":"success",
  "validated":true
}
```