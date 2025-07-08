package config

import (
	"os"
	"strconv"
	"time"

	"github.com/ArthurHlt/go-eureka-client/eureka"
)

type EurekaClient struct {
	Client   *eureka.Client
	Instance *eureka.InstanceInfo
	AppName  string
	HostName string
	Port     int
}

// NewEurekaClient cria e configura o cliente Eureka
func NewEurekaClient() *EurekaClient {
	eurekaURL := os.Getenv("EUREKA_URL")
	appName := os.Getenv("APP_NAME")
	host := os.Getenv("APP_HOST")
	portStr := os.Getenv("APP_PORT")
	port, _ := strconv.Atoi(portStr)

	client := eureka.NewClient([]string{eurekaURL})

	instance := eureka.NewInstanceInfo(
		host,
		appName,
		host,
		port,
		30,
		false,
	)

	instance.Metadata = &eureka.MetaData{
		Map: map[string]string{
			"language": "go",
		},
	}

	return &EurekaClient{
		Client:   client,
		Instance: instance,
		AppName:  appName,
		HostName: host,
		Port:     port,
	}
}

// Register registra o serviço no Eureka
func (ec *EurekaClient) Register() error {
	err := ec.Client.RegisterInstance(ec.AppName, ec.Instance)
	if err != nil {
		return err
	}

	return nil
}

func (ec *EurekaClient) StartHeartbeat() {
	go func() {
		for {
			time.Sleep(30 * time.Second)
			ec.Client.SendHeartbeat(ec.Instance.App, ec.Instance.HostName)
		}
	}()
}
