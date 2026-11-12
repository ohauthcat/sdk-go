package ohauthcat

import (
	"bytes"
	"context"
	"encoding/json"
	"iter"
)

type pageItem struct {
	Value    JSONValue
	Response rawResponse
}

func (c *Client) paginate(ctx context.Context, operation string, input requestInput) iter.Seq2[pageItem, error] {
	snapshot := input
	snapshot.Path = bytes.Clone(input.Path)
	snapshot.Query = bytes.Clone(input.Query)
	snapshot.Body = bytes.Clone(input.Body)
	return func(yield func(pageItem, error) bool) {
		if c == nil {
			yield(pageItem{}, requestFailure("client_closed"))
			return
		}
		queryData := snapshot.Query
		if queryData == nil {
			queryData = []byte("{}")
		}
		query, err := objectFields(queryData)
		if err != nil {
			yield(pageItem{}, requestFailure("request_invalid"))
			return
		}
		if _, exists := query["before"]; exists {
			yield(pageItem{}, requestFailure("invalid_pagination"))
			return
		}
		seen := make(map[string]bool)
		if raw, exists := query["after"]; exists {
			var after string
			if json.Unmarshal(raw, &after) != nil {
				yield(pageItem{}, requestFailure("request_invalid"))
				return
			}
			if after != "" {
				seen[after] = true
			}
		}
		args := snapshot
		for page := 0; page < c.config.MaxPages; page++ {
			args.Query, err = json.Marshal(query)
			if err != nil {
				yield(pageItem{}, requestFailure("request_invalid"))
				return
			}
			response, err := c.request(ctx, operation, args)
			if err != nil {
				yield(pageItem{}, err)
				return
			}
			var result struct {
				Data       []json.RawMessage `json:"data"`
				HasMore    bool              `json:"has_more"`
				NextCursor *string           `json:"next_cursor"`
			}
			if json.Unmarshal(response.Body, &result) != nil {
				yield(pageItem{}, responseFailure("response_invalid", response))
				return
			}
			cursor := ""
			if result.NextCursor != nil {
				cursor = *result.NextCursor
			}
			if result.HasMore && (cursor == "" || len(result.Data) == 0 || seen[cursor]) {
				yield(pageItem{}, responseFailure("invalid_pagination", response))
				return
			}
			for _, raw := range result.Data {
				value, err := ParseJSONValue(raw)
				if err != nil {
					yield(pageItem{}, responseFailure("response_invalid", response))
					return
				}
				if !yield(pageItem{Value: value, Response: response}, nil) {
					return
				}
			}
			if !result.HasMore {
				return
			}
			seen[cursor] = true
			query["after"], _ = json.Marshal(cursor)
		}
		yield(pageItem{}, requestFailure("pagination_limit"))
	}
}

// List iterates events forwards, without modifying the supplied query.
func (r *EventsResources) List(ctx context.Context, args ListEventsArgs) iter.Seq2[Event, error] {
	encoded, encodeErr := json.Marshal(args)
	return func(yield func(Event, error) bool) {
		var zero Event
		if r == nil || r.client == nil {
			yield(zero, requestFailure("client_closed"))
			return
		}
		if encodeErr != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		var copyArgs ListEventsArgs
		if json.Unmarshal(encoded, &copyArgs) != nil {
			yield(zero, requestFailure("request_invalid"))
			return
		}
		input := requestInput{}
		if copyArgs.Query != nil {
			var err error
			input.Query, err = json.Marshal(copyArgs.Query)
			if err != nil {
				yield(zero, requestFailure("request_invalid"))
				return
			}
		}
		for item, err := range r.client.paginate(ctx, "listEvents", input) {
			if err != nil {
				yield(zero, err)
				return
			}
			data, _ := item.Value.MarshalJSON()
			var event Event
			if json.Unmarshal(data, &event) != nil {
				yield(zero, responseFailure("response_invalid", item.Response))
				return
			}
			if !yield(event, nil) {
				return
			}
		}
	}
}

// CreateLink creates the typed portal link described by the pinned API.
func (r *PortalResources) CreateLink(ctx context.Context, args CreatePortalLinkArgs) (CreatePortalLinkResult, error) {
	return r.CreatePortalLink(ctx, args)
}
