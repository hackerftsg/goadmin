# goadmin 1.0.0

versão 1.0.0
estou trabalhando para melhorar bastante o script, em breve nova versão com bem mais opções e suportabilidade

o goadmin é uma ferramenta que bruta em massa ou somente uma url, procurando paginas administrativas
que retornam um status code e verifica se possui um formulario de login

você vai precisar instalar o [Golang](https://golang.org/dl/)

usagem, somente url (deve conter o protocolo na frente e a barra no final)

    go run goadmin.go http://target.com/
    
usagem, com arquivo

    go run goadmin.go urls.txt
