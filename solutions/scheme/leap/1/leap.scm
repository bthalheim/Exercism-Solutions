(import (rnrs))

(define (leap-year? x)
  (or
    (and
       (zero? (modulo x 4))
       (not (zero? (modulo x 100))))
    (zero? (modulo x 400)) ))

