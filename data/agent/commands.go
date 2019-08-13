/* Parle - a free customer communication platform
 * Copyright (C) 2017 Focus Centric inc.
 *
 */

package agent

import (
	"fmt"
	"strings"
)

func Auth(msg string) (*Agent, error) {
	if strings.Index(msg, "|") == -1 {
		return nil, fmt.Errorf("invalid token format")
	}

	pairs := strings.Split(msg, "|")
	if len(pairs) != 2 {
		return nil, fmt.Errorf("invalid token format")
	}

	return validateToken(pairs[0], pairs[1])
}
