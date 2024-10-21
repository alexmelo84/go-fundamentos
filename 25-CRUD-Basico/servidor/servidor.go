package servidor

import (
	"crud/db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID uint32 `json:"id"`
	Nome string `json:"nome"`
	Email string `json:"email"`
}

// CriarUsuario cria usuário
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpo, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.WriteHeader(400)
		w.Write([]byte("Corpo da requisição é obrigatório"))
		return
	}

	var usuario usuario

	if erro = json.Unmarshal(corpo, &usuario); erro != nil {
		w.WriteHeader(500)
		w.Write([]byte("Erro ao converter o Json em Struct"))
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		w.WriteHeader(500)
		w.Write([]byte("Erro ao se conectar ao banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("INSERT INTO usuarios (nome, email) VALUES (?, ?)")
	if erro != nil {
		w.WriteHeader(500)
		w.Write([]byte("Erro ao preparar a query de inserção!"))
		return
	}

	insert, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.WriteHeader(500)
		w.Write([]byte("Erro ao inserir usuário!"))
		return
	}

	idUsuario, erro := insert.LastInsertId()
	if erro != nil {
		w.WriteHeader(500)
		w.Write([]byte("Erro ao pegar a ID do usuário!"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso com a id %d", idUsuario)))
}

// BuscarUsuarios lista todos os usuários
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := db.Conectar()
	if erro != nil {
		w.WriteHeader(500)
		w.Write([]byte("Erro ao se conectar ao banco de dados"))
		return
	}
	defer db.Close()

	linhas, erro := db.Query("SELECT id, nome, email FROM usuarios")
	if erro != nil {
		w.WriteHeader(500)
		w.Write([]byte("Erro ao buscar usuários"))
		return
	}
	defer linhas.Close()

	var usuarios []usuario
	for linhas.Next() {
		var usuario usuario

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.WriteHeader(500)
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.WriteHeader(500)
		w.Write([]byte("Erro ao converter em JSON"))
		return
	}
	w.WriteHeader(http.StatusOK)
}

// BuscarUsuario lista todos um usuário
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.WriteHeader(500)
		w.Write([]byte("Erro ao converter o parâmetro para inteiro!"))
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		w.WriteHeader(500)
		w.Write([]byte("Erro ao se conectar ao banco de dados"))
		return
	}
	defer db.Close()

	linha, erro := db.Query("SELECT id, nome, email FROM usuarios WHERE id = ?", ID)
	if erro != nil {
		w.WriteHeader(500)
		w.Write([]byte("Erro ao buscar usuário"))
		return
	}

	var usuario usuario

	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.WriteHeader(500)
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}
	}

	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.WriteHeader(500)
		w.Write([]byte("Erro ao converter em JSON"))
		return
	}
	w.WriteHeader(http.StatusOK)
}

// AtualizarUsuario atualiza um usuário
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.WriteHeader(500)
		w.Write([]byte("Erro ao converter o parâmetro para inteiro!"))
		return
	}

	corpo, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.WriteHeader(400)
		w.Write([]byte("Corpo da requisição é obrigatório"))
		return
	}

	var usuario usuario

	if erro = json.Unmarshal(corpo, &usuario); erro != nil {
		w.WriteHeader(500)
		w.Write([]byte("Erro ao converter o Json em Struct"))
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		w.WriteHeader(500)
		w.Write([]byte("Erro ao se conectar ao banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("UPDATE usuarios SET nome = ?, email = ? WHERE id = ?")
	if erro != nil {
		w.WriteHeader(500)
		w.Write([]byte("Erro ao preparar a query de atualização!"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		w.WriteHeader(500)
		w.Write([]byte("Erro ao atualizar o usuário!"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeletarUsuario atualiza um usuário
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.WriteHeader(500)
		w.Write([]byte("Erro ao converter o parâmetro para inteiro!"))
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		w.WriteHeader(500)
		w.Write([]byte("Erro ao se conectar ao banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("DELETE FROM usuarios WHERE id = ?")
	if erro != nil {
		w.WriteHeader(500)
		w.Write([]byte("Erro ao preparar a query de deleção!"))
		return
	}

	if _, erro := statement.Exec(ID); erro != nil {
		w.WriteHeader(500)
		w.Write([]byte("Erro ao deletar o usuário!"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
