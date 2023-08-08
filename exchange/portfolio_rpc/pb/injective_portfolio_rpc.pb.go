// Code generated with goa v3.5.2, DO NOT EDIT.
//
// InjectivePortfolioRPC protocol buffer definition
//
// Command:
// $ goa gen github.com/InjectiveLabs/injective-indexer/api/design -o ../

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.19.4
// source: injective_portfolio_rpc.proto

package injective_portfolio_rpcpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AccountPortfolioRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Account address
	AccountAddress string `protobuf:"bytes,1,opt,name=account_address,json=accountAddress,proto3" json:"account_address,omitempty"`
}

func (x *AccountPortfolioRequest) Reset() {
	*x = AccountPortfolioRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_injective_portfolio_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountPortfolioRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountPortfolioRequest) ProtoMessage() {}

func (x *AccountPortfolioRequest) ProtoReflect() protoreflect.Message {
	mi := &file_injective_portfolio_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountPortfolioRequest.ProtoReflect.Descriptor instead.
func (*AccountPortfolioRequest) Descriptor() ([]byte, []int) {
	return file_injective_portfolio_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *AccountPortfolioRequest) GetAccountAddress() string {
	if x != nil {
		return x.AccountAddress
	}
	return ""
}

type AccountPortfolioResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The portfolio of this account
	Portfolio *Portfolio `protobuf:"bytes,1,opt,name=portfolio,proto3" json:"portfolio,omitempty"`
}

func (x *AccountPortfolioResponse) Reset() {
	*x = AccountPortfolioResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_injective_portfolio_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountPortfolioResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountPortfolioResponse) ProtoMessage() {}

func (x *AccountPortfolioResponse) ProtoReflect() protoreflect.Message {
	mi := &file_injective_portfolio_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountPortfolioResponse.ProtoReflect.Descriptor instead.
func (*AccountPortfolioResponse) Descriptor() ([]byte, []int) {
	return file_injective_portfolio_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *AccountPortfolioResponse) GetPortfolio() *Portfolio {
	if x != nil {
		return x.Portfolio
	}
	return nil
}

type Portfolio struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The account's portfolio address
	AccountAddress string `protobuf:"bytes,1,opt,name=account_address,json=accountAddress,proto3" json:"account_address,omitempty"`
	// Account available bank balances
	BankBalances []*Coin `protobuf:"bytes,2,rep,name=bank_balances,json=bankBalances,proto3" json:"bank_balances,omitempty"`
	// Subaccounts list
	Subaccounts []*SubaccountBalanceV2 `protobuf:"bytes,3,rep,name=subaccounts,proto3" json:"subaccounts,omitempty"`
	// All positions for all subaccounts, with unrealized PNL
	PositionsWithUpnl []*PositionsWithUPNL `protobuf:"bytes,4,rep,name=positions_with_upnl,json=positionsWithUpnl,proto3" json:"positions_with_upnl,omitempty"`
}

func (x *Portfolio) Reset() {
	*x = Portfolio{}
	if protoimpl.UnsafeEnabled {
		mi := &file_injective_portfolio_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Portfolio) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Portfolio) ProtoMessage() {}

func (x *Portfolio) ProtoReflect() protoreflect.Message {
	mi := &file_injective_portfolio_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Portfolio.ProtoReflect.Descriptor instead.
func (*Portfolio) Descriptor() ([]byte, []int) {
	return file_injective_portfolio_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *Portfolio) GetAccountAddress() string {
	if x != nil {
		return x.AccountAddress
	}
	return ""
}

func (x *Portfolio) GetBankBalances() []*Coin {
	if x != nil {
		return x.BankBalances
	}
	return nil
}

func (x *Portfolio) GetSubaccounts() []*SubaccountBalanceV2 {
	if x != nil {
		return x.Subaccounts
	}
	return nil
}

func (x *Portfolio) GetPositionsWithUpnl() []*PositionsWithUPNL {
	if x != nil {
		return x.PositionsWithUpnl
	}
	return nil
}

type Coin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Denom of the coin
	Denom  string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Amount string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Coin) Reset() {
	*x = Coin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_injective_portfolio_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coin) ProtoMessage() {}

