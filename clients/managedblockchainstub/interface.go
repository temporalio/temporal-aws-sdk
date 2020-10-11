// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package managedblockchainstub

import (
	"github.com/aws/aws-sdk-go/service/managedblockchain"
	"go.temporal.io/aws-sdk/clients"
	"go.temporal.io/sdk/workflow"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateMember(ctx workflow.Context, input *managedblockchain.CreateMemberInput) (*managedblockchain.CreateMemberOutput, error)
	CreateMemberAsync(ctx workflow.Context, input *managedblockchain.CreateMemberInput) *ManagedBlockchainCreateMemberFuture

	CreateNetwork(ctx workflow.Context, input *managedblockchain.CreateNetworkInput) (*managedblockchain.CreateNetworkOutput, error)
	CreateNetworkAsync(ctx workflow.Context, input *managedblockchain.CreateNetworkInput) *ManagedBlockchainCreateNetworkFuture

	CreateNode(ctx workflow.Context, input *managedblockchain.CreateNodeInput) (*managedblockchain.CreateNodeOutput, error)
	CreateNodeAsync(ctx workflow.Context, input *managedblockchain.CreateNodeInput) *ManagedBlockchainCreateNodeFuture

	CreateProposal(ctx workflow.Context, input *managedblockchain.CreateProposalInput) (*managedblockchain.CreateProposalOutput, error)
	CreateProposalAsync(ctx workflow.Context, input *managedblockchain.CreateProposalInput) *ManagedBlockchainCreateProposalFuture

	DeleteMember(ctx workflow.Context, input *managedblockchain.DeleteMemberInput) (*managedblockchain.DeleteMemberOutput, error)
	DeleteMemberAsync(ctx workflow.Context, input *managedblockchain.DeleteMemberInput) *ManagedBlockchainDeleteMemberFuture

	DeleteNode(ctx workflow.Context, input *managedblockchain.DeleteNodeInput) (*managedblockchain.DeleteNodeOutput, error)
	DeleteNodeAsync(ctx workflow.Context, input *managedblockchain.DeleteNodeInput) *ManagedBlockchainDeleteNodeFuture

	GetMember(ctx workflow.Context, input *managedblockchain.GetMemberInput) (*managedblockchain.GetMemberOutput, error)
	GetMemberAsync(ctx workflow.Context, input *managedblockchain.GetMemberInput) *ManagedBlockchainGetMemberFuture

	GetNetwork(ctx workflow.Context, input *managedblockchain.GetNetworkInput) (*managedblockchain.GetNetworkOutput, error)
	GetNetworkAsync(ctx workflow.Context, input *managedblockchain.GetNetworkInput) *ManagedBlockchainGetNetworkFuture

	GetNode(ctx workflow.Context, input *managedblockchain.GetNodeInput) (*managedblockchain.GetNodeOutput, error)
	GetNodeAsync(ctx workflow.Context, input *managedblockchain.GetNodeInput) *ManagedBlockchainGetNodeFuture

	GetProposal(ctx workflow.Context, input *managedblockchain.GetProposalInput) (*managedblockchain.GetProposalOutput, error)
	GetProposalAsync(ctx workflow.Context, input *managedblockchain.GetProposalInput) *ManagedBlockchainGetProposalFuture

	ListInvitations(ctx workflow.Context, input *managedblockchain.ListInvitationsInput) (*managedblockchain.ListInvitationsOutput, error)
	ListInvitationsAsync(ctx workflow.Context, input *managedblockchain.ListInvitationsInput) *ManagedBlockchainListInvitationsFuture

	ListMembers(ctx workflow.Context, input *managedblockchain.ListMembersInput) (*managedblockchain.ListMembersOutput, error)
	ListMembersAsync(ctx workflow.Context, input *managedblockchain.ListMembersInput) *ManagedBlockchainListMembersFuture

	ListNetworks(ctx workflow.Context, input *managedblockchain.ListNetworksInput) (*managedblockchain.ListNetworksOutput, error)
	ListNetworksAsync(ctx workflow.Context, input *managedblockchain.ListNetworksInput) *ManagedBlockchainListNetworksFuture

	ListNodes(ctx workflow.Context, input *managedblockchain.ListNodesInput) (*managedblockchain.ListNodesOutput, error)
	ListNodesAsync(ctx workflow.Context, input *managedblockchain.ListNodesInput) *ManagedBlockchainListNodesFuture

	ListProposalVotes(ctx workflow.Context, input *managedblockchain.ListProposalVotesInput) (*managedblockchain.ListProposalVotesOutput, error)
	ListProposalVotesAsync(ctx workflow.Context, input *managedblockchain.ListProposalVotesInput) *ManagedBlockchainListProposalVotesFuture

	ListProposals(ctx workflow.Context, input *managedblockchain.ListProposalsInput) (*managedblockchain.ListProposalsOutput, error)
	ListProposalsAsync(ctx workflow.Context, input *managedblockchain.ListProposalsInput) *ManagedBlockchainListProposalsFuture

	RejectInvitation(ctx workflow.Context, input *managedblockchain.RejectInvitationInput) (*managedblockchain.RejectInvitationOutput, error)
	RejectInvitationAsync(ctx workflow.Context, input *managedblockchain.RejectInvitationInput) *ManagedBlockchainRejectInvitationFuture

	UpdateMember(ctx workflow.Context, input *managedblockchain.UpdateMemberInput) (*managedblockchain.UpdateMemberOutput, error)
	UpdateMemberAsync(ctx workflow.Context, input *managedblockchain.UpdateMemberInput) *ManagedBlockchainUpdateMemberFuture

	UpdateNode(ctx workflow.Context, input *managedblockchain.UpdateNodeInput) (*managedblockchain.UpdateNodeOutput, error)
	UpdateNodeAsync(ctx workflow.Context, input *managedblockchain.UpdateNodeInput) *ManagedBlockchainUpdateNodeFuture

	VoteOnProposal(ctx workflow.Context, input *managedblockchain.VoteOnProposalInput) (*managedblockchain.VoteOnProposalOutput, error)
	VoteOnProposalAsync(ctx workflow.Context, input *managedblockchain.VoteOnProposalInput) *ManagedBlockchainVoteOnProposalFuture
}

func NewClient() Client {
	return &stub{}
}
