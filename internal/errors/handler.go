/*
Copyright (C) 2021-2023, Kubefirst

This program is licensed under MIT.
See the LICENSE file for more details.
*/
package errors

import (
	"github.com/kubefirst/kubefirst-api/internal/db"
	"github.com/kubefirst/kubefirst-api/internal/types"
)

// HandleClusterError implements an error handler for standalone cluster objects
func HandleClusterError(cl *types.Cluster, condition string) error {
	err := db.Client.UpdateCluster(cl.ClusterName, "in_progress", false)
	if err != nil {
		return err
	}
	err = db.Client.UpdateCluster(cl.ClusterName, "status", "error")
	if err != nil {
		return err
	}
	err = db.Client.UpdateCluster(cl.ClusterName, "last_condition", condition)
	if err != nil {
		return err
	}

	return nil
}