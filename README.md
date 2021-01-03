<h1 align="center">
    Dados publicos (CNPJ) para banco de dados PostgreSQL em Golang
</h1>
<p align="center">  
  <img alt="Made by Rui" src="https://img.shields.io/badge/Made%20by-ruiblaese-%2304D361">  
  <img alt="Made with Golang" src="https://img.shields.io/badge/Made%20with-Golang-%1f425f">  
  <img alt="Project top programing language" src="https://img.shields.io/github/languages/top/ruiblaese/dados-publicos-cnpj-para-postgresql">  
  <img alt="Repository size" src="https://img.shields.io/github/repo-size/ruiblaese/dados-publicos-cnpj-para-postgresql">   
</p>

<p align="center">
    <a href="#descricao">Descrição</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;    
    <a href="#contato">Contato</a>
</p>

## Descrição
Projeto criado para "bricar" com Golang, ideia é pegar os dados públicos disponibilizados pela receita e transformar em uma base dados em um banco de dados relacional, nesse caso PostgreSQL.

#### Site Dados Publicos CNPJ
http://idg.receita.fazenda.gov.br/orientacao/tributaria/cadastros/cadastro-nacional-de-pessoas-juridicas-cnpj/dados-publicos-cnpj

## Download Dados Publicos com todos CNPJ's

| Versão | Tempo Proc.     | Data Arq. Sefaz | Data Proc.    | Data Download | Tamanho Banco PostgreSQL | Tamanho Backup PostgreSQL | Registros | Link |
|:------:|:---------------:|:---------------:|:-------------:|:-------------:|:------------------------:|:-------------------------:|------------|------|
| 0.0.1  |  6h             |   xx-xx-xx      |  08-11-20     |   15-01-20    |          22,5 GB         |          3,15 GB          |            |  [Link para Download](https://drive.google.com/file/d/1oTWhFzPsJLMQwfLCUd38berjjy1cfmhq/view?usp=sharing)    |
| 0.0.2  |  12h            |   xx-xx-xx      |  11-11-20     |   10-01-20    |          21,7 GB         |          3,38 GB          | 45.153.134 |  [Link para Download](https://drive.google.com/file/d/1utdRqViqZlji8J2eVckB4bAI8BgSReI5/view?usp=sharing)    |
| 0.0.3  |  2h 37m         |   xx-xx-xx      |  12-11-20     |   10-01-20    |          - - -           |          - - -            | 45.153.134 |  v 0.0.2    |
| 0.1.0  |  5h 16m (total) |   23-11-20      |  02-01-21     |   02-01-21    |          - - -           |          3,56             | 46.536.906 |      |

Tempo Proc.: Tempo de leitura e gravação dos arquivos texto para PostgreSQL   
Tempo Total.: Download, Descompactação,leitura e gravação dos arquivos texto para PostgreSQL  

## Change Log
- **0.0.1**: Versão inicial, processado os arquivos baixado, todos os dados em uma tabela
- **0.0.2**: Corrigido os problemas com ORM, iniciado melhorias de modelagem. Encontrado problema em algumas linhas, valores fora da posição, ainda na análise de como resolver.
- **0.0.3**: 10 goroutines usando "wait group" com uma conexão com banco cada goroutines; Ainda com problema em algumas linhas que vieram erradas, mostra numero de caracteres certo mais vendo a linha no Notepad++ a diferença grande ao final da linha :-/
- **0.1.0**: Configuração Dockerfile e Docker Compose para executar rotina por completa pelo Docker; Não verificado as linhas com problema; Problema agora com tamanho das imagens geradas no Docker

### Proximas Implementações
- Melhorar forma de processamento para não criar imagem/container enorme (100GB++)
  - Individualmente descompactar, processar arquivo texto e apagar arquivos
- Encontrar problema existente em algumas linhas, quebra código da cidade e consequentemente não cadastra aquela empresa. 

## Ambiente
#### Softwares:
- Golang via Docker
- PostgreSQL: 13.x via Docker
#### Hardware:   
- i7 8700K   
- 32 GB RAM   
- 1 TB SSD M.2 NVME 3000mb/s+
     

## Contato

<a href="https://www.linkedin.com/in/ruiblaese/" target="_blank" >
  <img alt="Linkedin - Rui Richard Blaese" src="https://img.shields.io/badge/Linkedin--%23F8952D?style=social&logo=linkedin">
</a>&nbsp;&nbsp;&nbsp;
<a href="mailto:ruiblaese@gmail.com" target="_blank" >
  <img alt="Email - Rui Richard Blaese" src="https://img.shields.io/badge/Email--%23F8952D?style=social&logo=gmail">
</a> 
