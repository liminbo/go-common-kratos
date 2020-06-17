// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: store.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Store struct {
	StoreId              int64          `protobuf:"varint,1,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	Title                string         `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Type                 int64          `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Level                int32          `protobuf:"varint,4,opt,name=level,proto3" json:"level,omitempty"`
	Cover                string         `protobuf:"bytes,5,opt,name=cover,proto3" json:"cover,omitempty"`
	ProvinceName         string         `protobuf:"bytes,6,opt,name=province_name,json=provinceName,proto3" json:"province_name,omitempty"`
	CityName             string         `protobuf:"bytes,7,opt,name=city_name,json=cityName,proto3" json:"city_name,omitempty"`
	AreaName             string         `protobuf:"bytes,8,opt,name=area_name,json=areaName,proto3" json:"area_name,omitempty"`
	ProvinceCode         int64          `protobuf:"varint,9,opt,name=province_code,json=provinceCode,proto3" json:"province_code,omitempty"`
	CityCode             int64          `protobuf:"varint,10,opt,name=city_code,json=cityCode,proto3" json:"city_code,omitempty"`
	AreaCode             int64          `protobuf:"varint,11,opt,name=area_code,json=areaCode,proto3" json:"area_code,omitempty"`
	Address              string         `protobuf:"bytes,12,opt,name=address,proto3" json:"address,omitempty"`
	Tel                  string         `protobuf:"bytes,13,opt,name=tel,proto3" json:"tel,omitempty"`
	Gcj_02               string         `protobuf:"bytes,14,opt,name=gcj_02,json=gcj02,proto3" json:"gcj_02,omitempty"`
	Location             string         `protobuf:"bytes,15,opt,name=location,proto3" json:"location,omitempty"`
	DistrictId           int64          `protobuf:"varint,16,opt,name=district_id,json=districtId,proto3" json:"district_id,omitempty"`
	Wechat               string         `protobuf:"bytes,17,opt,name=wechat,proto3" json:"wechat,omitempty"`
	TagList              *Store_TagList `protobuf:"bytes,18,opt,name=tag_list,json=tagList,proto3" json:"tag_list,omitempty"`
	ImgList              *Store_ImgList `protobuf:"bytes,19,opt,name=img_list,json=imgList,proto3" json:"img_list,omitempty"`
	Distance             string         `protobuf:"bytes,20,opt,name=distance,proto3" json:"distance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Store) Reset()         { *m = Store{} }
func (m *Store) String() string { return proto.CompactTextString(m) }
func (*Store) ProtoMessage()    {}
func (*Store) Descriptor() ([]byte, []int) {
	return fileDescriptor_98bbca36ef968dfc, []int{0}
}
func (m *Store) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Store) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Store.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Store) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Store.Merge(m, src)
}
func (m *Store) XXX_Size() int {
	return m.Size()
}
func (m *Store) XXX_DiscardUnknown() {
	xxx_messageInfo_Store.DiscardUnknown(m)
}

var xxx_messageInfo_Store proto.InternalMessageInfo

func (m *Store) GetStoreId() int64 {
	if m != nil {
		return m.StoreId
	}
	return 0
}

func (m *Store) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Store) GetType() int64 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Store) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *Store) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *Store) GetProvinceName() string {
	if m != nil {
		return m.ProvinceName
	}
	return ""
}

func (m *Store) GetCityName() string {
	if m != nil {
		return m.CityName
	}
	return ""
}

func (m *Store) GetAreaName() string {
	if m != nil {
		return m.AreaName
	}
	return ""
}

func (m *Store) GetProvinceCode() int64 {
	if m != nil {
		return m.ProvinceCode
	}
	return 0
}

func (m *Store) GetCityCode() int64 {
	if m != nil {
		return m.CityCode
	}
	return 0
}

func (m *Store) GetAreaCode() int64 {
	if m != nil {
		return m.AreaCode
	}
	return 0
}

func (m *Store) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Store) GetTel() string {
	if m != nil {
		return m.Tel
	}
	return ""
}

func (m *Store) GetGcj_02() string {
	if m != nil {
		return m.Gcj_02
	}
	return ""
}

