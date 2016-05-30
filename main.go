
//line main.rl:1
// Ensure Ragel is installed (used v6.9), then run:
//   ragel -Z -T0 -o main.go main.rl
//   go build -o parser main.go 
//   ./parser 

package main

import (
    "fmt"
    "strconv"
    "time"
)


//line main.go:18
var _pricedb_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 1, 5, 
	2, 0, 5, 2, 3, 4, 
}

var _pricedb_key_offsets []byte = []byte{
	0, 0, 1, 3, 7, 9, 11, 13, 
	14, 16, 18, 19, 21, 23, 25, 28, 
	31, 34, 
}

var _pricedb_trans_keys []byte = []byte{
	80, 9, 32, 9, 32, 48, 57, 48, 
	57, 48, 57, 48, 57, 45, 48, 57, 
	48, 57, 45, 48, 57, 48, 57, 9, 
	32, 9, 32, 34, 10, 13, 34, 10, 
	13, 34, 
}

var _pricedb_single_lengths []byte = []byte{
	0, 1, 2, 2, 0, 0, 0, 1, 
	0, 0, 1, 0, 0, 2, 3, 3, 
	3, 0, 
}

var _pricedb_range_lengths []byte = []byte{
	0, 0, 0, 1, 1, 1, 1, 0, 
	1, 1, 0, 1, 1, 0, 0, 0, 
	0, 0, 
}

var _pricedb_index_offsets []byte = []byte{
	0, 0, 2, 5, 9, 11, 13, 15, 
	17, 19, 21, 23, 25, 27, 30, 34, 
	38, 42, 
}

var _pricedb_indicies []byte = []byte{
	0, 1, 2, 2, 1, 2, 2, 3, 
	1, 4, 1, 5, 1, 6, 1, 7, 
	1, 8, 1, 9, 1, 10, 1, 11, 
	1, 12, 1, 13, 13, 1, 14, 14, 
	15, 1, 1, 1, 17, 16, 1, 1, 
	19, 18, 20, 
}

var _pricedb_trans_targs []byte = []byte{
	2, 0, 3, 4, 5, 6, 7, 8, 
	9, 10, 11, 12, 13, 14, 14, 15, 
	16, 17, 16, 17, 17, 
}

var _pricedb_trans_actions []byte = []byte{
	0, 0, 0, 1, 0, 0, 0, 3, 
	1, 0, 5, 1, 0, 12, 0, 0, 
	1, 9, 0, 7, 0, 
}

const pricedb_start int = 1
const pricedb_first_final int = 17
const pricedb_error int = 0

const pricedb_en_main int = 1


//line main.rl:17


type Symbol struct {
    Symbol string
    Quoted bool
}

func pricedb(data string) (time.Time, Symbol) {
    cs, p, pe := 0, 0, len(data)
    mark := 0
    var year, month, day int
    var date time.Time
    var symbol Symbol
    
    
//line main.go:100
	{
	cs = pricedb_start
	}

//line main.go:105
	{
	var _klen int
	var _trans int
	var _acts int
	var _nacts uint
	var _keys int
	if p == pe {
		goto _test_eof
	}
	if cs == 0 {
		goto _out
	}
_resume:
	_keys = int(_pricedb_key_offsets[cs])
	_trans = int(_pricedb_index_offsets[cs])

	_klen = int(_pricedb_single_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + _klen - 1)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + ((_upper - _lower) >> 1)
			switch {
			case data[p] < _pricedb_trans_keys[_mid]:
				_upper = _mid - 1
			case data[p] > _pricedb_trans_keys[_mid]:
				_lower = _mid + 1
			default:
				_trans += int(_mid - int(_keys))
				goto _match
			}
		}
		_keys += _klen
		_trans += _klen
	}

	_klen = int(_pricedb_range_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + (_klen << 1) - 2)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + (((_upper - _lower) >> 1) & ^1)
			switch {
			case data[p] < _pricedb_trans_keys[_mid]:
				_upper = _mid - 2
			case data[p] > _pricedb_trans_keys[_mid + 1]:
				_lower = _mid + 2
			default:
				_trans += int((_mid - int(_keys)) >> 1)
				goto _match
			}
		}
		_trans += _klen
	}

_match:
	_trans = int(_pricedb_indicies[_trans])
	cs = int(_pricedb_trans_targs[_trans])

	if _pricedb_trans_actions[_trans] == 0 {
		goto _again
	}

	_acts = int(_pricedb_trans_actions[_trans])
	_nacts = uint(_pricedb_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _pricedb_actions[_acts-1] {
		case 0:
//line main.rl:32

            fmt.Println("Mark")
            mark = p
        
		case 1:
//line main.rl:36
 year, _ = strconv.Atoi(data[mark:p]) 
		case 2:
//line main.rl:37
 month, _ = strconv.Atoi(data[mark:p]) 
		case 3:
//line main.rl:38
 day, _ = strconv.Atoi(data[mark:p]) 
		case 4:
//line main.rl:39
 date = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local) 
		case 5:
//line main.rl:40

            symbol = Symbol {
                Symbol: data[mark:p],
                Quoted: true,
            }
        
//line main.go:210
		}
	}

_again:
	if cs == 0 {
		goto _out
	}
	p++
	if p != pe {
		goto _resume
	}
	_test_eof: {}
	_out: {}
	}

//line main.rl:65

    
    return date, symbol
}

func main() {
    date, symbol := pricedb("P 2016-05-26 \"TDB911\" $11.96")
    fmt.Println(date)
    fmt.Println(symbol)
}