func (x *Coin) ProtoReflect() protoreflect.Message {
	mi := &file_injective_portfolio_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coin.ProtoReflect.Descriptor instead.
func (*Coin) Descriptor() ([]byte, []int) {
	return file_injective_portfolio_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *Coin) GetDenom() string {
	if x != nil {
		return x.Denom
	}
	return ""
}

func (x *Coin) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

type SubaccountBalanceV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Related subaccount ID
	SubaccountId string `protobuf:"bytes,1,opt,name=subaccount_id,json=subaccountId,proto3" json:"subaccount_id,omitempty"`
	// Coin denom on the chain.
	Denom   string             `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
	Deposit *SubaccountDeposit `protobuf:"bytes,3,opt,name=deposit,proto3" json:"deposit,omitempty"`
}

func (x *SubaccountBalanceV2) Reset() {
	*x = SubaccountBalanceV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_injective_portfolio_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubaccountBalanceV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubaccountBalanceV2) ProtoMessage() {}

func (x *SubaccountBalanceV2) ProtoReflect() protoreflect.Message {
	mi := &file_injective_portfolio_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubaccountBalanceV2.ProtoReflect.Descriptor instead.
func (*SubaccountBalanceV2) Descriptor() ([]byte, []int) {
	return file_injective_portfolio_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *SubaccountBalanceV2) GetSubaccountId() string {
	if x != nil {
		return x.SubaccountId
	}
	return ""
}

func (x *SubaccountBalanceV2) GetDenom() string {
	if x != nil {
		return x.Denom
	}
	return ""
}

func (x *SubaccountBalanceV2) GetDeposit() *SubaccountDeposit {
	if x != nil {
		return x.Deposit
	}
	return nil
}

type SubaccountDeposit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalBalance     string `protobuf:"bytes,1,opt,name=total_balance,json=totalBalance,proto3" json:"total_balance,omitempty"`
	AvailableBalance string `protobuf:"bytes,2,opt,name=available_balance,json=availableBalance,proto3" json:"available_balance,omitempty"`
}

func (x *SubaccountDeposit) Reset() {
	*x = SubaccountDeposit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_injective_portfolio_rpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubaccountDeposit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubaccountDeposit) ProtoMessage() {}

func (x *SubaccountDeposit) ProtoReflect() protoreflect.Message {
	mi := &file_injective_portfolio_rpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubaccountDeposit.ProtoReflect.Descriptor instead.
func (*SubaccountDeposit) Descriptor() ([]byte, []int) {
	return file_injective_portfolio_rpc_proto_rawDescGZIP(), []int{5}
}

func (x *SubaccountDeposit) GetTotalBalance() string {
	if x != nil {
		return x.TotalBalance
	}
	return ""
}

func (x *SubaccountDeposit) GetAvailableBalance() string {
	if x != nil {
		return x.AvailableBalance
	}
	return ""
}

type PositionsWithUPNL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position *DerivativePosition `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	// Unrealized PNL
	UnrealizedPnl string `protobuf:"bytes,2,opt,name=unrealized_pnl,json=unrealizedPnl,proto3" json:"unrealized_pnl,omitempty"`
}

func (x *PositionsWithUPNL) Reset() {
	*x = PositionsWithUPNL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_injective_portfolio_rpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionsWithUPNL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionsWithUPNL) ProtoMessage() {}

