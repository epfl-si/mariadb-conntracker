# CHANGELOG

## v1.1.9 (2025-09-04)
- Bump modernc.org/sqlite from 1.38.2 to 1.40.0

## v1.1.8 (2025-08-25)
- Bump modernc.org/sqlite from 1.38.1 to 1.38.2

## v1.1.7 (2025-08-04)
- Bump modernc.org/sqlite from 1.38.0 to 1.38.1

## v1.1.6 (2025-06-16)
- Bump modernc.org/sqlite from 1.37.1 to 1.38.0

## v1.1.5 (2025-05-30)
- Bump modernc.org/sqlite from 1.37.0 to 1.37.1

## v1.1.4 (2025-04-03)
- Bump modernc.org/sqlite from 1.36.1 to 1.37.0
- Bump Go from 1.22 to 1.24

## v1.1.3 (2025-03-17)
- Bump modernc.org/sqlite from 1.34.5 to 1.36.1

## v1.1.2 (2025-03-14)
- Bump modernc.org/sqlite from 1.33.1 to 1.34.5

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
