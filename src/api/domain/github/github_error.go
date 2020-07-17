package github

type GithubErrorResponse struct {
	StatusCode      int          `json:"status_code"`
	Message         string       `json:"message"`
	DocumentationUr string       `json:"documentation_url"`
	Errors          []GitubError `json:"errors"`
}

type GitubError struct {
	Resource string `json:"resource"`
	Code     string `json:"code"`
	Field    string `json:"field"`
	Message  string `json:"message"`
}

func (r GithubErrorResponse) Error() string {
	return r.Message
}
