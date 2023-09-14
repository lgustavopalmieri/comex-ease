package ncm

type NcmInterface interface {
	IsValid() (bool, error)
	GetID() string
	GetCode() string
	GetDescription() string
	GetStartDate() string
	GetEndDate() string
	GetActType() string
	GetActNumber() int
	GetActYear() int
}
