package utils

import (
	"github.com/pkg/sftp"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
)
const (
	SSH_DEFAULT_PORT = "22"
)

func createSshClient(ip string, password string,port string) (*ssh.Client, error) {
	auth := []ssh.AuthMethod{ssh.Password(password)}
	addr := fmt.Sprintf("%s:%d", ip, port)
	config := &ssh.ClientConfig{
		User:            "root",
		Auth:            auth,
		Timeout:         30 * time.Second,
		HostKeyCallback: ssh.HostKeyCallback(func(hostname string, remote net.Addr, key ssh.PublicKey) error { return nil }),
	}
	return ssh.Dial("tcp", addr, config)
}

func createSftpClient(ip string, password string) (*sftp.Client, error) {
	sshClient, err := createSshClient(ip, password,SSH_DEFAULT_PORT)
	if err != nil {
		return nil, err
	}
	return sftp.NewClient(sshClient)
}

func copyFileToRemoteHost(ip string, password string, localFile string, remoteFile string) error {
	client, err := createSftpClient(ip, password)
	if err != nil {
		return err
	}
	defer client.Close()

	srcFile, err := os.Open(localFile)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := client.Create(remoteFile)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	buf := make([]byte, 1024)
	for {
		n, _ := srcFile.Read(buf)
		if n == 1024 {
			dstFile.Write(buf)
		} else {
			if n > 0 {
				dstFile.Write(buf[0:n])
			}
			break
		}
	}

	return nil
}

func runRemoteHostScript(ip string, password string, remoteFile string) (string, error) {
	client, err := createSshClient(ip, password,SSH_DEFAULT_PORT)
	if err != nil {
		return "", err
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	var stdout, stderr bytes.Buffer
	session.Stdout = &stdout
	session.Stderr = &stderr
	if err := session.Run(remoteFile); err != nil {
		logrus.Errorf("runRemoteHostScript stdout=%v,stderr=%v\n", stdout, stderr)
		return "", err
	}
	return stdout.String(), nil
}