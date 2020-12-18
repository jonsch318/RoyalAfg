package events

import "strings"

//ValidateEventName validates if the live event name equals the required with case insensetivity
func ValidateEventName(expected, live string) bool {
	return strings.EqualFold(expected, strings.ToUpper(live))
}
