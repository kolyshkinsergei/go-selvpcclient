package testing

import (
	"time"

	"github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/licenses"
	"github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/servers"
)

// TestGetLicenseResponseRaw represents a raw response from the Get request.
const TestGetLicenseResponseRaw = `
{
    "license": {
        "id": 123123,
        "project_id": "49338ac045f448e294b25d013f890317",
        "region": "ru-2",
        "servers": [
            {
                "id": "253b680c-89f6-4c85-afbf-c9a67c92d3fe",
                "name": "Node00",
                "status": "ACTIVE",
                "updated": "2018-02-20T22:02:21Z"
            }
        ],
        "status": "ACTIVE",
        "type": "license_windows_2012_standard"
    }
}
`

var licenseServerTimeStamp, _ = time.Parse(time.RFC3339, "2018-02-20T22:02:21Z")

// TestGetLicenseResponse represents an unmarshalled TestGetLicenseResponseRaw.
var TestGetLicenseResponse = &licenses.License{
	ID:        123123,
	ProjectID: "49338ac045f448e294b25d013f890317",
	Region:    "ru-2",
	Servers: []servers.Server{
		{
			ID:      "253b680c-89f6-4c85-afbf-c9a67c92d3fe",
			Name:    "Node00",
			Status:  "ACTIVE",
			Updated: licenseServerTimeStamp,
		},
	},
	Status: "ACTIVE",
	Type:   "license_windows_2012_standard",
}

// TestListLicensesResponseRaw represents a raw response from the List request.
const TestListLicensesResponseRaw = `
{
    "licenses": [
        {
            "id": 1123123,
            "project_id": "49338ac045f448e294b25d013f890317",
            "region": "ru-1",
            "status": "DOWN",
            "type": "license_windows_2012_standard"
        },
        {
            "id": 124123,
            "project_id": "49338ac045f448e294b25d013f890317",
            "region": "ru-3",
            "status": "DOWN",
            "type": "license_windows_2016_standard"
        },
        {
            "id": 13212,
            "project_id": "49338ac045f448e294b25d013f890317",
            "region": "ru-2",
            "status": "DOWN",
            "type": "license_windows_2016_standard"
        }
    ]
}
`

// TestListLicensesSingleResponseRaw represents a raw response with a single license from the List request.
const TestListLicensesSingleResponseRaw = `
{
    "licenses": [
        {
            "id": 1123123,
            "project_id": "49338ac045f448e294b25d013f890317",
            "region": "ru-1",
			"status": "DOWN",
            "type": "license_windows_2012_standard"
        }
    ]
}
`

// TestListLicensesSingleResponse represents the unmarshalled TestListLicensesSingleResponseRaw response.
var TestListLicensesSingleResponse = []*licenses.License{
	{
		ID:        1123123,
		ProjectID: "49338ac045f448e294b25d013f890317",
		Region:    "ru-1",
		Status:    "DOWN",
		Type:      "license_windows_2012_standard",
	},
}

// TestCreateLicenseOptsRaw represents marshalled options for the Create request.
const TestCreateLicenseOptsRaw = `
{
    "licenses": [
        {
            "region": "ru-2",
            "quantity": 1,
            "type": "license_windows_2016_standard"
        }
    ]
}
`

// TestCreateLicenseOpts represent options for the Create request.
var TestCreateLicenseOpts = licenses.LicenseOpts{
	Licenses: []licenses.LicenseOpt{
		{
			Region:   "ru-2",
			Quantity: 1,
			Type:     "license_windows_2016_standard",
		},
	},
}

// TestCreateLicenseResponseRaw represents a raw response from the Create request.
const TestCreateLicenseResponseRaw = `
{
    "licenses": [
        {
            "type": "license_windows_2016_standard",
            "status": "DOWN",
            "region": "ru-2",
            "project_id": "49338ac045f448e294b25d013f890317",
            "id": 1123123
        }
    ]
}
`

// TestCreateLicenseResponse represents the unmarshalled TestCreateLicenseResponseRaw response.
var TestCreateLicenseResponse = []*licenses.License{
	{
		ID:        1123123,
		ProjectID: "49338ac045f448e294b25d013f890317",
		Region:    "ru-2",
		Status:    "DOWN",
		Type:      "license_windows_2016_standard",
	},
}

// TestManyLicensesInvalidResponseRaw represents a raw invalid response with several licenses.
const TestManyLicensesInvalidResponseRaw = `
{
    "licenses": [
        {
            "id": "49338ac045f448e294b25d013f890317"
        }
    ]
}
`

// TestSingleLicenseInvalidResponseRaw represents a raw invalid response with a single license.
const TestSingleLicenseInvalidResponseRaw = `
{
    "license": {
        "id": "49338ac045f448e294b25d013f890317"
    }
}
`
