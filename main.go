package main

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/google/uuid"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	mensagemHttp string
)

func main() {
	mensagemHttp = "Aguarde, gerando banco de dados"

	go GeraBancoDeDadosComCNPJs()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": mensagemHttp,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func GeraBancoDeDadosComCNPJs() {
	time.Sleep(5 * time.Second)

	start := time.Now()
	godotenv.Load()

	dsn := os.Getenv("DSN_POSTGRES")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	downloadTodosArquivosDeCnpj()

	// Migrate the schema
	db.AutoMigrate(&Empresa{})

	listaArquivosDescompactados := descompactaArquivos()
	LerArquivoESalvarNoBanco(listaArquivosDescompactados, db)

	elapsed := time.Since(start)
	log.Printf("took %s", fmt.Sprint(elapsed))

	mensagemHttp = "Banco de dados Pronto! Tempo: " + fmt.Sprint(elapsed)
}

var wg sync.WaitGroup

func worker(input chan string) {
	defer wg.Done()

	dsn := os.Getenv("DSN_POSTGRES")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Consumer: Process items from the input channel and send results to output channel
	for line := range input {
		linhaStruct := ConverteLinhaParaStruct(db, line)
		err := db.Create(&linhaStruct).Error
		if err != nil {
			fmt.Println(err)
		}
	}
}

func LerArquivoESalvarNoBanco(arquivos []string, db *gorm.DB) {

	input := make(chan string)
	workers := 10

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go worker(input)
	}

	for _, arquivo := range arquivos {

		file, err := os.Open(arquivo)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := toUtf8(scanner.Bytes())
			if line[0] == '1' {
				input <- line

			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

	}

	close(input)
	wg.Wait()
}

func toUtf8(iso8859_1_buf []byte) string {
	buf := make([]rune, len(iso8859_1_buf))
	for i, b := range iso8859_1_buf {
		buf[i] = rune(b)
	}
	return string(buf)
}

type Estado struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	Sigla     string
}

func (Estado) TableName() string {
	return "estado"
}

type Municipio struct {
	ID           uint `gorm:"primaryKey"`
	CreatedAt    time.Time
	Codigo       int64
	Nome         string
	Estado       int64
	EstadoStruct Estado `gorm:"foreignKey:Estado"`
}

func (Municipio) TableName() string {
	return "municipio"
}

type Empresa struct {
	ID                         uint `gorm:"primaryKey"`
	CreatedAt                  time.Time
	TipoDeRegistro             string
	IndicadorFullDiario        string
	TipoAtualizacao            string
	Cnpj                       string `gorm:"index"`
	IdentificadorMatrizFilial  string
	RazaoSocialNomeEmpresarial string `gorm:"index"`
	NomeFantasia               string `gorm:"index"`
	SituacaoCadastral          string
	DataSituacaoCadastral      string
	MotivoSituacaoCadastral    string
	NomeCidadeExterior         string
	CodigoPais                 string
	NomePais                   string
	CodigoNaturezaJuridica     string
	DataInicioAtividade        string
	CnaeFiscal                 string
	DescricaoTipoLogradouro    string
	Logradouro                 string
	Numero                     string
	Complemento                string
	Bairro                     string
	Cep                        string
	Municipio                  int64
	MunicipioStruct            Municipio `gorm:"foreignKey:Municipio"`
	Telefone1                  string
	Telefone2                  string
	Fax                        string
	CorreioEletronicoEmail     string
	QualificacaoDoResponsavel  string
	CapitalSocialDaEmpresa     string
	PorteEmpresa               string
	OpcaoPeloSimples           string
	DataOpcaoSimples           string
	DataExclusaoSimples        string
	OpcaoPeloMei               string
	SituacaoEspecial           string
	DataSituacaoEspecial       string
}

type Tabler interface {
	TableName() string
}

func (Empresa) TableName() string {
	return "empresa"
}

func obtemValorDaLinha(linha string, posicao int, tamanho int) string {
	return strings.TrimSpace(linha[posicao-1 : (posicao - 1 + tamanho)])
}

func buscaCadastraMunicipio(db *gorm.DB, codigo string, nome string, estado Estado) Municipio {
	var municipio Municipio
	i, _ := strconv.ParseInt(codigo, 0, 64)
	result := db.Where(&Municipio{Codigo: i}).First(&municipio)

	if result.RowsAffected > 0 {
		return municipio
	}
	municipio = Municipio{Codigo: i, Nome: nome, EstadoStruct: estado}
	db.Create(&municipio)
	return municipio

}

func buscaCadastraEstado(db *gorm.DB, sigla string) Estado {
	var uf Estado

	result := db.Where(&Estado{Sigla: sigla}).First(&uf)

	if result.RowsAffected > 0 {
		return uf
	}
	uf = Estado{Sigla: sigla}
	db.Create(&uf)
	return uf
}

