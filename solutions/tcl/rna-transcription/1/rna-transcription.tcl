proc to-rna {dna} {

    if {$dna == ""} { return $dna}
    set newchar {}

    for {set i 0} {$i < [string length $dna]} {incr i} {

        set char [string index $dna $i]

        switch $char {
            G { lappend newchar C}
            C { lappend newchar G}
            T { lappend newchar A}
            A { lappend newchar U}
        }
    }

    set result [join $newchar ""]
    return $result 

}
