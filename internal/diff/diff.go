package diff

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

type Severity string

const (
	Breaking   Severity = "BREAKING"
	Warning    Severity = "WARNING"
	Compatible Severity = "COMPATIBLE"
)

type Change struct {
	Severity Severity `json:"severity"`
	Type     string   `json:"type"`
	Path     string   `json:"path"`
	Message  string   `json:"message"`
	Before   any      `json:"before,omitempty"`
	After    any      `json:"after,omitempty"`
}

func LoadDocument(path string) (map[string]any, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var value map[string]any
	trimmed := strings.TrimSpace(string(raw))
	if strings.HasPrefix(trimmed, "{") {
		if err := json.Unmarshal(raw, &value); err != nil {
			return nil, fmt.Errorf("parse JSON: %w", err)
		}
		return value, nil
	}
	if err := yaml.Unmarshal(raw, &value); err != nil {
		return nil, fmt.Errorf("parse YAML: %w", err)
	}
	return value, nil
}

func LoadSitemap(path string) (map[string]any, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var root struct {
		URLs []struct {
			Loc     string `xml:"loc"`
			Lastmod string `xml:"lastmod"`
		} `xml:"url"`
	}
	if err := xml.Unmarshal(raw, &root); err != nil {
		return nil, fmt.Errorf("parse sitemap XML: %w", err)
	}
	result := make(map[string]any, len(root.URLs))
	for _, item := range root.URLs {
		if item.Loc == "" {
			continue
		}
		value := any(true)
		if item.Lastmod != "" {
			value = item.Lastmod
		}
		result[item.Loc] = value
	}
	return result, nil
}

func LoadFeed(path string) (map[string]any, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(string(raw))
	if strings.HasPrefix(trimmed, "{") {
		return loadJSONFeed(raw)
	}

	var rss struct {
		Channel struct {
			Items []struct {
				GUID        string `xml:"guid"`
				Link        string `xml:"link"`
				Title       string `xml:"title"`
				Description string `xml:"description"`
			} `xml:"item"`
		} `xml:"channel"`
	}
	if err := xml.Unmarshal(raw, &rss); err == nil && len(rss.Channel.Items) > 0 {
		result := make(map[string]any, len(rss.Channel.Items))
		for _, item := range rss.Channel.Items {
			key := firstNonEmpty(item.GUID, item.Link, item.Title)
			if key != "" {
				result[key] = firstNonEmpty(item.Title, item.Description)
			}
		}
		return result, nil
	}

	var atom struct {
		Entries []struct {
			ID      string `xml:"id"`
			Link    string `xml:"link"`
			Title   string `xml:"title"`
			Summary string `xml:"summary"`
		} `xml:"entry"`
	}
	if err := xml.Unmarshal(raw, &atom); err != nil {
		return nil, fmt.Errorf("parse feed XML: %w", err)
	}
	result := make(map[string]any, len(atom.Entries))
	for _, item := range atom.Entries {
		key := firstNonEmpty(item.ID, item.Link, item.Title)
		if key != "" {
			result[key] = firstNonEmpty(item.Title, item.Summary)
		}
	}
	return result, nil
}

func loadJSONFeed(raw []byte) (map[string]any, error) {
	var payload struct {
		Items []map[string]any `json:"items"`
	}
	if err := json.Unmarshal(raw, &payload); err != nil {
		return nil, fmt.Errorf("parse JSON Feed: %w", err)
	}
	result := make(map[string]any, len(payload.Items))
	for _, item := range payload.Items {
		key := firstNonEmpty(stringValue(item["id"]), stringValue(item["url"]))
		if key != "" {
			result[key] = firstNonEmpty(stringValue(item["title"]), stringValue(item["summary"]))
		}
	}
	return result, nil
}

func stringValue(value any) string {
	result, _ := value.(string)
	return result
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}

