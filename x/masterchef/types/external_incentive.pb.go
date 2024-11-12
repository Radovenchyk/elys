// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elys/masterchef/external_incentive.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// ExternalIncentive defines the external incentives.
type ExternalIncentive struct {
	Id             uint64                      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	RewardDenom    string                      `protobuf:"bytes,2,opt,name=reward_denom,json=rewardDenom,proto3" json:"reward_denom,omitempty"`
	PoolId         uint64                      `protobuf:"varint,3,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	FromBlock      int64                       `protobuf:"varint,4,opt,name=from_block,json=fromBlock,proto3" json:"from_block,omitempty"`
	ToBlock        int64                       `protobuf:"varint,5,opt,name=to_block,json=toBlock,proto3" json:"to_block,omitempty"`
	AmountPerBlock cosmossdk_io_math.Int       `protobuf:"bytes,6,opt,name=amount_per_block,json=amountPerBlock,proto3,customtype=cosmossdk.io/math.Int" json:"amount_per_block"`
	Apr            cosmossdk_io_math.LegacyDec `protobuf:"bytes,7,opt,name=apr,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"apr"`
}

func (m *ExternalIncentive) Reset()         { *m = ExternalIncentive{} }
func (m *ExternalIncentive) String() string { return proto.CompactTextString(m) }
func (*ExternalIncentive) ProtoMessage()    {}
func (*ExternalIncentive) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d270e9faef9dc8d, []int{0}
}
func (m *ExternalIncentive) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExternalIncentive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExternalIncentive.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExternalIncentive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExternalIncentive.Merge(m, src)
}
func (m *ExternalIncentive) XXX_Size() int {
	return m.Size()
}
func (m *ExternalIncentive) XXX_DiscardUnknown() {
	xxx_messageInfo_ExternalIncentive.DiscardUnknown(m)
}

var xxx_messageInfo_ExternalIncentive proto.InternalMessageInfo

func (m *ExternalIncentive) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ExternalIncentive) GetRewardDenom() string {
	if m != nil {
		return m.RewardDenom
	}
	return ""
}

func (m *ExternalIncentive) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *ExternalIncentive) GetFromBlock() int64 {
	if m != nil {
		return m.FromBlock
	}
	return 0
}

func (m *ExternalIncentive) GetToBlock() int64 {
	if m != nil {
		return m.ToBlock
	}
	return 0
}

func init() {
	proto.RegisterType((*ExternalIncentive)(nil), "elys.masterchef.ExternalIncentive")
}

func init() {
	proto.RegisterFile("elys/masterchef/external_incentive.proto", fileDescriptor_3d270e9faef9dc8d)
}

var fileDescriptor_3d270e9faef9dc8d = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x86, 0xdb, 0xc2, 0x85, 0xcb, 0xdc, 0x1b, 0xee, 0xb5, 0xd1, 0x58, 0x35, 0x14, 0xd4, 0x4d,
	0x37, 0xb6, 0x0b, 0xe3, 0x0b, 0x10, 0x8c, 0x69, 0xe2, 0xc2, 0x74, 0xe9, 0xa6, 0x29, 0xed, 0xa1,
	0x4c, 0xe8, 0xcc, 0x69, 0xa6, 0x83, 0xc0, 0x5b, 0xb8, 0xf5, 0x8d, 0x58, 0xb2, 0x34, 0x2e, 0x88,
	0x81, 0x17, 0x31, 0xd3, 0xd6, 0x68, 0xe2, 0x6e, 0xce, 0xff, 0xff, 0xdf, 0x3f, 0xc9, 0x39, 0xc4,
	0x81, 0x6c, 0x55, 0x78, 0x2c, 0x2a, 0x24, 0x88, 0x78, 0x0a, 0x13, 0x0f, 0x96, 0x12, 0x04, 0x8f,
	0xb2, 0x90, 0xf2, 0x18, 0xb8, 0xa4, 0x4f, 0xe0, 0xe6, 0x02, 0x25, 0x9a, 0xff, 0x54, 0xd2, 0xfd,
	0x4a, 0x9e, 0x1e, 0xa6, 0x98, 0x62, 0xe9, 0x79, 0xea, 0x55, 0xc5, 0x2e, 0x5e, 0x0c, 0x72, 0x70,
	0x5b, 0x77, 0xf8, 0x9f, 0x15, 0x66, 0x97, 0x18, 0x34, 0xb1, 0xf4, 0x81, 0xee, 0x34, 0x03, 0x83,
	0x26, 0xe6, 0x39, 0xf9, 0x2b, 0x60, 0x11, 0x89, 0x24, 0x4c, 0x80, 0x23, 0xb3, 0x8c, 0x81, 0xee,
	0x74, 0x82, 0x3f, 0x95, 0x36, 0x52, 0x92, 0x79, 0x4c, 0xda, 0x39, 0x62, 0x16, 0xd2, 0xc4, 0x6a,
	0x94, 0x5c, 0x4b, 0x8d, 0x7e, 0x62, 0xf6, 0x08, 0x99, 0x08, 0x64, 0xe1, 0x38, 0xc3, 0x78, 0x66,
	0x35, 0x07, 0xba, 0xd3, 0x08, 0x3a, 0x4a, 0x19, 0x2a, 0xc1, 0x3c, 0x21, 0xbf, 0x25, 0xd6, 0xe6,
	0xaf, 0xd2, 0x6c, 0x4b, 0xac, 0xac, 0x3b, 0xf2, 0x3f, 0x62, 0x38, 0xe7, 0x32, 0xcc, 0x41, 0xd4,
	0x91, 0x96, 0xfa, 0x79, 0xd8, 0x5b, 0x6f, 0xfb, 0xda, 0xdb, 0xb6, 0x7f, 0x14, 0x63, 0xc1, 0xb0,
	0x28, 0x92, 0x99, 0x4b, 0xd1, 0x63, 0x91, 0x9c, 0xba, 0x3e, 0x97, 0x41, 0xb7, 0xc2, 0x1e, 0x40,
	0x54, 0x45, 0x37, 0xa4, 0x11, 0xe5, 0xc2, 0x6a, 0x97, 0xec, 0x65, 0xcd, 0x9e, 0xfd, 0x64, 0xef,
	0x21, 0x8d, 0xe2, 0xd5, 0x08, 0xe2, 0x40, 0xe5, 0x87, 0xfe, 0x7a, 0x67, 0xeb, 0x9b, 0x9d, 0xad,
	0xbf, 0xef, 0x6c, 0xfd, 0x79, 0x6f, 0x6b, 0x9b, 0xbd, 0xad, 0xbd, 0xee, 0x6d, 0xed, 0xd1, 0x4b,
	0xa9, 0x9c, 0xce, 0xc7, 0x6e, 0x8c, 0xcc, 0x53, 0x7b, 0xbe, 0xe2, 0x20, 0x17, 0x28, 0x66, 0xe5,
	0xe0, 0x2d, 0xbf, 0x1f, 0x48, 0xae, 0x72, 0x28, 0xc6, 0xad, 0x72, 0xdb, 0xd7, 0x1f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x5f, 0x0d, 0x78, 0xaf, 0xc0, 0x01, 0x00, 0x00,
}

