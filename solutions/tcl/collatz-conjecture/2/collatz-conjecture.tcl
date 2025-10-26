

proc recurse {number} {

    if {$number == 1} { return 0 }

    if {($number % 2) == 0} {
        return [expr {1 + [tailcall recurse [expr {$number / 2}]]}]
    } else {
        return [expr {1 + [tailcall recurse [expr {$number * 3 + 1}]]}]
    }

}


proc steps {start} {

    if {$start < 1} {
        error "Only positive integers are allowed"
    } 

    return [recurse $start]

}





