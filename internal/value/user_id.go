package value

import (
	"fmt"

	"github.com/rs/xid"
)

type UserID struct {
	xid.ID
}

func ParseUserID(s string) (UserID, error) {
	id, err := xid.FromString(s)
	if err != nil {
		return UserID{}, fmt.Errorf("xid.FromString: %w", err)
	}
	return UserID{ID: id}, nil
}
