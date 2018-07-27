// Copyright (c) 2018, Mark "Happy-Ferret" Bauermeister
//
// This software may be modified and distributed under the terms
// of the BSD license.  See the LICENSE file for details.

package xclient

import (
	"strings"
)

func (c *Client) matchWmClass(haystack []string) bool {
	instance := strings.ToLower(c.class.Instance)
	class := strings.ToLower(c.class.Class)
	for _, s := range haystack {
		if s == instance || s == class {
			return true
		}
	}
	return false
}
