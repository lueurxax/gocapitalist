package enums

type DocumentState string

const (
	Any       DocumentState = "any" // any document status
	InProcess DocumentState = "inprocess"
	Executed  DocumentState = "executed"
	Declined  DocumentState = "declined"
)
