package handlers

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jailtonjunior94/udemy-golang-fundamentos-adicionais/database/banco"
)

type usuario struct {
	ID    uint   `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// CriarUsuario - Essa função a funcionalidade de criar um novo usuário
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("500 - Erro ao procurar body!"))
		return
	}

	var usuario usuario
	if err := json.Unmarshal(body, &usuario); err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("400 - Erro ao converter!"))
		return
	}

	conn, err := banco.Conectar()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Erro ao conectar!"))
		return
	}
	defer conn.Close()

	statement, err := conn.Prepare("INSERT INTO dbo.Usuarios (Nome, Email) OUTPUT INSERTED.Id VALUES (@nome, @email)")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Erro ao criar statement!"))
		return
	}
	defer statement.Close()

	insert, err := statement.Exec(sql.Named("nome", usuario.Nome), sql.Named("email", usuario.Email))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Erro ao inserir!"))
		return
	}

	_, err = insert.LastInsertId()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Erro ao inserir!"))
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	conn, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar"))
		return
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT * FROM dbo.Usuarios (NOLOCK)")
	if err != nil {
		w.Write([]byte("Erro ao buscar usuários"))
		return
	}
	defer rows.Close()

	var usuarios []usuario
	for rows.Next() {
		var usuario usuario

		if err := rows.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); err != nil {
			w.Write([]byte("Erro ao buscar usuários"))
			return
		}

		usuarios = append(usuarios, usuario)
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(usuarios); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	conn, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar"))
		return
	}
	defer conn.Close()

	row, err := conn.Query("SELECT * FROM dbo.Usuarios (NOLOCK) WHERE Id = @id", sql.Named("id", id))
	if err != nil {
		w.Write([]byte("Erro ao buscar usuários"))
		return
	}

	var usuario usuario
	if row.Next() {
		if err := row.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); err != nil {
			w.Write([]byte("Erro ao buscar usuários"))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(usuario); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
