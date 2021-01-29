#include <linux/init.h>
#include <linux/kernel.h>
#include <linux/kprobes.h>
#include <linux/module.h>

#define MAX_SYMBOL_LEN 64
#define MOD_NAME "test: "
#define log_info(fmt, ...) pr_info(MOD_NAME fmt, ##__VA_ARGS__)
#define log_error(fmt, ...) pr_info(MOD_NAME fmt, ##__VA_ARGS__)

static char symbol[MAX_SYMBOL_LEN] = "get_task_mm";
module_param_string(symbol, symbol, sizeof(symbol), 0644);

static struct kprobe kp = {
    .symbol_name = symbol,
};

static int handler_pre(struct kprobe* p, struct pt_regs* regs) {
    log_info("<%s> pre_handler: p->addr = %pF, ip = %lx, flags = 0x%lx\n",
             p->symbol_name, p->addr, regs->ip, regs->flags);
    log_info("argument maybe <%d>\n", ((struct task_struct* )((unsigned long )regs + 61))->pid);
    return 0;
}

static void handler_post(struct kprobe* p, struct pt_regs* regs,
                         unsigned long flags) {
    log_info("<%s> post_handler: p->addr = %pF, flags = 0x%lx\n", p->symbol_name,
             p->addr, regs->flags);
}

static int handler_fault(struct kprobe* p, struct pt_regs* regs, int trapnr) {
    log_info("fault_handler: p->addr = %pF, trap #%dn", p->addr, trapnr);
    return 0;
}

static int __init kprobe_init(void) {
    int ret;
    kp.pre_handler = handler_pre;
    kp.post_handler = handler_post;
    kp.fault_handler = handler_fault;

    ret = register_kprobe(&kp);
    if (ret < 0) {
        log_error("register_kprobe failed, returned %d\n", ret);
        return ret;
    }
    log_info("Planted kprobe at %pF\n", kp.addr);
    return 0;
}

static void __exit kprobe_exit(void) {
    unregister_kprobe(&kp);
    log_info("kprobe at %pF unregistered\n", kp.addr);
}

module_init(kprobe_init);
module_exit(kprobe_exit);
MODULE_LICENSE("GPL");
MODULE_AUTHOR("xuanc");
MODULE_DESCRIPTION("hello_module");