func (m *Store) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *Store) GetDistrictId() int64 {
	if m != nil {
		return m.DistrictId
	}
	return 0
}

func (m *Store) GetWechat() string {
	if m != nil {
		return m.Wechat
	}
	return ""
}

func (m *Store) GetTagList() *Store_TagList {
	if m != nil {
		return m.TagList
	}
	return nil
}

func (m *Store) GetImgList() *Store_ImgList {
	if m != nil {
		return m.ImgList
	}
	return nil
}

func (m *Store) GetDistance() string {
	if m != nil {
		return m.Distance
	}
	return ""
}

type Store_TagList struct {
	Tag                  []string `protobuf:"bytes,1,rep,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Store_TagList) Reset()         { *m = Store_TagList{} }
func (m *Store_TagList) String() string { return proto.CompactTextString(m) }
func (*Store_TagList) ProtoMessage()    {}
func (*Store_TagList) Descriptor() ([]byte, []int) {
	return fileDescriptor_98bbca36ef968dfc, []int{0, 0}
}
func (m *Store_TagList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Store_TagList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Store_TagList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Store_TagList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Store_TagList.Merge(m, src)
}
func (m *Store_TagList) XXX_Size() int {
	return m.Size()
}
func (m *Store_TagList) XXX_DiscardUnknown() {
	xxx_messageInfo_Store_TagList.DiscardUnknown(m)
}

var xxx_messageInfo_Store_TagList proto.InternalMessageInfo

func (m *Store_TagList) GetTag() []string {
	if m != nil {
		return m.Tag
	}
	return nil
}

type Store_ImgList struct {
	Url                  []string `protobuf:"bytes,1,rep,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Store_ImgList) Reset()         { *m = Store_ImgList{} }
func (m *Store_ImgList) String() string { return proto.CompactTextString(m) }
func (*Store_ImgList) ProtoMessage()    {}
func (*Store_ImgList) Descriptor() ([]byte, []int) {
	return fileDescriptor_98bbca36ef968dfc, []int{0, 1}
}
func (m *Store_ImgList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Store_ImgList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Store_ImgList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Store_ImgList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Store_ImgList.Merge(m, src)
}
func (m *Store_ImgList) XXX_Size() int {
	return m.Size()
}
func (m *Store_ImgList) XXX_DiscardUnknown() {
	xxx_messageInfo_Store_ImgList.DiscardUnknown(m)
}

var xxx_messageInfo_Store_ImgList proto.InternalMessageInfo

func (m *Store_ImgList) GetUrl() []string {
	if m != nil {
		return m.Url
	}
	return nil
}

func init() {
	proto.RegisterType((*Store)(nil), "business.service.store.v1.Store")
	proto.RegisterType((*Store_TagList)(nil), "business.service.store.v1.Store.TagList")
	proto.RegisterType((*Store_ImgList)(nil), "business.service.store.v1.Store.ImgList")
}

func init() { proto.RegisterFile("store.proto", fileDescriptor_98bbca36ef968dfc) }

var fileDescriptor_98bbca36ef968dfc = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x6e, 0xd3, 0x40,
	0x14, 0xc6, 0x99, 0xa6, 0x8e, 0x9d, 0x97, 0x16, 0xca, 0x10, 0xd0, 0xb4, 0x95, 0x82, 0x05, 0x1b,
	0xaf, 0xac, 0x36, 0xdc, 0x80, 0xae, 0x22, 0x21, 0x16, 0x81, 0x15, 0x9b, 0x68, 0x3a, 0x7e, 0x32,
	0x53, 0xd9, 0x9e, 0x68, 0x66, 0x62, 0xd4, 0x9b, 0x70, 0x0a, 0xce, 0xc1, 0x92, 0x23, 0xa0, 0x70,
	0x11, 0x34, 0x6f, 0x6c, 0x60, 0x83, 0xd4, 0xdd, 0xfb, 0xbe, 0xdf, 0xe7, 0x6f, 0xfe, 0x19, 0xe6,
	0xce, 0x1b, 0x8b, 0xe5, 0xce, 0x1a, 0x6f, 0xf8, 0xf9, 0xed, 0xde, 0xe9, 0x0e, 0x9d, 0x2b, 0x1d,
	0xda, 0x5e, 0x2b, 0x2c, 0x23, 0xed, 0xaf, 0x5f, 0x7d, 0x4b, 0x20, 0xf9, 0x10, 0x04, 0x3f, 0x87,
	0x8c, 0xdc, 0xad, 0xae, 0x04, 0xcb, 0x59, 0x31, 0xd9, 0xa4, 0xa4, 0xd7, 0x15, 0x5f, 0x40, 0xe2,
	0xb5, 0x6f, 0x50, 0x1c, 0xe5, 0xac, 0x98, 0x6d, 0xa2, 0xe0, 0x1c, 0x8e, 0xfd, 0xfd, 0x0e, 0xc5,
	0x84, 0xc2, 0x34, 0x87, 0x64, 0x83, 0x3d, 0x36, 0xe2, 0x38, 0x67, 0x45, 0xb2, 0x89, 0x22, 0xb8,
	0xca, 0xf4, 0x68, 0x45, 0x12, 0xbf, 0x27, 0xc1, 0x5f, 0xc3, 0xe9, 0xce, 0x9a, 0x5e, 0x77, 0x0a,
	0xb7, 0x9d, 0x6c, 0x51, 0x4c, 0x89, 0x9e, 0x8c, 0xe6, 0x7b, 0xd9, 0x22, 0xbf, 0x84, 0x99, 0xd2,
	0xfe, 0x3e, 0x06, 0x52, 0x0a, 0x64, 0xc1, 0x18, 0xa1, 0xb4, 0x28, 0x23, 0xcc, 0x22, 0x0c, 0x06,
	0xc1, 0x7f, 0xeb, 0x95, 0xa9, 0x50, 0xcc, 0x68, 0x9f, 0x7f, 0xea, 0x6f, 0x4c, 0xf5, 0xb7, 0x9e,
	0x02, 0x40, 0x01, 0xaa, 0x1f, 0x21, 0xd5, 0x13, 0x9c, 0x47, 0x18, 0x0c, 0x82, 0x02, 0x52, 0x59,
	0x55, 0x16, 0x9d, 0x13, 0x27, 0xb4, 0xf2, 0x28, 0xf9, 0x19, 0x4c, 0x3c, 0x36, 0xe2, 0x94, 0xdc,
	0x30, 0xf2, 0xe7, 0x30, 0xad, 0xd5, 0xdd, 0xf6, 0x6a, 0x25, 0x1e, 0xc7, 0x0b, 0xa8, 0xd5, 0xdd,
	0xd5, 0x8a, 0x5f, 0x40, 0xd6, 0x18, 0x25, 0xbd, 0x36, 0x9d, 0x78, 0x12, 0x77, 0x3f, 0x6a, 0xfe,
	0x12, 0xe6, 0x95, 0x76, 0xde, 0x6a, 0xe5, 0xc3, 0x83, 0x9c, 0xd1, 0xea, 0x30, 0x5a, 0xeb, 0x8a,
	0xbf, 0x80, 0xe9, 0x17, 0x54, 0x9f, 0xa5, 0x17, 0x4f, 0xe9, 0xd3, 0x41, 0xf1, 0x1b, 0xc8, 0xbc,
	0xac, 0xb7, 0x8d, 0x76, 0x5e, 0xf0, 0x9c, 0x15, 0xf3, 0x55, 0x51, 0xfe, 0xf7, 0xf9, 0x4b, 0x7a,
	0xfa, 0xf2, 0xa3, 0xac, 0xdf, 0x69, 0xe7, 0x37, 0xa9, 0x8f, 0x43, 0x28, 0xd1, 0xed, 0x50, 0xf2,
	0xec, 0x81, 0x25, 0xeb, 0x76, 0x28, 0xd1, 0x71, 0x08, 0xc7, 0x0b, 0xfb, 0x95, 0x9d, 0x42, 0xb1,
	0x88, 0xc7, 0x1b, 0xf5, 0xc5, 0x25, 0xa4, 0xc3, 0xa2, 0x74, 0x5d, 0xb2, 0x16, 0x2c, 0x9f, 0xd0,
	0x75, 0xc9, 0x3a, 0xc0, 0xa1, 0x2c, 0xc0, 0xbd, 0x6d, 0x46, 0xb8, 0xb7, 0xcd, 0xdb, 0xc5, 0xf7,
	0xc3, 0x92, 0xfd, 0x38, 0x2c, 0xd9, 0xcf, 0xc3, 0x92, 0x7d, 0xfd, 0xb5, 0x7c, 0xf4, 0xe9, 0xa8,
	0xbf, 0xbe, 0x9d, 0xd2, 0x8f, 0xfe, 0xe6, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc3, 0xf8, 0x00,
	0xf5, 0xf7, 0x02, 0x00, 0x00,
}

