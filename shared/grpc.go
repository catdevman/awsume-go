package shared

import (
	"context"

	"github.com/catdevman/awsume-go/proto"
)

// ProfilesClient is an implementation of KV that talks over RPC.
type ProfilesClient struct{ client proto.ProfilesClient }

func (m *ProfilesClient) Pre() error {
	_, err := m.client.Pre(context.Background(), &proto.Empty{})
	return err
}

func (m *ProfilesClient) Get() (Profiles, error) {
	_, err := m.client.Get(context.Background(), &proto.Empty{})
	if err != nil {
		return Profiles{}, err
	}

	return Profiles{}, nil
}
func (m *ProfilesClient) Post(p Profiles) error {
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
	req *proto.Empty) (*proto.Empty, error) {
	return &proto.Empty{}, m.Impl.Pre()
}

func (m *ProfilesServer) Get(
	ctx context.Context,
	req *proto.Empty) (*proto.ProfilesMsg, error) {
	_, err := m.Impl.Get()
	return &proto.ProfilesMsg{}, err
}
func (m *ProfilesServer) Post(
	ctx context.Context,
	req *proto.ProfilesMsg) (*proto.Empty, error) {
	err := m.Impl.Post(Profiles{})
	return &proto.Empty{}, err
}

// ArgumentsClient is an implementation of KV that talks over RPC.
type ArgumentsClient struct{ client proto.ArgumentsClient }

func (m *ArgumentsClient) Pre() error {
	_, err := m.client.Pre(context.Background(), &proto.Empty{})
	return err
}

func (m *ArgumentsClient) Get() error {
	_, err := m.client.Get(context.Background(), &proto.Empty{})
	if err != nil {
		return err
	}

	return nil
}
func (m *ArgumentsClient) Post() error {
	_, err := m.client.Get(context.Background(), &proto.Empty{})
	if err != nil {
		return err
	}

	return nil
}

// Here is the gRPC server that ProfilesClient talks to.
type ArgumentsServer struct {
	// This is the real implementation
	Impl ArgumentsService
	proto.UnimplementedArgumentsServer
}

func (m *ArgumentsServer) Pre(
	ctx context.Context,
	req *proto.Empty) (*proto.Empty, error) {
	return &proto.Empty{}, m.Impl.Pre()
}

func (m *ArgumentsServer) Get(
	ctx context.Context,
	req *proto.Empty) (*proto.ArgumentsMsg, error) {
	err := m.Impl.Get()
	return &proto.ArgumentsMsg{}, err
}
func (m *ArgumentsServer) Post(
	ctx context.Context,
	req *proto.ArgumentsMsg) (*proto.ArgumentsMsg, error) {
	err := m.Impl.Post()
	return &proto.ArgumentsMsg{}, err
}
