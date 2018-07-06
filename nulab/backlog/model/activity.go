package model

const (
	TypeKindUnknown = "unknown"
	TypeKindIssue   = "issue"
	TypeKindPR      = "PR"
	TypeKindVCS     = "VCS"
	TypeKindWiki    = "Wiki"
	TypeKindFile    = "File"

	TypeValueUnknown = iota
	TypeValueCreate
	TypeValueUpdate
	TypeValueDelete
	TypeValueAction
)

type (
	// Activity represents activity.
	Activity struct {
		Created
		ID            int            `json:"id"`
		Project       Project        `json:"project"`
		Type          Type           `json:"type"`
		Content       Content        `json:"content"`
		Notifications []Notification `json:"notifications"`
		User          User           `json:"createdUser"`
	}

	Type int
)

func (a *Activity) TypeKind() string {
	switch a.Type {
	case 1, 2, 3, 4, 14, 17:
		return TypeKindIssue
	case 5, 6, 7:
		return TypeKindWiki
	case 8, 9, 10:
		return TypeKindFile
	case 11, 12, 13:
		return TypeKindVCS
	case 18, 19, 20:
		return TypeKindPR
	}
	return TypeKindUnknown
}

func (a *Activity) TypeValue() int {
	switch a.Type {
	case 1, 5, 8, 13, 18:
		return TypeValueCreate
	case 2, 6, 9, 11, 12, 19:
		return TypeValueUpdate
	case 4, 7, 10, 14:
		return TypeValueDelete
	case 3, 17, 20:
		return TypeValueAction
	}
	return TypeValueUnknown
}
