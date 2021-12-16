package day16

import "math/big"

func Decode(data *big.Int, offset uint) (Packet, uint) {
	var packet Packet
	version, offset := readBits(data, offset, 3)
	typeID, offset := readBits(data, offset, 3)
	switch t := typeID.Uint64(); t {
	case 4:
		value := new(big.Int)
		hasMoreGroups := true
		for hasMoreGroups {
			var group *big.Int
			group, offset = readBits(data, offset, 5)
			hasMoreGroups = group.Bit(4) == 1
			group.SetBit(group, 4, 0)
			value.Lsh(value, 4).Or(value, group)
		}
		packet = literalValue{header{version.Uint64(), t}, value}
	default:
		var operands []Packet
		offset--
		lengthTypeID := data.Bit(int(offset))
		if lengthTypeID == 0 {
			// total length
			var length *big.Int
			length, offset = readBits(data, offset, 15)
			until := offset - uint(length.Uint64())
			for offset > until {
				var subPacket Packet
				subPacket, offset = Decode(data, offset)
				operands = append(operands, subPacket)
			}
		} else {
			//number of sub-packets
			var number *big.Int
			number, offset = readBits(data, offset, 11)
			count := int(number.Uint64())
			for i := 0; i < count; i++ {
				var subPacket Packet
				subPacket, offset = Decode(data, offset)
				operands = append(operands, subPacket)
			}
		}
		packet = operatorValue{header{version.Uint64(), t}, operands}
	}

	return packet, offset
}

func readBits(data *big.Int, offset uint, count uint) (*big.Int, uint) {
	mask := new(big.Int).Set(one)
	mask.Lsh(mask, count).Sub(mask, one)
	bits := new(big.Int).Rsh(data, offset-count)
	return bits.And(bits, mask), offset - count
}

type Packet interface {
	GetVersionSum() int
	Calculate() *big.Int
}

type header struct {
	version uint64
	typeID  uint64
}

func (h header) GetVersionSum() int {
	return int(h.version)
}

type literalValue struct {
	header
	value *big.Int
}

func (l literalValue) GetVersionSum() int {
	return l.header.GetVersionSum()
}

func (l literalValue) Calculate() *big.Int {
	return l.value
}

type operatorValue struct {
	header
	operands []Packet
}

func (o operatorValue) GetVersionSum() int {
	sum := o.header.GetVersionSum()
	for _, operand := range o.operands {
		sum += operand.GetVersionSum()
	}

	return sum
}

func (o operatorValue) Calculate() *big.Int {
	operation := operations[o.typeID]
	result := new(big.Int)
	if len(o.operands) > 0 {
		result.Set(o.operands[0].Calculate())
		for _, operand := range o.operands[1:] {
			result = operation(result, operand.Calculate())
		}
	}
	return result
}

var one = big.NewInt(1)

var operations = [8]func(*big.Int, *big.Int) *big.Int{
	func(a *big.Int, b *big.Int) *big.Int {
		return (&big.Int{}).Add(a, b)
	},
	func(a *big.Int, b *big.Int) *big.Int {
		return (&big.Int{}).Mul(a, b)
	},
	func(a *big.Int, b *big.Int) *big.Int {
		if a.Cmp(b) <= 0 {
			return a
		}
		return b
	},
	func(a *big.Int, b *big.Int) *big.Int {
		if a.Cmp(b) >= 0 {
			return a
		}
		return b
	},
	nil,
	func(a *big.Int, b *big.Int) *big.Int {
		result := &big.Int{}
		if a.Cmp(b) > 0 {
			result.Set(one)
		}
		return result
	},
	func(a *big.Int, b *big.Int) *big.Int {
		if a.Cmp(b) < 0 {
			return big.NewInt(1)
		}
		return &big.Int{}
	},
	func(a *big.Int, b *big.Int) *big.Int {
		if a.Cmp(b) == 0 {
			return big.NewInt(1)
		}
		return &big.Int{}
	},
}
