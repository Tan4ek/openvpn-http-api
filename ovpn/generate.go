package ovpn

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func GenerateClientConfig(clientId string) (string, error) {
	var out bytes.Buffer

	cmd := exec.Command("ovpn_getclient", clientId)

	cmd.Stdout = &out
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("unable to find client \"%s\", please try again or generate the key first", clientId)
	}

	return out.String(), nil
}

func GenerateClientCerts(clientId string, password string) error {
	cmd := exec.Command("easyrsa", "--passin=file:capassfile", "build-client-full", clientId)

	var errBuf bytes.Buffer

	cmd.Stdout = os.Stdout
	cmd.Stderr = &errBuf

	buffer := bytes.Buffer{}
	buffer.Write([]byte(fmt.Sprintf("%s\n%s\n", password, password)))
	cmd.Stdin = &buffer

	err := cmd.Run()

	if err != nil {
		log.Printf("client \"%s\" certs generation command failed with %s ", clientId, err)
	}

	return errors.New(errBuf.String())
}
