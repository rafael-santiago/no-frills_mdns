# MELEIAMEEU...

Isso aqui é uma implementação simples, não inflada e sem frescura de um servidor minimalista
de `MDNS`. Fiz essa implementação para um projeto embarcado de final de semana, para facilitar
os trecos de `zero conf` que precisava para continuar com minhas loucuras. Porém, não queria
vender a minha alma usando um `zero conf` onde trouxesse para o projeto fadas, duendes,
insetos e doninhas. Só queria um código direto, sem muita cambalhota e manutenível em que
pudesse resolver algumas `tralhas.local`. E acaboOoOoU! Tchau!

O que? Ah é!

Para compilar é só rodar `go build`. Vai gerar um executável de demonstração bem sugestivo:
`nofrills-mdns`. Roda ele e você terá uma resolução `MDNS` por dois minutos em sua `LAN`.

O código é bem simples de entender como funciona. Você registra as resoluções numa
estrutura, passa essa estrutura para uma função principal do pacote `mdns` e acabou.
A função também recebe um canal `booleano`, quando você envia `true` para esse canal
o servidor `MDNS` para.

Aproveite!

Nota: Não aceito `PR` corrigindo `MELEIAMEEU`... para `LEIAME`... Assim tá `SERTO`! Valeu!

Oh! English please! Okay!

This is a no-frills, not bloated and well-simple implementation of a minimalist `MDNS` server.
I wrote this for one embed weekend project of mine for making easy some stuff related to `zero
conf`. However I did not want to sell my soul using a `zero conf` that could bring to the
project fairies, elfs, some insects and weasels. I just wanted a straightforward code, with
no somersaults and maintainable that could resolve some `trinket.local`. That's it! see ya!

What? Oh yeah!

In order to build it is only about run `go build` at the toplevel directory. It will generate
a sample executable quite meaningful: `nofrills-mdns`. Run it and you will get `MDNS` resolution
during two minutes in your `LAN`.

The code is fairly simple of understanding how works. You register the resolutions into
a struct, pass this struct to a main function of the package `mdns` and you done.
The function also receives a `boolean` channel, when you send `true` to this channel
the `MDNS` server stops.

Enjoy!

-- Rafael
