package main

import (
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

var username = "user"
var host = "examplehost.com:22"
var privateKeyFile = "/home/user/.ssh/id_rsa"
var commandToExecute = "cat /etc/passwd"

func getKeySigner(privateKeyFile string) ssh.Signer {
	privateKeyData, err := ioutil.ReadFile(privateKeyFile)
	if err != nil {
		log.Fatal("Error loading private key file. ", err)
	}
	privateKey, err := ssh.ParsePrivateKey(privateKeyData)
	if err != nil {
		log.Fatal("Error parsing private key.", err)
	}
	return privateKey
}

func main() {
	privateKey := getKeySigner(privateKeyFile)
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(privateKey),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", host, config)
	if err != nil {
		log.Fatal("Error dialing server. ", err)
	}
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session:", err)
	}
	defer session.Close()
	session.Stdout = os.Stdout

	err = session.Run(commandToExecute)
	if err != nil {
		log.Fatal("Error executing command.", err)
	}
}
