// Code generated by protoc-gen-go. DO NOT EDIT.
// source: trackable_object_graph.proto

package core_protos_go_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type TrackableObjectGraph struct {
	Nodes                []*TrackableObjectGraph_TrackableObject `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *TrackableObjectGraph) Reset()         { *m = TrackableObjectGraph{} }
func (m *TrackableObjectGraph) String() string { return proto.CompactTextString(m) }
func (*TrackableObjectGraph) ProtoMessage()    {}
func (*TrackableObjectGraph) Descriptor() ([]byte, []int) {
	return fileDescriptor_981c99d5aab8d094, []int{0}
}

func (m *TrackableObjectGraph) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrackableObjectGraph.Unmarshal(m, b)
}
func (m *TrackableObjectGraph) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrackableObjectGraph.Marshal(b, m, deterministic)
}
func (m *TrackableObjectGraph) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrackableObjectGraph.Merge(m, src)
}
func (m *TrackableObjectGraph) XXX_Size() int {
	return xxx_messageInfo_TrackableObjectGraph.Size(m)
}
func (m *TrackableObjectGraph) XXX_DiscardUnknown() {
	xxx_messageInfo_TrackableObjectGraph.DiscardUnknown(m)
}

var xxx_messageInfo_TrackableObjectGraph proto.InternalMessageInfo

func (m *TrackableObjectGraph) GetNodes() []*TrackableObjectGraph_TrackableObject {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type TrackableObjectGraph_TrackableObject struct {
	// Objects which this object depends on.
	Children []*TrackableObjectGraph_TrackableObject_ObjectReference `protobuf:"bytes,1,rep,name=children,proto3" json:"children,omitempty"`
	// Serialized data specific to this object.
	Attributes []*TrackableObjectGraph_TrackableObject_SerializedTensor `protobuf:"bytes,2,rep,name=attributes,proto3" json:"attributes,omitempty"`
	// Slot variables owned by this object.
	SlotVariables        []*TrackableObjectGraph_TrackableObject_SlotVariableReference `protobuf:"bytes,3,rep,name=slot_variables,json=slotVariables,proto3" json:"slot_variables,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                      `json:"-"`
	XXX_unrecognized     []byte                                                        `json:"-"`
	XXX_sizecache        int32                                                         `json:"-"`
}

func (m *TrackableObjectGraph_TrackableObject) Reset()         { *m = TrackableObjectGraph_TrackableObject{} }
func (m *TrackableObjectGraph_TrackableObject) String() string { return proto.CompactTextString(m) }
func (*TrackableObjectGraph_TrackableObject) ProtoMessage()    {}
func (*TrackableObjectGraph_TrackableObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_981c99d5aab8d094, []int{0, 0}
}

func (m *TrackableObjectGraph_TrackableObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrackableObjectGraph_TrackableObject.Unmarshal(m, b)
}
func (m *TrackableObjectGraph_TrackableObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrackableObjectGraph_TrackableObject.Marshal(b, m, deterministic)
}
func (m *TrackableObjectGraph_TrackableObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrackableObjectGraph_TrackableObject.Merge(m, src)
}
func (m *TrackableObjectGraph_TrackableObject) XXX_Size() int {
	return xxx_messageInfo_TrackableObjectGraph_TrackableObject.Size(m)
}
func (m *TrackableObjectGraph_TrackableObject) XXX_DiscardUnknown() {
	xxx_messageInfo_TrackableObjectGraph_TrackableObject.DiscardUnknown(m)
}

var xxx_messageInfo_TrackableObjectGraph_TrackableObject proto.InternalMessageInfo

func (m *TrackableObjectGraph_TrackableObject) GetChildren() []*TrackableObjectGraph_TrackableObject_ObjectReference {
	if m != nil {
		return m.Children
	}
	return nil
}

