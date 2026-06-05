(import (rnrs))

(define (two-fer . maybe-name)

  (let ((x (if (null? maybe-name) "you" (car maybe-name)))) 

  (format #f "One for ~a, one for me." x) ))

