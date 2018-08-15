// Copyright 2017-2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const (
	v2MetadataEndpoint     = "http://169.254.170.2/v3/metadata"
	v2StatsEndpoint        = "http://169.254.170.2/v2/stats"
	maxRetries             = 4
	durationBetweenRetries = time.Second
)

// TaskResponse defines the schema for the task response JSON object
type TaskResponse struct {
	Cluster       string
	TaskARN       string
	Family        string
	Revision      string
	DesiredStatus string `json:",omitempty"`
	KnownStatus   string
	Containers    []ContainerResponse `json:",omitempty"`
	Limits        LimitsResponse      `json:",omitempty"`
}

// ContainerResponse defines the schema for the container response
// JSON object
type ContainerResponse struct {
	ID            string `json:"DockerId"`
	Name          string
	DockerName    string
	Image         string
	ImageID       string
	Ports         []PortResponse    `json:",omitempty"`
	Labels        map[string]string `json:",omitempty"`
	DesiredStatus string
	KnownStatus   string
	ExitCode      *int `json:",omitempty"`
	Limits        LimitsResponse
	CreatedAt     *time.Time `json:",omitempty"`
	StartedAt     *time.Time `json:",omitempty"`
	FinishedAt    *time.Time `json:",omitempty"`
	Type          string
	Health        HealthStatus `json:"health,omitempty"`
	Networks      []Network    `json:",omitempty"`
}

type HealthStatus struct {
	Status   string     `json:"status,omitempty"`
	Since    *time.Time `json:"statusSince,omitempty"`
	ExitCode int        `json:"exitCode,omitempty"`
	Output   string     `json:"output,omitempty"`
}

// LimitsResponse defines the schema for task/cpu limits response
// JSON object
type LimitsResponse struct {
	CPU    uint
	Memory uint
}

// PortResponse defines the schema for portmapping response JSON
// object
type PortResponse struct {
	ContainerPort uint16
	Protocol      string
	HostPort      uint16 `json:",omitempty"`
}

// Network is a struct that keeps track of metadata of a network interface
type Network struct {
	NetworkMode   string   `json:"NetworkMode,omitempty"`
	IPv4Addresses []string `json:"IPv4Addresses,omitempty"`
	IPv6Addresses []string `json:"IPv6Addresses,omitempty"`
}

func containerMetadata(client *http.Client, endpoint string) (*TaskResponse, error) {
	body, err := metadataResponse(client, endpoint, "container metadata")
	if err != nil {
		return nil, err
	}

	fmt.Printf("Received task metadata: %s \n", string(body))

	var taskMetadata TaskResponse
	err = json.Unmarshal(body, &taskMetadata)
	if err != nil {
		return nil, fmt.Errorf("task metadata: unable to parse response body: %v", err)
	}

	return &taskMetadata, nil
}

func metadataResponse(client *http.Client, endpoint string, respType string) ([]byte, error) {
	var resp []byte
	var err error
	for i := 0; i < maxRetries; i++ {
		resp, err = metadataResponseOnce(client, endpoint, respType)
		if err == nil {
			return resp, nil
		}
		fmt.Fprintf(os.Stderr, "Attempt [%d/%d]: unable to get metadata response for '%s' from '%s': %v",
			i, maxRetries, respType, endpoint, err)
		time.Sleep(durationBetweenRetries)
	}

	return nil, err
}

func metadataResponseOnce(client *http.Client, endpoint string, respType string) ([]byte, error) {
	resp, err := client.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("%s: unable to get response: %v", respType, err)
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s: incorrect status code  %d", respType, resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("task metadata: unable to read response body: %v", err)
	}

	return body, nil
}

func main() {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	endpoint := os.Getenv("FOO")

	fmt.Println("endpoint:" + endpoint)

	containerMetadata(client, endpoint)

	// Wait for the Health information to be ready
	time.Sleep(5 * time.Second)

	os.Exit(42)
}
