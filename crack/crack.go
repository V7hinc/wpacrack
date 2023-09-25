package crack

import (
	"bytes"
	"fmt"
	"os/exec"
)

var aircrackPath = "aircrack-ng"

func Crack(passwordFile, capFile string) (string, bool, error) {
	foundKey := false
	cmd := exec.Command(aircrackPath, "-w", passwordFile, capFile)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", false, fmt.Errorf("cmd.Run() failed with %s\n", err)
	}
	hasHandShake := false
	ssid := ""
	key := ""
	for _, line := range bytes.Split(out, []byte("\n")) {
		line = bytes.TrimSpace(line)
		if !hasHandShake && bytes.Contains(line, []byte("handshake")) {
			hasHandShake = true
			ssid = string(line)
		}

		if bytes.Contains(line, []byte("KEY FOUND!")) {
			foundKey = true
			keySplit := bytes.SplitN(line, []byte("KEY FOUND!"), 2)
			if len(keySplit) == 2 {
				key = string(keySplit[1])
			} else {
				key = string(line)
			}
			break
		}
	}
	impMsg := ""
	if hasHandShake {
		message := fmt.Sprintf("%s -- %s \n", ssid, key)
		impMsg = message
	} else {
		err = fmt.Errorf("no hasHandShake")
	}
	return impMsg, foundKey, err

}
