Small utility to check a list of paths and return the last time they were
modified. Can also pass an individual path as an argument from the command line.

I built this as I need to check the last time a big number of backup files were
created/modified. It comes useful when you need to check if the daily backups
of 100s of servers have been completed.

Todo:
-----
1. Compare file size on disk with expected file size
2. Compare modify date with current date and report out of date files
   - i.e. in the case of checking backups, this can indicate which backups have 
	 failed or haven't finished yet