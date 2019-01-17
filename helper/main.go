package helper

import (
	"os/exec"
	"log"
	"encoding/json"
	"errors"
)

func RunShell(command string) ([]byte, error) {
	//cmd := exec.Command("/bin/bash", "-c", `ps -eaf|grep "nginx: master"|grep -v "grep"|awk '{print $2}'`)
	cmd := exec.Command("/bin/bash", "-c", command)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
		return nil, errors.New(string(out))
	}
	log.Println(string(out))
	return out, nil;
}

func JsonResp(status bool, msg string, data interface{}) []byte {
	bytes, _ := json.Marshal(struct {
		Status bool
		Msg    string
		Data   interface{}
	}{Status: status, Msg: msg, Data: data})
	return bytes
}
