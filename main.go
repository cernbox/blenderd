package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

var port int

func main() {
	flag.IntVar(&port, "p", 6666, "listening port")
	flag.Parse()

	http.HandleFunc("/rooteosuser", rootEOSUser)
	http.HandleFunc("/rooteosproject", rootEOSProject)
	addr := fmt.Sprintf("%s:%d", "0.0.0.0", port)
	http.ListenAndServe(addr, nil)
}

func rootEOSUser(w http.ResponseWriter, r *http.Request) {
	write(w, r, userData)
}

func rootEOSProject(w http.ResponseWriter, r *http.Request) {
	write(w, r, projectData)
}

func write(w http.ResponseWriter, r *http.Request, data string) {
	if r.Method == "PROPFIND" {
		w.Header().Set("Content-Type", "application/xml; charset=utf-8")
		w.Header().Set("Content-Length", fmt.Sprintf("%d", len(data)))
		w.WriteHeader(207)
		if _, err := w.Write([]byte(data)); err != nil {
			fmt.Fprintf(os.Stderr, "error writing propfind data: %+v", err)
			return
		}
	} else {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}
}

var userData string = `<?xml version="1.0" encoding="utf-8"?>
<d:multistatus xmlns:d="DAV:">
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/user</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/user/a</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/user/a</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/user/b</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/user/c</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/user/d</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/user/e</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/user/f</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/user/g</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/user/h</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/user/i</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/user/j</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/user/k</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/user/l</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/user/m</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/user/n</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/user/o</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/user/p</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/user/q</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/user/r</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/user/s</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/user/t</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/user/u</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/user/v</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/user/w</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/user/x</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/user/y</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/user/z</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
</d:multistatus>
`

var projectData string = `<?xml version="1.0" encoding="utf-8"?>
<d:multistatus xmlns:d="DAV:">
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/project</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/project/a</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/project/a</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/project/b</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/project/c</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/project/d</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/project/e</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/project/f</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/project/g</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/project/h</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/project/i</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/project/j</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/project/k</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/project/l</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/project/m</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/project/n</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/project/o</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/project/p</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/project/q</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/project/r</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/project/s</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/project/t</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/project/u</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/project/v</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/project/w</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/project/x</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/project/y</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
<d:response><d:href>/cernbox/desktop/remote.php/webdav/eos/project/z</d:href>
<d:propstat><d:status>HTTP/1.1 200 OK</d:status><d:prop><d:resourcetype><d:collection/></d:resourcetype></d:prop></d:propstat><d:propstat><d:status>HTTP/1.1 404 Not Found</d:status><d:prop/></d:propstat></d:response>
</d:multistatus>
`
