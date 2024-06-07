package identity

import (
	"context"
	"github/tronglv_authen_author/api/authentication"
	"github/tronglv_authen_author/api/authorization"
	"github/tronglv_authen_author/helper/client"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type Client struct {
	connection  *grpc.ClientConn
	authzClient authorization.AuthorizationClient
	authClient  authentication.AuthenticationClient
}

func New(c zrpc.RpcClientConf) *Client {
	connection := client.NewGRPCConnection(c).Conn()
	return &Client{
		connection:  connection,
		authzClient: authorization.NewAuthorizationClient(connection),
		authClient:  authentication.NewAuthenticationClient(connection),
	}
}

func (c *Client) Close() {
	logx.Infof("closing grpc connection of the identity service")
	_ = c.connection.Close()
}

func (c *Client) Validate(ctx context.Context, req *authentication.TokenAuthReq) (*authentication.TokenAuthResp, error) {
	return c.authClient.Validate(ctx, req)
}

func (c *Client) Permissions(ctx context.Context, req *authorization.PermissionListReq) (*authorization.PermissionListResp, error) {
	return c.authzClient.Permissions(ctx, req)
}

func (c *Client) Permission(ctx context.Context, req *authorization.PermissionReq) (*authorization.PermissionResp, error) {
	return c.authzClient.Permission(ctx, req)
}

func (c *Client) PermissionByCode(ctx context.Context, req *authorization.PermissionByCodeReq) (*authorization.PermissionResp, error) {
	return c.authzClient.PermissionByCode(ctx, req)
}
