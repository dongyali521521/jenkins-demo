package main

// Import the fmt for formatting strings
// Import os so we can read environment variables from the system
import (
	"fmt"
	"os"
)

func main() {
<<<<<<< HEAD
<<<<<<< HEAD
	fmt.Println("Hello, Kubernetes！I'm from Jenkins CI！4444")
=======
	fmt.Println("Hello, Kubernetes！I'm from Jenkins CI！66666")
>>>>>>> dev
=======
	fmt.Println("Hello, Kubernetes！I'm from Jenkins CI！8888")
>>>>>>> dev
	fmt.Println("BRANCH_NAME:", os.Getenv("branch"))
}
