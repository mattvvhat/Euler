(ns euler-005.core)

; GCD
(defn gcd [a b]
  (if (zero? b)
    a
    (recur b (mod a b)))
  )

; LCM
(defn lcm [a b]
  (/ (* a b) (gcd a b)))

; CLOJURE
(defn -main [& args]
  (println "Starting")
  (println (reduce lcm (range 1 20)))
  )
