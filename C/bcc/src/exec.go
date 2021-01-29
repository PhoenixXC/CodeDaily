package main

import (
    "bytes"
    "encoding/binary"
    "fmt"
    "os"
    "os/signal"
    
    "C"
    bpf "github.com/iovisor/gobpf/bcc"
)

const source string = `
#include <linux/fs.h>

struct exec_event {
	char filename[256];
};
BPF_PERF_OUTPUT(events);

int kprobe__do_execveat_common(struct pt_regs* ctx, int fd, struct filename* filename) {
	struct exec_event event = {};
	bpf_probe_read_kernel(&(event.filename), sizeof(event.filename), (void*)filename->name);
	events.perf_submit((void*)ctx, (void*)&event, sizeof(event));
	return 0;
}
`

type execEvent struct {
    Filename [256]byte
}

func main() {
    module := bpf.NewModule(source, []string{})
    defer module.Close()
    
    execKprobeFd, err := module.LoadKprobe("kprobe__do_execveat_common")
    if err != nil {
        _, _ = fmt.Fprintf(os.Stderr, "Failed to load kprobe__do_execveat_common: %s\n", err)
        os.Exit(1)
    }
    
    err = module.AttachKprobe("do_execveat_common", execKprobeFd, -1)
    if err != nil {
        _, _ = fmt.Fprintf(os.Stderr, "Failed to attach do_execveat_common: %s\n", err)
        os.Exit(1)
    }
    
    table := bpf.NewTable(module.TableId("events"), module)
    
    channel := make(chan []byte)
    
    perfMap, err := bpf.InitPerfMap(table, channel, nil)
    if err != nil {
        _, _ = fmt.Fprintf(os.Stderr, "Failed to init perf map: %s\n", err)
        os.Exit(1)
    }
    
    sig := make(chan os.Signal, 1)
    signal.Notify(sig, os.Interrupt)
    
    go func() {
       var event execEvent
       for {
           data := <-channel
           err := binary.Read(bytes.NewBuffer(data), binary.LittleEndian, &event)
           if err != nil {
               fmt.Printf("failed to decode received data: %s\n", err)
               continue
           }
           filename := string(event.Filename[:])
           fmt.Printf("name = %s\n", filename)
       }
    }()
    
    perfMap.Start()
    <-sig
    perfMap.Stop()
    
}
