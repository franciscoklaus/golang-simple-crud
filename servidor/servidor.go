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

// RECUPERA TODOS OS USUÁRIOS DO BANCO
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}
	defer db.Close()

	linhas, err := db.Query("SELECT * FROM usuarios")
	if err != nil {
		w.Write([]byte("Erro ao retornar os usuarios!"))
		return
	}

	defer linhas.Close()

	// SLICE DE USUARIOS
	var usuarios []Usuario

	for linhas.Next() {
		var usuario Usuario
		if err := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); err != nil {
			w.Write([]byte("Erro ao escanear o usuário!"))
			return
		}
		usuarios = append(usuarios, usuario)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(usuarios); err != nil {
		w.Write([]byte("Erro ao converter os usuarios para json!"))
	}

}

// RECUPERA USUÁRIO DO BANCO
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, err := strconv.ParseUint(parametros["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Erro ao converter o parametro para inteiro!"))
		return
	}
	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}

	defer db.Close()

	linha, err := db.Query("SELECT * FROM usuarios WHERE id = ?", ID)
	if err != nil {
		w.Write([]byte("Erro ao recuperar o usuario do banco!"))
	}

	var usuario Usuario

	if linha.Next() {
		if err := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); err != nil {
			w.Write([]byte("Erro ao escanear o usuario!"))
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(usuario); err != nil {
		w.Write([]byte("Erro ao escanear o usuario!"))
		return
	}

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

func ContarUsuario(w http.ResponseWriter, r *http.Request) {
	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}

	defer db.Close()

	contagem, err := db.Query("SELECT COUNT(*) FROM usuarios")
	if err != nil {
		w.Write([]byte("Erro ao realizar a query!"))
		return
	}
	var retorno int
	if contagem.Next() {
		if err := contagem.Scan(&retorno); err != nil {
			w.Write([]byte("Erro ao escanear a contagem!"))
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(retorno)

}
