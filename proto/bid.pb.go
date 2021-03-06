// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bid.proto

/*
Package sonm is a generated protocol buffer package.

It is generated from these files:
	bid.proto
	capabilities.proto
	hub.proto
	insonmnia.proto
	marketplace.proto
	miner.proto

It has these top-level messages:
	Geo
	Resources
	Slot
	Order
	Capabilities
	CPUDevice
	RAMDevice
	GPUDevice
	ListRequest
	ListReply
	HubInfoRequest
	TaskRequirements
	HubStartTaskRequest
	HubStartTaskReply
	HubStatusMapRequest
	HubStatusRequest
	HubStatusReply
	DealRequest
	DealReply
	PingRequest
	PingReply
	CPUUsage
	MemoryUsage
	NetworkUsage
	ResourceUsage
	InfoReply
	StopTaskRequest
	StopTaskReply
	TaskStatusRequest
	TaskStatusReply
	StatusMapReply
	ContainerRestartPolicy
	TaskLogsRequest
	TaskLogsChunk
	EmptyReply
	DiscoverHubRequest
	TaskResourceRequirements
	Timestamp
	GetOrdersReply
	GetOrderRequest
	Empty
	MinerInfoRequest
	MinerHandshakeRequest
	MinerHandshakeReply
	MinerStartRequest
	MinerStartReply
	TaskInfo
	MinerStatusMapRequest
*/
package sonm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type OrderType int32

const (
	OrderType_BID OrderType = 0
	OrderType_ASK OrderType = 1
)

var OrderType_name = map[int32]string{
	0: "BID",
	1: "ASK",
}
var OrderType_value = map[string]int32{
	"BID": 0,
	"ASK": 1,
}

func (x OrderType) String() string {
	return proto.EnumName(OrderType_name, int32(x))
}
func (OrderType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type NetworkType int32

const (
	NetworkType_NO_NETWORK NetworkType = 0
	NetworkType_OUTBOUND   NetworkType = 1
	NetworkType_INCOMING   NetworkType = 2
)

var NetworkType_name = map[int32]string{
	0: "NO_NETWORK",
	1: "OUTBOUND",
	2: "INCOMING",
}
var NetworkType_value = map[string]int32{
	"NO_NETWORK": 0,
	"OUTBOUND":   1,
	"INCOMING":   2,
}

func (x NetworkType) String() string {
	return proto.EnumName(NetworkType_name, int32(x))
}
func (NetworkType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// Geo represent GeoIP results for node
type Geo struct {
	Country string  `protobuf:"bytes,1,opt,name=country" json:"country,omitempty"`
	City    string  `protobuf:"bytes,2,opt,name=city" json:"city,omitempty"`
	Lat     float32 `protobuf:"fixed32,3,opt,name=lat" json:"lat,omitempty"`
	Lon     float32 `protobuf:"fixed32,4,opt,name=lon" json:"lon,omitempty"`
}

func (m *Geo) Reset()                    { *m = Geo{} }
func (m *Geo) String() string            { return proto.CompactTextString(m) }
func (*Geo) ProtoMessage()               {}
func (*Geo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Geo) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Geo) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Geo) GetLat() float32 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *Geo) GetLon() float32 {
	if m != nil {
		return m.Lon
	}
	return 0
}

