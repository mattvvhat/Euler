(ns euler-006.core)

; Square
(defn sqr [x] (* x x))

(defn -main [& args]
  (println
    (-
     (sqr (/ (* 100 101) 2))
     ; OR: 
     ; (sqr (reduce + (range 1 101)))
     (reduce + (map sqr (range 1 101)))
     ))
    )
