all: one_dns

one_dns: one_dns.go Makefile
	go build one_dns.go