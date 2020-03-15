package main

import (
	"testing"
)

func TestChecksum(t *testing.T) {
	ipHeader := []byte{
		0x45, 0x00, 0x00, 0x3c, 0x1c, 0x46, 0x40, 0x00, 0x40, 0x06,
		0x00, 0x00, 0xac, 0x10, 0x0a, 0x63, 0xac, 0x10, 0x0a, 0x0c,
	}

	icmpPacket := []byte{
		0x08, 0x00, // type and subtype
		0x00, 0x00, // checksum
		0x2C, 0xCC, // identifier
		0x00, 0x01, // sequence
		0x4F, 0x93, 0x65, 0x5E,
		0x00, 0x00, 0x00, 0x00,
		0x25, 0x33, 0x0A, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19,
		0x1A, 0x1B, 0x1C, 0x1D, 0x1E, 0x1F, 0x20, 0x21, 0x22, 0x23,
		0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2A, 0x2B, 0x2C, 0x2D,
		0x2E, 0x2F, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37,
	}

	tests := []struct {
		packet []byte
		want   uint16
	}{
		{packet: ipHeader, want: 0xB1E6},
		{packet: icmpPacket, want: 0x283B},
	}

	for _, tc := range tests {
		got := Checksum(tc.packet)
		if got != tc.want {
			t.Errorf("Got:% X, Want:% X", got, tc.want)
		}
	}
}
