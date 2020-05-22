* Vaiadic template and variable [guide](https://kevinushey.github.io/blog/2016/01/27/introduction-to-c++-variadic-templates/)(...)
* cv-qulification, ref-qualification [explaination](https://stackoverflow.com/questions/8610571/what-is-rvalue-reference-for-this). Consider `const mem_func() const && {}`:
  * returns const value(first const)
  * modifies no attribute of the class(second const)
  * &&/& implies the function only works when it's a rvalue/lvalue
