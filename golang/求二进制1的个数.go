// Copyright 2018 pfwu. All rights reserved.
//
// @Author: pfwu
// @Date: 2018-04-18 23:44

//任意给定一个32位无符号整数n，求n的二进制表示中1的个数，比如n = 5（0101）时，返回2，n = 15（1111）时，返回4
//对于二进制数来说，除一个2，就少一位，可以判断这个少的位来确定"1"的个数。

package main

import (
	"fmt"
	"os/exec"
	"syscall"
	"os"
	"strings"
)

func count(n byte) {
	IntNum := 0
	for n > 0 {
		if n%2 == 1 {
			IntNum++
		}
		n = n / 2
	}
	fmt.Print(IntNum)
}

func main() {
	count(15)
}

func ommonExec(command []string, ch chan Results) {
	//command := []string{"nohup tail -f /search/odin/nginx/log/access.log &"}
	defaultEnv := map[string]string{
		"PATH": "/usr/local/sbin:/usr/bin:/bin:/usr/sbin:/sbin:/usr/local/bin",
	}
	execCommand := exec.Command("bash")

	execCommand.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	execCommand.Stdout = os.Stdout
	execCommand.Stderr = os.Stderr
	execCommand.Stdin = nil
	execCommand.Env = mergeEnv(defaultEnv, commandArgs.env...)
	execCommand.Args = append([]string{"bash", "-c"}, strings.Join(command, " "))
	gCmds = append(gCmds, execCommand)
	err := execCommand.Run()
	retCode := 0
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		// try to get the exit code
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			retCode = ws.ExitStatus()
		} else {
			retCode = 127
		}
	} else {
		ws := execCommand.ProcessState.Sys().(syscall.WaitStatus)
		retCode = ws.ExitStatus()
	}
	ch <- Results{err: err, rc: retCode}
	return
}
