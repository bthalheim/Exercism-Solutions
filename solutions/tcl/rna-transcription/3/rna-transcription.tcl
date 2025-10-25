
proc to-rna {dna} {

    set result "" 
    for {set i 0} {$i < [string length $dna]} {incr i} {
        switch [string index $dna $i] {
            G { append result C}
            C { append result G}
            T { append result A}
            A { append result U}
            default { error "Input Error"}
        }
    }
    return $result 

}



