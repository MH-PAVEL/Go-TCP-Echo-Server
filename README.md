# Go TCP Echo Server

A simple TCP echo server written in Go that listens for incoming connections, reads data from clients, and sends the same data back.  
This repository also includes my personal learning notes on how TCP servers work in Go and the underlying networking concepts.

---

## Features

- Listens on a custom port provided via command-line arguments.
- Handles multiple clients concurrently using goroutines.
- Reads and writes raw TCP streams.
- Echoes received data back to the client.
- Uses `defer` for proper resource cleanup.

---

## How It Works (Go Code Overview)

1. **Listen for connections**

   ```go
   l, err := net.Listen("tcp4", ":"+port)
   ```

   - `tcp4` specifies IPv4 TCP.
   - Binds to `0.0.0.0:port` and starts listening.

2. **Accept new connections**

   ```go
   conn, err := l.Accept()
   ```

   - Blocks until a client completes the TCP handshake.
   - Returns a `net.Conn` object representing the connection.

3. **Handle connections concurrently**

   ```go
   go handleConnection(conn)
   ```

   - Each client gets its own goroutine.

4. **Read and write data**

   ```go
   conn.Read(buf)
   conn.Write(buf)
   ```

   - TCP is a stream of bytes — you decide how to parse it.

5. **Clean up**

   ```go
   defer conn.Close()
   ```

   - Ensures the OS socket is freed when done.

---

## TCP Theory (My Learning Notes)

### TCP Server Lifecycle

1. **Bind & Listen**
   `net.Listen("tcp4", ":PORT")`

   - Creates a server socket.
   - OS reserves the port for listening.

2. **Accept Connections**
   `listener.Accept()`

   - Waits until a TCP handshake completes.
   - Returns a new socket for that client.

3. **Data Transfer**
   `c.Read()` and `c.Write()`

   - TCP guarantees reliable, ordered delivery of bytes.

4. **Close Connection**
   `defer c.Close()`

   - Sends TCP FIN to close the connection gracefully.

---

### TCP Connection Stages

#### **1. Handshake**

```
Client → Server: SYN
Server → Client: SYN+ACK
Client → Server: ACK
```

#### **2. Data Transfer**

- Client sends a byte stream.
- Server reads from the OS socket buffer and processes it.
- Server writes bytes back to client.

#### **3. Teardown**

```
Client → Server: FIN
Server → Client: ACK
Server → Client: FIN
Client → Server: ACK
```

---

## Packet Flow Diagram

```
Client (nc)                     Server (Go TCP server)
   │   SYN   ────────────────────────►
   │ ◄───────────── SYN+ACK
   │   ACK   ────────────────────────►

   │ "Hello\n" ──────────────────────►  (c.Read)
   │ ◄───────────────────── "Hello\n"  (c.Write)

   │   FIN   ────────────────────────►
   │ ◄──────────────────────── ACK
   │ ◄──────────────────────── FIN
   │   ACK   ────────────────────────►
```

---

## How to Run

```bash
# Run server on port 8080
go run main.go 8080

# In another terminal, connect with netcat
nc localhost 8080
```

---

## Example Session

**Terminal 1 (Server)**

```
$ go run main.go 8080
Serving 127.0.0.1:52344
```

**Terminal 2 (Client)**

```
$ nc localhost 8080
Hello
Hello
```

---

## Future Improvements

- Add connection timeouts using `SetDeadline`.
- Limit the number of concurrent connections.
- Implement message framing (so each message is clearly separated).
- Log connections and data for debugging.

---

## References

- [Go net package](https://pkg.go.dev/net)
- [TCP Protocol RFC 793](https://www.rfc-editor.org/rfc/rfc793)
- [Wireshark TCP Analysis](https://www.wireshark.org/)

```

```
