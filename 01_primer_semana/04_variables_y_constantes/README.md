**Variables y constantes**

1. En el archivo **main.go** vemos distintas formas de declarar variables y constantes, pero antes de mirar el código, veamos cómo es la sintaxis que vamos a usar.

2. Constantes:
    i. Para definir constantes se hace de la siguiente forma:
**const <nombre> <tipo> = <valor>**

3. Variables:
La sintaxis completa para declarar e inicializar una variable será:
    i. **var <nombre> <tipo> = <valor>**

    ii. Go trabaja con inferencia de tipos si es que no lo definimos como en el siguiente caso:

**var <nombre> = <valor>**

    iii. Eso mismo también lo podemos hacer sin escribir la palabra reservada var:

**<nombre> := <valor>**

4. En golang tenemos el concepto de zero value que hace que cuando no inicializamos una variable, se le asigne un valor por defecto según el tipo de dato con el que estemos declarando la variable. 

5. Probemos ejecutar la función main para ver qué se imprime y analicemos un poco el código.

6. ¿Qué pasa si declaro una variable y no la uso? Probemos.
 

 