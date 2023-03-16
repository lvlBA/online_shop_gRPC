# user test map

## Create user

*succes
- [x] resource doesn't exist

*failed
- [x] resource exist
  * invalid argument:
    * URN
  - [x] empty 

## Get resource

*succes
- [x] Resource  exist

*failed
- [x] resource doesn't exist
    * invalid argument:
      * field URN
    - [x] is bad
    - [x] is empty
  
## Delete resource

*succes
- [x] resource exists

*failed
- [x] resource doesn't exist
    * invalid argument:
        * field URN
    - [x] is bad
    - [x] is empty

## List site

*succes:
   - [x] resource  exists, without pagination
    
  *pagination:
  - [x] with limit (page = 0, limit = 1)
  - [ ] with limit (page = 2, limit = 1)



