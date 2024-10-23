# Guia de Rotas para o Site no Server do Go

## Índice
1. [Requisitos](#)
2. [Adicionar Rota ao Servidor](#adicionando-routes-ao-servidor-gin-em-go)

## Requisitos
- O arquivo que deseja introduzir nas rotas precisa estar na [pasta html](/html) do projeto.

## Adicionando Routes ao Servidor Gin em Go

1. **Adicionar Rota na Construção do server**
    - No Arquivo [Main](/server.go#L21) Adicione o seguinte bloco
    ```go
        router.GET("/<endereço da pagina>", controller.<endereço da pagina>)
    ```
    - subistitua os `<endereço da pagina>` pelo real endereço da pagina.
        - Exemplo:
        ```go
            router.GET("/pagina-teste", controller.pagina_teste)
        ```

2. **Adicionar Arquivo HTML ao Sistema de Rota**
    - No Arquivo do [Controller](/controller/controller.go) Adicione o seguinte bloco
    ```go
        func endereço_da_pagina (c *gin.Context) {
            c.HTML(200, "arquivo.html", gin.H{
                "title": "Main website",
            })
        }
    ```
    - subistitua o `endereço_da_pagina` com o mesmo nome que adicionou apos o apos o `controller.` no arquivo [Main](/server.go#L21)
        - Exemplo:
        ```go
            func pagina_teste (c *gin.Context) {
                c.HTML(200, "", gin.H{
                    "title": "Main website",
                })
            }
        ```
    - introduza o nome do arquivo da pagina HTML subistituindo o `"arquivo.html"` no bloco de exemplo.
        - **o root do servidor é a pasta html então tudo precisa ser especificada apartir dela**
            - Exemplo:
            
                se o arquivo está em `challenge-espr/html/paginaTeste.html`
                
                vc so coloca `paginaTeste.html`.
        - Exemplo: 
        ```go
            func pagina_teste (c *gin.Context) {
                c.HTML(200, "paginaTeste.html", gin.H{
                    "title": "Main website",
                })
            }
        ```
    

