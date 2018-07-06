package model

type (
	// Content represents content
	Content struct {
		ID             int          `json:"id"`
		Number         int          `json:"number"`
		KeyID          int          `json:"key_id"`
		Summary        string       `json:"summary"`
		Description    string       `json:"description"`
		Comment        Comment      `json:"comment"`
		Changes        []Change     `json:"changes"`
		Issue          Issue        `json:"issue"`
		Repository     Repository   `json:"repository"`
		Change_type    string       `json:"change_type"`
		Revision_type  string       `json:"revision_type"`
		Ref            string       `json:"ref"`
		Revision_count int          `json:"revision_count"`
		Revisions      []Revision   `json:"revisions"`
		Attachments    []Attachment `json:"attachments"`
		SharedFiles    []SharedFile `json:"shared_files"`
	}

	// Revision represents revision
	Revision struct {
		Rev     string `json:"rev"`
		Comment string `json:"comment"`
	}

	// Change represents change
	Change struct {
		Field     string `json:"field"`
		NewValues string `json:"new_value"`
		OldValue  string `json:"old_value"`
	}
)
