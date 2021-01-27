package client

import (
	"github.com/gogo/protobuf/types"
	"github.com/pachyderm/pachyderm/src/client/version"
	"github.com/pachyderm/pachyderm/src/internal/grpcutil"
)

// Version returns the version of pachd as a string.
func (c APIClient) Version() (string, error) {
	v, err := c.VersionAPIClient.GetVersion(c.Ctx(), &types.Empty{})
	if err != nil {
		return "", grpcutil.ScrubGRPC(err)
	}
	return version.PrettyPrintVersion(v), nil
}
