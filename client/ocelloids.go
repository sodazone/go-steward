// Copyright 2024 team@soda.zone
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

const HTTP_URL = "https://api.ocelloids.net"
const PUB_KEY = "eyJhbGciOiJFZERTQSIsImtpZCI6Im92SFVDU3hRM0NiYkJmc01STVh1aVdjQkNZcDVydmpvamphT2J4dUxxRDQ9In0.ewogICJpc3MiOiAiYXBpLm9jZWxsb2lkcy5uZXQiLAogICJqdGkiOiAiMDEwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAiLAogICJzdWIiOiAicHVibGljQG9jZWxsb2lkcyIKfQo.qKSfxo6QYGxzv40Ox7ec6kpt2aVywKmhpg6lue4jqmZyY6y3SwfT-DyX6Niv-ine5k23E0RKGQdm_MbtyPp9CA"

type OcelloidsClient struct {
	apiKey  string
	httpUrl string
	//wsUrl   string
	enc        *json.Encoder
	pagination Pagination
}

type PageInfo struct {
	EndCursor   string
	HasNextPage bool
}

type QueryResult struct {
	PageInfo PageInfo
	Items    []any
}

type Pagination struct {
	Cursor string `json:"cursor"`
	Limit  uint16 `json:"limit"`
}

type QueryArgs struct {
	Op string `json:"op"`
}

type Query struct {
	Pagination Pagination `json:"pagination"`
	Args       QueryArgs  `json:"args"`
}

func NewOcelloidsClient(apiKey string, httpUrl string, pagination Pagination) *OcelloidsClient {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")

	if apiKey == "" {
		apiKey = PUB_KEY
	}

	return &OcelloidsClient{
		apiKey,
		httpUrl,
		enc,
		pagination,
	}
}

func (client OcelloidsClient) FetchAssets() error {
	return client.execOp("assets.list")
}

func (client OcelloidsClient) FetchChains() error {
	return client.execOp("chains.list")
}

func (client OcelloidsClient) execOp(op string) error {
	return client.post(Query{
		Pagination: client.pagination,
		Args: QueryArgs{
			Op: op,
		},
	})
}

func (client OcelloidsClient) printAsJson(item any) error {
	return client.enc.Encode(item)
}

func (client OcelloidsClient) post(query Query) error {
	url := fmt.Sprintf("%s/query/steward", client.httpUrl)
	bearer := fmt.Sprintf("Bearer %s", client.apiKey)

	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(query)
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode == 429 {
		delay, err := strconv.Atoi(res.Header.Get("Retry-After"))
		if err == nil {
			time.Sleep(time.Duration(delay) * time.Second)
			return client.post(query)
		}
	}

	if res.StatusCode >= 400 {
		msg, _ := io.ReadAll(res.Body)
		return fmt.Errorf("error status=%d, body=%s", res.StatusCode, string(msg))
	}

	dec := json.NewDecoder(res.Body)
	var qres QueryResult
	if err := dec.Decode(&qres); err != nil {
		return err
	}

	if len(qres.Items) > 0 {
		for _, item := range qres.Items {
			if err := client.printAsJson(item); err != nil {
				return err
			}
		}

		if qres.PageInfo.HasNextPage {
			client.post(Query{
				Pagination: Pagination{
					Cursor: qres.PageInfo.EndCursor,
					Limit:  query.Pagination.Limit,
				},
				Args: query.Args,
			})
		}
	}

	return nil
}
