// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: injective/insurance/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the insurance module's genesis state.
type GenesisState struct {
	// params defines all the parameters of related to insurance.
	Params             Params                `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	InsuranceFunds     []InsuranceFund       `protobuf:"bytes,2,rep,name=insurance_funds,json=insuranceFunds,proto3" json:"insurance_funds"`
	RedemptionSchedule []*RedemptionSchedule `protobuf:"bytes,3,rep,name=redemption_schedule,json=redemptionSchedule,proto3" json:"redemption_schedule,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_293324fee7d3f3b1, []int{0}
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

func (m *GenesisState) GetInsuranceFunds() []InsuranceFund {
	if m != nil {
		return m.InsuranceFunds
	}
	return nil
}

func (m *GenesisState) GetRedemptionSchedule() []*RedemptionSchedule {
	if m != nil {
		return m.RedemptionSchedule
	}
	return nil
}

type Redemption struct {
	Redeemer string             `protobuf:"bytes,1,opt,name=redeemer,proto3" json:"redeemer,omitempty"`
	Schedule RedemptionSchedule `protobuf:"bytes,2,opt,name=schedule,proto3" json:"schedule"`
}

func (m *Redemption) Reset()         { *m = Redemption{} }
func (m *Redemption) String() string { return proto.CompactTextString(m) }
func (*Redemption) ProtoMessage()    {}
func (*Redemption) Descriptor() ([]byte, []int) {
	return fileDescriptor_293324fee7d3f3b1, []int{1}
}
func (m *Redemption) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Redemption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Redemption.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Redemption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Redemption.Merge(m, src)
}
func (m *Redemption) XXX_Size() int {
	return m.Size()
}
func (m *Redemption) XXX_DiscardUnknown() {
	xxx_messageInfo_Redemption.DiscardUnknown(m)
}

var xxx_messageInfo_Redemption proto.InternalMessageInfo

func (m *Redemption) GetRedeemer() string {
	if m != nil {
		return m.Redeemer
	}
	return ""
}

func (m *Redemption) GetSchedule() RedemptionSchedule {
	if m != nil {
		return m.Schedule
	}
	return RedemptionSchedule{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "injective.insurance.v1beta1.GenesisState")
	proto.RegisterType((*Redemption)(nil), "injective.insurance.v1beta1.Redemption")
}

func init() {
	proto.RegisterFile("injective/insurance/v1beta1/genesis.proto", fileDescriptor_293324fee7d3f3b1)
}

