
# This version specifically not using built in function
# - string reverse $input

proc reverse {input} {

    set ret {}

    for {set i [string length $input]} { $i >= 0} {incr i -1} {
        lappend ret [string index $input $i]
    }
     
    return [join $ret ""]

}
