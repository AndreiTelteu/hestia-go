module github.com/andreitelteu/hestia-go/packages/example-plugin

go 1.21

require github.com/gofiber/fiber/v2 v2.48.0

require github.com/gofiber/fiber v1.14.6 // indirect

require (
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/gofiber/utils v0.0.10 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/schema v1.1.0 // indirect
	github.com/klauspost/compress v1.16.3 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.48.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
)

require github.com/andreitelteu/hestia-go/common v1.0.0

replace github.com/andreitelteu/hestia-go/common => ../../common
