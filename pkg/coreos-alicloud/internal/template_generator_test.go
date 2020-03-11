// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package internal_test

import (
	. "github.com/gardener/gardener-extension-os-coreos-alicloud/pkg/coreos-alicloud/internal"

	. "github.com/gardener/gardener-extension-os-coreos-alicloud/pkg/coreos-alicloud/coreos"
	"github.com/gobuffalo/packr"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	onlyOwnerPerm = int32(0600)
)

var _ = Describe("#TemplateBashGenerator", func() {

	box := packr.NewBox("./testfiles")
	var (
		ExpectedCloudInit []byte
		err               error
	)

	It("should render correctly", func() {
		ExpectedCloudInit, err = box.Find("cloud-init.sh")
		Expect(err).NotTo(HaveOccurred())

		gen := NewCloudInitGenerator(DefaultUnitsPath)

		cloudInit, err := gen.Generate(&OperatingSystemConfig{
			Files: []*File{
				{
					Path:        "/foo",
					Content:     []byte("bar"),
					Permissions: &onlyOwnerPerm,
				},
			},

			Units: []*Unit{
				{
					Name:    "docker.service",
					Content: []byte("unit"),
					DropIns: []*DropIn{
						{
							Name:    "10-docker-opts.conf",
							Content: []byte("override"),
						},
					},
				},
			},
			Bootstrap: true,
		})

		Expect(err).NotTo(HaveOccurred())
		Expect(cloudInit).To(Equal(ExpectedCloudInit))
	})

	It("should render correctly", func() {
		ExpectedCloudInit, err = box.Find("cloud-init-reconcile.sh")
		Expect(err).NotTo(HaveOccurred())

		gen := NewCloudInitGenerator(DefaultUnitsPath)

		cloudInit, err := gen.Generate(&OperatingSystemConfig{
			Files: []*File{
				{
					Path:        "/foo",
					Content:     []byte("bar"),
					Permissions: &onlyOwnerPerm,
				},
			},

			Units: []*Unit{
				{
					Name: "containerd.service",
					DropIns: []*DropIn{
						{
							Name:    "11-exec-start-config.conf",
							Content: []byte(ContainerDUnitDropInContent),
						},
					},
				},
				{
					Name: "docker.service",
					DropIns: []*DropIn{
						{
							Name:    "11-exec-start-config.conf",
							Content: []byte(DockerUnitDropInContent),
						},
					},
				},
			},
			Bootstrap: false,
		})

		Expect(err).NotTo(HaveOccurred())
		Expect(cloudInit).To(Equal(ExpectedCloudInit))
	})
})
