package stats_test

import (
	"awesomeProject2/pkg/stats"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleServersStatsRequest(t *testing.T) {
	addresses := []string{"1.1.1.1", "2.2.2.2", "3.3.3.3"}

	req, err := http.NewRequest("GET", "/servers/stat", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	i := 0
	handler := stats.HandleServersStatsRequest(addresses, func(address string) (error, stats.ClusterLoad) {
		i++
		clusterStats := stats.ClusterLoad{Server: address, Load: rand.Intn(i)}
		return nil, clusterStats
	})
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"server":"1.1.1.1","load":0},{"server":"2.2.2.2","load":1},{"server":"3.3.3.3","load":2}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
