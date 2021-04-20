// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: injective/oracle/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the oracle module's genesis state.
type GenesisState struct {
	// params defines all the parameters of related to oracle.
	Params                Params                  `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	BandRelayers          []string                `protobuf:"bytes,2,rep,name=band_relayers,json=bandRelayers,proto3" json:"band_relayers,omitempty"`
	BandRefDataState      []BandRefDataState      `protobuf:"bytes,3,rep,name=band_ref_data_state,json=bandRefDataState,proto3" json:"band_ref_data_state"`
	PriceFeedRelayerState []PriceFeedRelayerState `protobuf:"bytes,4,rep,name=price_feed_relayer_state,json=priceFeedRelayerState,proto3" json:"price_feed_relayer_state"`
	PriceFeedPriceState   []PriceFeedPriceState   `protobuf:"bytes,5,rep,name=price_feed_price_state,json=priceFeedPriceState,proto3" json:"price_feed_price_state"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7e14cf80151b4d2, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetBandRelayers() []string {
	if m != nil {
		return m.BandRelayers
	}
	return nil
}

func (m *GenesisState) GetBandRefDataState() []BandRefDataState {
	if m != nil {
		return m.BandRefDataState
	}
	return nil
}

func (m *GenesisState) GetPriceFeedRelayerState() []PriceFeedRelayerState {
	if m != nil {
		return m.PriceFeedRelayerState
	}
	return nil
}

func (m *GenesisState) GetPriceFeedPriceState() []PriceFeedPriceState {
	if m != nil {
		return m.PriceFeedPriceState
	}
	return nil
}

type BandRefDataState struct {
	Symbol string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Ref    *Ref   `protobuf:"bytes,2,opt,name=ref,proto3" json:"ref,omitempty"`
}

func (m *BandRefDataState) Reset()         { *m = BandRefDataState{} }
func (m *BandRefDataState) String() string { return proto.CompactTextString(m) }
func (*BandRefDataState) ProtoMessage()    {}
func (*BandRefDataState) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7e14cf80151b4d2, []int{1}
}
func (m *BandRefDataState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BandRefDataState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BandRefDataState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BandRefDataState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BandRefDataState.Merge(m, src)
}
func (m *BandRefDataState) XXX_Size() int {
	return m.Size()
}
func (m *BandRefDataState) XXX_DiscardUnknown() {
	xxx_messageInfo_BandRefDataState.DiscardUnknown(m)
}

var xxx_messageInfo_BandRefDataState proto.InternalMessageInfo

func (m *BandRefDataState) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *BandRefDataState) GetRef() *Ref {
	if m != nil {
		return m.Ref
	}
	return nil
}

type PriceFeedRelayerState struct {
	BaseQuoteHash string   `protobuf:"bytes,1,opt,name=base_quote_hash,json=baseQuoteHash,proto3" json:"base_quote_hash,omitempty"`
	Relayers      []string `protobuf:"bytes,2,rep,name=relayers,proto3" json:"relayers,omitempty"`
}

func (m *PriceFeedRelayerState) Reset()         { *m = PriceFeedRelayerState{} }
func (m *PriceFeedRelayerState) String() string { return proto.CompactTextString(m) }
func (*PriceFeedRelayerState) ProtoMessage()    {}
func (*PriceFeedRelayerState) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7e14cf80151b4d2, []int{2}
}
func (m *PriceFeedRelayerState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PriceFeedRelayerState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PriceFeedRelayerState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PriceFeedRelayerState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PriceFeedRelayerState.Merge(m, src)
}
func (m *PriceFeedRelayerState) XXX_Size() int {
	return m.Size()
}
func (m *PriceFeedRelayerState) XXX_DiscardUnknown() {
	xxx_messageInfo_PriceFeedRelayerState.DiscardUnknown(m)
}

var xxx_messageInfo_PriceFeedRelayerState proto.InternalMessageInfo

func (m *PriceFeedRelayerState) GetBaseQuoteHash() string {
	if m != nil {
		return m.BaseQuoteHash
	}
	return ""
}

