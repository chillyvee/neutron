// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: neutron/dex/genesis.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// GenesisState defines the dex module's genesis state.
type GenesisState struct {
	Params                        Params                   `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	TickLiquidityList             []*TickLiquidity         `protobuf:"bytes,2,rep,name=tick_liquidity_list,json=tickLiquidityList,proto3" json:"tick_liquidity_list,omitempty"`
	InactiveLimitOrderTrancheList []*LimitOrderTranche     `protobuf:"bytes,3,rep,name=inactive_limit_order_tranche_list,json=inactiveLimitOrderTrancheList,proto3" json:"inactive_limit_order_tranche_list,omitempty"`
	LimitOrderTrancheUserList     []*LimitOrderTrancheUser `protobuf:"bytes,4,rep,name=limit_order_tranche_user_list,json=limitOrderTrancheUserList,proto3" json:"limit_order_tranche_user_list,omitempty"`
	PoolMetadataList              []PoolMetadata           `protobuf:"bytes,5,rep,name=pool_metadata_list,json=poolMetadataList,proto3" json:"pool_metadata_list"`
	PoolCount                     uint64                   `protobuf:"varint,6,opt,name=pool_count,json=poolCount,proto3" json:"pool_count,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c051a8a0d58cd8b, []int{0}
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

func (m *GenesisState) GetTickLiquidityList() []*TickLiquidity {
	if m != nil {
		return m.TickLiquidityList
	}
	return nil
}

func (m *GenesisState) GetInactiveLimitOrderTrancheList() []*LimitOrderTranche {
	if m != nil {
		return m.InactiveLimitOrderTrancheList
	}
	return nil
}

func (m *GenesisState) GetLimitOrderTrancheUserList() []*LimitOrderTrancheUser {
	if m != nil {
		return m.LimitOrderTrancheUserList
	}
	return nil
}

func (m *GenesisState) GetPoolMetadataList() []PoolMetadata {
	if m != nil {
		return m.PoolMetadataList
	}
	return nil
}

func (m *GenesisState) GetPoolCount() uint64 {
	if m != nil {
		return m.PoolCount
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "neutron.dex.GenesisState")
}

func init() { proto.RegisterFile("neutron/dex/genesis.proto", fileDescriptor_0c051a8a0d58cd8b) }

var fileDescriptor_0c051a8a0d58cd8b = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x6e, 0xda, 0x30,
	0x18, 0xc7, 0x93, 0xc1, 0x90, 0x66, 0x76, 0xd8, 0xc2, 0x0e, 0x21, 0x12, 0x21, 0x43, 0x9a, 0x84,
	0x90, 0x96, 0x68, 0xec, 0x0d, 0xd8, 0xa4, 0x5d, 0x40, 0x43, 0x8c, 0x5d, 0x76, 0x89, 0x4c, 0x62,
	0x05, 0x8f, 0x24, 0x4e, 0x9d, 0x2f, 0x15, 0xbc, 0x45, 0x1f, 0x8b, 0x23, 0xc7, 0x5e, 0x5a, 0x55,
	0xf0, 0x22, 0x55, 0x6c, 0x23, 0x25, 0x2a, 0x6d, 0x6f, 0xd6, 0xf7, 0xfd, 0xfc, 0xff, 0xd9, 0x9f,
	0x8d, 0xba, 0x29, 0x29, 0x80, 0xb3, 0xd4, 0x0b, 0xc9, 0xd6, 0x8b, 0x48, 0x4a, 0x72, 0x9a, 0xbb,
	0x19, 0x67, 0xc0, 0x8c, 0xb6, 0x6a, 0xb9, 0x21, 0xd9, 0x5a, 0x9f, 0x22, 0x16, 0x31, 0x51, 0xf7,
	0xca, 0x95, 0x44, 0x2c, 0xb3, 0xba, 0x3b, 0xc3, 0x1c, 0x27, 0x6a, 0xb3, 0x35, 0xaa, 0x76, 0x62,
	0x9a, 0x50, 0xf0, 0x19, 0x0f, 0x09, 0xf7, 0x81, 0xe3, 0x34, 0x58, 0x13, 0xbf, 0xc8, 0x09, 0x57,
	0xec, 0x97, 0x57, 0x58, 0x85, 0x39, 0x55, 0x0c, 0x68, 0xb0, 0xf1, 0x63, 0x7a, 0x55, 0xd0, 0x90,
	0xc2, 0x4e, 0x11, 0xfd, 0xda, 0x71, 0x18, 0x8b, 0xfd, 0x84, 0x00, 0x0e, 0x31, 0x60, 0x09, 0x0c,
	0xee, 0x1a, 0xe8, 0xfd, 0x2f, 0x79, 0xc9, 0x3f, 0x80, 0x81, 0x18, 0xdf, 0x50, 0x4b, 0x1e, 0xdb,
	0xd4, 0x1d, 0x7d, 0xd8, 0x1e, 0x77, 0xdc, 0xca, 0xa5, 0xdd, 0xb9, 0x68, 0x4d, 0x9a, 0xfb, 0xfb,
	0xbe, 0xb6, 0x50, 0xa0, 0x31, 0x47, 0x9d, 0xba, 0xdc, 0x8f, 0x69, 0x0e, 0xe6, 0x1b, 0xa7, 0x31,
	0x6c, 0x8f, 0xad, 0xda, 0xfe, 0x25, 0x0d, 0x36, 0xd3, 0x33, 0x26, 0x62, 0xf4, 0xc5, 0x47, 0xa8,
	0x16, 0xa7, 0x34, 0x07, 0x23, 0x45, 0x9f, 0x69, 0x8a, 0x03, 0xa0, 0xd7, 0xc4, 0xbf, 0x34, 0x2a,
	0x91, 0xdf, 0x10, 0xf9, 0x76, 0x2d, 0x7f, 0x5a, 0xc2, 0xbf, 0x4b, 0x76, 0x29, 0x51, 0xe5, 0xe8,
	0x9d, 0xe3, 0x9e, 0x00, 0xc2, 0xf7, 0x1f, 0xf5, 0x9e, 0x7b, 0x11, 0xe9, 0x6a, 0x0a, 0xd7, 0xe0,
	0x65, 0xd7, 0xdf, 0x9c, 0x70, 0xe5, 0xeb, 0xc6, 0x97, 0x9a, 0xc2, 0x35, 0x43, 0x46, 0xed, 0x21,
	0xa4, 0xe0, 0xad, 0x10, 0x74, 0xeb, 0xc3, 0x66, 0x2c, 0x9e, 0x29, 0x4a, 0x8d, 0xfc, 0x43, 0x56,
	0xa9, 0x89, 0xb8, 0x1e, 0x42, 0x22, 0x2e, 0x60, 0x45, 0x0a, 0x66, 0xcb, 0xd1, 0x87, 0xcd, 0xc5,
	0xbb, 0xb2, 0xf2, 0xa3, 0x2c, 0x4c, 0x7e, 0xee, 0x8f, 0xb6, 0x7e, 0x38, 0xda, 0xfa, 0xc3, 0xd1,
	0xd6, 0x6f, 0x4e, 0xb6, 0x76, 0x38, 0xd9, 0xda, 0xed, 0xc9, 0xd6, 0xfe, 0x8d, 0x22, 0x0a, 0xeb,
	0x62, 0xe5, 0x06, 0x2c, 0xf1, 0x94, 0xf5, 0x2b, 0xe3, 0xd1, 0x79, 0xed, 0x6d, 0xe5, 0xaf, 0xda,
	0x65, 0x24, 0x5f, 0xb5, 0xc4, 0x67, 0xf9, 0xfe, 0x18, 0x00, 0x00, 0xff, 0xff, 0x62, 0xa7, 0x84,
	0x86, 0x1c, 0x03, 0x00, 0x00,
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
	if m.PoolCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.PoolCount))
		i--
		dAtA[i] = 0x30
	}
	if len(m.PoolMetadataList) > 0 {
		for iNdEx := len(m.PoolMetadataList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PoolMetadataList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.LimitOrderTrancheUserList) > 0 {
		for iNdEx := len(m.LimitOrderTrancheUserList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LimitOrderTrancheUserList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.InactiveLimitOrderTrancheList) > 0 {
		for iNdEx := len(m.InactiveLimitOrderTrancheList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InactiveLimitOrderTrancheList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.TickLiquidityList) > 0 {
		for iNdEx := len(m.TickLiquidityList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TickLiquidityList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.TickLiquidityList) > 0 {
		for _, e := range m.TickLiquidityList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.InactiveLimitOrderTrancheList) > 0 {
		for _, e := range m.InactiveLimitOrderTrancheList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LimitOrderTrancheUserList) > 0 {
		for _, e := range m.LimitOrderTrancheUserList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PoolMetadataList) > 0 {
		for _, e := range m.PoolMetadataList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.PoolCount != 0 {
		n += 1 + sovGenesis(uint64(m.PoolCount))
	}
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
				return fmt.Errorf("proto: wrong wireType = %d for field TickLiquidityList", wireType)
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
			m.TickLiquidityList = append(m.TickLiquidityList, &TickLiquidity{})
			if err := m.TickLiquidityList[len(m.TickLiquidityList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InactiveLimitOrderTrancheList", wireType)
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
			m.InactiveLimitOrderTrancheList = append(m.InactiveLimitOrderTrancheList, &LimitOrderTranche{})
			if err := m.InactiveLimitOrderTrancheList[len(m.InactiveLimitOrderTrancheList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LimitOrderTrancheUserList", wireType)
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
			m.LimitOrderTrancheUserList = append(m.LimitOrderTrancheUserList, &LimitOrderTrancheUser{})
			if err := m.LimitOrderTrancheUserList[len(m.LimitOrderTrancheUserList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolMetadataList", wireType)
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
			m.PoolMetadataList = append(m.PoolMetadataList, PoolMetadata{})
			if err := m.PoolMetadataList[len(m.PoolMetadataList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolCount", wireType)
			}
			m.PoolCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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