func (m *TrackableObjectGraph_TrackableObject) GetAttributes() []*TrackableObjectGraph_TrackableObject_SerializedTensor {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *TrackableObjectGraph_TrackableObject) GetSlotVariables() []*TrackableObjectGraph_TrackableObject_SlotVariableReference {
	if m != nil {
		return m.SlotVariables
	}
	return nil
}

type TrackableObjectGraph_TrackableObject_ObjectReference struct {
	// An index into `TrackableObjectGraph.nodes`, indicating the object
	// being referenced.
	NodeId int32 `protobuf:"varint,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// A user-provided name for the edge.
	LocalName            string   `protobuf:"bytes,2,opt,name=local_name,json=localName,proto3" json:"local_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrackableObjectGraph_TrackableObject_ObjectReference) Reset() {
	*m = TrackableObjectGraph_TrackableObject_ObjectReference{}
}
func (m *TrackableObjectGraph_TrackableObject_ObjectReference) String() string {
	return proto.CompactTextString(m)
}
func (*TrackableObjectGraph_TrackableObject_ObjectReference) ProtoMessage() {}
func (*TrackableObjectGraph_TrackableObject_ObjectReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_981c99d5aab8d094, []int{0, 0, 0}
}

func (m *TrackableObjectGraph_TrackableObject_ObjectReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrackableObjectGraph_TrackableObject_ObjectReference.Unmarshal(m, b)
}
func (m *TrackableObjectGraph_TrackableObject_ObjectReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrackableObjectGraph_TrackableObject_ObjectReference.Marshal(b, m, deterministic)
}
func (m *TrackableObjectGraph_TrackableObject_ObjectReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrackableObjectGraph_TrackableObject_ObjectReference.Merge(m, src)
}
func (m *TrackableObjectGraph_TrackableObject_ObjectReference) XXX_Size() int {
	return xxx_messageInfo_TrackableObjectGraph_TrackableObject_ObjectReference.Size(m)
}
func (m *TrackableObjectGraph_TrackableObject_ObjectReference) XXX_DiscardUnknown() {
	xxx_messageInfo_TrackableObjectGraph_TrackableObject_ObjectReference.DiscardUnknown(m)
}

var xxx_messageInfo_TrackableObjectGraph_TrackableObject_ObjectReference proto.InternalMessageInfo

func (m *TrackableObjectGraph_TrackableObject_ObjectReference) GetNodeId() int32 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *TrackableObjectGraph_TrackableObject_ObjectReference) GetLocalName() string {
	if m != nil {
		return m.LocalName
	}
	return ""
}

type TrackableObjectGraph_TrackableObject_SerializedTensor struct {
	// A name for the Tensor. Simple variables have only one
	// `SerializedTensor` named "VARIABLE_VALUE" by convention. This value may
	// be restored on object creation as an optimization.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The full name of the variable/tensor, if applicable. Used to allow
	// name-based loading of checkpoints which were saved using an
	// object-based API. Should match the checkpoint key which would have been
	// assigned by tf.train.Saver.
	FullName string `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	// The generated name of the Tensor in the checkpoint.
	CheckpointKey string `protobuf:"bytes,3,opt,name=checkpoint_key,json=checkpointKey,proto3" json:"checkpoint_key,omitempty"`
	// Whether checkpoints should be considered as matching even without this
	// value restored. Used for non-critical values which don't affect the
	// TensorFlow graph, such as layer configurations.
	OptionalRestore      bool     `protobuf:"varint,4,opt,name=optional_restore,json=optionalRestore,proto3" json:"optional_restore,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrackableObjectGraph_TrackableObject_SerializedTensor) Reset() {
	*m = TrackableObjectGraph_TrackableObject_SerializedTensor{}
}
func (m *TrackableObjectGraph_TrackableObject_SerializedTensor) String() string {
	return proto.CompactTextString(m)
}
func (*TrackableObjectGraph_TrackableObject_SerializedTensor) ProtoMessage() {}
func (*TrackableObjectGraph_TrackableObject_SerializedTensor) Descriptor() ([]byte, []int) {
	return fileDescriptor_981c99d5aab8d094, []int{0, 0, 1}
}

