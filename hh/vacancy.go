package hh

type Vacancy struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Area struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"area"`
	Salary struct {
		From     int    `json:"from"`
		To       int    `json:"to"`
		Currency string `json:"currency"`
		Gross    bool   `json:"gross"`
	} `json:"salary"`
	Address struct {
		Raw string `json:"raw"`
	} `json:"address"`
	Experience struct {
		Name string `json:"name"`
	} `json:"experience"`
	Employer struct {
		Name string `json:"name"`
	} `json:"employer"`
	Employment struct {
		Name string `json:"name"`
	} `json:"employment"`
	Snippet struct {
		Requirement    string `json:"requirement"`
		Responsibility string `json:"responsibility"`
	} `json:"snippet"`
	Schedule struct {
		Name string `json:"name"`
	} `json:"schedule"`
	KeySkills []struct {
		Name string `json:"name"`
	} `json:"key_skills"`
	Description  string `json:"description"`
	AlternateURL string `json:"alternate_url"`
}

type Response struct {
	Items []Vacancy `json:"items"`
	Found int       `json:"found"`
}
