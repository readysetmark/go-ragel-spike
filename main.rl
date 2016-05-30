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

%%{
    machine pricedb;
    write data;
}%%

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
    
    %%{
        action mark {
            fmt.Println("Mark")
            mark = p
        }
        action year { year, _ = strconv.Atoi(data[mark:p]) }
        action month { month, _ = strconv.Atoi(data[mark:p]) }
        action day { day, _ = strconv.Atoi(data[mark:p]) }
        action date { date = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local) }
        action quotedSymbol {
            symbol = Symbol {
                Symbol: data[mark:p],
                Quoted: true,
            }
        }
        
        mws = [ \t]+;  # mandatory whitespace
        
        year = digit{4} >mark %year;
        month = digit{2} >mark %month;
        day = digit{2} >mark %day;
        date = year '-' month '-' day %date;
        
        quoted_symbol = (any* -- [\"\r\n]) >mark %quotedSymbol;
        symbol = '\"' quoted_symbol '\"';
        
        price = 'P' mws date mws symbol;
        
        rest = any*;
        
        main := price rest;
        
        write init;
        write exec;
    }%%
    
    return date, symbol
}

func main() {
    date, symbol := pricedb("P 2016-05-26 \"TDB911\" $11.96")
    fmt.Println(date)
    fmt.Println(symbol)
}