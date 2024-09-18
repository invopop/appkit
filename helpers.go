package appkit

type ErrorProps struct {
	XBindError string
	Error      string
}

type ValueProps struct {
	XBindValue string
	Value      string
}

func GetError(props *ErrorProps) string {
	if props.XBindError != "" {
		return props.XBindError
	}

	if props.Error != "" {
		return "'" + props.Error + "'"
	}

	return "false"
}

func GetErrorClasses(props *ErrorProps) string {
	err := GetError(&ErrorProps{props.XBindError, props.Error})
	classes := "{" +
		"'text-danger-500 border-danger-400 outline-danger-400': " + err + "," +
		"'border-neutral-200 hover:enabled:border-neutral-300 text-neutral-800 outline-none focus:border-workspace-accent-500 group-hover:focus:border-workspace-accent-500 caret-workspace-accent-500 focus:shadow-active': !" + err +
		"}"
	return classes
}

func GetValue(props *ValueProps) string {
	if props.XBindValue != "" {
		return props.XBindValue
	}

	if props.Value != "" {
		return "'" + props.Value + "'"
	}

	return ""
}
