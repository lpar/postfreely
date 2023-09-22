/*
 * Copyright © 2020-2021 Musing Studio LLC.
 *
 * This file is part of WriteFreely.
 *
 * WriteFreely is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License, included
 * in the LICENSE file in this source code package.
 */

package writefreely_test

import (
	"testing"

	"github.com/guregu/null/zero"
	"github.com/stretchr/testify/assert"
	"github.com/postfreely/postfreely"
)

func TestPostSummary(t *testing.T) {
	testCases := map[string]struct {
		given    writefreely.Post
		expected string
	}{
		"no special chars":          {givenPost("Content."), "Content."},
		"HTML content":              {givenPost("Content <p>with a</p> paragraph."), "Content with a paragraph."},
		"content with escaped char": {givenPost("Content&#39;s all OK."), "Content's all OK."},
		"multiline content": {givenPost(`Content
in
multiple
lines.`), "Content in multiple lines."},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := test.given.Summary()
			assert.Equal(t, test.expected, actual)
		})
	}
}

func givenPost(content string) writefreely.Post {
	return writefreely.Post{Title: zero.StringFrom("Title"), Content: content}
}
