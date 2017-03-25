package songo

import "net/url"

type Songo interface {
	ParseURL(u *url.URL) error
	ParseRawURL(rawurl string) error

	Least(key string, value interface{}) Songo
	Must(key string, value interface{}) Songo

	Include(key string) Songo
	Exclude(key string) Songo
}
