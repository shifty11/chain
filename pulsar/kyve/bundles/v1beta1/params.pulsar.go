// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package bundlesv1beta1

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
	md_Params                protoreflect.MessageDescriptor
	fd_Params_upload_timeout protoreflect.FieldDescriptor
	fd_Params_storage_cost   protoreflect.FieldDescriptor
	fd_Params_network_fee    protoreflect.FieldDescriptor
	fd_Params_max_points     protoreflect.FieldDescriptor
)

func init() {
	file_kyve_bundles_v1beta1_params_proto_init()
	md_Params = File_kyve_bundles_v1beta1_params_proto.Messages().ByName("Params")
	fd_Params_upload_timeout = md_Params.Fields().ByName("upload_timeout")
	fd_Params_storage_cost = md_Params.Fields().ByName("storage_cost")
	fd_Params_network_fee = md_Params.Fields().ByName("network_fee")
	fd_Params_max_points = md_Params.Fields().ByName("max_points")
}

var _ protoreflect.Message = (*fastReflection_Params)(nil)

type fastReflection_Params Params

func (x *Params) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Params)(x)
}

func (x *Params) slowProtoReflect() protoreflect.Message {
	mi := &file_kyve_bundles_v1beta1_params_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Params_messageType fastReflection_Params_messageType
var _ protoreflect.MessageType = fastReflection_Params_messageType{}

type fastReflection_Params_messageType struct{}

func (x fastReflection_Params_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Params)(nil)
}
func (x fastReflection_Params_messageType) New() protoreflect.Message {
	return new(fastReflection_Params)
}
func (x fastReflection_Params_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Params
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Params) Descriptor() protoreflect.MessageDescriptor {
	return md_Params
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Params) Type() protoreflect.MessageType {
	return _fastReflection_Params_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Params) New() protoreflect.Message {
	return new(fastReflection_Params)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Params) Interface() protoreflect.ProtoMessage {
	return (*Params)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Params) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.UploadTimeout != uint64(0) {
		value := protoreflect.ValueOfUint64(x.UploadTimeout)
		if !f(fd_Params_upload_timeout, value) {
			return
		}
	}
	if x.StorageCost != "" {
		value := protoreflect.ValueOfString(x.StorageCost)
		if !f(fd_Params_storage_cost, value) {
			return
		}
	}
	if x.NetworkFee != "" {
		value := protoreflect.ValueOfString(x.NetworkFee)
		if !f(fd_Params_network_fee, value) {
			return
		}
	}
	if x.MaxPoints != uint64(0) {
		value := protoreflect.ValueOfUint64(x.MaxPoints)
		if !f(fd_Params_max_points, value) {
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
func (x *fastReflection_Params) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "kyve.bundles.v1beta1.Params.upload_timeout":
		return x.UploadTimeout != uint64(0)
	case "kyve.bundles.v1beta1.Params.storage_cost":
		return x.StorageCost != ""
	case "kyve.bundles.v1beta1.Params.network_fee":
		return x.NetworkFee != ""
	case "kyve.bundles.v1beta1.Params.max_points":
		return x.MaxPoints != uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: kyve.bundles.v1beta1.Params"))
		}
		panic(fmt.Errorf("message kyve.bundles.v1beta1.Params does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "kyve.bundles.v1beta1.Params.upload_timeout":
		x.UploadTimeout = uint64(0)
	case "kyve.bundles.v1beta1.Params.storage_cost":
		x.StorageCost = ""
	case "kyve.bundles.v1beta1.Params.network_fee":
		x.NetworkFee = ""
	case "kyve.bundles.v1beta1.Params.max_points":
		x.MaxPoints = uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: kyve.bundles.v1beta1.Params"))
		}
		panic(fmt.Errorf("message kyve.bundles.v1beta1.Params does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Params) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "kyve.bundles.v1beta1.Params.upload_timeout":
		value := x.UploadTimeout
		return protoreflect.ValueOfUint64(value)
	case "kyve.bundles.v1beta1.Params.storage_cost":
		value := x.StorageCost
		return protoreflect.ValueOfString(value)
	case "kyve.bundles.v1beta1.Params.network_fee":
		value := x.NetworkFee
		return protoreflect.ValueOfString(value)
	case "kyve.bundles.v1beta1.Params.max_points":
		value := x.MaxPoints
		return protoreflect.ValueOfUint64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: kyve.bundles.v1beta1.Params"))
		}
		panic(fmt.Errorf("message kyve.bundles.v1beta1.Params does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_Params) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "kyve.bundles.v1beta1.Params.upload_timeout":
		x.UploadTimeout = value.Uint()
	case "kyve.bundles.v1beta1.Params.storage_cost":
		x.StorageCost = value.Interface().(string)
	case "kyve.bundles.v1beta1.Params.network_fee":
		x.NetworkFee = value.Interface().(string)
	case "kyve.bundles.v1beta1.Params.max_points":
		x.MaxPoints = value.Uint()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: kyve.bundles.v1beta1.Params"))
		}
		panic(fmt.Errorf("message kyve.bundles.v1beta1.Params does not contain field %s", fd.FullName()))
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
func (x *fastReflection_Params) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "kyve.bundles.v1beta1.Params.upload_timeout":
		panic(fmt.Errorf("field upload_timeout of message kyve.bundles.v1beta1.Params is not mutable"))
	case "kyve.bundles.v1beta1.Params.storage_cost":
		panic(fmt.Errorf("field storage_cost of message kyve.bundles.v1beta1.Params is not mutable"))
	case "kyve.bundles.v1beta1.Params.network_fee":
		panic(fmt.Errorf("field network_fee of message kyve.bundles.v1beta1.Params is not mutable"))
	case "kyve.bundles.v1beta1.Params.max_points":
		panic(fmt.Errorf("field max_points of message kyve.bundles.v1beta1.Params is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: kyve.bundles.v1beta1.Params"))
		}
		panic(fmt.Errorf("message kyve.bundles.v1beta1.Params does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Params) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "kyve.bundles.v1beta1.Params.upload_timeout":
		return protoreflect.ValueOfUint64(uint64(0))
	case "kyve.bundles.v1beta1.Params.storage_cost":
		return protoreflect.ValueOfString("")
	case "kyve.bundles.v1beta1.Params.network_fee":
		return protoreflect.ValueOfString("")
	case "kyve.bundles.v1beta1.Params.max_points":
		return protoreflect.ValueOfUint64(uint64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: kyve.bundles.v1beta1.Params"))
		}
		panic(fmt.Errorf("message kyve.bundles.v1beta1.Params does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Params) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in kyve.bundles.v1beta1.Params", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Params) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_Params) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Params) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Params)
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
		if x.UploadTimeout != 0 {
			n += 1 + runtime.Sov(uint64(x.UploadTimeout))
		}
		l = len(x.StorageCost)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.NetworkFee)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.MaxPoints != 0 {
			n += 1 + runtime.Sov(uint64(x.MaxPoints))
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
		x := input.Message.Interface().(*Params)
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
		if x.MaxPoints != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.MaxPoints))
			i--
			dAtA[i] = 0x20
		}
		if len(x.NetworkFee) > 0 {
			i -= len(x.NetworkFee)
			copy(dAtA[i:], x.NetworkFee)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.NetworkFee)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.StorageCost) > 0 {
			i -= len(x.StorageCost)
			copy(dAtA[i:], x.StorageCost)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.StorageCost)))
			i--
			dAtA[i] = 0x12
		}
		if x.UploadTimeout != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.UploadTimeout))
			i--
			dAtA[i] = 0x8
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
		x := input.Message.Interface().(*Params)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Params: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field UploadTimeout", wireType)
				}
				x.UploadTimeout = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.UploadTimeout |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field StorageCost", wireType)
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
				x.StorageCost = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field NetworkFee", wireType)
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
				x.NetworkFee = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MaxPoints", wireType)
				}
				x.MaxPoints = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.MaxPoints |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
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
// source: kyve/bundles/v1beta1/params.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Params defines the bundles module parameters.
type Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// upload_timeout ...
	UploadTimeout uint64 `protobuf:"varint,1,opt,name=upload_timeout,json=uploadTimeout,proto3" json:"upload_timeout,omitempty"`
	// storage_cost ...
	StorageCost string `protobuf:"bytes,2,opt,name=storage_cost,json=storageCost,proto3" json:"storage_cost,omitempty"`
	// network_fee ...
	NetworkFee string `protobuf:"bytes,3,opt,name=network_fee,json=networkFee,proto3" json:"network_fee,omitempty"`
	// max_points ...
	MaxPoints uint64 `protobuf:"varint,4,opt,name=max_points,json=maxPoints,proto3" json:"max_points,omitempty"`
}

