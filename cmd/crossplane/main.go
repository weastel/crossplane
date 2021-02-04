/*
Copyright 2019 The Crossplane Authors.

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

package main

import (
	"strings"

	"github.com/alecthomas/kong"

	"github.com/crossplane/crossplane/cmd/crossplane/core"
	"github.com/crossplane/crossplane/cmd/crossplane/rbac"
)

var _ = kong.Must(&cli)

var cli struct {
	Core core.Cmd `cmd:"" help:"Start core Crossplane controllers."`
	Rbac rbac.Cmd `cmd:"" help:"Start Crossplane RBAC Manager controllers."`
}

func main() {

	// TODO: Fix Debugging issue
	// NOTE(negz): We must setup our logger after calling kingpin.MustParse in
	// order to ensure the debug flag has been parsed and set.
	// zl := zap.New(zap.UseDevMode(*debug))
	// if *debug {
	// 	// The controller-runtime runs with a no-op logger by default. It is
	// 	// *very* verbose even at info level, so we only provide it a real
	// 	// logger when we're running in debug mode.
	// 	ctrl.SetLogger(zl)
	// }

	ctx := kong.Parse(&cli,
		kong.Name("crossplane"),
		kong.Description("An open source multicloud control plane."),
		kong.Vars{
			"ManagementPolicyAll": rbac.ManagementPolicyAll,
			"ManagementPolicyEnum": strings.Join(
				[]string{
					rbac.ManagementPolicyAll,
					rbac.ManagementPolicyBasic,
				},
				","),
		},
		kong.UsageOnError())
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
