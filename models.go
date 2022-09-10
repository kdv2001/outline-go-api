package main

type Key struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Password  string `json:"password,omitempty"`
	Port      int    `json:"port,omitempty"`
	Method    string `json:"method,omitempty"`
	AccessUrl string `json:"accessUrl,omitempty"`
	DataLimit Bytes  `json:"dataLimit,omitempty"`
}

type AccessKeys struct {
	AccessKeys []Key `json:"accessKeys"`
}

type Bytes struct {
	Bytes int `json:"bytes,omitempty"`
}

type Limit struct {
	Limit Bytes `json:"limit,omitempty"`
}

type Server struct {
	Name               string `json:"name,omitempty"`
	ServerId           string `json:"serverId,omitempty"`
	MetricsEnabled     bool   `json:"metricsEnabled,omitempty"`
	CreatedTimestampMs int64  `json:"createdTimestampMs,omitempty"`
	Version            string `json:"version,omitempty"`
	AccessKeyDataLimit struct {
		Bytes int64 `json:"bytes,omitempty"`
	} `json:"accessKeyDataLimit,omitempty"`
	PortForNewAccessKeys  int    `json:"portForNewAccessKeys,omitempty"`
	HostnameForAccessKeys string `json:"hostnameForAccessKeys,omitempty"`
}

type Status struct {
	metricksStatus bool `json:"metricsEnabled,omitempty"`
}
