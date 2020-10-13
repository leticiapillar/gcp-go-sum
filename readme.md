## Curso: Desenvolvimento de Aplicações Modernas e Escaláveis com Microsserviços
### DevOps: Desafio de CI


Aplicação em GO com  uma função chamada soma que receberá dois parâmetros e retornará a adição desses dois valores.

Configuração de uma CI no GCP - Google Cloud Plataform

***Rodando o main***
```
> go run src/main/main.go
Resultado da soma é: 10
```

***Rodando os testes***
```
> go test -v src/soma/main_test.go src/soma/main.go 
=== RUN   TestMainSuccess
    main_test.go:10: Function sum success
--- PASS: TestMainSuccess (0.00s)
PASS
ok      command-line-arguments  1.003s
```

***Fontes de pesquisa***
- https://cloud.google.com/cloud-build/docs/building/build-go
- https://github.com/GoogleCloudPlatform/cloud-builders/blob/master/go/examples/hello_world/Dockerfile
- https://github.com/GoogleCloudPlatform/cloud-builders/tree/master/go
