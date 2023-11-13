package stats

import (
	"encoding/json"
	"log"
	"net/http"
)

// HandleServersStatsRequest handles the request for the cluster stats.
func HandleServersStatsRequest(addresses []string, fetchLoad FetchLoadFn) http.HandlerFunc {
	return func(w http.ResponseWriter, rq *http.Request) {
		err, clusterStats := BuildClusterStats(addresses, fetchLoad)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		statsJson, err := json.Marshal(clusterStats)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, err = w.Write(statsJson)
		if err != nil {
			log.Printf("internal error: %s", err.Error())
			return
		}
	}
}
