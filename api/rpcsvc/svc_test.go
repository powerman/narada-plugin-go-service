package rpcsvc

import (
	"os"
	"testing"

	"github.com/powerman/narada-go/narada"
	"github.com/powerman/narada-go/narada/staging"
)

func TestMain(m *testing.M) { os.Exit(staging.TearDown(m.Run())) }

func TestVERSION(t *testing.T) {
	api := RPC{}
	var args struct{}
	var res string
	err := api.Version(&args, &res)
	if err != nil {
		t.Errorf("Version(), err = %v", err)
	}
	if ver, _ := narada.Version(); res != ver {
		t.Errorf("Version() = %v, want %v", res, ver)
	}
}
