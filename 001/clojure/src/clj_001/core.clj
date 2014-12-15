; This solution is pretty literal... there are probably better ways
(ns clj-001.core)

(defn answer []
  (def limit 1000)
  (def sum3 (reduce + (range 0 limit 3)))
  (def sum5 (reduce + (range 0 limit 5)))
  (def sum15 (reduce + (range 0 limit 15)))
  (- (+ sum3 sum5) sum15)
  )

(defn answer2 []
  (def limit 1000)
  (+
   (reduce + (range 0 limit 3))
   (reduce + (range 0 limit 5))
   (- (reduce + (range 0 limit 15)))
   )
  )


(defn -main [& args]
  (println (answer))
  (println (answer2))
  )
