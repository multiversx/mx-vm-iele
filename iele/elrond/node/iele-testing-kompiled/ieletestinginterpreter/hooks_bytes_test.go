// File provided by the K Framework Go backend. Timestamp: 2019-06-24 20:04:33.113

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"
	"testing"
)

func TestBytesEmpty(t *testing.T) {
	interpreter := newTestInterpreter()
	var bs m.K
	var err error
	bs, err = bytesHooks.empty(m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{}, bs, err, interpreter)
}

func TestBytes2Int(t *testing.T) {
	interpreter := newTestInterpreter()
	var res m.K
	var err error
	var arg1, arg2, arg3 m.K

	// unsigned
	arg1, arg2, arg3 =
		&m.Bytes{Value: []byte{1, 0}},
		&m.KApply{Label: m.LblBigEndianBytes},
		&m.KApply{Label: m.LblUnsignedBytes}
	interpreter.backupInput(arg1, arg2, arg3)
	res, err = bytesHooks.bytes2int(arg1, arg2, arg3, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertIntOk(t, "256", res, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, arg3)

	arg1, arg2, arg3 =
		&m.Bytes{Value: []byte{1, 0}},
		&m.KApply{Label: m.LblLittleEndianBytes},
		&m.KApply{Label: m.LblUnsignedBytes}
	interpreter.backupInput(arg1, arg2, arg3)
	res, err = bytesHooks.bytes2int(arg1, arg2, arg3, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertIntOk(t, "1", res, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, arg3)

	arg1, arg2, arg3 =
		&m.Bytes{Value: []byte{255}},
		&m.KApply{Label: m.LblBigEndianBytes},
		&m.KApply{Label: m.LblUnsignedBytes}
	interpreter.backupInput(arg1, arg2, arg3)
	res, err = bytesHooks.bytes2int(arg1, arg2, arg3, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertIntOk(t, "255", res, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, arg3)

	// zero
	for _, b := range [][]byte{[]byte{}, []byte{0, 0}, []byte{0, 0, 0, 0, 0, 0}} {
		for _, c2 := range []m.K{&m.KApply{Label: m.LblBigEndianBytes}, &m.KApply{Label: m.LblLittleEndianBytes}} {
			for _, c3 := range []m.K{&m.KApply{Label: m.LblUnsignedBytes}, &m.KApply{Label: m.LblSignedBytes}} {
				c1 := &m.Bytes{Value: b}
				interpreter.backupInput(c1, c2, c3)
				res, err = bytesHooks.bytes2int(c1, c2, c3, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
				assertIntOk(t, "0", res, err, interpreter)
				interpreter.checkImmutable(t, c1, c2, c3)
			}
		}
	}

	// -1
	for _, b := range [][]byte{[]byte{255}, []byte{255, 255, 255, 255, 255}} {
		for _, c2 := range []m.K{&m.KApply{Label: m.LblBigEndianBytes}, &m.KApply{Label: m.LblLittleEndianBytes}} {
			c1 := &m.Bytes{Value: b}
			c3 := &m.KApply{Label: m.LblSignedBytes}
			interpreter.backupInput(c1, c2, c3)
			res, err = bytesHooks.bytes2int(c1, c2, c3, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
			assertIntOk(t, "-1", res, err, interpreter)
			interpreter.checkImmutable(t, c1, c2, c3)
		}
	}

	// other signed negative
	arg1, arg2, arg3 =
		&m.Bytes{Value: []byte{255, 254}},
		&m.KApply{Label: m.LblBigEndianBytes},
		&m.KApply{Label: m.LblSignedBytes}
	interpreter.backupInput(arg1, arg2, arg3)
	res, err = bytesHooks.bytes2int(arg1, arg2, arg3, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertIntOk(t, "-2", res, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, arg3)

	arg1, arg2, arg3 =
		&m.Bytes{Value: []byte{255, 254}},
		&m.KApply{Label: m.LblLittleEndianBytes},
		&m.KApply{Label: m.LblSignedBytes}
	interpreter.backupInput(arg1, arg2, arg3)
	res, err = bytesHooks.bytes2int(arg1, arg2, arg3, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertIntOk(t, "-257", res, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, arg3)
}

func TestInt2Bytes(t *testing.T) {
	interpreter := newTestInterpreter()
	var bs m.K
	var err error
	kappBigEndian := &m.KApply{Label: m.LblBigEndianBytes}
	var arg1, arg2 m.K

	// length 0, empty result
	arg1, arg2 = m.NewIntFromInt(0), m.NewIntFromInt(0)
	interpreter.backupInput(arg1, arg2, kappBigEndian)
	bs, err = bytesHooks.int2bytes(arg1, arg2, kappBigEndian, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, kappBigEndian)

	arg1, arg2 = m.NewIntFromInt(0), m.NewIntFromInt(12345)
	interpreter.backupInput(arg1, arg2, kappBigEndian)
	bs, err = bytesHooks.int2bytes(arg1, arg2, kappBigEndian, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, kappBigEndian)

	arg1, arg2 = m.NewIntFromInt(0), m.NewIntFromInt(-12345)
	interpreter.backupInput(arg1, arg2, kappBigEndian)
	bs, err = bytesHooks.int2bytes(arg1, arg2, kappBigEndian, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, kappBigEndian)

	// one byte
	arg1, arg2 = m.NewIntFromInt(1), m.NewIntFromInt(0)
	interpreter.backupInput(arg1, arg2, kappBigEndian)
	bs, err = bytesHooks.int2bytes(arg1, arg2, kappBigEndian, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{0}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, kappBigEndian)

	arg1, arg2 = m.NewIntFromInt(1), m.NewIntFromInt(1)
	interpreter.backupInput(arg1, arg2, kappBigEndian)
	bs, err = bytesHooks.int2bytes(arg1, arg2, kappBigEndian, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{1}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, kappBigEndian)

	arg1, arg2 = m.NewIntFromInt(1), m.NewIntFromInt(-1)
	interpreter.backupInput(arg1, arg2, kappBigEndian)
	bs, err = bytesHooks.int2bytes(arg1, arg2, kappBigEndian, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{255}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, kappBigEndian)

	arg1, arg2 = m.NewIntFromInt(1), m.NewIntFromInt(256)
	interpreter.backupInput(arg1, arg2, kappBigEndian)
	bs, err = bytesHooks.int2bytes(arg1, arg2, kappBigEndian, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{0}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, kappBigEndian)

	arg1, arg2 = m.NewIntFromInt(1), m.NewIntFromInt(257)
	interpreter.backupInput(arg1, arg2, kappBigEndian)
	bs, err = bytesHooks.int2bytes(arg1, arg2, kappBigEndian, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{1}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, kappBigEndian)

	arg1, arg2 = m.NewIntFromInt(1), m.NewIntFromInt(-256)
	interpreter.backupInput(arg1, arg2, kappBigEndian)
	bs, err = bytesHooks.int2bytes(arg1, arg2, kappBigEndian, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{0}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, kappBigEndian)

	arg1, arg2 = m.NewIntFromInt(1), m.NewIntFromInt(-257)
	interpreter.backupInput(arg1, arg2, kappBigEndian)
	bs, err = bytesHooks.int2bytes(arg1, arg2, kappBigEndian, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{255}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, kappBigEndian)

	// 2 bytes
	arg1, arg2 = m.NewIntFromInt(2), m.NewIntFromInt(0)
	interpreter.backupInput(arg1, arg2, kappBigEndian)
	bs, err = bytesHooks.int2bytes(arg1, arg2, kappBigEndian, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{0, 0}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, kappBigEndian)

	arg1, arg2 = m.NewIntFromInt(2), m.NewIntFromInt(1)
	interpreter.backupInput(arg1, arg2, kappBigEndian)
	bs, err = bytesHooks.int2bytes(arg1, arg2, kappBigEndian, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{0, 1}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, kappBigEndian)

	arg1, arg2 = m.NewIntFromInt(2), m.NewIntFromInt(-1)
	interpreter.backupInput(arg1, arg2, kappBigEndian)
	bs, err = bytesHooks.int2bytes(arg1, arg2, kappBigEndian, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{255, 255}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, kappBigEndian)

	arg1, arg2 = m.NewIntFromInt(2), m.NewIntFromInt(256)
	interpreter.backupInput(arg1, arg2, kappBigEndian)
	bs, err = bytesHooks.int2bytes(arg1, arg2, kappBigEndian, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{1, 0}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, kappBigEndian)

	arg1, arg2 = m.NewIntFromInt(2), m.NewIntFromInt(257)
	interpreter.backupInput(arg1, arg2, kappBigEndian)
	bs, err = bytesHooks.int2bytes(arg1, arg2, kappBigEndian, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{1, 1}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, kappBigEndian)

	arg1, arg2 = m.NewIntFromInt(2), m.NewIntFromInt(-256)
	interpreter.backupInput(arg1, arg2, kappBigEndian)
	bs, err = bytesHooks.int2bytes(arg1, arg2, kappBigEndian, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{255, 0}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, kappBigEndian)

	arg1, arg2 = m.NewIntFromInt(2), m.NewIntFromInt(-257)
	interpreter.backupInput(arg1, arg2, kappBigEndian)
	bs, err = bytesHooks.int2bytes(arg1, arg2, kappBigEndian, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{254, 255}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, kappBigEndian)

	arg1, arg2 = m.NewIntFromInt(2), m.NewIntFromInt(-255)
	interpreter.backupInput(arg1, arg2, kappBigEndian)
	bs, err = bytesHooks.int2bytes(arg1, arg2, kappBigEndian, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{255, 1}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, kappBigEndian)

	// more bytes
	arg1, arg2 = m.NewIntFromInt(5), m.NewIntFromInt(0)
	interpreter.backupInput(arg1, arg2, kappBigEndian)
	bs, err = bytesHooks.int2bytes(arg1, arg2, kappBigEndian, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{0, 0, 0, 0, 0}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, kappBigEndian)

	arg1, arg2 = m.NewIntFromInt(4), m.NewIntFromInt(1)
	interpreter.backupInput(arg1, arg2, kappBigEndian)
	bs, err = bytesHooks.int2bytes(arg1, arg2, kappBigEndian, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{0, 0, 0, 1}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, kappBigEndian)

	arg1, arg2 = m.NewIntFromInt(6), m.NewIntFromInt(-1)
	interpreter.backupInput(arg1, arg2, kappBigEndian)
	bs, err = bytesHooks.int2bytes(arg1, arg2, kappBigEndian, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{255, 255, 255, 255, 255, 255}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, kappBigEndian)

	// little endian
	kappLittleEndian := &m.KApply{Label: m.LblLittleEndianBytes}

	arg1, arg2 = m.NewIntFromInt(0), m.NewIntFromInt(-12345)
	interpreter.backupInput(arg1, arg2, kappLittleEndian)
	bs, err = bytesHooks.int2bytes(arg1, arg2, kappLittleEndian, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, kappLittleEndian)

	arg1, arg2 = m.NewIntFromInt(1), m.NewIntFromInt(-1)
	interpreter.backupInput(arg1, arg2, kappLittleEndian)
	bs, err = bytesHooks.int2bytes(arg1, arg2, kappLittleEndian, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{255}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, kappLittleEndian)

	arg1, arg2 = m.NewIntFromInt(2), m.NewIntFromInt(1)
	interpreter.backupInput(arg1, arg2, kappLittleEndian)
	bs, err = bytesHooks.int2bytes(arg1, arg2, kappLittleEndian, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{1, 0}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, kappLittleEndian)

	arg1, arg2 = m.NewIntFromInt(2), m.NewIntFromInt(-256)
	interpreter.backupInput(arg1, arg2, kappLittleEndian)
	bs, err = bytesHooks.int2bytes(arg1, arg2, kappLittleEndian, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{0, 255}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, kappLittleEndian)

	arg1, arg2 = m.NewIntFromInt(4), m.NewIntFromInt(1)
	interpreter.backupInput(arg1, arg2, kappLittleEndian)
	bs, err = bytesHooks.int2bytes(arg1, arg2, kappLittleEndian, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{1, 0, 0, 0}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, kappLittleEndian)
}

func TestBytesSubstr(t *testing.T) {
	interpreter := newTestInterpreter()
	var bs, arg1, arg2, arg3 m.K
	var err error

	arg1, arg2, arg3 = &m.Bytes{Value: []byte{1, 2, 3, 4, 5}}, m.NewIntFromInt(0), m.NewIntFromInt(5)
	interpreter.backupInput(arg1, arg2, arg3)
	bs, err = bytesHooks.substr(arg1, arg2, arg3, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{1, 2, 3, 4, 5}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, arg3)

	arg1, arg2, arg3 = &m.Bytes{Value: []byte{1, 2, 3, 4, 5}}, m.NewIntFromInt(3), m.NewIntFromInt(3)
	interpreter.backupInput(arg1, arg2, arg3)
	bs, err = bytesHooks.substr(arg1, arg2, arg3, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, arg3)

	arg1, arg2, arg3 = &m.Bytes{Value: []byte{1, 2, 3, 4, 5}}, m.NewIntFromInt(0), m.NewIntFromInt(0)
	interpreter.backupInput(arg1, arg2, arg3)
	bs, err = bytesHooks.substr(arg1, arg2, arg3, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, arg3)

	arg1, arg2, arg3 = &m.Bytes{Value: []byte{1, 2, 3, 4, 5}}, m.NewIntFromInt(5), m.NewIntFromInt(5)
	interpreter.backupInput(arg1, arg2, arg3)
	bs, err = bytesHooks.substr(arg1, arg2, arg3, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, arg3)

	arg1, arg2, arg3 = &m.Bytes{Value: []byte{1, 2, 3, 4, 5}}, m.NewIntFromInt(0), m.NewIntFromInt(2)
	interpreter.backupInput(arg1, arg2, arg3)
	bs, err = bytesHooks.substr(arg1, arg2, arg3, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{1, 2}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, arg3)

	arg1, arg2, arg3 = &m.Bytes{Value: []byte{1, 2, 3, 4, 5}}, m.NewIntFromInt(4), m.NewIntFromInt(5)
	interpreter.backupInput(arg1, arg2, arg3)
	bs, err = bytesHooks.substr(arg1, arg2, arg3, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{5}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, arg3)

	arg1, arg2, arg3 = &m.Bytes{Value: []byte{1, 2, 3, 4, 5}}, m.NewIntFromInt(1), m.NewIntFromInt(5)
	interpreter.backupInput(arg1, arg2, arg3)
	bs, err = bytesHooks.substr(arg1, arg2, arg3, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{2, 3, 4, 5}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, arg3)

	arg1, arg2, arg3 = &m.Bytes{Value: []byte{1, 2, 3, 4, 5}}, m.NewIntFromInt(1), m.NewIntFromInt(4)
	interpreter.backupInput(arg1, arg2, arg3)
	bs, err = bytesHooks.substr(arg1, arg2, arg3, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{2, 3, 4}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, arg3)
}

func TestBytesReplaceAt(t *testing.T) {
	interpreter := newTestInterpreter()
	var bs, arg1, arg2, arg3 m.K
	var err error

	arg1, arg2, arg3 =
		&m.Bytes{Value: []byte{10, 20, 30, 40, 50}},
		m.NewIntFromInt(0),
		&m.Bytes{Value: []byte{11, 21, 31, 41, 51}}
	interpreter.backupInput(arg1, arg2, arg3)
	bs, err = bytesHooks.replaceAt(arg1, arg2, arg3, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{11, 21, 31, 41, 51}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, arg3)

	arg1, arg2, arg3 =
		&m.Bytes{Value: []byte{10, 20, 30, 40, 50}},
		m.NewIntFromInt(2),
		&m.Bytes{Value: []byte{33, 34, 35}}
	interpreter.backupInput(arg1, arg2, arg3)
	bs, err = bytesHooks.replaceAt(arg1, arg2, arg3, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{10, 20, 33, 34, 35}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, arg3)

	arg1, arg2, arg3 =
		m.BytesEmpty,
		m.NewIntFromInt(0),
		m.BytesEmpty
	interpreter.backupInput(arg1, arg2, arg3)
	bs, err = bytesHooks.replaceAt(arg1, arg2, arg3, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, arg3)

	arg1, arg2, arg3 =
		&m.Bytes{Value: []byte{10, 20, 30}},
		m.NewIntFromInt(1),
		&m.Bytes{Value: []byte{100}}
	interpreter.backupInput(arg1, arg2, arg3)
	bs, err = bytesHooks.replaceAt(arg1, arg2, arg3, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{10, 100, 30}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2, arg3)
}

func TestBytesLength(t *testing.T) {
	interpreter := newTestInterpreter()
	var res m.K
	var err error

	interpreter.backupInput(m.BytesEmpty)
	res, err = bytesHooks.length(m.BytesEmpty, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertIntOk(t, "0", res, err, interpreter)
	interpreter.checkImmutable(t, m.BytesEmpty)

	arg := &m.Bytes{Value: []byte{1, 2, 3, 4, 5}}
	interpreter.backupInput(arg)
	res, err = bytesHooks.length(arg, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertIntOk(t, "5", res, err, interpreter)
	interpreter.checkImmutable(t, arg)
}

func TestBytesPadRight(t *testing.T) {
	interpreter := newTestInterpreter()
	var bs, argB, argLen m.K
	var err error
	padChar := m.NewIntFromInt(80)

	argB, argLen = &m.Bytes{Value: []byte{1, 2, 3, 4, 5}}, m.NewIntFromInt(3)
	interpreter.backupInput(argB, argLen, padChar)
	bs, err = bytesHooks.padRight(argB, argLen, padChar, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{1, 2, 3, 4, 5}, bs, err, interpreter)
	interpreter.checkImmutable(t, argB, argLen, padChar)

	argB, argLen = &m.Bytes{Value: []byte{1, 2, 3, 4, 5}}, m.NewIntFromInt(5)
	interpreter.backupInput(argB, argLen, padChar)
	bs, err = bytesHooks.padRight(argB, argLen, padChar, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{1, 2, 3, 4, 5}, bs, err, interpreter)
	interpreter.checkImmutable(t, argB, argLen, padChar)

	argB, argLen = &m.Bytes{Value: []byte{1, 2, 3, 4, 5}}, m.NewIntFromInt(7)
	interpreter.backupInput(argB, argLen, padChar)
	bs, err = bytesHooks.padRight(argB, argLen, padChar, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{1, 2, 3, 4, 5, 80, 80}, bs, err, interpreter)
	interpreter.checkImmutable(t, argB, argLen, padChar)

	argB, argLen = m.BytesEmpty, m.NewIntFromInt(3)
	interpreter.backupInput(argB, argLen, padChar)
	bs, err = bytesHooks.padRight(argB, argLen, padChar, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{80, 80, 80}, bs, err, interpreter)
	interpreter.checkImmutable(t, argB, argLen, padChar)
}

func TestBytesPadLeft(t *testing.T) {
	interpreter := newTestInterpreter()
	var bs, argB, argLen m.K
	var err error
	padChar := m.NewIntFromInt(80)

	argB, argLen = &m.Bytes{Value: []byte{1, 2, 3, 4, 5}}, m.NewIntFromInt(3)
	interpreter.backupInput(argB, argLen, padChar)
	bs, err = bytesHooks.padLeft(argB, argLen, padChar, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{1, 2, 3, 4, 5}, bs, err, interpreter)
	interpreter.checkImmutable(t, argB, argLen, padChar)

	argB, argLen = &m.Bytes{Value: []byte{1, 2, 3, 4, 5}}, m.NewIntFromInt(5)
	interpreter.backupInput(argB, argLen, padChar)
	bs, err = bytesHooks.padLeft(argB, argLen, padChar, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{1, 2, 3, 4, 5}, bs, err, interpreter)
	interpreter.checkImmutable(t, argB, argLen, padChar)

	argB, argLen = &m.Bytes{Value: []byte{1, 2, 3, 4, 5}}, m.NewIntFromInt(7)
	interpreter.backupInput(argB, argLen, padChar)
	bs, err = bytesHooks.padLeft(argB, argLen, padChar, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{80, 80, 1, 2, 3, 4, 5}, bs, err, interpreter)
	interpreter.checkImmutable(t, argB, argLen, padChar)

	argB, argLen = m.BytesEmpty, m.NewIntFromInt(3)
	interpreter.backupInput(argB, argLen, padChar)
	bs, err = bytesHooks.padLeft(argB, argLen, padChar, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{80, 80, 80}, bs, err, interpreter)
	interpreter.checkImmutable(t, argB, argLen, padChar)
}

func TestBytesReverse(t *testing.T) {
	interpreter := newTestInterpreter()
	var bs, arg m.K
	var err error

	arg = m.BytesEmpty
	interpreter.backupInput(arg)
	bs, err = bytesHooks.reverse(arg, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg)

	arg = &m.Bytes{Value: []byte{1, 2, 3, 4, 5}}
	interpreter.backupInput(arg)
	bs, err = bytesHooks.reverse(arg, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{5, 4, 3, 2, 1}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg)
}

func TestBytesConcat(t *testing.T) {
	interpreter := newTestInterpreter()
	var bs, arg1, arg2 m.K
	var err error

	arg1, arg2 = m.BytesEmpty, m.BytesEmpty
	interpreter.backupInput(arg1, arg2)
	bs, err = bytesHooks.concat(arg1, arg2, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2)

	arg1, arg2 = &m.Bytes{Value: []byte{1, 2, 3}}, &m.Bytes{Value: []byte{4, 5}}
	interpreter.backupInput(arg1, arg2)
	bs, err = bytesHooks.concat(arg1, arg2, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{1, 2, 3, 4, 5}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2)

	arg1, arg2 = &m.Bytes{Value: []byte{1, 2, 3}}, m.BytesEmpty
	interpreter.backupInput(arg1, arg2)
	bs, err = bytesHooks.concat(arg1, arg2, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{1, 2, 3}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2)

	arg1, arg2 = m.BytesEmpty, &m.Bytes{Value: []byte{1, 2, 3}}
	interpreter.backupInput(arg1, arg2)
	bs, err = bytesHooks.concat(arg1, arg2, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBytesOk(t, []byte{1, 2, 3}, bs, err, interpreter)
	interpreter.checkImmutable(t, arg1, arg2)
}

func assertBytesOk(t *testing.T, expectedBytes []byte, actual m.K, err error, interpreter *Interpreter) {
	if err != nil {
		t.Error(err, interpreter)
	}
	expected := &m.Bytes{Value: expectedBytes}
	if !interpreter.Model.Equals(expected, actual) {
		t.Errorf("Unexpected Bytes. Got: %s Want: %s.",
			interpreter.Model.PrettyPrint(actual),
			interpreter.Model.PrettyPrint(expected))
	}
}
