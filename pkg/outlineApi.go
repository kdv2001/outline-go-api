package pkg

type OutlineVpnMethods interface {
	GetKeys()
	CreateKey()
	DeleteKey()
	RenameKey()
	AddDataLimit()
	DeleteDataLimit()
	GetTransferredData()
	GetServerInformation()
	SetServerName()
	SetHostName()
	GetMetricsStatus()
	SetPortNewForAccessKeys()
	SetDataLimitForAllKeys()
	DeleteDataLimitForAllKeys()
}
