// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: interchainqueries/genesis.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

type RegisteredQuery struct {
	// The unique id of the registered query.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// The address that registered the query.
	Owner string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	// The query type identifier: `kv` or `tx` now
	QueryType string `protobuf:"bytes,3,opt,name=query_type,json=queryType,proto3" json:"query_type,omitempty"`
	// The KV-storage keys for which we want to get values from remote chain
	Keys []*KVKey `protobuf:"bytes,4,rep,name=keys,proto3" json:"keys,omitempty"`
	// The filter for transaction search ICQ
	TransactionsFilter string `protobuf:"bytes,5,opt,name=transactions_filter,json=transactionsFilter,proto3" json:"transactions_filter,omitempty"`
	// The IBC connection ID for getting ConsensusState to verify proofs
	ConnectionId string `protobuf:"bytes,6,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	// Parameter that defines how often the query must be updated.
	UpdatePeriod uint64 `protobuf:"varint,7,opt,name=update_period,json=updatePeriod,proto3" json:"update_period,omitempty"`
	// The local chain last block height when the query result was updated.
	LastSubmittedResultLocalHeight uint64 `protobuf:"varint,8,opt,name=last_submitted_result_local_height,json=lastSubmittedResultLocalHeight,proto3" json:"last_submitted_result_local_height,omitempty"`
	// The remote chain last block height when the query result was updated.
	LastSubmittedResultRemoteHeight uint64 `protobuf:"varint,10,opt,name=last_submitted_result_remote_height,json=lastSubmittedResultRemoteHeight,proto3" json:"last_submitted_result_remote_height,omitempty"`
	// Amount of coins deposited for the query.
	Deposit github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,11,rep,name=deposit,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"deposit"`
}

func (m *RegisteredQuery) Reset()         { *m = RegisteredQuery{} }
func (m *RegisteredQuery) String() string { return proto.CompactTextString(m) }
func (*RegisteredQuery) ProtoMessage()    {}
func (*RegisteredQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_68e6c14f58b92f58, []int{0}
}
func (m *RegisteredQuery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisteredQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisteredQuery.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisteredQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisteredQuery.Merge(m, src)
}
func (m *RegisteredQuery) XXX_Size() int {
	return m.Size()
}
func (m *RegisteredQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisteredQuery.DiscardUnknown(m)
}

var xxx_messageInfo_RegisteredQuery proto.InternalMessageInfo

func (m *RegisteredQuery) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RegisteredQuery) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *RegisteredQuery) GetQueryType() string {
	if m != nil {
		return m.QueryType
	}
	return ""
}

func (m *RegisteredQuery) GetKeys() []*KVKey {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *RegisteredQuery) GetTransactionsFilter() string {
	if m != nil {
		return m.TransactionsFilter
	}
	return ""
}

func (m *RegisteredQuery) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *RegisteredQuery) GetUpdatePeriod() uint64 {
	if m != nil {
		return m.UpdatePeriod
	}
	return 0
}

func (m *RegisteredQuery) GetLastSubmittedResultLocalHeight() uint64 {
	if m != nil {
		return m.LastSubmittedResultLocalHeight
	}
	return 0
}

func (m *RegisteredQuery) GetLastSubmittedResultRemoteHeight() uint64 {
	if m != nil {
		return m.LastSubmittedResultRemoteHeight
	}
	return 0
}

func (m *RegisteredQuery) GetDeposit() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Deposit
	}
	return nil
}

