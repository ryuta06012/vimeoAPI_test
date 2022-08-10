import "github.com/silentsokolov/go-vimeo/vimeo"


func main() {
	client := vimeo.NewClient(tokenContext, nil)

	// Specific optional parameters
	cats, _, err := client.Categories.List(OptPage(1), OptPerPage(2), OptFields([]string{"name"}))
}