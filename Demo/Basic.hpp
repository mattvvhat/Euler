/**
 * C++ Class Illustrating Rule of 3 (+2)
 *
 */

#pragma once

#include <iostream>

template <typename type>
class Basic {
public:

  Basic (type *val=NULL) : ptr(val) { }

  // Rule #1 (Copy Constructor)
  Basic (const Basic<type> &copy) {
    std::cout << "(1) Copy Constructor\n";
    ptr = copy.ptr != NULL ? new type(*copy.ptr) : NULL;
  }

  // Rule #4 (Move Constructor)
  Basic (Basic<type> &&orig) : ptr(NULL) {
    std::cout << "(4) Move Constructor\n";
    *this = std::move(orig);
  }

  // Rule #2 (Destructor)
  ~Basic () {
    if (ptr != NULL) {
      delete ptr;
    }
    std::cout << "(3) Destructor\n";
  }

  // Rule #3 (Copy Assignment)
  Basic& operator= (Basic &rhs) {
    std::cout << "(2) Copy Assignment\n";
    ptr = rhs.ptr ? new type(*rhs.ptr) : NULL;
    return *this;
  }

  // Rule #5 (Move Assignment)
  Basic& operator= (Basic &&rhs) {
    std::cout << "(5) Move Assignment\n";
    if (this == &rhs)
      return *this;

    if (ptr)
      delete ptr;

    ptr = rhs.ptr;
    rhs.ptr = NULL;

    return *this;
  }

  friend std::ostream& operator<< (std::ostream &os, const Basic &rhs) {
    os << *rhs.ptr;
    return os;
  }

protected:
  type *ptr;
};

