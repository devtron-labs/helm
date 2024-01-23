package registry // import "helm.sh/helm/v3/pkg/registry"
import (
	"oras.land/oras-go/pkg/registry"
	registryremote "oras.land/oras-go/pkg/registry/remote"
	"strings"
)

// GetTagsIgnoreSemVer implements Tags function but removes semver StrictNewVersion check
func (c *Client) GetTagsIgnoreSemVer(ref string) ([]string, error) {
	parsedReference, err := registry.ParseReference(ref)
	if err != nil {
		return nil, err
	}

	repository := registryremote.Repository{
		Reference: parsedReference,
		Client:    c.registryAuthorizer,
		PlainHTTP: c.plainHTTP,
	}

	var registryTags []string

	registryTags, err = registry.Tags(ctx(c.out, c.debug), &repository)
	if err != nil {
		return nil, err
	}

	var tagVersions []string
	for _, tag := range registryTags {
		// Change underscore (_) back to plus (+) for Helm
		// See https://github.com/helm/helm/issues/10166
		tagVersion := strings.ReplaceAll(tag, "_", "+")
		tagVersions = append(tagVersions, tagVersion)

	}

	return tagVersions, nil
}