func (x *Params) Reset() {
	*x = Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kyve_bundles_v1beta1_params_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Params) ProtoMessage() {}

// Deprecated: Use Params.ProtoReflect.Descriptor instead.
func (*Params) Descriptor() ([]byte, []int) {
	return file_kyve_bundles_v1beta1_params_proto_rawDescGZIP(), []int{0}
}

func (x *Params) GetUploadTimeout() uint64 {
	if x != nil {
		return x.UploadTimeout
	}
	return 0
}

func (x *Params) GetStorageCost() string {
	if x != nil {
		return x.StorageCost
	}
	return ""
}

func (x *Params) GetNetworkFee() string {
	if x != nil {
		return x.NetworkFee
	}
	return ""
}

func (x *Params) GetMaxPoints() uint64 {
	if x != nil {
		return x.MaxPoints
	}
	return 0
}

var File_kyve_bundles_v1beta1_params_proto protoreflect.FileDescriptor

var file_kyve_bundles_v1beta1_params_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6b, 0x79, 0x76, 0x65, 0x2f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6b, 0x79, 0x76, 0x65, 0x2e, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xf2, 0x01, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0d, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x12, 0x51, 0x0a, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x73,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2e, 0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f,
	0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x63, 0x52, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x43, 0x6f, 0x73, 0x74, 0x12, 0x4f, 0x0a, 0x0b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f,
	0x66, 0x65, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2e, 0xc8, 0xde, 0x1f, 0x00, 0xda,
	0xde, 0x1f, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x63, 0x52, 0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x46, 0x65, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x42, 0xe2, 0x01, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x6b, 0x79, 0x76,
	0x65, 0x2e, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x42, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4b, 0x59, 0x56,
	0x45, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x70,
	0x75, 0x6c, 0x73, 0x61, 0x72, 0x2f, 0x6b, 0x79, 0x76, 0x65, 0x2f, 0x62, 0x75, 0x6e, 0x64, 0x6c,
	0x65, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x62, 0x75, 0x6e, 0x64, 0x6c,
	0x65, 0x73, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x4b, 0x42, 0x58, 0xaa,
	0x02, 0x14, 0x4b, 0x79, 0x76, 0x65, 0x2e, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x2e, 0x56,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x14, 0x4b, 0x79, 0x76, 0x65, 0x5c, 0x42, 0x75,
	0x6e, 0x64, 0x6c, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x20,
	0x4b, 0x79, 0x76, 0x65, 0x5c, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x16, 0x4b, 0x79, 0x76, 0x65, 0x3a, 0x3a, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_kyve_bundles_v1beta1_params_proto_rawDescOnce sync.Once
	file_kyve_bundles_v1beta1_params_proto_rawDescData = file_kyve_bundles_v1beta1_params_proto_rawDesc
)

