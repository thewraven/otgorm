package otgorm

import (
	"fmt"

	"go.opentelemetry.io/otel/label"
)

//CreateSpanAttribute creates a KeyValue for use as a span attribute
func CreateSpanAttribute(k string, v interface{}) label.KeyValue {
	switch x := v.(type) {
	case string:
		return label.String(k, v.(string))
	case bool:
		return label.Bool(k, v.(bool))
	case int64:
		return label.Int64(k, v.(int64))
	case int32:
		return label.Int32(k, v.(int32))
	case int:
		return label.Int(k, v.(int))
	case float64:
		return label.Float64(k, v.(float64))
	case float32:
		return label.Float32(k, v.(float32))
	case uint:
		return label.Uint(k, v.(uint))
	case uint64:
		return label.Uint64(k, v.(uint64))
	case uint32:
		return label.Uint32(k, v.(uint32))
	default:
		return label.String("attribute.error", fmt.Sprintf("Couldn't convert %s into KeyValue", x))
	}
}
