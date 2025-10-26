
proc findAnagrams {subject candidates} {

    set maskPattern  [lsort [split [string tolower $subject] ""]]
    set ret {} 

    foreach word $candidates {

        # Word doesn't match itself
        if { [string tolower $word] == [string tolower $subject]} { continue }

        if { $maskPattern == [lsort [split [string tolower $word] ""]] } {
            lappend ret $word
        }        
    }

    return $ret

}
