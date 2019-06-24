// File provided by the K Framework Go backend. Timestamp: 2019-06-24 20:00:34.418

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/original/node/iele-testing-kompiled/ieletestingmodel"
	"math/big"
)

type bytesHooksType int

const bytesHooks bytesHooksType = 0

func reverseBytes(bytes []byte) []byte {
	length := len(bytes)
	revBytes := make([]byte, length)
	for i, b := range bytes {
		revBytes[length-1-i] = b
	}
	return revBytes
}

func (bytesHooksType) empty(lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	return m.BytesEmpty, nil
}

func (bytesHooksType) bytes2int(argBytes m.K, argEndian m.K, argSigned m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	kbytes, ok1 := argBytes.(*m.Bytes)
	kappEndian, ok2 := argEndian.(*m.KApply)
	kappSigned, ok3 := argSigned.(*m.KApply)
	if !ok1 || !ok2 || !ok3 {
		return invalidArgsResult()
	}
	if len(kappEndian.List) != 0 || len(kappSigned.List) != 0 {
		return invalidArgsResult()
	}

	var littleEndian bool
	if kappEndian.Label == m.LblLittleEndianBytes {
		littleEndian = true
	} else if kappEndian.Label == m.LblBigEndianBytes {
		littleEndian = false
	} else {
		return invalidArgsResult()
	}

	var signedBytes bool
	if kappSigned.Label == m.LblSignedBytes {
		signedBytes = true
	} else if kappSigned.Label == m.LblUnsignedBytes {
		signedBytes = false
	} else {
		return invalidArgsResult()
	}

	if kbytes.IsEmpty() {
		return m.IntZero, nil
	}

	var bytes []byte // will be the big endian version of the input
	if littleEndian {
		bytes = reverseBytes(kbytes.Value) // will return a fresh slice
	} else {
		bytes = make([]byte, len(kbytes.Value))
		copy(bytes, kbytes.Value) // copy slice, avoid altering the input
	}

	negative := false
	if signedBytes {
		if bytes[0]>>7 == 1 { // most significant bit is 1, this signals a 2's complement negative number
			negative = true
			for i := range bytes {
				bytes[i] = ^bytes[i]
			}
		}
	}

	result := new(big.Int)
	result.SetBytes(bytes)

	if negative {
		// complete 2's complement transformation
		result.Add(result, m.IntOne.Value)
		result.Neg(result)
	}

	return &m.Int{Value: result}, nil
}

func (bytesHooksType) int2bytes(argLen m.K, argI m.K, argEndian m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	klen, ok1 := argLen.(*m.Int)
	kint, ok2 := argI.(*m.Int)
	kappEndian, ok3 := argEndian.(*m.KApply)
	if !ok1 || !ok2 || !ok3 {
		return invalidArgsResult()
	}
	if len(kappEndian.List) != 0 {
		return invalidArgsResult()
	}
	var littleEndian bool
	if kappEndian.Label == m.LblLittleEndianBytes {
		littleEndian = true
	} else if kappEndian.Label == m.LblBigEndianBytes {
		littleEndian = false
	} else {
		return invalidArgsResult()
	}

	resLen := int(klen.Value.Int64())
	if resLen < 0 {
		return invalidArgsResult()
	}
	if resLen == 0 {
		return m.BytesEmpty, nil // len = 0 means 0 length result
	}

	var resultBytes []byte
	if kint.Value.Sign() < 0 {
		// compute 2's complement
		kintPlus1 := big.NewInt(0)
		kintPlus1.Add(kint.Value, m.IntOne.Value) // add 1
		kintPlus1Bytes := kintPlus1.Bytes()
		offset := len(kintPlus1Bytes) - resLen
		resultBytes = make([]byte, resLen)
		for i := 0; i < resLen; i++ {
			j := offset + i
			if j < 0 {
				resultBytes[i] = 255 // pad left with 11111111
			} else {
				resultBytes[i] = ^kintPlus1Bytes[j] // also negate every bit
			}
		}
	} else {
		originalBytes := kint.Value.Bytes()
		resultBytes = make([]byte, resLen)
		offset := len(originalBytes) - resLen
		for i := 0; i < resLen; i++ {
			j := offset + i
			if j < 0 {
				resultBytes[i] = 0 // pad left with 00000000
			} else {
				resultBytes[i] = originalBytes[j]
			}
		}
	}

	if littleEndian {
		// everything we did until here is big endian, so for little endian just reverse the byte slice
		for left, right := 0, len(resultBytes)-1; left < right; left, right = left+1, right-1 {
			resultBytes[left], resultBytes[right] = resultBytes[right], resultBytes[left]
		}
	}

	return &m.Bytes{Value: resultBytes}, nil
}

func (bytesHooksType) bytes2string(arg m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	kbytes, ok := arg.(*m.Bytes)
	if !ok {
		return invalidArgsResult()
	}
	return m.NewString(string(kbytes.Value)), nil
}

