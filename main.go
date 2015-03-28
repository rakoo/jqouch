package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var mappers = make([]mapper, 0)

type mapper struct {
	in  io.Writer
	out *bufio.Reader
}

var jsonOut = json.NewEncoder(os.Stdout)

func main() {
	log("Starting")
	lineReader := bufio.NewReader(os.Stdin)

	// We read data that looks like ["cmd", <blabla>]
	for {
		line, err := lineReader.ReadString('\n')
		if err != nil {
			logerror("error_readinput", fmt.Sprintf("Error reading input:%s", err))
		}

		parts := strings.SplitN(line, ",", 2)
		cmd := strings.Trim(parts[0], "[\" ")
		args := strings.Trim(parts[1], "] \n")

		switch cmd {
		case "reset":
			jsonOut.Encode(true)
		case "add_fun":
			filter, err := strconv.Unquote(args)
			if err != nil {
				logerror("invalid_function", "This function is invalid:"+err.Error())
				continue
			}
			jqCmd := exec.Command("jq", "-c", "--unbuffered", "-r", filter)
			in, err := jqCmd.StdinPipe()
			if err != nil {
				logerror("error_installing_jq", fmt.Sprintf("Error installing jq: %s", err))
				return
			}
			out, err := jqCmd.StdoutPipe()
			if err != nil {
				logerror("error_installing_jq", fmt.Sprintf("Error installing jq: %s", err))
				return
			}
			stdErr, err := jqCmd.StderrPipe()
			if err != nil {
				logerror("error_installing_jq", fmt.Sprintf("Error installing jq: %s", err))
				return
			}

			err = jqCmd.Start()
			if err != nil {
				logerror("error_installing_jq", fmt.Sprintf("Error installing jq: %s", err))
				return
			}

			mappers = append(mappers, mapper{in, bufio.NewReader(out)})
			go func() {
				for {
					rd := bufio.NewScanner(stdErr)
					for rd.Scan() {
						err := rd.Text()
						logerror("jq_process", fmt.Sprintf("Error in jq subprocess: %s", err))
					}
				}
			}()

			log("Installed jq")

			jsonOut.Encode(true)
		case "map_doc":

			var buf bytes.Buffer
			buf.WriteString("[")

		allmappers:
			for i, m := range mappers {
				io.WriteString(m.in, args+"\n")

				for {
					line, isPrefix, err := m.out.ReadLine()
					if err != nil {
						logerror("error_calling_jq", fmt.Sprintf("Error calling jq: %s", err))
						continue allmappers
					}
					buf.Write(line)
					if !isPrefix {
						break
					}
				}
				if i != len(mappers)-1 {
					buf.WriteString(",")
				}
			}

			io.Copy(os.Stdout, &buf)
			io.WriteString(os.Stdout, "]\n")

		default:
			logerror("unknown_command", fmt.Sprintf("%s is unknown", cmd))
			log(fmt.Sprintf("unknown command: %s", cmd))
		}
	}
}

func logerror(typ, message string) {
	fmt.Fprintln(os.Stdout, `["error", "`+typ+`", "`+message+`"]`)
}

func log(message string) {
	fmt.Fprintln(os.Stdout, `["log", "`+message+`"]`)
}
