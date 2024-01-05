package servidor

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gtihub.com/franciscoklaus/golango-simple-crud/banco"
	"io"
	"net/http"
	"strconv"
)

type Usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// INSERE USUARIO NO BANCO DE DADOS
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisicao!"))
		return
	}
	var usuario Usuario
	if err = json.Unmarshal(corpoRequisicao, &usuario); err != nil {
		w.Write([]byte("Erro ao converter o usuario da requisicao!"))
		return
	}
	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}

	defer db.Close()

	//PREPARE STATEMENT
	statement, err := db.Prepare("INSERT INTO usuarios (nome, email) values (?,?)")
	if err != nil {
		w.Write([]byte("Erro ao escrever o usuario no banco de dados!"))
		return
	}
	defer statement.Close()

	insercao, err := statement.Exec(usuario.Nome, usuario.Email)
	if err != nil {
		w.Write([]byte("Erro ao executar o statement!"))
		return
	}

	idInserido, err := insercao.LastInsertId()
	if err != nil {
		w.Write([]byte("Erro ao obter o id inserido!"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuario inserido com sucesso! Id: %d", idInserido)))

}

// REMOVE USUARIO DO BANCO DE DADOS
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, err := strconv.ParseUint(parametros["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro!"))
		return
	}
	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}
	defer db.Close()
	statement, err := db.Prepare("DELETE FROM usuarios where id = ?")
	if err != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}

	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		w.Write([]byte("Erro ao deletar o usuário!"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Usuário deletado com sucesso do banco de dados!"))
}
