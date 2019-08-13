/* Parle - a free customer communication platform
 * Copyright (C) 2017 Focus Centric inc.
 *
 */

package user

type Lead struct {
	Visitor
	Email string `json:"email"`
}

func updateLead(id uint64, l Lead) {
}
