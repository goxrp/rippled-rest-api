# ledger_data

XRPL.org describes this API as:

> The `ledger_data` method retrieves contents of the specified ledger. You can iterate through several calls to retrieve the entire contents of a single ledger version.

## Request

### Parameters

Request parameters can be included in the REST API request as `GET` query string parameters or a `POST` JSON body.

This API can be queried using `ledger_hash` or `ledger_index`.

It can also be used the `binary` flag set to true or false.

| Parameter | Type | Required | Description |
|-----------|------|--------- | ------------|
| `ledger_hash` | `string` | no | A 20-byte hex string for the ledger version to use. (See Specifying Ledgers) |
| `ledger_index` | `string` or `uint` | no | The ledger index of the ledger to use, or a shortcut string to choose a ledger automatically. In addition to specifying a specific numerical `ledger_index` version, three shortcuts can be used: `validated`, `closed`, and `current`. (See [Specifying Ledgers](https://xrpl.org/basic-data-types.html#specifying-ledgers)) |
| `binary` | `boolean` | no | If set to true, return ledger objects as hashed hex strings instead of JSON. |
| `limit` | `uint` | no | Limit the number of ledger objects to retrieve. The server is not required to honor this value. |
| `marker` | `string` | no | Value from a previous paginated response. Resume retrieving data where that response left off. |

### Get Example

```
GET /api/ledger/v1.0/ledger_data?ledger_index=62977748&limit=5
```

### Post Example

```
POST /api/ledger/v1.0/ledger_data
Content-Type: application/json

{
  "binary": false,
  "ledger_index": "62977748", 
  "limit": 5
}
```

### Rippled Example

```
POST https://s1.ripple.com:51234/
Content-Type: application/json

{
  "method":"ledger_data",
  "params":[
    {
      "binary":false,
      "ledger_index":"62977748",
      "limit":5
    }
  ]
}
```

## Response: Non-Binary

The following is an example pretty formatted non-binary response.

```json
{
  "result":{
    "ledger":{
      "accepted":true,
      "account_hash":"3A794AAA44E23C889CB372F080A284B75DDE9C068AB66452BA03AFD86AC57EB7",
      "close_flags":0,
      "close_time":672049480,
      "close_time_human":"2021-Apr-18 08:24:40.000000000 UTC",
      "close_time_resolution":10,
      "closed":true,
      "hash":"C358F71E2905B7C6E883AE91D344D60897596EC165832495255EED5D3672176B",
      "ledger_hash":"C358F71E2905B7C6E883AE91D344D60897596EC165832495255EED5D3672176B",
      "ledger_index":"62977748",
      "parent_close_time":672049471,
      "parent_hash":"0C0CAD5C8A312004CA9342726D6D192E1705682DEE5BCF9662F918E82F9C2832",
      "seqNum":"62977748",
      "totalCoins":"99990532842465484",
      "total_coins":"99990532842465484",
      "transaction_hash":"1A99C90E7D4FCFB27FCE2765C1C52FFA94BB98B789B205F7814A459FCA1AB618"
    },
    "ledger_hash":"C358F71E2905B7C6E883AE91D344D60897596EC165832495255EED5D3672176B",
    "ledger_index":62977748,
    "marker":"0000139EDA03EF58CE7176F1402035B5EB6AEE49724555DDB0EBA01432B009A7",
    "state":[
      {
        "Account":"rMj5DFATVxw91PDy3AM2wu7Uu1kgrhWypE",
        "Balance":"20722989",
        "Flags":0,
        "LedgerEntryType":"AccountRoot",
        "OwnerCount":0,
        "PreviousTxnID":"4C43C78081198AC96F6EC15315FDD16FC12458DC9A3F1B0BD0AA62F8D9F8FB51",
        "PreviousTxnLgrSeq":62869196,
        "Sequence":4,
        "index":"000003E6AFED1AADCC39AAE0727B354C2286F1503274F345FE661748F24366CF"
      },
      {
        "Balance":{
          "currency":"GCB",
          "issuer":"rrrrrrrrrrrrrrrrrrrrBZbvji",
          "value":"0"
        },
        "Flags":2162688,
        "HighLimit":{
          "currency":"GCB",
          "issuer":"rBfVgTnsdh8ckC19RM8aVGNuMZnpwrMP6n",
          "value":"0"
        },
        "HighNode":"283",
        "LedgerEntryType":"RippleState",
        "LowLimit":{
          "currency":"GCB",
          "issuer":"rhRFGCy2RJTA8oxkjjtYTvofPVGqcgvXWj",
          "value":"2000000"
        },
        "LowNode":"0",
        "PreviousTxnID":"C0C37CE200B509E0A529880634F7841A9EF4CB65F03C12E6004CFAD9718D6694",
        "PreviousTxnLgrSeq":24695242,
        "index":"0000041EFD027808D3F78C8352F97E324CB816318E00B977C74ECDDC7CD975B2"
      },
      {
        "Account":"rHeyw38ezc3LSLzYAYwaBci2KssDkYGVr9",
        "Balance":"20000000",
        "Flags":0,
        "LedgerEntryType":"AccountRoot",
        "OwnerCount":0,
        "PreviousTxnID":"7540CE04B966D67DBD39F3AA832274902B79AF4782F5AC9D4DC7CD18B1D9AE0D",
        "PreviousTxnLgrSeq":47846971,
        "Sequence":2,
        "index":"000004D417A9CE049C9A71A62B004659B5F1AAAB1BEA1EFDE4E01EB3497FD999"
      },
      {
        "Account":"rLeNL66BfgeszBsCBKJcLmYTRThiWrNTUL",
        "Balance":"19999988",
        "Flags":0,
        "LedgerEntryType":"AccountRoot",
        "OwnerCount":1,
        "PreviousTxnID":"E4BE6307E377590FF56BBF2F26DCBC4BA9682A4C141269352E4E2D4E53C1116E",
        "PreviousTxnLgrSeq":37851086,
        "Sequence":2,
        "index":"00000FB78838CA2CFA82E7438B4F54794A6783327326D58C46B4EF137C059038"
      },
      {
        "Account":"rUwXrQMa4HHBbfxQT71YJNbQXaxZPR8Uhp",
        "Balance":"20941276",
        "Flags":0,
        "LedgerEntryType":"AccountRoot",
        "OwnerCount":0,
        "PreviousTxnID":"4ECC554CC1DEFFD447EEA6FB068E6BED4179C44BAF10F7F61F163ED25FF7EA85",
        "PreviousTxnLgrSeq":60539737,
        "Sequence":2,
        "index":"000012F60C3F1E226D03F974AE8E77250B2BEA91C38AB4146B6055A048C7D540"
      }
    ],
    "status":"success",
    "validated":true
  }
}
```

## Response: Binary

`/api/ledger/v1.0/ledger_data?ledger_index=62977748`

curl -H 'Content-Type: application/json' -d '{"method":"ledger_data","params":[{"binary":true,"ledger_index":"62977748","limit":5}]}' https://s1.ripple.com:51234/



The following is an example pretty formatted binary response.

```json
{
  "result":{
    "ledger":{
      "closed":true,
      "ledger_data":"03C0F6D401633CDC1EFA28CC0C0CAD5C8A312004CA9342726D6D192E1705682DEE5BCF9662F918E82F9C28321A99C90E7D4FCFB27FCE2765C1C52FFA94BB98B789B205F7814A459FCA1AB6183A794AAA44E23C889CB372F080A284B75DDE9C068AB66452BA03AFD86AC57EB7280EA93F280EA9480A00"
    },
    "ledger_hash":"C358F71E2905B7C6E883AE91D344D60897596EC165832495255EED5D3672176B",
    "ledger_index":62977748,
    "marker":"0000139EDA03EF58CE7176F1402035B5EB6AEE49724555DDB0EBA01432B009A7",
    "state":[
      {
        "data":"110061220000000024000000042503BF4ECC2D00000000554C43C78081198AC96F6EC15315FDD16FC12458DC9A3F1B0BD0AA62F8D9F8FB516240000000013C352D8114E376654FF7B1F656D56462FB43E77E9776EE7396",
        "index":"000003E6AFED1AADCC39AAE0727B354C2286F1503274F345FE661748F24366CF"
      },
      {
        "data":"1100722200210000250178D1CA37000000000000000038000000000000028355C0C37CE200B509E0A529880634F7841A9EF4CB65F03C12E6004CFAD9718D66946280000000000000000000000000000000000000004743420000000000000000000000000000000000000000000000000166D6071AFD498D000000000000000000000000000047434200000000002599D1D255BCA61189CA64C84528F2FCBE4BFC3867800000000000000000000000000000000000000047434200000000006EEBB1D1852CE667876A0B3630861FB6C6AB358E",
        "index":"0000041EFD027808D3F78C8352F97E324CB816318E00B977C74ECDDC7CD975B2"
      },
      {
        "data":"110061220000000024000000022502DA163B2D00000000557540CE04B966D67DBD39F3AA832274902B79AF4782F5AC9D4DC7CD18B1D9AE0D624000000001312D008114B6B047F1FE00A59289D45CDDB0FE81F6BD07A267",
        "index":"000004D417A9CE049C9A71A62B004659B5F1AAAB1BEA1EFDE4E01EB3497FD999"
      },
      {
        "data":"110061220000000024000000022502418FCE2D0000000155E4BE6307E377590FF56BBF2F26DCBC4BA9682A4C141269352E4E2D4E53C1116E624000000001312CF48114D774EA776552E07F863D6BE94ADFD8735A28D82E",
        "index":"00000FB78838CA2CFA82E7438B4F54794A6783327326D58C46B4EF137C059038"
      },
      {
        "data":"1100612200000000240000000225039BC3592D00000000554ECC554CC1DEFFD447EEA6FB068E6BED4179C44BAF10F7F61F163ED25FF7EA856240000000013F89DC81147A762D01DEFA26F7EE16BFAD723468A366E8F4F0",
        "index":"000012F60C3F1E226D03F974AE8E77250B2BEA91C38AB4146B6055A048C7D540"
      }
    ],
    "status":"success",
    "validated":true
  }
}
```