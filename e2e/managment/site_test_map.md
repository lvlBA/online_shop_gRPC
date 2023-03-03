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
- [ ] site  exist

*failed
- [ ] site doesn't exist
    * invalid argument:
      * field name
    - [ ] is bad
    - [ ] is empty
  
## Delete site

*succes
- [ ] site exists

*failed
- [ ] site doesn't exist
    * invalid argument:
        * field name
    - [ ] is bad
    - [ ] is empty

## List site

*succes:
   - [ ] site  exists, without pagination
    
    *pagination:
    - [ ] with limit (page = 0, limit = 1)
    - [ ] with limit (page = 2, limit = 1)

