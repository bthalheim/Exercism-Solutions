#!/usr/bin/env tclsh

proc isLeapYear {year} {

    if {$year % 100 == 0} {
        return [expr {$year % 400 == 0}]
    } else {
        return [expr {$year % 4 == 0}]
    }

}
