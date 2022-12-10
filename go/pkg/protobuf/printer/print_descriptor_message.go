package printer

import (
	"bytes"
	"fmt"

	"github.com/mojo-lang/core/go/pkg/mojo"
	"github.com/mojo-lang/protobuf/go/pkg/mojo/protobuf/descriptor"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/mojo-lang/mojo/go/pkg/context"
)

func (p *Printer) PrintDescriptorMessage(ctx context.Context, message *descriptor.Message) *Printer {
	p.PrintLine("message ", message.GetName(), " {").BreakLine()
	p.Indent()

	firstItem := true

	// Build a structure more suitable for generating the text in one pass
	for _, enum := range message.Enums {
		if !firstItem {
			p.PrintBlankLine()
		} else {
			firstItem = false
		}

		p.PrintDescriptorEnum(ctx, enum)
	}

	for _, msg := range message.Messages {
		if msg.IsMapEntry() {
			continue
		}

		if !firstItem {
			p.PrintBlankLine()
		} else {
			firstItem = false
		}

		p.PrintDescriptorMessage(ctx, msg)
	}

	printField := func(field *descriptor.Field) {
		if field.IsMessageType() {
			desc := message.GetMessage(field.GetTypeName())
			if desc.IsMapEntry() {
				keyType := desc.Fields[0].GetTypeName()
				valType := desc.Fields[1].GetTypeName()
				p.PrintRaw("map<", keyType, ", ", valType, "> ", field.GetName(), " = ", field.GetNumber())
			} else {
				if field.IsRepeated() {
					p.PrintRaw("repeated ", field.GetTypeName(), " ", field.GetName(), " = ", field.GetNumber())
				} else {
					p.PrintRaw(field.GetTypeName(), " ", field.GetName(), " = ", field.GetNumber())
				}
			}
		} else if field.IsRepeated() {
			p.PrintRaw("repeated ", field.GetTypeName(), " ", field.GetName(), " = ", field.GetNumber())
		} else {
			p.PrintRaw(field.GetTypeName(), " ", field.GetName(), " = ", field.GetNumber())
		}

		if field.HasOptions() {
			fieldOptions := mojo.FieldOptionsExtensions()
			first := true
			buffer := bytes.NewBuffer(nil)
			for _, option := range fieldOptions {
				if !field.HasOption(option) {
					continue
				}

				switch option.TypeDescriptor().Kind() {
				case protoreflect.BoolKind:
					v := field.GetBoolOption(option)
					if !first {
						buffer.WriteString(", ")
					}
					if v {
						buffer.WriteString(fmt.Sprint("(", mojo.GetOptionFullName(option), ")=true"))
					} else {
						buffer.WriteString(fmt.Sprint("(", mojo.GetOptionFullName(option), ")=false"))
					}
					first = false
				case protoreflect.Int32Kind, protoreflect.Int64Kind,
					protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind,
					protoreflect.Fixed32Kind, protoreflect.Fixed64Kind:
					v := field.GetOption(option).(int64)
					if !first {
						buffer.WriteString(", ")
					}
					buffer.WriteString(fmt.Sprint("(", mojo.GetOptionFullName(option), ")=", v))
					first = false
				case protoreflect.StringKind:
					v := field.GetStringOption(option)
					if !first {
						buffer.WriteString(", ")
					}
					buffer.WriteString(fmt.Sprint("(", mojo.GetOptionFullName(option), ")=", "\"", v, "\""))
					first = false
				}
			}
			if buffer.Len() > 0 {
				p.PrintRaw(" [", buffer.String(), "];")
			} else {
				p.PrintRaw(";")
			}
		} else {
			p.PrintRaw(";")
		}

		// fieldDeprecated := ""
		// if field.GetOptions().IsDeprecated() {
		//	fieldDeprecated = deprecationComment
		// }

		// fieldFullPath := fmt.Sprintf("%s,%d,%d", message.path, messageFieldPath, i)
		// c, ok := p.makeComments(fieldFullPath)
		// if ok {
		//	c += "\n"
		// }
	}

	if !firstItem && len(message.Fields) > 0 {
		p.PrintBlankLine()
	}

	printedOneofs := make(map[string]bool)
	for _, field := range message.Fields {
		if field.Oneof == nil {
			p.PrintLine()
			printField(field)
		} else {
			if _, ok := printedOneofs[field.Oneof.GetName()]; !ok {
				p.PrintLine("oneof ", field.Oneof.GetName(), " {")
				p.Indent()
				fields := field.Oneof.Fields
				for _, f := range fields {
					p.PrintLine()
					printField(f)
				}
				p.Outdent()
				p.PrintLine("}")

				printedOneofs[field.Oneof.GetName()] = true
			}
		}
	}

	p.Outdent()
	p.PrintLine("}")

	return p
}
