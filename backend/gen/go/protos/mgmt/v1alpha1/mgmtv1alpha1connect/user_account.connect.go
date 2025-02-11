// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: mgmt/v1alpha1/user_account.proto

package mgmtv1alpha1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1alpha1 "github.com/nucleuscloud/neosync/backend/gen/go/protos/mgmt/v1alpha1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// UserAccountServiceName is the fully-qualified name of the UserAccountService service.
	UserAccountServiceName = "mgmt.v1alpha1.UserAccountService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// UserAccountServiceGetUserProcedure is the fully-qualified name of the UserAccountService's
	// GetUser RPC.
	UserAccountServiceGetUserProcedure = "/mgmt.v1alpha1.UserAccountService/GetUser"
	// UserAccountServiceSetUserProcedure is the fully-qualified name of the UserAccountService's
	// SetUser RPC.
	UserAccountServiceSetUserProcedure = "/mgmt.v1alpha1.UserAccountService/SetUser"
	// UserAccountServiceGetUserAccountsProcedure is the fully-qualified name of the
	// UserAccountService's GetUserAccounts RPC.
	UserAccountServiceGetUserAccountsProcedure = "/mgmt.v1alpha1.UserAccountService/GetUserAccounts"
	// UserAccountServiceSetPersonalAccountProcedure is the fully-qualified name of the
	// UserAccountService's SetPersonalAccount RPC.
	UserAccountServiceSetPersonalAccountProcedure = "/mgmt.v1alpha1.UserAccountService/SetPersonalAccount"
	// UserAccountServiceConvertPersonalToTeamAccountProcedure is the fully-qualified name of the
	// UserAccountService's ConvertPersonalToTeamAccount RPC.
	UserAccountServiceConvertPersonalToTeamAccountProcedure = "/mgmt.v1alpha1.UserAccountService/ConvertPersonalToTeamAccount"
	// UserAccountServiceCreateTeamAccountProcedure is the fully-qualified name of the
	// UserAccountService's CreateTeamAccount RPC.
	UserAccountServiceCreateTeamAccountProcedure = "/mgmt.v1alpha1.UserAccountService/CreateTeamAccount"
	// UserAccountServiceIsUserInAccountProcedure is the fully-qualified name of the
	// UserAccountService's IsUserInAccount RPC.
	UserAccountServiceIsUserInAccountProcedure = "/mgmt.v1alpha1.UserAccountService/IsUserInAccount"
	// UserAccountServiceGetAccountTemporalConfigProcedure is the fully-qualified name of the
	// UserAccountService's GetAccountTemporalConfig RPC.
	UserAccountServiceGetAccountTemporalConfigProcedure = "/mgmt.v1alpha1.UserAccountService/GetAccountTemporalConfig"
	// UserAccountServiceSetAccountTemporalConfigProcedure is the fully-qualified name of the
	// UserAccountService's SetAccountTemporalConfig RPC.
	UserAccountServiceSetAccountTemporalConfigProcedure = "/mgmt.v1alpha1.UserAccountService/SetAccountTemporalConfig"
	// UserAccountServiceGetTeamAccountMembersProcedure is the fully-qualified name of the
	// UserAccountService's GetTeamAccountMembers RPC.
	UserAccountServiceGetTeamAccountMembersProcedure = "/mgmt.v1alpha1.UserAccountService/GetTeamAccountMembers"
	// UserAccountServiceRemoveTeamAccountMemberProcedure is the fully-qualified name of the
	// UserAccountService's RemoveTeamAccountMember RPC.
	UserAccountServiceRemoveTeamAccountMemberProcedure = "/mgmt.v1alpha1.UserAccountService/RemoveTeamAccountMember"
	// UserAccountServiceInviteUserToTeamAccountProcedure is the fully-qualified name of the
	// UserAccountService's InviteUserToTeamAccount RPC.
	UserAccountServiceInviteUserToTeamAccountProcedure = "/mgmt.v1alpha1.UserAccountService/InviteUserToTeamAccount"
	// UserAccountServiceGetTeamAccountInvitesProcedure is the fully-qualified name of the
	// UserAccountService's GetTeamAccountInvites RPC.
	UserAccountServiceGetTeamAccountInvitesProcedure = "/mgmt.v1alpha1.UserAccountService/GetTeamAccountInvites"
	// UserAccountServiceRemoveTeamAccountInviteProcedure is the fully-qualified name of the
	// UserAccountService's RemoveTeamAccountInvite RPC.
	UserAccountServiceRemoveTeamAccountInviteProcedure = "/mgmt.v1alpha1.UserAccountService/RemoveTeamAccountInvite"
	// UserAccountServiceAcceptTeamAccountInviteProcedure is the fully-qualified name of the
	// UserAccountService's AcceptTeamAccountInvite RPC.
	UserAccountServiceAcceptTeamAccountInviteProcedure = "/mgmt.v1alpha1.UserAccountService/AcceptTeamAccountInvite"
	// UserAccountServiceGetSystemInformationProcedure is the fully-qualified name of the
	// UserAccountService's GetSystemInformation RPC.
	UserAccountServiceGetSystemInformationProcedure = "/mgmt.v1alpha1.UserAccountService/GetSystemInformation"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	userAccountServiceServiceDescriptor                            = v1alpha1.File_mgmt_v1alpha1_user_account_proto.Services().ByName("UserAccountService")
	userAccountServiceGetUserMethodDescriptor                      = userAccountServiceServiceDescriptor.Methods().ByName("GetUser")
	userAccountServiceSetUserMethodDescriptor                      = userAccountServiceServiceDescriptor.Methods().ByName("SetUser")
	userAccountServiceGetUserAccountsMethodDescriptor              = userAccountServiceServiceDescriptor.Methods().ByName("GetUserAccounts")
	userAccountServiceSetPersonalAccountMethodDescriptor           = userAccountServiceServiceDescriptor.Methods().ByName("SetPersonalAccount")
	userAccountServiceConvertPersonalToTeamAccountMethodDescriptor = userAccountServiceServiceDescriptor.Methods().ByName("ConvertPersonalToTeamAccount")
	userAccountServiceCreateTeamAccountMethodDescriptor            = userAccountServiceServiceDescriptor.Methods().ByName("CreateTeamAccount")
	userAccountServiceIsUserInAccountMethodDescriptor              = userAccountServiceServiceDescriptor.Methods().ByName("IsUserInAccount")
	userAccountServiceGetAccountTemporalConfigMethodDescriptor     = userAccountServiceServiceDescriptor.Methods().ByName("GetAccountTemporalConfig")
	userAccountServiceSetAccountTemporalConfigMethodDescriptor     = userAccountServiceServiceDescriptor.Methods().ByName("SetAccountTemporalConfig")
	userAccountServiceGetTeamAccountMembersMethodDescriptor        = userAccountServiceServiceDescriptor.Methods().ByName("GetTeamAccountMembers")
	userAccountServiceRemoveTeamAccountMemberMethodDescriptor      = userAccountServiceServiceDescriptor.Methods().ByName("RemoveTeamAccountMember")
	userAccountServiceInviteUserToTeamAccountMethodDescriptor      = userAccountServiceServiceDescriptor.Methods().ByName("InviteUserToTeamAccount")
	userAccountServiceGetTeamAccountInvitesMethodDescriptor        = userAccountServiceServiceDescriptor.Methods().ByName("GetTeamAccountInvites")
	userAccountServiceRemoveTeamAccountInviteMethodDescriptor      = userAccountServiceServiceDescriptor.Methods().ByName("RemoveTeamAccountInvite")
	userAccountServiceAcceptTeamAccountInviteMethodDescriptor      = userAccountServiceServiceDescriptor.Methods().ByName("AcceptTeamAccountInvite")
	userAccountServiceGetSystemInformationMethodDescriptor         = userAccountServiceServiceDescriptor.Methods().ByName("GetSystemInformation")
)

