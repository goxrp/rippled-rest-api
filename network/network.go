package network

const RipplePublicServer1JsonRpcURL = "https://s1.ripple.com:51234/"

var servers = []Server{
	{Hostname: "s1.ripple.com", Description: "Mainnet Public Cluster", JsonRpcPort: 51234},
	{Hostname: "xrpl.ws", Description: "Mainnet Full History Cluster"},
	{Hostname: "s2.ripple.com", Description: "Mainnet Full History Cluster"},
	{Hostname: "s.altnet.rippletest.net", Description: "Testnet Public Cluster"},
	{Hostname: "s.devnet.rippletest.net", Description: "Devnet Public Cluster"},
}

type Server struct {
	Hostname    string `json:"hostname"`
	Description string `json:"description"`
	JsonRpcPort int    `json:"jsonRpcPort"`
}

/*
s1.ripple.com (Mainnet Public Cluster)
xrpl.ws (Mainnet Full History Cluster)
s2.ripple.com (Mainnet Full History Cluster)
s.altnet.rippletest.net (Testnet Public Cluster)
s.devnet.rippletest.net (Devnet Public Cluster)
localhost:6006 (Local rippled Server on port 6006)
*/
