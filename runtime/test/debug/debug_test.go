// debug test is a convenient package
// you can paste your minimal code your
// to focus only the problemtic part of
// failing code

package debug

import (
	"context"
	"testing"

	"github.com/xhd2015/xgo/runtime/core"
	"github.com/xhd2015/xgo/runtime/mock"
	"github.com/xhd2015/xgo/runtime/test/mock_var/sub"
)

func TestMockVarInOtherPkg(t *testing.T) {
	mock.Mock(&sub.A, func(ctx context.Context, fn *core.FuncInfo, args, results core.Object) error {
		results.GetFieldIndex(0).Set("mockA")
		return nil
	})
	b := sub.A
	if b != "mockA" {
		t.Fatalf("expect sub.A to be %s, actual: %s", "mockA", b)
	}
}
