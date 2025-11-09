package core

import (
	"sync"

	jsoniter "github.com/json-iterator/go"
)

var registeredJsonEncoderTypesOnce = &sync.Once{}
var registeredJsonEncoderTypes map[string]jsoniter.ValEncoder

var registeredJsonEncoderTypeFieldsOnce = &sync.Once{}
var registeredJsonEncoderTypeFields map[string]jsoniter.ValEncoder

func RegisterJSONTypeDecoder(typ string, decoder jsoniter.ValDecoder) {
	jsoniter.RegisterTypeDecoder(typ, decoder)
}

func RegisterJSONTypeEncoder(typ string, encoder jsoniter.ValEncoder) {
	jsoniter.RegisterTypeEncoder(typ, encoder)

	registeredJsonEncoderTypesOnce.Do(func() {
		registeredJsonEncoderTypes = make(map[string]jsoniter.ValEncoder)
	})
	registeredJsonEncoderTypes[typ] = encoder
}

func RegisterJSONFieldDecoder(typ string, field string, decoder jsoniter.ValDecoder) {
	jsoniter.RegisterFieldDecoder(typ, field, decoder)
}

func RegisterJSONFieldEncoder(typ string, field string, encoder jsoniter.ValEncoder) {
	jsoniter.RegisterFieldEncoder(typ, field, encoder)

	registeredJsonEncoderTypeFieldsOnce.Do(func() {
		registeredJsonEncoderTypeFields = make(map[string]jsoniter.ValEncoder)
	})
	registeredJsonEncoderTypeFields[typ+"."+field] = encoder
}

func FormatJson(input []byte) ([]byte, error) {
	var obj interface{}
	if err := jsoniter.Unmarshal(input, &obj); err != nil {
		return nil, err
	}
	return jsoniter.MarshalIndent(obj, "", "    ")
}
