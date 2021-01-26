#include <linux/module.h>
#include <linux/kernel.h>
#include <linux/init.h>

#define LOG_INFO(args...) printk(KERN_INFO "test: " args);

static int __init mod_init(void) {
  LOG_INFO("Hello Kernel Module\n");
  return 0;
}

static void __exit mod_exit(void) {
  LOG_INFO("Goodbye\n");
}

module_init(mod_init);
module_exit(mod_exit);

MODULE_LICENSE("GPL");
MODULE_AUTHOR("xuanc");
MODULE_DESCRIPTION("hello_module");