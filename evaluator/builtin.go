package evaluator

import "interpreter/object"

var builtins map[string]*object.Builtin

func initBuiltins() {
	builtins = map[string]*object.Builtin{
		"len": {
			Fn: func(args ...object.Object) object.Object {
				if len(args) != 1 {
					return newError("wrong number of arguments. got=%d, want=1", len(args))
				}

				switch arg := args[0].(type) {
				case *object.String:
					return &object.Integer{Value: int64(len(arg.Value))}
				case *object.Array:
					return &object.Integer{Value: int64(len(arg.Elements))}
				default:
					return newError("argument to `len` not supported, got %s", args[0].Type())
				}
			},
		},

		"first": {
			Fn: func(args ...object.Object) object.Object {
				if len(args) != 1 {
					return newError("wrong number of arguments. got=%d, want=1", len(args))
				}

				if args[0].Type() != object.ARRAY_OBJ {
					return newError("argument to `first` must be ARRAY, got %s", args[0].Type())
				}

				arr := args[0].(*object.Array)
				if len(arr.Elements) > 0 {
					return arr.Elements[0]
				}

				return NULL
			},
		},

		"last": {
			Fn: func(args ...object.Object) object.Object {
				if len(args) != 1 {
					return newError("wrong number of arguments. got=%d, want=1", len(args))
				}

				if args[0].Type() != object.ARRAY_OBJ {
					return newError("argument to `first` must be ARRAY, got %s", args[0].Type())
				}
				arr := args[0].(*object.Array)
				length := len(arr.Elements)
				if length > 0 {
					return arr.Elements[length-1]
				}
				return NULL
			},
		},

		"rest": {
			Fn: func(args ...object.Object) object.Object {
				if len(args) != 1 {
					return newError("wrong number of arguments. got=%d, want=1", len(args))
				}

				if args[0].Type() != object.ARRAY_OBJ {
					return newError("argument to `rest` must be ARRAY, got %s", args[0].Type())
				}

				arr := args[0].(*object.Array)
				length := len(arr.Elements)
				if length > 0 {
					newElements := make([]object.Object, length-1, length-1)
					copy(newElements, arr.Elements[1:length])
					return &object.Array{Elements: newElements}
				}
				return NULL
			},
		},

		"push": {
			Fn: func(args ...object.Object) object.Object {
				if len(args) != 2 {
					return newError("wrong number of arguments. got=%d, want=2", len(args))
				}

				if args[0].Type() != object.ARRAY_OBJ {
					return newError("argument to `push` must be ARRAY, got %s", args[0].Type())
				}

				arr := args[0].(*object.Array)
				length := len(arr.Elements)
				newElements := make([]object.Object, length+1, length+1)
				copy(newElements, arr.Elements)
				newElements[length] = args[1]
				return &object.Array{Elements: newElements}

			},
		},

		"map": {
			Fn: func(args ...object.Object) object.Object {
				if len(args) != 2 {
					return newError("wrong number of arguments. got=%d, want=2", len(args))
				}

				if args[0].Type() != object.ARRAY_OBJ {
					return newError("first argument to `map` must be ARRAY, got %s", args[0].Type())
				}

				if args[1].Type() != object.FUNCTION_OBJ {
					return newError("second argument to `map` must be FUNCTION, got %s", args[1].Type())
				}

				arr := args[0].(*object.Array)
				fn := args[1].(*object.Function)

				length := len(arr.Elements)
				newElements := make([]object.Object, length, length)

				for i, el := range arr.Elements {
					evaluated := applyFunction(fn, []object.Object{el})
					if isError(evaluated) {
						return evaluated
					}
					newElements[i] = evaluated
				}

				return &object.Array{Elements: newElements}
			},
		},

		"reduce": {
			Fn: func(args ...object.Object) object.Object {
				if len(args) != 3 {
					return newError("wrong number of arguments. got=%d, want=3", len(args))
				}

				if args[0].Type() != object.ARRAY_OBJ {
					return newError("first argument to `reduce` must be ARRAY, got %s", args[0].Type())
				}
				if args[1].Type() != object.INTEGER_OBJ {
					return newError("second argument to `reduce` must be INTEGER, got %s", args[1].Type())
				}

				if args[2].Type() != object.FUNCTION_OBJ {
					return newError("third argument to `reduce` must be FUNCTION, got %s", args[2].Type())
				}

				arr := args[0].(*object.Array)
				acc := args[1].(*object.Integer).Value
				fn := args[2].(*object.Function)

				for _, el := range arr.Elements {
					evaluated := applyFunction(fn, []object.Object{&object.Integer{Value: acc}, el})
					if isError(evaluated) {
						return evaluated
					}
					acc = evaluated.(*object.Integer).Value
				}

				return &object.Integer{Value: acc}
			},
		},
	}
}

func init() {
	initBuiltins()
}
