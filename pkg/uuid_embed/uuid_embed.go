package uuid_embed

import "github.com/google/uuid"

type ID struct{ uuid.UUID }

var NilID = ID{uuid.Nil}
