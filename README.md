# hexagonal-go
Projeto de aprendizado de arquitetura hexagonal em go.

Para iniciar nosso ambiente localmente, utilize o comando a seguir na raiza do projeto:
```
docker compose up -d --build
```

Para abrir o container em modo iterativo utilize o comando a seguir:
```
docker exec -it appproduct bash
```

Para gerenciamento das dependências externas de forma automática foi utilizado o seguinte comando dentro do nosso container em modo iterativo:
```
go mod init github.com/augustopedro/hexagonal-go
```

Caso vá desenvolver em Go é recomendado que instale a linguagem seguindo os passos:
Run:
```
sudo apt update
sudo apt upgrade
```

Instale a versão Go:
```
sudo apt install golang-go
```