(ns euler-007.core)

; is this the best is prime that we can do?
(defn is-prime? [n]
  (loop [ d 2 ]
    (if (= d n)
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

  ; 
  (loop [ i 2 nums (range (inc n)) ]
    (if (> i n)
      nums
      (recur
        (first nums)
        (cons i (filter (fn [x] (not (zero? (mod x i)))) (rest nums)))
        )
      )
    )
  )

(defn -main [& args]
  (println (sieve-of-eratosthenes 10))
  )
