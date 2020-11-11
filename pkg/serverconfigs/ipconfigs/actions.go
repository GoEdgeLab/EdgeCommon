package ipconfigs

type IPListAction = string

const (
	IPListActionHTTP  IPListAction = "http"  // HTTP
	IPListActionIPSet IPListAction = "ipset" // ipset
)

