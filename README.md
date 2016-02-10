# httpratelimit [![Build Status](https://travis-ci.org/dougnukem/httpratelimit.png)](https://travis-ci.org/dougnukem/httpratelimit) [![Coverage](https://gocover.io/_badge/github.com/dougnukem/httpratelimit)](https://gocover.io/github.com/dougnukem/httpratelimit) [![GoDoc](https://godoc.org/github.com/dougnukem/httpratelimit?status.png)](https://godoc.org/github.com/dougnukem/httpratelimit)

allows HTTP Transport supporting rate limiting policies

inspired by: https://github.com/facebookgo/httpcontrol

# Goals

To provide HTTP rate limit policies.
- **MonitoringRateLimit** - only monitors rate limit status for debugging/MonitoringRateLimit
- **ResponseErrorRateLimit** - rate limit intercept request if exceeded and provide error HTTP response (avoid hitting real API server to avoid blacklisting and throttling)
- **ThrottlingRateLimit** - provide a throttling mechanism to spread out API requests evenly across rate limit time period
  - blocks requests and waits until it can schedule request that would result in an even request distribution across rate limit time period

Provide concrete rate limit configuration for specific HTTP API Services:
- [Twitter API Rate Limits](https://dev.twitter.com/rest/public/rate-limiting)
- [Github API Rate Limits](https://developer.github.com/v3/rate_limit/)