func (x *PositionsWithUPNL) ProtoReflect() protoreflect.Message {
	mi := &file_injective_portfolio_rpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PositionsWithUPNL.ProtoReflect.Descriptor instead.
func (*PositionsWithUPNL) Descriptor() ([]byte, []int) {
	return file_injective_portfolio_rpc_proto_rawDescGZIP(), []int{6}
}

func (x *PositionsWithUPNL) GetPosition() *DerivativePosition {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *PositionsWithUPNL) GetUnrealizedPnl() string {
	if x != nil {
		return x.UnrealizedPnl
	}
	return ""
}

type DerivativePosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Ticker of the derivative market
	Ticker string `protobuf:"bytes,1,opt,name=ticker,proto3" json:"ticker,omitempty"`
	// Derivative Market ID
	MarketId string `protobuf:"bytes,2,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty"`
	// The subaccountId that the position belongs to
	SubaccountId string `protobuf:"bytes,3,opt,name=subaccount_id,json=subaccountId,proto3" json:"subaccount_id,omitempty"`
	// Direction of the position
	Direction string `protobuf:"bytes,4,opt,name=direction,proto3" json:"direction,omitempty"`
	// Quantity of the position
	Quantity string `protobuf:"bytes,5,opt,name=quantity,proto3" json:"quantity,omitempty"`
	// Price of the position
	EntryPrice string `protobuf:"bytes,6,opt,name=entry_price,json=entryPrice,proto3" json:"entry_price,omitempty"`
	// Margin of the position
	Margin string `protobuf:"bytes,7,opt,name=margin,proto3" json:"margin,omitempty"`
	// LiquidationPrice of the position
	LiquidationPrice string `protobuf:"bytes,8,opt,name=liquidation_price,json=liquidationPrice,proto3" json:"liquidation_price,omitempty"`
	// MarkPrice of the position
	MarkPrice string `protobuf:"bytes,9,opt,name=mark_price,json=markPrice,proto3" json:"mark_price,omitempty"`
	// Aggregate Quantity of the Reduce Only orders associated with the position
	AggregateReduceOnlyQuantity string `protobuf:"bytes,11,opt,name=aggregate_reduce_only_quantity,json=aggregateReduceOnlyQuantity,proto3" json:"aggregate_reduce_only_quantity,omitempty"`
	// Position updated timestamp in UNIX millis.
	UpdatedAt int64 `protobuf:"zigzag64,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Position created timestamp in UNIX millis.
	CreatedAt int64 `protobuf:"zigzag64,13,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *DerivativePosition) Reset() {
	*x = DerivativePosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_injective_portfolio_rpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DerivativePosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DerivativePosition) ProtoMessage() {}

func (x *DerivativePosition) ProtoReflect() protoreflect.Message {
	mi := &file_injective_portfolio_rpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DerivativePosition.ProtoReflect.Descriptor instead.
func (*DerivativePosition) Descriptor() ([]byte, []int) {
	return file_injective_portfolio_rpc_proto_rawDescGZIP(), []int{7}
}

func (x *DerivativePosition) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *DerivativePosition) GetMarketId() string {
	if x != nil {
		return x.MarketId
	}
	return ""
}

func (x *DerivativePosition) GetSubaccountId() string {
	if x != nil {
		return x.SubaccountId
	}
	return ""
}

func (x *DerivativePosition) GetDirection() string {
	if x != nil {
		return x.Direction
	}
	return ""
}

func (x *DerivativePosition) GetQuantity() string {
	if x != nil {
		return x.Quantity
	}
	return ""
}

func (x *DerivativePosition) GetEntryPrice() string {
	if x != nil {
		return x.EntryPrice
	}
	return ""
}

func (x *DerivativePosition) GetMargin() string {
	if x != nil {
		return x.Margin
	}
	return ""
}

func (x *DerivativePosition) GetLiquidationPrice() string {
	if x != nil {
		return x.LiquidationPrice
	}
	return ""
}

func (x *DerivativePosition) GetMarkPrice() string {
	if x != nil {
		return x.MarkPrice
	}
	return ""
}

