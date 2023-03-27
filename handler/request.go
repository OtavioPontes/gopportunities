package handler

import "fmt"

func errorParmIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Link == "" && r.Remote == nil && r.Salary <= 0 {
		return fmt.Errorf("request body is empty")
	}
	if r.Role == "" {
		return errorParmIsRequired("role", "string")
	}
	if r.Company == "" {
		return errorParmIsRequired("company", "string")
	}
	if r.Location == "" {
		return errorParmIsRequired("location", "string")
	}
	if r.Link == "" {
		return errorParmIsRequired("link", "string")
	}
	if r.Remote == nil {
		return errorParmIsRequired("remote", "bool")
	}
	if r.Salary <= 0 {
		return errorParmIsRequired("salary", "int")
	}
	return nil
}

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Link != "" || r.Remote != nil || r.Salary > 0 {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}
