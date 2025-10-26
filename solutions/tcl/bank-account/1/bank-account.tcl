oo::class create BankAccount {

    variable open
    variable balance

    constructor {} {
        set open false
        set balance 0
    }

    method open {} {
        variable open
        
        if {$open} { error "account is already open" } 
        set open true
    }

    method close {} {
        variable open
        variable balance
        
        if {!$open} {error "account is not open"}
        set balance 0
        set open false
    }

    method balance {} {
        variable balance
        variable open

        if {!$open} { error "account is not open" }
        return $balance
    }

    method deposit {amount} {
        variable balance
        variable open

        if {!$open} { error "account is not open" }
        if {$amount < 0 } {error "invalid amount" }

        incr balance $amount 

    }

    method withdraw {amount} {
        variable balance
        variable open

        if {!$open} { error "account is not open" }
        if {$amount < 0 } {error "invalid amount" }
        if {$amount > $balance } { error "insufficient funds" }

        incr balance [expr {-1 * $amount}]

    }



}