func (m *PriceFeedRelayerState) GetRelayers() []string {
	if m != nil {
		return m.Relayers
	}
	return nil
}

type PriceFeedPriceState struct {
	BaseQuoteHash string                                 `protobuf:"bytes,1,opt,name=base_quote_hash,json=baseQuoteHash,proto3" json:"base_quote_hash,omitempty"`
	Price         github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=price,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"price"`
}

func (m *PriceFeedPriceState) Reset()         { *m = PriceFeedPriceState{} }
func (m *PriceFeedPriceState) String() string { return proto.CompactTextString(m) }
func (*PriceFeedPriceState) ProtoMessage()    {}
func (*PriceFeedPriceState) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7e14cf80151b4d2, []int{3}
}
func (m *PriceFeedPriceState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PriceFeedPriceState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PriceFeedPriceState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PriceFeedPriceState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PriceFeedPriceState.Merge(m, src)
}
func (m *PriceFeedPriceState) XXX_Size() int {
	return m.Size()
}
func (m *PriceFeedPriceState) XXX_DiscardUnknown() {
	xxx_messageInfo_PriceFeedPriceState.DiscardUnknown(m)
}

var xxx_messageInfo_PriceFeedPriceState proto.InternalMessageInfo

func (m *PriceFeedPriceState) GetBaseQuoteHash() string {
	if m != nil {
		return m.BaseQuoteHash
	}
	return ""
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "injective.oracle.v1beta1.GenesisState")
	proto.RegisterType((*BandRefDataState)(nil), "injective.oracle.v1beta1.BandRefDataState")
	proto.RegisterType((*PriceFeedRelayerState)(nil), "injective.oracle.v1beta1.PriceFeedRelayerState")
	proto.RegisterType((*PriceFeedPriceState)(nil), "injective.oracle.v1beta1.PriceFeedPriceState")
}

func init() {
	proto.RegisterFile("injective/oracle/v1beta1/genesis.proto", fileDescriptor_f7e14cf80151b4d2)
}

var fileDescriptor_f7e14cf80151b4d2 = []byte{
	// 483 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x6a, 0xdb, 0x40,
	0x10, 0xc6, 0xad, 0x28, 0x31, 0xf5, 0x26, 0xa1, 0x61, 0xdd, 0x04, 0x61, 0xa8, 0x62, 0x0c, 0x35,
	0x6e, 0xc1, 0x5a, 0x92, 0xde, 0x7b, 0x30, 0xa6, 0x7f, 0x20, 0x87, 0x56, 0xbd, 0x35, 0x07, 0x31,
	0x92, 0x47, 0x96, 0x5a, 0x4b, 0xab, 0x6a, 0xd7, 0x01, 0x9f, 0xfb, 0x02, 0x7d, 0x98, 0x3e, 0x44,
	0x8e, 0x39, 0x96, 0x1e, 0x42, 0xb1, 0x5f, 0xa4, 0x68, 0xb5, 0x36, 0x6a, 0x6a, 0x15, 0x9f, 0x34,
	0x3b, 0xf3, 0xed, 0xf7, 0x9b, 0x99, 0x45, 0xa4, 0x1f, 0xa7, 0x9f, 0x31, 0x90, 0xf1, 0x0d, 0x32,
	0x9e, 0x43, 0x30, 0x43, 0x76, 0x73, 0xe1, 0xa3, 0x84, 0x0b, 0x36, 0xc5, 0x14, 0x45, 0x2c, 0x9c,
	0x2c, 0xe7, 0x92, 0x53, 0x6b, 0xa3, 0x73, 0x4a, 0x9d, 0xa3, 0x75, 0x9d, 0x67, 0xb5, 0x0e, 0x5a,
	0xa8, 0x0c, 0x3a, 0x4f, 0xa6, 0x7c, 0xca, 0x55, 0xc8, 0x8a, 0xa8, 0xcc, 0xf6, 0x7e, 0x98, 0xe4,
	0xe8, 0x4d, 0x09, 0xfa, 0x28, 0x41, 0x22, 0x7d, 0x45, 0x9a, 0x19, 0xe4, 0x90, 0x08, 0xcb, 0xe8,
	0x1a, 0x83, 0xc3, 0xcb, 0xae, 0x53, 0x07, 0x76, 0xde, 0x2b, 0xdd, 0x68, 0xff, 0xf6, 0xfe, 0xbc,
	0xe1, 0xea, 0x5b, 0xf4, 0x39, 0x39, 0xf6, 0x21, 0x9d, 0x78, 0x39, 0xce, 0x60, 0x81, 0xb9, 0xb0,
	0xf6, 0xba, 0xe6, 0xa0, 0xa5, 0x45, 0x47, 0x45, 0xc9, 0xd5, 0x15, 0xea, 0x91, 0xb6, 0x96, 0x86,
	0xde, 0x04, 0x24, 0x78, 0xa2, 0xe8, 0xc0, 0x32, 0xbb, 0xe6, 0xe0, 0xf0, 0xf2, 0x45, 0x3d, 0x77,
	0xa4, 0x4c, 0xc2, 0x31, 0x48, 0x50, 0x3d, 0x6b, 0xf3, 0x13, 0xff, 0x41, 0x9e, 0xa6, 0xc4, 0xca,
	0xf2, 0x38, 0x40, 0x2f, 0x44, 0xdc, 0x74, 0xa4, 0x29, 0xfb, 0x8a, 0xc2, 0xfe, 0x33, 0x5d, 0x71,
	0xf3, 0x35, 0xe2, 0xba, 0xdf, 0x2a, 0xea, 0x34, 0xdb, 0x56, 0xa4, 0x11, 0x39, 0xab, 0xf0, 0xca,
	0xb0, 0xa4, 0x1d, 0x28, 0xda, 0x70, 0x07, 0x9a, 0x0a, 0xaa, 0xac, 0x76, 0xf6, 0x6f, 0xa9, 0x77,
	0x4d, 0x4e, 0x1e, 0x6e, 0x81, 0x9e, 0x91, 0xa6, 0x58, 0x24, 0x3e, 0x9f, 0xa9, 0x97, 0x6b, 0xb9,
	0xfa, 0x44, 0x19, 0x31, 0x73, 0x0c, 0xad, 0x3d, 0xf5, 0x9c, 0x4f, 0xeb, 0x5b, 0x70, 0x31, 0x74,
	0x0b, 0x65, 0xef, 0x9a, 0x9c, 0x6e, 0x1d, 0x9e, 0xf6, 0xc9, 0x63, 0x1f, 0x04, 0x7a, 0x5f, 0xe7,
	0x5c, 0xa2, 0x17, 0x81, 0x88, 0x34, 0xea, 0xb8, 0x48, 0x7f, 0x28, 0xb2, 0x6f, 0x41, 0x44, 0xb4,
	0x43, 0x1e, 0xfd, 0xfd, 0xfc, 0xee, 0xe6, 0xdc, 0xfb, 0x66, 0x90, 0xf6, 0x96, 0x61, 0x77, 0xf6,
	0x1e, 0x93, 0x03, 0xb5, 0x10, 0x35, 0x4f, 0x6b, 0xe4, 0x14, 0x3b, 0xfa, 0x75, 0x7f, 0xde, 0x9f,
	0xc6, 0x32, 0x9a, 0xfb, 0x4e, 0xc0, 0x13, 0x16, 0x70, 0x91, 0x70, 0xa1, 0x3f, 0x43, 0x31, 0xf9,
	0xc2, 0xe4, 0x22, 0x43, 0xe1, 0x8c, 0x31, 0x70, 0xcb, 0xcb, 0xa3, 0xf0, 0x76, 0x69, 0x1b, 0x77,
	0x4b, 0xdb, 0xf8, 0xbd, 0xb4, 0x8d, 0xef, 0x2b, 0xbb, 0x71, 0xb7, 0xb2, 0x1b, 0x3f, 0x57, 0x76,
	0xe3, 0xd3, 0x55, 0xc5, 0xe8, 0xdd, 0x7a, 0x55, 0x57, 0xe0, 0x0b, 0xb6, 0x59, 0xdc, 0x30, 0xe0,
	0x39, 0x56, 0x8f, 0x11, 0xc4, 0x29, 0x4b, 0xf8, 0x64, 0x3e, 0x43, 0xb1, 0xfe, 0x07, 0x15, 0xd2,
	0x6f, 0xaa, 0xbf, 0xec, 0xe5, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x9e, 0xf9, 0xab, 0xe6,
	0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PriceFeedPriceState) > 0 {
		for iNdEx := len(m.PriceFeedPriceState) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PriceFeedPriceState[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.PriceFeedRelayerState) > 0 {
		for iNdEx := len(m.PriceFeedRelayerState) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PriceFeedRelayerState[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.BandRefDataState) > 0 {
		for iNdEx := len(m.BandRefDataState) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BandRefDataState[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.BandRelayers) > 0 {
		for iNdEx := len(m.BandRelayers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.BandRelayers[iNdEx])
			copy(dAtA[i:], m.BandRelayers[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.BandRelayers[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *BandRefDataState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BandRefDataState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BandRefDataState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Ref != nil {
		{
			size, err := m.Ref.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PriceFeedRelayerState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PriceFeedRelayerState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PriceFeedRelayerState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Relayers) > 0 {
		for iNdEx := len(m.Relayers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Relayers[iNdEx])
			copy(dAtA[i:], m.Relayers[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.Relayers[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.BaseQuoteHash) > 0 {
		i -= len(m.BaseQuoteHash)
		copy(dAtA[i:], m.BaseQuoteHash)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.BaseQuoteHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PriceFeedPriceState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PriceFeedPriceState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PriceFeedPriceState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Price.Size()
		i -= size
		if _, err := m.Price.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.BaseQuoteHash) > 0 {
		i -= len(m.BaseQuoteHash)
		copy(dAtA[i:], m.BaseQuoteHash)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.BaseQuoteHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.BandRelayers) > 0 {
		for _, s := range m.BandRelayers {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BandRefDataState) > 0 {
		for _, e := range m.BandRefDataState {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PriceFeedRelayerState) > 0 {
		for _, e := range m.PriceFeedRelayerState {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PriceFeedPriceState) > 0 {
		for _, e := range m.PriceFeedPriceState {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *BandRefDataState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Ref != nil {
		l = m.Ref.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *PriceFeedRelayerState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BaseQuoteHash)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.Relayers) > 0 {
		for _, s := range m.Relayers {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *PriceFeedPriceState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BaseQuoteHash)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.Price.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BandRelayers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BandRelayers = append(m.BandRelayers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BandRefDataState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BandRefDataState = append(m.BandRefDataState, BandRefDataState{})
			if err := m.BandRefDataState[len(m.BandRefDataState)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceFeedRelayerState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PriceFeedRelayerState = append(m.PriceFeedRelayerState, PriceFeedRelayerState{})
			if err := m.PriceFeedRelayerState[len(m.PriceFeedRelayerState)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceFeedPriceState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PriceFeedPriceState = append(m.PriceFeedPriceState, PriceFeedPriceState{})
			if err := m.PriceFeedPriceState[len(m.PriceFeedPriceState)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BandRefDataState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BandRefDataState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BandRefDataState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ref", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Ref == nil {
				m.Ref = &Ref{}
			}
			if err := m.Ref.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PriceFeedRelayerState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PriceFeedRelayerState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PriceFeedRelayerState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseQuoteHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseQuoteHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Relayers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Relayers = append(m.Relayers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PriceFeedPriceState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PriceFeedPriceState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PriceFeedPriceState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseQuoteHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseQuoteHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)