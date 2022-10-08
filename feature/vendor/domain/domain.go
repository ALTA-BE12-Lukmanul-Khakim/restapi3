package domain

type Vendor struct {
	ID       uint
	NamaVen  string
	Expedisi string
}

type VendorData interface {
	Insert(newVendor Vendor) (Vendor, error)
	GetAll() ([]Vendor, error)
	GetVen(expedisi string) ([]Vendor, error)
	//GetIsDone(isDone bool) ([]Vendor, error)
}

type VendorService interface {
	AddVendor(newVendor Vendor) (Vendor, error)
	ShowAllVendor() ([]Vendor, error)
	ShowExpVendor(expedisi string) ([]Vendor, error)
}
