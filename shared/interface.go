package shared

import (
	"context"

	"github.com/catdevman/awsume-go/proto"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	// This isn't required when using VersionedPlugins
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	"arguments_grpc":   &ArgumentsPlugin{},
	"profiles_grpc":    &ProfilesPlugin{},
	"credentials_grpc": &CredentialsPlugin{},
	//	"profile_names_grpc": &ProfileNamesPlugin{},
}

type ArgumentsService interface {
	Pre() error
	Get() error
	Post() error
}

type ProfilesService interface {
	Pre() error
	Get() (Profiles, error)
	Post(Profiles) error
}

type CredentialsService interface {
	Pre() error
	Get() error
	Post() error
}

type ProfileNamesService interface {
	Pre() error
	Get() error
	Post() error
}

// This is the implementation of plugin.GRPCPlugin so we can serve/consume this.
type ArgumentsPlugin struct {
	// GRPCPlugin must still implement the Plugin interface
	plugin.Plugin
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl ArgumentsService
	proto.UnimplementedArgumentsServer
}

func (p *ArgumentsPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterArgumentsServer(s, &ArgumentsServer{Impl: p.Impl})
	return nil
}

func (p *ArgumentsPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &ArgumentsClient{client: proto.NewArgumentsClient(c)}, nil
}

// This is the implementation of plugin.GRPCPlugin so we can serve/consume this.
type ProfilesPlugin struct {
	// GRPCPlugin must still implement the Plugin interface
	plugin.Plugin
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl ProfilesService
	proto.UnimplementedProfilesServer
}

func (p *ProfilesPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterProfilesServer(s, &ProfilesServer{Impl: p.Impl})
	return nil
}

func (p *ProfilesPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &ProfilesClient{client: proto.NewProfilesClient(c)}, nil
}

// This is the implementation of plugin.GRPCPlugin so we can serve/consume this.
type CredentialsPlugin struct {
	// GRPCPlugin must still implement the Plugin interface
	plugin.Plugin
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl CredentialsService
	proto.UnimplementedCredentialsServer
}

func (p *CredentialsPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterCredentialsServer(s, &CredentialsServer{Impl: p.Impl})
	return nil
}

func (p *CredentialsPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &CredentialsClient{client: proto.NewCredentialsClient(c)}, nil
}

// This is the implementation of plugin.GRPCPlugin so we can serve/consume this.
//type ProfileNamesPlugin struct {
//	// GRPCPlugin must still implement the Plugin interface
//	plugin.Plugin
//	// Concrete implementation, written in Go. This is only used for plugins
//	// that are written in Go.
//	Impl ProfileNamesService
//	proto.UnimplementedProfileNamesServer
//}
//
//func (p *ProfileNamesPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
//	proto.RegisterProfileNamesServer(s, &GRPCServer{Impl: p.Impl})
//	return nil
//}
//
//func (p *ProfileNamesPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
//	return &GRPCClient{client: proto.NewProfileNamesClient(c)}, nil
//}