func ConverteLinhaParaStruct(db *gorm.DB, linha string) Empresa {

	estado := buscaCadastraEstado(db, obtemValorDaLinha(linha, 683, 2))
	empresa := Empresa{
		TipoDeRegistro:             obtemValorDaLinha(linha, 1, 1),
		IndicadorFullDiario:        obtemValorDaLinha(linha, 2, 1),
		TipoAtualizacao:            obtemValorDaLinha(linha, 3, 1),
		Cnpj:                       obtemValorDaLinha(linha, 4, 14),
		IdentificadorMatrizFilial:  obtemValorDaLinha(linha, 18, 1),
		RazaoSocialNomeEmpresarial: obtemValorDaLinha(linha, 19, 150),
		NomeFantasia:               obtemValorDaLinha(linha, 169, 55),
		SituacaoCadastral:          obtemValorDaLinha(linha, 224, 2),
		DataSituacaoCadastral:      obtemValorDaLinha(linha, 226, 8),
		MotivoSituacaoCadastral:    obtemValorDaLinha(linha, 234, 2),
		NomeCidadeExterior:         obtemValorDaLinha(linha, 236, 55),
		CodigoPais:                 obtemValorDaLinha(linha, 291, 3),
		NomePais:                   obtemValorDaLinha(linha, 294, 70),
		CodigoNaturezaJuridica:     obtemValorDaLinha(linha, 364, 4),
		DataInicioAtividade:        obtemValorDaLinha(linha, 368, 8),
		CnaeFiscal:                 obtemValorDaLinha(linha, 376, 7),
		DescricaoTipoLogradouro:    obtemValorDaLinha(linha, 383, 20),
		Logradouro:                 obtemValorDaLinha(linha, 403, 60),
		Numero:                     obtemValorDaLinha(linha, 463, 6),
		Complemento:                obtemValorDaLinha(linha, 469, 156),
		Bairro:                     obtemValorDaLinha(linha, 625, 50),
		Cep:                        obtemValorDaLinha(linha, 675, 8),
		MunicipioStruct:            buscaCadastraMunicipio(db, obtemValorDaLinha(linha, 685, 4), obtemValorDaLinha(linha, 689, 50), estado),
		Telefone1:                  obtemValorDaLinha(linha, 739, 12),
		Telefone2:                  obtemValorDaLinha(linha, 751, 12),
		Fax:                        obtemValorDaLinha(linha, 763, 12),
		CorreioEletronicoEmail:     obtemValorDaLinha(linha, 775, 115),
		QualificacaoDoResponsavel:  obtemValorDaLinha(linha, 890, 2),
		CapitalSocialDaEmpresa:     obtemValorDaLinha(linha, 892, 14),
		PorteEmpresa:               obtemValorDaLinha(linha, 906, 2),
		OpcaoPeloSimples:           obtemValorDaLinha(linha, 908, 1),
		DataOpcaoSimples:           obtemValorDaLinha(linha, 909, 8),
		DataExclusaoSimples:        obtemValorDaLinha(linha, 917, 8),
		OpcaoPeloMei:               obtemValorDaLinha(linha, 925, 2),
		SituacaoEspecial:           obtemValorDaLinha(linha, 926, 23),
		DataSituacaoEspecial:       obtemValorDaLinha(linha, 949, 8),
	}
	return empresa
}

func descompactaArquivos() []string {

	var listaArquivosDescompactados []string
	var files []string

	root := "./downloads"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if strings.Index(file, ".zip") > 0 {

			destino := strings.ReplaceAll(file, "downloads", "temp")
			destino = strings.ReplaceAll(destino, ".zip", "")

			fmt.Println("iniciando descompactacao", file, "->", destino)
			arquivosDescompactados, _ := Unzip(file, destino)
			fmt.Println(file, "->", destino, files)
			fmt.Println(file, "->", destino, "OK")

			listaArquivosDescompactados = append(listaArquivosDescompactados, arquivosDescompactados[0])
		}

	}
	return listaArquivosDescompactados

}

func downloadFile(filepath string, url string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		out.Close()
		os.Remove(filepath)
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func downloadTodosArquivosDeCnpj() {

	if _, err := os.Stat("downloads"); os.IsNotExist(err) {
		os.Mkdir("downloads", 0755)
	}

	const LinkDownload = "http://200.152.38.155/CNPJ/"

	for index := 1; index <= 20; index++ {

		var aux = ""
		if index <= 9 {
			aux = "0"
		}
		aux = aux + strconv.Itoa(index)

		fileName := "DADOS_ABERTOS_CNPJ_" + aux + ".zip"

		log.Printf("try download: %v", fileName)

		err := downloadFile("./downloads/"+fileName, LinkDownload+fileName)

		if err == nil {
			log.Printf("download: %v -> OK", fileName)
		} else {
			log.Printf("download: %v -> ERROR -> %v", fileName, err)
		}
	}
}

func Unzip(src string, dest string) ([]string, error) {

	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}
