package diff

import "testing"

func TestOpenAPIReportsRemovedPath(t *testing.T) {
	old := map[string]any{
		"paths": map[string]any{
			"/old": map[string]any{"get": map[string]any{}},
		},
	}
	new := map[string]any{"paths": map[string]any{}}

	changes := OpenAPI(old, new)
	if len(changes) != 1 {
		t.Fatalf("expected 1 change, got %d", len(changes))
	}
	if changes[0].Type != "PATH_REMOVED" || changes[0].Severity != Breaking {
		t.Fatalf("unexpected change: %#v", changes[0])
	}
}

func TestSitemapReportsAddedURL(t *testing.T) {
	old := map[string]any{"/old": true}
	new := map[string]any{"/old": true, "/new": true}

	changes := Sitemap(old, new)
	if len(changes) != 1 || changes[0].Type != "SITEMAP_URL_ADDED" {
		t.Fatalf("unexpected changes: %#v", changes)
	}
}
