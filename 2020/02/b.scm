#!/usr/bin/env scheme-script
#!r6rs

(import (rnrs base)
        (rnrs control)
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
      (let* ([a (read-integer)]
             [sep (read-char)]
             [b (read-integer)]
             [ws (read-char)]
             [ch (read-char)]
             [sep (read-char)]
             [ws (read-char)]
             [pw (get-line (current-input-port))])
        (each-pw (if (not (boolean=? (and (<= a (string-length pw))
                                          (char=? (string-ref pw (- a 1)) ch))
                                     (and (<= b (string-length pw))
                                          (char=? (string-ref pw (- b 1)) ch))))
                     (+ valid 1)
                     valid)))))
