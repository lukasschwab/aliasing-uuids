package uuid_alias

import (
	"github.com/google/uuid"
)

type ID = uuid.UUID

var NilID ID = uuid.Nil
