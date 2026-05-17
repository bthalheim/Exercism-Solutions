(import (rnrs))

(define (hamming-calc a b)

  (let loop (
      (chars_a (string->list a))
      (chars_b (string->list b))
      (count 0))
    (if (null? chars_a)
        count
        (loop
          (cdr chars_a)
          (cdr chars_b)
          (if (not (char=? (car chars_a) (car chars_b)))
              (+ count 1)
              count)))))



(define (hamming-distance strand-a strand-b)

  ; Guards First
  (cond

    ((string=? strand-a strand-b) 0)
    
    ; ((and (= (string-length strand-a) 0)
    ;       (= (string-length strand-b) 0)) 0)

    ((not (=
          (string-length strand-a)
          (string-length strand-b)))
     (error 'Argument-error "Arguments need same length")) 

    (else (hamming-calc strand-a strand-b))



))