func (bytesHooksType) string2bytes(arg m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	kstr, ok := arg.(*m.String)
	if !ok {
		return invalidArgsResult()
	}
	return &m.Bytes{Value: []byte(kstr.Value)}, nil
}

func (bytesHooksType) substr(argBytes m.K, argOffset1 m.K, argOffset2 m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	kbytes, ok1 := argBytes.(*m.Bytes)
	koff1, ok2 := argOffset1.(*m.Int)
	koff2, ok3 := argOffset2.(*m.Int)
	if !ok1 || !ok2 || !ok3 {
		return invalidArgsResult()
	}
	length := len(kbytes.Value)
	offset1, off1Ok := koff1.ToPositiveInt32()
	offset2, off2Ok := koff2.ToPositiveInt32()
	if !off1Ok || !off2Ok {
		return invalidArgsResult()
	}
	if offset1 > offset2 || offset1 > length || offset2 > length {
		return invalidArgsResult()
	}

	if offset1 == offset2 {
		return m.BytesEmpty, nil
	}
	if offset1 == 0 && offset2 == length {
		return kbytes, nil
	}

	return &m.Bytes{Value: kbytes.Value[offset1:offset2]}, nil
}

func (bytesHooksType) replaceAt(argBytes m.K, argOffset m.K, argReplacement m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	kbytes, ok1 := argBytes.(*m.Bytes)
	koff, ok2 := argOffset.(*m.Int)
	krepl, ok3 := argReplacement.(*m.Bytes)
	if !ok1 || !ok2 || !ok3 {
		return invalidArgsResult()
	}
	offset, offsetOk := koff.ToPositiveInt32()
	if !offsetOk {
		return invalidArgsResult()
	}
	if offset+len(krepl.Value) > len(kbytes.Value) {
		return invalidArgsResult()
	}
	result := make([]byte, len(kbytes.Value))
	copy(result, kbytes.Value)
	copy(result[offset:], krepl.Value)

	return &m.Bytes{Value: result}, nil
}

func (bytesHooksType) length(argBytes m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	kbytes, ok := argBytes.(*m.Bytes)
	if !ok {
		return invalidArgsResult()
	}
	return m.NewIntFromInt(len(kbytes.Value)), nil
}

func (bytesHooksType) padRight(argBytes m.K, argLen m.K, argWith m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	kbytes, ok1 := argBytes.(*m.Bytes)
	klen, ok2 := argLen.(*m.Int)
	kwith, ok3 := argWith.(*m.Int)
	if !ok1 || !ok2 || !ok3 {
		return invalidArgsResult()
	}
	length, lengthOk := klen.ToPositiveInt32()
	if !lengthOk {
		return invalidArgsResult()
	}
	if length <= len(kbytes.Value) {
		return kbytes, nil
	}

	padByte, padByteOk := kwith.ToByte()
	if !padByteOk {
		return invalidArgsResult()
	}

	result := make([]byte, length)
	for i := 0; i < len(kbytes.Value); i++ {
		result[i] = kbytes.Value[i]
	}
	for i := len(kbytes.Value); i < length; i++ {
		result[i] = padByte
	}

	return &m.Bytes{Value: result}, nil
}

func (bytesHooksType) padLeft(argBytes m.K, argLen m.K, argWith m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	kbytes, ok1 := argBytes.(*m.Bytes)
	klen, ok2 := argLen.(*m.Int)
	kwith, ok3 := argWith.(*m.Int)
	if !ok1 || !ok2 || !ok3 {
		return invalidArgsResult()
	}
	length, lengthOk := klen.ToPositiveInt32()
	if !lengthOk {
		return invalidArgsResult()
	}
	if length <= len(kbytes.Value) {
		return kbytes, nil
	}

	padByte, padByteOk := kwith.ToByte()
	if !padByteOk {
		return invalidArgsResult()
	}

	result := make([]byte, length)
	offset := len(kbytes.Value) - length
	for i := 0; i < length; i++ {
		j := offset + i
		if j < 0 {
			result[i] = padByte
		} else {
			result[i] = kbytes.Value[j]
		}
	}

	return &m.Bytes{Value: result}, nil
}

func (bytesHooksType) reverse(argBytes m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	kbytes, ok := argBytes.(*m.Bytes)
	if !ok {
		return invalidArgsResult()
	}
	kbLen := len(kbytes.Value)
	if kbLen == 0 {
		return m.BytesEmpty, nil
	}
	revBytes := reverseBytes(kbytes.Value)
	return &m.Bytes{Value: revBytes}, nil
}

func (bytesHooksType) concat(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	kbytes1, ok1 := c1.(*m.Bytes)
	kbytes2, ok2 := c2.(*m.Bytes)
	if !ok1 || !ok2 {
		return invalidArgsResult()
	}
	if kbytes2.IsEmpty() {
		return kbytes1, nil
	}
	if kbytes1.IsEmpty() {
		return kbytes2, nil
	}

	return &m.Bytes{Value: append(kbytes1.Value, kbytes2.Value...)}, nil
}
