# api
Todas as APIs relacionadas a Bolsa Capital


# Configuração

É necessário ter Go 1.15 ou superior instalado. É necessário rodar em ambiente Linux ou macOS.

Na primeira linha do Makefile, existe um alias para o local onde está instalado o [AIR](https://github.com/cosmtrek/air), que é o pacote responsável por fazer o live reload desse projeto.

Se você não tiver o AIR instalado, o projeto não irá rodar. Então é necessário [instalar ele primeiro](https://github.com/cosmtrek/air#installation).

Depois de instalado, você vai no Makefile, e coloca o caminho dele no alias da primeira linha. Exemplo:

```makefile
AIR=/caminho/para/o/binario/air

#... resto do arquivo Makefile...
```

Após isso, está tudo configurado e já é possível subir o servidor com o comando `make start`.

## Comandos

Os comandos disponíveis são:

```makefile
# Inicia o servidor e fica ouvindo por mudanças (live-reload);
make start

# Roda o projeto sem live-reload
make run

# Compila o projeto para o sistema operacional vigente
make build

# Compila o projeto para sistemas Linux
make build-linux

# Limpa arquivos binários de compilações passadas
make clean
```
