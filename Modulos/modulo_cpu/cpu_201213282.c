
// Linux Kernel/LKM headers: module.h es necesario para todos los modulos y kernel.h 
//y tambien para para KERN_INFO.
#include <linux/init.h>         // Incluido para __init y __exit macros
#include <linux/module.h>    // Incluido para todos los modulos de kernel
#include <linux/kernel.h>    // Incluido para KERN_INFO
#include <linux/sched.h>
#include <linux/proc_fs.h>
#include <linux/seq_file.h> 
#include <asm/uaccess.h> 
#include <linux/hugetlb.h>
#include <linux/fs.h>
#include <linux/sched/signal.h>
MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("Escribir informaciond del cpu");
MODULE_AUTHOR("Sergio De los Rios");

char buffer[256];
char * get_task_state(long state)
{
    switch (state) {
        case TASK_RUNNING:
            return "TASK_RUNNING";
        case TASK_INTERRUPTIBLE:
            return "TASK_INTERRUPTIBLE";
        case TASK_UNINTERRUPTIBLE:
            return "TASK_UNINTERRUPTIBLE";
        case __TASK_STOPPED:
            return "__TASK_STOPPED";
        case __TASK_TRACED:
            return "__TASK_TRACED";
        default:
        {
            sprintf(buffer, "Unknown Type:%ld\n", state);
            return buffer;
        }
    }
}

static int escritura_archivo(struct seq_file *s,void *v)
{
    unsigned int process_count = 0;
    pr_info("Number of processes:%u\n", process_count);
    seq_printf(s, "  Nombre: Sergio De los Ríos.\n");
    seq_printf(s, "  Carnet: 201213282.\n");
    seq_printf(s, "  Nombre: Randolph Muy.\n");
    seq_printf(s, "  Carnet: 201314112.\n");
    struct task_struct *task_list;
    for_each_process(task_list) {
        pr_info("Process: %s\t PID:[%d]\t State:%ld\t Padre:[%d]\n",task_list->comm, task_list->pid,task_list->state,task_list->parent);
        process_count++;    
    }
    return 0;
}
static int open_file(struct inode *inode, struct  file *file) {
  return single_open(file, escritura_archivo, NULL);
}

static struct file_operations fcpu =
{    
    .open = open_file,
    .read = seq_read
}; 
static int testmodulo_init(void)
{
    proc_create("memo_201213282_201314112", 0, NULL, &fcpu);
    printk(KERN_INFO "Sergio De los Rios\nRandolph Muy\n");//numero de carnet
 //Aqui iria el codigo a ejecutar
   return 0;    // Si el retorno no es 0 
                //quiere decir que el modulo no se ha podido cargar
}

static void testmodulo_cleanup(void)
{
   printk(KERN_INFO "Sistemas Operativos 1\n"); //se loga en el /var/log/messages
}
module_init(testmodulo_init);
module_exit(testmodulo_cleanup);
