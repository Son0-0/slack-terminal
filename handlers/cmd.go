package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"time"
)

func ExecCommand(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if HandleError(err) {
		fmt.Fprintln(w, "parse error", http.StatusInternalServerError)
	}

	length := len(r.Form["text"])
	if length != 0 {
		log.Printf("$ %s %s\n", r.Form["user_id"], r.Form["text"][0])

		ctx, cancel := context.WithTimeout(context.Background(), 2000*time.Millisecond)
		defer cancel()

		cmd := exec.CommandContext(ctx, "bash", "-c", r.Form["text"][0])
		stdout, err := cmd.Output()

		if HandleError(err) {
			fmt.Fprintf(w, "*ERROR*\n>msg: `%v`", err)
			return
		}
		fmt.Fprintf(w, "```"+string(stdout)+"```")
	}
}
