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
	_, err := m.client.Get(context.Background(), &proto.Empty{})
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
	err := m.Impl.Get()
	return &proto.ProfilesMsg{}, err
}
func (m *ProfilesServer) Post(
	ctx context.Context,
	req *proto.ProfilesMsg) (*proto.Empty, error) {
	err := m.Impl.Post()
	return &proto.Empty{}, err
}
