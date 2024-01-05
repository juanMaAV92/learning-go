# learning-go

## usefull commands

### go mod init
Este comando crea un archivo go.mod en el directorio actual, si no existe, y lo inicializa con el nombre del módulo, obtenido del argumento de ruta. El argumento de ruta debe ser un camino o un patrón de ruta (como ./...) correspondiente a los paquetes que se encuentran fuera de $GOPATH/src. Si el argumento de ruta no está presente, go mod init asume que el módulo se encuentra en el directorio actual.

```go
go mod init github.com/userName/repoName
```

### go get
Este comando agrega paquetes al módulo actual y lo actualiza (o usa la versión actual) si ya está presente.

```go
go get github.com/userName/repoName
```
> github.com/userName/repoName refiere al nombre del paquete que se desea agregar al módulo actual.


### go mod tidy
Este comando agrega cualquier nueva dependencia necesaria para construir el módulo actual y elimina las dependencias que no se necesitan. También actualiza go.mod para usar la nueva versión de cada dependencia que se conserva, si está disponible.

```go
go mod tidy
```
### go run
Este comando compila y ejecuta el programa principal del módulo actual.

```go
go run main.go
```

### go build
Este comando compila los paquetes y dependencias del módulo actual o los especificados en los argumentos, junto con los archivos de prueba, pero no los instala.

```go
go build main.go
```

### go test 
Este comando ejecuta las pruebas unitarias de los paquetes especificados.

```go
go test
```

Verifiquemos el coverage de las pruebas unitarias

```go
go test -cover
go test -coverprofile=coverage.out
```

Para tener metricas legibles, tenemos que ejecutar el siguiente comando

```go
go tool cover -func=coverage.out
```

Con el archivo generado, tenemos que ejecutar el siguiente comando para ver el coverage en el navegador

```go
go tool cover -html=coverage.out
```

Veamos el profile de las pruebas unitarias

```go
go test -cpuprofile=cpu.out
go tool pprof cpu.out
```

## Some tips

### channel 
 
- **unbuffered:** Este tipo de canal no tiene un búfer interno, lo que significa que solo puede contener un valor a la vez. Si un goroutine intenta enviar un valor a un canal sin búfer, se bloqueará hasta que otro goroutine reciba el valor del canal.

- **buffered:** Este tipo de canal tiene un búfer interno, lo que significa que puede contener cero o más valores. Si el búfer está lleno, el siguiente envío se bloqueará hasta que otro goroutine reciba un valor del canal.