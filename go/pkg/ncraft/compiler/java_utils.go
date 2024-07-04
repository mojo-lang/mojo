package compiler

import "strings"

func JavaRequestMappingName(methodName string) string {
	switch methodName {
	case "get":
		return "GetMapping"
	case "post":
		return "PostMapping"
	case "put":
		return "PutMapping"
	case "delete":
		return "DeleteMapping"
	default:
		return ""
	}
}

func JavaTypeName(name string) string {
	switch name {
	case "Int8":
		return "byte"
	case "Int16":
		return "short"
	case "Int", "Int32":
		return "int"
	case "Int64":
		return "long"
	case "Bool":
		return "boolean"
	case "Float", "Float32":
		return "float"
	case "Double", "Float64":
		return "double"
	case "StringValue":
		return "String"
	case "Ordering", "FieldMask":
		return "String"
	case "Object":
		return "Map<String, Object>"
	default:
		return name
	}
}

func JavaGRpcTypeName(name string) string {
	switch name {
	case "Int8":
		return "byte"
	case "Int16":
		return "short"
	case "Int", "Int32":
		return "int"
	case "Int64":
		return "long"
	case "Bool":
		return "boolean"
	case "Float", "Float32":
		return "float"
	case "Double", "Float64":
		return "double"
	case "Object":
		return "org.mojolang.mojo.core.Object"
	default:
		return name
	}
}

func JavaHttp2GRpcConvert(name string) string {
	switch name {
	case "FieldMask":
		return "FieldMasks.fromString"
	case "Ordering":
		return "Orderings.fromString"
	case "Object":
		return "Objects.of"
	}
	return ""
}

func JavaGRpc2HttpConvert(name string) string {
	switch name {
	case "FieldMask":
		return "FieldMasks.toString"
	case "Ordering":
		return "Orderings.toString"
	case "Object":
		return "Objects.to"
	}
	return ""
}

func JavaNeedConvert(name string) bool {
	switch name {
	case "FieldMask", "Ordering", "Object":
		return true
	}
	return false
}

func GetJavaPackageName(org string, name string) string {
	segments := strings.Split(org, ".")
	reverse := func(s []string) {
		for i := 0; i < len(s)/2; i++ {
			j := len(s) - i - 1
			s[i], s[j] = s[j], s[i]
		}
	}
	reverse(segments)
	return strings.Join(segments, ".") + "." + name
}
