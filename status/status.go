package status

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/pritunl/pritunl-link/config"
	"github.com/pritunl/pritunl-link/constants"
	"github.com/pritunl/pritunl-link/utils"
	"strings"
	"time"
)

var (
	offlineTime time.Time
	Status      = map[string]map[string]string{}
)

func Update(total int) (err error) {
	connected := 0
	status := map[string]map[string]string{}

	output, err := utils.ExecOutput("", "ipsec", "status")
	if err != nil {
		return
	}

	for _, line := range strings.Split(output, "\n") {
		lines := strings.SplitN(line, ":", 2)
		if len(lines) != 2 {
			continue
		}

		if !strings.HasSuffix(lines[0], "]") {
			continue
		}

		connId := strings.SplitN(strings.SplitN(lines[0], "[", 2)[0], "-", 2)
		connState := strings.SplitN(
			strings.TrimSpace(lines[1]), " ", 2)[0]

		if len(connId) != 2 {
			continue
		}

		switch connState {
		case "ESTABLISHED":
			connState = "connected"
			break
		case "CONNECTING":
			connState = "connecting"
			break
		default:
			connState = "disconnected"
		}

		if _, ok := status[connId[0]]; !ok {
			status[connId[0]] = map[string]string{}
		}

		if _, ok := status[connId[0]][connId[1]]; !ok {
			status[connId[0]][connId[1]] = connState
		}
	}

	for _, stat := range status {
		for _, conn := range stat {
			if conn == "connected" {
				connected += 1
			}
		}
	}

	Status = status

	if connected < total {
		if !offlineTime.IsZero() {
			timeout := constants.DefaultDiconnectedTimeout

			disconnectedTimeout := config.Config.DisconnectedTimeout
			if disconnectedTimeout != 0 {
				timeout = time.Duration(disconnectedTimeout) * time.Second
			}

			if !config.Config.DisableDisconnectedRestart {
				if time.Since(offlineTime) > timeout {
					logrus.Warn("status: Disconnected timeout restarting")

					e := utils.Exec("", "ipsec", "restart")
					if e != nil {
						logrus.WithFields(logrus.Fields{
							"error": e,
						}).Warn("status: Failed to restart")
					}
				}
			} else {
				offlineTime = time.Time{}
			}
		} else {
			offlineTime = time.Now()
		}
	} else {
		offlineTime = time.Time{}
	}

	return
}
