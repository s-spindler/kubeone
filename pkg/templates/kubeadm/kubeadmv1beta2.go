/*
Copyright 2019 The KubeOne Authors.

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

	kubeoneapi "k8c.io/kubeone/pkg/apis/kubeone"
	"k8c.io/kubeone/pkg/state"
	"k8c.io/kubeone/pkg/templates"
	"k8c.io/kubeone/pkg/templates/kubeadm/v1beta2"
)

type kubeadmv1beta2 struct {
	version string
}

func (*kubeadmv1beta2) Config(s *state.State, instance kubeoneapi.HostConfig) (string, error) {
	config, err := v1beta2.NewConfig(s, instance)
	if err != nil {
		return "", err
	}

	return templates.KubernetesToYAML(config, kubeletConf)
}

func (*kubeadmv1beta2) ConfigWorker(s *state.State, instance kubeoneapi.HostConfig) (string, error) {
	config, err := v1beta2.NewConfigWorker(s, instance)
	if err != nil {
		return "", err
	}

	return templates.KubernetesToYAML(config, kubeletConf)
}

func (k *kubeadmv1beta2) UpgradeLeaderCommand() string {
	return fmt.Sprintf("kubeadm upgrade apply -y --certificate-renewal=true %s", k.version)
}

func (*kubeadmv1beta2) UpgradeFollowerCommand() string {
	return kubeadmUpgradeNodeCommand
}

func (*kubeadmv1beta2) UpgradeStaticWorkerCommand() string {
	return kubeadmUpgradeNodeCommand
}
