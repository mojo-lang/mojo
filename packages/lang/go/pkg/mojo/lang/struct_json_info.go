package lang

import (
	"reflect"
	"strings"
	"unicode"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

type StructFieldJsonInfo struct {
	ToName    string
	Ignored   bool
	OmitEmpty bool
	Encoder   jsoniter.ValEncoder
}

type StructJsonInfo struct {
	Config jsoniter.Config
	API    jsoniter.API
	Fields []StructFieldJsonInfo
}

func getTagKey(config jsoniter.Config) string {
	if config.TagKey == "" {
		return "json"
	}
	return config.TagKey
}

func getFieldName(originalFieldName string, tagProvidedFieldName string, wholeTag string) string {
	// ignore?
	if wholeTag == "-" {
		return ""
	}

	// private?
	isNotExported := unicode.IsLower(rune(originalFieldName[0])) || originalFieldName[0] == '_'
	if isNotExported {
		return ""
	}

	// rename?
	if tagProvidedFieldName == "" {
		return originalFieldName
	} else {
		return tagProvidedFieldName
	}
}

func GetStructJsonInfo(config jsoniter.Config, t reflect.Type) StructJsonInfo {
	structInfo := StructJsonInfo{
		Config: config,
		API:    config.Froze(),
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag, hasTag := field.Tag.Lookup(getTagKey(config))
		if config.OnlyTaggedField && !hasTag && !field.Anonymous {
			continue
		}
		if tag == "-" || field.Name == "_" {
			structInfo.Fields = append(structInfo.Fields, StructFieldJsonInfo{
				Ignored: true,
			})
			continue
		}
		tagParts := strings.Split(tag, ",")
		if field.Anonymous && (tag == "" || tagParts[0] == "") {
			// skip for anonymous field
		}

		fieldName := getFieldName(field.Name, tagParts[0], tag)
		// fieldCacheKey := fmt.Sprintf("%s/%s", t.String(), field.Name)

		ignored := false
		if len(fieldName) == 0 {
			ignored = true
		}

		omitEmpty := false
		if len(tagParts) > 1 && tagParts[1] == "omitempty" {
			omitEmpty = true
		}

		structInfo.Fields = append(structInfo.Fields, StructFieldJsonInfo{
			ToName:    fieldName,
			Ignored:   ignored,
			OmitEmpty: omitEmpty,
			Encoder:   structInfo.API.EncoderOf(reflect2.Type2(field.Type)),
		})
	}
	return structInfo
}

func GetStructJsonInfos(config jsoniter.Config, types ...interface{}) map[reflect.Type]StructJsonInfo {
	infos := make(map[reflect.Type]StructJsonInfo)

	for _, t := range types {
		info := GetStructJsonInfo(config, reflect.TypeOf(t))
		infos[reflect.TypeOf(t)] = info
	}

	return infos
}
