(ns euler-004.core)

; convert int to string
;
;
(defn int-to-string [n]
  (format "%d" n))

; is-palindrome?
;
;
(defn is-palindrome? [string]
  (def front (first string))
  (def back (last string))

  (if (<= (count string) 1)
    ; Return True
    true

    (if (= front back)
      ; Is the subsection a palindrome?
      (is-palindrome? (subs string 1 (- (count string) 1)))

      ; Return false
      false)))


; main
;
;
(defn -main [& args]
  (println "Starting...")

  ; Print
  (println 
    ; the maximum
    (apply max
           ; filtered 
           (filter
             ; by is-paindrome(int-to-string(.))
             (comp is-palindrome? int-to-string)
             ; of the list a*b a, b in 100..1000
             (for [a (range 100 1000) b (range 100 1000)] (* a b)))))
  )
