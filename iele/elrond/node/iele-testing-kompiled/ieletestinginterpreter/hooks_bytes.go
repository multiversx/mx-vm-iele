// File provided by the K Framework Go backend. Timestamp: 2019-08-24 18:56:17.501

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"
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
	bytes, ok1 := interpreter.Model.GetBytes(argBytes)
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

	if len(bytes) == 0 {
		return m.IntZero, nil
	}

	var beBytes []byte // will be the big endian version of the input
	if littleEndian {
		beBytes = reverseBytes(bytes) // will return a fresh slice
	} else {
		beBytes = make([]byte, len(bytes))
		copy(beBytes, bytes) // copy slice, avoid altering the input
	}

	negative := false
	if signedBytes {
		if beBytes[0]>>7 == 1 { // most significant bit is 1, this signals a 2's complement negative number
			negative = true
			for i := range beBytes {
				beBytes[i] = ^beBytes[i]
			}
		}
	}

	result := new(big.Int)
	result.SetBytes(beBytes)

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

func (bytesHooksType) bytes2string(bytesRef m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	strRef, ok := interpreter.Model.Bytes2String(bytesRef)
	if !ok {
		return invalidArgsResult()
	}
	return strRef, nil
}

func (bytesHooksType) string2bytes(strRef m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	bytesRef, ok := interpreter.Model.String2Bytes(strRef)
	if !ok {
		return invalidArgsResult()
	}
	return bytesRef, nil
}

func (bytesHooksType) substr(kbytes, kfrom, kto m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	from, ok2 := interpreter.Model.GetUint64(kfrom) // from is inclusive
	to, ok3 := interpreter.Model.GetUint64(kto)     // to is exclusive
	if !ok2 || !ok3 {
		return invalidArgsResult()
	}
	subStr, ok := m.BytesSub(kbytes, from, to)
	if !ok {
		return invalidArgsResult()
	}
	return subStr, nil
}

func (bytesHooksType) replaceAt(argBytes m.KReference, argOffset m.KReference, argReplacement m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	bytes, ok1 := interpreter.Model.GetBytes(argBytes)
	offset, ok2 := interpreter.Model.GetPositiveInt(argOffset)
	krepl, ok3 := interpreter.Model.GetBytes(argReplacement)
	if !ok1 || !ok2 || !ok3 {
		return invalidArgsResult()
	}
	if offset+len(krepl) > len(bytes) {
		return invalidArgsResult()
	}
	result := make([]byte, len(bytes))
	copy(result, bytes)
	copy(result[offset:], krepl)

	return interpreter.Model.NewBytes(result), nil
}

func (bytesHooksType) length(argBytes m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	length, ok := m.BytesLength(argBytes)
	if !ok {
		return invalidArgsResult()
	}
	return interpreter.Model.FromUint64(length), nil
}

func (bytesHooksType) padRight(argBytes m.KReference, argLen m.KReference, argWith m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	bytes, ok1 := interpreter.Model.GetBytes(argBytes)
	length, ok2 := interpreter.Model.GetPositiveInt(argLen)
	padByte, ok3 := interpreter.Model.GetByte(argWith)
	if !ok1 || !ok2 || !ok3 {
		return invalidArgsResult()
	}
	if length <= len(bytes) {
		return argBytes, nil
	}
	result := make([]byte, length)
	for i := 0; i < len(bytes); i++ {
		result[i] = bytes[i]
	}
	for i := len(bytes); i < length; i++ {
		result[i] = padByte
	}

	return interpreter.Model.NewBytes(result), nil
}

func (bytesHooksType) padLeft(argBytes m.KReference, argLen m.KReference, argWith m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	bytes, ok1 := interpreter.Model.GetBytes(argBytes)
	length, ok2 := interpreter.Model.GetPositiveInt(argLen)
	padByte, ok3 := interpreter.Model.GetByte(argWith)
	if !ok1 || !ok2 || !ok3 {
		return invalidArgsResult()
	}
	if length <= len(bytes) {
		return argBytes, nil
	}
	result := make([]byte, length)
	offset := len(bytes) - length
	for i := 0; i < length; i++ {
		j := offset + i
		if j < 0 {
			result[i] = padByte
		} else {
			result[i] = bytes[j]
		}
	}

	return interpreter.Model.NewBytes(result), nil
}

func (bytesHooksType) reverse(argBytes m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	bytes, ok := interpreter.Model.GetBytes(argBytes)
	if !ok {
		return invalidArgsResult()
	}
	kbLen := len(bytes)
	if kbLen == 0 {
		return m.BytesEmpty, nil
	}
	revBytes := reverseBytes(bytes)
	return interpreter.Model.NewBytes(revBytes), nil
}

func (bytesHooksType) concat(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	bytes1, ok1 := interpreter.Model.GetBytes(c1)
	bytes2, ok2 := interpreter.Model.GetBytes(c2)
	if !ok1 || !ok2 {
		return invalidArgsResult()
	}
	if len(bytes2) == 0 {
		return c1, nil
	}
	if len(bytes1) == 0 {
		return c2, nil
	}

	return interpreter.Model.NewBytes(append(bytes1, bytes2...)), nil
}
