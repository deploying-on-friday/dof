# Problems in whoknows legacy codebase

## Critical priority issues

### 1. Database stored in /tmp directory
**Location:** `app.py` line 14

**Problem:** Database path is set to `/tmp/whoknows.db` which is deleted on server restart

**Impact:** All user data and content will be lost on restart

**Solution:** Use persistent storage location outside /tmp



### 2. SQL Injection vulnerability

**Location:** `app.py` lines 64, 80, 102, 142, 151, 181-182

**Problem:** User input is directly concatenated into SQL queries using string formatting

**Impact:** Attackers can execute arbitrary SQL commands, steal data, or delete entire database

**Solution:** Use parameterized queries with placeholders instead of string concatenation

**Example:** 
```python
#bad:
"select id from users where username = '%s'" % username

#good:
"select id from users where username = ?" (username)
```


### 3. Hardcoded SECRET_KEY

**Location:** `app.py` line 17

**Problem:** SECRET_KEY is hardcoded as 'development key' in source code

**Impact:** Anyone with access to code can forge sessions and impersonate users

**Solution:** Use environment variables to store secrets


### 4. MD5 hash

**Location:** `app.py` line 198

**Problem:** MD5 is unsafe


### 5. CSRF protection

**Location:** `app.py` line 147-160

**Problem:** Endpoints have no CSRF token validation


### 6. Plaintext error messages expose sys info

**Location:** `app.py` line 39

**Problem:** print ("Database not found") leaks internal system info to logs/console

**Impact:** Attackers learn about database file paths and system configuration



### 7. Missing input validation on user registration

**Location:** `app.py` line 163

**Problem:** Users can register with invalid emails or weak passwords



### 8. Session configuration missing security flags

**Location:** `app.py` 

**Problem:** Needs secure cookie flags for session cookies



### 9. No rate limiting on authenication endpoints

**Location:** `app.py` line 147 api_login(): and line 163 api_register():

**Problem:** No protection against brute force attacks. Unlimited login/registration attempts


### 10. No password salt

**Location:** `app.py` line 198 hash_password():

**Problem:** Passwords are hashed without adding salt


###  11. Missing form field validation (causes crashes)

**Location:** `app.py` line 147 api_login() and line 163 api_register()

**Problem:** `request.form['username']` and `request.form['password']` throw KeyError if fields are missing in POST request. Application crashes if attacker sends malformed POST requests without required form fields.


## End of critical issues




## High priority issues


### 1. Application crashes instead of handling error

**Location:** `app.py` line 35 check_db_exists()

**Problem:** `sys.exit(1)` stops application when database is not found


### 2. No logging or monitoring

**Location:** `app.py`

**Problem:** No logging of errors, security events, user actions. Cannot detect attacks, debug issues or audit user activities


### 3. Missing error handling on database operations

**Location:** `app.py` all query_db calls (lines 54, 80, 102, 142, 151, 180)

**Problem:** No try-catch blocks around database operations so when errors happen it crashes instead of showing user-friendly error messages


### 4. Debug mode configuration

**Location:** `app.py` line 16

**Problem:** Debug flag defined but set to false without environment variable option makes it harder to toggle debug mode for different environments


## End of high priority issues




## Medium priority issues


### 1. No database connection pooling

**Location:** `app.py` line 72 before_request(), 83 after_request()

**Problem:** New database connection opened and closed on every single request. Poor performance under high load


### 2. Logout route uses GET instead of POST

**Location:** `app.py` line 187 logout() uses GET method instead of POST

**Problem:** Users can be logged out by visiting malicious links


### 3. No HTTP method restrictions on page routes

**Location:** `app.py` line 94 search(), line 107 about(), line 113 login(), line 121 register()

**Problem:** Page routes accept all HTTP methods, not restricted by GET only


### 4. Missing validation on language parameter

**Location:** `app.py` lines 98, 138

**Problem:** Language parameter not validated against allowed values ('en', 'da'). Invalid language values could cause SQL errors or bypass language filter


### 5. Hardcoded host & port

**Location:** `app.py` line 214

**Problem:** Server runs on hardcoded port. Cannot easily change, needs to be in environment variables
























# Template


### 

**Location:** 

**Problem:** 