func (m *ExternalIncentive) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExternalIncentive) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExternalIncentive) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Apr.Size()
		i -= size
		if _, err := m.Apr.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExternalIncentive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.AmountPerBlock.Size()
		i -= size
		if _, err := m.AmountPerBlock.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExternalIncentive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if m.ToBlock != 0 {
		i = encodeVarintExternalIncentive(dAtA, i, uint64(m.ToBlock))
		i--
		dAtA[i] = 0x28
	}
	if m.FromBlock != 0 {
		i = encodeVarintExternalIncentive(dAtA, i, uint64(m.FromBlock))
		i--
		dAtA[i] = 0x20
	}
	if m.PoolId != 0 {
		i = encodeVarintExternalIncentive(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x18
	}
	if len(m.RewardDenom) > 0 {
		i -= len(m.RewardDenom)
		copy(dAtA[i:], m.RewardDenom)
		i = encodeVarintExternalIncentive(dAtA, i, uint64(len(m.RewardDenom)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintExternalIncentive(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintExternalIncentive(dAtA []byte, offset int, v uint64) int {
	offset -= sovExternalIncentive(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ExternalIncentive) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovExternalIncentive(uint64(m.Id))
	}
	l = len(m.RewardDenom)
	if l > 0 {
		n += 1 + l + sovExternalIncentive(uint64(l))
	}
	if m.PoolId != 0 {
		n += 1 + sovExternalIncentive(uint64(m.PoolId))
	}
	if m.FromBlock != 0 {
		n += 1 + sovExternalIncentive(uint64(m.FromBlock))
	}
	if m.ToBlock != 0 {
		n += 1 + sovExternalIncentive(uint64(m.ToBlock))
	}
	l = m.AmountPerBlock.Size()
	n += 1 + l + sovExternalIncentive(uint64(l))
	l = m.Apr.Size()
	n += 1 + l + sovExternalIncentive(uint64(l))
	return n
}

func sovExternalIncentive(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozExternalIncentive(x uint64) (n int) {
	return sovExternalIncentive(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ExternalIncentive) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExternalIncentive
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
			return fmt.Errorf("proto: ExternalIncentive: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExternalIncentive: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternalIncentive
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
				return fmt.Errorf("proto: wrong wireType = %d for field RewardDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternalIncentive
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
				return ErrInvalidLengthExternalIncentive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExternalIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RewardDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternalIncentive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromBlock", wireType)
			}
			m.FromBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternalIncentive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FromBlock |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToBlock", wireType)
			}
			m.ToBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternalIncentive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ToBlock |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountPerBlock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternalIncentive
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
				return ErrInvalidLengthExternalIncentive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExternalIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountPerBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Apr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternalIncentive
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
				return ErrInvalidLengthExternalIncentive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExternalIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Apr.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExternalIncentive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExternalIncentive
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
func skipExternalIncentive(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExternalIncentive
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
					return 0, ErrIntOverflowExternalIncentive
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
					return 0, ErrIntOverflowExternalIncentive
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
				return 0, ErrInvalidLengthExternalIncentive
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupExternalIncentive
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthExternalIncentive
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthExternalIncentive        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExternalIncentive          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupExternalIncentive = fmt.Errorf("proto: unexpected end of group")
)
