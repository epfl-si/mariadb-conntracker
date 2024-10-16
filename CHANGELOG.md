# CHANGELOG

## v1.1.1 (2024-10-16)

### Summary
Fix a bug introduced in v1.1.0 that prevent the script to run a second time.

### Fixes
- Fix date format from SQLite when retrieving the last processing date from the mysql@localhost account.

## v1.1.0 (2024-10-15)

### Summary
This release addresses inconsistencies in date handling throughout the program. We've standardized the date format across all components, including filesystem interactions (file modification dates), SQLite database storage, and audit file parsing.

### Changes
- Standardized date format to include time location across all program components
- Store the time zone indicator at the end of date in SQLite (E.G.: +02:00 for GMT+2)

## v1.0.0 (2024-10-14)

### Summary
Initial release
