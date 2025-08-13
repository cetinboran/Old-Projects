package core

import (
	"log"

	"golang.org/x/crypto/ssh"
)

type ISSH interface {
	Start() error
	GetInfo() string
}

type SSH struct {
	Server   string
	Port     string
	ConnType string
	Config   *ssh.ClientConfig
}

func NewSSH(server, port string) ISSH {
	return &SSH{
		Server:   server,
		Port:     port,
		ConnType: "tcp",
	}
}

func (s *SSH) AddClientConfig(username, password string) {
	clientConfig := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	s.Config = clientConfig
}

func (s *SSH) Start() error {
	client, err := ssh.Dial(s.ConnType, s.Server+":"+s.Port, s.Config)
	if err != nil {
		log.Fatalf("Failed to dial: %s", err)
	}
	defer client.Close()

	// Yeni bir SSH oturumu oluşturma
	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %s", err)
	}
	defer session.Close()

	return nil
}

func (s *SSH) GetInfo() string {
	return s.Server + ":" + s.Port
}
