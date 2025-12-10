package appstore

import (
	"encoding/json"
	"fmt"
)

type AccountInfoOutput struct {
	Account Account
}

type AccountsInfoOutput struct {
	Accounts []Account
	Current  string
}

func (t *appstore) AccountInfo() (AccountInfoOutput, error) {
	data, err := t.keychain.Get(AccountKey)
	if err != nil {
		return AccountInfoOutput{}, fmt.Errorf("failed to get account: %w", err)
	}

	var acc Account

	err = json.Unmarshal(data, &acc)
	if err != nil {
		return AccountInfoOutput{}, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return AccountInfoOutput{
		Account: acc,
	}, nil
}

func (t *appstore) AccountsInfo() (AccountsInfoOutput, error) {
	data, err := t.keychain.Get(AccountStorageKey)
	if err != nil {
		return AccountsInfoOutput{}, fmt.Errorf("failed to get account storage: %w", err)
	}

	var storage AccountStorage

	err = json.Unmarshal(data, &storage)
	if err != nil {
		return AccountsInfoOutput{}, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return AccountsInfoOutput(storage), nil
}
