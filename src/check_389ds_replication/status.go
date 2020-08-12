package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func parseStatus(s []replicationStatus) status {
	var result status
	var st updateStatus

	for _, r := range s {
		err := json.Unmarshal(r.ReplicationJSONStatus, &st)
		if err != nil {
			result.Unknown = append(result.Unknown, err.Error())
			continue
		}
		switch strings.ToLower(st.State) {
		case "green":
			// "Error (0) for a good state can be confusing, remove it instead
			result.Ok = append(result.Ok, fmt.Sprintf("Replication agreement %s with %s: %s", r.ReplicationAgreement, r.ReplicationHost, strings.Replace(st.Message, "Error (0) ", "", 1)))
		case "amber":
			result.Warning = append(result.Warning, fmt.Sprintf("Replication agreement %s with %s: %s", r.ReplicationAgreement, r.ReplicationHost, st.Message))
		case "red":
			result.Critical = append(result.Critical, fmt.Sprintf("Replication agreement %s with %s: %s", r.ReplicationAgreement, r.ReplicationHost, st.Message))
		default:
			// XXX: This should never happen
			result.Unknown = append(result.Unknown, fmt.Sprintf("Unknown replication state %s for replication agreement %s with %s: %s", st.State, r.ReplicationAgreement, r.ReplicationHost, st.Message))
		}
	}

	return result
}
