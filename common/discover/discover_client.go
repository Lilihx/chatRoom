package discover

import (
	"strconv"
	"sync"

	"github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
	"github.com/lilihx/chatRoom/common/log"
)

type DiscoveryClient interface {
	Register(serviceName, instanceId, healthCheckUrl string, instanceHost string, instancePort int, meta map[string]string) bool

	DeRegister(instanceId string) bool

	DiscoverServices(serviceName string) []interface{}
}

type KitDiscoverClient struct {
	Host   string
	Port   int
	client consul.Client
	// 连接 consul 的配置
	config      *api.Config
	mutex       sync.Mutex
	instanceMap sync.Map
}

func NewKitDiscoverClient(consulHost string, consulPort int) (DiscoveryClient, error) {
	consulConfig := api.DefaultConfig()
	consulConfig.Address = consulHost + ":" + strconv.Itoa(consulPort)
	apiClient, err := api.NewClient(consulConfig)
	if err != nil {
		log.Error(("init consul err" + err.Error()))
		return nil, err
	}
	client := consul.NewClient(apiClient)
	return &KitDiscoverClient{
		Host:   consulHost,
		Port:   consulPort,
		client: client,
		config: consulConfig,
	}, err
}

func (consulClient *KitDiscoverClient) Register(serviceName, instanceId, healthCheckUrl string, instanceHost string, instancePort int, meta map[string]string) bool {
	serviceRegistration := &api.AgentServiceRegistration{
		ID:      instanceId,
		Name:    serviceName,
		Address: instanceHost,
		Port:    instancePort,
		Meta:    meta,
		Check: &api.AgentServiceCheck{
			DeregisterCriticalServiceAfter: "3000s",
			HTTP:                           "http://host.docker.internal" + ":" + strconv.Itoa(instancePort) + healthCheckUrl,
			Interval:                       "15s",
		},
	}
	err := consulClient.client.Register(serviceRegistration)
	if err != nil {
		log.Error("Register service error")
		return false
	}
	log.Info("Register service success")
	return true
}

func (consulClient *KitDiscoverClient) DeRegister(instanceId string) bool {
	serviceRegistration := &api.AgentServiceRegistration{
		ID: instanceId,
	}
	err := consulClient.client.Deregister(serviceRegistration)
	if err != nil {
		log.Error("Deregister service error")
		return false
	}
	log.Info("Deregister service success")
	return true
}

func (consulClient *KitDiscoverClient) DiscoverServices(serviceName string) []interface{} {
	instanceList, ok := consulClient.instanceMap.Load(serviceName)
	if ok {
		return instanceList.([]interface{})
	}
	consulClient.mutex.Lock()
	defer consulClient.mutex.Unlock()
	instanceList, ok = consulClient.instanceMap.Load(serviceName)
	if ok {
		return instanceList.([]interface{})
	}
	entries, _, err := consulClient.client.Service(serviceName, "", false, nil)
	if err != nil {
		log.Error("Discover error")
		return nil
	}
	instances := make([]interface{}, len(entries))
	for i := 0; i < len(entries); i++ {
		instances[i] = entries[i].Service
	}
	consulClient.instanceMap.Store(serviceName, instances)
	return instances
}
