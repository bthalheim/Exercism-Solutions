namespace eval resistorColor {

    proc colorCode {args} {
         
        set bands [dict create \
            black   0 white   9  orange  3 brown   1 \
            red     2 yellow  4  green   5 blue    6 \
            violet  7 grey    8]


        if {![dict exists $bands $args ]} { error "Invalid color: $args"}

        return [dict get $bands $args]
        
    }

    proc colors {args} {
        return {black brown red orange yellow green blue violet grey white}
    }

}
