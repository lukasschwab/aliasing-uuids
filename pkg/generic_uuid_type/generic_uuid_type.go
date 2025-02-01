package generic_uuid_type

import (
	"encoding/json"

	"github.com/google/uuid"
)

type ID[T any] uuid.UUID

// MarshalJSON implements json.Marshaler
func (id ID[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(uuid.UUID(id))
}

// UnmarshalJSON implements json.Unmarshaler
func (id *ID[T]) UnmarshalJSON(data []byte) error {
	uuidPointer := (*uuid.UUID)(id)
	return json.Unmarshal(data, uuidPointer)
}

// String implements fmt.Stringer
func (id ID[T]) String() string {
	return uuid.UUID(id).String()
}

// Widget with an ID, for example.
type Widget struct {
	ID ID[Widget]
}

func NilID[T any]() ID[T] {
	return ID[T](uuid.Nil)
}