var fileDescriptor_293324fee7d3f3b1 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xdb, 0x4d, 0xc6, 0xcc, 0x44, 0xa1, 0x7a, 0x18, 0x13, 0xea, 0x98, 0x97, 0xa9, 0xd8,
	0xb0, 0xf9, 0x09, 0xdc, 0x41, 0x19, 0x08, 0x6a, 0x77, 0xd2, 0xcb, 0x4c, 0xdb, 0x67, 0x17, 0xb1,
	0x49, 0x49, 0xd2, 0x81, 0xf8, 0x1d, 0xc4, 0x8f, 0xb5, 0xe3, 0x8e, 0x9e, 0x44, 0xb6, 0x2f, 0x22,
	0xcb, 0xba, 0x6c, 0x20, 0x14, 0xbc, 0x25, 0x7f, 0x7e, 0xef, 0xf7, 0x5e, 0x5f, 0x83, 0x4e, 0x28,
	0x7b, 0x81, 0x50, 0xd1, 0x31, 0x60, 0xca, 0x64, 0x26, 0x08, 0x0b, 0x01, 0x8f, 0x3b, 0x01, 0x28,
	0xd2, 0xc1, 0x31, 0x30, 0x90, 0x54, 0x7a, 0xa9, 0xe0, 0x8a, 0x3b, 0x87, 0x06, 0xf5, 0x0c, 0xea,
	0xe5, 0x68, 0xe3, 0xac, 0xc8, 0xb3, 0xc6, 0xb5, 0xa9, 0x71, 0x10, 0xf3, 0x98, 0xeb, 0x23, 0x5e,
	0x9c, 0x96, 0x69, 0xeb, 0xa3, 0x84, 0x76, 0xae, 0x97, 0x1d, 0x07, 0x8a, 0x28, 0x70, 0x2e, 0x51,
	0x25, 0x25, 0x82, 0x24, 0xb2, 0x6e, 0x37, 0xed, 0x76, 0xad, 0x7b, 0xec, 0x15, 0x4c, 0xe0, 0xdd,
	0x69, 0xb4, 0xb7, 0x35, 0xf9, 0x3e, 0xb2, 0xfc, 0xbc, 0xd0, 0x79, 0x40, 0x7b, 0x86, 0x1c, 0x3e,
	0x67, 0x2c, 0x92, 0xf5, 0x52, 0xb3, 0xdc, 0xae, 0x75, 0x4f, 0x0b, 0x5d, 0xfd, 0x55, 0x72, 0x95,
	0xb1, 0x28, 0x57, 0xee, 0xd2, 0xcd, 0x50, 0x3a, 0x4f, 0x68, 0x5f, 0x40, 0x04, 0x49, 0xaa, 0x28,
	0x67, 0x43, 0x19, 0x8e, 0x20, 0xca, 0x5e, 0xa1, 0x5e, 0xd6, 0x7a, 0x5c, 0xa8, 0xf7, 0x4d, 0xdd,
	0x20, 0x2f, 0xf3, 0x1d, 0xf1, 0x27, 0x6b, 0xbd, 0x23, 0xb4, 0x26, 0x9d, 0x06, 0xaa, 0x2e, 0x18,
	0x48, 0x40, 0xe8, 0x7d, 0x6c, 0xfb, 0xe6, 0xee, 0xdc, 0xa3, 0xaa, 0x19, 0xa0, 0xa4, 0x77, 0xf5,
	0xdf, 0x01, 0xf2, 0x8f, 0x34, 0x9a, 0x1e, 0x9d, 0xcc, 0x5c, 0x7b, 0x3a, 0x73, 0xed, 0x9f, 0x99,
	0x6b, 0x7f, 0xce, 0x5d, 0x6b, 0x3a, 0x77, 0xad, 0xaf, 0xb9, 0x6b, 0x3d, 0xde, 0xc6, 0x54, 0x8d,
	0xb2, 0xc0, 0x0b, 0x79, 0x82, 0xfb, 0xab, 0x26, 0x37, 0x24, 0x90, 0xd8, 0xb4, 0x3c, 0x0f, 0xb9,
	0x80, 0xcd, 0xeb, 0x88, 0x50, 0x86, 0x13, 0xbe, 0x30, 0xcb, 0x8d, 0x07, 0xa2, 0xde, 0x52, 0x90,
	0x41, 0x45, 0xff, 0xff, 0x8b, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x91, 0x8e, 0x0c, 0x8b, 0x8c,
	0x02, 0x00, 0x00,
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
	if len(m.RedemptionSchedule) > 0 {
		for iNdEx := len(m.RedemptionSchedule) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RedemptionSchedule[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.InsuranceFunds) > 0 {
		for iNdEx := len(m.InsuranceFunds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InsuranceFunds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
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

func (m *Redemption) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Redemption) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Redemption) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Schedule.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Redeemer) > 0 {
		i -= len(m.Redeemer)
		copy(dAtA[i:], m.Redeemer)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Redeemer)))
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
	if len(m.InsuranceFunds) > 0 {
		for _, e := range m.InsuranceFunds {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.RedemptionSchedule) > 0 {
		for _, e := range m.RedemptionSchedule {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *Redemption) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Redeemer)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.Schedule.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field InsuranceFunds", wireType)
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
			m.InsuranceFunds = append(m.InsuranceFunds, InsuranceFund{})
			if err := m.InsuranceFunds[len(m.InsuranceFunds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RedemptionSchedule", wireType)
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
			m.RedemptionSchedule = append(m.RedemptionSchedule, &RedemptionSchedule{})
			if err := m.RedemptionSchedule[len(m.RedemptionSchedule)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *Redemption) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Redemption: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Redemption: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Redeemer", wireType)
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
			m.Redeemer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Schedule", wireType)
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
			if err := m.Schedule.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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