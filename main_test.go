package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComprehensiveHclParse(t *testing.T) {

	moduleContent := `
	variable "test1" {
		default = "variable default"
		description = "variable description"
	}

	variable "test2" {
		default = "variable default"
		description = "variable description"
	}

	output "test" {
		sensitive = true
		description = "output description"
	}
	data "null" "test" {
		provider = null.test
	}

	resource "null" "test" {}
	provider "null" {}
	module "test" {}
	terraform {
		required_providers {
			null = {
				source = "abc"
				version = "~> 1.2.3"
			}
		}
	}
	`

	module, err := parseHcl(moduleContent)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "<virtual>", module.Path)
	assert.Len(t, module.ModuleCalls, 1)
	assert.Len(t, module.Variables, 2)
	assert.Len(t, module.Outputs, 1)
	assert.Len(t, module.DataResources, 1)
	assert.Len(t, module.ManagedResources, 1)
	assert.Len(t, module.ProviderConfigs, 1)
	assert.Len(t, module.RequiredProviders, 1)

	assert.Equal(t, "test", module.ModuleCalls["test"].Name)

	assert.Equal(t, "test1", module.Variables["test1"].Name)
	assert.Equal(t, "variable default", module.Variables["test1"].Default)
	assert.Equal(t, "variable description", module.Variables["test1"].Description)

	assert.Equal(t, "test", module.Outputs["test"].Name)
	assert.Equal(t, true, module.Outputs["test"].Sensitive)
	assert.Equal(t, "output description", module.Outputs["test"].Description)

	assert.Equal(t, "test", module.DataResources["data.null.test"].Name)
	assert.Equal(t, "data", module.DataResources["data.null.test"].Mode.String())
	assert.Equal(t, "test", module.DataResources["data.null.test"].Provider.Alias)
	assert.Equal(t, "null", module.DataResources["data.null.test"].Provider.Name)

	assert.Equal(t, "test", module.ManagedResources["null.test"].Name)
	assert.Equal(t, "managed", module.ManagedResources["null.test"].Mode.String())
	assert.Equal(t, "", module.ManagedResources["null.test"].Provider.Alias)
	assert.Equal(t, "null", module.ManagedResources["null.test"].Provider.Name)

	assert.Equal(t, "null", module.ProviderConfigs["null"].Name)
	assert.Equal(t, "abc", module.RequiredProviders["null"].Source)
	assert.Equal(t, []string{"~> 1.2.3"}, module.RequiredProviders["null"].VersionConstraints)
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
