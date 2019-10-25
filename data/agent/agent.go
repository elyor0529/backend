/* Parle - a free customer communication platform
 * Copyright (C) 2017 Focus Centric inc.
 *
 */

package agent

import (
	"fmt"
	"time"

	"github.com/parle-io/backend/data"
)

var bucketAgents = []byte("agents")

type Agent struct {
	ID       uint64    `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:-`
	Token    string    `json:"token"`
	Created  time.Time `json:"created"`
}

func (a *Agent) SetID(id uint64) {
	a.ID = id
}

func Add(a Agent) (*Agent, error) {
	if err := data.CreateWithAutoIncrement(bucketAgents, &a); err != nil {
		return nil, err
	}
	return &a, nil
}

func GetByEmail(email string) (*Agent, error) {
	var agent Agent
	agents, err := data.List(bucketAgents, &agent)
	if err != nil {
		return nil, err
	}

	for _, v := range agents {
		a, ok := v.(Agent)
		if !ok {
			return nil, err
		}

		if a.Email == email {
			return &a, nil
		}
	}

	return nil, nil
}

func validateToken(email, token string) (*Agent, error) {
	a, err := GetByEmail(email)
	if err != nil {
		return nil, err
	} else if a.Token != token {
		return nil, fmt.Errorf("invalid email/token")
	}

	return a, nil
}
