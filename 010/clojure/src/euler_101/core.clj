(ns euler-101.core)

(defn naive-primes
   "Implementation of lazy but naive sieve algorithms"
   []
  (letfn [
          (next-prime [x xs] 
             (if (some #(zero? (rem x %))
                    (take-while #(<= (* % %) x) xs))
             (recur (+ x 2) xs)
             (cons x (lazy-seq (next-prime (+ x 2) (conj xs x))))))]
   (cons 2 (lazy-seq (next-prime 3 [])))))

(defn -main [& args]
  (time (println
    (apply + (take-while (partial > 2000000) (naive-primes))))))
