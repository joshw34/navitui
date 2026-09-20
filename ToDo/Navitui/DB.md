- Provide a name for the db
	- Based on username/url (possibly hash)
	- Offer to delete on logout

- Implement cache expiry
	- Pull from server and replace db cache

- Setting:
	- Disable caching (always pull)
	- Disable expiry (only refresh manually)
	- Ephemeral cache (hold db in memory)