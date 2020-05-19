// Copyright 2015 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package elasticsearch

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"

	info "github.com/google/cadvisor/info/v1"
	storage "github.com/google/cadvisor/storage"
	"github.com/pkg/errors"

	"github.com/olivere/elastic/v7"
)

func init() {
	storage.RegisterStorageDriver("elasticsearch", new)
}

type elasticStorage struct {
	client      *elastic.Client
	machineName string
	indexName   string
	typeName    string
	lock        sync.Mutex
}

type detailSpec struct {
	Timestamp      int64                `json:"timestamp"`
	MachineName    string               `json:"machine_name,omitempty"`
	ContainerName  string               `json:"container_Name,omitempty"`
	ContainerStats *info.ContainerStats `json:"container_stats,omitempty"`
}

var (
	argElasticHost   = flag.String("storage_driver_es_host", "http://localhost:9200", "ElasticSearch host:port")
	argIndexName     = flag.String("storage_driver_es_index", "cadvisor", "ElasticSearch index name")
	argTypeName      = flag.String("storage_driver_es_type", "stats", "ElasticSearch type name")
	argEnableSniffer = flag.Bool("storage_driver_es_enable_sniffer", false, "ElasticSearch uses a sniffing process to find all nodes of your cluster by default, automatically")
	argElasticCA     = flag.String("storage_driver_es_ca", "", "ElasticSearch tls ca")
)

func new() (storage.StorageDriver, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	username := os.Getenv("ES_USERNAME")
	password := os.Getenv("ES_PASSWORD")

	return newStorage(
		hostname,
		*argIndexName,
		*argTypeName,
		*argElasticHost,
		*argEnableSniffer,
		*argElasticCA,
		username,
		password,
	)
}

func (self *elasticStorage) containerStatsAndDefaultValues(
	cInfo *info.ContainerInfo, stats *info.ContainerStats) *detailSpec {
	timestamp := stats.Timestamp.UnixNano() / 1E3
	var containerName string
	if len(cInfo.ContainerReference.Aliases) > 0 {
		containerName = cInfo.ContainerReference.Aliases[0]
	} else {
		containerName = cInfo.ContainerReference.Name
	}
	detail := &detailSpec{
		Timestamp:      timestamp,
		MachineName:    self.machineName,
		ContainerName:  containerName,
		ContainerStats: stats,
	}
	return detail
}

func (self *elasticStorage) AddStats(cInfo *info.ContainerInfo, stats *info.ContainerStats) error {
	if stats == nil {
		return nil
	}
	func() {
		// AddStats will be invoked simultaneously from multiple threads and only one of them will perform a write.
		self.lock.Lock()
		defer self.lock.Unlock()
		// Add some default params based on ContainerStats
		detail := self.containerStatsAndDefaultValues(cInfo, stats)
		// Index a cadvisor (using JSON serialization)
		_, err := self.client.Index().
			Index(self.indexName).
			Type(self.typeName).
			BodyJson(detail).
			Do(context.Background())
		if err != nil {
			// Handle error
			fmt.Printf("failed to write stats to ElasticSearch - %s", err)
			return
		}
	}()
	return nil
}

func (self *elasticStorage) Close() error {
	self.client = nil
	return nil
}

// machineName: A unique identifier to identify the host that current cAdvisor
// instance is running on.
// ElasticHost: The host which runs ElasticSearch.
func newStorage(
	machineName,
	indexName,
	typeName,
	elasticHost string,
	enableSniffer bool,
	caPath string,
	username string,
	password string,
) (storage.StorageDriver, error) {
	httpClient, err := createHTTPClient(caPath)
	if err != nil {
		return nil, err
	}

	// Obtain a client and connect to the default Elasticsearch installation
	// on 127.0.0.1:9200. Of course you can configure your client to connect
	// to other hosts and configure it in various other ways.
	client, err := elastic.NewClient(
		elastic.SetHealthcheck(true),
		elastic.SetSniff(enableSniffer),
		elastic.SetHealthcheckInterval(30*time.Second),
		elastic.SetURL(elasticHost),
		elastic.SetHttpClient(httpClient),
		elastic.SetBasicAuth(username, password),
	)
	if err != nil {
		// Handle error
		return nil, fmt.Errorf("failed to create the elasticsearch client - %s", err)
	}

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping(elasticHost).Do(context.Background())
	if err != nil {
		// Handle error
		return nil, fmt.Errorf("failed to ping the elasticsearch - %s", err)

	}
	fmt.Printf("Elasticsearch returned with code %d and version %s", code, info.Version.Number)

	ret := &elasticStorage{
		client:      client,
		machineName: machineName,
		indexName:   indexName,
		typeName:    typeName,
	}
	return ret, nil
}

func createHTTPClient(caPath string) (*http.Client, error) {
	if caPath == "" {
		return http.DefaultClient, nil
	}

	caCert, err := ioutil.ReadFile(caPath)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not read Elasticsearch CA from '%s'", caPath)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
		},
	}
	return httpClient, nil
}
