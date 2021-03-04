#!/usr/bin/env scheme-script
#!r6rs

(import (rnrs base)
        (rnrs control)
        (rnrs hashtables)
        (rnrs io simple))

(call/cc
  (lambda (break)
    (do ([e (read) (read)]
         [es '() (cons e es)]
         [tab (make-eqv-hashtable)])
        ((eof-object? e))
      (let ([d (- 2020 e)])
        (when (hashtable-contains? tab d)
          (write (* (hashtable-ref tab d 0) e))
          (newline)
          (break)))
      (for-each
        (lambda (e0)
          (hashtable-set! tab (+ e0 e) (* e0 e)))
        es))))
