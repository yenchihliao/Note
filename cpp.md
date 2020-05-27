* `Variadic template` and `variadic variable` [guide](https://kevinushey.github.io/blog/2016/01/27/introduction-to-c++-variadic-templates/)(...)
* `cv-qulification`, `ref-qualification` [explaination](https://stackoverflow.com/questions/8610571/what-is-rvalue-reference-for-this). Consider `const mem_func() const && {}`:
  * returns const value(first const)
  * modifies no attribute of the class(second const)
  * &&/& implies the function only works when it's a rvalue/lvalue
* Smart pointers such as `unique_ptr`, `shared_ptr`, and `weak_ptr` is the pointers in c++ that has the ability to do garbage collection.
  * weak_ptr: a pointer that points to an address without affecting its lifetime.
* The situations that you might want to use `forward declarations` [here](https://stackoverflow.com/questions/8526819/c-header-files-including-each-other-mutually)
