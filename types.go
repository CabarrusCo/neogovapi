package neogovapi

type DepartmentResponse struct {
	ID     string `json:"id"`
	URL    string `json:"url"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type EmployeeQueryResponse struct {
	ID                  string             `json:"id"`
	URL                 string             `json:"url"`
	Active              string             `json:"active"`
	Firstname           string             `json:"firstname"`
	Lastname            string             `json:"lastname"`
	Employeenumber      string             `json:"employeenumber"`
	Email               string             `json:"email"`
	Createddate         string             `json:"createddate"`
	Evaluationcycledate string             `json:"evaluationcycledate"`
	Hiredate            string             `json:"hiredate"`
	Updateddate         string             `json:"updateddate"`
	Separationdate      string             `json:"separationdate"`
	Address1            string             `json:"address1"`
	Address2            string             `json:"address2"`
	City                string             `json:"city"`
	State               string             `json:"state"`
	Zipcode             string             `json:"zipcode"`
	Phone               string             `json:"phone"`
	Employeephotourl    string             `json:"employeephotourl"`
	User                User               `json:"user"`
	Directmanager       DirectManager      `json:"directmanager"`
	Position            Position           `json:"position"`
	Department          Department         `json:"department"`
	Division            Division           `json:"division"`
	Classspecification  ClassSpecification `json:"classspecification"`
}

type User struct {
	ID           string `json:"id"`
	URL          string `json:"url"`
	Username     string `json:"username"`
	Activeuser   string `json:"activeuser"`
	OnlineStatus string `json:"onlineStatus"`
}

type DirectManager struct {
	Fullname string `json:"fullname"`
	ID       string `json:"id"`
	URI      string `json:"uri"`
}

type Position struct {
	ID    string `json:"id"`
	URI   string `json:"uri"`
	Code  string `json:"code"`
	Title string `json:"title"`
}

type Department struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	URI  string `json:"uri"`
	Code string `json:"code"`
}

type Division struct {
	Fullname string `json:"fullname"`
	ID       string `json:"id"`
	URI      string `json:"uri"`
}

type ClassSpecification struct {
	ID    string `json:"id"`
	URI   string `json:"uri"`
	Code  string `json:"code"`
	Title string `json:"title"`
}

type Evaluation struct {
	ID                    string         `json:"id"`
	URL                   string         `json:"url"`
	Name                  string         `json:"name"`
	Type                  string         `json:"type"`
	Description           string         `json:"description"`
	DueDate               string         `json:"dueDate"`
	ShowNumericScoring    bool           `json:"showNumericScoring"`
	ShowOverallRating     bool           `json:"showOverallRating"`
	Status                string         `json:"status"`
	Score                 Score          `json:"score"`
	Settings              Settings       `json:"settings"`
	Employee              Employee       `json:"employee"`
	Raters                []Rater        `json:"raters"`
	Approvers             []Approver     `json:"approvers"`
	Acknowledgers         []Acknowledger `json:"acknowledgers"`
	Criticalelementfailed bool           `json:"criticalelementfailed"`
	CompletionDate        string         `json:"completionDate"`
}

type Numeric struct {
	Actual          string `json:"actual"`
	MaximumPossible string `json:"maximumPossible"`
	Percentage      string `json:"percentage"`
}

type Rating struct {
	Text  string `json:"text"`
	Value string `json:"value"`
}

type Score struct {
	Numeric Numeric `json:"numeric"`
	Rating  Rating  `json:"rating"`
}

type NumericScoring struct {
	Enabled        bool `json:"enabled"`
	ShowActual     bool `json:"showActual"`
	ShowMaximum    bool `json:"showMaximum"`
	ShowPercentage bool `json:"showPercentage"`
}

type RatingScale struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type OverallRating struct {
	Enabled     bool        `json:"enabled"`
	RatingScale RatingScale `json:"ratingScale"`
	ShowText    bool        `json:"showText"`
	ShowValue   bool        `json:"showValue"`
}

type Settings struct {
	NumericScoring NumericScoring `json:"numericScoring"`
	OverallRating  OverallRating  `json:"overallRating"`
}

type Employee struct {
	ID             string `json:"id"`
	URL            string `json:"url"`
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
	Employeenumber string `json:"employeenumber"`
	Email          string `json:"email"`
}

type Rater struct {
	Employee        Employee `json:"employee"`
	RaterOfRecord   bool     `json:"raterOfRecord"`
	Submitted       bool     `json:"submitted"`
	CalculatedScore float64  `json:"CalculatedScore"`
	OverallRating   string   `json:"overallRating"`
	DueDate         string   `json:"dueDate"`
	LanguageFlag    bool     `json:"languageFlag"`
	Weight          float64  `json:"weight"`
	StatusDate      string   `json:"statusDate"`
}

type Approver struct {
	Employee           Employee `json:"employee"`
	Type               string   `json:"type"`
	RequiredForRelease bool     `json:"requiredForRelease"`
	DueDate            string   `json:"dueDate"`
	Status             string   `json:"status"`
	StatusDate         string   `json:"statusDate"`
	Comment            string   `json:"comment"`
}

type Acknowledger struct {
	Employee   Employee `json:"employee"`
	DueDate    string   `json:"dueDate"`
	Status     string   `json:"status"`
	StatusDate string   `json:"statusDate"`
}
