//********* para cargar el modulo ********************
#include <linux/module.h>   /* Todos los modulos lo necesitan */
#include <linux/kernel.h>   /* Ofrece la macro KERN_INFO */
#include <linux/init.h>     /* Ofrece las macros de inicio y fin */
#include <linux/seq_file.h>  //Esta libreria funciona para el seq_file el cual escribe y lee
//******** para tomar informacion de la memoria *********
#include <linux/hugetlb.h> //si_meminfo
#include <linux/proc_fs.h> //proc_create
#include <linux/sys.h>
#define BUFSIZE  150
struct sysinfo si;

static int write_and_read(struct seq_file *s,void *v)
{
    si_meminfo(&si);
    int32_t total_memoria,memoria_libre;
    total_memoria = (si.totalram * si.mem_unit)/(1024 * 1024);
    memoria_libre = (si.freeram * si.mem_unit)/(1024*1024);
    seq_printf(s, "  ------- Info Est1 ----- \n");
    seq_printf(s, "  Nombre: Randolph Muy\n");
    seq_printf(s, "  Carnet: 201314112\n");
    seq_printf(s, "  ------- Info Est2------ \n");
    seq_printf(s, "  Nombre: Sergio De los Ríos\n");
    seq_printf(s, "  Carnet: 201213282\n");
    seq_printf(s, "  ----------------------- \n");
    seq_printf(s, "  Memoria Total : \t %d MB\n",total_memoria);
    seq_printf(s, "  Memoria Libre : \t %d MB \n",memoria_libre);
    seq_printf(s, "  Memoria en uso: \t %i %%\n", (memoria_libre * 100)/total_memoria) ;
    return 0;
}

static int open_file(struct inode *inode, struct  file *file) {
  return single_open(file, write_and_read, NULL);
}

static struct file_operations fops =
{    
    .open = open_file,
    .read = seq_read
};

static int __init inicio(void)
{
    proc_create("memo_201314112_201213282", 0, NULL, &fops); //creacion del archivo
    printk(KERN_ALERT "Carnets: 201314112_201213282\n");
    return 0;
}

static void __exit salida(void)
{
    printk(KERN_ALERT "Sistemas Operativos 1\n");
}

// se Indica cuales son las funciones de inicio y fin
module_init(inicio);
module_exit(salida);

// Documentacion del modulo

MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("Modulo de Pruebas SOPES1");
MODULE_AUTHOR("Randolph_Muy-201404243___Sergio_De_los_Ríos-201213282");