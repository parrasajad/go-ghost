package admin

import (
	"encoding/json"
	"github.com/parrasajad/go-ghost"
	"github.com/writeas/go-mobiledoc"
	"net/http"
)

func AddPost(gc *ghost.Client, p ghost.PostParams) error {
	c := getC(gc)

	// Convert params to mobiledoc
	if *p.Markdown != "" {
		mobdoc, err := json.Marshal(mobiledoc.FromMarkdown(*p.Markdown))
		if err == nil {
			p.Mobiledoc = ghost.String(string(mobdoc))
			p.Markdown = nil
		}
	}
	// TODO: also handle HTML

	_, err := c.Request(http.MethodPost, c.endpoint("posts"), map[string]interface{}{
		"posts": []ghost.PostParams{p},
	})
	return err
}