// UserAccountServiceClient is a client for the mgmt.v1alpha1.UserAccountService service.
type UserAccountServiceClient interface {
	GetUser(context.Context, *connect.Request[v1alpha1.GetUserRequest]) (*connect.Response[v1alpha1.GetUserResponse], error)
	SetUser(context.Context, *connect.Request[v1alpha1.SetUserRequest]) (*connect.Response[v1alpha1.SetUserResponse], error)
	GetUserAccounts(context.Context, *connect.Request[v1alpha1.GetUserAccountsRequest]) (*connect.Response[v1alpha1.GetUserAccountsResponse], error)
	SetPersonalAccount(context.Context, *connect.Request[v1alpha1.SetPersonalAccountRequest]) (*connect.Response[v1alpha1.SetPersonalAccountResponse], error)
	ConvertPersonalToTeamAccount(context.Context, *connect.Request[v1alpha1.ConvertPersonalToTeamAccountRequest]) (*connect.Response[v1alpha1.ConvertPersonalToTeamAccountResponse], error)
	CreateTeamAccount(context.Context, *connect.Request[v1alpha1.CreateTeamAccountRequest]) (*connect.Response[v1alpha1.CreateTeamAccountResponse], error)
	IsUserInAccount(context.Context, *connect.Request[v1alpha1.IsUserInAccountRequest]) (*connect.Response[v1alpha1.IsUserInAccountResponse], error)
	GetAccountTemporalConfig(context.Context, *connect.Request[v1alpha1.GetAccountTemporalConfigRequest]) (*connect.Response[v1alpha1.GetAccountTemporalConfigResponse], error)
	SetAccountTemporalConfig(context.Context, *connect.Request[v1alpha1.SetAccountTemporalConfigRequest]) (*connect.Response[v1alpha1.SetAccountTemporalConfigResponse], error)
	GetTeamAccountMembers(context.Context, *connect.Request[v1alpha1.GetTeamAccountMembersRequest]) (*connect.Response[v1alpha1.GetTeamAccountMembersResponse], error)
	RemoveTeamAccountMember(context.Context, *connect.Request[v1alpha1.RemoveTeamAccountMemberRequest]) (*connect.Response[v1alpha1.RemoveTeamAccountMemberResponse], error)
	InviteUserToTeamAccount(context.Context, *connect.Request[v1alpha1.InviteUserToTeamAccountRequest]) (*connect.Response[v1alpha1.InviteUserToTeamAccountResponse], error)
	GetTeamAccountInvites(context.Context, *connect.Request[v1alpha1.GetTeamAccountInvitesRequest]) (*connect.Response[v1alpha1.GetTeamAccountInvitesResponse], error)
	RemoveTeamAccountInvite(context.Context, *connect.Request[v1alpha1.RemoveTeamAccountInviteRequest]) (*connect.Response[v1alpha1.RemoveTeamAccountInviteResponse], error)
	AcceptTeamAccountInvite(context.Context, *connect.Request[v1alpha1.AcceptTeamAccountInviteRequest]) (*connect.Response[v1alpha1.AcceptTeamAccountInviteResponse], error)
	GetSystemInformation(context.Context, *connect.Request[v1alpha1.GetSystemInformationRequest]) (*connect.Response[v1alpha1.GetSystemInformationResponse], error)
}

