#!/usr/bin/env tclsh

proc isLeapYear {year} {

    if {[expr $year % 100] == 0} {
        if {[expr $year % 400] == 0} {
            return true
        }

        return false
        
    }

    if {[expr $year % 4] == 0} {
        return true
    }

    return false

}
