package shared

type Code string

const (
	Success           Code = "00"
	DuplicatedRecord  Code = "01"
	NonExistentRecord Code = "02"
	ThirdPartyFail    Code = "11"
	Unkown            Code = "12"
	InvalidOperation  Code = "13"
)
