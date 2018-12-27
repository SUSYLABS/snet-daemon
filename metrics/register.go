// Copyright 2018 singularitynet foundation.
// All rights reserved.
// <<add licence terms for code reuse>>

// package for monitoring and reporting the daemon metrics
package metrics

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/singnet/snet-daemon/config"
	log "github.com/sirupsen/logrus"
)

// generates DaemonID nad returns i.e. DaemonID = HASH (Org Name, Service Name, daemon endpoint)
func GetDaemonID() string {
	rawID := config.GetString(config.OrganizationId) + config.GetString(config.ServiceId) + daemonGroupId + config.GetString(config.DaemonEndPoint)
	//get hash of the string id combination
	hasher := sha256.New()
	hasher.Write([]byte(rawID))
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash
}

var daemonGroupId string

// setter method for daemonGroupID
func SetDaemonGrpId(grpId string) {
	daemonGroupId = grpId
}

// New Daemon registration. Generates the DaemonID and use that as getting access token
func RegisterDaemon(serviceURL string) bool {
	daemonID := GetDaemonID()
	status := false
	// call the service and get the result
	status = callRegisterService(daemonID, serviceURL)
	if !status {
		log.Infof("Daemon unable to register with the monitoring service. ")
	} else {
		log.Infof("Daemon successfully registered with the monitoring service. ")
	}
	return status
}

/*
service request
{"daemonID":"3a4ebeb75eace1857a9133c7a50bdbb841b35de60f78bc43eafe0d204e523dfe"}

service output
true/false
*/
