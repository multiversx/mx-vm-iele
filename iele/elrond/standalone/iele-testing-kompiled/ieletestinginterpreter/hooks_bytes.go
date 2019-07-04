// File provided by the K Framework Go backend. Timestamp: 2019-07-04 13:14:15.638

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/standalone/iele-testing-kompiled/ieletestingmodel"
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

func (bytesHooksType) empty(lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.BytesEmpty, nil
}

func (bytesHooksType) bytes2int(argBytes m.KReference, argEndian m.KReference, argSigned m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	kbytes, ok1 := interpreter.Model.GetBytesObject(argBytes)
	kappEndian, ok2 := interpreter.Model.GetKApplyObject(argEndian)
	kappSigned, ok3 := interpreter.Model.GetKApplyObject(argSigned)
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
		result.Add(result, bigIntOne)
		result.Neg(result)
	}

	return interpreter.Model.FromBigInt(result), nil
}

func (bytesHooksType) int2bytes(argLen m.KReference, argI m.KReference, argEndian m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	resLen, ok1 := interpreter.Model.GetPositiveInt(argLen)
	kint, ok2 := interpreter.Model.GetBigInt(argI)
	kappEndian, ok3 := interpreter.Model.GetKApplyObject(argEndian)
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

	if resLen == 0 {
		return m.BytesEmpty, nil // len = 0 means 0 length result
	}

	var resultBytes []byte
	if kint.Sign() < 0 {
		// compute 2's complement
		kintPlus1 := big.NewInt(0)
		kintPlus1.Add(kint, bigIntOne) // add 1
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
		originalBytes := kint.Bytes()
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

	return interpreter.Model.NewBytes(resultBytes), nil
}

func (bytesHooksType) bytes2string(arg m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	kbytes, ok := interpreter.Model.GetBytesObject(arg)
	if !ok {
		return invalidArgsResult()
	}
	return interpreter.Model.NewString(string(kbytes.Value)), nil
}

func (bytesHooksType) string2bytes(arg m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	kstr, ok := interpreter.Model.GetStringObject(arg)
	if !ok {
		return invalidArgsResult()
	}
	return interpreter.Model.NewBytes([]byte(kstr.Value)), nil
}

func (bytesHooksType) substr(argBytes m.KReference, argOffset1 m.KReference, argOffset2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	kbytes, ok1 := interpreter.Model.GetBytesObject(argBytes)
	offset1, ok2 := interpreter.Model.GetPositiveInt(argOffset1)
	offset2, ok3 := interpreter.Model.GetPositiveInt(argOffset2)
	if !ok1 || !ok2 || !ok3 {
		return invalidArgsResult()
	}
	length := len(kbytes.Value)

	if offset1 > offset2 || offset1 > length || offset2 > length {
		return invalidArgsResult()
	}

	if offset1 == offset2 {
		return m.BytesEmpty, nil
	}
	if offset1 == 0 && offset2 == length {
		return argBytes, nil
	}

	return interpreter.Model.NewBytes(kbytes.Value[offset1:offset2]), nil
}

func (bytesHooksType) replaceAt(argBytes m.KReference, argOffset m.KReference, argReplacement m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	kbytes, ok1 := interpreter.Model.GetBytesObject(argBytes)
	offset, ok2 := interpreter.Model.GetPositiveInt(argOffset)
	krepl, ok3 := interpreter.Model.GetBytesObject(argReplacement)
	if !ok1 || !ok2 || !ok3 {
		return invalidArgsResult()
	}
	if offset+len(krepl.Value) > len(kbytes.Value) {
		return invalidArgsResult()
	}
	result := make([]byte, len(kbytes.Value))
	copy(result, kbytes.Value)
	copy(result[offset:], krepl.Value)

	return interpreter.Model.NewBytes(result), nil
}

func (bytesHooksType) length(argBytes m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	kbytes, ok := interpreter.Model.GetBytesObject(argBytes)
	if !ok {
		return invalidArgsResult()
	}
	return interpreter.Model.FromInt(len(kbytes.Value)), nil
}

func (bytesHooksType) padRight(argBytes m.KReference, argLen m.KReference, argWith m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	kbytes, ok1 := interpreter.Model.GetBytesObject(argBytes)
	length, ok2 := interpreter.Model.GetPositiveInt(argLen)
	padByte, ok3 := interpreter.Model.GetByte(argWith)
	if !ok1 || !ok2 || !ok3 {
		return invalidArgsResult()
	}
	if length <= len(kbytes.Value) {
		return argBytes, nil
	}
	result := make([]byte, length)
	for i := 0; i < len(kbytes.Value); i++ {
		result[i] = kbytes.Value[i]
	}
	for i := len(kbytes.Value); i < length; i++ {
		result[i] = padByte
	}

	return interpreter.Model.NewBytes(result), nil
}

func (bytesHooksType) padLeft(argBytes m.KReference, argLen m.KReference, argWith m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	kbytes, ok1 := interpreter.Model.GetBytesObject(argBytes)
	length, ok2 := interpreter.Model.GetPositiveInt(argLen)
	padByte, ok3 := interpreter.Model.GetByte(argWith)
	if !ok1 || !ok2 || !ok3 {
		return invalidArgsResult()
	}
	if length <= len(kbytes.Value) {
		return argBytes, nil
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

	return interpreter.Model.NewBytes(result), nil
}

func (bytesHooksType) reverse(argBytes m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	kbytes, ok := interpreter.Model.GetBytesObject(argBytes)
	if !ok {
		return invalidArgsResult()
	}
	kbLen := len(kbytes.Value)
	if kbLen == 0 {
		return m.BytesEmpty, nil
	}
	revBytes := reverseBytes(kbytes.Value)
	return interpreter.Model.NewBytes(revBytes), nil
}

func (bytesHooksType) concat(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	kbytes1, ok1 := interpreter.Model.GetBytesObject(c1)
	kbytes2, ok2 := interpreter.Model.GetBytesObject(c2)
	if !ok1 || !ok2 {
		return invalidArgsResult()
	}
	if kbytes2.IsEmpty() {
		return c1, nil
	}
	if kbytes1.IsEmpty() {
		return c2, nil
	}

	return interpreter.Model.NewBytes(append(kbytes1.Value, kbytes2.Value...)), nil
}
