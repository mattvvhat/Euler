(ns euler-007.core)

; is this the best is prime that we can do?
(defn is-prime? [n]
  (loop [ d 2 ]
    (if (> (* d d) n)
      true
      (if (zero? (mod n d))
        false
        (recur (inc d))))
    )
  )


(defn sieve-of-erat [total]
  "
  Sieve of Eratosthenes
  Needs to be patched for performance...
  what is wrong with the inner loop? It does a better job
  of filtering the list than the other version, but it has *way*
  worse runtime.
  "
  ; Divides
  (defn divides? [value divider] (zero? (mod value divider)))

  ; Does divider divide value without being equal to it?
  ; (defn is-multiple? [value divider] (and (not= value divider) (divides? value divider)))
  (defn is-multiple? [value divider] (divides? value divider))

  ; Filter all the multiples of a particular number
  (defn filter-multiples [ values number ]
    (filter (fn [x] (not (is-multiple? x number))) values))

  ; Tail-recursive loop
  (loop [primes '() values (range 2 (inc total))]

     (println values)
    ; (if (empty? values)
    (if (empty? values)
      ; Return
      primes

      (let [ div (first values) ]
        (recur
          (cons div primes)
          (filter-multiples values div)))))


  ; END: sieve-of-erat
  )

; 
;
;
(defn sieve-of-eratosthenes [n]

  ; Divides
  (defn divides? [value divider] (zero? (mod value divider)))

  ; Does divider divide value without being equal to it?
  (defn is-multiple? [value divider] (and (not= value divider) (divides? value divider)))

  ; Filter all the multiples of a particular number
  (defn filter-multiples [ values number ]
    (filter (fn [x] (not (is-multiple? x number))) values))

  ; Actual recursion
  (loop [i 2 values (range 2 (inc n))]
    (if (< n (* i i))
      ; Return value
      values
      ; Tail recursion
      (recur (inc i) (filter-multiples values i)))))

(defn -main [& args]
  ; Sieve 2
  (time 
    (println
      (sieve-of-erat 10)))

  ; Sieve
;   (time 
;     (println
;       (nth (sieve-of-eratosthenes 200000) 1)))

  (time 
    (println
      (sieve-of-eratosthenes 10000)))
  )
