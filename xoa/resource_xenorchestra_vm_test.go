package xoa

import (
	"encoding/json"
	"testing"
)

var vmObjectData string = `
{
  "type": "VM",
  "addresses": {},
  "auto_poweron": false,
  "boot": {
    "order": "ncd"
  },
  "CPUs": {
    "max": 1,
    "number": 1
  },
  "current_operations": {},
  "expNestedHvm": false,
  "high_availability": "",
  "memory": {
    "dynamic": [
      1073741824,
      1073741824
    ],
    "static": [
      536870912,
      1073741824
    ],
    "size": 1073733632
  },
  "installTime": 1552287083,
  "name_description": "Testingsdfsdf",
  "name_label": "Hello from terraform!",
  "other": {
    "base_template_name": "Ubuntu Bionic Beaver 18.04",
    "import_task": "OpaqueRef:a1c9c64b-eeec-48cd-b587-51a8fc7924d0",
    "mac_seed": "bc583da8-d7a4-9437-7630-ce5ecce7efd0",
    "install-methods": "cdrom,nfs,http,ftp",
    "linux_template": "true"
  },
  "os_version": {},
  "power_state": "Running",
  "hasVendorDevice": false,
  "snapshots": [],
  "startTime": 1552445802,
  "tags": [],
  "VIFs": [
    "13793e84-110e-7f0d-8544-cb2f39adf2f4"
  ],
  "virtualizationMode": "hvm",
  "xenTools": false,
  "$container": "a5c7d15c-2724-47ce-8e30-77f21f08bb4c",
  "$VBDs": [
    "43202d8a-c2ba-963e-54d0-b8cf770b2725",
    "5b8adb53-36a7-8489-3a5f-f4ec8aa79568",
    "5a12a8dd-f217-9f90-5aab-8f434893a6e3",
    "567d6474-62ac-2d98-66ed-53cf468fc8b3",
    "42eeb5e0-b14e-99c0-36a6-61b671112353"
  ],
  "VGPUs": [],
  "$VGPUs": [],
  "vga": "cirrus",
  "videoram": "8",
  "id": "77c6637c-fa3d-0a46-717e-296208c40169",
  "uuid": "77c6637c-fa3d-0a46-717e-296208c40169",
  "$pool": "cadf25ab-91ff-6fc0-041f-5a7033c4bc78",
  "$poolId": "cadf25ab-91ff-6fc0-041f-5a7033c4bc78"
}`

var data string = `
{
  "6944cce9-5ce0-a853-ee9c-bcc8281b597f": {
    "body": "VM 'Test VM' started on host: xenserver-ddelnano (uuid: a5c7d15c-2724-47ce-8e30-77f21f08bb4c)",
    "name": "VM_STARTED",
    "time": 1547577637,
    "$object": "5f318ba2-2300-cc34-3710-f64f53634ac0",
    "id": "6944cce9-5ce0-a853-ee9c-bcc8281b597f",
    "type": "message",
    "uuid": "6944cce9-5ce0-a853-ee9c-bcc8281b597f",
    "$pool": "cadf25ab-91ff-6fc0-041f-5a7033c4bc78",
    "$poolId": "cadf25ab-91ff-6fc0-041f-5a7033c4bc78"
  },
  "52e21132-2f4f-3b80-ef65-7d43750eb6db": {
    "body": "VM 'XOA' started on host: xenserver-ddelnano (uuid: a5c7d15c-2724-47ce-8e30-77f21f08bb4c)",
    "name": "VM_STARTED",
    "time": 1547578119,
    "$object": "9df00260-c6d8-6cd4-4b24-d9e23602bf95",
    "id": "52e21132-2f4f-3b80-ef65-7d43750eb6db",
    "type": "message",
    "uuid": "52e21132-2f4f-3b80-ef65-7d43750eb6db",
    "$pool": "cadf25ab-91ff-6fc0-041f-5a7033c4bc78",
    "$poolId": "cadf25ab-91ff-6fc0-041f-5a7033c4bc78"
  },
  "3bdc3bed-2a91-12af-5154-097706009593": {
    "body": "VM 'XOA' started on host: xenserver-ddelnano (uuid: a5c7d15c-2724-47ce-8e30-77f21f08bb4c)",
    "name": "VM_STARTED",
    "time": 1547578434,
    "$object": "2e793135-3b3d-354e-46c4-ef99fe0a37d0",
    "id": "3bdc3bed-2a91-12af-5154-097706009593",
    "type": "message",
    "uuid": "3bdc3bed-2a91-12af-5154-097706009593",
    "$pool": "cadf25ab-91ff-6fc0-041f-5a7033c4bc78",
    "$poolId": "cadf25ab-91ff-6fc0-041f-5a7033c4bc78"
  },
  "ad97f23d-3bf0-7fac-e70c-ab98661450c6": {
    "body": "VM 'XOA' shutdown forcibly",
    "name": "VM_SHUTDOWN",
    "time": 1547578721,
    "$object": "9df00260-c6d8-6cd4-4b24-d9e23602bf95",
    "id": "ad97f23d-3bf0-7fac-e70c-ab98661450c6",
    "type": "message",
    "uuid": "ad97f23d-3bf0-7fac-e70c-ab98661450c6",
    "$pool": "cadf25ab-91ff-6fc0-041f-5a7033c4bc78",
    "$poolId": "cadf25ab-91ff-6fc0-041f-5a7033c4bc78"
  },
  "77c6637c-fa3d-0a46-717e-296208c40169": {
    "body": "VM 'XOA' shutdown forcibly",
    "name": "VM_SHUTDOWN",
    "time": 1547578721,
    "$object": "9df00260-c6d8-6cd4-4b24-d9e23602bf95",
    "id": "77c6637c-fa3d-0a46-717e-296208c40169",
    "type": "VM",
    "uuid": "ad97f23d-3bf0-7fac-e70c-ab98661450c6",
    "$pool": "cadf25ab-91ff-6fc0-041f-5a7033c4bc78",
    "$poolId": "cadf25ab-91ff-6fc0-041f-5a7033c4bc78"
  }
}
`

func TestThatUnmarshalingWorks(t *testing.T) {
	var allObjectRes allObjectResponse
	err := json.Unmarshal([]byte(data), &allObjectRes.Objects)

	if err != nil || allObjectRes.Objects["77c6637c-fa3d-0a46-717e-296208c40169"].Id != "77c6637c-fa3d-0a46-717e-296208c40169" {
		t.Fatalf("error: %v. Need to have VM object: %v", err, allObjectRes)
	}
}

func TestUnmarshalingVmObject(t *testing.T) {
	var vmObj VmObject

	err := json.Unmarshal([]byte(vmObjectData), &vmObj)

	if err != nil {
		t.Fatalf("error: %v. Need to have VM object: %v", err, vmObj)
	}

	if !validateVmObject(vmObj) {
		t.Fatalf("VmObject has not passed validation")
	}

}

func validateVmObject(o VmObject) bool {
	if o.Type != "VM" {
		return false
	}

	if o.CPUs.Number != 1 {
		return false
	}

	if o.Memory.Size != 1073733632 {
		return false
	}

	if o.PowerState != "Running" {
		return false
	}
	if o.VirtualizationMode != "hvm" {
		return false
	}

	if o.VIFs[0] != "13793e84-110e-7f0d-8544-cb2f39adf2f4" {
		return false
	}

	if o.PoolId != "cadf25ab-91ff-6fc0-041f-5a7033c4bc78" {
		return false
	}

	return true
}