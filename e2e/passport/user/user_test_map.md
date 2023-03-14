# user test map

## Create user

*succes
- [x] user doesn't exist

*failed
- [x] user exist
  * invalid argument:
    * first_name
  - [x] empty 
    * last_name
  - [x] empty
    * sex
  - [x] empty
    * age
  - [x] empty
    * login
  - [x] empty
    * pass
  - [x] empty

## Get user

*succes
- [x] user  exist

*failed
- [x] site doesn't exist
    * invalid argument:
      * field name
    - [x] is bad
    - [x] is empty
  
## Delete user

*succes
- [x] user exists

*failed
- [x] user doesn't exist
    * invalid argument:
        * field name
    - [x] is bad
    - [x] is empty

## List site

*succes:
   - [x] user  exists, without pagination
    
  *pagination:
  - [x] with limit (page = 0, limit = 1)
  - [ ] with limit (page = 2, limit = 1)

## Change password request

*succes:
- [x] user  exists, correct password

*failed:
- [x] user  exists, wrong userid
- [x] user  exists, wrong old password
- [x] user  exists, empty new password


