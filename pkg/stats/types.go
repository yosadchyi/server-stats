package stats

// ClusterLoad is a struct that holds the cluster address and the load of the cluster.
type ClusterLoad struct {
	Server string `json:"server"`
	Load   int    `json:"load"`
}

// FetchLoadFn is a function that fetches the stats from a server.
type FetchLoadFn func(string) (error, ClusterLoad)

type ClustersStats []ClusterLoad
