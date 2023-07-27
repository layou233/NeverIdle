# NeverIdle

**Español** | [**English**](README_en.md) | [**简体中文**](README.md)

*Te quiero, pero ¿podrías no detener mi máquina?*

**Este archivo Readme podría no estar actualizado, considere revisar el Readme en [**ingles**](README_en.md)**

---

**A los usuarios no chinos:**

¡Gracias por interesarte en este programa! :-P  
Inicialmente, inventé esto para compartirlo con mis amigos chinos, pero no esperaba que se volviera popular en todo el mundo.  

Si necesitas ayuda, primero busca en Google y luego consulta los problemas reportados (issues).
Hablo chino e inglés. Para otros idiomas, por favor, traduce antes de hacer preguntas. :)

---

## Uso

Descarga el archivo ejecutable desde "Release". Ten en cuenta la diferencia entre las versiones para amd64 y arm64.

Inicia una sesión "screen" en el servidor y ejecútalo.
Si deseas aprender sobre el comando "screen", simplemente busca en Google.

Argumentos del comando:

```shell
./NeverIdle -cp 0.15 -m 2 -n 4h
```

Donde:

-c activa el desperdicio periódico de la CPU, seguido del intervalo entre los desperdicios.
Por ejemplo, para desperdiciar CPU cada 12 horas, 23 minutos y 34 segundos, el argumento sería `-c 12h23m34s`.
Solo sigue esta plantilla.

-cp activa el desperdicio de porcentaje de CPU de granulación gruesa, y la tasa de desperdicio cambiará en tiempo real según el nivel de uso de la máquina.
Si el desperdicio máximo del 20% de la CPU es `-cp 0.2`. El rango de valores del porcentaje es [0, 1], y ten cuidado de no usarlo junto con `-c`.

-m activa el desperdicio de memoria, seguido de un número en GiB.
Después de iniciarse, se ocupará la cantidad de memoria especificada y no se liberará hasta que el proceso sea detenido.

-n activa el desperdicio periódico de la red (ancho de banda), seguido del intervalo entre los desperdicios.
El formato del argumento es igual al de `-c`. ¡Se realizará una prueba de velocidad de Ookla periódicamente (y los resultados serán mostrados)!

-t especifica el número de conexiones simultáneas para el desperdicio periódico de la red.
El valor predeterminado es 10. Cuanto mayor sea el valor, más recursos se consumirán. En la mayoría de las situaciones, no es necesario cambiarlo.

-p especifica la prioridad del proceso, seguida de un valor de prioridad. Si no se especifica, se utilizará la prioridad más baja de la plataforma.
Para sistemas similares a UNIX (como Linux, FreeBSD y macOS), el rango de valores es [-20, 19], y cuanto mayor sea el número, menor será la prioridad.
Para Windows, consulta [la documentación oficial](https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-setpriorityclass).
Se recomienda no especificar un valor, ya que el valor predeterminado es la prioridad más baja, lo que permitirá que todos los demás procesos tengan prioridad.

*Todas las funciones que hayas configurado se ejecutarán inmediatamente una vez que inicies el programa, para que puedas ver el efecto.*
