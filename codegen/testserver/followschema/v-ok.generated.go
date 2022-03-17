// Code generated by github.com/00security/gqlgen, DO NOT EDIT.

package followschema

import (
	"context"
	"strconv"

	"github.com/00security/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _VOkCaseNil_value(ctx context.Context, field graphql.CollectedField, obj *VOkCaseNil) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "VOkCaseNil",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		v, ok := obj.Value()
		if !ok {
			return nil, nil
		}
		return v, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalOString2string(ctx, field.Selections, res)
}

func (ec *executionContext) _VOkCaseValue_value(ctx context.Context, field graphql.CollectedField, obj *VOkCaseValue) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "VOkCaseValue",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		v, ok := obj.Value()
		if !ok {
			return nil, nil
		}
		return v, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalOString2string(ctx, field.Selections, res)
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var vOkCaseNilImplementors = []string{"VOkCaseNil"}

func (ec *executionContext) _VOkCaseNil(ctx context.Context, sel ast.SelectionSet, obj *VOkCaseNil) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, vOkCaseNilImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("VOkCaseNil")
		case "value":
			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				return ec._VOkCaseNil_value(ctx, field, obj)
			}

			out.Values[i] = innerFunc(ctx)

		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var vOkCaseValueImplementors = []string{"VOkCaseValue"}

func (ec *executionContext) _VOkCaseValue(ctx context.Context, sel ast.SelectionSet, obj *VOkCaseValue) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, vOkCaseValueImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("VOkCaseValue")
		case "value":
			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				return ec._VOkCaseValue_value(ctx, field, obj)
			}

			out.Values[i] = innerFunc(ctx)

		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) marshalOVOkCaseNil2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐVOkCaseNil(ctx context.Context, sel ast.SelectionSet, v *VOkCaseNil) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._VOkCaseNil(ctx, sel, v)
}

func (ec *executionContext) marshalOVOkCaseValue2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐVOkCaseValue(ctx context.Context, sel ast.SelectionSet, v *VOkCaseValue) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._VOkCaseValue(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************
