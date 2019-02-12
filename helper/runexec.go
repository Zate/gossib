/*
MIT License

Copyright (c) 2018 Aris Ripandi

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package helper

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

func RunExec(command string, progress bool) {

	x := exec.Command("bash", "-c", command)

	if progress == true {
		stdoutIn, _ := x.StdoutPipe()
		stderrIn, _ := x.StderrPipe()

		var stdoutBuf, stderrBuf bytes.Buffer
		var errStdout, errStderr error
		stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
		stderr := io.MultiWriter(os.Stderr, &stderrBuf)
		err := x.Start()
		if err != nil {
			log.Fatalf("Error: '%s'\n", err)
		}

		go func() {
			_, errStdout = io.Copy(stdout, stdoutIn)
		}()

		go func() {
			_, errStderr = io.Copy(stderr, stderrIn)
		}()

		err = x.Wait()
		if err != nil {
			log.Fatalf("Error: '%s'\n", err)
		}
		if errStdout != nil || errStderr != nil {
			log.Fatal("failed to capture stdout or stderr\n")
		}
		// outStr := string(stdoutBuf.Bytes())
		// fmt.Printf("out:\n%s\n", outStr)
	} else {
		out, err := x.CombinedOutput()
		if err != nil {
			log.Fatalf("Error: %s\n", err)
		}
		WriteLog(string(out))
	}
}

// Run external command and get the single line value from result
func RunExecV(c string) string {
	x := exec.Command("bash", "-c", c)
	out, err := x.CombinedOutput()
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}
	r := fmt.Sprintf("%s", out)
	z := strings.Replace(r, "\n", "", -1)
	return z
}
