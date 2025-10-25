#!/usr/bin/env tclsh

proc isLeapYear {y} {
    expr $y % 400 == 0 | $y % 100 != 0 && $y % 4 == 0
}
