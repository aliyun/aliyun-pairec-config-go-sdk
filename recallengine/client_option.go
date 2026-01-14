package recallengine

type ClientOption func(c *Client)

func WithLogger(l Logger) ClientOption {
	return func(e *Client) {
		e.Logger = l
	}
}

func WithErrorLogger(l Logger) ClientOption {
	return func(e *Client) {
		e.ErrorLogger = l
	}
}