type KVKey struct {
	// Path (storage prefix) to the storage where you want to read value by key (usually name of cosmos-sdk module: 'staking', 'bank', etc.)
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Key you want to read from the storage
	Key []byte `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *KVKey) Reset()         { *m = KVKey{} }
func (m *KVKey) String() string { return proto.CompactTextString(m) }
func (*KVKey) ProtoMessage()    {}
func (*KVKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_68e6c14f58b92f58, []int{1}
}
func (m *KVKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KVKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KVKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KVKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KVKey.Merge(m, src)
}
func (m *KVKey) XXX_Size() int {
	return m.Size()
}
func (m *KVKey) XXX_DiscardUnknown() {
	xxx_messageInfo_KVKey.DiscardUnknown(m)
}

var xxx_messageInfo_KVKey proto.InternalMessageInfo

func (m *KVKey) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *KVKey) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

// GenesisState defines the interchainadapter module's genesis state.
type GenesisState struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_68e6c14f58b92f58, []int{2}
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

func init() {
	proto.RegisterType((*RegisteredQuery)(nil), "neutron.interchainadapter.interchainqueries.RegisteredQuery")
	proto.RegisterType((*KVKey)(nil), "neutron.interchainadapter.interchainqueries.KVKey")
	proto.RegisterType((*GenesisState)(nil), "neutron.interchainadapter.interchainqueries.GenesisState")
}

func init() { proto.RegisterFile("interchainqueries/genesis.proto", fileDescriptor_68e6c14f58b92f58) }

var fileDescriptor_68e6c14f58b92f58 = []byte{
	// 555 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcd, 0x6e, 0xd3, 0x4e,
	0x10, 0x8f, 0x1b, 0x27, 0xfd, 0x67, 0x93, 0x3f, 0x1f, 0xdb, 0x1e, 0x4c, 0x25, 0x9c, 0x28, 0xbd,
	0x44, 0x42, 0xb1, 0x69, 0x7a, 0xe1, 0x1c, 0x44, 0xf9, 0x68, 0x0f, 0xad, 0x8b, 0x38, 0x70, 0xb1,
	0x36, 0xf6, 0xe0, 0xac, 0x92, 0xec, 0x9a, 0xdd, 0x35, 0xe0, 0xb7, 0xe0, 0x39, 0x78, 0x92, 0x8a,
	0x53, 0x8f, 0x9c, 0x00, 0x25, 0x2f, 0x82, 0x3c, 0xb6, 0xd5, 0x4a, 0xe9, 0xa5, 0xa7, 0x4c, 0xe6,
	0xf7, 0x31, 0x1e, 0xed, 0x6f, 0x48, 0x9f, 0x0b, 0x03, 0x2a, 0x9a, 0x33, 0x2e, 0x3e, 0x67, 0xa0,
	0x38, 0x68, 0x3f, 0x01, 0x01, 0x9a, 0x6b, 0x2f, 0x55, 0xd2, 0x48, 0xfa, 0x4c, 0x40, 0x66, 0x94,
	0x14, 0xde, 0x0d, 0x91, 0xc5, 0x2c, 0x35, 0xa0, 0xbc, 0x2d, 0xe9, 0xc1, 0x7e, 0x22, 0x13, 0x89,
	0x3a, 0xbf, 0xa8, 0x4a, 0x8b, 0x03, 0x77, 0x7b, 0x46, 0xca, 0x14, 0x5b, 0xe9, 0x1a, 0x8f, 0xa4,
	0x5e, 0x49, 0xed, 0xcf, 0x98, 0x06, 0xff, 0xcb, 0xd1, 0x0c, 0x0c, 0x3b, 0xf2, 0x23, 0xc9, 0x45,
	0x89, 0x0f, 0x7f, 0xda, 0xe4, 0x61, 0x00, 0x09, 0xd7, 0x06, 0x14, 0xc4, 0x17, 0x19, 0xa8, 0x9c,
	0x3e, 0x20, 0x3b, 0x3c, 0x76, 0xac, 0x81, 0x35, 0xb2, 0x83, 0x1d, 0x1e, 0xd3, 0x7d, 0xd2, 0x92,
	0x5f, 0x05, 0x28, 0x67, 0x67, 0x60, 0x8d, 0x3a, 0x41, 0xf9, 0x87, 0x3e, 0x25, 0xa4, 0x98, 0x98,
	0x87, 0x26, 0x4f, 0xc1, 0x69, 0x22, 0xd4, 0xc1, 0xce, 0xfb, 0x3c, 0x05, 0x7a, 0x42, 0xec, 0x05,
	0xe4, 0xda, 0xb1, 0x07, 0xcd, 0x51, 0x77, 0x32, 0xf1, 0xee, 0xb1, 0xaa, 0x77, 0xfa, 0xe1, 0x14,
	0xf2, 0x00, 0xf5, 0xd4, 0x27, 0x7b, 0x46, 0x31, 0xa1, 0x59, 0x64, 0xb8, 0x14, 0x3a, 0xfc, 0xc4,
	0x97, 0x06, 0x94, 0xd3, 0xc2, 0x79, 0xf4, 0x36, 0x74, 0x82, 0x08, 0x3d, 0x24, 0xff, 0x47, 0x52,
	0x08, 0xc0, 0x66, 0xc8, 0x63, 0xa7, 0x8d, 0xd4, 0xde, 0x4d, 0xf3, 0x6d, 0x5c, 0x90, 0xb2, 0x34,
	0x66, 0x06, 0xc2, 0x14, 0x14, 0x97, 0xb1, 0xb3, 0x8b, 0xdb, 0xf6, 0xca, 0xe6, 0x39, 0xf6, 0xa8,
	0x47, 0xf6, 0x96, 0x4c, 0x9b, 0x10, 0x56, 0xdc, 0x18, 0x88, 0xc3, 0x39, 0xf0, 0x64, 0x6e, 0x9c,
	0xff, 0x90, 0xfa, 0xb8, 0x80, 0x5e, 0x95, 0xc8, 0x1b, 0x04, 0xe8, 0x3b, 0x32, 0x44, 0xbe, 0xce,
	0x66, 0x95, 0x42, 0x81, 0xce, 0x96, 0x26, 0x5c, 0xca, 0x88, 0x2d, 0x6b, 0x79, 0x07, 0xe5, 0x6e,
	0xc1, 0xbc, 0xac, 0x89, 0x01, 0xf2, 0xce, 0x0a, 0x5a, 0xe5, 0x75, 0x46, 0x0e, 0xef, 0xf6, 0x52,
	0xb0, 0x92, 0x06, 0x6a, 0x33, 0x82, 0x66, 0xfd, 0x3b, 0xcc, 0x02, 0xe4, 0x55, 0x6e, 0x40, 0x76,
	0x63, 0x48, 0xa5, 0xe6, 0xc6, 0xe9, 0xe2, 0x7b, 0x3c, 0xf1, 0xca, 0x5c, 0x78, 0x45, 0x2e, 0xbc,
	0x2a, 0x17, 0xde, 0x4b, 0xc9, 0xc5, 0xf4, 0xf9, 0xd5, 0xef, 0x7e, 0xe3, 0xc7, 0x9f, 0xfe, 0x28,
	0xe1, 0x66, 0x9e, 0xcd, 0xbc, 0x48, 0xae, 0xfc, 0x2a, 0x44, 0xe5, 0xcf, 0x58, 0xc7, 0x0b, 0xbf,
	0x78, 0x74, 0x8d, 0x02, 0x1d, 0xd4, 0xde, 0xc3, 0x31, 0x69, 0xe1, 0xd3, 0x51, 0x4a, 0xec, 0x94,
	0x99, 0x39, 0x66, 0xa8, 0x13, 0x60, 0x4d, 0x1f, 0x91, 0xe6, 0x02, 0x72, 0xcc, 0x50, 0x2f, 0x28,
	0xca, 0x21, 0x23, 0xbd, 0xd7, 0xe5, 0x3d, 0x5c, 0x1a, 0x66, 0x80, 0x5e, 0x90, 0x76, 0x99, 0x5d,
	0xd4, 0x75, 0x27, 0xc7, 0xf7, 0x0a, 0xcd, 0x39, 0x4a, 0xa7, 0x76, 0xf1, 0xf9, 0x41, 0x65, 0x34,
	0x0d, 0xae, 0xd6, 0xae, 0x75, 0xbd, 0x76, 0xad, 0xbf, 0x6b, 0xd7, 0xfa, 0xbe, 0x71, 0x1b, 0xd7,
	0x1b, 0xb7, 0xf1, 0x6b, 0xe3, 0x36, 0x3e, 0xbe, 0xb8, 0xb5, 0x5e, 0x35, 0x66, 0x2c, 0x55, 0x52,
	0xd7, 0xfe, 0x37, 0x7f, 0xfb, 0xb2, 0x70, 0xe9, 0x59, 0x1b, 0x2f, 0xe7, 0xf8, 0x5f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x1b, 0x3e, 0x68, 0xc9, 0xdf, 0x03, 0x00, 0x00,
}

func (m *RegisteredQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisteredQuery) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisteredQuery) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Deposit) > 0 {
		for iNdEx := len(m.Deposit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Deposit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x5a
		}
	}
	if m.LastSubmittedResultRemoteHeight != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastSubmittedResultRemoteHeight))
		i--
		dAtA[i] = 0x48
	}
	if m.LastSubmittedResultLocalHeight != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastSubmittedResultLocalHeight))
		i--
		dAtA[i] = 0x40
	}
	if m.UpdatePeriod != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.UpdatePeriod))
		i--
		dAtA[i] = 0x38
	}
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ConnectionId)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.TransactionsFilter) > 0 {
		i -= len(m.TransactionsFilter)
		copy(dAtA[i:], m.TransactionsFilter)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.TransactionsFilter)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Keys) > 0 {
		for iNdEx := len(m.Keys) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Keys[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.QueryType) > 0 {
		i -= len(m.QueryType)
		copy(dAtA[i:], m.QueryType)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.QueryType)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *KVKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KVKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KVKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
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
func (m *RegisteredQuery) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovGenesis(uint64(m.Id))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.QueryType)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.Keys) > 0 {
		for _, e := range m.Keys {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = len(m.TransactionsFilter)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.UpdatePeriod != 0 {
		n += 1 + sovGenesis(uint64(m.UpdatePeriod))
	}
	if m.LastSubmittedResultLocalHeight != 0 {
		n += 1 + sovGenesis(uint64(m.LastSubmittedResultLocalHeight))
	}
	if m.LastSubmittedResultRemoteHeight != 0 {
		n += 1 + sovGenesis(uint64(m.LastSubmittedResultRemoteHeight))
	}
	if len(m.Deposit) > 0 {
		for _, e := range m.Deposit {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *KVKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RegisteredQuery) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: RegisteredQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisteredQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
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
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryType", wireType)
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
			m.QueryType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Keys", wireType)
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
			m.Keys = append(m.Keys, &KVKey{})
			if err := m.Keys[len(m.Keys)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransactionsFilter", wireType)
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
			m.TransactionsFilter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
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
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatePeriod", wireType)
			}
			m.UpdatePeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdatePeriod |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastEmittedHeight", wireType)
			}
			m.LastEmittedHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastEmittedHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastSubmittedResultLocalHeight", wireType)
			}
			m.LastSubmittedResultLocalHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastSubmittedResultLocalHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastSubmittedResultRemoteHeight", wireType)
			}
			m.LastSubmittedResultRemoteHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastSubmittedResultRemoteHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposit", wireType)
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
			m.Deposit = append(m.Deposit, types.Coin{})
			if err := m.Deposit[len(m.Deposit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *KVKey) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: KVKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KVKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
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
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
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
