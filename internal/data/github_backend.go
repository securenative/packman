package data

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/go-github/v25/github"
	"golang.org/x/oauth2"
	"os/exec"
	"strings"
)

type GithubConfig struct {
	Username    string
	Token       string
	PrivatePush bool
}

type GithubBackend struct {
	cfg          *GithubConfig
	configLoader ConfigStore
	client       *github.Client
}

func NewGithubBackend(configLoader ConfigStore) *GithubBackend {
	return &GithubBackend{configLoader: configLoader}
}

func (this *GithubBackend) Push(name string, source string) error {
	gh, cfg, err := this.loadClient()
	if err != nil {
		return err
	}

	splitName := strings.Split(name, "/")
	if len(splitName) != 2 {
		return errors.New("github repository name should be formatted as <owner>/<repo name>")
	}

	err = this.getOrCreateRepository(gh, splitName, cfg)

	addRemote := exec.Command("git", "remote", "add", "origin", githubUrl(name))
	addRemote.Dir = source
	err = addRemote.Run()
	if err != nil {
		return err
	}

	push := exec.Command("git", "push", "-u", "origin", "master")
	push.Dir = source
	err = push.Run()
	return err
}

func (this *GithubBackend) Pull(name string, destination string) error {
	url := githubUrl(name)
	cmd := exec.Command("git", "clone", url, destination)
	return cmd.Run()
}

func (this *GithubBackend) ConfigKey() string {
	return "github"
}

func (this *GithubBackend) getOrCreateRepository(gh *github.Client, splitName []string, cfg *GithubConfig) error {

	_, _, err := gh.Repositories.Get(context.Background(), splitName[0], splitName[1])
	if err == nil {
		err = this.createRepository(splitName, cfg, gh)
	}

	return err
}

func (this *GithubBackend) createRepository(splitName []string, cfg *GithubConfig, gh *github.Client) error {
	repo := &github.Repository{
		Name:    &splitName[1],
		Private: &cfg.PrivatePush,
	}
	_, _, err := gh.Repositories.Create(context.Background(), splitName[0], repo)
	return err
}

func githubUrl(name string) string {
	url := fmt.Sprintf("https://github.com/%s.git", name)
	return url
}

func (this *GithubBackend) loadClient() (*github.Client, *GithubConfig, error) {
	if this.client != nil {
		cfg, err := this.loadConfig()
		if err != nil {
			return nil, nil, err
		}

		ctx := context.Background()
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: cfg.Token})
		tc := oauth2.NewClient(ctx, ts)
		this.client = github.NewClient(tc)
	}

	return this.client, this.cfg, nil
}

func (this *GithubBackend) loadConfig() (*GithubConfig, error) {
	if this.cfg == nil {
		var temp GithubConfig
		found := this.configLoader.Get(this.ConfigKey(), &temp)
		if !found {
			return nil, errors.New("cannot find github configuration")
		}
		this.cfg = &temp
	}

	return this.cfg, nil
}
