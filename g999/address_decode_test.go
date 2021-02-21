/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package g999

import (
	"encoding/hex"
	"testing"
)

func TestAddressDecoder_AddressEncode(t *testing.T) {

	p2pk, _ := hex.DecodeString("164fb7968b11c173b2cc9e683d0f46b251770158")
	p2pkAddr, _ := tw.DecoderV2.AddressEncode(p2pk)
	t.Logf("p2pkAddr: %s", p2pkAddr)
}

func TestAddressDecoder_AddressDecode(t *testing.T) {

	p2pkAddr := "GKstnC4KHbVWFL8F37wZgsRxFUDEBN3tUo"
	p2pkHash, _ := tw.DecoderV2.AddressDecode(p2pkAddr)
	t.Logf("p2pkHash: %s", hex.EncodeToString(p2pkHash))
}
