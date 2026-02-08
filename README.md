# gronart-galery-website
Hi! This is a website that I made for my mother, who has her own gallery. Enjoy!

# TODOs

1. Cleanup price still showing when painting is indicated as sold
2. 404 page

# Project structure
More so for me than anyone. But if you're curious about the login behind the code then here's the place!

## Principles
Type = "what purpose does it serve?" = models, handlers, repository, helpers
Domain = "what object is it?" = paintings, media, routes
1. Separate with folders by domain, not type
2. Have one init folder that initializes all domains, separating using files by type.
3. In the end call the files in the init folder in main.go

### Default files inside each domain's folder (with exceptions ofc)
1. `models.go` -> stores database schemas, and all sorts of struct types that you use
    - Use pointers for optional fields (put/patch requests)
    - Use regular types for read fields (read requests)
    - Use regular + `binding: required` for required fields (post requests)
2. `handlers.go` -> stores all http request handlers.
    - Handles authentication, authorization and input validation
3. `repository.go` -> stores all database crud operations related stuff
    - Only concerned about performing the database query. Assumes that all the input is already valid.
        - MAKE SURE TO DO VALIDATION IN `handlers.go` THEN!!!
4. `helpers.go` -> stores all functions that aren't directly involved in the way the file performs, but still valuable.


# Future things

## Security stuff

Please just ignore this!!!! This is just for me for now for later yknow

Ways to protect the upload endpoint:
1. Authentication/Authorization

Most important: Only allow logged-in admins to upload
Use JWT tokens, session cookies, or API keys
Check Authorization header before allowing upload

2. Rate Limiting

Limit uploads per IP (e.g., 10 per hour)
Libraries like github.com/ulule/limiter can help
Prevents spam and DoS

3. File Validation

Check file type (only allow .jpg, .png, etc.)
Check file size (max 5MB or whatever)
Validate it's actually an image (not a disguised executable)
Use http.DetectContentType() to verify

4. Orphaned File Cleanup

If someone uploads but never creates a painting, you have a dead file
Run a cleanup job that deletes uploads not referenced in the database
Or: generate a temporary token that expires

5. Virus Scanning

For production: scan uploaded files with antivirus
Services like ClamAV or cloud scanning APIs


For GET request DoS protection:
1. Rate Limiting (again)

Limit requests per IP
Different limits for different endpoints

2. Caching

Cache GET responses so database isn't hit constantly
Use Redis or in-memory cache

3. Pagination

Never allow unlimited results
Force a max limit (e.g., LIMIT 100)

4. Request Timeouts

Set max query execution time
Prevent slow queries from hogging resources

5. Use a reverse proxy

Nginx, Cloudflare can block malicious traffic before it hits your app


For your project right now:
Minimum viable security:

Add authentication (most critical - we can discuss how)
Validate file uploads (type, size)
Add rate limiting to upload endpoint
Set max limit on GET queries (you already have this!)

### Authentication

Here's the plan, we're going to use session tokens while using cookies to store them. For now though we simplify.

## Other changes

### Cleanup

Since upload and painting stuff is separate make sure that painting creation and editing cleans up everything properly in case that there's a problem, so no hanging images are left!
