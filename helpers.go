package appkit

type ErrorProps struct {
	XBindError string
	Error      string
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
