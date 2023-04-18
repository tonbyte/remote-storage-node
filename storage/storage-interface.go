package storage

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/tonbyte/remote-storage-node/config"

	"github.com/labstack/gommon/log"
)

const bagIDString string = "BagID = "
const duplicateHash string = "duplicate hash "

func execStorageCliCommand(command string) ([]byte, error) {
	cmd := exec.Command(config.StorageConfig.SPCliPath,
		`-I`, `127.0.0.1:`+strconv.Itoa(config.StorageConfig.SPCliPort),
		`-k`, config.StorageConfig.StorageDBPath+`/cli-keys/client`,
		`-p`, config.StorageConfig.StorageDBPath+`/cli-keys/server.pub`,
		`-c`, command)

	return cmd.CombinedOutput()
}

func AddBag(bagID string) bool {
	cliQuery := fmt.Sprintf(`"add-by-hash" "%s"`, bagID)
	output, err := execStorageCliCommand(cliQuery)
	fmt.Println(string(output))

	if err != nil {
		log.Warn(fmt.Sprintf("ExecQuery() error: %s\noutput: %s", err.Error(), output))
	}

	return parseCreateBagOutput(string(output)) == bagID
}

func RemoveBag(bagID string) bool {
	cliQuery := fmt.Sprintf(`"remove" "%s"`, bagID)
	output, err := execStorageCliCommand(cliQuery)
	fmt.Println(string(output))

	if err != nil {
		log.Warn(fmt.Sprintf("ExecQuery() error: %s\noutput: %s", err.Error(), output))
	}

	return parseRemoveBagOutput(string(output))
}

func parseRemoveBagOutput(output string) bool {
	return strings.Contains(output, "No such torrent") || strings.Contains(output, "Success")
}

func parseCreateBagOutput(output string) string {
	bagIdBegin := strings.Index(output, bagIDString)
	if bagIdBegin != -1 {
		bagIdBegin += len(bagIDString)
	} else {
		bagIdBegin = strings.Index(output, duplicateHash) + len(duplicateHash)
	}

	if bagIdBegin == -1 {
		return ""
	}

	return output[bagIdBegin : bagIdBegin+64]
}
