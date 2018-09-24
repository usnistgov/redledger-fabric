/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package server

import (
	"github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric/protos/token"
	"github.com/pkg/errors"
)

//go:generate counterfeiter -o mock/acl_provider.go -fake-name ACLProvider . ACLProvider

type ACLProvider interface {
	// CheckACL checks access control for the resource for the given channel.
	// idinfo is an object such as []*common.SignedData from which
	// an id can be extracted for testing against a policy
	CheckACL(resName string, channelID string, idinfo interface{}) error
}

type ACLResources struct {
	IssueTokens    string
	TransferTokens string
	ListTokens     string
}

// PolicyBasedAccessControl implements token command access control functions.
type PolicyBasedAccessControl struct {
	ACLProvider  ACLProvider
	ACLResources *ACLResources
}

func (ac *PolicyBasedAccessControl) Check(sc *token.SignedCommand, c *token.Command) error {
	signedData := []*common.SignedData{{
		Identity:  c.Header.Creator,
		Data:      sc.Command,
		Signature: sc.Signature,
	}}

	switch t := c.GetPayload().(type) {

	case *token.Command_ImportRequest:
		return ac.ACLProvider.CheckACL(
			ac.ACLResources.IssueTokens,
			c.Header.ChannelId,
			signedData,
		)
	case *token.Command_ListRequest:
		return ac.ACLProvider.CheckACL(
			ac.ACLResources.ListTokens,
			c.Header.ChannelId,
			signedData,
		)
	case *token.Command_TransferRequest:
		return ac.ACLProvider.CheckACL(
			ac.ACLResources.TransferTokens,
			c.Header.ChannelId,
			signedData,
		)
	case *token.Command_RedeemRequest:
		// Redeem has same policy as transfer
		return ac.ACLProvider.CheckACL(
			ac.ACLResources.TransferTokens,
			c.Header.ChannelId,
			signedData,
		)
	default:
		return errors.Errorf("command type not recognized: %T", t)
	}
}
