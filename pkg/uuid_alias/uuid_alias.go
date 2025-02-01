package uuid_alias

import (
	"encoding/json"

	"github.com/google/uuid"
)

type ID uuid.UUID

var NilID = ID(uuid.Nil)

// MarshalJSON implements json.Marshaler
func (id ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(uuid.UUID(id))
}

// UnmarshalJSON implements json.Unmarshaler
func (id *ID) UnmarshalJSON(data []byte) error {
	uuidPointer := (*uuid.UUID)(id)
	return json.Unmarshal(data, uuidPointer)
}

// String implements fmt.Stringer
func (id ID) String() string {
	return uuid.UUID(id).String()
}
