package finally

import "fmt"

type audit struct {
	message string
}

func (a *audit) setMessage(msg string) {
	a.message = msg
}

func AuditLog() {
	audit := audit{}
	isError := true
	audit.setMessage("success")

	defer func() {
		fmt.Println(audit.message)
	}()

	fmt.Println("start...")
	fmt.Println("process...")

	if isError {
		audit.setMessage("failed")
		return
	}

	fmt.Println("done")
}
