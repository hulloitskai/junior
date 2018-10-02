package main

import (
	"bytes"
	"os"
	"path/filepath"

	fhttp "github.com/valyala/fasthttp"
)

// DefaultFile is the default file of a directory that will be returned if
// the path has no file extension.
const DefaultFile = "index.html"

// HandleFastHTTP implements fasthttp.Handler for Config.
func (cfg *Config) HandleFastHTTP(ctx *fhttp.RequestCtx) {
	var (
		uri  = ctx.URI()
		path = uri.Path()
		lseg = uri.LastPathSegment()
	)

	// Perform trailing slash redirection if necessary.
	if cfg.TrailingSlashes {
		if bytes.IndexRune(lseg, '.') == -1 {
			path = append(path, '/')
			ctx.RedirectBytes(path, 301) // redirect with trailing slash
			return
		}
	} else if len(lseg) == 0 && len(path) > 1 { // path ends in a '/'
		path = path[:len(path)-1]    // pop last element
		ctx.RedirectBytes(path, 301) // redirect without trailing slash
		return
	}

	// Interpret path to target file.
	fpath := filepath.Join(cfg.RootDir, string(path))
	if len(lseg) == 0 || bytes.IndexRune(lseg, '.') == -1 {
		fpath = filepath.Join(fpath, DefaultFile)
	}

	// Return 404 if target file does not exist.
	if _, err := os.Stat(fpath); err != nil {
		if os.IsNotExist(err) {
			ctx.NotFound()

			if cfg.Has404 {
				ctx.SendFile(filepath.Join(cfg.RootDir, NFName))
				return
			}

			ctx.WriteString(NotFoundBody)
			return
		}
	}

	ctx.SendFile(fpath)
}