func file_kyve_bundles_v1beta1_params_proto_rawDescGZIP() []byte {
	file_kyve_bundles_v1beta1_params_proto_rawDescOnce.Do(func() {
		file_kyve_bundles_v1beta1_params_proto_rawDescData = protoimpl.X.CompressGZIP(file_kyve_bundles_v1beta1_params_proto_rawDescData)
	})
	return file_kyve_bundles_v1beta1_params_proto_rawDescData
}

var file_kyve_bundles_v1beta1_params_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_kyve_bundles_v1beta1_params_proto_goTypes = []interface{}{
	(*Params)(nil), // 0: kyve.bundles.v1beta1.Params
}
var file_kyve_bundles_v1beta1_params_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_kyve_bundles_v1beta1_params_proto_init() }
func file_kyve_bundles_v1beta1_params_proto_init() {
	if File_kyve_bundles_v1beta1_params_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kyve_bundles_v1beta1_params_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Params); i {
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
			RawDescriptor: file_kyve_bundles_v1beta1_params_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kyve_bundles_v1beta1_params_proto_goTypes,
		DependencyIndexes: file_kyve_bundles_v1beta1_params_proto_depIdxs,
		MessageInfos:      file_kyve_bundles_v1beta1_params_proto_msgTypes,
	}.Build()
	File_kyve_bundles_v1beta1_params_proto = out.File
	file_kyve_bundles_v1beta1_params_proto_rawDesc = nil
	file_kyve_bundles_v1beta1_params_proto_goTypes = nil
	file_kyve_bundles_v1beta1_params_proto_depIdxs = nil
}
