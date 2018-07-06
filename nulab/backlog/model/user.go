package model

type (
	// User represents user.
	User struct {
		ID           int          `json:"id"`
		UserID       string       `json:"userId"`
		Name         string       `json:"name"`
		RoleType     int          `json:"roleType"`
		Lang         string       `json:"lang"`
		MailAddress  string       `json:"mailAddress"`
		NulabAccount NulabAccount `json:"nulabAccount"`
	}

	// NulabAccount represents account of nulab
	NulabAccount struct {
		NulabId  string `json:"nulabId"`
		Name     string `json:"name"`
		UniqueId string `json:"uniqueId"`
	}
)