func asMap(value any) map[string]any {
	switch typed := value.(type) {
	case map[string]any:
		return typed
	case map[any]any:
		result := make(map[string]any, len(typed))
		for key, item := range typed {
			result[fmt.Sprint(key)] = item
		}
		return result
	}
	return nil
}

func asList(value any) []any {
	list, _ := value.([]any)
	return list
}

func changed(before, after any) bool {
	return fmt.Sprintf("%v", before) != fmt.Sprintf("%v", after)
}

func OpenAPI(old, new map[string]any) []Change {
	changes := []Change{}
	oldPaths := asMap(old["paths"])
	newPaths := asMap(new["paths"])
	if oldPaths == nil || newPaths == nil {
		return changes
	}

	for _, path := range sortedKeys(oldPaths) {
		if _, ok := newPaths[path]; !ok {
			changes = append(changes, Change{
				Severity: Breaking,
				Type:     "PATH_REMOVED",
				Path:     "paths." + path,
				Message:  fmt.Sprintf("path %q was removed", path),
			})
		}
	}
	for _, path := range sortedKeys(newPaths) {
		if _, ok := oldPaths[path]; !ok {
			changes = append(changes, Change{
				Severity: Breaking,
				Type:     "PATH_ADDED",
				Path:     "paths." + path,
				Message:  fmt.Sprintf("new path %q was added", path),
				After:    newPaths[path],
			})
		}
	}

	for _, path := range sortedKeys(oldPaths) {
		oldItem := asMap(oldPaths[path])
		newItem := asMap(newPaths[path])
		if oldItem == nil || newItem == nil {
			continue
		}
		for _, method := range sortedKeys(oldItem) {
			if _, ok := newItem[method]; !ok {
				changes = append(changes, Change{
					Severity: Breaking,
					Type:     "OPERATION_REMOVED",
					Path:     "paths." + path + "." + method,
					Message:  fmt.Sprintf("%s %s was removed", strings.ToUpper(method), path),
				})
			}
		}
		for _, method := range sortedKeys(newItem) {
			if _, ok := oldItem[method]; !ok {
				changes = append(changes, Change{
					Severity: Breaking,
					Type:     "OPERATION_ADDED",
					Path:     "paths." + path + "." + method,
					Message:  fmt.Sprintf("%s %s was added", strings.ToUpper(method), path),
				})
			}
		}
		for _, method := range sortedKeys(oldItem) {
			compareOperation(changes, "paths."+path+"."+method, asMap(oldItem[method]), asMap(newItem[method]))
		}
	}

	compareSecurity(changes, old, new)
	return changes
}

func Schema(old, new map[string]any) []Change {
	changes := []Change{}
	compareMaps(changes, "schema", old, new)
	return changes
}

func compareOperation(changes []Change, base string, old, new map[string]any) {
	if old == nil || new == nil {
		return
	}

	oldRequired := requiredParameterNames(asList(old["parameters"]))
	newRequired := requiredParameterNames(asList(new["parameters"]))
	for name := range oldRequired {
		if !newRequired[name] {
			changes = append(changes, Change{
				Severity: Compatible,
				Type:     "PARAMETER_OPTIONAL",
				Path:     base + ".parameters." + name,
				Message:  fmt.Sprintf("parameter %q became optional", name),
			})
		}
	}
	for name := range newRequired {
		if !oldRequired[name] {
			changes = append(changes, Change{
				Severity: Breaking,
				Type:     "PARAMETER_REQUIRED",
				Path:     base + ".parameters." + name,
				Message:  fmt.Sprintf("parameter %q became required", name),
			})
		}
	}

	oldResponses := asMap(old["responses"])
	newResponses := asMap(new["responses"])
	if oldResponses != nil && newResponses != nil {
		for code := range oldResponses {
			if _, ok := newResponses[code]; !ok {
				changes = append(changes, Change{
					Severity: Breaking,
					Type:     "RESPONSE_REMOVED",
					Path:     base + ".responses." + code,
					Message:  fmt.Sprintf("response %s was removed", code),
				})
			}
		}
	}
}

