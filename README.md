# aliasing-uuids

Companion to a blog post: [Deriving Safe ID Types in Go](https://lukasschwab.me/blog/gen/deriving-safe-id-types-in-go.html).

## Usage

1. Select an implementation to test: comment/uncomment `widget` imports in [widget_test.go](widget_test.go).
2. Run the tests: `go test widget_test.go`.

| Approach (blog post)            | Package                            |
|:--------------------------------|:-----------------------------------|
| Defining a new `string` type    | [pkg/string_type](pkg/string_type) |
| Embedding `uuid.UUID`           | [pkg/uuid_embed](pkg/uuid_embed)   |
| Defining a new `uuid.UUID` type | [pkg/uuid_type](pkg/uuid_type)     |
| Aliasing `uuid.UUID`            | [pkg/uuid_alias](pkg/uuid_alias)   |
