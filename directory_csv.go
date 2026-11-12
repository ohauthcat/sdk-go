package ohauthcat

import (
	"context"
	"encoding/json"
)

// rawBodyMedia lists the operations whose request body is not JSON, with the
// media type sent as Content-Type. The pinned contract does not carry request
// media, so a test checks this map against the spec.
var rawBodyMedia = map[string]string{
	"confirmDirectorySourceOAuth":            "application/x-www-form-urlencoded",
	"importDirectoryCSV":                     "text/csv",
	"importPortalDirectoryCSV":               "text/csv",
	"removeFromWaitlistOneClick":             "application/x-www-form-urlencoded",
	"unsubscribeSubprocessorNoticesOneClick": "application/x-www-form-urlencoded",
}

// ImportDirectoryCSVOptions controls ImportDirectoryCSV.
type ImportDirectoryCSVOptions struct {
	// DryRun reports the changes without applying them.
	DryRun *bool
	// IdempotencyKey is sent as the Idempotency-Key header when set.
	IdempotencyKey string
}

// ImportDirectoryCSV replaces the users and groups of a directory whose
// provider is csv with the content of the file, sent as a text/csv body.
// The generated Directories.ImportDirectoryCsv cannot send this body.
func (c *Client) ImportDirectoryCSV(ctx context.Context, directory DirectoryID, csv []byte, options ImportDirectoryCSVOptions) (DirectorySyncResult, error) {
	var result DirectorySyncResult
	if c == nil {
		return result, requestFailure("client_closed")
	}
	if directory == "" || len(csv) == 0 {
		return result, requestFailure("request_invalid")
	}
	path, err := json.Marshal(ImportDirectoryCsvPath{Dir: directory})
	if err != nil {
		return result, requestFailure("request_invalid")
	}
	input := requestInput{Path: path, IdempotencyKey: options.IdempotencyKey, RawBody: csv, RawMedia: rawBodyMedia["importDirectoryCSV"]}
	if options.DryRun != nil {
		input.Query, err = json.Marshal(ImportDirectoryCsvQuery{DryRun: options.DryRun})
		if err != nil {
			return result, requestFailure("request_invalid")
		}
	}
	response, err := c.request(ctx, "importDirectoryCSV", input)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return result, responseFailure("response_invalid", response)
	}
	return result, nil
}
