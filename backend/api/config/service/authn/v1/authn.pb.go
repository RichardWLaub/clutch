// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.17.3
// source: config/service/authn/v1/authn.proto

package authnv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type OIDC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Issuer       string   `protobuf:"bytes,1,opt,name=issuer,proto3" json:"issuer,omitempty"`
	ClientId     string   `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret string   `protobuf:"bytes,3,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	RedirectUrl  string   `protobuf:"bytes,4,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
	Scopes       []string `protobuf:"bytes,5,rep,name=scopes,proto3" json:"scopes,omitempty"`
	// The subject is mapped from the JWT token's email claim by default.
	// Set this field to the JWT token's claim name to override the subject.
	SubjectClaimNameOverride string `protobuf:"bytes,6,opt,name=subject_claim_name_override,json=subjectClaimNameOverride,proto3" json:"subject_claim_name_override,omitempty"`
	GroupsClaimNameOverride  string `protobuf:"bytes,7,opt,name=groups_claim_name_override,json=groupsClaimNameOverride,proto3" json:"groups_claim_name_override,omitempty"`
}

func (x *OIDC) Reset() {
	*x = OIDC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_service_authn_v1_authn_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OIDC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OIDC) ProtoMessage() {}

func (x *OIDC) ProtoReflect() protoreflect.Message {
	mi := &file_config_service_authn_v1_authn_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OIDC.ProtoReflect.Descriptor instead.
func (*OIDC) Descriptor() ([]byte, []int) {
	return file_config_service_authn_v1_authn_proto_rawDescGZIP(), []int{0}
}

func (x *OIDC) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *OIDC) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *OIDC) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

func (x *OIDC) GetRedirectUrl() string {
	if x != nil {
		return x.RedirectUrl
	}
	return ""
}

func (x *OIDC) GetScopes() []string {
	if x != nil {
		return x.Scopes
	}
	return nil
}

func (x *OIDC) GetSubjectClaimNameOverride() string {
	if x != nil {
		return x.SubjectClaimNameOverride
	}
	return ""
}

func (x *OIDC) GetGroupsClaimNameOverride() string {
	if x != nil {
		return x.GroupsClaimNameOverride
	}
	return ""
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Used to sign the nonce or any other JWT secrets.
	SessionSecret string `protobuf:"bytes,1,opt,name=session_secret,json=sessionSecret,proto3" json:"session_secret,omitempty"`
	// Types that are assignable to Type:
	//
	//	*Config_Oidc
	Type isConfig_Type `protobuf_oneof:"type"`
	// Whether to permit service tokens to be issued. In addition to setting this flag
	// a token store must be configured.
	EnableServiceTokenCreation bool `protobuf:"varint,3,opt,name=enable_service_token_creation,json=enableServiceTokenCreation,proto3" json:"enable_service_token_creation,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_service_authn_v1_authn_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_config_service_authn_v1_authn_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_config_service_authn_v1_authn_proto_rawDescGZIP(), []int{1}
}

func (x *Config) GetSessionSecret() string {
	if x != nil {
		return x.SessionSecret
	}
	return ""
}

func (m *Config) GetType() isConfig_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Config) GetOidc() *OIDC {
	if x, ok := x.GetType().(*Config_Oidc); ok {
		return x.Oidc
	}
	return nil
}

func (x *Config) GetEnableServiceTokenCreation() bool {
	if x != nil {
		return x.EnableServiceTokenCreation
	}
	return false
}

type isConfig_Type interface {
	isConfig_Type()
}

type Config_Oidc struct {
	Oidc *OIDC `protobuf:"bytes,2,opt,name=oidc,proto3,oneof"`
}

func (*Config_Oidc) isConfig_Type() {}

var File_config_service_authn_v1_authn_proto protoreflect.FileDescriptor

var file_config_service_authn_v1_authn_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb,
	0x02, 0x0a, 0x04, 0x4f, 0x49, 0x44, 0x43, 0x12, 0x1f, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01,
	0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x20, 0x01, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2c,
	0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x0c,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x2a, 0x0a, 0x0c,
	0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x0b, 0x72, 0x65, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73,
	0x12, 0x3d, 0x0a, 0x1b, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x63, 0x6c, 0x61, 0x69,
	0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6c,
	0x61, 0x69, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x12,
	0x3b, 0x0a, 0x1a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x5f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x17, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x43, 0x6c, 0x61, 0x69, 0x6d,
	0x4e, 0x61, 0x6d, 0x65, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x22, 0xbf, 0x01, 0x0a,
	0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2e, 0x0a, 0x0e, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x0d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x3a, 0x0a, 0x04, 0x6f, 0x69, 0x64, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x49, 0x44, 0x43, 0x48, 0x00, 0x52, 0x04, 0x6f,
	0x69, 0x64, 0x63, 0x12, 0x41, 0x0a, 0x1d, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1a, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x42, 0x44,
	0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x79, 0x66,
	0x74, 0x2f, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x75, 0x74,
	0x68, 0x6e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_service_authn_v1_authn_proto_rawDescOnce sync.Once
	file_config_service_authn_v1_authn_proto_rawDescData = file_config_service_authn_v1_authn_proto_rawDesc
)

func file_config_service_authn_v1_authn_proto_rawDescGZIP() []byte {
	file_config_service_authn_v1_authn_proto_rawDescOnce.Do(func() {
		file_config_service_authn_v1_authn_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_service_authn_v1_authn_proto_rawDescData)
	})
	return file_config_service_authn_v1_authn_proto_rawDescData
}

var file_config_service_authn_v1_authn_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_config_service_authn_v1_authn_proto_goTypes = []interface{}{
	(*OIDC)(nil),   // 0: clutch.config.service.authn.v1.OIDC
	(*Config)(nil), // 1: clutch.config.service.authn.v1.Config
}
var file_config_service_authn_v1_authn_proto_depIdxs = []int32{
	0, // 0: clutch.config.service.authn.v1.Config.oidc:type_name -> clutch.config.service.authn.v1.OIDC
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_config_service_authn_v1_authn_proto_init() }
func file_config_service_authn_v1_authn_proto_init() {
	if File_config_service_authn_v1_authn_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_service_authn_v1_authn_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OIDC); i {
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
		file_config_service_authn_v1_authn_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
	file_config_service_authn_v1_authn_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Config_Oidc)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_config_service_authn_v1_authn_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_service_authn_v1_authn_proto_goTypes,
		DependencyIndexes: file_config_service_authn_v1_authn_proto_depIdxs,
		MessageInfos:      file_config_service_authn_v1_authn_proto_msgTypes,
	}.Build()
	File_config_service_authn_v1_authn_proto = out.File
	file_config_service_authn_v1_authn_proto_rawDesc = nil
	file_config_service_authn_v1_authn_proto_goTypes = nil
	file_config_service_authn_v1_authn_proto_depIdxs = nil
}
