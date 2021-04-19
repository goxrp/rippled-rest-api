# Ledger APIs Overview

## The XRP Ledger

The Ledger APIs represent APIs on the XRP Ledger.

XRPL.org describes the ledger as [the following](https://xrpl.org/ledgers.html):

> The XRP Ledger is a shared, global ledger that is open to all. Individual participants can trust the integrity of the ledger without having to trust any single institution to manage it. The `rippled` server software accomplishes this by managing a ledger database that can only be updated according to very specific rules. Each instance of `rippled` keeps a full copy of the ledger, and the peer-to-peer network of rippled servers distributes candidate transactions among themselves. The consensus process determines which transactions get applied to each new version of the ledger. See also: [The Consensus Process](https://xrpl.org/consensus.html).

## Ledger Versions, Ledger Hash, Ledger Index

The XRP Ledger is versioned and a new version is avaialable every few seconds. A particular ledger version is referred to via the `ledger_hash` or `ledger_index`.

[The following describes Ledger Versions](https://xrpl.org/consensus.html):

> The XRP Ledger has a new ledger version every several seconds. When the network agrees on the contents of a ledger version, that ledger version is validated, and its contents can never change. The validated ledger versions that preceded it form the ledger history. Even the most recent validated ledger is part of history, as it represents the state of the network as of a short time ago. In the present, the network is evaluating transactions which may be applied and finalized in the next ledger version. While this evaluation is happening, the network has candidate ledger versions that are not yet validated.

There are three shortcuts for ledger versions:

1. `validated`
1. `closed`
1. `current`

## Ledger API List

There are 5 Ledger APIs:

| API | Description |
|-----|-------------|
| `ledger` | Retrieve information about the public ledger. |
| [`ledger_closed`](ledger_closed/) | The ledger_closed method returns the unique identifiers of the most recently closed ledger. (This ledger is not necessarily validated and immutable yet.) |
| [`ledger_current`](ledger_current/) | The ledger_current method returns the unique identifiers of the current in-progress ledger. This command is mostly useful for testing, because the ledger returned is still in flux. |
| [`ledger_data`](ledger_data/) | The ledger_data method retrieves contents of the specified ledger. You can iterate through several calls to retrieve the entire contents of a single ledger version. |
| `ledger_entry` | The ledger_entry method returns a single ledger object from the XRP Ledger in its raw format. See ledger format for information on the different types of objects you can retrieve. |