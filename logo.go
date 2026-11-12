package ohauthcat

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"io"
	"os"
	"unicode/utf8"
)

// MaxLogoBytes is the largest logo the API accepts (256 KiB).
const MaxLogoBytes = 256 * 1024

// Logo content types recognized from a file's bytes.
const (
	LogoTypePNG = "image/png"
	LogoTypeSVG = "image/svg+xml"
)

var pngSignature = []byte("\x89PNG\r\n\x1a\n")

// UploadLogoOptions carries optional request settings for logo uploads.
type UploadLogoOptions struct {
	IdempotencyKey string
}

func logoRefused(message string) *APIError {
	return &APIError{Code: CodeLogoInvalid, Type: "invalid_request", Param: "data", Message: message}
}

func trimXMLSpace(rest []byte) []byte {
	return bytes.TrimLeft(rest, " \t\r\n")
}

func hasSVGRoot(data []byte) bool {
	if !utf8.Valid(data) {
		return false
	}
	rest := bytes.TrimPrefix(data, []byte("\xef\xbb\xbf"))
	for {
		rest = trimXMLSpace(rest)
		switch {
		case bytes.HasPrefix(rest, []byte("<?")):
			end := bytes.Index(rest, []byte("?>"))
			if end < 0 {
				return false
			}
			rest = rest[end+2:]
		case bytes.HasPrefix(rest, []byte("<!--")):
			end := bytes.Index(rest[4:], []byte("-->"))
			if end < 0 {
				return false
			}
			rest = rest[4+end+3:]
		default:
			if !bytes.HasPrefix(rest, []byte("<svg")) || len(rest) == 4 {
				return false
			}
			return bytes.IndexByte([]byte(" \t\r\n/>"), rest[4]) >= 0
		}
	}
}

// CheckLogo checks a logo before upload: at most 256 KiB and a PNG signature or
// an SVG root element. It returns the detected type, or an *APIError with code
// logo_invalid and status 0. The API still validates dimensions and SVG content.
func CheckLogo(data []byte) (string, error) {
	switch {
	case len(data) == 0:
		return "", logoRefused("The logo is empty")
	case len(data) > MaxLogoBytes:
		return "", logoRefused("The logo is larger than 256 KiB")
	case bytes.HasPrefix(data, pngSignature):
		return LogoTypePNG, nil
	case hasSVGRoot(data):
		return LogoTypeSVG, nil
	}
	return "", logoRefused("The logo must be a PNG or an SVG")
}

// UploadLogo checks, base64-encodes and uploads logo bytes through
// Branding.UploadBrandingLogo. Refused bytes never reach the network.
func (c *Client) UploadLogo(ctx context.Context, data []byte, options UploadLogoOptions) (BrandingLogo, error) {
	if c == nil {
		return BrandingLogo{}, requestFailure("client_closed")
	}
	if _, err := CheckLogo(data); err != nil {
		return BrandingLogo{}, err
	}
	return c.Branding.UploadBrandingLogo(ctx, UploadBrandingLogoArgs{
		IdempotencyKey: options.IdempotencyKey,
		Body:           UploadBrandingLogoBody{Data: base64.StdEncoding.EncodeToString(data)},
	})
}

// UploadLogoFile reads a PNG or SVG file of at most 256 KiB and uploads it like UploadLogo.
func (c *Client) UploadLogoFile(ctx context.Context, path string, options UploadLogoOptions) (BrandingLogo, error) {
	if path == "" {
		return BrandingLogo{}, logoRefused("The logo path is empty")
	}
	data, err := readLogoFile(path)
	if err != nil {
		return BrandingLogo{}, err
	}
	return c.UploadLogo(ctx, data, options)
}

func readLogoFile(path string) ([]byte, error) {
	file, err := os.Open(path) //nolint:gosec // The caller chooses which local logo file to upload.
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, logoRefused("The logo path is not a file")
		}
		return nil, err
	}
	defer func() { _ = file.Close() }()
	info, err := file.Stat()
	if err != nil {
		return nil, err
	}
	if !info.Mode().IsRegular() {
		return nil, logoRefused("The logo path is not a file")
	}
	data, err := io.ReadAll(io.LimitReader(file, MaxLogoBytes+1))
	if err != nil {
		return nil, err
	}
	if len(data) > MaxLogoBytes {
		return nil, logoRefused("The logo is larger than 256 KiB")
	}
	return data, nil
}
