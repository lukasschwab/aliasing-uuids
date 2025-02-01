package generic_uuid_type_test

import (
	"reflect"
	"testing"

	generic "github.com/lukasschwab/aliasing-uuids/pkg/generic_uuid_type"

	"github.com/peterldowns/testy/assert"
)

func TestID_Generic(t *testing.T) {
	type Order interface{}
	orderID := generic.NilID[Order]()
	widgetID := generic.NilID[generic.Widget]()

	// Not comparable!
	assert.NotEqual(t, reflect.TypeOf(orderID).Name(), reflect.TypeOf(widgetID).Name())
	// Won't compile, as desired:
	// assert.NotEqual(t, orderID, widgetID)
}
