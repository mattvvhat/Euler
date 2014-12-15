(ns euler-003.core)
(defn largest-prime-factor [n]
  (loop [x n d 2]
    (if (= x 1)
      ; Return Value
      d

      ; Recursion: This must be at the end for tail recursion to work
      (if (zero? (mod x d))
        (recur (/ x d) d)
        (recur x (inc d)))
      ))
  )


(defn -main [& args]
  (def VAL 600851475143)

  (println "Starting...")
  (println (largest-prime-factor VAL))
  )
