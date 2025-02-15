// Code generated by github.com/00security/gqlgen, DO NOT EDIT.

package followschema

import (
	"context"
	"strconv"

	"github.com/00security/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

type MutationResolver interface {
	DefaultInput(ctx context.Context, input DefaultInput) (*DefaultParametersMirror, error)
	UpdateSomething(ctx context.Context, input SpecialInput) (string, error)
	UpdatePtrToPtr(ctx context.Context, input UpdatePtrToPtrOuter) (*PtrToPtrOuter, error)
}

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

func (ec *executionContext) field_Mutation_defaultInput_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 DefaultInput
	if tmp, ok := rawArgs["input"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
		arg0, err = ec.unmarshalNDefaultInput2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐDefaultInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_updatePtrToPtr_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 UpdatePtrToPtrOuter
	if tmp, ok := rawArgs["input"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
		arg0, err = ec.unmarshalNUpdatePtrToPtrOuter2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐUpdatePtrToPtrOuter(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_updateSomething_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 SpecialInput
	if tmp, ok := rawArgs["input"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
		arg0, err = ec.unmarshalNSpecialInput2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐSpecialInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
	return args, nil
}

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _DefaultParametersMirror_falsyBoolean(ctx context.Context, field graphql.CollectedField, obj *DefaultParametersMirror) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "DefaultParametersMirror",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.FalsyBoolean, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*bool)
	fc.Result = res
	return ec.marshalOBoolean2ᚖbool(ctx, field.Selections, res)
}

func (ec *executionContext) _DefaultParametersMirror_truthyBoolean(ctx context.Context, field graphql.CollectedField, obj *DefaultParametersMirror) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "DefaultParametersMirror",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.TruthyBoolean, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*bool)
	fc.Result = res
	return ec.marshalOBoolean2ᚖbool(ctx, field.Selections, res)
}

func (ec *executionContext) _Mutation_defaultInput(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Mutation_defaultInput_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().DefaultInput(rctx, args["input"].(DefaultInput))
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*DefaultParametersMirror)
	fc.Result = res
	return ec.marshalNDefaultParametersMirror2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐDefaultParametersMirror(ctx, field.Selections, res)
}

func (ec *executionContext) _Mutation_updateSomething(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Mutation_updateSomething_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().UpdateSomething(rctx, args["input"].(SpecialInput))
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Mutation_updatePtrToPtr(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Mutation_updatePtrToPtr_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp := ec._fieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().UpdatePtrToPtr(rctx, args["input"].(UpdatePtrToPtrOuter))
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*PtrToPtrOuter)
	fc.Result = res
	return ec.marshalNPtrToPtrOuter2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐPtrToPtrOuter(ctx, field.Selections, res)
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputDefaultInput(ctx context.Context, obj interface{}) (DefaultInput, error) {
	var it DefaultInput
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	if _, present := asMap["falsyBoolean"]; !present {
		asMap["falsyBoolean"] = false
	}
	if _, present := asMap["truthyBoolean"]; !present {
		asMap["truthyBoolean"] = true
	}

	for k, v := range asMap {
		switch k {
		case "falsyBoolean":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("falsyBoolean"))
			it.FalsyBoolean, err = ec.unmarshalOBoolean2ᚖbool(ctx, v)
			if err != nil {
				return it, err
			}
		case "truthyBoolean":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("truthyBoolean"))
			it.TruthyBoolean, err = ec.unmarshalOBoolean2ᚖbool(ctx, v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var defaultParametersMirrorImplementors = []string{"DefaultParametersMirror"}

func (ec *executionContext) _DefaultParametersMirror(ctx context.Context, sel ast.SelectionSet, obj *DefaultParametersMirror) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, defaultParametersMirrorImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("DefaultParametersMirror")
		case "falsyBoolean":
			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				return ec._DefaultParametersMirror_falsyBoolean(ctx, field, obj)
			}

			out.Values[i] = innerFunc(ctx)

		case "truthyBoolean":
			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				return ec._DefaultParametersMirror_truthyBoolean(ctx, field, obj)
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

var mutationImplementors = []string{"Mutation"}

func (ec *executionContext) _Mutation(ctx context.Context, sel ast.SelectionSet) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, mutationImplementors)
	ctx = graphql.WithFieldContext(ctx, &graphql.FieldContext{
		Object: "Mutation",
	})

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		innerCtx := graphql.WithRootFieldContext(ctx, &graphql.RootFieldContext{
			Object: field.Name,
			Field:  field,
		})

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Mutation")
		case "defaultInput":
			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_defaultInput(ctx, field)
			}

			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, innerFunc)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "updateSomething":
			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_updateSomething(ctx, field)
			}

			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, innerFunc)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "updatePtrToPtr":
			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_updatePtrToPtr(ctx, field)
			}

			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, innerFunc)

			if out.Values[i] == graphql.Null {
				invalids++
			}
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

func (ec *executionContext) unmarshalNDefaultInput2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐDefaultInput(ctx context.Context, v interface{}) (DefaultInput, error) {
	res, err := ec.unmarshalInputDefaultInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNDefaultParametersMirror2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐDefaultParametersMirror(ctx context.Context, sel ast.SelectionSet, v DefaultParametersMirror) graphql.Marshaler {
	return ec._DefaultParametersMirror(ctx, sel, &v)
}

func (ec *executionContext) marshalNDefaultParametersMirror2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐDefaultParametersMirror(ctx context.Context, sel ast.SelectionSet, v *DefaultParametersMirror) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	return ec._DefaultParametersMirror(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************
