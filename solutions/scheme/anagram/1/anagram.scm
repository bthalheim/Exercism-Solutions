(import (rnrs))

(define (anagram target words)


  (let loop ((w words)
             (acc '()))
    (if (null? w)
        (reverse acc)
        (loop
          (cdr w)

          (if
           (and
              (equal? (sort char<? (string->list (string-downcase target)))
                      (sort char<? (string->list (string-downcase (car w)))))
              (not (equal? (string-downcase target)
                           (string-downcase (car w)))))


              (cons (car w) acc)
              acc)))))









