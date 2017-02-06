package command

const (
	/* Closes the active ro file (if any) and open/stat a new one */
	OpenFileCode = 0x1224 + iota
	/* Reads the active ro file. Offsets and sizes in bytes. If file read fails, client is exited. Only read data is returned. */
	ReadFileCriticalCode
	/* Reads 2048 sectors in a 2352 sectors iso.
	 * Offsets and sizes in sectors. If file read fails, client is exited */
	ReadCD2048CriticalCode
	/* Reads the active ro file. Offsets and sizes in bytes. It returns number of bytes read to client, -1 on error, and after that, the data read (if any)
	   Only up to the BUFFER_SIZE used by server can be red at one time*/
	ReadFileCode
	/* Closes the active wo file (if any) and opens+truncates or creates a new one */
	CreateFileCode
	/* Writes to the active wo file. After command, data is sent. It returns number of bytes written to client, -1 on erro.
	   If more than BUFFER_SIZE used by server is specified in command, connection is aborted. */
	WriteFileCode
	/* Closes the active directory (if any) and opens a new one */
	OpenDirCode
	/* Reads a directory entry. and returns result. If no more entries or an error happens, the directory is automatically closed. . and .. are automatically ignored */
	ReadDirEntryCode
	/* Deletes a file. */
	DeleteFileCode
	/* Creates a directory */
	MakeDirCode
	/* Removes a directory (if empty) */
	RemoveDirCode
	/* Reads a directory entry (v2). and returns result. If no more entries or an error happens, the directory is automatically closed. . and .. are automatically ignored */
	ReadDirEntryV2Code
	/* Stats a file or directory */
	StatFileCode
	/* Gets a directory size */
	GetDirSizeCode

	/* Get complete directory contents */
	ReadDirCode

	/* Replace this with any custom command */
	Custom0Code = 0x2412
)
