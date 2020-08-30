package internal

import (
	"fmt"
	"github.com/ackerr/lab/utils"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"io/ioutil"
	"os"

	"github.com/go-git/go-git/v5"
)

// Clone : git clone the gitlab project
func Clone(gitURL, path string, useHTTPS bool) {
	var auth transport.AuthMethod
	if !useHTTPS {
		sshPath := fmt.Sprintf("%s/.ssh/id_rsa", os.Getenv("HOME"))
		keyFile := utils.GetEnv("PUBLIC_KEY_PATH", sshPath)
		sshKey, err := ioutil.ReadFile(keyFile)
		utils.Check(err)
		auth, err = ssh.NewPublicKeys("git", sshKey, "")
		utils.Check(err)
	}
	_, err := git.PlainClone(path, false, &git.CloneOptions{
		Auth:     auth,
		URL:      gitURL,
		Progress: os.Stdout,
	})
	utils.Check(err)
}
