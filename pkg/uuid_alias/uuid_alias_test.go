// uuid_alias_test demonstrates the shortcomings of type-alias IDs:
// this test represents desirable behavior, but it fails.
package uuid_alias_test

import (
	"reflect"
	"testing"

	widget "github.com/lukasschwab/aliasing-uuids/pkg/uuid_alias"

	"github.com/google/uuid"
	"github.com/peterldowns/testy/check"
)

func TestID_MultiAlias(t *testing.T) {
	type OrderID = uuid.UUID
	var orderID OrderID = uuid.Nil

	// Both of these assertions fail; these IDs are equal!
	check.NotEqual(t, reflect.TypeOf(orderID).Name(), reflect.TypeOf(widget.NilID).Name())
	check.NotEqual(t, orderID, widget.NilID)
}
