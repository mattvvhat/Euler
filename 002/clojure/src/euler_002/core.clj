(ns euler-002.core)

(defn fib [n]
  ; Final results
  (def results [ 1 1 ])
  (while (< (last results) n)

    (def results (conj results (+
                                (last results)
                                (last (pop results)))))
    )

  results
  )


(defn -main [& args]
  (println (reduce + (filter even? (fib 4000000))))
  )
