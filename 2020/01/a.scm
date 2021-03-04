#!/usr/bin/env scheme-script
#!r6rs

(import (rnrs base)
        (rnrs control)
        (rnrs hashtables)
        (rnrs io simple))

(call/cc
  (lambda (break)
    (do ([e (read) (read)]
         [es (make-eqv-hashtable)])
        ((eof-object? e))
      (let ([d (- 2020 e)])
        (when (hashtable-contains? es d)
          (write (* d e))
          (newline)
          (break)))
      (hashtable-set! es e #t))))
