# Problems in whoknows legacy codebase

## Critical prio

### 1. Database stored in /tmp directory
**Location:** `app.py` line 14

**Problem:** Database path is set to `/tmp/whoknows.db` which is deleted on server restart

**Impact:** All user data and content will be lost on restart

**Solution:** Use persistent storage location outside /tmp

**Example:** 
```python
DATABASE_PATH = os.getenv('DB_PATH', '/var/lib/whoknows/whoknows.db')
```


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
"select id from users where username = ?", (username,)
```

### 3. Hardcoded SECRET_KEY

**Location:** `app.py` line 17

**Problem:** SECRET_KEY is hardcoded as 'development key' in source code

**Impact:** Anyone with access to code can forge sessions and impersonate users

**Solution:** Use environment variables to store secrets

**Example:**
```python
SECRET_KEY = os.getenv('SECRET_KEY', 'fallback-for-dev-only')
```


### To be continued...