func requiredParameterNames(params []any) map[string]bool {
	result := map[string]bool{}
	for _, item := range params {
		param := asMap(item)
		if param == nil {
			continue
		}
		name, _ := param["name"].(string)
		if name == "" {
			continue
		}
		required, _ := param["required"].(bool)
		if required {
			result[name] = true
		}
	}
	return result
}

func compareSecurity(changes []Change, old, new map[string]any) {
	if old["security"] == nil && new["security"] == nil {
		return
	}
	if changed(old["security"], new["security"]) {
		changes = append(changes, Change{
			Severity: Breaking,
			Type:     "SECURITY_CHANGED",
			Path:     "security",
			Message:  "security requirements changed",
			Before:   old["security"],
			After:    new["security"],
		})
	}
}

func compareMaps(changes []Change, base string, old, new map[string]any) {
	for _, key := range sortedKeys(old) {
		oldValue := old[key]
		newValue, newOK := new[key]
		if !newOK {
			changes = append(changes, Change{
				Severity: Breaking,
				Type:     "FIELD_REMOVED",
				Path:     base + "." + key,
				Message:  fmt.Sprintf("field %q was removed", key),
				Before:   oldValue,
			})
			continue
		}
		oldMap := asMap(oldValue)
		newMap := asMap(newValue)
		if oldMap != nil && newMap != nil {
			compareMaps(changes, base+"."+key, oldMap, newMap)
			continue
		}
		if changed(oldValue, newValue) {
			changes = append(changes, Change{
				Severity: Breaking,
				Type:     "VALUE_CHANGED",
				Path:     base + "." + key,
				Message:  fmt.Sprintf("value changed for %q", key),
				Before:   oldValue,
				After:    newValue,
			})
		}
	}
	for _, key := range sortedKeys(new) {
		if _, ok := old[key]; !ok {
			changes = append(changes, Change{
				Severity: Breaking,
				Type:     "FIELD_ADDED",
				Path:     base + "." + key,
				Message:  fmt.Sprintf("new field %q was added", key),
				After:    new[key],
			})
		}
	}
}

func Sitemap(old, new map[string]any) []Change {
	return diffStringSets("SITEMAP_URL", old, new)
}

func Feed(old, new map[string]any) []Change {
	return diffStringSets("FEED_ENTRY", old, new)
}

func diffStringSets(kind string, old, new map[string]any) []Change {
	oldKeys := topLevelKeys(old)
	newKeys := topLevelKeys(new)
	changes := []Change{}

	for _, key := range sortedKeys(oldKeys) {
		if _, ok := newKeys[key]; !ok {
			changes = append(changes, Change{
				Severity: Warning,
				Type:     kind + "_REMOVED",
				Path:     key,
				Message:  fmt.Sprintf("%q was removed", key),
				Before:   oldKeys[key],
			})
		}
	}
	for _, key := range sortedKeys(newKeys) {
		if _, ok := oldKeys[key]; !ok {
			changes = append(changes, Change{
				Severity: Warning,
				Type:     kind + "_ADDED",
				Path:     key,
				Message:  fmt.Sprintf("%q was added", key),
				After:    newKeys[key],
			})
		}
	}
	for _, key := range sortedKeys(oldKeys) {
		if newValue, ok := newKeys[key]; ok && changed(oldKeys[key], newValue) {
			changes = append(changes, Change{
				Severity: Warning,
				Type:     kind + "_CHANGED",
				Path:     key,
				Message:  fmt.Sprintf("%q changed", key),
				Before:   oldKeys[key],
				After:    newValue,
			})
		}
	}
	return changes
}

func topLevelKeys(value map[string]any) map[string]any {
	result := map[string]any{}
	for key, item := range value {
		result[key] = item
	}
	return result
}

func sortedKeys[T any](value map[string]T) []string {
	keys := make([]string, 0, len(value))
	for key := range value {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
