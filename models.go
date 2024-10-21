package todoist_go

// ViewStyle is a custom type to restrict the view style to either "list" or "board".
type ViewStyle string

const (
	ViewStyleList  ViewStyle = "list"
	ViewStyleBoard ViewStyle = "board"
)

// Project represents a project from the Todoist API.
type Project struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	CommentCount   int       `json:"comment_count"`
	Order          int       `json:"order"`
	Color          string    `json:"color"`
	IsShared       bool      `json:"is_shared"`
	IsFavorite     bool      `json:"is_favorite"`
	IsInboxProject bool      `json:"is_inbox_project"`
	IsTeamInbox    bool      `json:"is_team_inbox"`
	ViewStyle      ViewStyle `json:"view_style"`
	URL            string    `json:"url"`
	ParentID       *string   `json:"parent_id"`
}

// ProjectParams defines the parameters available for creating and updating a project.
type ProjectParams struct {
	Name       string    `json:"name,omitempty"`
	ParentID   string    `json:"parent_id,omitempty"`
	Color      string    `json:"color,omitempty"`
	IsFavorite bool      `json:"is_favorite,omitempty"`
	ViewStyle  ViewStyle `json:"view_style,omitempty"`
}

// Section represents a section from the Todoist API.
type Section struct {
	ID        string `json:"id"`
	ProjectID string `json:"project_id"`
	Order     int    `json:"order"`
	Name      string `json:"name"`
}

// SectionParams defines the parameters available for creating and updating a section.
type SectionParams struct {
	ProjectID string `json:"project_id"`
	Name      string `json:"name"`
	Order     int    `json:"order,omitempty"`
}

// Task represents a task from the Todoist API.
type Task struct {
	ID           string        `json:"id"`
	ProjectID    string        `json:"project_id"`
	SectionID    string        `json:"section_id"`
	Content      string        `json:"content"`
	Description  string        `json:"description,omitempty"`
	IsCompleted  bool          `json:"is_completed"`
	Labels       []string      `json:"labels,omitempty"`
	Order        int           `json:"order"`
	Priority     int           `json:"priority"`
	AssigneeID   string        `json:"assignee_id,omitempty"`
	AssignerID   string        `json:"assigner_id,omitempty"`
	CommentCount int           `json:"comment_count,omitempty"`
	Due          *TaskDue      `json:"due,omitempty"`
	Duration     *TaskDuration `json:"duration,omitempty"`
	URL          string        `json:"url"`
}

// TaskDue represents the due date/time details of a task.
type TaskDue struct {
	Date        string `json:"date"`
	IsRecurring bool   `json:"is_recurring"`
	Datetime    string `json:"datetime,omitempty"`
	String      string `json:"string"`
	Timezone    string `json:"timezone,omitempty"`
}

// TaskDuration represents the duration of a task.
type TaskDuration struct {
	Amount int    `json:"amount"`
	Unit   string `json:"unit"`
}

// TaskParams defines the parameters for creating and updating a task.
type TaskParams struct {
	Content      string   `json:"content,omitempty"`
	Description  string   `json:"description,omitempty"`
	ProjectID    string   `json:"project_id,omitempty"`
	SectionID    string   `json:"section_id,omitempty"`
	Priority     int      `json:"priority,omitempty"`
	Labels       []string `json:"labels,omitempty"`
	DueString    string   `json:"due_string,omitempty"`
	DueDate      string   `json:"due_date,omitempty"`
	DueDatetime  string   `json:"due_datetime,omitempty"`
	DueLang      string   `json:"due_lang,omitempty"`
	AssigneeID   string   `json:"assignee_id,omitempty"`
	Duration     int      `json:"duration,omitempty"`
	DurationUnit string   `json:"duration_unit,omitempty"`
}

// Comment represents a comment from the Todoist API.
type Comment struct {
	ID         string      `json:"id"`
	TaskID     string      `json:"task_id,omitempty"`
	ProjectID  string      `json:"project_id,omitempty"`
	Content    string      `json:"content"`
	PostedAt   string      `json:"posted_at"`
	Attachment *Attachment `json:"attachment,omitempty"`
}

// Attachment represents an optional attachment in a comment.
type Attachment struct {
	FileName     string `json:"file_name,omitempty"`
	FileType     string `json:"file_type,omitempty"`
	FileURL      string `json:"file_url,omitempty"`
	ResourceType string `json:"resource_type,omitempty"`
}

// CommentParams defines the parameters for creating and updating a comment.
type CommentParams struct {
	TaskID     string      `json:"task_id,omitempty"`
	ProjectID  string      `json:"project_id,omitempty"`
	Content    string      `json:"content,omitempty"`
	Attachment *Attachment `json:"attachment,omitempty"`
}

// Label represents a personal label from the Todoist API.
type Label struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Color      string `json:"color"`
	Order      int    `json:"order"`
	IsFavorite bool   `json:"is_favorite"`
}

// LabelParams defines the parameters for creating and updating a label.
type LabelParams struct {
	Name       string `json:"name,omitempty"`
	Color      string `json:"color,omitempty"`
	Order      int    `json:"order,omitempty"`
	IsFavorite bool   `json:"is_favorite,omitempty"`
}

// SharedLabelParams defines the parameters for renaming or removing shared labels.
type SharedLabelParams struct {
	Name    string `json:"name,omitempty"`
	NewName string `json:"new_name,omitempty"` // For renaming
}
