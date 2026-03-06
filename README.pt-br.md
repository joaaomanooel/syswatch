# SysWatch

<p align="center">
  <img src="https://img.shields.io/badge/versão-0.0.1-blue" alt="Versão">
  <img src="https://img.shields.io/badge/go-1.24-blue" alt="Versão do Go">
  <img src="https://img.shields.io/badge/cobertura-90%25-green" alt="Cobertura">
  <img src="https://img.shields.io/github/license/joaaomanooel/syswatch" alt="Licença">
</p>

SysWatch é uma ferramenta CLI poderosa para monitoramento de sistema em tempo real. Ela fornece visibilidade instantânea do uso da CPU, consumo de memória, uso do disco e contagem de processos diretamente no seu terminal.

## Recursos

- **Monitoramento em Tempo Real**: Assista métricas de CPU, memória, disco e processos atualizarem em tempo real
- **Multiplataforma**: Funciona no macOS, Linux e Windows
- **Leve**: Binário único, sem dependências externas
- **Intervalos Personalizáveis**: Ajuste a frequência de atualização conforme sua necessidade

## Instalação

### macOS

```bash
# Usando Homebrew
brew install joaaomanooel/homebrew-syswatch/syswatch

# Ou baixe manualmente
curl -L https://github.com/joaaomanooel/syswatch/releases/latest/download/syswatch-darwin-amd64.zip -o syswatch.zip
unzip syswatch.zip
./syswatch
```

### Linux

```bash
# Usando APT (Debian/Ubuntu)
# Adicione o repositório (em breve)

# Usando Snap
sudo snap install syswatch --classic

# Usando RPM (Fedora/RHEL)
sudo rpm -i https://github.com/joaaomanooel/syswatch/releases/latest/download/syswatch-*.rpm

# Baixe manualmente
curl -L https://github.com/joaaomanooel/syswatch/releases/latest/download/syswatch-linux-amd64.zip -o syswatch.zip
unzip syswatch.zip
./syswatch
```

### Windows

```bash
# Baixe pelos Lançamentos do GitHub
# Extraia o arquivo ZIP e execute syswatch.exe
```

## Uso

### Uso Básico

```bash
syswatch monitor
```

### Intervalo Personalizado

```bash
# Atualizar a cada 5 segundos
syswatch monitor -i 5s

# Atualizar a cada 500 milissegundos
syswatch monitor -i 500ms
```

### Versão

```bash
syswatch version
```

## Desenvolvimento

### Pré-requisitos

- Go 1.24 ou superior
- Make

### Compilar do Código-Fonte

```bash
# Clone o repositório
git clone https://github.com/joaaomanooel/syswatch.git
cd syswatch

# Compile para sua plataforma atual
make build

# Ou execute diretamente
make run
```

### Executar Testes

```bash
make test
```

### Executar Linter

```bash
make lint
```

### Compilar para Todas as Plataformas

```bash
# Compila para macOS, Linux, Windows (amd64 + arm64)
make build
```

## Configuração

O SysWatch usa padrões sensíveis que funcionam bem para a maioria dos casos de uso. Nenhum arquivo de configuração é necessário.

## Contribuindo

Contribuições são bem-vindas! Sinta-se livre para enviar um Pull Request.

## Licença

Este projeto está licenciado sob a Licença MIT - veja o arquivo LICENSE para detalhes.

## Agradecimentos

- [gopsutil](https://github.com/shirou/gopsutil) para métricas do sistema
- [Cobra](https://github.com/spf13/cobra) para framework CLI
