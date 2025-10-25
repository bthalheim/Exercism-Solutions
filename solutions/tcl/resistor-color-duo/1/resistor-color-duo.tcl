namespace eval ::resistorColor {

    proc value {args} {

        lassign $args first second

        set num1 [expr [colorCode $first] * 10]
        set num2 [colorCode $second]

        return [expr $num1 + $num2]

    }

    proc colorCode {args} {
    
        set i [lsearch -exact [colors] $args]

        if {$i == -1} {
            error "Invalid color: $args"
        }

        return $i 
    }

    proc colors {args} {
        return [list black brown red orange yellow green blue violet grey white ] 
    }

}
