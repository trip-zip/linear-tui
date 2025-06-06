package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const linearAPIURL = "https://api.linear.app/graphql"

type LinearClient struct {
	apiKey string
	client *http.Client
}

func NewLinearClient(apiKey string) *LinearClient {
	return &LinearClient{
		apiKey: apiKey,
		client: &http.Client{},
	}
}

func (lc *LinearClient) GetCurrentUser() (*User, error) {
	query := `
		query {
			viewer {
				id
				name
				email
			}
		}
	`

	reqBody := GraphQLRequest{Query: query}
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", linearAPIURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", lc.apiKey)

	resp, err := lc.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	var viewerResp ViewerResponse
	if err := json.NewDecoder(resp.Body).Decode(&viewerResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &viewerResp.Data.Viewer, nil
}

func (lc *LinearClient) GetIssues() ([]Issue, error) {
	query := `
		query {
			issues(first: 20) {
				nodes {
					id
					title
					description
					state {
						name
					}
					team {
						name
					}
					assignee {
						id
						name
					}
					priority
				}
			}
		}
	`

	reqBody := GraphQLRequest{Query: query}
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", linearAPIURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", lc.apiKey)

	resp, err := lc.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	var issuesResp IssuesResponse
	if err := json.NewDecoder(resp.Body).Decode(&issuesResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return issuesResp.Data.Issues.Nodes, nil
}

func (lc *LinearClient) GetMyIssues() ([]Issue, error) {
	// First get current user info
	user, err := lc.GetCurrentUser()
	if err != nil {
		return nil, fmt.Errorf("failed to get current user: %w", err)
	}

	// Query for issues assigned to the current user
	query := fmt.Sprintf(`
		query {
			issues(first: 50, filter: { assignee: { id: { eq: "%s" } } }) {
				nodes {
					id
					title
					description
					state {
						name
					}
					team {
						name
					}
					assignee {
						id
						name
					}
					priority
				}
			}
		}
	`, user.ID)

	reqBody := GraphQLRequest{Query: query}
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", linearAPIURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", lc.apiKey)

	resp, err := lc.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	var issuesResp IssuesResponse
	if err := json.NewDecoder(resp.Body).Decode(&issuesResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return issuesResp.Data.Issues.Nodes, nil
}

func (lc *LinearClient) GetIssuesByStatus(status string) ([]Issue, error) {
	query := fmt.Sprintf(`
		query {
			issues(first: 50, filter: { state: { name: { eq: "%s" } } }) {
				nodes {
					id
					title
					description
					state {
						name
					}
					team {
						name
					}
					assignee {
						id
						name
					}
					priority
				}
			}
		}
	`, status)

	reqBody := GraphQLRequest{Query: query}
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", linearAPIURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", lc.apiKey)

	resp, err := lc.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	var issuesResp IssuesResponse
	if err := json.NewDecoder(resp.Body).Decode(&issuesResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return issuesResp.Data.Issues.Nodes, nil
}

func (lc *LinearClient) GetMyIssuesByStatus(status string) ([]Issue, error) {
	// First get current user info
	user, err := lc.GetCurrentUser()
	if err != nil {
		return nil, fmt.Errorf("failed to get current user: %w", err)
	}

	// Query for issues assigned to the current user with specific status
	query := fmt.Sprintf(`
		query {
			issues(first: 50, filter: { 
				assignee: { id: { eq: "%s" } },
				state: { name: { eq: "%s" } }
			}) {
				nodes {
					id
					title
					description
					state {
						name
					}
					team {
						name
					}
					assignee {
						id
						name
					}
					priority
				}
			}
		}
	`, user.ID, status)

	reqBody := GraphQLRequest{Query: query}
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", linearAPIURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", lc.apiKey)

	resp, err := lc.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	var issuesResp IssuesResponse
	if err := json.NewDecoder(resp.Body).Decode(&issuesResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return issuesResp.Data.Issues.Nodes, nil
}

func getLinearClient() (*LinearClient, error) {
	apiKey := os.Getenv("LINEAR_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("LINEAR_API_KEY environment variable not set")
	}
	return NewLinearClient(apiKey), nil
}
