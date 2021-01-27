// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: server/worker/pipeline/transform/transform.proto

package transform

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	pfs "github.com/pachyderm/pachyderm/src/pfs"
	datum "github.com/pachyderm/pachyderm/src/server/worker/datum"
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

type DatumSet struct {
	// Inputs
	JobID        string      `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	FileSet      string      `protobuf:"bytes,2,opt,name=file_set,json=fileSet,proto3" json:"file_set,omitempty"`
	OutputCommit *pfs.Commit `protobuf:"bytes,3,opt,name=output_commit,json=outputCommit,proto3" json:"output_commit,omitempty"`
	MetaCommit   *pfs.Commit `protobuf:"bytes,4,opt,name=meta_commit,json=metaCommit,proto3" json:"meta_commit,omitempty"`
	// Outputs
	Stats                *datum.Stats `protobuf:"bytes,5,opt,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DatumSet) Reset()         { *m = DatumSet{} }
func (m *DatumSet) String() string { return proto.CompactTextString(m) }
func (*DatumSet) ProtoMessage()    {}
func (*DatumSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_21583a759eb7fa97, []int{0}
}
func (m *DatumSet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DatumSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DatumSet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DatumSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatumSet.Merge(m, src)
}
func (m *DatumSet) XXX_Size() int {
	return m.Size()
}
func (m *DatumSet) XXX_DiscardUnknown() {
	xxx_messageInfo_DatumSet.DiscardUnknown(m)
}

var xxx_messageInfo_DatumSet proto.InternalMessageInfo

func (m *DatumSet) GetJobID() string {
	if m != nil {
		return m.JobID
	}
	return ""
}

func (m *DatumSet) GetFileSet() string {
	if m != nil {
		return m.FileSet
	}
	return ""
}

func (m *DatumSet) GetOutputCommit() *pfs.Commit {
	if m != nil {
		return m.OutputCommit
	}
	return nil
}

func (m *DatumSet) GetMetaCommit() *pfs.Commit {
	if m != nil {
		return m.MetaCommit
	}
	return nil
}

func (m *DatumSet) GetStats() *datum.Stats {
	if m != nil {
		return m.Stats
	}
	return nil
}

func init() {
	proto.RegisterType((*DatumSet)(nil), "pachyderm.worker.pipeline.transform.DatumSet")
}

func init() {
	proto.RegisterFile("server/worker/pipeline/transform/transform.proto", fileDescriptor_21583a759eb7fa97)
}

var fileDescriptor_21583a759eb7fa97 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcf, 0x4e, 0x02, 0x31,
	0x10, 0xc6, 0x53, 0x75, 0x11, 0x0a, 0x5c, 0x36, 0x1e, 0x56, 0x0e, 0x40, 0xf0, 0xc2, 0xc1, 0xb4,
	0x44, 0xdf, 0x00, 0xb8, 0xe0, 0xcd, 0xe5, 0xe6, 0x85, 0xec, 0x9f, 0xee, 0xb2, 0x48, 0x99, 0xa6,
	0x9d, 0xd5, 0xf8, 0x86, 0x26, 0x5e, 0x7c, 0x02, 0x63, 0xf6, 0x49, 0x4c, 0xb7, 0x80, 0x1a, 0x0f,
	0x5e, 0x9a, 0xef, 0x9b, 0xf9, 0x7d, 0x69, 0x3b, 0x43, 0x27, 0x46, 0xe8, 0x27, 0xa1, 0xf9, 0x33,
	0xe8, 0x47, 0xa1, 0xb9, 0x2a, 0x94, 0xd8, 0x16, 0x3b, 0xc1, 0x51, 0x47, 0x3b, 0x93, 0x81, 0x96,
	0xdf, 0x8a, 0x29, 0x0d, 0x08, 0xfe, 0x95, 0x8a, 0x92, 0xf5, 0x4b, 0x2a, 0xb4, 0x64, 0x2e, 0xc4,
	0x0e, 0x21, 0x76, 0x44, 0x7b, 0x17, 0x39, 0xe4, 0x50, 0xf3, 0xdc, 0x2a, 0x17, 0xed, 0x75, 0x55,
	0x66, 0xb8, 0xca, 0xcc, 0xde, 0x0e, 0x7e, 0xdf, 0x9d, 0x46, 0x58, 0x4a, 0x77, 0x3a, 0x60, 0xf4,
	0x46, 0x68, 0x73, 0x6e, 0xfd, 0x52, 0xa0, 0x3f, 0xa4, 0x8d, 0x0d, 0xc4, 0xab, 0x22, 0x0d, 0xc8,
	0x90, 0x8c, 0x5b, 0xd3, 0x56, 0xf5, 0x31, 0xf0, 0xee, 0x20, 0x5e, 0xcc, 0x43, 0x6f, 0x03, 0xf1,
	0x22, 0xf5, 0x2f, 0x69, 0x33, 0x2b, 0xb6, 0x62, 0x65, 0x04, 0x06, 0x27, 0x96, 0x09, 0xcf, 0xad,
	0xb7, 0xe1, 0x09, 0xed, 0x42, 0x89, 0xaa, 0xc4, 0x55, 0x02, 0x52, 0x16, 0x18, 0x9c, 0x0e, 0xc9,
	0xb8, 0x7d, 0xd3, 0x66, 0xf6, 0x35, 0xb3, 0xba, 0x14, 0x76, 0x1c, 0xe1, 0x9c, 0x7f, 0x4d, 0xdb,
	0x52, 0x60, 0x74, 0xe0, 0xcf, 0xfe, 0xf2, 0xd4, 0xf6, 0xf7, 0xf4, 0x88, 0x7a, 0x06, 0x23, 0x34,
	0x81, 0x57, 0x73, 0x1d, 0xe6, 0xbe, 0xb1, 0xb4, 0xb5, 0xd0, 0xb5, 0xa6, 0xf7, 0xaf, 0x55, 0x9f,
	0xbc, 0x57, 0x7d, 0xf2, 0x59, 0xf5, 0xc9, 0xc3, 0x2c, 0x2f, 0x70, 0x5d, 0xc6, 0x2c, 0x01, 0xc9,
	0x8f, 0x13, 0xfd, 0xa1, 0x8c, 0x4e, 0xf8, 0x7f, 0xbb, 0x89, 0x1b, 0xf5, 0x9c, 0x6e, 0xbf, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xc1, 0x70, 0x23, 0x60, 0xc6, 0x01, 0x00, 0x00,
}

