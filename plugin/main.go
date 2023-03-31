package main

import (
	"fmt"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	pb "github.com/kostyay/protoc-gen-go-private/private/v1"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, file := range gen.Files {
			if !file.Generate {
				continue
			}
			generateFile(file)
		}
		return nil
	})
}
func generateFile(file *protogen.File) {
	for _, message := range file.Messages {
		generateMessage(message)
	}
}

func generateMessage(message *protogen.Message) {

	for _, field := range message.Fields {
		if privateOption := getPrivateOption(field); privateOption != nil {
			fmt.Printf("Got one: %v", privateOption)
		}
	}

}

func getPrivateOption(field *protogen.Field) *pb.FieldOption {
	if !field.Desc.IsExtension() {
		return nil
	}

	return nil
}

func zeroValue(field *descriptor.FieldDescriptorProto) string {
	switch field.GetType() {
	case descriptor.FieldDescriptorProto_TYPE_BOOL,
		descriptor.FieldDescriptorProto_TYPE_ENUM:
		return "false"
	case descriptor.FieldDescriptorProto_TYPE_INT32,
		descriptor.FieldDescriptorProto_TYPE_UINT32,
		descriptor.FieldDescriptorProto_TYPE_FIXED32,
		descriptor.FieldDescriptorProto_TYPE_SINT32,
		descriptor.FieldDescriptorProto_TYPE_SFIXED32:
		return "0"
	case descriptor.FieldDescriptorProto_TYPE_INT64,
		descriptor.FieldDescriptorProto_TYPE_UINT64,
		descriptor.FieldDescriptorProto_TYPE_FIXED64,
		descriptor.FieldDescriptorProto_TYPE_SINT64,
		descriptor.FieldDescriptorProto_TYPE_SFIXED64:
		return "0"
	case descriptor.FieldDescriptorProto_TYPE_FLOAT,
		descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		return "0.0"
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		return "\"\""
	case descriptor.FieldDescriptorProto_TYPE_BYTES:
		return "[]byte{}"
	default:
		panic(fmt.Sprintf("unknown type: %v", field.GetType()))
	}
}
