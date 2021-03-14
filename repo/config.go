package repo

import (
	"os/exec"
)

// GlobalUser sets the global git author email.
func GlobalUser(email string) *exec.Cmd {
	cmd := exec.Command(
		"git",
		"config",
		"--global",
		"user.email",
		email)
	return cmd
}

func configGlobalUser(name string, email string) [2]*exec.Cmd {
	var cmd [2]*exec.Cmd
	cmd[0] = exec.Command(
		"git",
		"config",
		"--global",
		"user.email",
		email)
	cmd[1] = exec.Command(
		"git",
		"config",
		"--global",
		"user.name",
		name)
	return cmd
}

// GlobalName sets the global git author name.
func GlobalName(author string) *exec.Cmd {
	cmd := exec.Command(
		"git",
		"config",
		"--global",
		"user.name",
		author)

	return cmd
}
