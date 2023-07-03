# POC sobre visão computacional <img src=https://github.com/eltonCasacio/gocv/blob/main/statics/images/gocvlogo.jpg width=30/>
![Badge](https://img.shields.io/static/v1?label=go&message=1.20&color=blue&style=for-the-badge&logo=Go)

## Projeto para aplicar conhecimento de visão computacional utilizando Golang

### Executando o a aplicação

1. clone o repositório em sua máquina.
2. rode `go mod tidy` para instalar as dependencias
3. rode `go run cmd/main.go`

irá abrir uma tela com a imagem da webcam.

com o app rodando voce pode enviar uma requisição REST via POST para aplicar o filtro.

  > POST http://localhost:3000/border HTTP/1.1 <br />
  > Content-Type:  application/json <br />
  > { <br />
  >    "filters": ["SOBEL", "CANNY"] <br />
  > }


podemos testar utilizando os arquivos .http ou através do postman por exemplo.<br />
