package opml

import (
	"github.com/gilliek/go-opml/opml"
	"github.com/ninedraft/thermalic/pkg/models"
)

func ParseOPMLFile(path string) ([]models.OPMLentry, error) {
	OPMLdata, err := opml.NewOPMLFromFile(path)
	if err != nil {
		return nil, err
	}
	outlines := OPMLdata.Outlines()
	feeds := make([]models.OPMLentry, 0, len(outlines))
	for _, outline := range outlines {
		feeds = append(feeds, models.OPMLentry{
			Name: outline.Text,
			URL:  outline.HTMLURL,
		})
	}
	return feeds, nil
}
