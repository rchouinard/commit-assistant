diff --git a/response_writer.go b/response_writer.go
index a46d90e..ba791c7 100644
--- a/response_writer.go
+++ b/response_writer.go
@@ -1,6 +1,9 @@
 package middlewares

 import (
+       "bufio"
+       "fmt"
+       "net"
        "net/http"
 )

@@ -66,3 +69,19 @@ func (rw *responseWriter) Size() int {
 func (rw *responseWriter) Written() bool {
        return rw.status != 0
 }
+
+// Hijack implements the [http.Hijacker] interface.
+func (rw *responseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
+       if orw, ok := rw.ResponseWriter.(http.Hijacker); ok {
+               return orw.Hijack()
+       }
+
+       return nil, nil, fmt.Errorf("Hijacker interface not implemented")
+}
+
+// Flush implements the [http.Flusher] interface.
+func (rw *responseWriter) Flush() {
+       if orw, ok := rw.ResponseWriter.(http.Flusher); ok {
+               orw.Flush()
+       }
+}
