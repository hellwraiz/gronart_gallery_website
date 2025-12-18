# gronart-galery-website
Hi! This is a website that I made for my mother, who has her own gallery. Enjoy!

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
