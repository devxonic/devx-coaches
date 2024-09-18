package api

type Student struct {
	Id string `json:"id"`
	// FirstName string `json:"firstName"`
	// LastName  string `json:"lastName"`
	Name string `json:"name"`
	// Number int64  `json:"number"`
	// CreatedAt    time.Time `json:"createdAt"`
	Gender string `json:"gender"`
	// Ccode      string `json:"ccode"`
	FeesAmount string `json:"fees_amount"`
	Phone      string `json:"phone"`
	Class_id   string `json:"classId"`
	S_of       string `json:"s_of"`
	Email      string `json:"email"`
	Batch_id   string `json:"batchId"`
	Subject_id string `json:"subjectId"`
	Nic_no     string `json:"cnic"`
	CAddress   string `json:"address"`
	Adm_fee    string `json:"admissionFee"`
	Status     int    `json:"status"`
	// Code       string `json:"code"`
	// Cid        string `json:"cid"`
	// Secid      string `json:"secid"`
	// Subid      string `json:"subid"`
	// Subcode    string `json:"subcode"`
	IsActive  int `json:"is_active"`
	IsDeleted int `json:"is_deleted"`
}

type Class struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	// Code        string `json:"code"`
	Status int `json:"status"`
}

type Subject struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	// Code        string `json:"code"`
	Status int `json:"status"`
}

type Period struct {
	Id                string `json:"id"`
	YearId            string `json:"yearId"`
	Month             string `json:"month"`
	PeriodDescription string `json:"periodDescription"`
	PeriodDate        string `json:"periodDate"`
	Status            int    `json:"status"`
	Isactive          int    `json:"isactive"`
	Isdeleted         int    `json:"isdeleted"`
}

type Teachers struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

type Year struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Status      int    `json:"status"`
	Remarks     string `json:"remarks"`
}

type YearMonth struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	YearId      string
	Status      int `json:"status"`
}

type Batch struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

type YearMonthJoin struct {
	Id                   string `json:"id"`
	Description          string `json:"description"`
	Status               string `json:"status"`
	Yearmonthstatus      string `json:"yearmonthstatus"`
	Yearmonthid          string `json:"yearmonthid"`
	Yearmonthdescription string `json:"yearmonthdescription"`
	Months               string `json:"months"`
}
