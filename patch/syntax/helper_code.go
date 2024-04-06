//go:build ignore
// +build ignore

package syntax

const __xgo_local_pkg_name = "" // filled later

type __xgo_local_func_stub struct {
	PkgPath      string
	Kind         int // 0 = func, 1 = var, 2=var_ptr 3 = const
	Fn           interface{}
	Var          interface{} // pointer to a variable if this is a declare variable
	PC           uintptr     // filled later
	Interface    bool
	Generic      bool
	Closure      bool // is the given function a closure
	RecvTypeName string
	RecvPtr      bool
	Name         string
	IdentityName string // name without pkgPath

	RecvName string
	ArgNames []string
	ResNames []string

	// Deprecated
	// these two fields can be retrieved at runtime
	FirstArgCtx bool // first argument is context.Context or sub type?
	// Deprecated
	LastResErr bool // last res is error or sub type?

	File string
	Line int
}

func init() {
	__xgo_link_generate_init_regs_body()
}

// TODO: ensure safety for this
func __xgo_link_generate_init_regs_body() {
	// linked later by compiler
	// NOTE: if the function is called by init, don't include
	// any extra statements, they will be executed
	// no matther whether you have replaced it or not
	// panic("failed to link __xgo_link_generate_init_regs_body")
}

func __xgo_link_trap_for_generated(pkgPath string, pc uintptr, identityName string, generic bool, recv interface{}, args []interface{}, results []interface{}) (func(), bool) {
	// linked by compiler
	return nil, false
}
func __xgo_link_trap_var_for_generated(pkgPath string, name string, tmpVarAddr interface{}, takeAddr bool) {
	// linked by compiler
}

func __xgo_link_generated_register_func(fn interface{}) {
	// linked later by compiler
	panic("failed to link __xgo_link_generated_register_func")
}

func __xgo_local_register_func(pkgPath string, identityName string, fn interface{}, closure bool, recvName string, argNames []string, resNames []string, file string, line int) {
	__xgo_link_generated_register_func(__xgo_local_func_stub{PkgPath: pkgPath, IdentityName: identityName, Fn: fn, Closure: closure, RecvName: recvName, ArgNames: argNames, ResNames: resNames, File: file, Line: line})
}

// not used
// func __xgo_local_register_interface(pkgPath string, interfaceName string, file string, line int) {
// 	__xgo_link_generated_register_func(__xgo_local_func_stub{PkgPath: pkgPath, Interface: true, File: file, Line: line})
// }