func (x *DerivativePosition) GetAggregateReduceOnlyQuantity() string {
	if x != nil {
		return x.AggregateReduceOnlyQuantity
	}
	return ""
}

func (x *DerivativePosition) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *DerivativePosition) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type StreamAccountPortfolioRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The account's portfolio address
	AccountAddress string `protobuf:"bytes,1,opt,name=account_address,json=accountAddress,proto3" json:"account_address,omitempty"`
	// Related subaccount ID
	SubaccountId string `protobuf:"bytes,2,opt,name=subaccount_id,json=subaccountId,proto3" json:"subaccount_id,omitempty"`
	Type         string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *StreamAccountPortfolioRequest) Reset() {
	*x = StreamAccountPortfolioRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_injective_portfolio_rpc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamAccountPortfolioRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamAccountPortfolioRequest) ProtoMessage() {}

func (x *StreamAccountPortfolioRequest) ProtoReflect() protoreflect.Message {
	mi := &file_injective_portfolio_rpc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamAccountPortfolioRequest.ProtoReflect.Descriptor instead.
func (*StreamAccountPortfolioRequest) Descriptor() ([]byte, []int) {
	return file_injective_portfolio_rpc_proto_rawDescGZIP(), []int{8}
}

func (x *StreamAccountPortfolioRequest) GetAccountAddress() string {
	if x != nil {
		return x.AccountAddress
	}
	return ""
}

func (x *StreamAccountPortfolioRequest) GetSubaccountId() string {
	if x != nil {
		return x.SubaccountId
	}
	return ""
}

func (x *StreamAccountPortfolioRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type StreamAccountPortfolioResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// type of portfolio entry
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// denom of portfolio entry
	Denom string `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
	// amount of portfolio entry
	Amount string `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	// subaccount id of portfolio entry
	SubaccountId string `protobuf:"bytes,4,opt,name=subaccount_id,json=subaccountId,proto3" json:"subaccount_id,omitempty"`
	// Operation timestamp in UNIX millis.
	Timestamp int64 `protobuf:"zigzag64,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *StreamAccountPortfolioResponse) Reset() {
	*x = StreamAccountPortfolioResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_injective_portfolio_rpc_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamAccountPortfolioResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamAccountPortfolioResponse) ProtoMessage() {}

func (x *StreamAccountPortfolioResponse) ProtoReflect() protoreflect.Message {
	mi := &file_injective_portfolio_rpc_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamAccountPortfolioResponse.ProtoReflect.Descriptor instead.
func (*StreamAccountPortfolioResponse) Descriptor() ([]byte, []int) {
	return file_injective_portfolio_rpc_proto_rawDescGZIP(), []int{9}
}

func (x *StreamAccountPortfolioResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *StreamAccountPortfolioResponse) GetDenom() string {
	if x != nil {
		return x.Denom
	}
	return ""
}

func (x *StreamAccountPortfolioResponse) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *StreamAccountPortfolioResponse) GetSubaccountId() string {
	if x != nil {
		return x.SubaccountId
	}
	return ""
}

func (x *StreamAccountPortfolioResponse) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_injective_portfolio_rpc_proto protoreflect.FileDescriptor

var file_injective_portfolio_rpc_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74,
	0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x17, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x66,
	0x6f, 0x6c, 0x69, 0x6f, 0x5f, 0x72, 0x70, 0x63, 0x22, 0x42, 0x0a, 0x17, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x5c, 0x0a, 0x18,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x09, 0x70, 0x6f, 0x72, 0x74,
	0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e,
	0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69,
	0x6f, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x52,
	0x09, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x22, 0xa4, 0x02, 0x0a, 0x09, 0x50,
	0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x42, 0x0a, 0x0d, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x69, 0x6e, 0x6a, 0x65, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x5f, 0x72,
	0x70, 0x63, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x0c, 0x62, 0x61, 0x6e, 0x6b, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x4e, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x69, 0x6e, 0x6a,
	0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f,
	0x5f, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x75, 0x62, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x56, 0x32, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x5a, 0x0a, 0x13, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x75, 0x70, 0x6e, 0x6c, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70,
	0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x57, 0x69, 0x74, 0x68, 0x55, 0x50, 0x4e, 0x4c, 0x52, 0x11,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x57, 0x69, 0x74, 0x68, 0x55, 0x70, 0x6e,
	0x6c, 0x22, 0x34, 0x0a, 0x04, 0x43, 0x6f, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x6e,
	0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x96, 0x01, 0x0a, 0x13, 0x53, 0x75, 0x62, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x56, 0x32, 0x12,
	0x23, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x44, 0x0a, 0x07, 0x64, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x69, 0x6e,
	0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69,
	0x6f, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x75, 0x62, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x07, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x22, 0x65, 0x0a, 0x11, 0x53, 0x75, 0x62, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x61, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x83, 0x01, 0x0a, 0x11, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x57, 0x69, 0x74, 0x68, 0x55, 0x50, 0x4e, 0x4c, 0x12, 0x47, 0x0a,
	0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2b, 0x2e, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74,
	0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x69, 0x76, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x64, 0x5f, 0x70, 0x6e, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x75, 0x6e, 0x72, 0x65, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x50, 0x6e, 0x6c, 0x22, 0xb0, 0x03,
	0x0a, 0x12, 0x44, 0x65, 0x72, 0x69, 0x76, 0x61, 0x74, 0x69, 0x76, 0x65, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x75, 0x62,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x73, 0x75, 0x62, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x74, 0x72,
	0x79, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x61, 0x72,
	0x67, 0x69, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x61, 0x72, 0x67, 0x69,
	0x6e, 0x12, 0x2b, 0x0a, 0x11, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6c, 0x69,
	0x71, 0x75, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x6d, 0x61, 0x72, 0x6b, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x72, 0x6b, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a,
	0x1e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x64, 0x75, 0x63,
	0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1b, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x12, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x12, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0x81, 0x01, 0x0a, 0x1d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x73,
	0x75, 0x62, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x22, 0xa5, 0x01, 0x0a, 0x1e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x64,
	0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x65, 0x6e, 0x6f,
	0x6d, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x75, 0x62,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x73, 0x75, 0x62, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x12, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x32, 0x9e, 0x02, 0x0a,
	0x15, 0x49, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f,
	0x6c, 0x69, 0x6f, 0x52, 0x50, 0x43, 0x12, 0x77, 0x0a, 0x10, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x12, 0x30, 0x2e, 0x69, 0x6e, 0x6a,
	0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f,
	0x5f, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x6f, 0x72, 0x74,
	0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x69,
	0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c,
	0x69, 0x6f, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x6f,
	0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x8b, 0x01, 0x0a, 0x16, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x12, 0x36, 0x2e, 0x69, 0x6e, 0x6a,
	0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f,
	0x5f, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x37, 0x2e, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70,
	0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f,
	0x6c, 0x69, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x1c, 0x5a,
	0x1a, 0x2f, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74,
	0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x5f, 0x72, 0x70, 0x63, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_injective_portfolio_rpc_proto_rawDescOnce sync.Once
	file_injective_portfolio_rpc_proto_rawDescData = file_injective_portfolio_rpc_proto_rawDesc
)

func file_injective_portfolio_rpc_proto_rawDescGZIP() []byte {
	file_injective_portfolio_rpc_proto_rawDescOnce.Do(func() {
		file_injective_portfolio_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_injective_portfolio_rpc_proto_rawDescData)
	})
	return file_injective_portfolio_rpc_proto_rawDescData
}

var file_injective_portfolio_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_injective_portfolio_rpc_proto_goTypes = []interface{}{
	(*AccountPortfolioRequest)(nil),        // 0: injective_portfolio_rpc.AccountPortfolioRequest
	(*AccountPortfolioResponse)(nil),       // 1: injective_portfolio_rpc.AccountPortfolioResponse
	(*Portfolio)(nil),                      // 2: injective_portfolio_rpc.Portfolio
	(*Coin)(nil),                           // 3: injective_portfolio_rpc.Coin
	(*SubaccountBalanceV2)(nil),            // 4: injective_portfolio_rpc.SubaccountBalanceV2
	(*SubaccountDeposit)(nil),              // 5: injective_portfolio_rpc.SubaccountDeposit
	(*PositionsWithUPNL)(nil),              // 6: injective_portfolio_rpc.PositionsWithUPNL
	(*DerivativePosition)(nil),             // 7: injective_portfolio_rpc.DerivativePosition
	(*StreamAccountPortfolioRequest)(nil),  // 8: injective_portfolio_rpc.StreamAccountPortfolioRequest
	(*StreamAccountPortfolioResponse)(nil), // 9: injective_portfolio_rpc.StreamAccountPortfolioResponse
}
var file_injective_portfolio_rpc_proto_depIdxs = []int32{
	2, // 0: injective_portfolio_rpc.AccountPortfolioResponse.portfolio:type_name -> injective_portfolio_rpc.Portfolio
	3, // 1: injective_portfolio_rpc.Portfolio.bank_balances:type_name -> injective_portfolio_rpc.Coin
	4, // 2: injective_portfolio_rpc.Portfolio.subaccounts:type_name -> injective_portfolio_rpc.SubaccountBalanceV2
	6, // 3: injective_portfolio_rpc.Portfolio.positions_with_upnl:type_name -> injective_portfolio_rpc.PositionsWithUPNL
	5, // 4: injective_portfolio_rpc.SubaccountBalanceV2.deposit:type_name -> injective_portfolio_rpc.SubaccountDeposit
	7, // 5: injective_portfolio_rpc.PositionsWithUPNL.position:type_name -> injective_portfolio_rpc.DerivativePosition
	0, // 6: injective_portfolio_rpc.InjectivePortfolioRPC.AccountPortfolio:input_type -> injective_portfolio_rpc.AccountPortfolioRequest
	8, // 7: injective_portfolio_rpc.InjectivePortfolioRPC.StreamAccountPortfolio:input_type -> injective_portfolio_rpc.StreamAccountPortfolioRequest
	1, // 8: injective_portfolio_rpc.InjectivePortfolioRPC.AccountPortfolio:output_type -> injective_portfolio_rpc.AccountPortfolioResponse
	9, // 9: injective_portfolio_rpc.InjectivePortfolioRPC.StreamAccountPortfolio:output_type -> injective_portfolio_rpc.StreamAccountPortfolioResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_injective_portfolio_rpc_proto_init() }
func file_injective_portfolio_rpc_proto_init() {
	if File_injective_portfolio_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_injective_portfolio_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountPortfolioRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_injective_portfolio_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountPortfolioResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_injective_portfolio_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Portfolio); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_injective_portfolio_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Coin); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_injective_portfolio_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubaccountBalanceV2); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_injective_portfolio_rpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubaccountDeposit); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_injective_portfolio_rpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PositionsWithUPNL); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_injective_portfolio_rpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DerivativePosition); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_injective_portfolio_rpc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamAccountPortfolioRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_injective_portfolio_rpc_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamAccountPortfolioResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_injective_portfolio_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_injective_portfolio_rpc_proto_goTypes,
		DependencyIndexes: file_injective_portfolio_rpc_proto_depIdxs,
		MessageInfos:      file_injective_portfolio_rpc_proto_msgTypes,
	}.Build()
	File_injective_portfolio_rpc_proto = out.File
	file_injective_portfolio_rpc_proto_rawDesc = nil
	file_injective_portfolio_rpc_proto_goTypes = nil
	file_injective_portfolio_rpc_proto_depIdxs = nil
}