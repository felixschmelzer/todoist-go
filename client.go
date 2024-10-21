package todoist

import (
	"fmt"
	"net/http"
)

// TodoistClient represents the Todoist API client.
type TodoistClient struct {
	BaseURL    string
	Token      string
	HTTPClient *http.Client
}

// NewClient initializes a new Todoist API client.
func NewTodoistClient(token string) *TodoistClient {
	return &TodoistClient{
		BaseURL:    "https://api.todoist.com/rest/v2",
		Token:      token,
		HTTPClient: &http.Client{},
	}
}

// CreateProject creates a new project on Todoist.
func (c *TodoistClient) CreateProject(params ProjectParams) (*Project, error) {
	url := fmt.Sprintf("%s/projects", c.BaseURL)

	resp, err := sendRequest(c.HTTPClient, "POST", url, c.Token, params)
	if err != nil {
		return nil, err
	}

	var project Project
	if err := parseResponse(resp, &project); err != nil {
		return nil, err
	}

	return &project, nil
}

// UpdateProject updates an existing project on Todoist.
func (c *TodoistClient) UpdateProject(id string, params ProjectParams) (*Project, error) {
	url := fmt.Sprintf("%s/projects/%s", c.BaseURL, id)

	resp, err := sendRequest(c.HTTPClient, "POST", url, c.Token, params)
	if err != nil {
		return nil, err
	}

	var project Project
	if err := parseResponse(resp, &project); err != nil {
		return nil, err
	}

	return &project, nil
}

// GetProject fetches a specific project by its ID from Todoist.
func (c *TodoistClient) GetProject(id string) (*Project, error) {
	url := fmt.Sprintf("%s/projects/%s", c.BaseURL, id)

	resp, err := sendRequest(c.HTTPClient, "GET", url, c.Token, nil)
	if err != nil {
		return nil, err
	}

	var project Project
	if err := parseResponse(resp, &project); err != nil {
		return nil, err
	}

	return &project, nil
}

// DeleteProject deletes a specific project by its ID from Todoist.
func (c *TodoistClient) DeleteProject(id string) (bool, error) {
	url := fmt.Sprintf("%s/projects/%s", c.BaseURL, id)

	resp, err := sendRequest(c.HTTPClient, "DELETE", url, c.Token, nil)
	if err != nil {
		return false, err
	}

	if resp.StatusCode != http.StatusNoContent {
		return false, fmt.Errorf("failed to delete project, status code: %d", resp.StatusCode)
	}

	return true, nil
}

// GetSections fetches all sections for a given project from Todoist.
func (c *TodoistClient) GetSections(projectID string) ([]Section, error) {
	url := fmt.Sprintf("%s/sections", c.BaseURL)
	if projectID != "" {
		url = fmt.Sprintf("%s?project_id=%s", url, projectID)
	}

	resp, err := sendRequest(c.HTTPClient, "GET", url, c.Token, nil)
	if err != nil {
		return nil, err
	}

	var sections []Section
	if err := parseResponse(resp, &sections); err != nil {
		return nil, err
	}

	return sections, nil
}

// CreateSection creates a new section on Todoist.
func (c *TodoistClient) CreateSection(params SectionParams) (*Section, error) {
	url := fmt.Sprintf("%s/sections", c.BaseURL)

	resp, err := sendRequest(c.HTTPClient, "POST", url, c.Token, params)
	if err != nil {
		return nil, err
	}

	var section Section
	if err := parseResponse(resp, &section); err != nil {
		return nil, err
	}

	return &section, nil
}

// GetSection fetches a specific section by its ID from Todoist.
func (c *TodoistClient) GetSection(id string) (*Section, error) {
	url := fmt.Sprintf("%s/sections/%s", c.BaseURL, id)

	resp, err := sendRequest(c.HTTPClient, "GET", url, c.Token, nil)
	if err != nil {
		return nil, err
	}

	var section Section
	if err := parseResponse(resp, &section); err != nil {
		return nil, err
	}

	return &section, nil
}

