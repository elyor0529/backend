/* Parle - a free customer communication platform
 * Copyright (C) 2017 Focus Centric inc.
 *
 */

package user

import "time"

// Company represents a company
type Company struct {
	ID           uint64            `json:"id"`
	YourID       string            `json:"yourId"`
	Name         string            `json:"name"`
	Website      string            `json:"website"`
	UserCount    int               `json:"userCount"`
	Attributes   map[string]string `json:"attributes"`
	Created      time.Time         `json:"created"`
	Updated      time.Time         `json:"updated"`
	LastActivity time.Time         `json:"lastActivity"`
}
