package main

import (
	"embed"
)

// Assets contains project assets.
//
//go:embed app/build
var Assets embed.FS
