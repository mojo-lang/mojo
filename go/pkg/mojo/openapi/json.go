package openapi

import jsoniter "github.com/json-iterator/go"

// decodeJSONValue retains the iterator configuration and reports nested errors,
// which jsoniter.Any.ToVal otherwise discards.
func decodeJSONValue(iter *jsoniter.Iterator, value jsoniter.Any, target interface{}) bool {
	decoder := iter.Pool().BorrowIterator([]byte(value.ToString()))
	defer iter.Pool().ReturnIterator(decoder)
	decoder.ReadVal(target)
	if decoder.Error != nil {
		iter.ReportError("openapi JSON", decoder.Error.Error())
		return false
	}
	return true
}
