// package appdir provides a facility for determining the system-dependent
// paths for application resources.
package appdir

// General returns the path for general aplication resources (e.g.
// ~/Library/<App>).
func General(app string) string {
	return general(app)
}

// Logs returns the path for log files (e.g. ~/Library/Logs/<App>).
func Logs(app string) string {
	return logs(app)
}