func (m *Store) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Store) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Store) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Distance) > 0 {
		i -= len(m.Distance)
		copy(dAtA[i:], m.Distance)
		i = encodeVarintStore(dAtA, i, uint64(len(m.Distance)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xa2
	}
	if m.ImgList != nil {
		{
			size, err := m.ImgList.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintStore(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x9a
	}
	if m.TagList != nil {
		{
			size, err := m.TagList.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintStore(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x92
	}
	if len(m.Wechat) > 0 {
		i -= len(m.Wechat)
		copy(dAtA[i:], m.Wechat)
		i = encodeVarintStore(dAtA, i, uint64(len(m.Wechat)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x8a
	}
	if m.DistrictId != 0 {
		i = encodeVarintStore(dAtA, i, uint64(m.DistrictId))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if len(m.Location) > 0 {
		i -= len(m.Location)
		copy(dAtA[i:], m.Location)
		i = encodeVarintStore(dAtA, i, uint64(len(m.Location)))
		i--
		dAtA[i] = 0x7a
	}
	if len(m.Gcj_02) > 0 {
		i -= len(m.Gcj_02)
		copy(dAtA[i:], m.Gcj_02)
		i = encodeVarintStore(dAtA, i, uint64(len(m.Gcj_02)))
		i--
		dAtA[i] = 0x72
	}
	if len(m.Tel) > 0 {
		i -= len(m.Tel)
		copy(dAtA[i:], m.Tel)
		i = encodeVarintStore(dAtA, i, uint64(len(m.Tel)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintStore(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x62
	}
	if m.AreaCode != 0 {
		i = encodeVarintStore(dAtA, i, uint64(m.AreaCode))
		i--
		dAtA[i] = 0x58
	}
	if m.CityCode != 0 {
		i = encodeVarintStore(dAtA, i, uint64(m.CityCode))
		i--
		dAtA[i] = 0x50
	}
	if m.ProvinceCode != 0 {
		i = encodeVarintStore(dAtA, i, uint64(m.ProvinceCode))
		i--
		dAtA[i] = 0x48
	}
	if len(m.AreaName) > 0 {
		i -= len(m.AreaName)
		copy(dAtA[i:], m.AreaName)
		i = encodeVarintStore(dAtA, i, uint64(len(m.AreaName)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.CityName) > 0 {
		i -= len(m.CityName)
		copy(dAtA[i:], m.CityName)
		i = encodeVarintStore(dAtA, i, uint64(len(m.CityName)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.ProvinceName) > 0 {
		i -= len(m.ProvinceName)
		copy(dAtA[i:], m.ProvinceName)
		i = encodeVarintStore(dAtA, i, uint64(len(m.ProvinceName)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Cover) > 0 {
		i -= len(m.Cover)
		copy(dAtA[i:], m.Cover)
		i = encodeVarintStore(dAtA, i, uint64(len(m.Cover)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Level != 0 {
		i = encodeVarintStore(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x20
	}
	if m.Type != 0 {
		i = encodeVarintStore(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintStore(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x12
	}
	if m.StoreId != 0 {
		i = encodeVarintStore(dAtA, i, uint64(m.StoreId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Store_TagList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Store_TagList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Store_TagList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Tag) > 0 {
		for iNdEx := len(m.Tag) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Tag[iNdEx])
			copy(dAtA[i:], m.Tag[iNdEx])
			i = encodeVarintStore(dAtA, i, uint64(len(m.Tag[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Store_ImgList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Store_ImgList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Store_ImgList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Url) > 0 {
		for iNdEx := len(m.Url) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Url[iNdEx])
			copy(dAtA[i:], m.Url[iNdEx])
			i = encodeVarintStore(dAtA, i, uint64(len(m.Url[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintStore(dAtA []byte, offset int, v uint64) int {
	offset -= sovStore(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Store) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StoreId != 0 {
		n += 1 + sovStore(uint64(m.StoreId))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovStore(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovStore(uint64(m.Type))
	}
	if m.Level != 0 {
		n += 1 + sovStore(uint64(m.Level))
	}
	l = len(m.Cover)
	if l > 0 {
		n += 1 + l + sovStore(uint64(l))
	}
	l = len(m.ProvinceName)
	if l > 0 {
		n += 1 + l + sovStore(uint64(l))
	}
	l = len(m.CityName)
	if l > 0 {
		n += 1 + l + sovStore(uint64(l))
	}
	l = len(m.AreaName)
	if l > 0 {
		n += 1 + l + sovStore(uint64(l))
	}
	if m.ProvinceCode != 0 {
		n += 1 + sovStore(uint64(m.ProvinceCode))
	}
	if m.CityCode != 0 {
		n += 1 + sovStore(uint64(m.CityCode))
	}
	if m.AreaCode != 0 {
		n += 1 + sovStore(uint64(m.AreaCode))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovStore(uint64(l))
	}
	l = len(m.Tel)
	if l > 0 {
		n += 1 + l + sovStore(uint64(l))
	}
	l = len(m.Gcj_02)
	if l > 0 {
		n += 1 + l + sovStore(uint64(l))
	}
	l = len(m.Location)
	if l > 0 {
		n += 1 + l + sovStore(uint64(l))
	}
	if m.DistrictId != 0 {
		n += 2 + sovStore(uint64(m.DistrictId))
	}
	l = len(m.Wechat)
	if l > 0 {
		n += 2 + l + sovStore(uint64(l))
	}
	if m.TagList != nil {
		l = m.TagList.Size()
		n += 2 + l + sovStore(uint64(l))
	}
	if m.ImgList != nil {
		l = m.ImgList.Size()
		n += 2 + l + sovStore(uint64(l))
	}
	l = len(m.Distance)
	if l > 0 {
		n += 2 + l + sovStore(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Store_TagList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Tag) > 0 {
		for _, s := range m.Tag {
			l = len(s)
			n += 1 + l + sovStore(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Store_ImgList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Url) > 0 {
		for _, s := range m.Url {
			l = len(s)
			n += 1 + l + sovStore(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovStore(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStore(x uint64) (n int) {
	return sovStore(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Store) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStore
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
			return fmt.Errorf("proto: Store: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Store: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoreId", wireType)
			}
			m.StoreId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StoreId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Level", wireType)
			}
			m.Level = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Level |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cover", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cover = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProvinceName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProvinceName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CityName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CityName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AreaName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AreaName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProvinceCode", wireType)
			}
			m.ProvinceCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProvinceCode |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CityCode", wireType)
			}
			m.CityCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CityCode |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AreaCode", wireType)
			}
			m.AreaCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AreaCode |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gcj_02", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Gcj_02 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Location", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Location = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistrictId", wireType)
			}
			m.DistrictId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DistrictId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Wechat", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Wechat = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 18:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TagList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TagList == nil {
				m.TagList = &Store_TagList{}
			}
			if err := m.TagList.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 19:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImgList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ImgList == nil {
				m.ImgList = &Store_ImgList{}
			}
			if err := m.ImgList.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 20:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Distance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Distance = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStore(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStore
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStore
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
func (m *Store_TagList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStore
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
			return fmt.Errorf("proto: TagList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TagList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tag", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tag = append(m.Tag, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStore(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStore
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStore
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
func (m *Store_ImgList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStore
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
			return fmt.Errorf("proto: ImgList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ImgList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = append(m.Url, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStore(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStore
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStore
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
func skipStore(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStore
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
					return 0, ErrIntOverflowStore
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
					return 0, ErrIntOverflowStore
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
				return 0, ErrInvalidLengthStore
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStore
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStore
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStore        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStore          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStore = fmt.Errorf("proto: unexpected end of group")
)