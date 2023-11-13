package stats

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// FetchLoadOverHTTP fetches the stats from a server over HTTP.
func FetchLoadOverHTTP(address string) (error, ClusterLoad) {
	resp, err := http.Get("http://" + address + "/stats")
	if err != nil {
		return err, ClusterLoad{}
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("error closing response body: %s", err.Error())
		}
	}(resp.Body)

	var clusterStats ClusterLoad
	if err := json.NewDecoder(resp.Body).Decode(&clusterStats); err != nil {
		return err, ClusterLoad{}
	}

	return nil, clusterStats
}

// BuildClusterStats builds a slice of ClusterLoad from a slice of addresses.
func BuildClusterStats(addresses []string, fetchLoad FetchLoadFn) (error, ClustersStats) {
	var clusterStats ClustersStats

	for _, address := range addresses {
		err, clusterStat := fetchLoad(address)

		if err != nil {
			// fail fast, return error
			return err, nil
		}
		clusterStats = append(clusterStats, clusterStat)
	}
	return nil, clusterStats
}
