// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/appengine/v1/location.proto

package appengine

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "google.golang.org/genproto/googleapis/type/latlng"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Metadata for the given [google.cloud.location.Location][google.cloud.location.Location].
type LocationMetadata struct {
	// App Engine Standard Environment is available in the given location.
	//
	// @OutputOnly
	StandardEnvironmentAvailable bool `protobuf:"varint,2,opt,name=standard_environment_available,json=standardEnvironmentAvailable" json:"standard_environment_available,omitempty"`
	// App Engine Flexible Environment is available in the given location.
	//
	// @OutputOnly
	FlexibleEnvironmentAvailable bool `protobuf:"varint,4,opt,name=flexible_environment_available,json=flexibleEnvironmentAvailable" json:"flexible_environment_available,omitempty"`
}

func (m *LocationMetadata) Reset()                    { *m = LocationMetadata{} }
func (m *LocationMetadata) String() string            { return proto.CompactTextString(m) }
func (*LocationMetadata) ProtoMessage()               {}
func (*LocationMetadata) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *LocationMetadata) GetStandardEnvironmentAvailable() bool {
	if m != nil {
		return m.StandardEnvironmentAvailable
	}
	return false
}

func (m *LocationMetadata) GetFlexibleEnvironmentAvailable() bool {
	if m != nil {
		return m.FlexibleEnvironmentAvailable
	}
	return false
}

func init() {
	proto.RegisterType((*LocationMetadata)(nil), "google.appengine.v1.LocationMetadata")
}

func init() { proto.RegisterFile("google/appengine/v1/location.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x89, 0x88, 0x48, 0x40, 0x90, 0x7a, 0xb0, 0x94, 0x22, 0xd2, 0x93, 0xa7, 0x5d, 0x8a,
	0x47, 0xbd, 0x58, 0xf4, 0xa6, 0x50, 0x3c, 0x7a, 0x29, 0x93, 0x66, 0x1c, 0x17, 0xa6, 0x33, 0x4b,
	0x32, 0x04, 0xfd, 0x33, 0xfe, 0x56, 0x69, 0x36, 0x1b, 0x11, 0xea, 0x71, 0x79, 0xdf, 0xfb, 0xd8,
	0x37, 0xe5, 0x82, 0x54, 0x89, 0xd1, 0x43, 0x8c, 0x28, 0x14, 0x04, 0x7d, 0xb7, 0xf4, 0xac, 0x5b,
	0xb0, 0xa0, 0xe2, 0x62, 0xa3, 0xa6, 0x93, 0x8b, 0xc4, 0xb8, 0x91, 0x71, 0xdd, 0x72, 0x36, 0x1f,
	0x8b, 0xc1, 0x83, 0x88, 0x5a, 0xdf, 0x68, 0x53, 0x65, 0x36, 0x1d, 0x52, 0xfb, 0x8a, 0xe8, 0x19,
	0x8c, 0x85, 0x52, 0xb2, 0xf8, 0x2e, 0xca, 0xf3, 0xe7, 0xc1, 0xff, 0x82, 0x06, 0x35, 0x18, 0x4c,
	0x1e, 0xcb, 0xab, 0xd6, 0x40, 0x6a, 0x68, 0xea, 0x0d, 0x4a, 0x17, 0x1a, 0x95, 0x1d, 0x8a, 0x6d,
	0xa0, 0x83, 0xc0, 0x50, 0x31, 0x4e, 0x8f, 0xae, 0x8b, 0x9b, 0xd3, 0xd7, 0x79, 0xa6, 0x9e, 0x7e,
	0xa1, 0x87, 0xcc, 0xec, 0x2d, 0xef, 0x8c, 0x9f, 0xa1, 0x62, 0xfc, 0xc7, 0x72, 0x9c, 0x2c, 0x99,
	0x3a, 0x64, 0x59, 0x7d, 0x94, 0x97, 0x5b, 0xdd, 0xb9, 0x03, 0x9b, 0x57, 0x67, 0xf9, 0xe3, 0xeb,
	0xfd, 0x94, 0x75, 0xf1, 0x76, 0x3f, 0x50, 0xa4, 0x0c, 0x42, 0x4e, 0x1b, 0xf2, 0x84, 0xd2, 0x0f,
	0xf5, 0x29, 0x82, 0x18, 0xda, 0x3f, 0xc7, 0xbd, 0x1b, 0x1f, 0xd5, 0x49, 0x0f, 0xde, 0xfe, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x93, 0x9b, 0x7c, 0xf8, 0x84, 0x01, 0x00, 0x00,
}
