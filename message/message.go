package message

import (
	"fmt"
)

func JobSuccess(clusterName, jobName string, timeSinceCompletion float64) string {
	return "*" + clusterName + ":" + jobName + "* succeeded " + fmt.Sprintf("%.1f", timeSinceCompletion) + " minutes ago :tada:"
}

func JobFailure(clusterName, jobName string, message string) string {
	if len(message) > 0 {
		return "*" + clusterName + ":" + jobName + "* " + message
	} else {
		return "*" + clusterName + ":" + jobName + "* failed :alert:"
	}
}
