package tedlassoquotes

import (
	"fmt"
)

// buildURL from Author or random author
func buildURL(tag string) string {
	if tag == "" {
		return "https://tedlassoquotes.com/v1/quote"
	}
	return fmt.Sprintf("https://tedlassoquotes.com/v1/%s/quote", tag)
}

// APIResponse returned by the Ted Lasso Quotes API
type APIResponse struct {
	Quote      string `json:"quote"`
	Author     string `json:"author"`
	Tag        string `json:"tag"`
	ProfileImg string `json:"profile_img"`
}
