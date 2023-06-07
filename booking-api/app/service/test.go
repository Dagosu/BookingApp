package service

import (
	"regexp"
	"strings"
)

var (
	isUpperLetter = regexp.MustCompile(`^[A-Z]+$`).MatchString
)

func IsTimeField(field string) bool {
	timeSubfields := []string{
		"created_at", "updated_at", "deleted_at", "date_", "valid_", "operations.", "est_or_sct", "flight_time", "tow_time",
		"roll_off_", "checkin.start", "checkin.end", "start", "end", "datetime", "expires_at", "triggered_at", "act_", "sct_",
		"act_or_est_or_sct",
	}

	for _, sf := range timeSubfields {
		if strings.Contains(field, sf) {
			return true
		}
	}

	return false
}

func IsSliceField(field string) bool {
	return field == "rules_applied" || field == "routing" || field == "alerts" || field == "vias" || field == "boarding_bridges" ||
		field == "gpus" || field == "pcas"
}

func MapContainsKey(key string, mapValues map[string]interface{}) bool {
	for k := range mapValues {
		if strings.Contains(key, k) {
			return true
		}
	}

	return false
}
