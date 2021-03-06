// Copyright 2018 Oath, Inc.
// Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

package api

import (
	"github.com/yahoo/yfuzz/pkg/schema"
)

// GetServerVersion retrieves the version of the YFuzz server.
func GetServerVersion() (string, error) {
	var response schema.GetVersionResponse

	err := get("/version", &response)
	if err != nil {
		return "", err
	}

	return response.Version, nil
}
