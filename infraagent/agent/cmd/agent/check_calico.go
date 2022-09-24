// Copyright (c) 2022 Intel Corporation.  All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License")
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package agent

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/ipdk-io/k8s-infra-offload/pkg/types"
	"github.com/ipdk-io/k8s-infra-offload/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	var timeoutString string
	checkCalicoConfig.PersistentFlags().StringVar(&timeoutString, "timeout", "0s", "timeout after which program return error if file does not exist, passing duration <= 0 means wait forever.")
	if err := viper.BindPFlag("timeout", checkCalicoConfig.PersistentFlags().Lookup("timeout")); err != nil {
		fmt.Fprintf(os.Stderr, "There was an error while binding flags '%s'", err)
		os.Exit(1)
	}
	rootCmd.AddCommand(checkCalicoConfig)
}

var checkCalicoConfig = &cobra.Command{
	Use:           "checkCalicoConfig",
	Short:         "Wait until Calico config exist then exit",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		timeout, err := time.ParseDuration(viper.GetString("timeout"))
		if err != nil {
			return err
		}
		if timeout == 0 {
			fmt.Fprintf(os.Stdout, "Waiting for Calico config file in %s\n", path.Dir(types.DefaultCalicoConfig))
		} else {
			fmt.Fprintf(os.Stdout, "Waiting for Calico config file in %s for %s\n", path.Dir(types.DefaultCalicoConfig), viper.GetString("timeout"))
		}

		if err := utils.WaitForCalicoConfig(timeout, types.DefaultCalicoConfig); err != nil {
			return err
		}
		fmt.Fprintf(os.Stdout, "Calico config file exist. Exit.\n")
		return nil
	},
}
