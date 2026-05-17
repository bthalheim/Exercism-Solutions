(import (rnrs))


(define (dna->rna s)
  (let loop (
      (chars (string->list s))
      (acc '()))
    
    (if (null? chars)
        (list->string (reverse acc))
        (loop
          (cdr chars)
          (cons
            (cond
              ((char=? (car chars) #\G) #\C)
              ((char=? (car chars) #\C) #\G)              
              ((char=? (car chars) #\T) #\A)
              ((char=? (car chars) #\A) #\U)
              (else (car cars))) acc))))) 
              

