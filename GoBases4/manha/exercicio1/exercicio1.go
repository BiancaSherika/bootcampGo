package exercicio1

import (
	"fmt"
	"net/http"
)

type meuErroCustomizado struct {
	status int
	msg    string
}

func (m *meuErroCustomizado) Error() string {
	return fmt.Sprintf("Salario: %d - %v", m.status, m.msg)
}

func verificarSalario(status int) (int, error) {
	if status < 15000 {
		return http.StatusInternalServerError, &meuErroCustomizado{
			status: status,
			msg:    "Erro: O salário digitado não está dentro do valor mínimo",
		}
	}
	return http.StatusOK, &meuErroCustomizado{
		status: status,
		msg:    "Necessário pagamento de imposto",
	}
}

func printarErro(erro error, status int) {
	if erro != nil {
		fmt.Println(erro)
	}
}

func Run() {
	fmt.Println("exercicio 1")
	status, erro := verificarSalario(10000)
	printarErro(erro, status)

	status, erro = verificarSalario(20000)
	printarErro(erro, status)

}
