<!--
order: 1
-->

# Concepts

## EVM Gas Denom

In order to use the EVM and be compatible with existing clients, the gas denom used by the EVM must be in 18 decimals. Since `ugert` has 6 decimals of precision, it cannot be used as the EVM gas denom directly.

In order to use the gert token on the EVM, the `evmutil` module provides an `EvmBankKeeper` that is responsible for the conversion of `ugert` and `agert`. A user's `agert` balance is stored in the `evmutil` store, while it's `ugert` balance remains stored in the cosmos-sdk `bank` module.

## `EvmBankKeeper` Overview

The `EvmBankKeeper` provides access to an account's **total** `agert` balance and the ability to transfer, mint, and burn `agert`. If anything other than the `agert` denom is requested, the `EvmBankKeeper` will panic.

```go
type BankKeeper interface {
	evmtypes.BankKeeper
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
}
```

The keeper implements the `x/evm` module's `BankKeeper` interface to enable the usage of `agert` denom on the EVM.

### `x/evm` Parameter

Since the EVM denom `agert` is required to use the `EvmBankKeeper`, it is necessary to set the `EVMDenom` param of the `x/evm` module to `agert`.

### `agert` Balance Calculation

The `agert` balance of an account is derived from an account's **spendable** `ugert` balance times 10^12 (to derive its `agert` equivalent), plus the account's excess `agert` balance that can be accessed by the module `Keeper`.

### Conversions Between `agert` & `ugert`

When an account does not have sufficient `agert` to cover a transfer or burn, the `EvmBankKeeper` will try to swap 1 `ugert` to its equivalent `agert` amount. It does this by transferring 1 `ugert` from the sender to the `evmutil` module account, then adding the equivalent `agert` amount to the sender's balance via the keeper's `AddBalance`.

In reverse, if an account has enough `agert` balance for one or more `ugert`, the excess `agert` balance will be converted to `ugert`. This is done by removing the excess `agert` balance tracked by `evmutil` module store, then transferring the equivalent `ugert` coins from the `evmutil` module account to the target account.

The swap logic ensures that all `agert` is backed by the equivalent `ugert` balance stored in the module account.

## Module Keeper

The module Keeper provides access to an account's excess `agert` balance and the ability to update the balance.
