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
Projeto criado para "bricar" com Golang, ideia é pegar os dados publicos disponibilizados pela receita e transformar em uma base dados em um banco de dados relacional, nesse caso PostgreSQL.

#### Site Dados Publicos CNPJ
http://idg.receita.fazenda.gov.br/orientacao/tributaria/cadastros/cadastro-nacional-de-pessoas-juridicas-cnpj/dados-publicos-cnpj

## Download Dados Publicos com todos CNPJ's

| Versão | Tempo Proc. | Data Proc.    | Data Download | Tamanho Download | Tamanho Descompactado | Tamanho Banco PostgreSQL | Tamanho Backup PostgreSQL | Registros | Link |
|:------:|:-----------:|:-------------:|:-------------:|:----------------:|:---------------------:|:------------------------:|:-------------------------:|------------|------|
| 0.0.1  |  6h         |  08-11-20     |   15-01-20    |      6,06 GB     |        96,7 GB        |          22,5 GB         |          3,15 GB          |            |  [Link para Download](https://drive.google.com/file/d/1oTWhFzPsJLMQwfLCUd38berjjy1cfmhq/view?usp=sharing)    |
| 0.0.2  |  12h        |  11-11-20     |   10-01-20    |      6,46 GB     |        102,0 GB       |          21,7 GB         |          3,38 GB          | 45.153.134 |  [Link para Download](https://drive.google.com/file/d/1utdRqViqZlji8J2eVckB4bAI8BgSReI5/view?usp=sharing)    |
| 0.0.3  |  2h 37m     |  12-11-20     |   10-01-20    |      6,46 GB     |        102,0 GB       |          - - -           |          - - -            | 45.153.134 |  v 0.0.2    |
| 0.1.0  |  5h 16m     |  02-01-21     |   02-01-21    |      -,-- GB     |        ---,- GB       |          - - -           |          - - -            | 46.536.906 |  v 0.1.0    |

Tempo Proc.: Tempo de leitura e gravação dos arquivos texto para PostgreSQL  

## Change Log
- **0.0.1**: Versão inicial, processado os arquivos baixado, todos os dados em uma tabela
- **0.0.2**: Corrigido os problemas com ORM, iniciado melhorias de modelagem. Encontrado problema em algumas linhas, valores fora da posição, ainda na análise de como resolver.
- **0.0.3**: 10 goroutines usando "wait group" com uma conexão com banco cada goroutines; Ainda com problema em algumas linhas que vieram erradas, mostra numero de caracteres certo mais vendo a linha no Notepad++ a diferença grande ao final da linha :-/ 

## Ambiente
#### Softwares:
- Golang: 1.15.4 windows/amd64
- PostgreSQL: 13.0, compiled by Visual C++ build 1914, 64-bit
- Windows 10 Pro: Build 19042.630
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
