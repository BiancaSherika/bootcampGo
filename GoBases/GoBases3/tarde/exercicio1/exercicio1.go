package exercicio1

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
	email     string
	senha     string
}

type Pessoa interface {
	mudarNome(nome string, sobrenome string)
	mudarIdade(idade int)
	mudarEmail(email string)
	mudarSenha(senha string)
}

func (p *pessoa) mudarNome(nome string, sobrenome string) {
	p.nome = nome
	p.sobrenome = sobrenome
}

func (p *pessoa) mudarIdade(idade int) {
	p.idade = idade
}

func (p *pessoa) mudarEmail(email string) {
	p.email = email
}

func (p *pessoa) mudarSenha(senha string) {
	p.senha = senha
}

func Run() {
	fmt.Println("Exercicio 1")
	usuario := pessoa{"Bianca", "Claro", 26, "bianca@email.com", "Senha321"}
	var p *pessoa = &usuario
	printarInfos(p)

	p.mudarNome("Amanda", "Lima")
	p.mudarIdade(30)
	p.mudarEmail("amanda@email.com")
	p.mudarSenha("SenhaTeste")
	printarInfos(p)
}

func printarInfos(usuario *pessoa) {
	fmt.Println("Inicio da execução")
	fmt.Println(usuario.nome, &usuario.nome)
	fmt.Println(usuario.sobrenome, &usuario.sobrenome)
	fmt.Println(usuario.idade, &usuario.idade)
	fmt.Println(usuario.email, &usuario.email)
	fmt.Println(usuario.senha, &usuario.senha)
}
