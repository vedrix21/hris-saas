package services

import (
	"encoding/json"
)

// 🔥 ambil semua fitur tenant
func GetFeatures(tenant string) (map[string]bool, error) {
	account, err := GetAccountByCode(tenant)
	if err != nil {
		return nil, err
	}

	var features map[string]bool
	err = json.Unmarshal([]byte(account.Features), &features)

	return features, err
}

// 🔥 cek 1 fitur
func HasFeature(tenant string, feature string) bool {
	features, err := GetFeatures(tenant)
	if err != nil {
		return false
	}

	return features[feature]
}
