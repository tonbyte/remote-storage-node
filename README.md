# remote-storage-node

# Build:

`git clone https://github.com/tonbyte/remote-storage-node.git`

`cd remote-storage-node`

Install latest golang if need: https://go.dev/doc/install

`go build .`

# Setup

Edit config.json according to your environment. 

```
{
	"sp_cli_path": "home/tonbyte/.../storage-daemon-cli",
	"sp_cli_port": 5555,
	"storage_db_path": "path to storage-db folder",
	"port": 34312,
	"whitelist_ip": ["195.133.147.156", "47.87.160.118", "46.17.248.81"]
}
```

Where:

sp_cli_path - path to storage-daemon-cli

storage_db_path - path to storage-db folder

sp_cli_port - storage-daemon-cli port. Usualy 5555

whitelist_ip - do not remove this field. Add your IPs to this list if need

# Run:

`./remote-storage-node > log.txt &`

# Usage:

To check if node is working, send request to http://YOUR_IP:34312/status

To add file to storage, send request to http://YOUR_IP:34312/addBag/BAG_ID

To remove file from storage, send request to http://YOUR_IP:34312/removeBag/BAG_ID
