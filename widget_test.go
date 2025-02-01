package widget_test

import (
	"encoding/json"
	"fmt"
	"testing"

	// widget "github.com/lukasschwab/aliasing-uuids/pkg/string_alias"
	// widget "github.com/lukasschwab/aliasing-uuids/pkg/uuid_embed"
	widget "github.com/lukasschwab/aliasing-uuids/pkg/uuid_alias"

	"github.com/peterldowns/testy/assert"
)

func TestID_Marshal(t *testing.T) {
	bytes, err := json.Marshal(widget.NilID)
	assert.Nil(t, err)
	assert.Equal(t, "\"00000000-0000-0000-0000-000000000000\"", string(bytes))
}

func TestID_Unmarshal(t *testing.T) {
	widgetID := new(widget.ID)
	err := json.Unmarshal([]byte("\"00000000-0000-0000-0000-000000000000\""), widgetID)
	assert.Nil(t, err)
	assert.Equal(t, widget.NilID, *widgetID)
}

func TestID_String(t *testing.T) {
	assert.Equal(t, "00000000-0000-0000-0000-000000000000", fmt.Sprint(widget.NilID))
}
