package shared

import (
	"context"

	"github.com/catdevman/awsume-go/proto"
	"github.com/hashicorp/go-hclog"
)

// ArgumentsClient is an implementation of KV that talks over RPC.
type ArgumentsClient struct {
	client proto.ArgumentsClient
	logger hclog.Logger
}

func (m *ArgumentsClient) Pre() error {
	_, err := m.client.Pre(context.Background(), &proto.Empty{})
	return err
}

func (m *ArgumentsClient) Get() (*proto.ArgumentsMsg, error) {
	argsMsg, err := m.client.Get(context.Background(), &proto.Empty{})
	if err != nil {
		m.logger.Error(err.Error())
		return &proto.ArgumentsMsg{}, err
	}
	return argsMsg, nil
}

func (m *ArgumentsClient) Post(argsMsg *proto.ArgumentsMsg) (*proto.ArgumentsMsg, error) {
	args, err := m.client.Post(context.Background(), argsMsg)
	if err != nil {
		return &proto.ArgumentsMsg{}, err
	}

	return args, nil
}

// Here is the gRPC server that ProfilesClient talks to.
type ArgumentsServer struct {
	// This is the real implementation
	Impl ArgumentsService
	proto.UnimplementedArgumentsServer
	logger hclog.Logger
}

func (m *ArgumentsServer) Pre(ctx context.Context, req *proto.Empty) (*proto.Empty, error) {
	return &proto.Empty{}, m.Impl.Pre()
}

func (m *ArgumentsServer) Get(ctx context.Context, req *proto.Empty) (*proto.ArgumentsMsg, error) {
	args, err := m.Impl.Get()
	return args, err
}

func (m *ArgumentsServer) Post(ctx context.Context, req *proto.ArgumentsMsg) (*proto.ArgumentsMsg, error) {
	argsMsg, err := m.Impl.Post(req)
	// TODO: Turn shared Arguments into proto ArgumentsMsg
	return argsMsg, err
}

// ProfilesClient is an implementation of KV that talks over RPC.
type ProfilesClient struct{ client proto.ProfilesClient }

func (m *ProfilesClient) Pre() error {
	_, err := m.client.Pre(context.Background(), &proto.Empty{})
	return err
}

func (m *ProfilesClient) Get() (*proto.ProfilesMsg, error) {
	_, err := m.client.Get(context.Background(), &proto.Empty{})
	if err != nil {
		return &proto.ProfilesMsg{}, err
	}

	return &proto.ProfilesMsg{}, nil
}

func (m *ProfilesClient) Post(p *proto.ProfilesMsg) error {
	_, err := m.client.Post(context.Background(), &proto.ProfilesMsg{})
	if err != nil {
		return err
	}

	return nil
}

// Here is the gRPC server that ProfilesClient talks to.
type ProfilesServer struct {
	// This is the real implementation
	Impl ProfilesService
	proto.UnimplementedProfilesServer
}

func (m *ProfilesServer) Pre(
	ctx context.Context,
	req *proto.Empty,
) (*proto.Empty, error) {
	return &proto.Empty{}, m.Impl.Pre()
}

func (m *ProfilesServer) Get(
	ctx context.Context,
	req *proto.Empty,
) (*proto.ProfilesMsg, error) {
	_, err := m.Impl.Get()
	return &proto.ProfilesMsg{}, err
}

func (m *ProfilesServer) Post(
	ctx context.Context,
	req *proto.ProfilesMsg,
) (*proto.Empty, error) {
	err := m.Impl.Post(&proto.ProfilesMsg{})
	return &proto.Empty{}, err
}

// CredentialsClient is an implementation of KV that talks over RPC.
type CredentialsClient struct{ client proto.CredentialsClient }

func (m *CredentialsClient) Pre() error {
	_, err := m.client.Pre(context.Background(), &proto.Empty{})
	return err
}

func (m *CredentialsClient) Get() error {
	_, err := m.client.Get(context.Background(), &proto.Empty{})
	if err != nil {
		return err
	}

	return nil
}

func (m *CredentialsClient) Post() error {
	_, err := m.client.Post(context.Background(), &proto.Empty{})
	if err != nil {
		return err
	}

	return nil
}

// Here is the gRPC server that ProfilesClient talks to.
type CredentialsServer struct {
	// This is the real implementation
	Impl CredentialsService
	proto.UnimplementedCredentialsServer
}

func (m *CredentialsServer) Pre(
	ctx context.Context,
	req *proto.Empty,
) (*proto.Empty, error) {
	return &proto.Empty{}, m.Impl.Pre()
}

func (m *CredentialsServer) Get(
	ctx context.Context,
	req *proto.Empty,
) (*proto.Empty, error) {
	return &proto.Empty{}, m.Impl.Get()
}

func (m *CredentialsServer) Post(
	ctx context.Context,
	req *proto.Empty,
) (*proto.Empty, error) {
	return &proto.Empty{}, m.Impl.Post()
}

// ProfileNamesClient is an implementation of KV that talks over RPC.
type ProfileNamesClient struct{ client proto.ProfileNamesClient }

func (m *ProfileNamesClient) Get() error {
	_, err := m.client.Get(context.Background(), &proto.Empty{})
	if err != nil {
		return err
	}

	return nil
}

// Here is the gRPC server that ProfilesClient talks to.
type ProfileNamesServer struct {
	// This is the real implementation
	Impl ProfileNamesService
	proto.UnimplementedProfileNamesServer
}

func (m *ProfileNamesServer) Get(
	ctx context.Context,
	req *proto.Empty,
) (*proto.Empty, error) {
	return &proto.Empty{}, m.Impl.Get()
}
