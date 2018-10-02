package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	fhttp "github.com/valyala/fasthttp"
)

// HandleFastHTTP implements fasthttp.Handler for Config.
func (cfg *Config) HandleFastHTTP(ctx *fhttp.RequestCtx) {
	var (
		uri  = ctx.URI()
		path = uri.Path()
		lseg = uri.LastPathSegment()
	)

	// Perform trailing slash redirection if necessary.
	if len(path) > 1 {
		switch cfg.TrailingSlash {
		case "true", "1":
			if len(lseg) > 0 && bytes.IndexRune(lseg, '.') == -1 {
				path = append(path, '/')
				ctx.RedirectBytes(path, 301) // redirect with trailing slash
				return
			}
		case "false", "0":
			if len(lseg) == 0 { // path ends in a '/'
				path = path[:len(path)-1]    // pop last element
				ctx.RedirectBytes(path, 301) // redirect without trailing slash
				return
			}
		}
	}

	// Interpret path to target file.
	fpath := filepath.Join(cfg.RootDir, string(path))
	if len(lseg) == 0 || bytes.IndexRune(lseg, '.') == -1 {
		fpath = filepath.Join(fpath, DefaultFile)
	}

	// Return 404 if target file does not exist.
	if _, err := os.Stat(fpath); err != nil {
		if os.IsNotExist(err) {
			ctx.SetStatusCode(fhttp.StatusNotFound)

			if _, err := os.Stat(cfg.NotFound); err != nil {
				if os.IsNotExist(err) {
					ctx.SetContentType(NotFoundType)
					ctx.WriteString(NotFoundBody)
					return
				}

				ctx.Response.SetStatusCode(fhttp.StatusInternalServerError)
				fmt.Fprintf(ctx, "Failed to check 404 file: %v", err)
				return
			}

			ctx.SendFile(cfg.NotFound)
			return
		}

		ctx.Response.SetStatusCode(fhttp.StatusInternalServerError)
		fmt.Fprintf(ctx, "Failed to check file: %v", err)
		return
	}

	ctx.SendFile(fpath)
}
