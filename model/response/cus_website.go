package response

type WebsiteResponse struct {
	TargetUrl   string `json:"argetUrl"`
	Protocol    string `json:"protocol"`
	Domain      string `json:"domain"`
	Path        string `json:"path"`
	Query       string `json:"query"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
