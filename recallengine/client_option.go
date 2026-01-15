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

func WithRequestHeader(key string, value string) ClientOption {
	return func(e *Client) {
		if e.RequestHeaders == nil {
			e.RequestHeaders = make(map[string]string)
		}
		e.RequestHeaders[key] = value
	}
}

func WithRetryTimes(times int) ClientOption {
	return func(e *Client) {
		e.RetryTimes = times
	}
}