// NewUserAccountServiceClient constructs a client for the mgmt.v1alpha1.UserAccountService service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUserAccountServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) UserAccountServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &userAccountServiceClient{
		getUser: connect.NewClient[v1alpha1.GetUserRequest, v1alpha1.GetUserResponse](
			httpClient,
			baseURL+UserAccountServiceGetUserProcedure,
			connect.WithSchema(userAccountServiceGetUserMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		setUser: connect.NewClient[v1alpha1.SetUserRequest, v1alpha1.SetUserResponse](
			httpClient,
			baseURL+UserAccountServiceSetUserProcedure,
			connect.WithSchema(userAccountServiceSetUserMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getUserAccounts: connect.NewClient[v1alpha1.GetUserAccountsRequest, v1alpha1.GetUserAccountsResponse](
			httpClient,
			baseURL+UserAccountServiceGetUserAccountsProcedure,
			connect.WithSchema(userAccountServiceGetUserAccountsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		setPersonalAccount: connect.NewClient[v1alpha1.SetPersonalAccountRequest, v1alpha1.SetPersonalAccountResponse](
			httpClient,
			baseURL+UserAccountServiceSetPersonalAccountProcedure,
			connect.WithSchema(userAccountServiceSetPersonalAccountMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		convertPersonalToTeamAccount: connect.NewClient[v1alpha1.ConvertPersonalToTeamAccountRequest, v1alpha1.ConvertPersonalToTeamAccountResponse](
			httpClient,
			baseURL+UserAccountServiceConvertPersonalToTeamAccountProcedure,
			connect.WithSchema(userAccountServiceConvertPersonalToTeamAccountMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createTeamAccount: connect.NewClient[v1alpha1.CreateTeamAccountRequest, v1alpha1.CreateTeamAccountResponse](
			httpClient,
			baseURL+UserAccountServiceCreateTeamAccountProcedure,
			connect.WithSchema(userAccountServiceCreateTeamAccountMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		isUserInAccount: connect.NewClient[v1alpha1.IsUserInAccountRequest, v1alpha1.IsUserInAccountResponse](
			httpClient,
			baseURL+UserAccountServiceIsUserInAccountProcedure,
			connect.WithSchema(userAccountServiceIsUserInAccountMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getAccountTemporalConfig: connect.NewClient[v1alpha1.GetAccountTemporalConfigRequest, v1alpha1.GetAccountTemporalConfigResponse](
			httpClient,
			baseURL+UserAccountServiceGetAccountTemporalConfigProcedure,
			connect.WithSchema(userAccountServiceGetAccountTemporalConfigMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		setAccountTemporalConfig: connect.NewClient[v1alpha1.SetAccountTemporalConfigRequest, v1alpha1.SetAccountTemporalConfigResponse](
			httpClient,
			baseURL+UserAccountServiceSetAccountTemporalConfigProcedure,
			connect.WithSchema(userAccountServiceSetAccountTemporalConfigMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getTeamAccountMembers: connect.NewClient[v1alpha1.GetTeamAccountMembersRequest, v1alpha1.GetTeamAccountMembersResponse](
			httpClient,
			baseURL+UserAccountServiceGetTeamAccountMembersProcedure,
			connect.WithSchema(userAccountServiceGetTeamAccountMembersMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		removeTeamAccountMember: connect.NewClient[v1alpha1.RemoveTeamAccountMemberRequest, v1alpha1.RemoveTeamAccountMemberResponse](
			httpClient,
			baseURL+UserAccountServiceRemoveTeamAccountMemberProcedure,
			connect.WithSchema(userAccountServiceRemoveTeamAccountMemberMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		inviteUserToTeamAccount: connect.NewClient[v1alpha1.InviteUserToTeamAccountRequest, v1alpha1.InviteUserToTeamAccountResponse](
			httpClient,
			baseURL+UserAccountServiceInviteUserToTeamAccountProcedure,
			connect.WithSchema(userAccountServiceInviteUserToTeamAccountMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getTeamAccountInvites: connect.NewClient[v1alpha1.GetTeamAccountInvitesRequest, v1alpha1.GetTeamAccountInvitesResponse](
			httpClient,
			baseURL+UserAccountServiceGetTeamAccountInvitesProcedure,
			connect.WithSchema(userAccountServiceGetTeamAccountInvitesMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		removeTeamAccountInvite: connect.NewClient[v1alpha1.RemoveTeamAccountInviteRequest, v1alpha1.RemoveTeamAccountInviteResponse](
			httpClient,
			baseURL+UserAccountServiceRemoveTeamAccountInviteProcedure,
			connect.WithSchema(userAccountServiceRemoveTeamAccountInviteMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		acceptTeamAccountInvite: connect.NewClient[v1alpha1.AcceptTeamAccountInviteRequest, v1alpha1.AcceptTeamAccountInviteResponse](
			httpClient,
			baseURL+UserAccountServiceAcceptTeamAccountInviteProcedure,
			connect.WithSchema(userAccountServiceAcceptTeamAccountInviteMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getSystemInformation: connect.NewClient[v1alpha1.GetSystemInformationRequest, v1alpha1.GetSystemInformationResponse](
			httpClient,
			baseURL+UserAccountServiceGetSystemInformationProcedure,
			connect.WithSchema(userAccountServiceGetSystemInformationMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// userAccountServiceClient implements UserAccountServiceClient.
type userAccountServiceClient struct {
	getUser                      *connect.Client[v1alpha1.GetUserRequest, v1alpha1.GetUserResponse]
	setUser                      *connect.Client[v1alpha1.SetUserRequest, v1alpha1.SetUserResponse]
	getUserAccounts              *connect.Client[v1alpha1.GetUserAccountsRequest, v1alpha1.GetUserAccountsResponse]
	setPersonalAccount           *connect.Client[v1alpha1.SetPersonalAccountRequest, v1alpha1.SetPersonalAccountResponse]
	convertPersonalToTeamAccount *connect.Client[v1alpha1.ConvertPersonalToTeamAccountRequest, v1alpha1.ConvertPersonalToTeamAccountResponse]
	createTeamAccount            *connect.Client[v1alpha1.CreateTeamAccountRequest, v1alpha1.CreateTeamAccountResponse]
	isUserInAccount              *connect.Client[v1alpha1.IsUserInAccountRequest, v1alpha1.IsUserInAccountResponse]
	getAccountTemporalConfig     *connect.Client[v1alpha1.GetAccountTemporalConfigRequest, v1alpha1.GetAccountTemporalConfigResponse]
	setAccountTemporalConfig     *connect.Client[v1alpha1.SetAccountTemporalConfigRequest, v1alpha1.SetAccountTemporalConfigResponse]
	getTeamAccountMembers        *connect.Client[v1alpha1.GetTeamAccountMembersRequest, v1alpha1.GetTeamAccountMembersResponse]
	removeTeamAccountMember      *connect.Client[v1alpha1.RemoveTeamAccountMemberRequest, v1alpha1.RemoveTeamAccountMemberResponse]
	inviteUserToTeamAccount      *connect.Client[v1alpha1.InviteUserToTeamAccountRequest, v1alpha1.InviteUserToTeamAccountResponse]
	getTeamAccountInvites        *connect.Client[v1alpha1.GetTeamAccountInvitesRequest, v1alpha1.GetTeamAccountInvitesResponse]
	removeTeamAccountInvite      *connect.Client[v1alpha1.RemoveTeamAccountInviteRequest, v1alpha1.RemoveTeamAccountInviteResponse]
	acceptTeamAccountInvite      *connect.Client[v1alpha1.AcceptTeamAccountInviteRequest, v1alpha1.AcceptTeamAccountInviteResponse]
	getSystemInformation         *connect.Client[v1alpha1.GetSystemInformationRequest, v1alpha1.GetSystemInformationResponse]
}

// GetUser calls mgmt.v1alpha1.UserAccountService.GetUser.
func (c *userAccountServiceClient) GetUser(ctx context.Context, req *connect.Request[v1alpha1.GetUserRequest]) (*connect.Response[v1alpha1.GetUserResponse], error) {
	return c.getUser.CallUnary(ctx, req)
}

// SetUser calls mgmt.v1alpha1.UserAccountService.SetUser.
func (c *userAccountServiceClient) SetUser(ctx context.Context, req *connect.Request[v1alpha1.SetUserRequest]) (*connect.Response[v1alpha1.SetUserResponse], error) {
	return c.setUser.CallUnary(ctx, req)
}

// GetUserAccounts calls mgmt.v1alpha1.UserAccountService.GetUserAccounts.
func (c *userAccountServiceClient) GetUserAccounts(ctx context.Context, req *connect.Request[v1alpha1.GetUserAccountsRequest]) (*connect.Response[v1alpha1.GetUserAccountsResponse], error) {
	return c.getUserAccounts.CallUnary(ctx, req)
}

// SetPersonalAccount calls mgmt.v1alpha1.UserAccountService.SetPersonalAccount.
func (c *userAccountServiceClient) SetPersonalAccount(ctx context.Context, req *connect.Request[v1alpha1.SetPersonalAccountRequest]) (*connect.Response[v1alpha1.SetPersonalAccountResponse], error) {
	return c.setPersonalAccount.CallUnary(ctx, req)
}

// ConvertPersonalToTeamAccount calls mgmt.v1alpha1.UserAccountService.ConvertPersonalToTeamAccount.
func (c *userAccountServiceClient) ConvertPersonalToTeamAccount(ctx context.Context, req *connect.Request[v1alpha1.ConvertPersonalToTeamAccountRequest]) (*connect.Response[v1alpha1.ConvertPersonalToTeamAccountResponse], error) {
	return c.convertPersonalToTeamAccount.CallUnary(ctx, req)
}

// CreateTeamAccount calls mgmt.v1alpha1.UserAccountService.CreateTeamAccount.
func (c *userAccountServiceClient) CreateTeamAccount(ctx context.Context, req *connect.Request[v1alpha1.CreateTeamAccountRequest]) (*connect.Response[v1alpha1.CreateTeamAccountResponse], error) {
	return c.createTeamAccount.CallUnary(ctx, req)
}

// IsUserInAccount calls mgmt.v1alpha1.UserAccountService.IsUserInAccount.
func (c *userAccountServiceClient) IsUserInAccount(ctx context.Context, req *connect.Request[v1alpha1.IsUserInAccountRequest]) (*connect.Response[v1alpha1.IsUserInAccountResponse], error) {
	return c.isUserInAccount.CallUnary(ctx, req)
}

// GetAccountTemporalConfig calls mgmt.v1alpha1.UserAccountService.GetAccountTemporalConfig.
func (c *userAccountServiceClient) GetAccountTemporalConfig(ctx context.Context, req *connect.Request[v1alpha1.GetAccountTemporalConfigRequest]) (*connect.Response[v1alpha1.GetAccountTemporalConfigResponse], error) {
	return c.getAccountTemporalConfig.CallUnary(ctx, req)
}

// SetAccountTemporalConfig calls mgmt.v1alpha1.UserAccountService.SetAccountTemporalConfig.
func (c *userAccountServiceClient) SetAccountTemporalConfig(ctx context.Context, req *connect.Request[v1alpha1.SetAccountTemporalConfigRequest]) (*connect.Response[v1alpha1.SetAccountTemporalConfigResponse], error) {
	return c.setAccountTemporalConfig.CallUnary(ctx, req)
}

// GetTeamAccountMembers calls mgmt.v1alpha1.UserAccountService.GetTeamAccountMembers.
func (c *userAccountServiceClient) GetTeamAccountMembers(ctx context.Context, req *connect.Request[v1alpha1.GetTeamAccountMembersRequest]) (*connect.Response[v1alpha1.GetTeamAccountMembersResponse], error) {
	return c.getTeamAccountMembers.CallUnary(ctx, req)
}

// RemoveTeamAccountMember calls mgmt.v1alpha1.UserAccountService.RemoveTeamAccountMember.
func (c *userAccountServiceClient) RemoveTeamAccountMember(ctx context.Context, req *connect.Request[v1alpha1.RemoveTeamAccountMemberRequest]) (*connect.Response[v1alpha1.RemoveTeamAccountMemberResponse], error) {
	return c.removeTeamAccountMember.CallUnary(ctx, req)
}

// InviteUserToTeamAccount calls mgmt.v1alpha1.UserAccountService.InviteUserToTeamAccount.
func (c *userAccountServiceClient) InviteUserToTeamAccount(ctx context.Context, req *connect.Request[v1alpha1.InviteUserToTeamAccountRequest]) (*connect.Response[v1alpha1.InviteUserToTeamAccountResponse], error) {
	return c.inviteUserToTeamAccount.CallUnary(ctx, req)
}

// GetTeamAccountInvites calls mgmt.v1alpha1.UserAccountService.GetTeamAccountInvites.
func (c *userAccountServiceClient) GetTeamAccountInvites(ctx context.Context, req *connect.Request[v1alpha1.GetTeamAccountInvitesRequest]) (*connect.Response[v1alpha1.GetTeamAccountInvitesResponse], error) {
	return c.getTeamAccountInvites.CallUnary(ctx, req)
}

// RemoveTeamAccountInvite calls mgmt.v1alpha1.UserAccountService.RemoveTeamAccountInvite.
func (c *userAccountServiceClient) RemoveTeamAccountInvite(ctx context.Context, req *connect.Request[v1alpha1.RemoveTeamAccountInviteRequest]) (*connect.Response[v1alpha1.RemoveTeamAccountInviteResponse], error) {
	return c.removeTeamAccountInvite.CallUnary(ctx, req)
}

// AcceptTeamAccountInvite calls mgmt.v1alpha1.UserAccountService.AcceptTeamAccountInvite.
func (c *userAccountServiceClient) AcceptTeamAccountInvite(ctx context.Context, req *connect.Request[v1alpha1.AcceptTeamAccountInviteRequest]) (*connect.Response[v1alpha1.AcceptTeamAccountInviteResponse], error) {
	return c.acceptTeamAccountInvite.CallUnary(ctx, req)
}

// GetSystemInformation calls mgmt.v1alpha1.UserAccountService.GetSystemInformation.
func (c *userAccountServiceClient) GetSystemInformation(ctx context.Context, req *connect.Request[v1alpha1.GetSystemInformationRequest]) (*connect.Response[v1alpha1.GetSystemInformationResponse], error) {
	return c.getSystemInformation.CallUnary(ctx, req)
}

// UserAccountServiceHandler is an implementation of the mgmt.v1alpha1.UserAccountService service.
type UserAccountServiceHandler interface {
	GetUser(context.Context, *connect.Request[v1alpha1.GetUserRequest]) (*connect.Response[v1alpha1.GetUserResponse], error)
	SetUser(context.Context, *connect.Request[v1alpha1.SetUserRequest]) (*connect.Response[v1alpha1.SetUserResponse], error)
	GetUserAccounts(context.Context, *connect.Request[v1alpha1.GetUserAccountsRequest]) (*connect.Response[v1alpha1.GetUserAccountsResponse], error)
	SetPersonalAccount(context.Context, *connect.Request[v1alpha1.SetPersonalAccountRequest]) (*connect.Response[v1alpha1.SetPersonalAccountResponse], error)
	ConvertPersonalToTeamAccount(context.Context, *connect.Request[v1alpha1.ConvertPersonalToTeamAccountRequest]) (*connect.Response[v1alpha1.ConvertPersonalToTeamAccountResponse], error)
	CreateTeamAccount(context.Context, *connect.Request[v1alpha1.CreateTeamAccountRequest]) (*connect.Response[v1alpha1.CreateTeamAccountResponse], error)
	IsUserInAccount(context.Context, *connect.Request[v1alpha1.IsUserInAccountRequest]) (*connect.Response[v1alpha1.IsUserInAccountResponse], error)
	GetAccountTemporalConfig(context.Context, *connect.Request[v1alpha1.GetAccountTemporalConfigRequest]) (*connect.Response[v1alpha1.GetAccountTemporalConfigResponse], error)
	SetAccountTemporalConfig(context.Context, *connect.Request[v1alpha1.SetAccountTemporalConfigRequest]) (*connect.Response[v1alpha1.SetAccountTemporalConfigResponse], error)
	GetTeamAccountMembers(context.Context, *connect.Request[v1alpha1.GetTeamAccountMembersRequest]) (*connect.Response[v1alpha1.GetTeamAccountMembersResponse], error)
	RemoveTeamAccountMember(context.Context, *connect.Request[v1alpha1.RemoveTeamAccountMemberRequest]) (*connect.Response[v1alpha1.RemoveTeamAccountMemberResponse], error)
	InviteUserToTeamAccount(context.Context, *connect.Request[v1alpha1.InviteUserToTeamAccountRequest]) (*connect.Response[v1alpha1.InviteUserToTeamAccountResponse], error)
	GetTeamAccountInvites(context.Context, *connect.Request[v1alpha1.GetTeamAccountInvitesRequest]) (*connect.Response[v1alpha1.GetTeamAccountInvitesResponse], error)
	RemoveTeamAccountInvite(context.Context, *connect.Request[v1alpha1.RemoveTeamAccountInviteRequest]) (*connect.Response[v1alpha1.RemoveTeamAccountInviteResponse], error)
	AcceptTeamAccountInvite(context.Context, *connect.Request[v1alpha1.AcceptTeamAccountInviteRequest]) (*connect.Response[v1alpha1.AcceptTeamAccountInviteResponse], error)
	GetSystemInformation(context.Context, *connect.Request[v1alpha1.GetSystemInformationRequest]) (*connect.Response[v1alpha1.GetSystemInformationResponse], error)
}

// NewUserAccountServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUserAccountServiceHandler(svc UserAccountServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	userAccountServiceGetUserHandler := connect.NewUnaryHandler(
		UserAccountServiceGetUserProcedure,
		svc.GetUser,
		connect.WithSchema(userAccountServiceGetUserMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userAccountServiceSetUserHandler := connect.NewUnaryHandler(
		UserAccountServiceSetUserProcedure,
		svc.SetUser,
		connect.WithSchema(userAccountServiceSetUserMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userAccountServiceGetUserAccountsHandler := connect.NewUnaryHandler(
		UserAccountServiceGetUserAccountsProcedure,
		svc.GetUserAccounts,
		connect.WithSchema(userAccountServiceGetUserAccountsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userAccountServiceSetPersonalAccountHandler := connect.NewUnaryHandler(
		UserAccountServiceSetPersonalAccountProcedure,
		svc.SetPersonalAccount,
		connect.WithSchema(userAccountServiceSetPersonalAccountMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userAccountServiceConvertPersonalToTeamAccountHandler := connect.NewUnaryHandler(
		UserAccountServiceConvertPersonalToTeamAccountProcedure,
		svc.ConvertPersonalToTeamAccount,
		connect.WithSchema(userAccountServiceConvertPersonalToTeamAccountMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userAccountServiceCreateTeamAccountHandler := connect.NewUnaryHandler(
		UserAccountServiceCreateTeamAccountProcedure,
		svc.CreateTeamAccount,
		connect.WithSchema(userAccountServiceCreateTeamAccountMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userAccountServiceIsUserInAccountHandler := connect.NewUnaryHandler(
		UserAccountServiceIsUserInAccountProcedure,
		svc.IsUserInAccount,
		connect.WithSchema(userAccountServiceIsUserInAccountMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userAccountServiceGetAccountTemporalConfigHandler := connect.NewUnaryHandler(
		UserAccountServiceGetAccountTemporalConfigProcedure,
		svc.GetAccountTemporalConfig,
		connect.WithSchema(userAccountServiceGetAccountTemporalConfigMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userAccountServiceSetAccountTemporalConfigHandler := connect.NewUnaryHandler(
		UserAccountServiceSetAccountTemporalConfigProcedure,
		svc.SetAccountTemporalConfig,
		connect.WithSchema(userAccountServiceSetAccountTemporalConfigMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userAccountServiceGetTeamAccountMembersHandler := connect.NewUnaryHandler(
		UserAccountServiceGetTeamAccountMembersProcedure,
		svc.GetTeamAccountMembers,
		connect.WithSchema(userAccountServiceGetTeamAccountMembersMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userAccountServiceRemoveTeamAccountMemberHandler := connect.NewUnaryHandler(
		UserAccountServiceRemoveTeamAccountMemberProcedure,
		svc.RemoveTeamAccountMember,
		connect.WithSchema(userAccountServiceRemoveTeamAccountMemberMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userAccountServiceInviteUserToTeamAccountHandler := connect.NewUnaryHandler(
		UserAccountServiceInviteUserToTeamAccountProcedure,
		svc.InviteUserToTeamAccount,
		connect.WithSchema(userAccountServiceInviteUserToTeamAccountMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userAccountServiceGetTeamAccountInvitesHandler := connect.NewUnaryHandler(
		UserAccountServiceGetTeamAccountInvitesProcedure,
		svc.GetTeamAccountInvites,
		connect.WithSchema(userAccountServiceGetTeamAccountInvitesMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userAccountServiceRemoveTeamAccountInviteHandler := connect.NewUnaryHandler(
		UserAccountServiceRemoveTeamAccountInviteProcedure,
		svc.RemoveTeamAccountInvite,
		connect.WithSchema(userAccountServiceRemoveTeamAccountInviteMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userAccountServiceAcceptTeamAccountInviteHandler := connect.NewUnaryHandler(
		UserAccountServiceAcceptTeamAccountInviteProcedure,
		svc.AcceptTeamAccountInvite,
		connect.WithSchema(userAccountServiceAcceptTeamAccountInviteMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userAccountServiceGetSystemInformationHandler := connect.NewUnaryHandler(
		UserAccountServiceGetSystemInformationProcedure,
		svc.GetSystemInformation,
		connect.WithSchema(userAccountServiceGetSystemInformationMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/mgmt.v1alpha1.UserAccountService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case UserAccountServiceGetUserProcedure:
			userAccountServiceGetUserHandler.ServeHTTP(w, r)
		case UserAccountServiceSetUserProcedure:
			userAccountServiceSetUserHandler.ServeHTTP(w, r)
		case UserAccountServiceGetUserAccountsProcedure:
			userAccountServiceGetUserAccountsHandler.ServeHTTP(w, r)
		case UserAccountServiceSetPersonalAccountProcedure:
			userAccountServiceSetPersonalAccountHandler.ServeHTTP(w, r)
		case UserAccountServiceConvertPersonalToTeamAccountProcedure:
			userAccountServiceConvertPersonalToTeamAccountHandler.ServeHTTP(w, r)
		case UserAccountServiceCreateTeamAccountProcedure:
			userAccountServiceCreateTeamAccountHandler.ServeHTTP(w, r)
		case UserAccountServiceIsUserInAccountProcedure:
			userAccountServiceIsUserInAccountHandler.ServeHTTP(w, r)
		case UserAccountServiceGetAccountTemporalConfigProcedure:
			userAccountServiceGetAccountTemporalConfigHandler.ServeHTTP(w, r)
		case UserAccountServiceSetAccountTemporalConfigProcedure:
			userAccountServiceSetAccountTemporalConfigHandler.ServeHTTP(w, r)
		case UserAccountServiceGetTeamAccountMembersProcedure:
			userAccountServiceGetTeamAccountMembersHandler.ServeHTTP(w, r)
		case UserAccountServiceRemoveTeamAccountMemberProcedure:
			userAccountServiceRemoveTeamAccountMemberHandler.ServeHTTP(w, r)
		case UserAccountServiceInviteUserToTeamAccountProcedure:
			userAccountServiceInviteUserToTeamAccountHandler.ServeHTTP(w, r)
		case UserAccountServiceGetTeamAccountInvitesProcedure:
			userAccountServiceGetTeamAccountInvitesHandler.ServeHTTP(w, r)
		case UserAccountServiceRemoveTeamAccountInviteProcedure:
			userAccountServiceRemoveTeamAccountInviteHandler.ServeHTTP(w, r)
		case UserAccountServiceAcceptTeamAccountInviteProcedure:
			userAccountServiceAcceptTeamAccountInviteHandler.ServeHTTP(w, r)
		case UserAccountServiceGetSystemInformationProcedure:
			userAccountServiceGetSystemInformationHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedUserAccountServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedUserAccountServiceHandler struct{}

func (UnimplementedUserAccountServiceHandler) GetUser(context.Context, *connect.Request[v1alpha1.GetUserRequest]) (*connect.Response[v1alpha1.GetUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.UserAccountService.GetUser is not implemented"))
}

func (UnimplementedUserAccountServiceHandler) SetUser(context.Context, *connect.Request[v1alpha1.SetUserRequest]) (*connect.Response[v1alpha1.SetUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.UserAccountService.SetUser is not implemented"))
}

func (UnimplementedUserAccountServiceHandler) GetUserAccounts(context.Context, *connect.Request[v1alpha1.GetUserAccountsRequest]) (*connect.Response[v1alpha1.GetUserAccountsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.UserAccountService.GetUserAccounts is not implemented"))
}

func (UnimplementedUserAccountServiceHandler) SetPersonalAccount(context.Context, *connect.Request[v1alpha1.SetPersonalAccountRequest]) (*connect.Response[v1alpha1.SetPersonalAccountResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.UserAccountService.SetPersonalAccount is not implemented"))
}

func (UnimplementedUserAccountServiceHandler) ConvertPersonalToTeamAccount(context.Context, *connect.Request[v1alpha1.ConvertPersonalToTeamAccountRequest]) (*connect.Response[v1alpha1.ConvertPersonalToTeamAccountResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.UserAccountService.ConvertPersonalToTeamAccount is not implemented"))
}

func (UnimplementedUserAccountServiceHandler) CreateTeamAccount(context.Context, *connect.Request[v1alpha1.CreateTeamAccountRequest]) (*connect.Response[v1alpha1.CreateTeamAccountResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.UserAccountService.CreateTeamAccount is not implemented"))
}

func (UnimplementedUserAccountServiceHandler) IsUserInAccount(context.Context, *connect.Request[v1alpha1.IsUserInAccountRequest]) (*connect.Response[v1alpha1.IsUserInAccountResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.UserAccountService.IsUserInAccount is not implemented"))
}

func (UnimplementedUserAccountServiceHandler) GetAccountTemporalConfig(context.Context, *connect.Request[v1alpha1.GetAccountTemporalConfigRequest]) (*connect.Response[v1alpha1.GetAccountTemporalConfigResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.UserAccountService.GetAccountTemporalConfig is not implemented"))
}

func (UnimplementedUserAccountServiceHandler) SetAccountTemporalConfig(context.Context, *connect.Request[v1alpha1.SetAccountTemporalConfigRequest]) (*connect.Response[v1alpha1.SetAccountTemporalConfigResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.UserAccountService.SetAccountTemporalConfig is not implemented"))
}

func (UnimplementedUserAccountServiceHandler) GetTeamAccountMembers(context.Context, *connect.Request[v1alpha1.GetTeamAccountMembersRequest]) (*connect.Response[v1alpha1.GetTeamAccountMembersResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.UserAccountService.GetTeamAccountMembers is not implemented"))
}

func (UnimplementedUserAccountServiceHandler) RemoveTeamAccountMember(context.Context, *connect.Request[v1alpha1.RemoveTeamAccountMemberRequest]) (*connect.Response[v1alpha1.RemoveTeamAccountMemberResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.UserAccountService.RemoveTeamAccountMember is not implemented"))
}

func (UnimplementedUserAccountServiceHandler) InviteUserToTeamAccount(context.Context, *connect.Request[v1alpha1.InviteUserToTeamAccountRequest]) (*connect.Response[v1alpha1.InviteUserToTeamAccountResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.UserAccountService.InviteUserToTeamAccount is not implemented"))
}

func (UnimplementedUserAccountServiceHandler) GetTeamAccountInvites(context.Context, *connect.Request[v1alpha1.GetTeamAccountInvitesRequest]) (*connect.Response[v1alpha1.GetTeamAccountInvitesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.UserAccountService.GetTeamAccountInvites is not implemented"))
}

func (UnimplementedUserAccountServiceHandler) RemoveTeamAccountInvite(context.Context, *connect.Request[v1alpha1.RemoveTeamAccountInviteRequest]) (*connect.Response[v1alpha1.RemoveTeamAccountInviteResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.UserAccountService.RemoveTeamAccountInvite is not implemented"))
}

func (UnimplementedUserAccountServiceHandler) AcceptTeamAccountInvite(context.Context, *connect.Request[v1alpha1.AcceptTeamAccountInviteRequest]) (*connect.Response[v1alpha1.AcceptTeamAccountInviteResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.UserAccountService.AcceptTeamAccountInvite is not implemented"))
}

func (UnimplementedUserAccountServiceHandler) GetSystemInformation(context.Context, *connect.Request[v1alpha1.GetSystemInformationRequest]) (*connect.Response[v1alpha1.GetSystemInformationResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.UserAccountService.GetSystemInformation is not implemented"))
}
