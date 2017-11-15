package main

import (
	"fmt"
	"net"
	"os"
)

func unpackDns(msg []byte, dnsType uint16) (domain string, id uint16, values []string) {
	d := new(dnsMsg)

	if !d.Unpack(msg) {
		fmt.Fprintf(os.Stderr, "dns error (unpacking)")
		return
	}

	//fmt.Fprintf(os.Stderr, d.String())

	id = d.id

	if len(d.question) < 1 {
		// fmt.Fprintf(os.Stderr, "dns error (wrong question section)\n")
		return
	}

	domain = d.question[0].Name
	if len(domain) < 1 {
		// fmt.Fprintf(os.Stderr, "dns error (wrong domain in question)\n")
		return
	}

	_, addrs, err := answer(domain, "server", d, dnsType)
	if err == nil {
		switch dnsType {
		case dnsTypeA:
			values = ipsToStrings(convertRR_A(addrs))
		case dnsTypeAAAA:
			values = ipsToStrings(convertRR_AAAA(addrs))
		case dnsTypeNS:
			values = convertRR_Ns(addrs)
		case dnsTypeMX:
			values = convertRR_Mx(addrs)
		case dnsTypeTXT:
			values = convertRR_Txt(addrs)
		}
	}
	return
}

func ipsToStrings(ips []net.IP) (result []string) {
	result = make([]string, len(ips))
	for i, ip :=range ips  {
		result[i] = ip.String()
	}

	return
}

func packDns(domain string, id uint16, dnsType uint16) []byte {

	out := new(dnsMsg)
	out.id = id
	out.recursion_desired = true
	out.question = []dnsQuestion{
		{domain, dnsType, dnsClassINET},
	}

	msg, ok := out.Pack()
	if !ok {
		fmt.Fprintf(os.Stderr, "can't pack domain %s\n", domain)
		os.Exit(1)
	}
	return msg
}
