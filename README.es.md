# SysWatch

<p align="center">
  <img src="https://img.shields.io/badge/versión-0.0.1-blue" alt="Versión">
  <img src="https://img.shields.io/badge/go-1.24-blue" alt="Versión de Go">
  <img src="https://img.shields.io/badge/cobertura-90%25-green" alt="Cobertura">
  <img src="https://img.shields.io/github/license/joaaomanooel/syswatch" alt="Licencia">
</p>

SysWatch es una herramienta CLI poderosa para el monitoreo del sistema en tiempo real. Proporciona visibilidad instantánea del uso de CPU, consumo de memoria, uso del disco y conteo de procesos directamente en tu terminal.

## Características

- **Monitoreo en Tiempo Real**: Observa las métricas de CPU, memoria, disco y procesos actualizándose en tiempo real
- **Multiplataforma**: Funciona en macOS, Linux y Windows
- **Ligero**: Binario único, sin dependencias externas
- **Intervalos Personalizables**: Ajusta la frecuencia de actualización según tus necesidades

## Instalación

### macOS

```bash
# Usando Homebrew
brew install joaaomanooel/homebrew-syswatch/syswatch

# O descarga manualmente
curl -L https://github.com/joaaomanooel/syswatch/releases/latest/download/syswatch-darwin-amd64.zip -o syswatch.zip
unzip syswatch.zip
./syswatch
```

### Linux

```bash
# Usando APT (Debian/Ubuntu)
# Agrega el repositorio (próximamente)

# Usando Snap
sudo snap install syswatch --classic

# Usando RPM (Fedora/RHEL)
sudo rpm -i https://github.com/joaaomanooel/syswatch/releases/latest/download/syswatch-*.rpm

# Descarga manualmente
curl -L https://github.com/joaaomanooel/syswatch/releases/latest/download/syswatch-linux-amd64.zip -o syswatch.zip
unzip syswatch.zip
./syswatch
```

### Windows

```bash
# Descarga desde GitHub Releases
# Extrae el archivo ZIP y ejecuta syswatch.exe
```

## Uso

### Uso Básico

```bash
syswatch monitor
```

### Intervalo Personalizado

```bash
# Actualizar cada 5 segundos
syswatch monitor -i 5s

# Actualizar cada 500 milisegundos
syswatch monitor -i 500ms
```

### Versión

```bash
syswatch version
```

## Desarrollo

### Requisitos Previos

- Go 1.24 o superior
- Make

### Compilar desde el Código Fuente

```bash
# Clona el repositorio
git clone https://github.com/joaaomanooel/syswatch.git
cd syswatch

# Compila para tu plataforma actual
make build

# O ejecuta directamente
make run
```

### Ejecutar Pruebas

```bash
make test
```

### Ejecutar Linter

```bash
make lint
```

### Compilar para Todas las Plataformas

```bash
# Compila para macOS, Linux, Windows (amd64 + arm64)
make build
```

## Configuración

SysWatch usa valores predeterminados que funcionan bien para la mayoría de los casos de uso. No se requiere archivo de configuración.

## Contribuyendo

¡Las contribuciones son bienvenidas! Siéntete libre de enviar un Pull Request.

## Licencia

Este proyecto está licenciado bajo la Licencia MIT - consulta el archivo LICENSE para más detalles.

## Agradecimientos

- [gopsutil](https://github.com/shirou/gopsutil) por las métricas del sistema
- [Cobra](https://github.com/spf13/cobra) por el framework CLI
