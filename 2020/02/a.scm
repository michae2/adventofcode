#!/usr/bin/env scheme-script
#!r6rs

(import (rnrs base)
        (rnrs control)
        (rnrs lists)
        (rnrs io simple)
        (rnrs io ports)
        (rnrs unicode))

(define (read-integer)
  (do ([digits '() (cons digit digits)]
       [digit (peek-char) (peek-char)])
      ((or (eof-object? digit) (not (char-numeric? digit)))
       (if (null? digits)
           0
           (string->number (list->string (reverse digits)))))
    (read-char)))

(let each-pw ([valid 0])
  (if (port-eof? (current-input-port))
      (begin
        (write valid)
        (newline))
      (let* ([lo (read-integer)]
             [sep (read-char)]
             [hi (read-integer)]
             [ws (read-char)]
             [ch (read-char)]
             [sep (read-char)]
             [ws (read-char)]
             [pw (get-line (current-input-port))]
             [n (fold-left
                  (lambda (n pwc)
                    (if (char=? pwc ch)
                        (+ n 1)
                        n))
                  0
                  (string->list pw))])
        (each-pw (if (and (<= lo n) (<= n hi))
                     (+ valid 1)
                     valid)))))
