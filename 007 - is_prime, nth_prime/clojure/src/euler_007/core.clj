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
  ; Naive-naive bad way
  (time 
    (println (filter is-prime? (range 2 1000))))

  ; Sieve
  (time 
    (println
      (sieve-of-eratosthenes 1000)))
  )
