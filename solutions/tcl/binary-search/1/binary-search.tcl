proc binarySearch {haystack needle} {


    if {[llength $haystack] == 0} {
        return -1
    }

    set low 0
    set high [expr {[llength $haystack] - 1}]

    while {$low <= $high} {
        set mid [expr {($low + $high) / 2}]
        set midVal [lindex $haystack $mid]

        if {$midVal == $needle} {
            return $mid
        } elseif {$midVal < $needle} {
            set low [expr {$mid + 1}]
        } else {
            set high [expr {$mid - 1}]
        }
            
    }

    return -1

}