func (m *TrackableObjectGraph_TrackableObject_SerializedTensor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrackableObjectGraph_TrackableObject_SerializedTensor.Unmarshal(m, b)
}
func (m *TrackableObjectGraph_TrackableObject_SerializedTensor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrackableObjectGraph_TrackableObject_SerializedTensor.Marshal(b, m, deterministic)
}
func (m *TrackableObjectGraph_TrackableObject_SerializedTensor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrackableObjectGraph_TrackableObject_SerializedTensor.Merge(m, src)
}
func (m *TrackableObjectGraph_TrackableObject_SerializedTensor) XXX_Size() int {
	return xxx_messageInfo_TrackableObjectGraph_TrackableObject_SerializedTensor.Size(m)
}
func (m *TrackableObjectGraph_TrackableObject_SerializedTensor) XXX_DiscardUnknown() {
	xxx_messageInfo_TrackableObjectGraph_TrackableObject_SerializedTensor.DiscardUnknown(m)
}

var xxx_messageInfo_TrackableObjectGraph_TrackableObject_SerializedTensor proto.InternalMessageInfo

func (m *TrackableObjectGraph_TrackableObject_SerializedTensor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TrackableObjectGraph_TrackableObject_SerializedTensor) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *TrackableObjectGraph_TrackableObject_SerializedTensor) GetCheckpointKey() string {
	if m != nil {
		return m.CheckpointKey
	}
	return ""
}

func (m *TrackableObjectGraph_TrackableObject_SerializedTensor) GetOptionalRestore() bool {
	if m != nil {
		return m.OptionalRestore
	}
	return false
}

