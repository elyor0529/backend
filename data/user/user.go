/* Parle - a free customer communication platform
 * Copyright (C) 2017 Focus Centric inc.
 *
 */

package user

import "time"

type User struct {
	Lead
	YourID         string            `json:"yourId"`
	FirstName      string            `json:"fname"`
	LastName       string            `json:"lname"`
	Avatar         string            `json:"avatar"`
	Tags           []string          `json:"tags"`
	Attributes     map[string]string `json:"customTags"`
	CompanyID      uint64            `json:"companyId"`
	CompanyName    string            `json:"companyName"`
	IsUnsubscribed bool              `json:"unsubscribed"`
}

type Tracking struct {
	Referrer string    `json:"referrer"`
	Source   string    `json:"utmSource"`
	Medium   string    `json:"utmMedium"`
	Campaign string    `json:"utmCampaign"`
	Terms    string    `json:"utmTerms"`
	Content  string    `json:"utmContent"`
	Tracked  time.Time `json:"tracked"`
}

func Add(u User) (*User, error) {
	return nil, nil
}

func updateUser(id uint64, u User) {

}
