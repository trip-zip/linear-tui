package cmd

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Issue struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	State       struct {
		Name string `json:"name"`
	} `json:"state"`
	Team struct {
		Name string `json:"name"`
	} `json:"team"`
	Assignee struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"assignee"`
	Priority int `json:"priority"`
}

type IssuesResponse struct {
	Data struct {
		Issues struct {
			Nodes []Issue `json:"nodes"`
		} `json:"issues"`
	} `json:"data"`
}

type ViewerResponse struct {
	Data struct {
		Viewer User `json:"viewer"`
	} `json:"data"`
}

type GraphQLRequest struct {
	Query string `json:"query"`
}
