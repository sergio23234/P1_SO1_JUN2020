
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
    struct task_struct *task_list;
    struct task_struct *task_child;
    struct list_head *list;
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
    seq_printf(s, "  Nombre: Sergio De los Ríos.\n");
    seq_printf(s, "  Carnet: 201213282.\n");
    seq_printf(s, "  Nombre: Randolph Muy.\n");
    seq_printf(s, "  Carnet: 201314112.\n");
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
    unsigned int process_count = 0;
    printk(KERN_INFO "---Inicio-----\n");
    for_each_process(task_list) {
         printk(KERN_INFO "PID: %d\tPROCESS: %s\tSTATE: %ld\n",task_list->pid, task_list->comm, task_list->state);
         process_count++;    
         list_for_each(list, &task_list->children){                        
           	task_child = list_entry( list, struct task_struct, sibling );   
         	printk(KERN_INFO "hijo:\tPID: %d\tPROCESS: %s\tSTATE: %ld\n",task_child->pid,task_child->comm, task_child->state);
        }
        printk("-----------------------------------------------------");    /*for aesthetics*/
    }
    pr_info("Number of processes:%u\n", process_count);
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
