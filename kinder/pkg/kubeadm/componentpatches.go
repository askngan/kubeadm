/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package kubeadm

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"k8s.io/kubeadm/kinder/pkg/constants"
)

// GetPatchesDirectoryPatches returns the kubeadm config patches that will instruct kubeadm
// to use patches directory.
func GetPatchesDirectoryPatches(kubeadmConfigVersion string) ([]string, error) {
	// select the patches for the kubeadm config version
	log.Debugf("Preparing patches directory for kubeadm config %s", kubeadmConfigVersion)
	if _, err := os.Stat(constants.PatchesDir); os.IsNotExist(err) {
		return []string{}, nil
	}

	var patchInit, patchJoin string
	switch kubeadmConfigVersion {
	case "v1beta3":
		patchInit = patchesDirectoryPatchInitv1beta3
		patchJoin = patchesDirectoryPatchJoinv1beta3
	default:
		return []string{}, errors.Errorf("unknown kubeadm config version: %s", kubeadmConfigVersion)
	}
	return []string{
		fmt.Sprintf(patchInit, constants.PatchesDir),
		fmt.Sprintf(patchJoin, constants.PatchesDir),
	}, nil
}

const patchesDirectoryPatchInitv1beta3 = `apiVersion: kubeadm.k8s.io/v1beta3
kind: InitConfiguration
metadata:
  name: config
patches:
  directory: %s`

const patchesDirectoryPatchJoinv1beta3 = `apiVersion: kubeadm.k8s.io/v1beta3
kind: JoinConfiguration
metadata:
  name: config
patches:
  directory: %s`
