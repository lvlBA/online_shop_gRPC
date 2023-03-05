# site test map

## Create site

*succes
- [x] site doesn't exist

*failed
- [x] site exist
  * invalid argument:
    * field name
  - [x] empty 

## Get site

*succes
- [x] site  exist

*failed
- [x] site doesn't exist
    * invalid argument:
      * field name
    - [x] is bad
    - [x] is empty
  
## Delete site

*succes
- [x] site exists

*failed
- [x] site doesn't exist
    * invalid argument:
        * field name
    - [x] is bad
    - [x] is empty

## List site

*succes:
   - [ ] site  exists, without pagination
    
  *pagination:
  - [ ] with limit (page = 0, limit = 1)
  - [ ] with limit (page = 2, limit = 1)

