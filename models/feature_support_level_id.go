// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// FeatureSupportLevelID feature support level id
//
// swagger:model feature-support-level-id
type FeatureSupportLevelID string

func NewFeatureSupportLevelID(value FeatureSupportLevelID) *FeatureSupportLevelID {
	return &value
}

// Pointer returns a pointer to a freshly-allocated FeatureSupportLevelID.
func (m FeatureSupportLevelID) Pointer() *FeatureSupportLevelID {
	return &m
}

const (

	// FeatureSupportLevelIDSNO captures enum value "SNO"
	FeatureSupportLevelIDSNO FeatureSupportLevelID = "SNO"

	// FeatureSupportLevelIDVIPAUTOALLOC captures enum value "VIP_AUTO_ALLOC"
	FeatureSupportLevelIDVIPAUTOALLOC FeatureSupportLevelID = "VIP_AUTO_ALLOC"

	// FeatureSupportLevelIDCUSTOMMANIFEST captures enum value "CUSTOM_MANIFEST"
	FeatureSupportLevelIDCUSTOMMANIFEST FeatureSupportLevelID = "CUSTOM_MANIFEST"

	// FeatureSupportLevelIDSINGLENODEEXPANSION captures enum value "SINGLE_NODE_EXPANSION"
	FeatureSupportLevelIDSINGLENODEEXPANSION FeatureSupportLevelID = "SINGLE_NODE_EXPANSION"

	// FeatureSupportLevelIDLVM captures enum value "LVM"
	FeatureSupportLevelIDLVM FeatureSupportLevelID = "LVM"

	// FeatureSupportLevelIDODF captures enum value "ODF"
	FeatureSupportLevelIDODF FeatureSupportLevelID = "ODF"

	// FeatureSupportLevelIDLSO captures enum value "LSO"
	FeatureSupportLevelIDLSO FeatureSupportLevelID = "LSO"

	// FeatureSupportLevelIDCNV captures enum value "CNV"
	FeatureSupportLevelIDCNV FeatureSupportLevelID = "CNV"

	// FeatureSupportLevelIDMCE captures enum value "MCE"
	FeatureSupportLevelIDMCE FeatureSupportLevelID = "MCE"

	// FeatureSupportLevelIDNUTANIXINTEGRATION captures enum value "NUTANIX_INTEGRATION"
	FeatureSupportLevelIDNUTANIXINTEGRATION FeatureSupportLevelID = "NUTANIX_INTEGRATION"

	// FeatureSupportLevelIDBAREMETALPLATFORM captures enum value "BAREMETAL_PLATFORM"
	FeatureSupportLevelIDBAREMETALPLATFORM FeatureSupportLevelID = "BAREMETAL_PLATFORM"

	// FeatureSupportLevelIDNONEPLATFORM captures enum value "NONE_PLATFORM"
	FeatureSupportLevelIDNONEPLATFORM FeatureSupportLevelID = "NONE_PLATFORM"

	// FeatureSupportLevelIDVSPHEREINTEGRATION captures enum value "VSPHERE_INTEGRATION"
	FeatureSupportLevelIDVSPHEREINTEGRATION FeatureSupportLevelID = "VSPHERE_INTEGRATION"

	// FeatureSupportLevelIDDUALSTACKVIPS captures enum value "DUAL_STACK_VIPS"
	FeatureSupportLevelIDDUALSTACKVIPS FeatureSupportLevelID = "DUAL_STACK_VIPS"

	// FeatureSupportLevelIDCLUSTERMANAGEDNETWORKING captures enum value "CLUSTER_MANAGED_NETWORKING"
	FeatureSupportLevelIDCLUSTERMANAGEDNETWORKING FeatureSupportLevelID = "CLUSTER_MANAGED_NETWORKING"

	// FeatureSupportLevelIDUSERMANAGEDNETWORKING captures enum value "USER_MANAGED_NETWORKING"
	FeatureSupportLevelIDUSERMANAGEDNETWORKING FeatureSupportLevelID = "USER_MANAGED_NETWORKING"

	// FeatureSupportLevelIDMINIMALISO captures enum value "MINIMAL_ISO"
	FeatureSupportLevelIDMINIMALISO FeatureSupportLevelID = "MINIMAL_ISO"

	// FeatureSupportLevelIDFULLISO captures enum value "FULL_ISO"
	FeatureSupportLevelIDFULLISO FeatureSupportLevelID = "FULL_ISO"

	// FeatureSupportLevelIDEXTERNALPLATFORMOCI captures enum value "EXTERNAL_PLATFORM_OCI"
	FeatureSupportLevelIDEXTERNALPLATFORMOCI FeatureSupportLevelID = "EXTERNAL_PLATFORM_OCI"

	// FeatureSupportLevelIDDUALSTACK captures enum value "DUAL_STACK"
	FeatureSupportLevelIDDUALSTACK FeatureSupportLevelID = "DUAL_STACK"

	// FeatureSupportLevelIDPLATFORMMANAGEDNETWORKING captures enum value "PLATFORM_MANAGED_NETWORKING"
	FeatureSupportLevelIDPLATFORMMANAGEDNETWORKING FeatureSupportLevelID = "PLATFORM_MANAGED_NETWORKING"

	// FeatureSupportLevelIDSKIPMCOREBOOT captures enum value "SKIP_MCO_REBOOT"
	FeatureSupportLevelIDSKIPMCOREBOOT FeatureSupportLevelID = "SKIP_MCO_REBOOT"
)

// for schema
var featureSupportLevelIdEnum []interface{}

func init() {
	var res []FeatureSupportLevelID
	if err := json.Unmarshal([]byte(`["SNO","VIP_AUTO_ALLOC","CUSTOM_MANIFEST","SINGLE_NODE_EXPANSION","LVM","ODF","LSO","CNV","MCE","NUTANIX_INTEGRATION","BAREMETAL_PLATFORM","NONE_PLATFORM","VSPHERE_INTEGRATION","DUAL_STACK_VIPS","CLUSTER_MANAGED_NETWORKING","USER_MANAGED_NETWORKING","MINIMAL_ISO","FULL_ISO","EXTERNAL_PLATFORM_OCI","DUAL_STACK","PLATFORM_MANAGED_NETWORKING","SKIP_MCO_REBOOT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		featureSupportLevelIdEnum = append(featureSupportLevelIdEnum, v)
	}
}

func (m FeatureSupportLevelID) validateFeatureSupportLevelIDEnum(path, location string, value FeatureSupportLevelID) error {
	if err := validate.EnumCase(path, location, value, featureSupportLevelIdEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this feature support level id
func (m FeatureSupportLevelID) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateFeatureSupportLevelIDEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this feature support level id based on context it is used
func (m FeatureSupportLevelID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
