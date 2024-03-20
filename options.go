package main

type Options struct {
	domains     []string
	concurrency int
	probes      probeArgs
	skipDefault bool
	to          int
	preferHTTPS bool
	method      string
}

func CreateOptions(domains []string, concurrency int,
	probes probeArgs, skipDefault bool, to int, preferHTTPS bool, method string) *Options {
	return &Options{
		domains:     domains,
		concurrency: concurrency,
		probes:      probes,
		skipDefault: skipDefault,
		to:          to,
		preferHTTPS: preferHTTPS,
		method:      method,
	}
}

func CreateDefaultOptions(domains []string) *Options {
	return CreateOptions(domains, 20, []string{}, false, 10000, false, "GET")
}
