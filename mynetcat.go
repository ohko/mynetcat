package main

import (
	"flag"
	"log"
	"net"
	"os/exec"
	"runtime"
	// "golang.org/x/text/encoding/simplifiedchinese"
	// "golang.org/x/text/transform"
)

var (
	p = flag.String("port", ":22", "Listen port")
)

func main() {
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.Lshortfile)

	srv, err := net.Listen("tcp", *p)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Listen:", *p)

	for {
		conn, err := srv.Accept()
		if err != nil {
			break
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	defer log.Println("close:", conn.RemoteAddr().String())
	log.Println("connect:", conn.RemoteAddr().String())

	var bash *exec.Cmd
	if runtime.GOOS == "windows" {
		bash = exec.Command("cmd")
	} else {
		bash = exec.Command("bash")
	}
	// c := &convert{conn: conn}
	bash.Stdin = conn
	bash.Stdout = conn
	bash.Stderr = conn
	if err := bash.Run(); err != nil {
		log.Println(err)
		return
	}
}

// func translate(s []byte) []byte {
// 	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
// 	d, e := ioutil.ReadAll(reader)
// 	if e != nil {
// 		log.Println(e)
// 		return nil
// 	}
// 	return d
// }

// type convert struct {
// 	conn net.Conn
// }

// func (o *convert) Write(p []byte) (n int, err error) {
// 	switch runtime.GOOS {
// 	case "windows":
// 		reader := transform.NewReader(bytes.NewReader(p), simplifiedchinese.GBK.NewDecoder())
// 		d, e := ioutil.ReadAll(reader)
// 		if e != nil {
// 			return 0, e
// 		}

// 		m, err := o.conn.Write(translate(p))
// 		if m != len(d) {
// 			return m, err
// 		}
// 		return len(p), err
// 	default:
// 		return o.conn.Write(p)
// 	}
// }

// func (o *convert) Read(p []byte) (n int, err error) {
// 	// m, err := o.conn.Read(p)
// 	// switch runtime.GOOS {
// 	// case "windows":
// 	// 	p = translate(p[:m])
// 	// 	return len(p), err
// 	// default:
// 	// 	return m, err
// 	// }
// 	return o.conn.Read(p)
// }
