// generated by gen.go using 'go generate'; DO NOT EDIT.

// === [ Bitwise expressions ] =================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#bitwise-binary-operations

package constant

import (
	"fmt"

	"github.com/llir/llvm/ir/types"
)

// --- [ shl ] -----------------------------------------------------------------

// ExprShl represents a shift left expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#shl-instruction
type ExprShl struct {
	// Operands.
	X, Y Constant
}

// NewShl returns a new shl expression based on the given operands.
func NewShl(x, y Constant) *ExprShl {
	return &ExprShl{
		X: x,
		Y: y,
	}
}

// Type returns the type of the constant expression.
func (expr *ExprShl) Type() types.Type {
	return expr.X.Type()
}

// Ident returns the string representation of the constant expression.
func (expr *ExprShl) Ident() string {
	return fmt.Sprintf("shl (%s %s, %s %s)",
		expr.X.Type(),
		expr.X.Ident(),
		expr.Y.Type(),
		expr.Y.Ident())
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprShl) Immutable() {}

// Simplify returns a simplified version of the constant expression.
func (expr *ExprShl) Simplify() Constant {
	panic("not yet implemented")
}

// MetadataNode ensures that only metadata nodes can be assigned to the
// ir.MetadataNode interface.
func (*ExprShl) MetadataNode() {}

// --- [ lshr ] ----------------------------------------------------------------

// ExprLShr represents a logical shift right expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#lshr-instruction
type ExprLShr struct {
	// Operands.
	X, Y Constant
}

// NewLShr returns a new lshr expression based on the given operands.
func NewLShr(x, y Constant) *ExprLShr {
	return &ExprLShr{
		X: x,
		Y: y,
	}
}

// Type returns the type of the constant expression.
func (expr *ExprLShr) Type() types.Type {
	return expr.X.Type()
}

// Ident returns the string representation of the constant expression.
func (expr *ExprLShr) Ident() string {
	return fmt.Sprintf("lshr (%s %s, %s %s)",
		expr.X.Type(),
		expr.X.Ident(),
		expr.Y.Type(),
		expr.Y.Ident())
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprLShr) Immutable() {}

// Simplify returns a simplified version of the constant expression.
func (expr *ExprLShr) Simplify() Constant {
	panic("not yet implemented")
}

// MetadataNode ensures that only metadata nodes can be assigned to the
// ir.MetadataNode interface.
func (*ExprLShr) MetadataNode() {}

// --- [ ashr ] ----------------------------------------------------------------

// ExprAShr represents an arithmetic shift right expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#ashr-instruction
type ExprAShr struct {
	// Operands.
	X, Y Constant
}

// NewAShr returns a new ashr expression based on the given operands.
func NewAShr(x, y Constant) *ExprAShr {
	return &ExprAShr{
		X: x,
		Y: y,
	}
}

// Type returns the type of the constant expression.
func (expr *ExprAShr) Type() types.Type {
	return expr.X.Type()
}

// Ident returns the string representation of the constant expression.
func (expr *ExprAShr) Ident() string {
	return fmt.Sprintf("ashr (%s %s, %s %s)",
		expr.X.Type(),
		expr.X.Ident(),
		expr.Y.Type(),
		expr.Y.Ident())
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprAShr) Immutable() {}

// Simplify returns a simplified version of the constant expression.
func (expr *ExprAShr) Simplify() Constant {
	panic("not yet implemented")
}

// MetadataNode ensures that only metadata nodes can be assigned to the
// ir.MetadataNode interface.
func (*ExprAShr) MetadataNode() {}

// --- [ and ] -----------------------------------------------------------------

// ExprAnd represents an AND expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#and-instruction
type ExprAnd struct {
	// Operands.
	X, Y Constant
}

// NewAnd returns a new and expression based on the given operands.
func NewAnd(x, y Constant) *ExprAnd {
	return &ExprAnd{
		X: x,
		Y: y,
	}
}

// Type returns the type of the constant expression.
func (expr *ExprAnd) Type() types.Type {
	return expr.X.Type()
}

// Ident returns the string representation of the constant expression.
func (expr *ExprAnd) Ident() string {
	return fmt.Sprintf("and (%s %s, %s %s)",
		expr.X.Type(),
		expr.X.Ident(),
		expr.Y.Type(),
		expr.Y.Ident())
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprAnd) Immutable() {}

// Simplify returns a simplified version of the constant expression.
func (expr *ExprAnd) Simplify() Constant {
	panic("not yet implemented")
}

// MetadataNode ensures that only metadata nodes can be assigned to the
// ir.MetadataNode interface.
func (*ExprAnd) MetadataNode() {}

// --- [ or ] ------------------------------------------------------------------

// ExprOr represents an OR expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#or-instruction
type ExprOr struct {
	// Operands.
	X, Y Constant
}

// NewOr returns a new or expression based on the given operands.
func NewOr(x, y Constant) *ExprOr {
	return &ExprOr{
		X: x,
		Y: y,
	}
}

// Type returns the type of the constant expression.
func (expr *ExprOr) Type() types.Type {
	return expr.X.Type()
}

// Ident returns the string representation of the constant expression.
func (expr *ExprOr) Ident() string {
	return fmt.Sprintf("or (%s %s, %s %s)",
		expr.X.Type(),
		expr.X.Ident(),
		expr.Y.Type(),
		expr.Y.Ident())
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprOr) Immutable() {}

// Simplify returns a simplified version of the constant expression.
func (expr *ExprOr) Simplify() Constant {
	panic("not yet implemented")
}

// MetadataNode ensures that only metadata nodes can be assigned to the
// ir.MetadataNode interface.
func (*ExprOr) MetadataNode() {}

// --- [ xor ] -----------------------------------------------------------------

// ExprXor represents an exclusive-OR expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#xor-instruction
type ExprXor struct {
	// Operands.
	X, Y Constant
}

// NewXor returns a new xor expression based on the given operands.
func NewXor(x, y Constant) *ExprXor {
	return &ExprXor{
		X: x,
		Y: y,
	}
}

// Type returns the type of the constant expression.
func (expr *ExprXor) Type() types.Type {
	return expr.X.Type()
}

// Ident returns the string representation of the constant expression.
func (expr *ExprXor) Ident() string {
	return fmt.Sprintf("xor (%s %s, %s %s)",
		expr.X.Type(),
		expr.X.Ident(),
		expr.Y.Type(),
		expr.Y.Ident())
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprXor) Immutable() {}

// Simplify returns a simplified version of the constant expression.
func (expr *ExprXor) Simplify() Constant {
	panic("not yet implemented")
}

// MetadataNode ensures that only metadata nodes can be assigned to the
// ir.MetadataNode interface.
func (*ExprXor) MetadataNode() {}
