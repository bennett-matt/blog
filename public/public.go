package public

import (
	"embed"
)

//go:embed layouts partials views
var Files embed.FS