func (m *DatumSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DatumSet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DatumSet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Stats != nil {
		{
			size, err := m.Stats.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTransform(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.MetaCommit != nil {
		{
			size, err := m.MetaCommit.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTransform(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.OutputCommit != nil {
		{
			size, err := m.OutputCommit.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTransform(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.FileSet) > 0 {
		i -= len(m.FileSet)
		copy(dAtA[i:], m.FileSet)
		i = encodeVarintTransform(dAtA, i, uint64(len(m.FileSet)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.JobID) > 0 {
		i -= len(m.JobID)
		copy(dAtA[i:], m.JobID)
		i = encodeVarintTransform(dAtA, i, uint64(len(m.JobID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTransform(dAtA []byte, offset int, v uint64) int {
	offset -= sovTransform(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DatumSet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.JobID)
	if l > 0 {
		n += 1 + l + sovTransform(uint64(l))
	}
	l = len(m.FileSet)
	if l > 0 {
		n += 1 + l + sovTransform(uint64(l))
	}
	if m.OutputCommit != nil {
		l = m.OutputCommit.Size()
		n += 1 + l + sovTransform(uint64(l))
	}
	if m.MetaCommit != nil {
		l = m.MetaCommit.Size()
		n += 1 + l + sovTransform(uint64(l))
	}
	if m.Stats != nil {
		l = m.Stats.Size()
		n += 1 + l + sovTransform(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTransform(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTransform(x uint64) (n int) {
	return sovTransform(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DatumSet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransform
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
			return fmt.Errorf("proto: DatumSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DatumSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JobID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransform
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
				return ErrInvalidLengthTransform
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransform
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JobID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileSet", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransform
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
				return ErrInvalidLengthTransform
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransform
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FileSet = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutputCommit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransform
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
				return ErrInvalidLengthTransform
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTransform
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OutputCommit == nil {
				m.OutputCommit = &pfs.Commit{}
			}
			if err := m.OutputCommit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetaCommit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransform
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
				return ErrInvalidLengthTransform
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTransform
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MetaCommit == nil {
				m.MetaCommit = &pfs.Commit{}
			}
			if err := m.MetaCommit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransform
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
				return ErrInvalidLengthTransform
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTransform
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Stats == nil {
				m.Stats = &datum.Stats{}
			}
			if err := m.Stats.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTransform(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTransform
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTransform(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTransform
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
					return 0, ErrIntOverflowTransform
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
					return 0, ErrIntOverflowTransform
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
				return 0, ErrInvalidLengthTransform
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTransform
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTransform
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTransform        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTransform          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTransform = fmt.Errorf("proto: unexpected end of group")
)
