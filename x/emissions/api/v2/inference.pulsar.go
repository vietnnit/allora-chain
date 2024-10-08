// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package emissionsv2

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_RegretInformedWeight        protoreflect.MessageDescriptor
	fd_RegretInformedWeight_worker protoreflect.FieldDescriptor
	fd_RegretInformedWeight_weight protoreflect.FieldDescriptor
)

func init() {
	file_emissions_v2_inference_proto_init()
	md_RegretInformedWeight = File_emissions_v2_inference_proto.Messages().ByName("RegretInformedWeight")
	fd_RegretInformedWeight_worker = md_RegretInformedWeight.Fields().ByName("worker")
	fd_RegretInformedWeight_weight = md_RegretInformedWeight.Fields().ByName("weight")
}

var _ protoreflect.Message = (*fastReflection_RegretInformedWeight)(nil)

type fastReflection_RegretInformedWeight RegretInformedWeight

func (x *RegretInformedWeight) ProtoReflect() protoreflect.Message {
	return (*fastReflection_RegretInformedWeight)(x)
}

func (x *RegretInformedWeight) slowProtoReflect() protoreflect.Message {
	mi := &file_emissions_v2_inference_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_RegretInformedWeight_messageType fastReflection_RegretInformedWeight_messageType
var _ protoreflect.MessageType = fastReflection_RegretInformedWeight_messageType{}

type fastReflection_RegretInformedWeight_messageType struct{}

func (x fastReflection_RegretInformedWeight_messageType) Zero() protoreflect.Message {
	return (*fastReflection_RegretInformedWeight)(nil)
}
func (x fastReflection_RegretInformedWeight_messageType) New() protoreflect.Message {
	return new(fastReflection_RegretInformedWeight)
}
func (x fastReflection_RegretInformedWeight_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_RegretInformedWeight
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_RegretInformedWeight) Descriptor() protoreflect.MessageDescriptor {
	return md_RegretInformedWeight
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_RegretInformedWeight) Type() protoreflect.MessageType {
	return _fastReflection_RegretInformedWeight_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_RegretInformedWeight) New() protoreflect.Message {
	return new(fastReflection_RegretInformedWeight)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_RegretInformedWeight) Interface() protoreflect.ProtoMessage {
	return (*RegretInformedWeight)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_RegretInformedWeight) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Worker != "" {
		value := protoreflect.ValueOfString(x.Worker)
		if !f(fd_RegretInformedWeight_worker, value) {
			return
		}
	}
	if x.Weight != "" {
		value := protoreflect.ValueOfString(x.Weight)
		if !f(fd_RegretInformedWeight_weight, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_RegretInformedWeight) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "emissions.v2.RegretInformedWeight.worker":
		return x.Worker != ""
	case "emissions.v2.RegretInformedWeight.weight":
		return x.Weight != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: emissions.v2.RegretInformedWeight"))
		}
		panic(fmt.Errorf("message emissions.v2.RegretInformedWeight does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_RegretInformedWeight) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "emissions.v2.RegretInformedWeight.worker":
		x.Worker = ""
	case "emissions.v2.RegretInformedWeight.weight":
		x.Weight = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: emissions.v2.RegretInformedWeight"))
		}
		panic(fmt.Errorf("message emissions.v2.RegretInformedWeight does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_RegretInformedWeight) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "emissions.v2.RegretInformedWeight.worker":
		value := x.Worker
		return protoreflect.ValueOfString(value)
	case "emissions.v2.RegretInformedWeight.weight":
		value := x.Weight
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: emissions.v2.RegretInformedWeight"))
		}
		panic(fmt.Errorf("message emissions.v2.RegretInformedWeight does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_RegretInformedWeight) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "emissions.v2.RegretInformedWeight.worker":
		x.Worker = value.Interface().(string)
	case "emissions.v2.RegretInformedWeight.weight":
		x.Weight = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: emissions.v2.RegretInformedWeight"))
		}
		panic(fmt.Errorf("message emissions.v2.RegretInformedWeight does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_RegretInformedWeight) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "emissions.v2.RegretInformedWeight.worker":
		panic(fmt.Errorf("field worker of message emissions.v2.RegretInformedWeight is not mutable"))
	case "emissions.v2.RegretInformedWeight.weight":
		panic(fmt.Errorf("field weight of message emissions.v2.RegretInformedWeight is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: emissions.v2.RegretInformedWeight"))
		}
		panic(fmt.Errorf("message emissions.v2.RegretInformedWeight does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_RegretInformedWeight) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "emissions.v2.RegretInformedWeight.worker":
		return protoreflect.ValueOfString("")
	case "emissions.v2.RegretInformedWeight.weight":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: emissions.v2.RegretInformedWeight"))
		}
		panic(fmt.Errorf("message emissions.v2.RegretInformedWeight does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_RegretInformedWeight) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in emissions.v2.RegretInformedWeight", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_RegretInformedWeight) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_RegretInformedWeight) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_RegretInformedWeight) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_RegretInformedWeight) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*RegretInformedWeight)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Worker)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Weight)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*RegretInformedWeight)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Weight) > 0 {
			i -= len(x.Weight)
			copy(dAtA[i:], x.Weight)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Weight)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Worker) > 0 {
			i -= len(x.Worker)
			copy(dAtA[i:], x.Worker)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Worker)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*RegretInformedWeight)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: RegretInformedWeight: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: RegretInformedWeight: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Worker", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Worker = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Weight = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: emissions/v2/inference.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegretInformedWeight struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Worker string `protobuf:"bytes,1,opt,name=worker,proto3" json:"worker,omitempty"` // worker who created the value
	Weight string `protobuf:"bytes,2,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (x *RegretInformedWeight) Reset() {
	*x = RegretInformedWeight{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emissions_v2_inference_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegretInformedWeight) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegretInformedWeight) ProtoMessage() {}

// Deprecated: Use RegretInformedWeight.ProtoReflect.Descriptor instead.
func (*RegretInformedWeight) Descriptor() ([]byte, []int) {
	return file_emissions_v2_inference_proto_rawDescGZIP(), []int{0}
}

func (x *RegretInformedWeight) GetWorker() string {
	if x != nil {
		return x.Worker
	}
	return ""
}

func (x *RegretInformedWeight) GetWeight() string {
	if x != nil {
		return x.Weight
	}
	return ""
}

var File_emissions_v2_inference_proto protoreflect.FileDescriptor

var file_emissions_v2_inference_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x65, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x69,
	0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x65, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x32, 0x1a, 0x18, 0x65, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x6e, 0x6f, 0x6e, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x01, 0x0a,
	0x14, 0x52, 0x65, 0x67, 0x72, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x64, 0x57,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x4f, 0x0a,
	0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x37, 0xc8,
	0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x6c, 0x6c, 0x6f, 0x72, 0x61, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x2f, 0x61, 0x6c, 0x6c, 0x6f, 0x72, 0x61, 0x2d, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x61,
	0x74, 0x68, 0x2e, 0x44, 0x65, 0x63, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x04,
	0xe8, 0xa0, 0x1f, 0x01, 0x42, 0xc4, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x32, 0x42, 0x0e, 0x49, 0x6e, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x6c, 0x6f, 0x72, 0x61, 0x2d, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x61, 0x6c, 0x6c, 0x6f, 0x72, 0x61, 0x2d, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x2f, 0x78, 0x2f, 0x65, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x65, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x32,
	0x3b, 0x65, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x76, 0x32, 0xa2, 0x02, 0x03, 0x45,
	0x58, 0x58, 0xaa, 0x02, 0x0c, 0x45, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x56,
	0x32, 0xca, 0x02, 0x0c, 0x45, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x5c, 0x56, 0x32,
	0xe2, 0x02, 0x18, 0x45, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x5c, 0x56, 0x32, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x45, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_emissions_v2_inference_proto_rawDescOnce sync.Once
	file_emissions_v2_inference_proto_rawDescData = file_emissions_v2_inference_proto_rawDesc
)

func file_emissions_v2_inference_proto_rawDescGZIP() []byte {
	file_emissions_v2_inference_proto_rawDescOnce.Do(func() {
		file_emissions_v2_inference_proto_rawDescData = protoimpl.X.CompressGZIP(file_emissions_v2_inference_proto_rawDescData)
	})
	return file_emissions_v2_inference_proto_rawDescData
}

var file_emissions_v2_inference_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_emissions_v2_inference_proto_goTypes = []interface{}{
	(*RegretInformedWeight)(nil), // 0: emissions.v2.RegretInformedWeight
}
var file_emissions_v2_inference_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_emissions_v2_inference_proto_init() }
func file_emissions_v2_inference_proto_init() {
	if File_emissions_v2_inference_proto != nil {
		return
	}
	file_emissions_v2_nonce_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_emissions_v2_inference_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegretInformedWeight); i {
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
			RawDescriptor: file_emissions_v2_inference_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_emissions_v2_inference_proto_goTypes,
		DependencyIndexes: file_emissions_v2_inference_proto_depIdxs,
		MessageInfos:      file_emissions_v2_inference_proto_msgTypes,
	}.Build()
	File_emissions_v2_inference_proto = out.File
	file_emissions_v2_inference_proto_rawDesc = nil
	file_emissions_v2_inference_proto_goTypes = nil
	file_emissions_v2_inference_proto_depIdxs = nil
}
