(import (rnrs))

(define (square n)
  (cond
    ((<= n 0) (error 'ValueError "Cant Use Zero"))
    ((> n 64) (error 'ValueError "Must be less than 64"))
    (else (ash 1 (- n 1)))))


(define total

  (let loop (
      (i 1)
      (acc 0))

    (if (> i 64)
        acc
        (loop
          (+ i 1)
          (+ acc (square i))))))
