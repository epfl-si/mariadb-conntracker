# CHANGELOG

## v1.1.0 (2024-10-15)

### Summary
This release addresses inconsistencies in date handling throughout the program. We've standardized the date format across all components, including filesystem interactions (file modification dates), SQLite database storage, and audit file parsing.

### Changes
- Standardized date format to include time location across all program components
- Store the time zone indicator at end of date in SQLite (+02:00 CEST)

### Breaking Changes
- Modified `TimeFormatDB` to incorporate time location

### Action Required
- Update your `/etc/conntracker/conntracker.ini` file to use the new date format that includes time location. Search for `TimeFormatDB` in `conntracker.ini-dist`.

## v1.0.0 (2024-10-14)

### Summary
Initial release