type TrackableObjectGraph_TrackableObject_SlotVariableReference struct {
	// An index into `TrackableObjectGraph.nodes`, indicating the
	// variable object this slot was created for.
	OriginalVariableNodeId int32 `protobuf:"varint,1,opt,name=original_variable_node_id,json=originalVariableNodeId,proto3" json:"original_variable_node_id,omitempty"`
	// The name of the slot (e.g. "m"/"v").
	SlotName string `protobuf:"bytes,2,opt,name=slot_name,json=slotName,proto3" json:"slot_name,omitempty"`
	// An index into `TrackableObjectGraph.nodes`, indicating the
	// `Object` with the value of the slot variable.
	SlotVariableNodeId   int32    `protobuf:"varint,3,opt,name=slot_variable_node_id,json=slotVariableNodeId,proto3" json:"slot_variable_node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrackableObjectGraph_TrackableObject_SlotVariableReference) Reset() {
	*m = TrackableObjectGraph_TrackableObject_SlotVariableReference{}
}
func (m *TrackableObjectGraph_TrackableObject_SlotVariableReference) String() string {
	return proto.CompactTextString(m)
}
func (*TrackableObjectGraph_TrackableObject_SlotVariableReference) ProtoMessage() {}
func (*TrackableObjectGraph_TrackableObject_SlotVariableReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_981c99d5aab8d094, []int{0, 0, 2}
}

func (m *TrackableObjectGraph_TrackableObject_SlotVariableReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrackableObjectGraph_TrackableObject_SlotVariableReference.Unmarshal(m, b)
}
func (m *TrackableObjectGraph_TrackableObject_SlotVariableReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrackableObjectGraph_TrackableObject_SlotVariableReference.Marshal(b, m, deterministic)
}
func (m *TrackableObjectGraph_TrackableObject_SlotVariableReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrackableObjectGraph_TrackableObject_SlotVariableReference.Merge(m, src)
}
func (m *TrackableObjectGraph_TrackableObject_SlotVariableReference) XXX_Size() int {
	return xxx_messageInfo_TrackableObjectGraph_TrackableObject_SlotVariableReference.Size(m)
}
func (m *TrackableObjectGraph_TrackableObject_SlotVariableReference) XXX_DiscardUnknown() {
	xxx_messageInfo_TrackableObjectGraph_TrackableObject_SlotVariableReference.DiscardUnknown(m)
}

var xxx_messageInfo_TrackableObjectGraph_TrackableObject_SlotVariableReference proto.InternalMessageInfo

func (m *TrackableObjectGraph_TrackableObject_SlotVariableReference) GetOriginalVariableNodeId() int32 {
	if m != nil {
		return m.OriginalVariableNodeId
	}
	return 0
}

func (m *TrackableObjectGraph_TrackableObject_SlotVariableReference) GetSlotName() string {
	if m != nil {
		return m.SlotName
	}
	return ""
}

func (m *TrackableObjectGraph_TrackableObject_SlotVariableReference) GetSlotVariableNodeId() int32 {
	if m != nil {
		return m.SlotVariableNodeId
	}
	return 0
}

func init() {
	proto.RegisterType((*TrackableObjectGraph)(nil), "tensorflow.TrackableObjectGraph")
	proto.RegisterType((*TrackableObjectGraph_TrackableObject)(nil), "tensorflow.TrackableObjectGraph.TrackableObject")
	proto.RegisterType((*TrackableObjectGraph_TrackableObject_ObjectReference)(nil), "tensorflow.TrackableObjectGraph.TrackableObject.ObjectReference")
	proto.RegisterType((*TrackableObjectGraph_TrackableObject_SerializedTensor)(nil), "tensorflow.TrackableObjectGraph.TrackableObject.SerializedTensor")
	proto.RegisterType((*TrackableObjectGraph_TrackableObject_SlotVariableReference)(nil), "tensorflow.TrackableObjectGraph.TrackableObject.SlotVariableReference")
}

func init() {
	proto.RegisterFile("trackable_object_graph.proto", fileDescriptor_981c99d5aab8d094)
}

var fileDescriptor_981c99d5aab8d094 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xb5, 0x4d, 0x13, 0x92, 0x41, 0x6d, 0xaa, 0x15, 0x05, 0x93, 0x82, 0x14, 0x21, 0x21,
	0x85, 0x8b, 0xcb, 0x9f, 0x13, 0x37, 0xe0, 0x50, 0xa8, 0x10, 0x45, 0x32, 0x15, 0x07, 0x84, 0xb4,
	0x5a, 0xaf, 0x27, 0xce, 0x92, 0x8d, 0x27, 0xda, 0xdd, 0x80, 0xca, 0x7b, 0xf0, 0x02, 0x1c, 0x79,
	0x13, 0xde, 0x88, 0x23, 0xf2, 0x3a, 0x6e, 0x12, 0x2b, 0x3d, 0xe4, 0x62, 0x8d, 0x3f, 0x7f, 0xf3,
	0x1b, 0x7d, 0x33, 0x32, 0x3c, 0xf0, 0x56, 0xaa, 0xa9, 0x4c, 0x0d, 0x0a, 0x4a, 0xbf, 0xa1, 0xf2,
	0x22, 0xb7, 0x72, 0x3e, 0x89, 0xe7, 0x96, 0x3c, 0x71, 0xf0, 0x58, 0x38, 0xb2, 0x63, 0x43, 0x3f,
	0x1e, 0xfd, 0xe9, 0xc0, 0x9d, 0xcb, 0xda, 0xfc, 0x31, 0x78, 0xdf, 0x96, 0x56, 0x7e, 0x06, 0xed,
	0x82, 0x32, 0x74, 0x11, 0x1b, 0xb6, 0x46, 0xb7, 0x9f, 0x3f, 0x8d, 0x57, 0x4d, 0xf1, 0xb6, 0x86,
	0xa6, 0x98, 0x54, 0xed, 0x83, 0xbf, 0x6d, 0xe8, 0x37, 0x3e, 0xf1, 0xaf, 0xd0, 0x55, 0x13, 0x6d,
	0x32, 0x8b, 0xc5, 0x12, 0xff, 0x6a, 0x57, 0x7c, 0xbc, 0x9c, 0x82, 0x63, 0xb4, 0x58, 0x28, 0x4c,
	0xae, 0x89, 0x5c, 0x02, 0x48, 0xef, 0xad, 0x4e, 0x17, 0x1e, 0x5d, 0xb4, 0x17, 0xf8, 0xaf, 0x77,
	0xe6, 0x7f, 0x42, 0xab, 0xa5, 0xd1, 0x3f, 0x31, 0xbb, 0x0c, 0x9d, 0xc9, 0x1a, 0x94, 0xcf, 0xe0,
	0xd0, 0x19, 0xf2, 0xe2, 0xbb, 0xb4, 0xba, 0xec, 0x71, 0x51, 0x2b, 0x8c, 0x39, 0xdb, 0x7d, 0x8c,
	0x21, 0xff, 0x79, 0x49, 0x59, 0x85, 0x39, 0x70, 0x6b, 0xb2, 0x1b, 0x9c, 0x43, 0xbf, 0x11, 0x97,
	0xdf, 0x83, 0x5b, 0xe5, 0x7e, 0x85, 0xce, 0x22, 0x36, 0x64, 0xa3, 0x76, 0xd2, 0x29, 0x5f, 0xcf,
	0x33, 0xfe, 0x10, 0xc0, 0x90, 0x92, 0x46, 0x14, 0x72, 0x86, 0xd1, 0xde, 0x90, 0x8d, 0x7a, 0x49,
	0x2f, 0x28, 0x17, 0x72, 0x86, 0x83, 0x5f, 0x0c, 0x8e, 0x9a, 0xd1, 0x38, 0x87, 0xfd, 0xe0, 0x66,
	0xc1, 0x1d, 0x6a, 0x7e, 0x02, 0xbd, 0xf1, 0xc2, 0x6c, 0x60, 0xba, 0xa5, 0x50, 0x52, 0xf8, 0x63,
	0x38, 0x54, 0x13, 0x54, 0xd3, 0x39, 0xe9, 0xc2, 0x8b, 0x29, 0x5e, 0x45, 0xad, 0xe0, 0x38, 0x58,
	0xa9, 0xef, 0xf1, 0x8a, 0x3f, 0x81, 0x23, 0x9a, 0x7b, 0x4d, 0x85, 0x34, 0xc2, 0xa2, 0xf3, 0x64,
	0x31, 0xda, 0x1f, 0xb2, 0x51, 0x37, 0xe9, 0xd7, 0x7a, 0x52, 0xc9, 0x83, 0xdf, 0x0c, 0x8e, 0xb7,
	0xee, 0x82, 0xbf, 0x84, 0xfb, 0x64, 0x75, 0xae, 0x4b, 0x48, 0xbd, 0x6f, 0xb1, 0x99, 0xfd, 0x6e,
	0x6d, 0xa8, 0xbb, 0x2f, 0xaa, 0x5d, 0x9c, 0x40, 0x2f, 0x9c, 0x69, 0x3d, 0x43, 0x29, 0x84, 0x0c,
	0xcf, 0xe0, 0x78, 0xe3, 0x86, 0xd7, 0xcc, 0x56, 0x60, 0xf2, 0xf5, 0x13, 0x54, 0xbc, 0x37, 0x1f,
	0xbe, 0xbc, 0xcb, 0xb5, 0x9f, 0x2c, 0xd2, 0x58, 0xd1, 0xec, 0x74, 0x75, 0xea, 0x1b, 0xca, 0x9c,
	0x4e, 0x15, 0x59, 0x0c, 0x0f, 0x11, 0x7e, 0x3d, 0x27, 0x72, 0xaa, 0xaa, 0x7f, 0x8c, 0xa5, 0x9d,
	0x50, 0xbd, 0xf8, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x76, 0x4a, 0xa1, 0xae, 0x03, 0x00, 0x00,
}
