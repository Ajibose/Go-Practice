package enums

type UserRole int

const (
	Guest UserRole = iota
	Member
	Moderator
	Admin
)

func (r UserRole) String() string {
	stateName := map[UserRole]string{
		Guest:     "Role: Guest",
		Member:    "Role: Member",
		Moderator: "Role: Moderator",
		Admin:     "Role: Administrator",
	}

	return stateName[r]
}
