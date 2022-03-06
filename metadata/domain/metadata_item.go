package domain

import (
	"strconv"
	"strings"

	traits "github.com/alephao/nftool/traits/domain"
)

type MetadataItem struct {
	Name            string            `json:"name"`
	Description     string            `json:"description"`
	Image           string            `json:"image"`
	ExternalLink    string            `json:"external_link"`
	BackgroundColor string         `json:"background_color"`
	AnimationUrl    string            `json:"animation_url"`
	Attributes   traits.TraitGroup `json:"attributes"`
}

// Template
// My Collection #{id}
// https://mycollection.com/token/{id}

func GenerateMetadata(collection traits.TraitCollection, name, description, image, externalLink, backgroundColor, animationUrl string) []MetadataItem {
	metadata := []MetadataItem{}

	for i, attrs := range collection {
		id := strconv.Itoa(i)
		item := MetadataItem{
			Name:         subPlaceholders(name, id),
			Description:  subPlaceholders(description, id),
			Image:        subPlaceholders(image, id),
			ExternalLink: subPlaceholders(externalLink, id),
			BackgroundColor: subPlaceholders(backgroundColor, id),
			AnimationUrl: subPlaceholders(animationUrl, id),
			Attributes:   attrs,
		}
		metadata = append(metadata, item)
	}

	return metadata
}

func subPlaceholders(str, id string) string {
	return strings.ReplaceAll(str, "{id}", id)
}
