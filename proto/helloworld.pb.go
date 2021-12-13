// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: helloworld.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PlanetaCiudad struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *PlanetaCiudad) Reset() {
	*x = PlanetaCiudad{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlanetaCiudad) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanetaCiudad) ProtoMessage() {}

func (x *PlanetaCiudad) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanetaCiudad.ProtoReflect.Descriptor instead.
func (*PlanetaCiudad) Descriptor() ([]byte, []int) {
	return file_helloworld_proto_rawDescGZIP(), []int{0}
}

func (x *PlanetaCiudad) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type Numero struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num string `protobuf:"bytes,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *Numero) Reset() {
	*x = Numero{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Numero) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Numero) ProtoMessage() {}

func (x *Numero) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Numero.ProtoReflect.Descriptor instead.
func (*Numero) Descriptor() ([]byte, []int) {
	return file_helloworld_proto_rawDescGZIP(), []int{1}
}

func (x *Numero) GetNum() string {
	if x != nil {
		return x.Num
	}
	return ""
}

type Comando struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comando string `protobuf:"bytes,1,opt,name=comando,proto3" json:"comando,omitempty"`
}

func (x *Comando) Reset() {
	*x = Comando{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comando) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comando) ProtoMessage() {}

func (x *Comando) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comando.ProtoReflect.Descriptor instead.
func (*Comando) Descriptor() ([]byte, []int) {
	return file_helloworld_proto_rawDescGZIP(), []int{2}
}

func (x *Comando) GetComando() string {
	if x != nil {
		return x.Comando
	}
	return ""
}

type Redirigido struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valor string `protobuf:"bytes,1,opt,name=valor,proto3" json:"valor,omitempty"`
}

func (x *Redirigido) Reset() {
	*x = Redirigido{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Redirigido) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Redirigido) ProtoMessage() {}

func (x *Redirigido) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Redirigido.ProtoReflect.Descriptor instead.
func (*Redirigido) Descriptor() ([]byte, []int) {
	return file_helloworld_proto_rawDescGZIP(), []int{3}
}

func (x *Redirigido) GetValor() string {
	if x != nil {
		return x.Valor
	}
	return ""
}

type Estructura struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Planeta  string `protobuf:"bytes,1,opt,name=planeta,proto3" json:"planeta,omitempty"`
	Ciudad   string `protobuf:"bytes,2,opt,name=ciudad,proto3" json:"ciudad,omitempty"`
	Rebeldes string `protobuf:"bytes,3,opt,name=rebeldes,proto3" json:"rebeldes,omitempty"`
}

func (x *Estructura) Reset() {
	*x = Estructura{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Estructura) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Estructura) ProtoMessage() {}

func (x *Estructura) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Estructura.ProtoReflect.Descriptor instead.
func (*Estructura) Descriptor() ([]byte, []int) {
	return file_helloworld_proto_rawDescGZIP(), []int{4}
}

func (x *Estructura) GetPlaneta() string {
	if x != nil {
		return x.Planeta
	}
	return ""
}

func (x *Estructura) GetCiudad() string {
	if x != nil {
		return x.Ciudad
	}
	return ""
}

func (x *Estructura) GetRebeldes() string {
	if x != nil {
		return x.Rebeldes
	}
	return ""
}

type Estructura2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Planeta  string `protobuf:"bytes,1,opt,name=planeta,proto3" json:"planeta,omitempty"`
	NomViejo string `protobuf:"bytes,2,opt,name=nom_viejo,json=nomViejo,proto3" json:"nom_viejo,omitempty"`
	NomNuevo string `protobuf:"bytes,3,opt,name=nom_nuevo,json=nomNuevo,proto3" json:"nom_nuevo,omitempty"`
}

func (x *Estructura2) Reset() {
	*x = Estructura2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Estructura2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Estructura2) ProtoMessage() {}

func (x *Estructura2) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Estructura2.ProtoReflect.Descriptor instead.
func (*Estructura2) Descriptor() ([]byte, []int) {
	return file_helloworld_proto_rawDescGZIP(), []int{5}
}

func (x *Estructura2) GetPlaneta() string {
	if x != nil {
		return x.Planeta
	}
	return ""
}

func (x *Estructura2) GetNomViejo() string {
	if x != nil {
		return x.NomViejo
	}
	return ""
}

func (x *Estructura2) GetNomNuevo() string {
	if x != nil {
		return x.NomNuevo
	}
	return ""
}

type Estructura3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Planeta string `protobuf:"bytes,1,opt,name=planeta,proto3" json:"planeta,omitempty"`
	Ciudad  string `protobuf:"bytes,2,opt,name=ciudad,proto3" json:"ciudad,omitempty"`
}

func (x *Estructura3) Reset() {
	*x = Estructura3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Estructura3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Estructura3) ProtoMessage() {}

func (x *Estructura3) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Estructura3.ProtoReflect.Descriptor instead.
func (*Estructura3) Descriptor() ([]byte, []int) {
	return file_helloworld_proto_rawDescGZIP(), []int{6}
}

func (x *Estructura3) GetPlaneta() string {
	if x != nil {
		return x.Planeta
	}
	return ""
}

func (x *Estructura3) GetCiudad() string {
	if x != nil {
		return x.Ciudad
	}
	return ""
}

type Vector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	Z int32 `protobuf:"varint,3,opt,name=z,proto3" json:"z,omitempty"`
}

func (x *Vector) Reset() {
	*x = Vector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vector) ProtoMessage() {}

func (x *Vector) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vector.ProtoReflect.Descriptor instead.
func (*Vector) Descriptor() ([]byte, []int) {
	return file_helloworld_proto_rawDescGZIP(), []int{7}
}

func (x *Vector) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Vector) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Vector) GetZ() int32 {
	if x != nil {
		return x.Z
	}
	return 0
}

var File_helloworld_proto protoreflect.FileDescriptor

var file_helloworld_proto_rawDesc = []byte{
	0x0a, 0x10, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x22, 0x23,
	0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x43, 0x69, 0x75, 0x64, 0x61, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x22, 0x1a, 0x0a, 0x06, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x6f, 0x12, 0x10, 0x0a,
	0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22,
	0x23, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x61, 0x6e, 0x64, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6d, 0x61, 0x6e, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d,
	0x61, 0x6e, 0x64, 0x6f, 0x22, 0x22, 0x0a, 0x0a, 0x52, 0x65, 0x64, 0x69, 0x72, 0x69, 0x67, 0x69,
	0x64, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x22, 0x5a, 0x0a, 0x0a, 0x45, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x75, 0x72, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x69, 0x75, 0x64, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x63, 0x69, 0x75, 0x64, 0x61, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x62, 0x65,
	0x6c, 0x64, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x62, 0x65,
	0x6c, 0x64, 0x65, 0x73, 0x22, 0x61, 0x0a, 0x0b, 0x45, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75,
	0x72, 0x61, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x12, 0x1b, 0x0a,
	0x09, 0x6e, 0x6f, 0x6d, 0x5f, 0x76, 0x69, 0x65, 0x6a, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6e, 0x6f, 0x6d, 0x56, 0x69, 0x65, 0x6a, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f,
	0x6d, 0x5f, 0x6e, 0x75, 0x65, 0x76, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e,
	0x6f, 0x6d, 0x4e, 0x75, 0x65, 0x76, 0x6f, 0x22, 0x3f, 0x0a, 0x0b, 0x45, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x75, 0x72, 0x61, 0x33, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x69, 0x75, 0x64, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x63, 0x69, 0x75, 0x64, 0x61, 0x64, 0x22, 0x32, 0x0a, 0x06, 0x56, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78,
	0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x12, 0x0c,
	0x0a, 0x01, 0x7a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x7a, 0x32, 0x8a, 0x01, 0x0a,
	0x06, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x12, 0x42, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x19, 0x2e, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x43,
	0x69, 0x75, 0x64, 0x61, 0x64, 0x1a, 0x12, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x2e, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x6f, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0b, 0x51,
	0x75, 0x69, 0x65, 0x72, 0x6f, 0x48, 0x61, 0x63, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x43, 0x6f, 0x6d, 0x61, 0x6e, 0x64, 0x6f, 0x1a,
	0x16, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x52, 0x65, 0x64,
	0x69, 0x72, 0x69, 0x67, 0x69, 0x64, 0x6f, 0x22, 0x00, 0x32, 0xc3, 0x02, 0x0a, 0x07, 0x46, 0x75,
	0x6c, 0x63, 0x72, 0x75, 0x6d, 0x12, 0x47, 0x0a, 0x14, 0x50, 0x72, 0x65, 0x67, 0x75, 0x6e, 0x74,
	0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x74, 0x65, 0x73, 0x12, 0x19, 0x2e,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x65,
	0x74, 0x61, 0x43, 0x69, 0x75, 0x64, 0x61, 0x64, 0x1a, 0x12, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x6f, 0x22, 0x00, 0x12, 0x37,
	0x0a, 0x07, 0x41, 0x64, 0x64, 0x43, 0x69, 0x74, 0x79, 0x12, 0x16, 0x2e, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x45, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72,
	0x61, 0x1a, 0x12, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x56,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x2e, 0x45, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x61, 0x32, 0x1a, 0x12,
	0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x56, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x2e, 0x45, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x61, 0x32, 0x1a, 0x12, 0x2e,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x69, 0x74,
	0x79, 0x12, 0x16, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x45,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x61, 0x1a, 0x12, 0x2e, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x00, 0x42,
	0x0c, 0x5a, 0x0a, 0x4c, 0x61, 0x62, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_helloworld_proto_rawDescOnce sync.Once
	file_helloworld_proto_rawDescData = file_helloworld_proto_rawDesc
)

func file_helloworld_proto_rawDescGZIP() []byte {
	file_helloworld_proto_rawDescOnce.Do(func() {
		file_helloworld_proto_rawDescData = protoimpl.X.CompressGZIP(file_helloworld_proto_rawDescData)
	})
	return file_helloworld_proto_rawDescData
}

var file_helloworld_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_helloworld_proto_goTypes = []interface{}{
	(*PlanetaCiudad)(nil), // 0: helloworld.PlanetaCiudad
	(*Numero)(nil),        // 1: helloworld.Numero
	(*Comando)(nil),       // 2: helloworld.Comando
	(*Redirigido)(nil),    // 3: helloworld.Redirigido
	(*Estructura)(nil),    // 4: helloworld.Estructura
	(*Estructura2)(nil),   // 5: helloworld.Estructura2
	(*Estructura3)(nil),   // 6: helloworld.Estructura3
	(*Vector)(nil),        // 7: helloworld.Vector
}
var file_helloworld_proto_depIdxs = []int32{
	0, // 0: helloworld.Broker.GetNumberRebels:input_type -> helloworld.PlanetaCiudad
	2, // 1: helloworld.Broker.QuieroHacer:input_type -> helloworld.Comando
	0, // 2: helloworld.Fulcrum.PreguntarInformantes:input_type -> helloworld.PlanetaCiudad
	4, // 3: helloworld.Fulcrum.AddCity:input_type -> helloworld.Estructura
	5, // 4: helloworld.Fulcrum.UpdateName:input_type -> helloworld.Estructura2
	5, // 5: helloworld.Fulcrum.UpdateNumber:input_type -> helloworld.Estructura2
	4, // 6: helloworld.Fulcrum.DeleteCity:input_type -> helloworld.Estructura
	1, // 7: helloworld.Broker.GetNumberRebels:output_type -> helloworld.Numero
	3, // 8: helloworld.Broker.QuieroHacer:output_type -> helloworld.Redirigido
	1, // 9: helloworld.Fulcrum.PreguntarInformantes:output_type -> helloworld.Numero
	7, // 10: helloworld.Fulcrum.AddCity:output_type -> helloworld.Vector
	7, // 11: helloworld.Fulcrum.UpdateName:output_type -> helloworld.Vector
	7, // 12: helloworld.Fulcrum.UpdateNumber:output_type -> helloworld.Vector
	7, // 13: helloworld.Fulcrum.DeleteCity:output_type -> helloworld.Vector
	7, // [7:14] is the sub-list for method output_type
	0, // [0:7] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_helloworld_proto_init() }
func file_helloworld_proto_init() {
	if File_helloworld_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_helloworld_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlanetaCiudad); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_helloworld_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Numero); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_helloworld_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comando); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_helloworld_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Redirigido); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_helloworld_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Estructura); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_helloworld_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Estructura2); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_helloworld_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Estructura3); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_helloworld_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vector); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_helloworld_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_helloworld_proto_goTypes,
		DependencyIndexes: file_helloworld_proto_depIdxs,
		MessageInfos:      file_helloworld_proto_msgTypes,
	}.Build()
	File_helloworld_proto = out.File
	file_helloworld_proto_rawDesc = nil
	file_helloworld_proto_goTypes = nil
	file_helloworld_proto_depIdxs = nil
}
