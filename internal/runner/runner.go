package runner

import (
	"fmt"
	"os/exec"
)

func Run() {
	cmd := exec.Command("la", "-al")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Xato:", err)
		return
	}
	fmt.Println(string(output))
}
