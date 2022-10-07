package tool

import (
	"database/sql"
	"embed"
	"fmt"

	"github.com/Jeffail/gabs/v2"
	"github.com/charlieegan3/toolbelt/pkg/apis"
	"github.com/gorilla/mux"

	"github.com/charlieegan3/tool-inoreader-github-actions-trigger/pkg/api"
	"github.com/charlieegan3/tool-inoreader-github-actions-trigger/pkg/tool/handlers"
)

// InoreaderGithubActions is a tool for receiving webhooks from Inoreader and
// triggering Github Actions workflows
type InoreaderGithubActions struct {
	targets map[string]api.Target
}

func (i *InoreaderGithubActions) Name() string {
	return "inoreader-github-actions"
}

func (i *InoreaderGithubActions) FeatureSet() apis.FeatureSet {
	return apis.FeatureSet{
		Config: true,
		HTTP:   true,
	}
}

func (i *InoreaderGithubActions) HTTPPath() string {
	return i.Name()
}

func (i *InoreaderGithubActions) SetConfig(config map[string]any) error {
	cfg := gabs.Wrap(config)

	configTargets, ok := cfg.Path("targets").Data().(map[string]interface{})
	if !ok {
		return fmt.Errorf("missing required config path: targets (array)")
	}

	i.targets = make(map[string]api.Target)
	for k, c := range configTargets {
		targetData, ok := c.(map[string]interface{})
		if !ok {
			return fmt.Errorf("invalid target config, not map[string]interface{}")
		}

		url, ok := targetData["url"].(string)
		if !ok {
			return fmt.Errorf("invalid target config, missing url (string)")
		}
		token, ok := targetData["token"].(string)
		if !ok {
			return fmt.Errorf("invalid target config, missing token (string)")
		}
		eventType, ok := targetData["event_type"].(string)
		if !ok {
			return fmt.Errorf("invalid target config, missing event_type (string)")
		}
		i.targets[k] = api.Target{
			URL:       url,
			Token:     token,
			EventType: eventType,
		}
	}

	return nil
}

func (i *InoreaderGithubActions) DatabaseMigrations() (*embed.FS, string, error) {
	return &embed.FS{}, "migrations", nil
}

func (i *InoreaderGithubActions) DatabaseSet(db *sql.DB) {}

func (i *InoreaderGithubActions) HTTPAttach(router *mux.Router) error {
	router.HandleFunc(
		"/targets/{target}",
		handlers.BuildGetHandler(i.targets),
	).Methods("POST")

	return nil
}

func (i *InoreaderGithubActions) Jobs() ([]apis.Job, error) {
	return []apis.Job{}, nil
}
