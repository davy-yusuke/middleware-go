package middleware

import (
	"fmt"
	"regexp"
	"strings"
)

type regexRoute struct {
	methodBit int
	pattern   string
	re        *regexp.Regexp
	names     []string
	handlers  []HandlerFunc
}

type RegexRouter struct {
	routes      []regexRoute
	middlewares []HandlerFunc
}

const maxRegexPatternLen = 2048

var nestedQuantifier = regexp.MustCompile(`\(\?:?.*[\+\*\{].*\).*[+\*\{]`)

func NewRegexRouter() *RegexRouter { return &RegexRouter{} }

func (rr *RegexRouter) Use(m HandlerFunc) {
	rr.middlewares = append(rr.middlewares, m)
}

func (rr *RegexRouter) AddRoute(methodBit int, pattern string, handlers []HandlerFunc) error {
	if len(pattern) > maxRegexPatternLen {
		return fmt.Errorf("pattern too long")
	}

	p := strings.Trim(pattern, "/")
	var segs []string
	if p == "" {
		segs = []string{}
	} else {
		segs = strings.Split(p, "/")
	}
	parts := make([]string, 0, len(segs))
	names := make([]string, 0)
	for _, s := range segs {
		if strings.HasPrefix(s, ":") {
			parts = append(parts, "([^/]+)")
			names = append(names, s[1:])
		} else if s == "*" {
			parts = append(parts, "(.*)")
			names = append(names, "*")
		} else {
			parts = append(parts, regexp.QuoteMeta(s))
		}
	}
	regexStr := "^/"
	if len(parts) > 0 {
		regexStr += strings.Join(parts, "/")
	}
	regexStr += "$"

	if nestedQuantifier.MatchString(regexStr) {
		return fmt.Errorf("pattern rejected: nested quantifier suspects")
	}
	re, err := regexp.Compile(regexStr)
	if err != nil {
		return err
	}
	rr.routes = append(rr.routes, regexRoute{
		methodBit: methodBit,
		pattern:   pattern,
		re:        re,
		names:     names,
		handlers:  handlers,
	})
	return nil
}

func (rr *RegexRouter) Find(methodBit int, p string) ([]HandlerFunc, map[string]string, bool) {

	for _, rt := range rr.routes {
		if rt.methodBit != methodBit {
			continue
		}
		if m := rt.re.FindStringSubmatch(p); m != nil {
			params := map[string]string{}
			for i, name := range rt.names {
				if 1+i < len(m) {
					params[name] = m[1+i]
				}
			}
			all := make([]HandlerFunc, 0, len(rr.middlewares)+len(rt.handlers))
			all = append(all, rr.middlewares...)
			all = append(all, rt.handlers...)
			return all, params, true
		}
	}
	return nil, nil, false
}
