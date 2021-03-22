package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrOrderInvalid                 = sdkerrors.Register(ModuleName, 1, "failed to validate order")
	ErrOrderNotFound                = sdkerrors.Register(ModuleName, 2, "no active order found for the hash")
	ErrSpotMarketSuspended          = sdkerrors.Register(ModuleName, 3, "spot market suspended")
	ErrSpotMarketNotFound           = sdkerrors.Register(ModuleName, 4, "spot market not found")
	ErrSpotMarketExists             = sdkerrors.Register(ModuleName, 5, "spot market exists")
	ErrSpotMarketMismatch           = sdkerrors.Register(ModuleName, 6, "spot market mismatch")
	ErrBadField                     = sdkerrors.Register(ModuleName, 7, "struct field error")
	ErrMarketNotFound               = sdkerrors.Register(ModuleName, 8, "derivative market not found")
	ErrMarketInvalid                = sdkerrors.Register(ModuleName, 9, "failed to validate derivative market")
	ErrMarketExists                 = sdkerrors.Register(ModuleName, 10, "market exists")
	ErrMarketSuspended              = sdkerrors.Register(ModuleName, 11, "market suspended")
	ErrBadUpdateEvent               = sdkerrors.Register(ModuleName, 12, "order update event not confirmed")
	ErrUpdateSameValue              = sdkerrors.Register(ModuleName, 13, "cannot update the record's field with the same value")
	ErrOverLeveragedOrder           = sdkerrors.Register(ModuleName, 14, "cannot add overlevered order")
	ErrSubaccountNotFound           = sdkerrors.Register(ModuleName, 15, "subaccount not found")
	ErrOrderAlreadyExists           = sdkerrors.Register(ModuleName, 16, "order already exists")
	ErrInsufficientDeposit          = sdkerrors.Register(ModuleName, 17, "subaccount has insufficient margin")
	ErrMarketExpired                = sdkerrors.Register(ModuleName, 18, "market has already expired")
	ErrOrderExpired                 = sdkerrors.Register(ModuleName, 19, "order has already expired")
	ErrInsufficientOrderQuantity    = sdkerrors.Register(ModuleName, 20, "order quantity invalid")
	ErrUnrecognizedOrderType        = sdkerrors.Register(ModuleName, 21, "unrecognized order type")
	ErrUnfundedPosition             = sdkerrors.Register(ModuleName, 22, "unfunded position for order type")
	ErrInsufficientPositionQuantity = sdkerrors.Register(ModuleName, 23, "position quantity insufficient for order type")
	ErrMarginNotBreached            = sdkerrors.Register(ModuleName, 24, "margin hold is not breached")
	ErrInsufficientTakerMargin      = sdkerrors.Register(ModuleName, 25, "taker has insufficient available margin")
	ErrInsufficientLiquidity        = sdkerrors.Register(ModuleName, 26, "insufficient liquidity in the orderbook")
	ErrReplayTecTransaction         = sdkerrors.Register(ModuleName, 27, "cannot replay TEC transaction")
	ErrOrderAlreadyArchived         = sdkerrors.Register(ModuleName, 28, "order already in archive store")
	ErrAddressNotContract           = sdkerrors.Register(ModuleName, 29, "address is not a smart contract")
	ErrOrderIDInvalid               = sdkerrors.Register(ModuleName, 30, "order id is not valid")
	ErrBadSubaccountID              = sdkerrors.Register(ModuleName, 31, "subaccount id is not valid")
	ErrInvalidTicker                = sdkerrors.Register(ModuleName, 32, "invalid ticker")
	ErrInvalidBaseDenom             = sdkerrors.Register(ModuleName, 33, "invalid base denom")
	ErrInvalidQuoteDenom            = sdkerrors.Register(ModuleName, 34, "invalid quote denom")
	ErrInvalidOracle                = sdkerrors.Register(ModuleName, 35, "invalid oracle")
	ErrInvalidExpiry                = sdkerrors.Register(ModuleName, 36, "invalid expiry")
	ErrInvalidPrice                 = sdkerrors.Register(ModuleName, 37, "invalid price")
	ErrInvalidQuantity              = sdkerrors.Register(ModuleName, 38, "invalid quantity")

	ErrOrderDoesntExist           = sdkerrors.Register(ModuleName, 40, "order doesnt exist")
	ErrDepositDoesntExist         = sdkerrors.Register(ModuleName, 41, "deposit doesnt exist")
	ErrOrderbookDoesntExist       = sdkerrors.Register(ModuleName, 42, "spot limit orderbook doesnt exist")
	ErrOrderbookFillInvalid       = sdkerrors.Register(ModuleName, 43, "spot limit orderbook fill invalid")
	ErrPerpetualMarketExists      = sdkerrors.Register(ModuleName, 44, "perpetual market exists")
	ErrExpiryFuturesMarketExists  = sdkerrors.Register(ModuleName, 45, "expiry futures market exists")
	ErrExpiryFuturesMarketExpired = sdkerrors.Register(ModuleName, 46, "expiry futures market expired")
	ErrNoLiquidity                = sdkerrors.Register(ModuleName, 47, "no liquidity on the orderbook!")
	ErrSlippageExceedsWorstPrice  = sdkerrors.Register(ModuleName, 48, "Orderbook liquidity cannot satisfy current worst price")
	ErrInsufficientOrderMargin    = sdkerrors.Register(ModuleName, 49, "Order has insufficient margin")
	ErrDerivativeMarketNotFound   = sdkerrors.Register(ModuleName, 50, "Derivative market not found")
)