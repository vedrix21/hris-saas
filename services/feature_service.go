package services

import (
	"encoding/json"
	"hris/models"
)

func HasFeature(account models.Account, feature string) bool {

	var features map[string]bool
	json.Unmarshal([]byte(account.Features), &features)

	return features[feature]
}