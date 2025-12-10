package appstore

import (
	"fmt"
)

func (t *appstore) Revoke() error {
	err := t.keychain.Remove(AccountKey)
	if err != nil {
		return fmt.Errorf("failed to remove account from keychain: %w", err)
	}

	return nil
}
