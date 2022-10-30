package user

const (
	Valid       byte = 0
	PlayBanned  byte = 1
	ChatBanned  byte = 2
	BothBanned  byte = 3
	LoginBanned byte = 4
)

// Returns whether the user is eligable to play based on the banned status and email verification.
func IsPlayEligible(verified uint8, bannedStatus uint8) bool {
	if bannedStatus == PlayBanned || bannedStatus == BothBanned || bannedStatus == LoginBanned {
		return false
	}

	if verified == 0b0111 {
		//Ceck for Email, Age and Name Verification
		return true
	}
	return false
}