// UpdateSection updates a specific section by its ID on Todoist.
func (c *TodoistClient) UpdateSection(id string, params SectionParams) (*Section, error) {
	url := fmt.Sprintf("%s/sections/%s", c.BaseURL, id)

	resp, err := sendRequest(c.HTTPClient, "POST", url, c.Token, params)
	if err != nil {
		return nil, err
	}

	var section Section
	if err := parseResponse(resp, &section); err != nil {
		return nil, err
	}

	return &section, nil
}

// DeleteSection deletes a specific section by its ID from Todoist.
func (c *TodoistClient) DeleteSection(id string) (bool, error) {
	url := fmt.Sprintf("%s/sections/%s", c.BaseURL, id)

	resp, err := sendRequest(c.HTTPClient, "DELETE", url, c.Token, nil)
	if err != nil {
		return false, err
	}

	if resp.StatusCode != http.StatusNoContent {
		return false, fmt.Errorf("failed to delete section, status code: %d", resp.StatusCode)
	}

	return true, nil
}

// GetTasks fetches all active tasks from Todoist, optionally filtered by project, section, or label.
func (c *TodoistClient) GetTasks(projectID, sectionID, label string) ([]Task, error) {
	url := fmt.Sprintf("%s/tasks", c.BaseURL)

	// Apply filters
	if projectID != "" {
		url = fmt.Sprintf("%s?project_id=%s", url, projectID)
	} else if sectionID != "" {
		url = fmt.Sprintf("%s?section_id=%s", url, sectionID)
	} else if label != "" {
		url = fmt.Sprintf("%s?label=%s", url, label)
	}

	resp, err := sendRequest(c.HTTPClient, "GET", url, c.Token, nil)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	if err := parseResponse(resp, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

// CreateTask creates a new task on Todoist.
func (c *TodoistClient) CreateTask(params TaskParams) (*Task, error) {
	url := fmt.Sprintf("%s/tasks", c.BaseURL)

	resp, err := sendRequest(c.HTTPClient, "POST", url, c.Token, params)
	if err != nil {
		return nil, err
	}

	var task Task
	if err := parseResponse(resp, &task); err != nil {
		return nil, err
	}

	return &task, nil
}

// GetTask fetches a specific task by its ID from Todoist.
func (c *TodoistClient) GetTask(id string) (*Task, error) {
	url := fmt.Sprintf("%s/tasks/%s", c.BaseURL, id)

	resp, err := sendRequest(c.HTTPClient, "GET", url, c.Token, nil)
	if err != nil {
		return nil, err
	}

	var task Task
	if err := parseResponse(resp, &task); err != nil {
		return nil, err
	}

	return &task, nil
}

// UpdateTask updates a specific task by its ID on Todoist.
func (c *TodoistClient) UpdateTask(id string, params TaskParams) (*Task, error) {
	url := fmt.Sprintf("%s/tasks/%s", c.BaseURL, id)

	resp, err := sendRequest(c.HTTPClient, "POST", url, c.Token, params)
	if err != nil {
		return nil, err
	}

	var task Task
	if err := parseResponse(resp, &task); err != nil {
		return nil, err
	}

	return &task, nil
}

// CloseTask closes a specific task by its ID on Todoist.
func (c *TodoistClient) CloseTask(id string) (bool, error) {
	url := fmt.Sprintf("%s/tasks/%s/close", c.BaseURL, id)

	resp, err := sendRequest(c.HTTPClient, "POST", url, c.Token, nil)
	if err != nil {
		return false, err
	}

	if resp.StatusCode != http.StatusNoContent {
		return false, fmt.Errorf("failed to close task, status code: %d", resp.StatusCode)
	}

	return true, nil
}

// ReopenTask reopens a specific task by its ID on Todoist.
func (c *TodoistClient) ReopenTask(id string) (bool, error) {
	url := fmt.Sprintf("%s/tasks/%s/reopen", c.BaseURL, id)

	resp, err := sendRequest(c.HTTPClient, "POST", url, c.Token, nil)
	if err != nil {
		return false, err
	}

	if resp.StatusCode != http.StatusNoContent {
		return false, fmt.Errorf("failed to reopen task, status code: %d", resp.StatusCode)
	}

	return true, nil
}

// DeleteTask deletes a specific task by its ID from Todoist.
func (c *TodoistClient) DeleteTask(id string) (bool, error) {
	url := fmt.Sprintf("%s/tasks/%s", c.BaseURL, id)

	resp, err := sendRequest(c.HTTPClient, "DELETE", url, c.Token, nil)
	if err != nil {
		return false, err
	}

	if resp.StatusCode != http.StatusNoContent {
		return false, fmt.Errorf("failed to delete task, status code: %d", resp.StatusCode)
	}

	return true, nil
}

// GetComments fetches all comments for a given task or project from Todoist.
func (c *TodoistClient) GetComments(taskID, projectID string) ([]Comment, error) {
	url := fmt.Sprintf("%s/comments", c.BaseURL)

	// Apply filters
	if taskID != "" {
		url = fmt.Sprintf("%s?task_id=%s", url, taskID)
	} else if projectID != "" {
		url = fmt.Sprintf("%s?project_id=%s", url, projectID)
	}

	resp, err := sendRequest(c.HTTPClient, "GET", url, c.Token, nil)
	if err != nil {
		return nil, err
	}

	var comments []Comment
	if err := parseResponse(resp, &comments); err != nil {
		return nil, err
	}

	return comments, nil
}

// CreateComment creates a new comment on a task or project.
func (c *TodoistClient) CreateComment(params CommentParams) (*Comment, error) {
	url := fmt.Sprintf("%s/comments", c.BaseURL)

	resp, err := sendRequest(c.HTTPClient, "POST", url, c.Token, params)
	if err != nil {
		return nil, err
	}

	var comment Comment
	if err := parseResponse(resp, &comment); err != nil {
		return nil, err
	}

	return &comment, nil
}

// GetComment fetches a specific comment by its ID from Todoist.
func (c *TodoistClient) GetComment(id string) (*Comment, error) {
	url := fmt.Sprintf("%s/comments/%s", c.BaseURL, id)

	resp, err := sendRequest(c.HTTPClient, "GET", url, c.Token, nil)
	if err != nil {
		return nil, err
	}

	var comment Comment
	if err := parseResponse(resp, &comment); err != nil {
		return nil, err
	}

	return &comment, nil
}

// UpdateComment updates a specific comment by its ID on Todoist.
func (c *TodoistClient) UpdateComment(id string, params CommentParams) (*Comment, error) {
	url := fmt.Sprintf("%s/comments/%s", c.BaseURL, id)

	resp, err := sendRequest(c.HTTPClient, "POST", url, c.Token, params)
	if err != nil {
		return nil, err
	}

	var comment Comment
	if err := parseResponse(resp, &comment); err != nil {
		return nil, err
	}

	return &comment, nil
}

// DeleteComment deletes a specific comment by its ID from Todoist.
func (c *TodoistClient) DeleteComment(id string) (bool, error) {
	url := fmt.Sprintf("%s/comments/%s", c.BaseURL, id)

	resp, err := sendRequest(c.HTTPClient, "DELETE", url, c.Token, nil)
	if err != nil {
		return false, err
	}

	if resp.StatusCode != http.StatusNoContent {
		return false, fmt.Errorf("failed to delete comment, status code: %d", resp.StatusCode)
	}

	return true, nil
}

// GetLabels fetches all personal labels from Todoist.
func (c *TodoistClient) GetLabels() ([]Label, error) {
	url := fmt.Sprintf("%s/labels", c.BaseURL)

	resp, err := sendRequest(c.HTTPClient, "GET", url, c.Token, nil)
	if err != nil {
		return nil, err
	}

	var labels []Label
	if err := parseResponse(resp, &labels); err != nil {
		return nil, err
	}

	return labels, nil
}

// CreateLabel creates a new personal label on Todoist.
func (c *TodoistClient) CreateLabel(params LabelParams) (*Label, error) {
	url := fmt.Sprintf("%s/labels", c.BaseURL)

	resp, err := sendRequest(c.HTTPClient, "POST", url, c.Token, params)
	if err != nil {
		return nil, err
	}

	var label Label
	if err := parseResponse(resp, &label); err != nil {
		return nil, err
	}

	return &label, nil
}

// GetLabel fetches a specific personal label by its ID from Todoist.
func (c *TodoistClient) GetLabel(id string) (*Label, error) {
	url := fmt.Sprintf("%s/labels/%s", c.BaseURL, id)

	resp, err := sendRequest(c.HTTPClient, "GET", url, c.Token, nil)
	if err != nil {
		return nil, err
	}

	var label Label
	if err := parseResponse(resp, &label); err != nil {
		return nil, err
	}

	return &label, nil
}

// UpdateLabel updates a specific personal label by its ID on Todoist.
func (c *TodoistClient) UpdateLabel(id string, params LabelParams) (*Label, error) {
	url := fmt.Sprintf("%s/labels/%s", c.BaseURL, id)

	resp, err := sendRequest(c.HTTPClient, "POST", url, c.Token, params)
	if err != nil {
		return nil, err
	}

	var label Label
	if err := parseResponse(resp, &label); err != nil {
		return nil, err
	}

	return &label, nil
}

// DeleteLabel deletes a specific personal label by its ID from Todoist.
func (c *TodoistClient) DeleteLabel(id string) (bool, error) {
	url := fmt.Sprintf("%s/labels/%s", c.BaseURL, id)

	resp, err := sendRequest(c.HTTPClient, "DELETE", url, c.Token, nil)
	if err != nil {
		return false, err
	}

	if resp.StatusCode != http.StatusNoContent {
		return false, fmt.Errorf("failed to delete label, status code: %d", resp.StatusCode)
	}

	return true, nil
}

// GetSharedLabels fetches all shared labels from Todoist.
func (c *TodoistClient) GetSharedLabels(omitPersonal bool) ([]string, error) {
	url := fmt.Sprintf("%s/labels/shared", c.BaseURL)

	if omitPersonal {
		url = fmt.Sprintf("%s?omit_personal=true", url)
	}

	resp, err := sendRequest(c.HTTPClient, "GET", url, c.Token, nil)
	if err != nil {
		return nil, err
	}

	var sharedLabels []string
	if err := parseResponse(resp, &sharedLabels); err != nil {
		return nil, err
	}

	return sharedLabels, nil
}

// RenameSharedLabel renames a shared label on Todoist.
func (c *TodoistClient) RenameSharedLabel(params SharedLabelParams) (bool, error) {
	url := fmt.Sprintf("%s/labels/shared/rename", c.BaseURL)

	resp, err := sendRequest(c.HTTPClient, "POST", url, c.Token, params)
	if err != nil {
		return false, err
	}

	if resp.StatusCode != http.StatusNoContent {
		return false, fmt.Errorf("failed to rename shared label, status code: %d", resp.StatusCode)
	}

	return true, nil
}

// RemoveSharedLabel removes a shared label from Todoist.
func (c *TodoistClient) RemoveSharedLabel(params SharedLabelParams) (bool, error) {
	url := fmt.Sprintf("%s/labels/shared/remove", c.BaseURL)

	resp, err := sendRequest(c.HTTPClient, "POST", url, c.Token, params)
	if err != nil {
		return false, err
	}

	if resp.StatusCode != http.StatusNoContent {
		return false, fmt.Errorf("failed to remove shared label, status code: %d", resp.StatusCode)
	}

	return true, nil
}
