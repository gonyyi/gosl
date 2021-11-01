# Go Small Library (gosl)

Copyright Gon Y. Yi 2021 <https://gonyyi.com/copyright>


## Goal

__General__

- [ ] No import of any library whatsoever including standard library.
- [ ] All the code should have zero memory allocation.
- [ ] 99%+ code coverage
- [ ] All the code should have tests, benchmarks and examples
- [ ] Minimize allocation caused by importing the library
    - Currently `bufp` is allocated at global level when importing the library
      as this is required for the logger. 

__Functionality__

- [ ] StringJoin
- [ ] FilePath
    - Append path
    - Remove last path (eg. `/abc/def/ghi` --> `/abc/def`)

