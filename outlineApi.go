package main

type OutlineVpnMethods interface {
	GetKeys() (AccessKeys, error)
	CreateKey(keyName string) (Key, error)
	DeleteKey(keyId string) error
	RenameKey(keyId string, keyName string) error
	AddDataLimit(keyId string, byteLimit int) error
	DeleteDataLimit(keeId string) error
	GetTransferredData() error
	GetServerInformation() (Server, error)
	SetServerName(serverName string) error
	SetHostName(hostName string) error
	GetMetricsStatus() (bool, error)
	SetPortNewForAccessKeys(port string) error
	SetDataLimitForAllKeys(bytesLimit int) error
	DeleteDataLimitForAllKeys() error
}
