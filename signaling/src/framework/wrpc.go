package framework

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"signaling/src/framework/wrpc"
	"strconv"
	"strings"
	"time"
)

// Client communicate with wrtcserver
var wrpcClients map[string]*wrpc.Client = make(map[string]*wrpc.Client)

func loadWrpc() error {
	sections := configFile.GetSectionList()
	for _, section := range sections {
		if !strings.HasPrefix(section, "wrpc.") {
			continue
		}

		mSection, err := configFile.GetSection(section)
		if err != nil {
			return err
		}

		values, ok := mSection["server"]
		if !ok {
			return errors.New("no server field in config file")
		}

		arrServer := strings.Split(values, ",")
		for i, server := range arrServer {
			arrServer[i] = strings.TrimSpace(server)
		}

		client := wrpc.NewClient(arrServer)

		if values, ok := mSection["connectTimeout"]; ok {
			if connectTimeout, err := strconv.Atoi(values); err == nil {
				client.ConnectTimeout = time.Duration(connectTimeout) * time.Millisecond
			}
		}

		if values, ok := mSection["readTimeout"]; ok {
			if readTimeout, err := strconv.Atoi(values); err == nil {
				client.ReadTimeout = time.Duration(readTimeout) * time.Millisecond
			}
		}
		if values, ok := mSection["writeTimeout"]; ok {
			if writeTimeout, err := strconv.Atoi(values); err == nil {
				client.WriteTimeout = time.Duration(writeTimeout) * time.Millisecond
			}
		}

		wrpcClients[section] = client
	}

	return nil
}

func Call(serviceName string, request interface{}, response interface{}, logId uint32) error {
	fmt.Println("call " + serviceName)

	client, ok := wrpcClients["wrpc."+serviceName]
	if !ok {
		return fmt.Errorf("[%s] service not found", serviceName)
	}

	content, err := json.Marshal(request)
	if err != nil {
		return err
	}

	req := wrpc.NewRequest(bytes.NewReader(content), logId)
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	fmt.Println(resp)

	return nil
}