type Resources struct {
	// CPU core count
	CpuCores uint64 `protobuf:"varint,1,opt,name=cpuCores" json:"cpuCores,omitempty"`
	// RAM, in bytes
	RamBytes uint64 `protobuf:"varint,2,opt,name=ramBytes" json:"ramBytes,omitempty"`
	// GPU devices count
	GpuCount uint64 `protobuf:"varint,3,opt,name=gpuCount" json:"gpuCount,omitempty"`
	// todo: discuss
	// storage volume, in Megabytes
	Storage uint64 `protobuf:"varint,4,opt,name=storage" json:"storage,omitempty"`
	// Inbound network traffic (the higher value), in bytes
	NetTrafficIn uint64 `protobuf:"varint,5,opt,name=netTrafficIn" json:"netTrafficIn,omitempty"`
	// Outbound network traffic (the higher value), in bytes
	NetTrafficOut uint64 `protobuf:"varint,6,opt,name=netTrafficOut" json:"netTrafficOut,omitempty"`
	// Allowed network connections
	NetworkType NetworkType `protobuf:"varint,7,opt,name=networkType,enum=sonm.NetworkType" json:"networkType,omitempty"`
	// Other properties/benchmarks. The higher means better
	Props map[string]string `protobuf:"bytes,8,rep,name=props" json:"props,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Resources) Reset()                    { *m = Resources{} }
func (m *Resources) String() string            { return proto.CompactTextString(m) }
func (*Resources) ProtoMessage()               {}
func (*Resources) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Resources) GetCpuCores() uint64 {
	if m != nil {
		return m.CpuCores
	}
	return 0
}

func (m *Resources) GetRamBytes() uint64 {
	if m != nil {
		return m.RamBytes
	}
	return 0
}

func (m *Resources) GetGpuCount() uint64 {
	if m != nil {
		return m.GpuCount
	}
	return 0
}

func (m *Resources) GetStorage() uint64 {
	if m != nil {
		return m.Storage
	}
	return 0
}

func (m *Resources) GetNetTrafficIn() uint64 {
	if m != nil {
		return m.NetTrafficIn
	}
	return 0
}

func (m *Resources) GetNetTrafficOut() uint64 {
	if m != nil {
		return m.NetTrafficOut
	}
	return 0
}

func (m *Resources) GetNetworkType() NetworkType {
	if m != nil {
		return m.NetworkType
	}
	return NetworkType_NO_NETWORK
}

func (m *Resources) GetProps() map[string]string {
	if m != nil {
		return m.Props
	}
	return nil
}

type Slot struct {
	// Virtual computer start of life (hour grained).
	StartTime *Timestamp `protobuf:"bytes,1,opt,name=startTime" json:"startTime,omitempty"`
	// Virtual computer end of life (hour grained).
	EndTime *Timestamp `protobuf:"bytes,2,opt,name=endTime" json:"endTime,omitempty"`
	// Buyer’s rating. Got from Buyer’s profile for BID orders rating_supplier.
	BuyerRating int64 `protobuf:"varint,3,opt,name=buyerRating" json:"buyerRating,omitempty"`
	// Supplier’s rating. Got from Supplier’s profile for ASK orders.
	SupplierRating int64 `protobuf:"varint,4,opt,name=supplierRating" json:"supplierRating,omitempty"`
	// Geo represent Worker's position
	Geo *Geo `protobuf:"bytes,5,opt,name=geo" json:"geo,omitempty"`
	// Hardware resources requirements
	Resources *Resources `protobuf:"bytes,6,opt,name=resources" json:"resources,omitempty"`
}

func (m *Slot) Reset()                    { *m = Slot{} }
func (m *Slot) String() string            { return proto.CompactTextString(m) }
func (*Slot) ProtoMessage()               {}
func (*Slot) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Slot) GetStartTime() *Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Slot) GetEndTime() *Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *Slot) GetBuyerRating() int64 {
	if m != nil {
		return m.BuyerRating
	}
	return 0
}

func (m *Slot) GetSupplierRating() int64 {
	if m != nil {
		return m.SupplierRating
	}
	return 0
}

func (m *Slot) GetGeo() *Geo {
	if m != nil {
		return m.Geo
	}
	return nil
}

func (m *Slot) GetResources() *Resources {
	if m != nil {
		return m.Resources
	}
	return nil
}

type Order struct {
	// Order ID, UUIDv4
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Buyer's EtherumID
	ByuerID string `protobuf:"bytes,2,opt,name=byuerID" json:"byuerID,omitempty"`
	// Supplier's is EtherumID
	SupplierID string `protobuf:"bytes,3,opt,name=supplierID" json:"supplierID,omitempty"`
	// Order price
	Price int64 `protobuf:"varint,4,opt,name=price" json:"price,omitempty"`
	// Order type (Bid or Ask)
	OrderType OrderType `protobuf:"varint,5,opt,name=orderType,enum=sonm.OrderType" json:"orderType,omitempty"`
	// Slot describe resource requiements
	Slot *Slot `protobuf:"bytes,6,opt,name=slot" json:"slot,omitempty"`
}

func (m *Order) Reset()                    { *m = Order{} }
func (m *Order) String() string            { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()               {}
func (*Order) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Order) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Order) GetByuerID() string {
	if m != nil {
		return m.ByuerID
	}
	return ""
}

func (m *Order) GetSupplierID() string {
	if m != nil {
		return m.SupplierID
	}
	return ""
}

func (m *Order) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Order) GetOrderType() OrderType {
	if m != nil {
		return m.OrderType
	}
	return OrderType_BID
}

func (m *Order) GetSlot() *Slot {
	if m != nil {
		return m.Slot
	}
	return nil
}

func init() {
	proto.RegisterType((*Geo)(nil), "sonm.Geo")
	proto.RegisterType((*Resources)(nil), "sonm.Resources")
	proto.RegisterType((*Slot)(nil), "sonm.Slot")
	proto.RegisterType((*Order)(nil), "sonm.Order")
	proto.RegisterEnum("sonm.OrderType", OrderType_name, OrderType_value)
	proto.RegisterEnum("sonm.NetworkType", NetworkType_name, NetworkType_value)
}

func init() { proto.RegisterFile("bid.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 559 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0x8d, 0x3e, 0x1c, 0x47, 0xa3, 0xd4, 0x51, 0x87, 0x1e, 0x84, 0x4b, 0x83, 0x31, 0xa5, 0xb8,
	0x81, 0x9a, 0xe2, 0x5c, 0xd2, 0xde, 0xea, 0x38, 0x18, 0x11, 0x2a, 0x95, 0x8d, 0x43, 0x8e, 0x45,
	0x96, 0x37, 0x46, 0xc4, 0xde, 0x15, 0xab, 0x55, 0x8b, 0xfe, 0x58, 0x7f, 0x51, 0x7f, 0x43, 0xcf,
	0x65, 0x57, 0x5f, 0x4e, 0xe9, 0xc9, 0xf3, 0xde, 0x3c, 0xf1, 0x66, 0xdf, 0x8c, 0xc1, 0x59, 0xa7,
	0x9b, 0x69, 0x26, 0xb8, 0xe4, 0x68, 0xe7, 0x9c, 0xed, 0x87, 0x67, 0x29, 0x53, 0xbf, 0x2c, 0x8d,
	0x2b, 0x7a, 0xfc, 0x00, 0xd6, 0x92, 0x72, 0xf4, 0xa1, 0x9f, 0xf0, 0x82, 0x49, 0x51, 0xfa, 0xc6,
	0xc8, 0x98, 0x38, 0xa4, 0x81, 0x88, 0x60, 0x27, 0xa9, 0x2c, 0x7d, 0x53, 0xd3, 0xba, 0x46, 0x0f,
	0xac, 0x5d, 0x2c, 0x7d, 0x6b, 0x64, 0x4c, 0x4c, 0xa2, 0x4a, 0xcd, 0x70, 0xe6, 0xdb, 0x35, 0xc3,
	0xd9, 0xf8, 0xb7, 0x09, 0x0e, 0xa1, 0x39, 0x2f, 0x44, 0x42, 0x73, 0x1c, 0xc2, 0x49, 0x92, 0x15,
	0xd7, 0x5c, 0xd0, 0x5c, 0x1b, 0xd8, 0xa4, 0xc5, 0xaa, 0x27, 0xe2, 0xfd, 0xbc, 0x94, 0x34, 0xd7,
	0x2e, 0x36, 0x69, 0xb1, 0xea, 0x6d, 0x95, 0xae, 0x60, 0x95, 0x9d, 0x4d, 0x5a, 0xac, 0x66, 0xce,
	0x25, 0x17, 0xf1, 0x96, 0x6a, 0x5f, 0x9b, 0x34, 0x10, 0xc7, 0x70, 0xca, 0xa8, 0x5c, 0x89, 0xf8,
	0xf1, 0x31, 0x4d, 0x02, 0xe6, 0xf7, 0x74, 0xfb, 0x19, 0x87, 0x6f, 0xe1, 0x45, 0x87, 0xa3, 0x42,
	0xfa, 0xc7, 0x5a, 0xf4, 0x9c, 0xc4, 0x4b, 0x70, 0x19, 0x95, 0x3f, 0xb9, 0x78, 0x5a, 0x95, 0x19,
	0xf5, 0xfb, 0x23, 0x63, 0x32, 0x98, 0xbd, 0x9c, 0xaa, 0x0c, 0xa7, 0x61, 0xd7, 0x20, 0x87, 0x2a,
	0xfc, 0x08, 0xbd, 0x4c, 0xf0, 0x2c, 0xf7, 0x4f, 0x46, 0xd6, 0xc4, 0x9d, 0x0d, 0x2b, 0x79, 0x1b,
	0xc6, 0xf4, 0x9b, 0x6a, 0xde, 0xa8, 0x74, 0x49, 0x25, 0x1c, 0x5e, 0x01, 0x74, 0xa4, 0x0a, 0xf3,
	0x89, 0x36, 0x8b, 0x50, 0x25, 0xbe, 0x82, 0xde, 0x8f, 0x78, 0x57, 0xd0, 0x7a, 0x0b, 0x15, 0xf8,
	0x6c, 0x5e, 0x19, 0xe3, 0x3f, 0x06, 0xd8, 0x77, 0x3b, 0x2e, 0xf1, 0x03, 0x38, 0xb9, 0x8c, 0x85,
	0x5c, 0xa5, 0x7b, 0xaa, 0x3f, 0x75, 0x67, 0x67, 0x95, 0xb1, 0x62, 0x72, 0x19, 0xef, 0x33, 0xd2,
	0x29, 0xf0, 0x3d, 0xf4, 0x29, 0xdb, 0x68, 0xb1, 0xf9, 0x7f, 0x71, 0xd3, 0xc7, 0x11, 0xb8, 0xeb,
	0xa2, 0xa4, 0x82, 0xc4, 0x32, 0x65, 0x5b, 0xbd, 0x06, 0x8b, 0x1c, 0x52, 0xf8, 0x0e, 0x06, 0x79,
	0x91, 0x65, 0xbb, 0xb4, 0x15, 0xd9, 0x5a, 0xf4, 0x0f, 0x8b, 0xaf, 0xc1, 0xda, 0x52, 0xae, 0xd7,
	0xe1, 0xce, 0x9c, 0xca, 0x70, 0x49, 0x39, 0x51, 0xac, 0x7a, 0x80, 0x68, 0x22, 0xd2, 0xcb, 0x68,
	0x67, 0x6a, 0x93, 0x23, 0x9d, 0x62, 0xfc, 0xcb, 0x80, 0x5e, 0x24, 0x36, 0x54, 0xe0, 0x00, 0xcc,
	0x74, 0x53, 0xa7, 0x65, 0xa6, 0x1b, 0x75, 0x17, 0xeb, 0xb2, 0xa0, 0x22, 0x58, 0xd4, 0x71, 0x35,
	0x10, 0xcf, 0x01, 0x9a, 0x89, 0x82, 0x85, 0x7e, 0x88, 0x43, 0x0e, 0x18, 0x15, 0x73, 0x26, 0xd2,
	0x84, 0xd6, 0xe3, 0x57, 0x40, 0x0d, 0xc6, 0x95, 0x91, 0xbe, 0x80, 0x9e, 0xbe, 0x80, 0x7a, 0xb0,
	0xa8, 0xa1, 0x49, 0xa7, 0xc0, 0x73, 0xb0, 0xf3, 0x1d, 0x97, 0xf5, 0x13, 0xa0, 0x52, 0xaa, 0x15,
	0x11, 0xcd, 0x5f, 0xbc, 0x01, 0xa7, 0xfd, 0x0e, 0xfb, 0x60, 0xcd, 0x83, 0x85, 0x77, 0xa4, 0x8a,
	0x2f, 0x77, 0xb7, 0x9e, 0x71, 0xf1, 0x09, 0xdc, 0x83, 0xc3, 0xc2, 0x01, 0x40, 0x18, 0x7d, 0x0f,
	0x6f, 0x56, 0x0f, 0x11, 0xb9, 0xf5, 0x8e, 0xf0, 0x14, 0x4e, 0xa2, 0xfb, 0xd5, 0x3c, 0xba, 0x0f,
	0x17, 0x9e, 0xa1, 0x50, 0x10, 0x5e, 0x47, 0x5f, 0x83, 0x70, 0xe9, 0x99, 0xeb, 0x63, 0xfd, 0x97,
	0xbe, 0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x47, 0x14, 0xa4, 0xc8, 0xf6, 0x03, 0x00, 0x00,
}
